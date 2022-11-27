package io.alpha.support.core.eth.web3;

import lombok.Data;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.annotation.Order;
import org.springframework.stereotype.Component;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.exceptions.ClientConnectionException;
import org.web3j.protocol.http.HttpService;

import javax.annotation.PostConstruct;
import java.util.ArrayList;
import java.util.List;

@Slf4j
@Data
@Component
@Order(1)
public class Web3jFactory {

    private List<Web3j> web3Services = new ArrayList<>();
    @Value("${ethServiceUrl}")
    private String ethServiceUrl;

    public Web3jFactory(String ethServiceUrl) {
        this.ethServiceUrl = ethServiceUrl;
    }

    public Web3jFactory() {
    }

    @PostConstruct
    public void setupWeb3j() {
        log.info("eth service url is {}", this.ethServiceUrl);
        String[] serviceUrls = this.ethServiceUrl.split(";");
        for (String sUrl : serviceUrls) {
            web3Services.add(createWeb3j(sUrl));
        }
    }

    public Web3j createWeb3j(String ethServiceUrl) {
        Web3j web3j = Web3j.build(new HttpService(ethServiceUrl));
        try {
            log.info(
                    "Connected to Ethereum client version: {}",
                    web3j.web3ClientVersion().send().getWeb3ClientVersion());
        } catch (Exception e) {
            if (e instanceof ClientConnectionException) {
                log.error("Check your eth node service ");
            }
            log.error("Exception -> Exit", e);
            System.exit(0);
        }
        return web3j;
    }

    //TODO: GET Web3 by availability
    public Web3j getWeb3j(){
        return this.web3Services.get(0);
    }
}
