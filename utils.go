package main

import (
	"encoding/hex"
)

// validateHex checks if the provided string is a valid hexadecimal string
func validateHex(input string) bool {
	_, err := hex.DecodeString(input)
	return err == nil
}

// validatePubKey checks if the provided public key is a valid hexadecimal and has a valid length
func validatePubKey(pubkey string) bool {
	if !validateHex(pubkey) {
		return false
	}
	// Check valid public key lengths (compressed or uncompressed)
	length := len(pubkey)
	return length == 66 || length == 130
}

func validatePubKeyHash(pkh string) bool {
	if !validateHex(pkh) {
		return false
	}
	// Check valid public key hash length
	return len(pkh) == 40
}
