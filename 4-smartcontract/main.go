package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	ERC20Token "github.com/web3-talk/4-smartcontract/gen"
)

var ganacheURL = "http://localhost:7545"

func main() {
	client, err := ethclient.Dial(ganacheURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	address := common.HexToAddress("0x52bb70735c91D7A19B32B949bAA80BEa15d538b9")

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("wei balance:", balance)

	// wei / 10^18
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println("ether balance: ", ethValue)

	// contract balance
	cAdd := common.HexToAddress("0xb53A19f611cB09A48BE2773d2D6561a429c21FC5")
	contract, err := ERC20Token.NewERC20Token(cAdd, client)
	if err != nil {
		fmt.Println(err)
	}

	add := common.HexToAddress("0x52bb70735c91D7A19B32B949bAA80BEa15d538b9")
	aTokenBalance, err := contract.BalanceOf(&bind.CallOpts{}, add)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("token balance: ", aTokenBalance)
}
