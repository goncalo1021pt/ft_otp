package main

import (
	"flag"
	"fmt"
)

// parseArgs parses command-line arguments and returns which mode to run
// Returns: generateFile, keyFile, error
func parseArgs() (string, string, error) {
	generateKey := flag.String("g", "", "generate key from hex file")
	keyFile := flag.String("k", "", "generate OTP from key file")
	
	flag.Parse()

	// Check that exactly one flag is provided
	if (*generateKey != "" && *keyFile != "") || (*generateKey == "" && *keyFile == "") {
		return "", "", fmt.Errorf("use either -g or -k flag")
	}

	return *generateKey, *keyFile, nil
}

// validateHexKey reads a file and validates it contains at least 64 hex characters
// Returns the decoded bytes and an error if validation fails
func validateHexKey(filename string) ([]byte, error) {
	// TODO: Read file content
	// TODO: Trim whitespace
	// TODO: Check if valid hex
	// TODO: Check length >= 64
	// TODO: Return decoded bytes
	
	return nil, fmt.Errorf("not implemented yet")
}
