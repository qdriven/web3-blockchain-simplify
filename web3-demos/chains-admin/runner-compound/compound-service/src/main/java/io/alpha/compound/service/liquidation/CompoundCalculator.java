package io.alpha.compound.service.liquidation;

import io.alpha.compound.intergation.response.AccountCToken;
import io.alpha.compound.intergation.response.CmpAccount;
import io.alpha.compound.service.CompoundDataService;
import io.alpha.compound.service.liquidation.domain.AssetForLiquidation;
import io.alpha.compound.service.liquidation.domain.LiqTransactionInfo;
import io.alpha.compound.service.liquidation.domain.LiquidationAccount;
import io.alpha.configs.protocols.compound.CompoundConfigService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.ArrayList;
import java.util.List;

import static io.alpha.support.numbers.NumberUtil.multipleBigDecimalString;

@Component
@Slf4j
public class CompoundCalculator {
    private final CompoundConfigService config;
    private final CompoundDataService dataService;

    public CompoundCalculator(CompoundConfigService config, CompoundDataService dataService) {
        this.config = config;
        this.dataService = dataService;
    }
//    private String LIQ_CLOSE_FACTOR = "0.48";
//    private String Collateral_Factor = "0.6";

    /**
     * 1. 获取borrow balance数据，作为待清算的数量
     * 2. 获取supply balance数据，作为可以获得的抵押品数量
     *
     * @param account
     * @return
     */
    public LiquidationAccount getLiquidationAccountInfo(CmpAccount account) {
        List<AssetForLiquidation> supplies = new ArrayList<>();
        List<AssetForLiquidation> borrows = new ArrayList<>();

        log.info("account information: {} ", account);

        String borrowerAddr = account.getAddress();
        for (AccountCToken token : account.getTokens()) {
            AssetForLiquidation suppliedAccount = convertToLiquidationAccount(token, true);
            if (!suppliedAccount.getAssetUnderlyingAmount().equalsIgnoreCase("0")) {
                supplies.add(suppliedAccount);
            }
            AssetForLiquidation borrowAccount = convertToLiquidationAccount(token, false);
            if (!borrowAccount.getAssetUnderlyingAmount().equalsIgnoreCase("0")) {
                borrows.add(borrowAccount);
            }
        }
        //获取borrower,获取可清算的金额和资产名称，获取抵押品的金额和资产
        log.info(borrows.toString());
        log.info(supplies.toString());
        LiquidationAccount liqAccount = new LiquidationAccount();
        liqAccount.setBorrowerAddress(borrowerAddr);
        liqAccount.setHealth(account.getHealth().getValue());
        liqAccount.setBorrows(borrows);
        liqAccount.setSuppliers(supplies);
        return liqAccount;
    }

    private AssetForLiquidation convertToLiquidationAccount(AccountCToken token, boolean isSupply) {
        AssetForLiquidation account = new AssetForLiquidation();
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

    public LiqTransactionInfo calculateLiqTxInfo(LiquidationAccount liqAccount) {
        List<String> allowedAssetList = config.getLiqAllowedAssets();
        if (liqAccount.getBorrowAssetsToLiq(allowedAssetList).size() < 1) return null;

        AssetForLiquidation borrowAccount = liqAccount.getBorrowAssetsToLiq(allowedAssetList).get(0);
        AssetForLiquidation supplierAccount = liqAccount.getSupplyAssetsToLiq(allowedAssetList).get(0);

        LiqTransactionInfo tx = new LiqTransactionInfo();
        tx.setCollateralAssetName(supplierAccount.getAssetName());
        tx.setCollateralAssetAddress(supplierAccount.getAssetAddress());
        tx.setBorrowAddress(liqAccount.getBorrowerAddress());
        tx.setRepayAssetName(borrowAccount.getAssetName());
        tx.setRepayAssetAddress(borrowAccount.getAssetAddress());

        //检查可清算资产和一种抵押资产的金额对比
        BigDecimal availableForLiq = new BigDecimal(borrowAccount.getAssetValueInEth())
                .multiply(new BigDecimal(config.getLIQ_CLOSE_FACTOR()));
        log.info("可清算borrow资产为{},价值为{}", borrowAccount.getAssetName(), borrowAccount.getAssetUnderlyingAmount());
        BigDecimal collateralAmount = new BigDecimal(supplierAccount.getAssetValueInEth())
                .multiply(new BigDecimal(config.getCollateral_Factor()));
        log.info("可清算抵押资产为{},价值使用Callateral factor {},价值为{}", supplierAccount.getAssetName(),
                config.getCollateral_Factor(), supplierAccount.getAssetUnderlyingAmount().toString());
        if (availableForLiq.compareTo(collateralAmount) > 0) {
            BigDecimal repayAmount = collateralAmount.divide(new BigDecimal(
                    dataService.getMarketDataBySymbol(borrowAccount.getAssetName()).getUnderlyingPrice()), RoundingMode.DOWN);
            tx.setRepayAmount(repayAmount.toString());
        } else {
            BigDecimal repayAmount = new BigDecimal(borrowAccount.getAssetUnderlyingAmount())
                    .multiply(new BigDecimal(config.getLIQ_CLOSE_FACTOR()));
            tx.setRepayAmount(repayAmount.toString()); //TODO: recalculate the repay amount
        }
        tx.setRepayAmountInEth(borrowAccount.getAssetValueInEth());
        return tx; //TODO: 清理已经清算过的账户
    }

    public LiqTransactionInfo calculateLiqTx(CmpAccount account) {
        LiquidationAccount liqAccount = getLiquidationAccountInfo(account);
        return calculateLiqTxInfo(liqAccount);
    }
}