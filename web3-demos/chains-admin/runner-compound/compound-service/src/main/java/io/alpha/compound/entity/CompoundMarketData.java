package io.alpha.compound.entity;

import lombok.Data;
import xyz.erupt.annotation.Erupt;
import xyz.erupt.annotation.EruptField;
import xyz.erupt.annotation.sub_field.View;
import xyz.erupt.upms.model.base.HyperModel;

import javax.persistence.*;

@Entity
@Table(name = "cmp_market_data")
@Erupt(
        name = "compound市场数据"
//        dataProxy = {CompoundMarketData.class}  //增删改查等行为数据代理
//        rowOperation = {                //自定义按钮
//                @RowOperation(title = "多行操作", code = "MULTI", icon = "fa fa-check-square",
//                        operationHandler = OperationHandlerImpl.class),
//                @RowOperation(title = "按钮操作", code = "BUTTON", tip = "不依赖任何数据即可执行", icon = "fa fa-google-wallet",
//                        operationHandler = OperationHandlerImpl.class, mode = RowOperation.Mode.BUTTON),
//        }
)
@Data
public class CompoundMarketData extends HyperModel {

    @EruptField(
            views = @View(title = "借款率", sortable = true)
    )
    private String borrowRate;
    //    @EruptField(
//            views = @View(title = "cash", sortable = true)
//    )
    private String cash;
    @EruptField(
            views = @View(title = "抵押率", sortable = true)
    )
    private String collateralFactor;
    //    @EruptField(
//            views = @View(title = "抵押率", sortable = true)
//    )
    private String interest_rate_model_address;
    @EruptField(
            views = @View(title = "市场名称", sortable = true)
    )
    private String name;
    //    @EruptField(
//            views = @View(title = "借款人数", sortable = true)
//    )
    private int numberOfBorrowers;
    //    @EruptField(
//            views = @View(title = "supply人数", sortable = true)
//    )
    private int numberOfSuppliers;
    //    @EruptField(
//            views = @View(title = "reverser总额", sortable = true)
//    )
    private String reserves;
    @EruptField(
            views = @View(title = "reverser_factor", sortable = true)
    )
    private String reserveFactor;
    //    @EruptField(
//            views = @View(title = "supply rate", sortable = true)
//    )
    private String supplyRate;
    @EruptField(
            views = @View(title = "symbol", sortable = true)
    )
    private String symbol;
    private String tokenAddress;
    private String totalBorrows;
    //    @EruptField(
//            views = @View(title = "总量", sortable = true)
//    )
    private String totalSupply;
    private String underlyingAddress;
    //    @EruptField(
//            views = @View(title = "底层资产名称", sortable = true)
//    )
    private String underlyingName;
    private String underlyingSymbol;
    @EruptField(
            views = @View(title = "底层资产价格/ETH", sortable = true)
    )
    private String underlyingPrice;

}
