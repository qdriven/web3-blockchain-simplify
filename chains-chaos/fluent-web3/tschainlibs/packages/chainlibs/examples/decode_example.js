import axios from "axios";

import InputDataDecoder from "ethereum-input-data-decoder";

console.log("test")

 function ethApiCall(
    url, method, params
) {
    url = "http://10.100.90.82:8545"
    return axios({
        method: "post",
        url: url,
        data: {
            jsonrpc: "2.0",
            method: method,
            params: params,
            id: 1,
        },
        headers: {"Content-Type": "application/json"},
        timeout: 2000,
    });
}

async function getTransactionDetail(url,
                 txHash
) {
    // console.log(req.body);
    const txnHash = txHash
    const rpcUrl = url
    let quorumTxn = {
        blockHash: "error",
        blockNumber: -1,
        from: "",
        gas: -1,
        gasPrice: -1,
        hash: "",
        input: "",
        nonce: -1,
        to: "",
        transactionIndex: -1,
        value: "",
        r: "",
        s: "",
        v: "",
    };

    try {
        const ethTxnByHash =  await ethApiCall(rpcUrl, "eth_getTransactionByHash", [
            txnHash,
        ]);
        quorumTxn["blockHash"] = ethTxnByHash.data.result.blockHash;
        quorumTxn["blockNumber"] = ethTxnByHash.data.result.blockNumber;
        quorumTxn["from"] = ethTxnByHash.data.result.from;
        quorumTxn["gas"] = ethTxnByHash.data.result.gas;
        quorumTxn["gasPrice"] = ethTxnByHash.data.result.gasPrice;
        quorumTxn["hash"] = ethTxnByHash.data.result.hash;
        quorumTxn["input"] = ethTxnByHash.data.result.input;
        quorumTxn["nonce"] = ethTxnByHash.data.result.nonce;
        quorumTxn["to"] = ethTxnByHash.data.result.to;
        quorumTxn["transactionIndex"] = ethTxnByHash.data.result.transactionIndex;
        quorumTxn["value"] = ethTxnByHash.data.result.value;
        quorumTxn["r"] = ethTxnByHash.data.result.r;
        quorumTxn["s"] = ethTxnByHash.data.result.s;
        quorumTxn["v"] = ethTxnByHash.data.result.v;
        return quorumTxn

    } catch (e) {
        console.error(e);
        console.error(
            "Node is unreachable. Ensure ports are open and client is running!"
        );

    } finally {
        console.log("do nothing")
    }
}

let result  = await getTransactionDetail("http://10.100.90.82:8545",
    "0x714679e57ce99a58e5265a96a9d8d9710163b2b28682312b192f1316f5f7cf58")

console.log(result["input"])


const abi = [
    {
        "inputs": [],
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "address",
                "name": "owner",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "bytes",
                "name": "hash",
                "type": "bytes"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "data",
                "type": "string"
            }
        ],
        "name": "AddEmissionDataEvent",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "address",
                "name": "user",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "bool",
                "name": "targetVal",
                "type": "bool"
            }
        ],
        "name": "AuthChangeEvent",
        "type": "event"
    },
    {
        "inputs": [],
        "name": "_owner",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address",
                "value": "0x0000000000000000000000000000000000000000"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "emissonData",
                "type": "string"
            },
            {
                "internalType": "bytes",
                "name": "hashedData",
                "type": "bytes",
                "value": "bytes"
            }
        ],
        "name": "addEmissionData",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "user",
                "type": "address",
                "value": "0x0000000000000000000000000000000000000000"
            }
        ],
        "name": "addToWhiteList",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getEmisionData",
        "outputs": [
            {
                "components": [
                    {
                        "internalType": "bytes",
                        "name": "hash",
                        "type": "bytes"
                    },
                    {
                        "internalType": "address",
                        "name": "owner",
                        "type": "address"
                    },
                    {
                        "internalType": "string",
                        "name": "data",
                        "type": "string"
                    }
                ],
                "internalType": "struct EmissionDataStore.EmissionData[]",
                "name": "",
                "type": "tuple[]",
                "value": "[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "user",
                "type": "address",
                "value": "0x0000000000000000000000000000000000000000"
            }
        ],
        "name": "removeFromWhiteList",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    }
]
const decoder = new InputDataDecoder(abi);

const decoded_result= decoder.decodeData(result["input"])
console.log(decoded_result)
console.log(decoded_result["inputs"][0])


