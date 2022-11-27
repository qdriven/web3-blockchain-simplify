package io.alpha.compound.repository;

import io.alpha.compound.entity.CompoundUnHealthAsset;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface CompoundUnhealthAssetRepo extends JpaRepository<CompoundUnHealthAsset, Long>,
        JpaSpecificationExecutor<CompoundUnHealthAsset> {

    List<CompoundUnHealthAsset> findByAccountAddressAndAddress(String accountAddress, String address);

}
