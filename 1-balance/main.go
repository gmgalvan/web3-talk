package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ganacheURL = "http://localhost:7545"

func main() {
	// Connect to blockchain
	client, err := ethclient.DialContext(context.Background(), ganacheURL)
	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	defer client.Close()

	/*
		Exploring the client:
		- BlockByNumber
		- BalanceAt
	*/

	// Get block number
	getBlockNumber(client)

	// Get balance
	addr := "0x03128FB7c18738de6985f40Bc632d921b77EA83d"
	getBalance(client, addr)

}

func getBlockNumber(client *ethclient.Client) {
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("block number: ", block.Number())
}

func getBalance(client *ethclient.Client, addr string) {

	address := common.HexToAddress(addr)

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
}
