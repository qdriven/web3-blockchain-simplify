package io.alpha.support.core.eth.gas;

import lombok.Data;

import java.util.Map;

@Data
//todo: gas station gas
public class EthStationGasPrice {
    private float fast;
    private float fastest;
    private float safeLow;
    private float average;
    private float block_time;
    private float speed;
    private Long blockNum;
    private float safeLowWait;
    private long avgWait;
    private long fastestWait;
    private Map<String,Float> gasPriceRange;
}
