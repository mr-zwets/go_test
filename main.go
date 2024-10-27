package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

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

	// Convert the hexadecimal string to bytes
	inputBytes, err := hex.DecodeString(pubkey)
	if err != nil {
		log.Fatalf("Failed to decode hex string: %v", err)
	}

	// Perform Hash160 (SHA-256 + RIPEMD-160)
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

		// check length of pkh
		pkhLength := len(pkh)
		if pkhLength != 40 {
			fmt.Println("Invalid input. Please enter a valid number.")
		} else {
			break // Exit the loop if the input is valid
		}
	}

	prefixP2pkh := "76a914"
	suffixP2pkh := "88ac"
	fmt.Printf("The lockingbytecode for your pkh is " + prefixP2pkh + pkh + suffixP2pkh + "\n")
}