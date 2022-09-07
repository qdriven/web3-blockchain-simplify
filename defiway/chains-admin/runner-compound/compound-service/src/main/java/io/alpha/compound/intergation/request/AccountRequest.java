package io.alpha.compound.intergation.request;

import io.alpha.compound.intergation.response.CompoundValue;
import lombok.Builder;
import lombok.Data;

import java.util.List;

/**
 * {
 * "addresses": [] // returns all accounts if empty or not included
 * "block_number": 0 // returns latest if given 0
 * "max_health": { "value": "10.0" }
 * "min_borrow_value_in_eth": { "value": "0.002" }
 * "page_number": 1
 * "page_size": 10
 * }
 */
@Data
@Builder
public class AccountRequest {
    private List<String> addresses;
    private long block_number;
    private CompoundValue max_health;
    private CompoundValue min_borrow_value_in_eth;
    private int page_number;
    private int page_size;
    private long block_timestamp;
    private String network;
}
