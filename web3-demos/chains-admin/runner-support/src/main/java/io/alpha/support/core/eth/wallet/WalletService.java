package io.alpha.support.core.eth.wallet;
/*
 *  Copyright 2020-2021 this repo owner
 *
 *  Licensed under unknown License
 */

import com.alibaba.excel.EasyExcel;
import com.alibaba.excel.event.SyncReadListener;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;
import org.web3j.crypto.Credentials;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

@Slf4j
@Component
public class WalletService {

    public static Map<String, Credentials> credentialsMap = new ConcurrentHashMap<>();

    public Credentials getWalletByUsage(String walletUsage) {
        return credentialsMap.get(walletUsage.toUpperCase());
    }

    public void loadWallets(String excelFile) {
        SyncReadListener listener = new SyncReadListener();
        EasyExcel.read(excelFile, WalletEntity.class, listener).sheet().doRead();
        for (Object o : listener.getList()) {
            WalletEntity e = (WalletEntity) o;
            if (e != null) {
                Credentials credentials = WalletProvider.getWalletProvider(e.getJsonFile(), e.getPassword()).getCredentials();
                credentialsMap.put(e.getUsage().toUpperCase(), credentials);
            }
        }
    }

    //todo: backend thread to getting wallet status
}
