package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
)

// parseArgs parses command-line arguments and returns which mode to run.
// Returns: generateFile, keyFile, qrFlag, webFlag, error
func parseArgs() (string, string, bool, bool, error) {
	generateKey := flag.String("g", "", "generate key from hex file")
	keyFile := flag.String("k", "", "generate OTP from key file")
	qrFlag := flag.Bool("qr", false, "generate QR code for Authenticator app")
	webFlag := flag.Bool("web", false, "start web interface")

	flag.Parse()

	if *webFlag {
		return "", "", false, true, nil
	}

	if (*generateKey != "" && *keyFile != "") || (*generateKey == "" && *keyFile == "") {
		return "", "", false, false, fmt.Errorf("use either -g -web or -k flag")
	}

	return *generateKey, *keyFile, *qrFlag, false, nil
}

// validateHexKey reads a file and validates it contains at least 64 hex characters
func validateHexKey(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %v", err)
	}

	hexStr := strings.TrimSpace(string(content))

	if len(hexStr) < 64 {
		return nil, fmt.Errorf("key must be 64 hexadecimal characters")
	}

	keyBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("key must be 64 hexadecimal characters")
	}

	return keyBytes, nil
}
