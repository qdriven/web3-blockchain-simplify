package io.alpha.compound.service;

import cn.hutool.core.bean.BeanUtil;
import cn.hutool.core.bean.copier.CopyOptions;
import io.alpha.compound.intergation.CompoundApi;
import io.alpha.compound.intergation.request.AccountRequest;
import io.alpha.compound.intergation.response.*;
import io.alpha.compound.repository.CompoundLiqTxRepo;
import io.alpha.compound.repository.CompoundMarketDataRepo;
import io.alpha.compound.repository.CompoundUnhealthAccountRepo;
import io.alpha.compound.repository.CompoundUnhealthAssetRepo;
import io.alpha.configs.protocols.compound.CompoundConfigService;
import io.alpha.support.core.http.RetrofitWrapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.annotation.Order;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;
import retrofit2.Retrofit;
import io.alpha.compound.entity.CompoundLiqTx;
import io.alpha.compound.entity.CompoundMarketData;
import io.alpha.compound.entity.CompoundUnHealthAsset;
import io.alpha.compound.entity.CompoundUnhealthAccount;

import javax.annotation.PostConstruct;
import javax.transaction.Transactional;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.PriorityBlockingQueue;

@Service
@Slf4j
@Order(value = 2)
public class CompoundDataService {

    Retrofit retrofit;
    CompoundApi apiService;
    @Autowired
    private CompoundMarketDataRepo marketDataRepo;
    @Autowired
    CompoundConfigService configService;
    @Autowired
    CompoundUnhealthAccountRepo unhealthAccountRepo;
    @Autowired
    CompoundUnhealthAssetRepo unhealthAssetRepo;
    @Autowired
    CompoundLiqTxRepo liqTxRepo;

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
    @Async
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
}
