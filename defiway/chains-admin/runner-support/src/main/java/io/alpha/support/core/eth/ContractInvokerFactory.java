package io.alpha.support.core.eth;


import io.alpha.support.core.eth.wallet.WalletService;
import io.alpha.support.core.eth.gas.GasProviderFactory;
import io.alpha.support.core.eth.web3.Web3jFactory;
import lombok.Data;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.annotation.Order;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;
import org.web3j.crypto.Credentials;


@Component
@Slf4j
@Data
@Order
public class ContractInvokerFactory {

    @Autowired
    private Web3jFactory web3jFactory;
    @Autowired
    private GasProviderFactory gasProviderFactory;
    @Value("${compound.testPrivateKey:ignore}")
    private String testPrivateKey;
    @Autowired
    private WalletService walletService;


    public ContractInvoker createProvider(String walletUsage, String gasUsage) {
        log.info("create contract invoker provider by {},{}", walletUsage, gasUsage);
        ContractInvoker provider = new ContractInvoker();
        provider.setWeb3j(web3jFactory.getWeb3j());
        if (testPrivateKey != null && !testPrivateKey.equalsIgnoreCase("ignore")) {
            provider.setCredentials(Credentials.create(testPrivateKey));
        } else {
            //TODO: use different wallets
            provider.setCredentials(walletService.getWalletByUsage(walletUsage));
        }
        provider.setGasProvider(gasProviderFactory.getGasProviderByGasUsage(gasUsage));
        log.info("create contract invoker provider by {},{} successfully", walletUsage, gasUsage);
        return provider;
    }

}
