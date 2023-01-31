- Generate abi and binary files
solc --bin --abi contract/token.sol -o build

- Generate go code
abigen --bin=build/ERC20Token.bin --abi=build/ERC20Token.abi --pkg=ERC20Token --out=gen/ERC20Token.go