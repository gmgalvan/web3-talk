package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

	// Create an Ethereum wallet using ECDS in order to create, private, public keys and public address.
	createSimpleWallet()

	// Create wallet with password and keystore
	createWalletWithPassword()
}

func createSimpleWallet() {
	// Generate private key
	pvk, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println(err)
	}

	// Show Private Key Hex
	privateKeyBytes := crypto.FromECDSA(pvk)
	fmt.Println("private key hex: ", hexutil.Encode(privateKeyBytes))

	// Show Public Key Hex
	pubKey := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println("public key hex: ", hexutil.Encode(pubKey))

	// Generate Public Address
	pubAddr := crypto.PubkeyToAddress(pvk.PublicKey)

	// Show public address Hex
	fmt.Println("public address: ", pubAddr.Hex())
	fmt.Println("--")
}

func createWalletWithPassword() {
	// Creating a kesysstore a file containing an encrypted wallet private key
	keyFile := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "this-is-my-password"

	// NewAccount creates a new Ethereum account by taking a password as input and returning
	// the public address, encrypted private key, and a keystore file to securely store
	account, err := keyFile.NewAccount(password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ethereum account address derived from the key: ", account.Address)

	// Generate private, public keys and public Address from keystore
	b, err := ioutil.ReadFile(account.URL.Path)
	if err != nil {
		fmt.Println(err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	pubKey := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public Key Hex: ", hexutil.Encode(pubKey))

	pubAddr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	fmt.Println("Public Address: ", pubAddr.Hex())

}
