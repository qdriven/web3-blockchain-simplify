package io.alpha.compound.controller;

import io.alpha.compound.service.CompoundDataService;
import io.alpha.compound.service.CompoundLiquidationService;
import io.alpha.configs.protocols.compound.CompoundConfigService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

@RestController
@Slf4j
//@OpenAPIDefinition(info = @Info(title = "Compound Protocol接口", version = "1.0"))
public class CompoundController {

    @Autowired
    CompoundDataService service;
    @Autowired
    CompoundConfigService configService;
    @Autowired
    CompoundLiquidationService liqService;

    /**
     * 1. 分页处理
     */
    @GetMapping("/compound/unhealthy")
//    @Operation(summary = "接近清算账户详细")
    public ResponseEntity<Map> getUnhealthyAccount() {
        return ResponseEntity.ok(service.getUnHealthyAccounts());
    }

    /**
     * TODO: token管理,数据库管理
     *
     * @return
     */
    @GetMapping("/compound/contracts")
//    @Operation(summary = "compound合约详细")
    public ResponseEntity<Map> getAllContracts() {
        return ResponseEntity.ok(configService.getConfig().getAllContracts());
    }

    /**
     * TODO: 清算任务的启动，停止
     */
    @GetMapping("/compound/job/stop")
//    @Operation(summary = "暂停清算任务")
    public void stopLiqMonitor() {
        liqService.shutdownLiqMonitorJob();
    }

    /**
     * TODO: 分页处理，
     *
     * @return
     */
    @GetMapping("/Compound/tx")
//    @Operation(summary = "获取清算交易详情")
    public ResponseEntity getCompoundLiqTx() {
        return ResponseEntity.ok(service.getAllCompoundLixTx());
    }

    @GetMapping("/Compound/configuration")
//    @Operation(summary = "获取运行时Compound相关配置")
    public ResponseEntity getCompoundConfiguration() {
        return ResponseEntity.ok(configService.getAllConfig());
    }

    @PostMapping("/Compound/configuration")
//    @Operation(summary = "设置运行时Compound相关配置")
    public ResponseEntity updateCompoundConfiguration(@RequestBody Map<String, String> config) {
        for (Map.Entry<String, String> entry : config.entrySet()) {
            configService.updateConfig(entry.getKey(), entry.getValue());
        }
        return ResponseEntity.ok(configService.getAllConfig());
    }

}
