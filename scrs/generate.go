package main

import (
	"fmt"
	"os"
)

// generateAndStoreKey takes hex key bytes and stores them encrypted in ft_otp.key
func generateAndStoreKey(keyBytes []byte, filename string) error {
	password, err := promptPassword("Enter encryption password: ")
	if err != nil {
		return fmt.Errorf("failed to read password: %v", err)
	}

	if len(password) < 6 {
		return fmt.Errorf("password must be at least 6 characters")
	}

	confirm, err := promptPassword("Confirm password: ")
	if err != nil {
		return fmt.Errorf("failed to read password: %v", err)
	}

	if password != confirm {
		return fmt.Errorf("passwords do not match")
	}

	encrypted, err := encryptKey(keyBytes, password)
	if err != nil {
		return fmt.Errorf("failed to encrypt key: %v", err)
	}

	err = os.WriteFile(filename, encrypted, 0600)
	if err != nil {
		return fmt.Errorf("failed to write key file: %v", err)
	}

	return nil
}

// handleGenerateMode validates hex key and stores it encrypted
func handleGenerateMode(filename string, qrFlag bool) error {
	keyBytes, err := validateHexKey(filename)
	if err != nil {
		return err
	}

	err = generateAndStoreKey(keyBytes, "ft_otp.key")
	if err != nil {
		return err
	}

	fmt.Println("Key was successfully saved in ft_otp.key.")

	if qrFlag {
		err = generateQRWithLogo(keyBytes, "qrcode.png", "assets/42_Logo.png")
		if err != nil {
			return fmt.Errorf("failed to generate QR code: %v", err)
		}
		fmt.Println("QR code generated as qrcode.png")
	}

	return nil
}
