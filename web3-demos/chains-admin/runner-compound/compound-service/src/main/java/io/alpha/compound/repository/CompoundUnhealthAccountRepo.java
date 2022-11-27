package io.alpha.compound.repository;

import io.alpha.compound.entity.CompoundUnhealthAccount;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface CompoundUnhealthAccountRepo extends JpaRepository<CompoundUnhealthAccount, Long>,
        JpaSpecificationExecutor<CompoundUnhealthAccount> {

    List<CompoundUnhealthAccount> findByAddress(String address);
}
