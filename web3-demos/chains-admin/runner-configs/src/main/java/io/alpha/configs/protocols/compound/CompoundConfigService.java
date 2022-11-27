package io.alpha.configs.protocols.compound;

import cn.hutool.core.bean.BeanUtil;
import cn.hutool.core.util.StrUtil;
import cn.hutool.json.JSONUtil;
import io.alpha.configs.operation.entity.RunnerConfig;
import io.alpha.configs.operation.repo.RunnerConfigRepo;
import io.alpha.configs.protocols.compound.domain.CompoundConfigEnum;
import lombok.Data;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.annotation.Order;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.util.List;
import java.util.Map;

/**
 * 动态修改运行时参数，保存Compound相关的所有参数
 */
@Component
@Data
@Order(1)
public class CompoundConfigService {

    @Autowired
    RunnerConfigRepo runnerConfigRepo;

    //TODO：move to profile properties
    private String COMPOUND_API_URL = "https://api.compound.finance/";
    //Search Compound API Threshold,default is 1.00
    private String DEFAULT_HEALTH_THRESHOLD = "1.00";
    //Search Compound API MIN_BORROW Amount in ETH
    private String MIN_BORROW_VALUE_IN_ETH = "1.0";
    //Search Compound API Page Size
    private int COMPOUND_API_PAGE_SIZE = 50;
    private int COMPOUND_API_PAGE_NUM = 10;
    //Liquidate threshold in eth value
    private String LIQUIDATE_IN_ETH_THRESHOLD = "1.0";
    private int COMPOUND_LIQ_THREADS = 1;
    private String LIQ_ASSET_LIST = "CETH，CDAI，CUSDC";
    private boolean ENABLE_LIQUIDATION = false;
    private String LIQ_CLOSE_FACTOR = "0.48";
    private String Collateral_Factor = "0.6";

    public List<String> getLiqAllowedAssets(){
        return StrUtil.split(LIQ_ASSET_LIST,',',true,true);
    }

    @Autowired
    CompoundContractConfig config;

    @PostConstruct
    public void initCompoundConfiguration(){
        List<RunnerConfig> configs = runnerConfigRepo.findByProtocolNameAndActive(CompoundConstants.PROTOCOL_NAME,true);

        for (RunnerConfig runnerConfig : configs) {
            updateConfig(runnerConfig.getParameterName(),runnerConfig.getParameterValue());
        }
    }

    /**
     * 更具单个key，value更新配置
     *
     * @param key
     * @param value
     */
    public void updateConfig(String key, String value) {
        BeanUtil.setFieldValue(this, key, value);
    }

    public Object getConfig(String key) {
        return BeanUtil.getFieldValue(this, key);
    }

    public Map<String, Object> getAllConfig() {
        return BeanUtil.beanToMap(this);
    }

    /**
     * 根据JSON更新配置
     *
     * @param jsonString
     * @return
     */
    public Map<String, Object> updateConfig(String jsonString) {
        Map<String, Object> configJson = JSONUtil.toBean(jsonString, Map.class);
        for (Map.Entry<String, Object> entry : configJson.entrySet()) {
            if (entry.getValue() != null) {
                updateConfig(entry.getKey(), entry.getValue().toString());
            }
        }
        return getAllConfig();
    }

    /**
     * for init configuration data
     * //TODO: refresh after database data changed
     */
    public void initConfigData() {
        for (CompoundConfigEnum configEnum : CompoundConfigEnum.values()) {
            RunnerConfig runnerConfig = new RunnerConfig();
            runnerConfig.setActive(true);
            runnerConfig.setParameterName(configEnum.name());
            runnerConfig.setDescription(configEnum.getDesc());
            runnerConfig.setProtocolName(CompoundConstants.PROTOCOL_NAME);
            runnerConfig.setParameterValue(getConfig(configEnum.name()).toString());
            runnerConfigRepo.save(runnerConfig);
        }
    }
}
