package io.alpha.compound.service;

import io.alpha.compound.CompoundConfigService;
import io.alpha.compound.intergation.response.CmpAccount;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import java.math.BigDecimal;
import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

@Service
@Slf4j
public class CompoundLiquidationService {

    @Autowired
    CompoundConfigService configService;
    @Autowired
    CompoundCalculator calculator;
    @Autowired
    CompoundDataService service;
    @Autowired
    private CompoundLiquidator contractInvoker;

    ExecutorService executor;

    private Map<String, CmpAccount> unLiqAccount = new ConcurrentHashMap<>();
    private List<String> liqBlacklist = Arrays.asList("0xa4c4b50f0bd21efa6ac571e76091e677c65d72ca");
    private Map<String, Object> liquidatedAccount = new ConcurrentHashMap<>();

    /**
     * Default not to start liq job
     */
    @PostConstruct
    public void setupLiqTask() {
        executor = Executors.newFixedThreadPool(Runtime.getRuntime().availableProcessors() * 2);
        for (int i = 0; i < configService.getCOMPOUND_LIQ_THREADS(); i++) { //TODO: 线程数量可以修改
            executor.submit(this::liqTask);
        }
    }

    /**
     * 不进行清算
     */
    public void shutdownLiqMonitorJob() {
        this.executor.shutdown();
    }

    /**
     * 重新开始清算
     */
    public void startLiqJob() {
        this.executor = Executors.newFixedThreadPool(Runtime.getRuntime().availableProcessors() * 2);
        for (int i = 0; i < configService.getCOMPOUND_LIQ_THREADS(); i++) { //TODO: 线程数量可以修改
            executor.submit(this::liqTask);
        }
    }

    //TODO: move to JOB Implementation
    private void liqTask() {
        while (true) { // job status to a global variable
            CmpAccount account = service.takeUnHealthyAddress();
            if (account == null) continue;
            if (new BigDecimal(account.getHealth().getValue())
                    .compareTo(new BigDecimal(configService.getDEFAULT_HEALTH_THRESHOLD())) < 0) {
                if (liqBlacklist.contains(account.getAddress())) continue;
                liqUnhealthAccount(account);
            }
            log.info("account is {},health is {},account detail is {}", account.getAddress(), account.getHealth().getValue(), account);
        }
    }

    /**
     * 清算发起
     *
     * @param account
     */
    private void liqUnhealthAccount(CmpAccount account) {
        //1. 是否已经在黑名单/清算列表中
        //2. 是否符合清算盈利目标, 确定使用那个手续费
        //3. 确定清算金额，使用哪个钱包和手续费策略
        //计算可以清算的数量，尽量一笔清算
        if (unLiqAccount.getOrDefault(account.getAddress(), null) != null) return;
        if (liquidatedAccount.getOrDefault(account.getAddress(), null) != null) return;

        CompoundCalculator.LiqTransactionInfo txInfo = calculator.calculateLiqTx(account);
        if (txInfo == null) {
            unLiqAccount.put(account.getAddress(), account);
            return;
        } else { //防止re-enter 造成代码错误
            if (txInfo.getRepayAssetName().equalsIgnoreCase(txInfo.getCollateralAssetName()) &&
                    !txInfo.getRepayAssetName().equalsIgnoreCase("cDAI")) { //只有dai才能清算抵押dai/同时借dai的数据

                //ceth清算不能自己清算
                unLiqAccount.put(account.getAddress(), account);
                return;
            }
        }

        unLiqAccount.put(account.getAddress(), account);
        if (configService.isENABLE_LIQUIDATION()) {
            contractInvoker.liquidateBorrow(txInfo); //liquidateBorrow -> liquidate Borrow in local first
        }
        liquidatedAccount.put(txInfo.getBorrowAddress(), txInfo);
    }
}
