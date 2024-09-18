package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/coreyvan/the_crypt/pkg/encipher"
)

func main() {
	// use the flag package to read in the key from a file and the plaintext as a string
	keyFile := flag.String("keyfile", "", "Path to the key file")
	ciphertextFile := flag.String("ciphertext", "", "File to decrypt")
	flag.Parse()

	// read the key from the file
	key, err := os.ReadFile(*keyFile)
	if err != nil {
		log.Fatalf("Failed to read key file: %v", err)
	}

	decoded, err := hex.DecodeString(string(key))
	if err != nil {
		log.Fatalf("Failed to decode key: %v", err)
	}

	ciphertext, err := os.ReadFile(*ciphertextFile)
	if err != nil {
		log.Fatalf("Failed to read ciphertext file: %v", err)
	}

	// use the encipher package to encrypt the plaintext
	decrypted, err := encipher.Decrypt(decoded, ciphertext)
	if err != nil {
		log.Fatalf("Failed to encrypt plaintext: %v", err)
	}

	fmt.Println(string(decrypted))
}
