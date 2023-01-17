docker run -d --name ipfs_node_1 -e IPFS_SWARM_KEY_FILE=~/.ipfs/swarm.key -v ~/tmp/ipfs/node1/staging:/export -v ~/tmp/ipfs/node1/data:/data/ipfs -p 4001:4001 -p 4001:4001/udp -p 127.0.0.1:8080:8080 -p 127.0.0.1:5001:5001 ipfs/go-ipfs:latest