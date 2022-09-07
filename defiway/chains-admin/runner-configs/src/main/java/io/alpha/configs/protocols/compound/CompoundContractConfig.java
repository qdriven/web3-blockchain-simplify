package io.alpha.configs.protocols.compound;

import io.alpha.configs.protocols.compound.domain.CTokenInfo;
import io.alpha.configs.tokens.entity.BlockchainToken;
import io.alpha.configs.tokens.repo.TokenRepo;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.DependsOn;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import static io.alpha.configs.protocols.compound.CompoundConstants.PROTOCOL_NAME;

@Component
@Slf4j
@DependsOn("tokenRepo")
public class CompoundContractConfig {

    @Autowired
    TokenRepo tokenRepo;

    public Map<String, CTokenInfo> getAllContracts() {
        return contractsMap;
    }
    private Map<String, CTokenInfo> contractsMap = new ConcurrentHashMap<>();

    @PostConstruct
    public void initContractMap() {
        //TODO: testnet
        List<BlockchainToken> tokens= tokenRepo.findByChainNameAndProtocolName("ETH", PROTOCOL_NAME);
//        if(tokens.size() ==0) throw new RuntimeException("Compound Token Information is not configured");
        for (BlockchainToken token : tokens) {
            CTokenInfo cTokenInfo = new CTokenInfo();
            cTokenInfo.setContract_address(token.getAddress());
            cTokenInfo.setCtoken_symbol(token.getSymbol());
            cTokenInfo.setDecimal(token.getDecimalValue());
            cTokenInfo.setAbi_href(token.getAbiUrl());
            contractsMap.put(cTokenInfo.getCtoken_symbol().toUpperCase(), cTokenInfo);
        }
    }

    public CTokenInfo getAssetInfoByContractAddress(String assetAddress) {
        for (Map.Entry<String, CTokenInfo> entry : contractsMap.entrySet()) {
            if (entry.getValue().getContract_address().equalsIgnoreCase(assetAddress)) {
                return entry.getValue();
            }
        }
        return null;
    }
}
