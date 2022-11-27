package io.alpha.configs.protocols.compound.domain;

import lombok.Data;

@Data
public  class CTokenInfo {
    private String contract_address;
    private String abi_href;
    private String ctoken_symbol;
    private String abi_json_file;
    private String decimal;
    //TODO: 抵押数量/清算值
}