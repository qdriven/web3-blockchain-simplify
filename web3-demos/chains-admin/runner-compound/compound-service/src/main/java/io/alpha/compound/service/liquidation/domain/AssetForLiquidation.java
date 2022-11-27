package io.alpha.compound.service.liquidation.domain;

import lombok.Data;

@Data
public class AssetForLiquidation {
    private String assetAddress;
    private String assetName;
    private String assetUnderlyingAmount;
    private String assetValueInEth;
}
