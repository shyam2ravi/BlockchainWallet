package main

import (
	"fmt"
	"goblockchain/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println("Private Key: ", w.PrivateKeyStr())
	fmt.Println("Public Key: ", w.PublicKeyStr())
	fmt.Println("Blockchain Address: ", w.BlockchainAddress())

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)
	fmt.Printf("signature: %s\n", t.GenerateSignature())
}
