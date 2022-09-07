package io.alpha.compound.intergation.response;

import lombok.Data;

/**
 * {
 *   "address": "0xf5dce57282a584d2746faf1593d3121fcac444dc"
 *   "borrow_balance_underlying": {"value": "131.4682716123015"}
 *   "lifetime_borrow_interest_accrued": {"value": "0.44430505829286"}
 *   "lifetime_supply_interest_accrued": {"value": "0.0000021671829864899976"}
 *   "supply_balance_underlying": {"value": "0.0"}
 * }
 */
@Data
public class AccountCToken {
    private String address;
    private String symbol;
    private CompoundValue borrow_balance_underlying;
    private CompoundValue lifetime_borrow_interest_accrued;
    private CompoundValue lifetime_supply_interest_accrued;
    private CompoundValue supply_balance_underlying;
}
