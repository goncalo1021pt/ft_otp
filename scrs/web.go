package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const webPort = ":8080"

var cachedKeyBytes []byte

// serveHomePage serves the HTML page
func serveHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/index.html")
}

// serveTOTPAPI returns current TOTP code as JSON
func serveTOTPAPI(w http.ResponseWriter, r *http.Request) {
	code, err := generateTOTP(cachedKeyBytes)
	if err != nil {
		respondJSON(w, map[string]string{"error": "Failed to generate TOTP"})
		return
	}

	now := time.Now().Unix()
	expiresAt := ((now / 30) + 1) * 30

	respondJSON(w, map[string]interface{}{
		"code":       code,
		"expires_at": expiresAt,
	})
}

// serveQRCode serves the QR code image
func serveQRCode(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("qrcode.png"); err == nil {
		http.ServeFile(w, r, "qrcode.png")
		return
	}

	uri := generateTOTPURI(cachedKeyBytes, "user", "ft_otp")
	err := generateQRCode(uri, "qrcode.png")
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, "qrcode.png")
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// handleWebMode starts a web server displaying TOTP codes and QR code
func handleWebMode() error {
	if _, err := os.Stat("ft_otp.key"); os.IsNotExist(err) {
		return fmt.Errorf("ft_otp.key not found. Generate a key first with -g")
	}

	password, err := promptPassword("Enter decryption password: ")
	if err != nil {
		return fmt.Errorf("failed to read password: %v", err)
	}

	encrypted, err := os.ReadFile("ft_otp.key")
	if err != nil {
		return fmt.Errorf("failed to read key file: %v", err)
	}

	cachedKeyBytes, err = decryptKey(encrypted, password)
	if err != nil {
		return fmt.Errorf("failed to decrypt key: %v", err)
	}

	fmt.Println("\nKey decrypted successfully!")

	http.HandleFunc("/", serveHomePage)
	http.HandleFunc("/api/totp", serveTOTPAPI)
	http.HandleFunc("/qrcode", serveQRCode)

	fmt.Printf("Starting web server on http://localhost%s\n", webPort)
	fmt.Println("Press Ctrl+C to stop")

	return http.ListenAndServe(webPort, nil)
}
