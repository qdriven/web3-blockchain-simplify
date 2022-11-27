package io.alpha.compound.service;


import cn.hutool.core.collection.CollectionUtil;

import io.alpha.compound.intergation.response.AccountCToken;
import io.alpha.compound.intergation.response.CmpAccount;
import io.alpha.configs.protocols.compound.CompoundConfigService;
import io.alpha.support.numbers.NumberUtil;
import lombok.Data;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.function.Predicate;
import java.util.stream.Collectors;

import static io.alpha.support.numbers.NumberUtil.multipleBigDecimalString;

/**
 * TODO:
 * 1. 获取钱包余额信息，放入缓存
 * 2. 获取待清算账户信息: 合约调用
 * 3. 新钱包初始化操作-approve记录
 * 4. 手续费研究
 */
@Component
@Slf4j
public class CompoundCalculator {
    @Autowired
    private CompoundConfigService config;
    @Autowired
    private CompoundDataService dataService;
//    private String LIQ_CLOSE_FACTOR = "0.48";
//    private String Collateral_Factor = "0.6";

    /**
     * 1. 获取borrow balance数据，作为待清算的数量
     * 2. 获取supply balance数据，作为可以获得的抵押品数量
     *
     * @param account
     * @return
     */
    public LiqAccount getLiquidationAccountInfo(CmpAccount account) {
        List<AccountForLiquidation> supplies = new ArrayList<>();
        List<AccountForLiquidation> borrows = new ArrayList<>();

        log.info("account information: {} ", account);

        String borrowerAddr = account.getAddress();
        for (AccountCToken token : account.getTokens()) {
            AccountForLiquidation suppliedAccount = convertToLiquidationAccount(token, true);
            if (!suppliedAccount.getAssetUnderlyingAmount().equalsIgnoreCase("0")) {
                supplies.add(suppliedAccount);
            }
            AccountForLiquidation borrowAccount = convertToLiquidationAccount(token, false);
            if (!borrowAccount.getAssetUnderlyingAmount().equalsIgnoreCase("0")) {
                borrows.add(borrowAccount);
            }
        }
        //获取borrower,获取可清算的金额和资产名称，获取抵押品的金额和资产
        log.info(borrows.toString());
        log.info(supplies.toString());
        LiqAccount liqAccount = new LiqAccount();
        liqAccount.setBorrowerAddress(borrowerAddr);
        liqAccount.setHealth(account.getHealth().getValue());
        liqAccount.setBorrows(borrows);
        liqAccount.setSuppliers(supplies);
        return liqAccount;
    }

    private AccountForLiquidation convertToLiquidationAccount(AccountCToken token, boolean isSupply) {
        AccountForLiquidation account = new AccountForLiquidation();
        account.setAssetAddress(token.getAddress());
        account.setAssetName(config.getConfig().getAssetInfoByContractAddress(token.getAddress()).getCtoken_symbol());
        if (isSupply) {
            account.setAssetUnderlyingAmount(token.getSupply_balance_underlying().getValue());
        } else {
            account.setAssetUnderlyingAmount(token.getBorrow_balance_underlying().getValue());
        }
        String sUnderlyingAmountInEth = multipleBigDecimalString(dataService.getMarketDataBySymbol(account.getAssetName())
                        .getUnderlyingPrice(), //ETH Price is 1
                account.getAssetUnderlyingAmount());
        account.setAssetValueInEth(sUnderlyingAmountInEth);
        return account;
    }

    public LiqTransactionInfo calculateLiqTxInfo(LiqAccount liqAccount) {
        if (liqAccount.getBorrowAssetsToLiq(config.getLIQ_ASSET_LIST()).size() < 1) return null;

        AccountForLiquidation borrowAccount = liqAccount.getBorrowAssetsToLiq(config.getLIQ_ASSET_LIST()).get(0);
        AccountForLiquidation supplierAccount = liqAccount.getSupplyAssetsToLiq(config.getLIQ_ASSET_LIST()).get(0);

        LiqTransactionInfo tx = new LiqTransactionInfo();
        tx.setCollateralAssetName(supplierAccount.getAssetName());
        tx.setCollateralAssetAddress(supplierAccount.getAssetAddress());
        tx.setBorrowAddress(liqAccount.getBorrowerAddress());
        tx.setRepayAssetName(borrowAccount.assetName);
        tx.setRepayAssetAddress(borrowAccount.assetAddress);

        //检查可清算资产和一种抵押资产的金额对比
        BigDecimal availableForLiq = new BigDecimal(borrowAccount.assetValueInEth).multiply(new BigDecimal(LIQ_CLOSE_FACTOR));
        log.info("可清算borrow资产为{},价值为{}", borrowAccount.assetName, borrowAccount.getAssetUnderlyingAmount());
        BigDecimal collateralAmount = new BigDecimal(supplierAccount.assetValueInEth).multiply(new BigDecimal(Collateral_Factor));
        log.info("可清算抵押资产为{},价值使用Callateral factor {},价值为{}", supplierAccount.assetName, Collateral_Factor, supplierAccount.getAssetUnderlyingAmount().toString());
        if (availableForLiq.compareTo(collateralAmount) > 0) {
            BigDecimal repayAmount = collateralAmount.divide(new BigDecimal(
                    dataService.getMarketDataBySymbol(borrowAccount.getAssetName()).getUnderlyingPrice()), RoundingMode.DOWN);
            tx.setRepayAmount(repayAmount.toString());
        } else {
            BigDecimal repayAmount = new BigDecimal(borrowAccount.assetUnderlyingAmount).multiply(new BigDecimal(LIQ_CLOSE_FACTOR));
            tx.setRepayAmount(repayAmount.toString()); //TODO: recalculate the repay amount
        }
        tx.setRepayAmountInEth(borrowAccount.getAssetValueInEth());
        return tx; //TODO: 清理已经清算过的账户
    }

    public LiqTransactionInfo calculateLiqTx(CmpAccount account) {
        LiqAccount liqAccount = getLiquidationAccountInfo(account);
        return calculateLiqTxInfo(liqAccount);
    }

    @Data
    public static class LiqTransactionInfo {
        private String borrowAddress;
        private String repayAmount;
        private String repayAmountInEth;
        private String repayAssetName;
        private String repayAssetAddress;
        private String collateralAssetAddress;
        private String collateralAssetName;
    }

    @Data
    public static class AccountForLiquidation {
        private String assetAddress;
        private String assetName;
        private String assetUnderlyingAmount;
        private String assetValueInEth;
    }


    @Data
    public static class LiqAccount {
        private String borrowerAddress;
        private String health;
        private List<AccountForLiquidation> borrows;
        private List<AccountForLiquidation> suppliers;

        //1. 根据获得清算的抵押品计算清算数量
        //2. 根据获得清算

        //获取排好序的底层资产,按照ETH
        public List<AccountForLiquidation> getBorrowAssetsToLiq(List<String> liqAssetList) {

            List<AccountForLiquidation> filtered = borrows.stream().filter(new Predicate<AccountForLiquidation>() {
                @Override
                public boolean test(AccountForLiquidation borrower) {
                    return CollectionUtil.contains(liqAssetList,
                            borrower.getAssetName().toUpperCase()) &&
                            NumberUtil.isGreater(borrower.getAssetValueInEth(), "0.6");
                }
            }).sorted(new Comparator<AccountForLiquidation>() {
                @Override
                public int compare(AccountForLiquidation o1, AccountForLiquidation o2) {
                    return new BigDecimal(o2.assetValueInEth)
                            .compareTo(new BigDecimal(o1.assetValueInEth));
                }
            }).collect(Collectors.toList());

            return filtered;
        }

        /**
         * TODO: 需要修正
         * 1. 计算抵押资产情况,设置那些有些资产清算时优先处理
         * 2.
         *
         * @return
         */
        public List<AccountForLiquidation> getSupplyAssetsToLiq(List<String> liqAssetList) {

            List<AccountForLiquidation> filtered = suppliers.stream().filter(
                    supplier -> CollectionUtil.contains(liqAssetList, supplier.getAssetName().toUpperCase())
                            && NumberUtil.isGreater(supplier.getAssetUnderlyingAmount(), "0.0")
            ).sorted(new Comparator<AccountForLiquidation>() {
                @Override
                public int compare(AccountForLiquidation o1, AccountForLiquidation o2) {
                    return new BigDecimal(o2.assetValueInEth)
                            .compareTo(new BigDecimal(o1.assetValueInEth));
                }
            }).collect(Collectors.toList());

            return filtered;
        }
    }
}
