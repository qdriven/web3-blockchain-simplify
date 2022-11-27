package io.alpha.compound.repository;


import io.alpha.compound.entity.CompoundLiqTx;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.stereotype.Repository;

@Repository
public interface CompoundLiqTxRepo extends JpaRepository<CompoundLiqTx, Long>,
        JpaSpecificationExecutor<CompoundLiqTx> {

}
