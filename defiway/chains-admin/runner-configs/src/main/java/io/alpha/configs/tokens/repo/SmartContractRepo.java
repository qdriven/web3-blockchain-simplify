package io.alpha.configs.tokens.repo;

import io.alpha.configs.tokens.entity.SmartContract;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.stereotype.Repository;

@Repository
public interface SmartContractRepo extends JpaRepository<SmartContract, Long>,
        JpaSpecificationExecutor<SmartContract> {

}