package io.alpha.compound.service.liquidation;

import io.alpha.compound.intergation.response.CmpAccount;
import io.alpha.compound.service.liquidation.domain.LiqTransactionInfo;

public interface LiquidateService {

    /**
     * setup liquidation job;
     */
//    public void setupLiquidationJob();
    /**
     * Check if account is eligible to liquidate
     */
    public boolean preLiquidationCheck(CmpAccount account);

    /**
     * calculate liquidation data for a given account
     */
    public LiqTransactionInfo calculateLiquidationAccount(CmpAccount account);

    /**
     *  liquidate account
     */
    public void liquidate(CmpAccount account);

    /**
     * handlers after liquidation
     */
    public void afterLiquidation(CmpAccount account);
}