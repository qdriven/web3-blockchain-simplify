package io.alpha.configs.protocols.compound;

import cn.hutool.core.collection.CollectionUtil;

import java.util.List;

public class CompoundConstants {

    public static String PROTOCOL_NAME ="COMPOUND";
    //Search Compound API Threshold,default is 1.00
    private String DEFAULT_HEALTH_THRESHOLD = "1.00";
    //Search Compound API MIN_BORROW Amount in ETH
    private String MIN_BORROW_VALUE_IN_ETH = "1.0";
    //Search Compound API Page Size
    private int COMPOUND_API_PAGE_SIZE = 50;
    private int COMPOUND_API_PAGE_NUM = 30;
    //Liquidate threshold in eth value
    private String LIQUIDATE_IN_ETH_THRESHOLD = "1.0";
    private int COMPOUND_LIQ_THREADS = 1;
    private List<String> LIQ_ASSET_LIST = CollectionUtil.newArrayList( "CETH", "CDAI","CUSDC");
    private boolean ENABLE_LIQUIDATION = false;
    private String LIQ_CLOSE_FACTOR = "0.48";
    private String Collateral_Factor = "0.6";


    public static String contractInfo = "[\n" +
            "  {\n" +
            "    \"ctoken_symbol\": \"cBAT\",\n" +
            "    \"abi_href\": \"https://compound.finance/docs/abi/mainnet/cBAT\",\n" +
            "    \"contract_address\": \"0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e\",\n" +
            "    \"abi_json_file\": \"mainnet_cBAT.json\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"ctoken_symbol\": \"cCOMP\",\n" +
            "    \"abi_href\": \"https://compound.finance/docs/abi/mainnet/cCOMP\",\n" +
            "    \"contract_address\": \"0x70e36f6bf80a52b3b46b3af8e106cc0ed743e8e4\",\n" +
            "    \"abi_json_file\": \"mainnet_cCOMP.json\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"ctoken_symbol\": \"cDAI\",\n" +
            "    \"abi_href\": \"https://compound.finance/docs/abi/mainnet/cDAI\",\n" +
            "    \"contract_address\": \"0x5d3a536e4d6dbd6114cc1ead35777bab948e3643\",\n" +
            "    \"abi_json_file\": \"mainnet_cDAI.json\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"ctoken_symbol\": \"cETH\",\n" +
            "    \"abi_href\": \"https://compound.finance/docs/abi/mainnet/cETH\",\n" +
            "    \"contract_address\": \"0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5\",\n" +
            "    \"abi_json_file\": \"mainnet_cETH.json\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"ctoken_symbol\": \"cREP\",\n" +
            "    \"abi_href\": \"https://compound.finance/docs/abi/mainnet/cREP\",\n" +
            "    \"contract_address\": \"0x158079ee67fce2f58472a96584a73c7ab9ac95c1\",\n" +
            "    \"abi_json_file\": \"mainnet_cREP.json\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"ctoken_symbol\": \"cSAI\",\n" +
            "    \"abi_href\": \"https://compound.finance/docs/abi/mainnet/cSAI\",\n" +
            "    \"contract_address\": \"0xf5dce57282a584d2746faf1593d3121fcac444dc\",\n" +
            "    \"abi_json_file\": \"mainnet_cSAI.json\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"ctoken_symbol\": \"cUNI\",\n" +
            "    \"abi_href\": \"https://compound.finance/docs/abi/mainnet/cUNI\",\n" +
            "    \"contract_address\": \"0x35a18000230da775cac24873d00ff85bccded550\",\n" +
            "    \"abi_json_file\": \"mainnet_cUNI.json\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"ctoken_symbol\": \"cUSDC\",\n" +
            "    \"abi_href\": \"https://compound.finance/docs/abi/mainnet/cUSDC\",\n" +
            "    \"contract_address\": \"0x39aa39c021dfbae8fac545936693ac917d5e7563\",\n" +
            "    \"abi_json_file\": \"mainnet_cUSDC.json\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"ctoken_symbol\": \"cUSDT\",\n" +
            "    \"abi_href\": \"https://compound.finance/docs/abi/mainnet/cUSDT\",\n" +
            "    \"contract_address\": \"0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9\",\n" +
            "    \"abi_json_file\": \"mainnet_cUSDT.json\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"ctoken_symbol\": \"cWBTC\",\n" +
            "    \"abi_href\": \"https://compound.finance/docs/abi/mainnet/cWBTC\",\n" +
            "    \"contract_address\": \"0xc11b1268c1a384e55c48c2391d8d480264a3a7f4\",\n" +
            "    \"abi_json_file\": \"mainnet_cWBTC.json\"\n" +
            "  },\n" +
            "  {\n" +
            "    \"ctoken_symbol\": \"cZRX\",\n" +
            "    \"abi_href\": \"https://compound.finance/docs/abi/mainnet/cZRX\",\n" +
            "    \"contract_address\": \"0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407\",\n" +
            "    \"abi_json_file\": \"mainnet_cZRX.json\"\n" +
            "  }\n" +
            "]";

}
