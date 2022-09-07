package io.alpha.support.core.eth.gas;
/*
 *  Copyright 2020-2021 this repo owner
 *
 *  Licensed under unknown License
 */

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

@Component
@Slf4j
public class GasProviderFactory {
    public enum GasUsageEnum {
        FAST, STANDARD
    }

    private Map<String, GenericGasProvider> gasProvider = new ConcurrentHashMap<>();

    @PostConstruct
    public void setup() {
        log.info("setup gas provider ");
        this.gasProvider.put(GasUsageEnum.STANDARD.name(), new GenericGasProvider(GasUsageEnum.STANDARD.name()));
        this.gasProvider.put(GasUsageEnum.FAST.name(), new GenericGasProvider(GasUsageEnum.FAST.name()));
    }

    public GenericGasProvider getGasProviderByGasUsage(String gasUsage) {
        return gasProvider.get(gasUsage.toUpperCase());
    }

    public GenericGasProvider getGasProvider() {
        return gasProvider.get(GasUsageEnum.STANDARD.toString());
    }
}
