package io.alpha.compound.service.liquidation.response;

import io.alpha.compound.intergation.response.CmpAccount;
import io.alpha.compound.service.liquidation.domain.LiqTransactionInfo;
import io.alpha.compound.service.liquidation.domain.LiquidationAccount;
import lombok.Data;

@Data
public class UnHealthyAccountResponse {
    private String comments;
    private boolean adviceForLiquidating;
    private CmpAccount originAccount;
    private LiqTransactionInfo advisedLiquidationInfo;
}
