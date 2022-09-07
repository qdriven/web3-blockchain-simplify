package io.alpha.compound.service.liquidation;

import cn.hutool.core.bean.BeanUtil;
import io.alpha.compound.contracts.CompoundContractFactory;
import io.alpha.compound.contracts.ctoken.CERC20Contract;
import io.alpha.compound.contracts.ctoken.CETHContract;
import io.alpha.compound.intergation.CexOrderApiService;
import io.alpha.compound.repository.CompoundLiqTxRepo;
import io.alpha.compound.service.CompoundDataService;
import io.alpha.compound.service.liquidation.domain.LiqTransactionInfo;
import io.alpha.configs.protocols.compound.CompoundContractConfig;
import io.alpha.compound.entity.CompoundLiqTx;
import io.alpha.support.core.eth.ContractInvoker;
import io.alpha.support.core.eth.ContractInvokerFactory;
import io.alpha.support.core.eth.gas.GasProviderFactory;
import io.alpha.support.core.eth.wallet.WalletUsageEnum;
import io.alpha.support.numbers.NumberUtil;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.utils.Convert;

@Component
@Slf4j
public class CompoundLiquidator {

    @Autowired
    CompoundContractConfig config;
    @Autowired
    CompoundContractFactory factory;
    @Autowired
    ContractInvokerFactory providers;
    @Autowired
    CompoundLiqTxRepo liqTxRepo;
    @Autowired
    CompoundDataService dataService;

    @Autowired
    CexOrderApiService orderApiService;

    private ContractInvoker createInvoker(LiqTransactionInfo tx) {
        if (NumberUtil.isGreater(tx.getRepayAmountInEth(), "25.000")) {
            return providers.createProvider(WalletUsageEnum.BIG.name(), GasProviderFactory.GasUsageEnum.FAST.name());
        } else {
            return providers.createProvider(WalletUsageEnum.SMALL.name(), GasProviderFactory.GasUsageEnum.STANDARD.name());
        }
    }

    private TransactionReceipt cerc20LiquidateBorrow(LiqTransactionInfo tx) throws Exception {
        //TODO:数值转换都在CompoundCalculate中处理
        ContractInvoker provider = createInvoker(tx);
        CERC20Contract borrowedTokenContract = factory.createErc20Contract(tx.getRepayAssetAddress(), //borrow asset hash address
                provider);
        String decimal = config.getAssetInfoByContractAddress(tx.getRepayAssetAddress()).getDecimal();
        return borrowedTokenContract.liquidateBorrow(
                tx.getBorrowAddress(),
                NumberUtil.toBigInteger(tx.getRepayAmount(), decimal), //todo: for different decimal
                tx.getCollateralAssetAddress()).send();
    }

    private TransactionReceipt cethLiquidateBorrow(LiqTransactionInfo tx) throws Exception {
        ContractInvoker invoker = createInvoker(tx);
        CETHContract borrowedTokenContract = factory.createETHContract(tx.getRepayAssetAddress(),
                invoker);
        return borrowedTokenContract.liquidateBorrow(
                tx.getBorrowAddress(),
                tx.getCollateralAssetAddress(),
                Convert.toWei(tx.getRepayAmount(), Convert.Unit.ETHER).toBigInteger() //这个一直成功
        ).send();

    }
//    TransactionManager txManager = new FastRawTransactionManager(web3j, creds);
//
//    String txHash = txManager.sendTransaction(DefaultGasProvider.GAS_PRICE, DefaultGasProvider.GAS_LIMIT,
//            documentRegistry.getContractAddress(), txData, BigInteger.ZERO).getTra
//    TransactionReceiptProcessor receiptProcessor =
//            new PollingTransactionReceiptProcessor(web3j, TransactionManager.DEFAULT_POLLING_FREQUENCY,
//                    TransactionManager.DEFAULT_POLLING_ATTEMPTS_PER_TX_HASH);
//
//    TransactionReceipt txReceipt = receiptProcessor.waitForTransactionReceipt(txHash);

    public void liquidateBorrow(LiqTransactionInfo tx) {
        // 1. TODO: APPROVE 给ctoken合约
        // 3. TODO: CompleteFuture 处理
        try {
            TransactionReceipt txReceipt;
            if (tx.getRepayAssetName().toUpperCase().contains("CETH")) {
                txReceipt = cethLiquidateBorrow(tx); //确定成功 TODO: 异步处理
            } else {
                txReceipt = cerc20LiquidateBorrow(tx); //会失败？为什么？//TODO: 处理异常
            }
            log.info("transaction receipt is {}", txReceipt);
            CompoundLiqTx liqTx = new CompoundLiqTx();
            BeanUtil.copyProperties(tx, liqTx);
            liqTx.setTransactionHash(txReceipt.getTransactionHash());
            liqTx.setTxStatus(txReceipt.getStatus().replaceAll("0x", ""));
            liqTxRepo.save(liqTx);
            orderApiService.sendCexOrder(liqTx);
            //TODO: Add Job to Check

        } catch (Exception e) {
            log.error("liq transaction failed,{},{}", tx, e);
            CompoundLiqTx liqTx = new CompoundLiqTx();
            BeanUtil.copyProperties(tx, liqTx);
            liqTx.setTxStatus("0"); //FAILED TX
            liqTxRepo.save(liqTx);
        }

    }
}
