# create pnpm vite proejct
pnpm create vite $1 -- --template vanilla-ts

## make packages
mkdir -p packages/$1
mkdir -p packages/$1/src
mkdir -p packages/$1/test
mkdir -p packages/$1/docs
mkdir -p packages/$1/examples

## create libs
cp scripts/templates/index.ts ./packages/$1/src
cp scripts/templates/tsconfig-tpl.json ./packages/$1/tsconfig.json
cp scripts/templates/package-cli-tpl.json ./packages/$1/package.json
