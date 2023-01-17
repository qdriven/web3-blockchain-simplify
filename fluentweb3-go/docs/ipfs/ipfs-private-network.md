# IPFS private network

## setup

```shell
 go get -u github.com/Kubuxu/go-ipfs-swarm-key-gen/ipfs-swarm-key-gen  
```

## key generation

```shell
 ipfs-swarm-key-gen > ~/.ipfs/swarm.key  

```

## nodes

```shell
docker run -d --name ipfs_node_1 -e IPFS_SWARM_KEY_FILE=~/.ipfs/swarm.key -v ~/tmp/ipfs/node1/staging:/export -v ~/tmp/ipfs/node1/data:/data/ipfs -p 4001:4001 -p 4001:4001/udp -p 127.0.0.1:8080:8080 -p 127.0.0.1:5001:5001 ipfs/go-ipfs:latest
docker run -d --name ipfs_node_2 -e IPFS_SWARM_KEY_FILE=/Users/sixdays/tmp/ipfs/swarm.key -v /Users/sixdays/tmp/ipfs/node2/staging:/export -v /Users/sixdays/tmp/ ipfs/node2/data:/data/ipfs -p 4002:4001 -p 4002:4001/udp -p 127.0.0.1:8081:8080 -p 127.0.0.1:5002:5001 ipfs/go-ipfs:latest
```

```shell
docker exec ipfs_node_1 ipfs bootstrap rm all
docker exec ipfs_node_2 ipfs bootstrap rm all  
docker exec ipfs_node_1 ipfs id
docker exec ipfs_node_2 ipfs id  

  
```

address:
```shell
/ip4/172.17.0.3/tcp/4001/p2p/12D3KooWEVo8FqH8YUT1noXvca5hgSRWBRcDQomEcFY2zXwA7dbw
```

## reference

- [ipfs private network](https://mp.weixin.qq.com/s?__biz=MzU5NzUwODcyMw==&mid=2247485626&idx=1&sn=51328d315ee3236474ec81e03774cd6c&chksm=fe531fa6c92496b0b716c396c0a8163661373f58c3f59aa577b8696adf49d2796ee5aa08f98a&scene=0&xtrack=1)