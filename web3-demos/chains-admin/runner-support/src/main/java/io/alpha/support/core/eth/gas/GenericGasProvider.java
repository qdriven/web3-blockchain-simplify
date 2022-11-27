package io.alpha.support.core.eth.gas;


import io.alpha.support.numbers.Wad18;
import org.web3j.tx.gas.ContractGasProvider;
import org.web3j.utils.Convert;

import java.math.BigInteger;

public class GenericGasProvider implements ContractGasProvider {
    GasApi api = new GasApi();
    private final String name;

    public GenericGasProvider(String name) {
        this.name = name;
    }

    @Override
    public BigInteger getGasPrice(String contractFunc) {
        return getGasPrice();
    }

    /**
     * faster gas price
     *
     * @return
     */
    @Deprecated
    @Override
    public BigInteger getGasPrice() {

        EthGasPrice price = api.getGasPriceByEthApi();
        Wad18 gasPrice;
        if (this.name.equalsIgnoreCase(GasProviderFactory.GasUsageEnum.STANDARD.name())) {
            gasPrice = new Wad18(Convert.toWei(price.getStandard()
                            .toString(),
                    Convert.Unit.GWEI).toBigInteger());
        } else {
            gasPrice = new Wad18(Convert.toWei(price.getFast()
                            .toString(),
                    Convert.Unit.GWEI).toBigInteger());
        }
        return gasPrice.toBigInteger();
    }

    @Override
    public BigInteger getGasLimit(String contractFunc) {
        return getGasLimit();
    }

    @Override
    @Deprecated
    public BigInteger getGasLimit() {
        //不要修改,已经足够高 90_000/200_000/400_000
        //100_000:out-of-gas
        //TODO: Getting historical gas limit data/ 1000_000*1000 * 10 9 次方 gwei
        if (this.name.equalsIgnoreCase(GasProviderFactory.GasUsageEnum.STANDARD.name())) {
            return BigInteger.valueOf(1000_000);
        } else {
            return BigInteger.valueOf(1500_000);
        }

    }

    /**
     * (gas limit x gas price) = transaction fee
     */
    public Wad18 getEstimatedEthValue() {
        BigInteger gasUsed = getGasPrice().multiply(getGasLimit());
        Wad18 ethValue = new Wad18(gasUsed);
        System.out.println(ethValue.toString());
        return ethValue;
    }
}
