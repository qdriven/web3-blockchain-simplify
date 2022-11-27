package io.alpha.compound.entity;

import lombok.Data;
import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.annotations.DynamicInsert;
import org.hibernate.annotations.DynamicUpdate;
import org.hibernate.annotations.UpdateTimestamp;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "cmp_unhealth_assets")
@Data
@DynamicInsert
@DynamicUpdate
public class CompoundUnHealthAsset {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String address;
    private String accountAddress;
    private String symbol;
    private String borrow_balance_underlying;
    private String lifetime_borrow_interest_accrued;
    private String lifetime_supply_interest_accrued;
    private String supply_balance_underlying;
    @CreationTimestamp
    private Date createdTime = new Date();
    /**
     * 更新时间
     */
    @UpdateTimestamp
    private Date updatedTime = new Date();
}
