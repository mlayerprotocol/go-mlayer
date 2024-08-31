
#!/bin/bash
mkdir -p ./contracts/evm/$1
abigen --abi $2 --pkg $1 --type $1Contract --out ./contracts/evm/$1/$1.go