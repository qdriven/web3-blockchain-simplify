package io.alpha.compound.service;

import cn.hutool.core.bean.BeanUtil;
import cn.hutool.core.bean.copier.CopyOptions;
import io.alpha.compound.CompoundConfigService;
import io.alpha.compound.contracts.CompoundContractFactory;
import io.alpha.compound.intergation.CompoundApi;
import io.alpha.compound.intergation.request.AccountRequest;
import io.alpha.compound.intergation.response.*;
import io.alpha.compound.repository.CompoundLiqTxRepo;
import io.alpha.compound.repository.CompoundMarketDataRepo;
import io.alpha.compound.repository.CompoundUnhealthAccountRepo;
import io.alpha.compound.repository.CompoundUnhealthAssetRepo;
import io.alpha.compound.CompoundLiqTx;
import io.alpha.compound.CompoundMarketData;
import io.alpha.compound.CompoundUnHealthAsset;
import io.alpha.compound.CompoundUnhealthAccount;
import io.alpha.support.core.http.RetrofitWrapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.annotation.Order;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;
import retrofit2.Retrofit;

import javax.annotation.PostConstruct;
import javax.transaction.Transactional;
import java.util.*;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.PriorityBlockingQueue;
import java.util.stream.Collectors;


@Service
@Slf4j
@Order(value = 2)
public class CompoundDataService {
    @Autowired
    CompoundConfigService configService;
    @Autowired
    private CompoundCalculator calculator;
    @Autowired
    private CompoundUnhealthAssetRepo unhealthAssetRepo;
    @Autowired
    private CompoundUnhealthAccountRepo unhealthAccountRepo;
    @Autowired
    private CompoundMarketDataRepo marketDataRepo;
    @Autowired
    CompoundLiqTxRepo liqTxRepo;
    @Autowired
    CompoundContractFactory factory;
    Retrofit retrofit;
    CompoundApi apiService;

    public Map<String, CompoundMarketData> compoundMarketData = new ConcurrentHashMap<>();
    private final Map<String, CmpAccount> unHealthyAccountMap = new ConcurrentHashMap<>();
    private final PriorityBlockingQueue<CmpAccount> queue = new PriorityBlockingQueue(50);

    /**
     * TODO:
     * 1. 开始启动之前确认钱包是否已经被Approve过
     * 2.REDEEM cToken
     * <p>
     * 1. 初始化Compound API监控
     * 2. 获取监控
     */
    @PostConstruct
    public void setupService() {
        this.retrofit = RetrofitWrapper.getRetrofit(configService.getCOMPOUND_API_URL());
        apiService = retrofit.create(CompoundApi.class);
        compoundMarketDataJob();
    }

    /**
     * @param request
     * @return
     */
    public AccountServiceResponse pullLiquidateInfo(AccountRequest request) {
        return RetrofitWrapper.sendAndReturnBody(apiService.accountService(
                request
        ));
    }

    /**
     * @return
     */
    public CTokenServiceResponse getCtokenServiceResponse() {
        return RetrofitWrapper.sendAndReturnBody(apiService.ctokenService());
    }

    /**
     * 保存Compound 各个Market数据
     */
    @Transactional(rollbackOn = Exception.class)
    public void compoundMarketDataJob() {
        CTokenServiceResponse response = getCtokenServiceResponse();
        List<CompoundMarketData> marketDataList = new ArrayList<>();
        for (CTokenResponse cTokenResponse : response.getCToken()) {
            CompoundMarketData data = new CompoundMarketData();
            BeanUtil.copyProperties(cTokenResponse, data, CopyOptions.create().ignoreError());
            data.setBorrowRate(cTokenResponse.getBorrow_rate().getValue());
            data.setCollateralFactor(cTokenResponse.getCollateral_factor().getValue());
            data.setReserveFactor(cTokenResponse.getReserve_factor().getValue());
            data.setReserves(cTokenResponse.getReserves().getValue());
            data.setSupplyRate(cTokenResponse.getSupply_rate().getValue());
            data.setBorrowRate(cTokenResponse.getBorrow_rate().getValue());
            data.setTotalBorrows(cTokenResponse.getTotal_borrows().getValue());
            data.setTotalSupply(cTokenResponse.getTotal_supply().getValue());
            data.setTokenAddress(cTokenResponse.getToken_address());
            data.setUnderlyingPrice(cTokenResponse.getUnderlying_price().getValue());
            data.setCash(cTokenResponse.getCash().getValue());
            marketDataList.add(data);
            compoundMarketData.put(cTokenResponse.getSymbol().toUpperCase(), data);
        }
        marketDataRepo.saveAll(marketDataList);
    }

    /**
     * 获取内存中的Compound市场信息
     *
     * @param symbol
     * @return
     */
    public CompoundMarketData getMarketDataBySymbol(String symbol) {
        return compoundMarketData.get(symbol.toUpperCase());
    }

    /**
     * compound监控清算数据的接口
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
            AccountServiceResponse response = pullLiquidateInfo(request);
            log.info("current compound page information,{}", response.getPagination_summary().toString());
            int accountSize = response.getAccounts().size();
            List<CmpAccount> unhealthyAccount = response.getAccounts().stream()
                    .filter(cmpAccount -> {
                        if (cmpAccount.getHealth() != null) {
                            unHealthyAccountMap.put(cmpAccount.getAddress(), cmpAccount);
                            queue.add(cmpAccount);
                            saveCompoundAssetInfo(cmpAccount);
                            return true;
                        }
                        return false;
                    }).collect(Collectors.toList());
            log.info(unhealthyAccount.toString());
            if (accountSize < configService.getCOMPOUND_API_PAGE_SIZE()) break;
        }
    }


    /**
     * 获取全部接近不安全账户数据接口
     *
     * @return
     */
    public Map<String, CompoundCalculator.LiqAccount> getUnHealthyAccounts() {
        Map<String, CompoundCalculator.LiqAccount> liqAccountMap = new HashMap<>();
        for (Map.Entry<String, CmpAccount> entry : unHealthyAccountMap.entrySet()) {
            liqAccountMap.put(entry.getKey(), calculator.getLiquidationAccountInfo(entry.getValue()));
        }
        return liqAccountMap;
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


    /**
     * 保存潜在不良资产的信息
     * TODO: 如何保存需要再考虑一下
     *
     * @param account
     */
    @Async
    @Transactional(rollbackOn = Exception.class)
    public void saveCompoundAssetInfo(CmpAccount account) {
        saveCompoundUnhealthAccount(account);
        saveCompoundUnhealthAsset(account);
    }

    @Async
    @Transactional(rollbackOn = Exception.class)
    void saveCompoundUnhealthAccount(CmpAccount account) {
        CompoundUnhealthAccount acc = new CompoundUnhealthAccount();
        acc.setAddress(account.getAddress());
        acc.setHealth(account.getHealth().getValue());
        acc.setTotalBorrowValueInEth(account.getTotal_borrow_value_in_eth().getValue());
        acc.setTotalCollateralValueInEth(account.getTotal_collateral_value_in_eth().getValue());
        unhealthAccountRepo.save(acc);
    }

    @Async
    @Transactional(rollbackOn = Exception.class)
    void saveCompoundUnhealthAsset(CmpAccount account) {
        List<CompoundUnHealthAsset> unhealthAssets = new ArrayList<>();
        for (AccountCToken token : account.getTokens()) {
            CompoundUnHealthAsset asset = new CompoundUnHealthAsset();
            asset.setAccountAddress(account.getAddress());
            asset.setSymbol(token.getSymbol());
            asset.setLifetime_borrow_interest_accrued(token.getLifetime_borrow_interest_accrued().getValue());
            asset.setLifetime_supply_interest_accrued(token.getLifetime_supply_interest_accrued().getValue());
            asset.setAddress(token.getAddress());
            asset.setBorrow_balance_underlying(token.getBorrow_balance_underlying().getValue());
            asset.setSupply_balance_underlying(token.getSupply_balance_underlying().getValue());
            unhealthAssets.add(asset);
        }

        unhealthAssetRepo.saveAll(unhealthAssets);
    }

    /**
     * 获取清算数据
     * TODO: 分页计算
     *
     * @return
     */
    public List<CompoundLiqTx> getAllCompoundLixTx() {
        return liqTxRepo.findAll();
    }

//    /**
//     * TODO: 获取所有相关币种的额度,移到WalletService中去
//     *
//     * @return
//     * @throws Exception
//     */
//    public Map getCompoundBalance() throws Exception {
//
//        String ownerAddress = provider.getCredentials().getAddress();
//        Map<String, String> walletBalanceResult = new HashMap<>();
//        BigInteger ethBalance = provider.getWeb3j().ethGetBalance(ownerAddress, DefaultBlockParameterName.LATEST)
//                .send().getBalance();
//        walletBalanceResult.put("ETH", ethBalance.toString());
//        for (Map.Entry<String, CompoundContractConfig.CTokenInfo> entry : configService.getConfig().getAllContracts().entrySet()) {
//            if (entry.getKey().equalsIgnoreCase("CETH")) {
//                CETHContract cethContract = factory.createETHContract(entry.getValue().getContract_address(), provider);
//                BigInteger cethBalance = cethContract.balanceOf(ownerAddress).send();
//                walletBalanceResult.put("CETH", cethBalance.toString());
//            } else {
//                CERC20Contract erc20 = factory.createErc20Contract(entry.getValue().getContract_address(), provider);
//                BigInteger balance = erc20.balanceOf(ownerAddress).send();
//                walletBalanceResult.put(entry.getKey(), balance.toString());
//            }
//        }
//        return walletBalanceResult;
//    }

}
