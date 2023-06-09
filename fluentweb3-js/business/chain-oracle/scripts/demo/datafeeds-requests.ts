
import { chainlink } from "hardhat";



async function getBtcEthPrice() {

    const contractAddress ="0x5fb1616F78dA7aFC9FF79e0371741a747D2a7F22";
    const btcEthPrice = await chainlink.getLatestPrice(contractAddress);
    console.log(btcEthPrice); 
}

getBtcEthPrice().catch((error)=>{
        console.error(error);
        process.exitCode = 1;
})

