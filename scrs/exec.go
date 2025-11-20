package main

import (
	"fmt"
)

// handleGenerateMode validates hex key and stores it encrypted
func handleGenerateMode(filename string) error {
	// TODO: Validate hex key from file
	// TODO: Store encrypted key
	// TODO: Print success message
	
	return fmt.Errorf("not implemented yet")
}

// handleOTPMode generates and prints OTP
func handleOTPMode(keyFile string) error {
	// TODO: Read encrypted key file
	// TODO: Generate OTP
	// TODO: Print OTP
	
	return fmt.Errorf("not implemented yet")
}

// generateAndStoreKey takes hex key bytes and stores them encrypted in ft_otp.key
func generateAndStoreKey(keyBytes []byte, filename string) error {
	// TODO: Encrypt the key
	// TODO: Write to ft_otp.key file
	
	return fmt.Errorf("not implemented yet")
}

// generateOTP reads the encrypted key and generates a 6-digit OTP
func generateOTP(keyFile string) (string, error) {
	// TODO: Read encrypted key file
	// TODO: Decrypt the key
	// TODO: Get current time counter
	// TODO: Generate HOTP
	
	return "", fmt.Errorf("not implemented yet")
}
