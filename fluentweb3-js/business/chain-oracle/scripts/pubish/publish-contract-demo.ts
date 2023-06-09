import { deployContract } from "./deploy-utils";


//0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6
async function publishDemo(){
    const publishContract = await  deployContract("PublishProxy")
    const jsonString = "{\"a\":\"b\"}"
    const publishResult = await publishContract.publish("a/b",jsonString)
    console.log(publishResult["hash"])
    console.log(JSON.stringify(publishResult))
    const result = await publishContract.GetMyTopics("a/b")
    console.log(JSON.stringify(result))
}


publishDemo().then((event)=>{
    console.log(event)
}).catch((error)=>{
    console.log(error)
})