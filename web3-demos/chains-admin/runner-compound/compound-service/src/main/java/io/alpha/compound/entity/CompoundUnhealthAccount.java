package io.alpha.compound.entity;

import lombok.Data;
import xyz.erupt.annotation.Erupt;
import xyz.erupt.annotation.EruptField;
import xyz.erupt.annotation.fun.DataProxy;
import xyz.erupt.annotation.sub_field.View;
import xyz.erupt.upms.model.base.HyperModel;

import javax.persistence.*;


@Table(name = "cmp_unhealth_accounts")
@Entity
@Erupt(
        name = "compound待清算数据",
        dataProxy = {CompoundUnhealthAccount.class}  //增删改查等行为数据代理
//        rowOperation = {                //自定义按钮
//                @RowOperation(title = "多行操作", code = "MULTI", icon = "fa fa-check-square",
//                        operationHandler = OperationHandlerImpl.class),
//                @RowOperation(title = "按钮操作", code = "BUTTON", tip = "不依赖任何数据即可执行", icon = "fa fa-google-wallet",
//                        operationHandler = OperationHandlerImpl.class, mode = RowOperation.Mode.BUTTON),
//        }
)
@Data
public class CompoundUnhealthAccount extends HyperModel implements DataProxy<CompoundUnhealthAccount> {

    @EruptField(
            views = @View(title = "借款人账户", sortable = true)
    )
    private String address;
    @EruptField(
            views = @View(title = "健康指数", sortable = true)
    )
    private String health;
    @EruptField(
            views = @View(title = "总计借款总额InEth", sortable = true)
    )
    private String totalBorrowValueInEth;
    @EruptField(
            views = @View(title = "总计抵押总额InEth", sortable = true)
    )
    private String totalCollateralValueInEth;

//
//    private String total_borrow_value_in_eth;
//    private String total_collateral_value_in_eth;
}
