package io.alpha.configs.protocols.compound.domain;


public enum CompoundConfigEnum {

    DEFAULT_HEALTH_THRESHOLD("默认的健康指数"),
    MIN_BORROW_VALUE_IN_ETH("最小借款ETH金额"),
    COMPOUND_API_PAGE_SIZE("compound api 返回PageSize"),
    COMPOUND_API_PAGE_NUM("compound api 抓取的页数"),
    LIQUIDATE_IN_ETH_THRESHOLD("清算ETH的最低数量"),
    COMPOUND_LIQ_THREADS("清算机器人线程数"),
    LIQ_ASSET_LIST("可清算资产"), //逗号间隔
    ENABLE_LIQUIDATION("是否启动清算机器人"),
    LIQ_CLOSE_FACTOR("清算金额系数"),
    Collateral_Factor("抵押系数"); //TODO: move to Contract Config;

    private String desc;
    CompoundConfigEnum(String desc) {
        this.desc=desc;
    }


    public String getDesc() {
        return desc;
    }
}
