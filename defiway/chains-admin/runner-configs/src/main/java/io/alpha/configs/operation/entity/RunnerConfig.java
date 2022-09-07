package io.alpha.configs.operation.entity;

import lombok.Data;
import xyz.erupt.annotation.Erupt;
import xyz.erupt.annotation.EruptField;
import xyz.erupt.annotation.sub_field.Edit;
import xyz.erupt.annotation.sub_field.EditType;
import xyz.erupt.annotation.sub_field.View;
import xyz.erupt.upms.model.base.HyperModel;

import javax.persistence.Entity;
import javax.persistence.Table;

@Data
@Entity
@Table(name = "runner_configs")
@Erupt(name = "运行参数")
public class RunnerConfig extends HyperModel {
    @EruptField(
            views = @View(title = "协议"),
            edit = @Edit(title = "协议") //相当于 type = EditType.INPUT
    )
    private String protocolName;
    @EruptField(
            views = @View(title = "参数名"),
            edit = @Edit(title = "参数名") //相当于 type = EditType.INPUT
    )
    private String parameterName;
    @EruptField(
            views = @View(title = "参数值"),
            edit = @Edit(title = "参数值") //相当于 type = EditType.INPUT
    )
    private String parameterValue;
    @EruptField(
            views = @View(title = "是否可用"),
            edit = @Edit(title = "是否可用", type = EditType.BOOLEAN) //相当于 type = EditType.INPUT
    )
    //TODO: Adding audit history
    private boolean active;
    @EruptField(
            views = @View(title = "参数说明"),
            edit = @Edit(title = "参数说明") //相当于 type = EditType.INPUT
    )
    private String description;
}
