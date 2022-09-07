import Block from "./Block"
const createHash = require("crypto")

interface Transaction {
    from: string;
    to: string;
    amount: number;
}


export const calculateHash = (block: Block): string => {
    const data = JSON.stringify(block.data);
    const blockData =
        data +
        block.previousHash +
        block.timestamp.toISOString() +
        block.pow.toString();
    return createHash("sha256").update(blockData).digest("hex");
};


export default Transaction;