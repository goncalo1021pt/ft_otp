package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

// generateTOTP generates a 6-digit TOTP code using the given secret key
func generateTOTP(secret []byte) (string, error) {
	counter := time.Now().Unix() / 30

	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, uint64(counter))

	h := hmac.New(sha1.New, secret)
	h.Write(counterBytes)
	hash := h.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F

	truncatedHash := binary.BigEndian.Uint32(hash[offset : offset+4])

	truncatedHash &= 0x7FFFFFFF

	code := truncatedHash % 1000000

	return fmt.Sprintf("%06d", code), nil
}

// generateOTP reads the key file and generates a 6-digit TOTP
func generateOTP(keyFile string) (string, error) {
	if keyFile != "ft_otp.key" {
		return "", fmt.Errorf("key file must be ft_otp.key")
	}

	encryptedKey, err := os.ReadFile(keyFile)
	if err != nil {
		return "", fmt.Errorf("cannot read key file: %v", err)
	}

	if len(encryptedKey) < 28 {
		return "", fmt.Errorf("invalid key file format (file too small)")
	}

	password, err := promptPassword("Enter password: ")
	if err != nil {
		return "", fmt.Errorf("failed to read password: %v", err)
	}

	keyBytes, err := decryptKey(encryptedKey, password)
	if err != nil {
		return "", fmt.Errorf("cannot decrypt key (wrong password or invalid file)")
	}

	otp, err := generateTOTP(keyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate OTP: %v", err)
	}

	return otp, nil
}

// handleOTPMode generates and prints OTP
func handleOTPMode(keyFile string) error {
	otp, err := generateOTP(keyFile)
	if err != nil {
		return err
	}

	fmt.Println(otp)
	return nil
}
