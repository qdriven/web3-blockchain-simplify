# IPFS cluster

```
By default, IPFS and IPFS-Cluster use the following ports:

IPFS
4001 – Communication with other nodes
5001 – API server
8080 – Gateway server

IPFS-CLUSTER
9094 – HTTP API endpoint
9095 – IPFS proxy endpoint
9096 – Cluster swarm, used for communication between cluster nodes
```

## Setup

1. install ipfs
2. create private network
   1. swam-key
   2. ipfs bootstrap rm –all
   3. ipfs bootstrap add /ip4/192.168.10.1/tcp/4001/ipfs/QmQVvZEmvjhYgsyEC7NvMn8EWf131EcgTXFFJQYGSz4Y83
