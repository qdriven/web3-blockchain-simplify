# docker pull remixproject/remix-ide:latest
# Docker 方式
docker run -p 8089:80 remixproject/remix-ide:latest

## npm 方式
npm install remixd -g
yarn global add remixd

remixd -s [path/ur/solidity/files] --remix-ide http://localhost:8080 #最后一个参数的意思是哪个服务访问后端，因为我们是映射到8080的Remix-project，所以填写http://localhost:8080


