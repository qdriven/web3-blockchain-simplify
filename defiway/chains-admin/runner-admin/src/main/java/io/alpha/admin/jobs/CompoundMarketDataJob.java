package io.alpha.admin.jobs;


import io.alpha.compound.service.CompoundDataService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import xyz.erupt.job.handler.EruptJobHandler;

import java.util.Date;

@Component
@Slf4j
//@Profile("prod")
//TODO: move to job
public class CompoundMarketDataJob implements EruptJobHandler {

    @Autowired
    CompoundDataService compoundService;

    @Override
    public String exec(String code, String param) throws Exception {
        log.info("start of compound market data job at {} ", new Date());
        compoundService.compoundMarketDataJob();
        log.info("end of compound liquidation job at {}", new Date());
        return "market data job successfully";
    }

    @Override
    public void success(String result, String param) {
        log.info("market data job result is {},params are {}", result, param);
    }

    @Override
    public void error(Throwable throwable, String param) {
        log.error(throwable.getMessage(), throwable.getCause());
    }

//    @PostConstruct
//    public void initCompoundMarketData() {
//        getCompoundMarketData();  //第一次加载数据到内存
//    }

//    //TODO: 动态任务配置
//    @Scheduled(fixedRate = 5 * 1000)
//    public void compoundLiquidationJob() {
//        log.info("start of compound liquidation job at {} ", new Date());
//        compoundService.compoundLiquidationJobTask();
//        log.info("end of compound liquidation job at {}", new Date());
//    }

//    @Scheduled(fixedRate = 60 * 1000 * 60 * 24)
//    public void getCompoundMarketData() {
//        log.info("start of compound market data job at {} ", new Date());
//        compoundService.compoundMarketDataJob();
//        log.info("end of compound liquidation job at {}", new Date());
//    }
}