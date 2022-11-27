# Flashloan

Flashloan:
a flash loan is a risk-free instantaneous loan whose time of borrowing and payback happens within one blockchain transaction.

## Mechaninc

Flashloan take advantage of Ethereum's state-reversing feature.

```
EIP-140: revert instruction

The REVERT instruction provides a way to stop execution and revert state changes, without consuming all provided gas and with the ability to return a reason.

```

A simple way to describe how it works:

1. smart contract call codes where the borrow function is at the begining
2. use the loan in middle function
3. then replay the loan
4. if replayment is fine, then transaction is done
   
## Use Cases

- Arbitrage
- Debt Refinancing
- wash trading
![img](https://hackingdistributed.com/images/2020/flash-loan.png)

- aave
- dydx

## Flashloan Post-Mortem

  


