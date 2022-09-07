var Web3 = require("web3")
var web3 = new Web3("https://matic-mumbai.chainstacklabs.com")

addresses =[
    "0x3610136664105E94b4325b482b94Af8fC960Cf0A",
    "0xa06C2B5679DFcDbBAd04D9018a405B0DC67C38a7"
]

async  function getBalance(address){
    await web3.eth.getBalance(address).then(
        function(response){
            return response
        }
    )
}

for (let i = 0; i <addresses.length ; i++) {
    let result =   getBalance(addresses[i])
    console.log(addresses[i]+" balance is ",result)
}