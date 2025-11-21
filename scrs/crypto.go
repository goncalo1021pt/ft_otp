package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"syscall"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/term"
)

// promptPassword prompts the user for a password without echoing
func promptPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

// deriveKeyFromPassword derives a 32-byte encryption key from password using PBKDF2
func deriveKeyFromPassword(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)
}

// encryptKey encrypts the key bytes using password-based AES-GCM
func encryptKey(plaintext []byte, password string) ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}

	key := deriveKeyFromPassword(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	result := make([]byte, 0, len(salt)+len(nonce)+len(ciphertext))
	result = append(result, salt...)
	result = append(result, nonce...)
	result = append(result, ciphertext...)

	return result, nil
}

// decryptKey decrypts the key bytes using password-based AES-GCM
func decryptKey(encrypted []byte, password string) ([]byte, error) {
	if len(encrypted) < 16 {
		return nil, fmt.Errorf("encrypted data too short")
	}
	salt := encrypted[:16]

	key := deriveKeyFromPassword(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(encrypted) < 16+nonceSize {
		return nil, fmt.Errorf("encrypted data too short")
	}
	nonce := encrypted[16 : 16+nonceSize]
	ciphertext := encrypted[16+nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed (wrong password?): %v", err)
	}

	return plaintext, nil
}
