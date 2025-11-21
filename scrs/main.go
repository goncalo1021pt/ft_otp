package main

import (
	"fmt"
	"os"
)

func main() {
	generateFile, keyFile, qrFlag, webFlag, err := parseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "./ft_otp: error: %v\n", err)
		os.Exit(1)
	}

	if webFlag {
		if err := handleWebMode(); err != nil {
			fmt.Fprintf(os.Stderr, "./ft_otp: error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if generateFile != "" {
		if err := handleGenerateMode(generateFile, qrFlag); err != nil {
			fmt.Fprintf(os.Stderr, "./ft_otp: error: %v\n", err)
			os.Exit(1)
		}

	}

	if keyFile != "" {
		if err := handleOTPMode(keyFile); err != nil {
			fmt.Fprintf(os.Stderr, "./ft_otp: error: %v\n", err)
			os.Exit(1)
		}
	}
}
