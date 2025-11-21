package main

import (
	"encoding/base32"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

// generateTOTPURI creates a TOTP URI
func generateTOTPURI(secret []byte, accountName, issuer string) string {
	base32Secret := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(secret)

	return fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s",
		issuer, accountName, base32Secret, issuer)
}

// generateQRCode generates a QR code PNG file from TOTP URI
func generateQRCode(uri, filename string) error {
	err := qrcode.WriteFile(uri, qrcode.High, 512, filename)
	if err != nil {
		return fmt.Errorf("failed to generate QR code: %v", err)
	}
	return nil
}

// overlayLogo overlays a logo image in the center of the QR code
func overlayLogo(qrPath, logoPath, outputPath string) error {

	qrFile, err := os.Open(qrPath)
	if err != nil {
		return fmt.Errorf("failed to open QR code: %v", err)
	}
	defer qrFile.Close()

	qrImg, err := png.Decode(qrFile)
	if err != nil {
		return fmt.Errorf("failed to decode QR code: %v", err)
	}

	logoFile, err := os.Open(logoPath)
	if err != nil {
		return fmt.Errorf("failed to open logo: %v", err)
	}
	defer logoFile.Close()

	logoImg, err := png.Decode(logoFile)
	if err != nil {
		return fmt.Errorf("failed to decode logo: %v", err)
	}

	bounds := qrImg.Bounds()
	output := image.NewRGBA(bounds)
	draw.Draw(output, bounds, qrImg, image.Point{}, draw.Src)

	qrSize := bounds.Dx()
	logoSize := qrSize / 5

	logoBounds := logoImg.Bounds()
	logoResized := image.NewRGBA(image.Rect(0, 0, logoSize, logoSize))

	for y := 0; y < logoSize; y++ {
		for x := 0; x < logoSize; x++ {
			srcX := x * logoBounds.Dx() / logoSize
			srcY := y * logoBounds.Dy() / logoSize
			logoResized.Set(x, y, logoImg.At(logoBounds.Min.X+srcX, logoBounds.Min.Y+srcY))
		}
	}

	centerX := (qrSize - logoSize) / 2
	centerY := (qrSize - logoSize) / 2
	logoRect := image.Rect(centerX, centerY, centerX+logoSize, centerY+logoSize)

	draw.Draw(output, logoRect, logoResized, image.Point{}, draw.Over)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, output)
	if err != nil {
		return fmt.Errorf("failed to encode PNG: %v", err)
	}

	return nil
}

// generateQRWithLogo generates a QR code with 42 logo overlay
func generateQRWithLogo(secret []byte, outputPath, logoPath string) error {
	uri := generateTOTPURI(secret, "user", "ft_otp")

	tempQR := "temp_qr.png"
	err := generateQRCode(uri, tempQR)
	if err != nil {
		return err
	}
	defer os.Remove(tempQR)

	if _, err := os.Stat(logoPath); err == nil {
		err = overlayLogo(tempQR, logoPath, outputPath)
		if err != nil {
			return err
		}
	} else {
		err = os.Rename(tempQR, outputPath)
		if err != nil {
			return fmt.Errorf("failed to save QR code: %v", err)
		}
	}

	return nil
}
