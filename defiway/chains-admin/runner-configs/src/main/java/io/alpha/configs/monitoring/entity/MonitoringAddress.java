package io.alpha.configs.monitoring.entity;

import xyz.erupt.upms.model.base.HyperModel;

import javax.persistence.Entity;
import javax.persistence.Table;

@Entity
@Table(name = "monitoring_addresses")
public class MonitoringAddress extends HyperModel {

    private String address;
    private String protocolName;
    private String description;
    private String action;
    private String expression;
}
