var Web3 = require("web3");
const axios = require('axios');
var web3 = new Web3('https://matic-mumbai.chainstacklabs.com'); // your geth
const addresses = [
    "0xe9Ef427487EA3d7797a65077b1F8CeC554965734",
]

https://faucet.polygon.technology/
for (let i = 0, len = addresses.length; i < len; i++){
    // var account = web3.eth.accounts.create();
// console.log(account.address );
// console.log(account.privateKey  );
    axios.post('https://api.faucet.matic.network/transferTokens', {
        "network": "mumbai",
        "address": addresses[i] ,
        "token": "maticToken"})
        .then(function (response) {
            console.log(response);
        })
        .catch(function (error) {
            console.log(error);
        });
// const  address = account.address +"\t" +account.privateKey.substring(2)
// console.log(address );
}

