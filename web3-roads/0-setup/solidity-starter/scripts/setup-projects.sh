# setup project

PROJECT_NAME=$1
mkdir -p ${PROJECT_NAME}
cd ${PROJECT_NAME}
npm install --save-dev hardhat
npx hardhat

## create typescript project

