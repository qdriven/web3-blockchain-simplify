package io.alpha.compound.entity;

import lombok.Data;
import xyz.erupt.annotation.Erupt;
import xyz.erupt.annotation.EruptField;
import xyz.erupt.annotation.fun.DataProxy;;
import xyz.erupt.annotation.sub_field.View;
import xyz.erupt.upms.model.base.HyperModel;

import javax.persistence.*;

@Entity
@Table(name = "compound_liq_tx")
@Erupt(
        name = "清算数据",
        dataProxy = {CompoundLiqTx.class}  //增删改查等行为数据代理
//        rowOperation = {                //自定义按钮
//                @RowOperation(title = "多行操作", code = "MULTI", icon = "fa fa-check-square",
//                        operationHandler = OperationHandlerImpl.class),
//                @RowOperation(title = "按钮操作", code = "BUTTON", tip = "不依赖任何数据即可执行", icon = "fa fa-google-wallet",
//                        operationHandler = OperationHandlerImpl.class, mode = RowOperation.Mode.BUTTON),
//        }
)
@Data
public class CompoundLiqTx extends HyperModel implements DataProxy<CompoundLiqTx> {

    private String borrowAddress;
    @EruptField(
            views = @View(title = "抵押品资产地址", sortable = true)
    )
    private String collateralAssetAddress;
    @EruptField(
            views = @View(title = "偿还数量", sortable = true)
    )
    private String repayAmount;
    @EruptField(
            views = @View(title = "偿还品资产地址", sortable = true)
    )
    private String repayAssetName;
    @EruptField(
            views = @View(title = "抵押品资产", sortable = true)
    )
    private String collateralAssetName;
    @EruptField(
            views = @View(title = "交易状态", sortable = true)
    )
    private String txStatus;
    @EruptField(
            views = @View(title = "交易hash", sortable = true)
    )
    private String transactionHash;

    @EruptField(
            views = @View(title = "已对冲", sortable = true)
    )
    private int hedgeStatus = 0;

    //    @EruptField(
//            views = @View(title = "交易备注", sortable = true)
//    )
    private String txMsg;

    private String repayAmountInEth;

    private String cexOrderId;
}
