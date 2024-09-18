package main

import (
	"encoding/hex"
	"flag"
	"log"
	"os"

	"github.com/coreyvan/the_crypt/pkg/encipher"
)

func main() {
	// use the flag package to read in the key from a file and the plaintext as a string
	keyFile := flag.String("keyfile", "", "Path to the key file")
	plaintext := flag.String("plaintext", "", "Plaintext to encrypt")
	outputFile := flag.String("output", "", "Path to the output file")
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

	// use the encipher package to encrypt the plaintext
	encryptedText, err := encipher.Encrypt(decoded, []byte(*plaintext))
	if err != nil {
		log.Fatalf("Failed to encrypt plaintext: %v", err)
	}

	// write the encrypted text to the output file
	err = os.WriteFile(*outputFile, []byte(encryptedText), 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}
}
