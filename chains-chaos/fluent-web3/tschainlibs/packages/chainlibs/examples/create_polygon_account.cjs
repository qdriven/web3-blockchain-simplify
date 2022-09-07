var Web3 = require("web3")
const axios = require("axios")
faucet_api_url = 'https://api.faucet.matic.network/transferTokens'
var web3 = new Web3("https://matic-mumbai.chainstacklabs.com")
for (let i = 0; i < 10; i++) {
    var account = web3.eth.accounts.create()
    axios.post(faucet_api_url, {
        "network": "mumbai",
        "address": account.address,
        "token": "maticToken"
    }).then(function (response) {
        console.log(response.data)
    }).catch(function (error) {
        console.log(account.address + " failed to get test token")
    });
    const addressInfo = account.address + ":" + account.privateKey.substring(2)
    console.log(addressInfo)
}