package io.alpha.compound.wallet;//package io.alpha.defirunner.modules.compound.wallet;
//
//import lombok.Data;
//import org.springframework.stereotype.Component;
//import org.web3j.crypto.Credentials;
//
//import javax.annotation.PostConstruct;
//import java.util.Map;
//import java.util.concurrent.ConcurrentHashMap;
//
///**
// * TODO: Wallet管理
// */
//@Component
//public class LocalWalletManager {
//
//    private Map<String, WalletWithUseFlag> wallets = new ConcurrentHashMap<>();
//
//    //1. Load All Wallets
//    //2. Assign Wallet to different threads
//    //3. Assign Wallet as different Role
//    //4. 正常钱包/大额清算钱包/手续费钱包/
//    @PostConstruct
//    public void initWallet() {
//
//    }
//
//    public synchronized WalletWithUseFlag getWallet(String walletName) {
//        //
//        WalletWithUseFlag wallet = wallets.get(walletName);
//        wallet.setInUse(true);
//        return wallet;
//    }
//
//    public synchronized WalletWithUseFlag returnWallet(String walletName) {
//        wallets.get(walletName).setInUse(false);
//        return wallets.get(walletName);
//    }
//
//    @Data
//    public static class WalletWithUseFlag {
//        private boolean inUse = false;
//        private Credentials credentials;
//        private String walletName;
//    }
//}
