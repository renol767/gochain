package main

import (
	"fmt"
	"wallet"
)

func main() {
	// myBlockchainAddress := "my_blockchain_address"
	// blockChain := NewBlockchain(myBlockchainAddress)

	// blockChain.AddTransaction("Ren", "Nol", 2.0)
	// blockChain.Mining()

	// blockChain.AddTransaction("Nol", "Ren", 8.0)
	// blockChain.AddTransaction("Nol", "Ren", 4.0)
	// blockChain.Mining()
	// blockChain.Print()

	// fmt.Printf("Ren %.1f\n", blockChain.CalculateTotalAmount("Ren"))
	// fmt.Printf("Nol %.1f\n", blockChain.CalculateTotalAmount("Nol"))
	// fmt.Printf("My Blockchain Address %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))

	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
}
