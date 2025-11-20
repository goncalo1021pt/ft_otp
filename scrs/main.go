package main

import (
	"fmt"
	"os"
)

func main() {
	generateFile, keyFile, err := parseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "./ft_otp: error: %v\n", err)
		os.Exit(1)
	}

	if generateFile != "" {
		if err := handleGenerateMode(generateFile); err != nil {
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