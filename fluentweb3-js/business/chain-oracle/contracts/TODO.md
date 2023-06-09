# TODO

1. [] Analysis


## Analysis

- ChangeEvent


## Spec

```sh
Field	Type	Description
makerToken	address	The ERC20 token the maker is selling and the maker is selling to the taker. [required]
takerToken	address	The ERC20 token the taker is selling and the taker is selling to the maker. [required]
makerAmount	uint128	The amount of makerToken being sold by the maker. [required]
takerAmount	uint128	The amount of takerToken being sold by the taker. [required]
takerTokenFeeAmount	uint128	Amount of takerToken paid by the taker to the feeRecipient. [optional; default 0]
maker	address	The address of the maker, and signer, of this order. [required]
taker	address	Allowed taker address. Set to zero to allow any taker. [optional; default 0]
sender	address	Allowed address to call fillLimitOrder() (msg.sender). This is the same as taker, expect when using meta-transactions. Set to zero to allow any caller. [optional; default 0]
feeRecipient	address	Recipient of maker token or taker token fees (if non-zero). [optional; default 0]
pool	bytes32	The staking pool to attribute the 0x protocol fee from this order. Set to zero to attribute to the default pool, not owned by anyone. [optional; default 0]
expiry	uint64	The Unix timestamp in seconds when this order expires. [required]
salt	uint256	Arbitrary number to enforce uniqueness of the order hash. [required]
```

```json
{
    "type": "update",
    "channel": "orders",
    "payload": [
        {
            "order": {
                "chainId": 1,
                "verifyingContract": "0xdef1c0ded9bec7f1a1670819833240f027b25eff",
                "makerToken": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
                "takerToken": "0xe41d2489571d322189246dafa5ebde1f4699f498",
                "makerAmount": "1",
                "takerAmount": "1000000000000000",
                "takerTokenFeeAmount": "0",
                "maker": "0x56eb0ad2dc746540fab5c02478b31e2aa9ddc38c",
                "taker": "0x0000000000000000000000000000000000000000",
                "sender": "0x0000000000000000000000000000000000000000",
                "feeRecipient": "0x0000000000000000000000000000000000000000",
                "pool": "0x0000000000000000000000000000000000000000000000000000000000000000",
                "expiry": "1614962196",
                "salt": "12213086377898595950156985364849399280012820680806192474044093697918905947103",
                "signature": {
                    "signatureType": 3,
                    "v": 28,
                    "r": "0x828110c66e0b5435ad2b2a4479f2ca5ea6afaf2846a7d5218aeb4e7f2519ce7b",
                    "s": "0x06360f6d55d86006c090855f5467925707788bdfb4bfb7aae7304f0504c93ba0"
                }
            },
            "metaData": {
                "orderHash": "0xa3aeb3b1aed8438c0398ba070bf4f365378ef74e4fc966f09df87854484b0012",
                "remainingFillableTakerAmount": "1000000000000000",
                "state": "ADDED"
            }
        }
    ],
    "requestId": "123e4567-e89b-12d3-a456-426655440000"
}
```


```sh
	Type	Description
makerToken	address	The ERC20 token the maker is selling and the maker is selling to the taker. [required]
takerToken	address	The ERC20 token the taker is selling and the taker is selling to the maker. [required]
makerAmount	uint128	The amount of makerToken being sold by the maker. [required]
takerAmount	uint128	The amount of takerToken being sold by the taker. [required]
maker	address	The address of the maker, and signer, of this order. [required]
taker	address	Allowed taker address. Set to zero to allow any taker. [optional; default 0]
txOrigin	address	The allowed address of the EOA that submitted the Ethereum transaction. This must be set. Multiple addresses are supported via registerAllowedRfqOrigins. [required]
pool	bytes32	The staking pool to attribute the 0x protocol fee from this order. Set to zero to attribute to the default pool, not owned by anyone. [optional; default 0]
expiry	uint64	The Unix timestamp in seconds when this order expires. [required]
salt	uint256	Arbitrary number to enforce uniqueness of the order hash. [required]
```