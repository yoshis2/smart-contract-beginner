package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	// this would be your generated smart contract bindings

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yoshis2/smart-contract-beginner/contracts"
)

// Ganacheのprivate keyの頭文字0xを外したキー
const PRIVATE_KEY = "22aabb811efca4e6f4748bd18a46b502fa85549df9fa07da649c0a148d7d5530"

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	address, tx, instance, err := contracts.DeployContracts(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())

	_, _ = instance, tx
}
