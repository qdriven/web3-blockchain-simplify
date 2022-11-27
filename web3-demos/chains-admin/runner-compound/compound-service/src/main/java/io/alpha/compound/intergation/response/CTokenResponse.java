package io.alpha.compound.intergation.response;
/*
 *  Copyright 2020-2021 QMETA - this repo owner
 *
 *  Licensed under unkown License
 */

import lombok.Data;

@Data
public class CTokenResponse {
    private CompoundValue borrow_rate;
    private CompoundValue cash;
    private CompoundValue collateral_factor;
    private String interest_rate_model_address;
    private String name;
    private int number_of_borrowers;
    private int number_of_suppliers;
    private CompoundValue reserves;
    private CompoundValue reserve_factor;
    private CompoundValue supply_rate;
    private String symbol;
    private String token_address;
    private CompoundValue total_borrows;
    private CompoundValue total_supply;
    private String underlying_address;
    private String underlying_name ;
    private String underlying_symbol ;
    private CompoundValue underlying_price ;
}
