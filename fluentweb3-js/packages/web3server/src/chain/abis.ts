export const publishAbi= [
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
          "name": "sender",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "string",
          "name": "content",
          "type": "string"
        }
      ],
      "name": "PublishEvent",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "topic",
          "type": "string"
        },
        {
          "internalType": "address",
          "name": "otherAddress",
          "type": "address"
        }
      ],
      "name": "GetAddressTopics",
      "outputs": [
        {
          "internalType": "string[]",
          "name": "",
          "type": "string[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "topic",
          "type": "string"
        }
      ],
      "name": "GetMyTopics",
      "outputs": [
        {
          "internalType": "string[]",
          "name": "",
          "type": "string[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "topic",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "message",
          "type": "string"
        }
      ],
      "name": "publish",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]