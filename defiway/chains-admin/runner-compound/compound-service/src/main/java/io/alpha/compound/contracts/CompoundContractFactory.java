package io.alpha.compound.contracts;

import io.alpha.compound.contracts.ctoken.CERC20Contract;
import io.alpha.compound.contracts.ctoken.CETHContract;
import io.alpha.support.core.eth.ContractInvoker;
import io.alpha.support.core.eth.contract.ContractFactory;
import org.springframework.stereotype.Component;

@Component
public class CompoundContractFactory implements ContractFactory<CERC20Contract, CETHContract> {
    @Override
    public CERC20Contract createErc20Contract(String contractAddress, ContractInvoker provider) {
        return CERC20Contract.load(contractAddress, provider.getWeb3j(), provider.getCredentials(), provider.getGasProvider());
    }

    @Override
    public CETHContract createETHContract(String contractAddress, ContractInvoker provider) {
        return CETHContract.load(contractAddress, provider.getWeb3j(), provider.getCredentials(), provider.getGasProvider());
    }

    /**
     * 创建CompControll合约
     *
     * @param contractAddress
     * @param provider
     * @return
     */
    public ComptrollerContract createComptrollerContract(String contractAddress, ContractInvoker provider) {
        return ComptrollerContract.load(contractAddress, provider.getWeb3j(), provider.getCredentials(), provider.getGasProvider());
    }

    /**
     * 创建USDT合约, USDT合约不能工作
     *
     * @param contractAddress
     * @param provider
     * @return
     */
    @Deprecated
    public UsdtContract createUsdtContract(String contractAddress, ContractInvoker provider) {
        return UsdtContract.load(contractAddress, provider.getWeb3j(), provider.getCredentials(), provider.getGasProvider());
    }
}
