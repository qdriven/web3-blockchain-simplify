package io.alpha.configs.protocols.compound;

import cn.hutool.json.JSONUtil;
import io.alpha.configs.protocols.compound.domain.CTokenInfo;
import io.alpha.configs.tokens.entity.BlockchainToken;
import io.alpha.configs.tokens.repo.TokenRepo;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.List;

import static io.alpha.configs.protocols.compound.CompoundConstants.*;

//@Component
//TODO: only for first time use to save data to database
public class CompoundDataInitializer {

    @Autowired
    private TokenRepo tokenRepo;

    public void initCompoundData() {
        List<CTokenInfo> contractArray = JSONUtil.toList(JSONUtil.parseArray(contractInfo), CTokenInfo.class);
        for (CTokenInfo cTokenInfo : contractArray) {
            if (cTokenInfo.getCtoken_symbol().equalsIgnoreCase("cUSDC")
                    || cTokenInfo.getCtoken_symbol().equalsIgnoreCase("cUSDT")) {
                cTokenInfo.setDecimal("6");
            } else if (cTokenInfo.getCtoken_symbol().equalsIgnoreCase("cWBTC")) {
                cTokenInfo.setDecimal("8");
            } else {
                cTokenInfo.setDecimal("18");
            }
            BlockchainToken token = new BlockchainToken();
            token.setChainName("ETH");
            token.setChainType("MAIN");
            token.setDecimalValue(cTokenInfo.getDecimal());
            token.setAbiUrl(cTokenInfo.getAbi_href());
            token.setSymbol(cTokenInfo.getCtoken_symbol());
            token.setAddress(cTokenInfo.getContract_address());
            token.setProtocolName(PROTOCOL_NAME);
            token.setName(cTokenInfo.getCtoken_symbol());
            token.setDescription("Compound CToken "+cTokenInfo.getAbi_href());
            tokenRepo.save(token);
        }

    }
}
