package io.qdriven.chains.config.tokens.entity;

import lombok.Data;
import xyz.erupt.annotation.EruptField;
import xyz.erupt.annotation.sub_field.Edit;
import xyz.erupt.annotation.sub_field.View;

import javax.persistence.Entity;
import javax.persistence.Table;

import xyz.erupt.annotation.Erupt;
import xyz.erupt.upms.model.base.HyperModel;

@Erupt(name = "tokens")
@Entity
@Table(name = "tokens")
@Data
public class BlockchainToken extends HyperModel {
    @EruptField(
            views = @View(title = "symbol"),
            edit = @Edit(title = "symbol") //相当于 type = EditType.INPUT
    )
    private String symbol;
    @EruptField(
            views = @View(title = "decimal"),
            edit = @Edit(title = "decimal") //相当于 type = EditType.INPUT
    )
    private String decimalValue;
    @EruptField(
            views = @View(title = "名称"),
            edit = @Edit(title = "name") //相当于 type = EditType.INPUT
    )
    private String name;
    @EruptField(
            views = @View(title = "协议名"),
            edit = @Edit(title = "协议名") //相当于 type = EditType.INPUT
    )
    private String protocolName;
    @EruptField(
            views = @View(title = "链名"),
            edit = @Edit(title = "链名") //相当于 type = EditType.INPUT
    )
    private String chainName;
    @EruptField(
            views = @View(title = "链种类"),
            edit = @Edit(title = "链种类") //相当于 type = EditType.INPUT
    )
    private String chainType;
    @EruptField(
            views = @View(title = "token地址"),
            edit = @Edit(title = "token地址") //相当于 type = EditType.INPUT
    )
    private String address;
    @EruptField(
            views = @View(title = "abi地址"),
            edit = @Edit(title = "abi地址") //相当于 type = EditType.INPUT
    )
    private String abiUrl;

    @EruptField(
            views = @View(title = "描述"),
            edit = @Edit(title = "描述") //相当于 type = EditType.INPUT
    )
    private String description;
    @EruptField(
            views = @View(title = "启用"),
            edit = @Edit(title = "启动") //相当于 type = EditType.INPUT
    )
    private boolean enable;

}
