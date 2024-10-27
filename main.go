package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

func main() {
	var inputInt int

	fmt.Print("CLI program to convert a hexadecimal publickey / publickeyhash to lockingbytecode\n")
	fmt.Print("Enter 0 to convert a pubkey, 1 to convert a pubkeyhash: ")
	fmt.Scanln(&inputInt)
	if inputInt == 0 {
		convertPk()
	} else if inputInt == 1 {
		convertPkh()
	} else {
		fmt.Print("Invalid input.")
	}
}

// hash160 performs a Hash160 (SHA-256 followed by RIPEMD-160) on the input bytes

func hash160(data []byte) []byte {
	// First, SHA-256 hash
	shaHash := sha256.Sum256(data)

	// Then, RIPEMD-160 hash
	ripemdHasher := ripemd160.New()
	ripemdHasher.Write(shaHash[:])
	ripemdHash := ripemdHasher.Sum(nil)

	return ripemdHash
}

func convertPk() {
	var pubkey string

	fmt.Print("Enter a public key: ")
	fmt.Scanln(&pubkey)

	if !validatePubKey(pubkey) {
		fmt.Println("Invalid public key format.")
		return
	}

	// Convert the hexadecimal string to bytes & perform Hash160
	inputBytes, _ := hex.DecodeString(pubkey)
	hash160Result := hash160(inputBytes)

	// Convert the result back to a hexadecimal string
	pkh := hex.EncodeToString(hash160Result)
	fmt.Printf("Your pkh is " + pkh + "\n")

	prefixP2pkh := "76a914"
	suffixP2pkh := "88ac"
	fmt.Printf("The lockingbytecode for your pkh is " + prefixP2pkh + pkh + suffixP2pkh + "\n")
}

func convertPkh() {
	var pkh string

	// Get the public key hash from the user
	for {
		fmt.Print("Enter a public key hash: ")
		fmt.Scanln(&pkh)

		if !validatePubKeyHash(pkh) {
			fmt.Println("Invalid publickeyhash format.")
		} else {
			break // Exit the loop if the input is valid
		}
	}

	prefixP2pkh := "76a914"
	suffixP2pkh := "88ac"
	fmt.Printf("The lockingbytecode for your pkh is " + prefixP2pkh + pkh + suffixP2pkh + "\n")
}
