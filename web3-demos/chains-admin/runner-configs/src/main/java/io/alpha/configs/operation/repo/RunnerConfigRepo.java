package io.alpha.configs.operation.repo;

import io.alpha.configs.operation.entity.RunnerConfig;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface RunnerConfigRepo extends JpaRepository<RunnerConfig, Long>,
        JpaSpecificationExecutor<RunnerConfig> {
    List<RunnerConfig> findByProtocolNameAndActive(String protocolName,boolean active);
}