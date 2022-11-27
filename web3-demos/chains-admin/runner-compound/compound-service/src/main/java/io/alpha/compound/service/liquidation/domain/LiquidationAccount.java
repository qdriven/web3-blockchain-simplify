package io.alpha.compound.service.liquidation.domain;

import cn.hutool.core.collection.CollectionUtil;
import io.alpha.support.numbers.NumberUtil;
import lombok.Data;

import java.math.BigDecimal;
import java.util.Comparator;
import java.util.List;
import java.util.function.Predicate;
import java.util.stream.Collectors;

@Data
public class LiquidationAccount {
    private String borrowerAddress;
    private String health;
    private List<AssetForLiquidation> borrows;
    private List<AssetForLiquidation> suppliers;

    //1. 根据获得清算的抵押品计算清算数量
    //2. 根据获得清算

    //获取排好序的底层资产,按照ETH
    public List<AssetForLiquidation> getBorrowAssetsToLiq(List<String> liqAssetList) {

        List<AssetForLiquidation> filtered = borrows.stream().filter(new Predicate<AssetForLiquidation>() {
            @Override
            public boolean test(AssetForLiquidation borrower) {
                return CollectionUtil.contains(liqAssetList,
                        borrower.getAssetName().toUpperCase()) &&
                        NumberUtil.isGreater(borrower.getAssetValueInEth(), "0.6");
            }
        }).sorted(new Comparator<AssetForLiquidation>() {
            @Override
            public int compare(AssetForLiquidation o1, AssetForLiquidation o2) {
                return new BigDecimal(o2.getAssetValueInEth())
                        .compareTo(new BigDecimal(o1.getAssetValueInEth()));
            }
        }).collect(Collectors.toList());

        return filtered;
    }

    /**
     * TODO: 需要修正
     * 1. 计算抵押资产情况,设置那些有些资产清算时优先处理
     * 2. 优先处理的资产排序靠前
     *
     * @return
     */
    public List<AssetForLiquidation> getSupplyAssetsToLiq(List<String> liqAssetList) {

        List<AssetForLiquidation> filtered = suppliers.stream().filter(
                supplier -> CollectionUtil.contains(liqAssetList, supplier.getAssetName().toUpperCase())
                        && NumberUtil.isGreater(supplier.getAssetUnderlyingAmount(), "0.0")
        ).sorted(new Comparator<AssetForLiquidation>() {
            @Override
            public int compare(AssetForLiquidation o1, AssetForLiquidation o2) {
                return new BigDecimal(o2.getAssetValueInEth())
                        .compareTo(new BigDecimal(o1.getAssetValueInEth()));
            }
        }).collect(Collectors.toList());

        return filtered;
    }
}