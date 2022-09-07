package io.qdriven.chains.config.tokens.repo;

import io.qdriven.chains.config.tokens.entity.BlockchainToken;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface TokenRepo extends JpaRepository<BlockchainToken, Long>,
        JpaSpecificationExecutor<BlockchainToken> {

    List<BlockchainToken> findByChainNameAndProtocolName(String chainName, String protocolName);

}