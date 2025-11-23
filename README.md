# ft_otp - Time-based One-Time Password Generator

A secure TOTP (Time-based One-Time Password) implementation in Go, featuring password-based encryption, QR code generation with custom logo overlay, and a modern web interface.

## Features

- 🔐 **Secure Key Storage**: Password-based encryption using PBKDF2 (100,000 iterations) + AES-256-GCM
- 🔑 **TOTP Generation**: RFC 6238 compliant time-based one-time passwords
- 📱 **QR Code Generation**: Create QR codes with custom 42 logo overlay for Google Authenticator
- 🌐 **Web Interface**: Beautiful, responsive web UI with auto-refreshing TOTP codes
- ⚡ **Cross-platform**: Pure Go implementation, no external dependencies required

## Installation

### Prerequisites

- Go 1.24.0 or higher
- Make (optional, for using Makefile)

### Build

```bash
# Clone the repository
git clone https://github.com/goncalo1021pt/ft_otp.git
cd ft_otp

# Build using Make
make

# Or build directly with Go
go build -o ft_otp scrs/*.go
```

## Usage

### Generate and Store a Key

Create an encrypted key file from a hexadecimal key (minimum 64 characters):

```bash
./ft_otp -g key.hex
```

You'll be prompted to enter and confirm an encryption password. The encrypted key will be saved to `ft_otp.key`.

### Generate and Store a Key with QR Code

Generate a key and create a QR code with 42 logo overlay:

```bash
./ft_otp -g key.hex --qr
```

This creates both `ft_otp.key` (encrypted key) and `qrcode.png` (QR code with logo).

### Generate a One-Time Password

Generate a TOTP code from your stored key:

```bash
./ft_otp -k ft_otp.key
```

You'll be prompted for your decryption password, then receive a 6-digit code valid for 30 seconds.

### Web Interface

Start the web server to view TOTP codes in your browser:

```bash
./ft_otp --web
```

Then open http://localhost:8080 in your browser. Features include:
- Live TOTP code display with countdown timer
- Visual progress bar (turns red in last 5 seconds)
- Automatic code refresh every 30 seconds
- QR code display for mobile setup

## Command-Line Options

| Flag | Description |
|------|-------------|
| `-g <file>` | Generate and store encrypted key from hex file |
| `-k <file>` | Generate TOTP code from encrypted key file |
| `--qr` | Generate QR code (use with `-g`) |
| `--web` | Start web interface |

## Project Structure

```
ft_otp/
├── scrs/
│   ├── main.go          # Entry point
│   ├── parser.go        # Command-line argument parsing
│   ├── generate.go      # Key generation and OTP generation
│   ├── totp.go          # TOTP algorithm implementation
│   ├── crypto.go        # Encryption/decryption functions
│   ├── qr.go            # QR code generation with logo overlay
│   └── web.go           # Web server implementation
├── assets/
│   ├── index.html       # Web interface HTML
│   └── 42_Logo.png      # 42 logo for QR overlay
├── tests/
│   └── key.hex          # Example hex key file
├── go.mod               # Go module definition
├── Makefile             # Build automation
└── README.md            # This file
```

## Security Features

### Encryption

- **Key Derivation**: PBKDF2 with SHA-256, 100,000 iterations
- **Encryption**: AES-256-GCM for authenticated encryption
- **Salt**: Random 32-byte salt generated per encryption
- **Nonce**: Random 12-byte nonce for GCM mode

### Key Storage

- Keys are never stored in plaintext
- File permissions set to 0600 (owner read/write only)
- Password validation (minimum 6 characters)
- Password confirmation on key generation

## Technical Details

### TOTP Algorithm

- Follows RFC 6238 specification
- HMAC-SHA1 based
- 30-second time step
- 6-digit output
- Dynamic truncation

### QR Code

- Uses High error correction level (30% tolerance)
- 512x512 pixel resolution
- Logo overlay at 20% size (1/5th of QR code)
- Compatible with Google Authenticator and similar apps

### Web Server

- Built-in Go HTTP server (net/http)
- Port 8080 by default
- RESTful API endpoints
- Single-page application architecture

## Testing

Verify TOTP codes using `oathtool`:

```bash
# Generate code with ft_otp
./ft_otp -k ft_otp.key

# Verify with oathtool (using original hex key)
oathtool --totp -b <hex_key>
```

Both should produce identical codes.

## Development

### Build Commands

```bash
make          # Build the project
make clean    # Remove built binaries
make re       # Rebuild from scratch
```

### Dependencies

- `golang.org/x/crypto` - Cryptographic functions (PBKDF2)
- `golang.org/x/term` - Terminal password input
- `github.com/skip2/go-qrcode` - QR code generation

## License

This project is part of the 42 Cybersecurity Piscine curriculum.

## Author

**Gonçalo Pereira**
- GitHub: [@goncalo1021pt](https://github.com/goncalo1021pt)

## Acknowledgments

- 42 School for the project specification
- RFC 4226 (HOTP) and RFC 6238 (TOTP) specifications
- Go community for excellent standard libraries
