package io.alpha.compound.intergation;

import io.alpha.compound.repository.CompoundLiqTxRepo;
import io.alpha.compound.entity.CompoundLiqTx;
import io.alpha.support.core.http.OkHttpClientWrapper;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Response;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import javax.transaction.Transactional;

@Service
@Slf4j
public class CexOrderApiService {
    @Value("${orderServiceUrl}")
    private String orderServiceUrl;
    @Autowired
    CompoundLiqTxRepo repo;

    @Transactional(rollbackOn = Exception.class)
    /**
     * 1. collateralAssetName = ETH, 需要对冲, 清算交易获得了ETH，金额为repayAmountInEth*0.05
     * 2. repayAssetName=ETH, 是否需要买回等额的ETH？清算获取稳定了币，支付了ETH，金额为repayAmountInEth
     */
    public Response sendCexOrder(CompoundLiqTx tx) {
        log.info("send cex order to hedge");
        OkHttpClientWrapper client = new OkHttpClientWrapper();
        Response response = client.url(orderServiceUrl).post(tx).execute();
        if (response.isSuccessful() && response.body() != null) {
            tx.setHedgeStatus(1);
            tx.setTxMsg(response.body().toString());
        }
        return response;
    }
}
