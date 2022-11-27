package io.alpha.support.core.eth.wallet;

import cn.hutool.core.io.FileUtil;
import cn.hutool.core.util.StrUtil;
import lombok.extern.slf4j.Slf4j;
import org.web3j.crypto.CipherException;
import org.web3j.crypto.Credentials;
import org.web3j.crypto.WalletUtils;

import java.io.File;
import java.io.IOException;
import java.nio.charset.Charset;
import java.security.InvalidAlgorithmParameterException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;

/**
 * TODO: Credential Factory to Create multiple accounts
 */
@Slf4j
public class WalletProvider {

    private Credentials credentials;

    public WalletProvider(String password, String wallet) {
        if (wallet.equals("") || password.equals("")) {
            log.error("WALLET OR PASSWORD IS EMPTY");
        }

        try {
            credentials = WalletUtils.loadJsonCredentials(password, wallet);
            log.trace("Credentials loaded");
        } catch (Exception e) {
            log.error("Exception", e);
        }
    }

    public Credentials getCredentials() {
        return credentials;
    }


    public static WalletProvider getWalletProvider(String walletPath, String pwd) {
        String jsonSource = FileUtil.readString(new File(walletPath), Charset.defaultCharset());
        return new WalletProvider(pwd, jsonSource);
    }

    public static String generateWallet(String path, String pwd) {
        File file = new File(path);
        try {
            return WalletUtils.generateNewWalletFile(pwd, file);
        } catch (CipherException | InvalidAlgorithmParameterException | NoSuchAlgorithmException | NoSuchProviderException | IOException e) {
            log.error("generate wallet file failed, {}", e.getMessage());
        }
        return StrUtil.EMPTY;
    }
}