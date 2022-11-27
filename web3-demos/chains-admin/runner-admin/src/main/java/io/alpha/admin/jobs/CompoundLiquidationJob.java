package io.alpha.admin.jobs;

import io.alpha.compound.service.liquidation.CompoundLiquidateService;
import io.alpha.compound.service.liquidation.LiquidateService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import xyz.erupt.job.handler.EruptJobHandler;

import java.util.Date;

@Service
@Slf4j
public class CompoundLiquidationJob implements EruptJobHandler {
    @Autowired
    CompoundLiquidateService compoundService;


    @Override
    public String exec(String code, String param) throws Exception {
        log.info("start of compound liquidation job at {} ", new Date());
        compoundService.compoundLiquidationJobTask();
        log.info("end of compound liquidation job at {}", new Date());
        return "compoundLiquidationJobTask executed";
    }

    @Override
    public void success(String result, String param) {
        log.info("execute liquidation job {},param: {}", result, param);
    }

    @Override
    public void error(Throwable throwable, String param) {
        log.error("execute liquidation job failed,err:{},param: {}", throwable.getCause(), param);
        log.error(throwable.getMessage());
    }
}
