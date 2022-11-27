package io.alpha.support.core.eth;


import lombok.Data;
import lombok.extern.slf4j.Slf4j;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.tx.gas.ContractGasProvider;
import org.web3j.tx.response.QueuingTransactionReceiptProcessor;


@Slf4j
@Data
public class ContractInvoker {
    private Web3j web3j;
    private Credentials credentials;
    private ContractGasProvider gasProvider;
    private QueuingTransactionReceiptProcessor txProcessor;
}
