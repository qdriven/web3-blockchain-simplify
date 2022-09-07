package io.alpha.configs.tokens.entity;

import lombok.Data;
import xyz.erupt.annotation.EruptField;
import xyz.erupt.annotation.sub_field.Edit;

import javax.persistence.Entity;
import javax.persistence.Table;

import xyz.erupt.annotation.Erupt;
import xyz.erupt.upms.model.base.HyperModel;

@Erupt(name = "智能合约")
@Entity
@Table(name = "smart_contracts")
@Data
public class SmartContract extends HyperModel {
    @EruptField(
            edit = @Edit(title = "合约名称") //相当于 type = EditType.INPUT
    )
    private String name;
    @EruptField(
            edit = @Edit(title = "协议") //相当于 type = EditType.INPUT
    )
    private String protocolName;
    @EruptField(
            edit = @Edit(title = "链名") //相当于 type = EditType.INPUT
    )
    private String chainName;
    @EruptField(
            edit = @Edit(title = "链种类") //相当于 type = EditType.INPUT
    )
    private String chainType;
    @EruptField(
            edit = @Edit(title = "token地址") //相当于 type = EditType.INPUT
    )
    private String address;
}
