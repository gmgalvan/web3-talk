package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	// Address account generated on 2-wallet exercise
	address := "0x03128FB7c18738de6985f40Bc632d921b77EA83d"
	from := common.HexToAddress(address)

	// Address from ganache blockchain
	to := common.HexToAddress("0xeE11943C25AA03FFB8193CB223B9056a8F0A6432")

	// PendingNonceAt retrieves the current pending nonce associated with an account.
	//The pending nonce is the nonce that will be used for the next transaction.
	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1000000000000000000 wei
	amount := big.NewInt(1000000000000000000)

	// Get suggested gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// create new transaction
	tx := types.NewTransaction(nonce, to, amount, 21000, gasPrice, nil)

	//chainID is the unique identifier to specify the Ethereum network the client is connected to.
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt created wallet with password on 2-wallet exercise
	password := "this-is-my-password"

	// get keystore file on 2-wallet exercise
	b, err := ioutil.ReadFile("./wallet/UTC--2023-01-30T23-17-20.923280898Z--03128fb7c18738de6985f40bc632d921b77ea83d")
	if err != nil {
		fmt.Println(err)
	}

	// decrypt using password
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	// signs a given transaction with a private key and returns the signed transaction
	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
