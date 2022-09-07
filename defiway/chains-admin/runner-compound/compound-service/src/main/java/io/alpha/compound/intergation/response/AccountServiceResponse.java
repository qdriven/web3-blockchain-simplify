package io.alpha.compound.intergation.response;


import io.alpha.compound.intergation.request.AccountRequest;
import lombok.Data;

import java.util.ArrayList;
import java.util.List;

/**
 * {
 * "address": "0xbac065be2e8ca097e9ac924e94af000dd3a5663"
 * "health": { "value": "1.07264275673050348990755599431194797431802239523113293682619605751591901" }
 * "tokens": [
 * {
 * "address": "0xf5dce57282a584d2746faf1593d3121fcac444dc"
 * "borrow_balance_underlying": {"value": "131.4682716123015"}
 * "lifetime_borrow_interest_accrued": {"value": "0.44430505829286"}
 * "lifetime_supply_interest_accrued": {"value": "0.0000021671829864899976"}
 * "supply_balance_underlying": {"value": "0.0"}
 * }
 * ],
 * "total_borrow_value_in_eth": {"value": "0.5100157047140227313856015174794473200000000000000000000000000000" }
 * "total_collateral_value_in_eth": {"value": "0.54706465148029978664135447293587201124121731200000000000000000000000000" }
 * }
 */
@Data
public class AccountServiceResponse {
    private AccountRequest request;
    private PaginationSummary pagination_summary;
    private CompoundValue health;
    private List<CmpAccount> accounts = new ArrayList<>();
    private CompoundValue total_borrow_value_in_eth;
    private CompoundValue total_collateral_value_in_eth;
}
