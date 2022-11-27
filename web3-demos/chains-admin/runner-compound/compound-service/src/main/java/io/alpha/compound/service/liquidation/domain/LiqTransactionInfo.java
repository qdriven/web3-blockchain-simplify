package io.alpha.compound.service.liquidation.domain;

import lombok.Data;

@Data
public class LiqTransactionInfo {
    private String borrowAddress;
    private String repayAmount;
    private String repayAmountInEth;
    private String repayAssetName;
    private String repayAssetAddress;
    private String collateralAssetAddress;
    private String collateralAssetName;
    private String comments;
    private boolean liquidatable=false;
}