import { deployContract } from "./deploy-utils";


async function runContract(){
    const contract = await  deployContract("AnyConsumerFeed")
    const data = await contract.requestVolumeData()
    console.log(data.toString())
}

runContract().then(()=>{
    
}).catch((error)=>{
    console.log(error)
})