package io.alpha.admin.modules;//package io.alpha.defirunner.modules;
//
//import org.springframework.stereotype.Component;
//
//import java.util.HashMap;
//import java.util.Map;
//import java.util.concurrent.ConcurrentHashMap;
//
///**
// * TODO: 统一管理各个协议的配置
// * 1. 不同协议配置注册进ConfigManager
// * 2. 获取配置也统一通过ConfigManage获取
// * 3. 保存配置到数据库，确保下次启动时使用的是上次修改的参数
// */
//@Component
//public class ConfigManager {
//    public Map<String, Map<String, String>> configMap = new ConcurrentHashMap<>();
//    public Map<String, String> getModuleConfig(String moduleName) {
//        return configMap.get(moduleName);
//    }
//
//    /**
//     * register module config
//     * @param moduleName
//     * @param configs
//     * @return
//     */
//    public ConfigManager registerModuleConfig(String moduleName, Map<String, String> configs) {
//        configMap.put(moduleName, configs);
//        return this;
//    }
//
//    /**
//     * update module configuration setting
//     * @param moduleName
//     * @param key
//     * @param value
//     * @return
//     */
//    public Map<String, String> updateModuleConfig(String moduleName, String key, String value) {
//        if (configMap.get(moduleName) == null) {
//            Map<String, String> map = new HashMap<>();
//            map.put(key, value);
//            configMap.put(moduleName, map);
//        } else {
//            configMap.get(moduleName).put(key, value);
//        }
//        return configMap.get(moduleName);
//    }
//}
