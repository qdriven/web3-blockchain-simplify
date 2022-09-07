package io.alpha.support.core.eth.gas;

import cn.hutool.json.JSONUtil;
import io.alpha.support.core.http.OkHttpClientWrapper;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Response;

import java.util.Objects;


@Slf4j
public class GasApi {

    private String ETH_GAS_URL = "https://www.etherchain.org/api/gasPriceOracle";
    private String ETH_GASSTATION_URL = "https://ethgasstation.info/json/ethgasAPI.json";

    /**
     * {"safeLow":77,"standard":85,"fast":92,"fastest":101}
     */
    public EthGasPrice getGasPriceByEthApi() {
        Response response = new OkHttpClientWrapper().url(ETH_GAS_URL)
                .get().execute();
        try {
            assert response.body() != null;
            return JSONUtil.toBean(response.body().string(), EthGasPrice.class);
        } catch (Exception e) {
            log.error("call eth gas api failed");
            throw new GasProviderException(e);
        }
    }

    /**
     * {"fast": 940.0, "fastest": 1020.0, "safeLow": 800.0, "average": 820.0, "block_time": 11.258064516129032, "blockNum": 11151719, "speed": 0.9960360156825414, "safeLowWait": 9.0, "avgWait": 1.3, "fastWait": 0.5, "fastestWait": 0.4, "gasPriceRange": {"1020": 0.4, "1000": 0.4, "980": 0.4, "960": 0.4, "940": 0.5, "920": 0.5, "900": 0.7, "880": 0.7, "860": 0.8, "840": 1.0, "820": 1.3, "800": 9.0, "780": 187.6, "760": 187.6, "740": 187.6, "720": 187.6, "700": 187.6, "680": 187.6, "660": 187.6, "640": 187.6, "620": 187.6, "600": 187.6, "580": 187.6, "560": 187.6, "540": 187.6, "520": 187.6, "500": 187.6, "480": 187.6, "460": 187.6, "440": 187.6, "420": 187.6, "400": 187.6, "380": 187.6, "360": 187.6, "340": 187.6, "320": 187.6, "300": 187.6, "280": 187.6, "260": 187.6, "240": 187.6, "220": 187.6, "200": 187.6, "190": 187.6, "180": 187.6, "170": 187.6, "160": 187.6, "150": 187.6, "140": 187.6, "130": 187.6, "120": 187.6, "110": 187.6, "100": 187.6, "90": 187.6, "80": 187.6, "70": 187.6, "60": 187.6, "50": 187.6, "40": 187.6, "30": 187.6, "20": 187.6, "10": 187.6, "8": 187.6, "6": 187.6, "4": 187.6}}
     *
     * @return
     */

    public EthStationGasPrice getGasPriceByGasStationApi() {
        Response response = new OkHttpClientWrapper().url(ETH_GASSTATION_URL)
                .get().execute();
        try {
            Objects.requireNonNull(response.body());
            return JSONUtil.toBean(response.body().string(),
                    EthStationGasPrice.class);
        } catch (Exception e) {
            log.error("call eth gas station api failed");
            throw new GasProviderException(e);
        }
    }
}
