package io.alpha.compound.repository;


import io.alpha.compound.entity.CompoundMarketData;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.stereotype.Repository;

@Repository
public interface CompoundMarketDataRepo extends JpaRepository<CompoundMarketData, Long>,
        JpaSpecificationExecutor<CompoundMarketData> {

}
