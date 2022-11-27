package io.alpha.support.core.eth.gas;

import lombok.Data;

@Data
public class EthGasPrice {
    private Integer safeLow;
    private Integer standard;
    private Integer fast;
    private Integer fastest;
}