package io.alpha.compound.service.liquidation;

import cn.hutool.core.util.StrUtil;
import io.alpha.compound.intergation.request.AccountRequest;
import io.alpha.compound.intergation.response.AccountServiceResponse;
import io.alpha.compound.intergation.response.CmpAccount;
import io.alpha.compound.intergation.response.CompoundValue;
import io.alpha.compound.service.CompoundDataService;
import io.alpha.compound.service.liquidation.domain.LiqTransactionInfo;
import io.alpha.compound.service.liquidation.response.UnHealthyAccountResponse;
import io.alpha.configs.protocols.compound.CompoundConfigService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.PriorityBlockingQueue;

@Service
@Slf4j
public class CompoundLiquidateService implements LiquidateService {
    private List<String> liqBlacklist = Arrays.asList("0xa4c4b50f0bd21efa6ac571e76091e677c65d72ca");
    private Map<String, CmpAccount> falseLiqAccount = new ConcurrentHashMap<>();
    private Map<String, Object> liquidatedAccount = new ConcurrentHashMap<>();
    private final Map<String, CmpAccount> unHealthyAccountMap = new ConcurrentHashMap<>();
    private final PriorityBlockingQueue<CmpAccount> queue = new PriorityBlockingQueue(50);
    @Autowired
    CompoundConfigService configService;
    @Autowired
    CompoundDataService dataService;
    @Autowired
    CompoundCalculator calculator;
    @Autowired
    CompoundLiquidator liquidator;

    ExecutorService executor;

    /**
     * compound监控清算数据的接口
     * TODO:
     * 1. move out to work as event
     */
    public void compoundLiquidationJobTask() {
        //处理多页情况
        log.info("current compound configuration is {}", configService.getAllConfig().toString());
        for (int i = 0; i < configService.getCOMPOUND_API_PAGE_NUM(); i++) {
            AccountRequest request = AccountRequest.builder()
                    .max_health(CompoundValue.fromValue(configService.getDEFAULT_HEALTH_THRESHOLD()))
                    .min_borrow_value_in_eth(
                            CompoundValue.fromValue(configService.getMIN_BORROW_VALUE_IN_ETH()))
                    .page_size(configService.getCOMPOUND_API_PAGE_SIZE()).page_number(i)
                    .build();
            AccountServiceResponse response = dataService.pullLiquidateInfo(request);
            log.info("current compound page information,{}", response.getPagination_summary().toString());
            int accountSize = response.getAccounts().size();
            response.getAccounts().stream()
                    .filter(cmpAccount -> {
                        if (cmpAccount.getHealth() != null) {
                            unHealthyAccountMap.put(cmpAccount.getAddress(), cmpAccount);
                            queue.add(cmpAccount);
                            dataService.saveCompoundAssetInfo(cmpAccount);
                            return true;
                        }
                        return false;
                    });
            if (accountSize < configService.getCOMPOUND_API_PAGE_SIZE()) break;
        }
    }

    @PostConstruct
    public void setupLiquidationJob() {
        executor = Executors.newFixedThreadPool(Runtime.getRuntime().availableProcessors() * 2);
        for (int i = 0; i < configService.getCOMPOUND_LIQ_THREADS(); i++) { //TODO: 线程数量可以修改
            executor.submit(this::liqTask);
        }
    }

    /**
     * 重新开始清算
     */
    public void restartJob() {
        this.executor.shutdownNow();
        this.executor = Executors.newFixedThreadPool(Runtime.getRuntime().availableProcessors() * 2);
        for (int i = 0; i < configService.getCOMPOUND_LIQ_THREADS(); i++) { //TODO: 线程数量可以修改
            executor.submit(this::liqTask);
        }
    }

    private void liqTask() {
        while (true) { // job status to a global variable
            CmpAccount account = takeUnHealthyAddress();
            log.info("account is {},health is {},account detail is {}", account.getAddress(), account.getHealth().getValue(), account);
            liquidate(account);
        }
    }

    @Override
    public boolean preLiquidationCheck(CmpAccount account) {
        //清算黑名单
        if (liqBlacklist.contains(account.getAddress())) return false;
        //不能清算的账户
        if (falseLiqAccount.getOrDefault(account.getAddress(), null) != null) return false;
        //已经清算清算账户
        if (liquidatedAccount.getOrDefault(account.getAddress(), null) != null) return false;
        //TODO: 检查实际的account liquidity
        if (new BigDecimal(account.getHealth().getValue())
                .compareTo(new BigDecimal(configService.getDEFAULT_HEALTH_THRESHOLD())) < 0) {
            return true;
        }
        return false;
    }

    @Override
    public LiqTransactionInfo calculateLiquidationAccount(CmpAccount account) {
        if (preLiquidationCheck(account)) {
            falseLiqAccount.put(account.getAddress(), account);
            return null;
        }
        LiqTransactionInfo txInfo = calculator.calculateLiqTx(account);
        if (txInfo == null) {
            falseLiqAccount.put(account.getAddress(), account);
        } else { //防止re-enter 造成代码错误
            if (txInfo.getRepayAssetName().equalsIgnoreCase(txInfo.getCollateralAssetName()) &&
                    !txInfo.getRepayAssetName().equalsIgnoreCase("cDAI")) { //只有dai才能清算抵押dai/同时借dai的数据
                //ceth清算不能自己清算
                txInfo.setComments("只有dai才能清算抵押dai/同时借dai的数据");
                falseLiqAccount.put(account.getAddress(), account);
            }
        }
        txInfo.setLiquidatable(true);
        return txInfo;
    }

    @Override
    public void liquidate(CmpAccount account) {

        LiqTransactionInfo txInfo = calculateLiquidationAccount(account);
        if (txInfo == null) {
            falseLiqAccount.put(account.getAddress(), account);
            return;
        }
        if (txInfo.isLiquidatable()) {
            if (configService.isENABLE_LIQUIDATION()) {
                liquidatedAccount.put(txInfo.getBorrowAddress(), txInfo);
                liquidator.liquidateBorrow(txInfo);
            }
        }
    }

    @Override
    public void afterLiquidation(CmpAccount account) {
        log.info("account {} is liquidated", account.getAddress());
    }

    /**
     * 获取不安全账户信息
     *
     * @return
     */
    public CmpAccount takeUnHealthyAddress() {
        try {
            return queue.take();
        } catch (InterruptedException e) {
            log.info("获取不监控账户失败,", e);
        }
        return null;
    }

    //GetAllLiquidationData:Refresh
    public List<UnHealthyAccountResponse> getUnhealthyAccounts() {
        List<UnHealthyAccountResponse> responses = new ArrayList<>();
        for (Map.Entry<String, CmpAccount> entry : unHealthyAccountMap.entrySet()) {
            UnHealthyAccountResponse response = new UnHealthyAccountResponse();
            LiqTransactionInfo info = calculateLiquidationAccount(entry.getValue());
            if (falseLiqAccount.get(entry.getKey()) != null) {
                response.setComments("不建议清算,可能不会盈利");
                if (info != null && StrUtil.isNotEmpty(info.getComments())) {
                    response.setComments(info.getComments());
                }
                response.setAdviceForLiquidating(false);
            }
            if (liquidatedAccount.get(entry.getKey()) != null) {
                response.setComments("已经清算");
                response.setAdviceForLiquidating(false);
            }

            response.setOriginAccount(entry.getValue());
            response.setAdvisedLiquidationInfo(info);
            responses.add(response);
        }
        return responses;
    }
}
