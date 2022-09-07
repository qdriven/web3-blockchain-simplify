package io.alpha.support.core.eth.contract;


import io.alpha.support.core.eth.ContractInvoker;

/**
 * ERC20 合约Factory
 * @param <ERC20>
 * @param <ETH>
 */
public interface ContractFactory<ERC20, ETH> {

    ERC20 createErc20Contract(String contractAddress, ContractInvoker provider);

    ETH createETHContract(String contractAddress, ContractInvoker provider);
}
