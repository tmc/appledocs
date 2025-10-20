# CryptoKit Framework Overview

This example provides an educational overview of Apple's CryptoKit framework and explains its relationship to Go cryptography.

## Overview

CryptoKit is Apple's modern cryptographic framework that provides:
- Secure hashing (SHA-256, SHA-384, SHA-512)
- Message Authentication Codes (HMAC)
- Symmetric encryption (AES-GCM, ChaCha20-Poly1305)
- Asymmetric cryptography (P-256, P-384, P-521, Curve25519)
- Digital signatures (ECDSA, EdDSA)
- Key agreement protocols (ECDH, X25519)

## Building

```bash
go build .
```

## Running

```bash
# Show comprehensive framework overview
./crypto-overview

# Show overview with e2e flag (no difference, framework is educational only)
./crypto-overview -e2e
```

## Features Demonstrated

1. **Framework Overview** - Core cryptographic capabilities
2. **Hashing Algorithms** - SHA family and legacy MD5
3. **Symmetric Encryption** - AES-GCM, ChaCha20-Poly1305
4. **Asymmetric Curves** - NIST P-curves and Curve25519
5. **Key Types** - Private/public key types
6. **Message Authentication** - HMAC variants
7. **Digital Signatures** - ECDSA and EdDSA
8. **Key Agreement** - ECDH and X25519
9. **Use Cases** - Real-world applications
10. **Security Best Practices** - Recommended patterns
11. **CryptoKit Advantages** - Hardware acceleration, Secure Enclave
12. **Swift-Only Limitations** - Why no Go bindings
13. **Go Equivalents** - Mapping CryptoKit → Go crypto
14. **Example Code** - Go crypto patterns
15. **When to Use Each** - CryptoKit vs Go crypto
16. **Resources** - Documentation links

## Important: Swift-Only Framework

**CryptoKit is a Swift-only framework with no Objective-C APIs.** This means:

- No Go bindings are available via purego/objc
- CryptoKit uses Swift-specific types (struct, enum, generics)
- Swift ABI is not stable for external language bindings
- The generated `cryptokit` package is a stub with no symbols

## Go Cryptography Alternatives

For Go applications, use the standard library crypto packages:

### Direct Equivalents

| CryptoKit | Go Package |
|-----------|------------|
| SHA256 | `crypto/sha256` |
| SHA512 | `crypto/sha512` |
| HMAC | `crypto/hmac` |
| AES-GCM | `crypto/aes` + `crypto/cipher` |
| ChaCha20-Poly1305 | `golang.org/x/crypto/chacha20poly1305` |
| P256 ECDSA | `crypto/ecdsa` + `crypto/elliptic` |
| Ed25519 | `crypto/ed25519` |
| X25519 | `golang.org/x/crypto/curve25519` |
| Random | `crypto/rand` |

### Example: AES-GCM Encryption in Go

```go
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
)

func encrypt(plaintext, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err := rand.Read(nonce); err != nil {
        return nil, err
    }

    // Seal appends the ciphertext to the nonce
    return gcm.Seal(nonce, nonce, plaintext, nil), nil
}
```

## When to Use Each Approach

### Use CryptoKit (Swift)
- Building native macOS/iOS/iPadOS apps
- Need Secure Enclave integration
- Want hardware acceleration on Apple Silicon
- Prefer Swift's type safety and modern API

### Use Go crypto
- Building cross-platform command-line tools
- Server-side applications
- Need cgo-free deployment
- Want consistent behavior across platforms

## Security Best Practices

1. **Always use authenticated encryption** - AES-GCM or ChaCha20-Poly1305
2. **Never reuse nonces** - Generate a new random nonce for each encryption
3. **Use appropriate key sizes** - 256-bit for symmetric, P-256+ for asymmetric
4. **Don't roll your own crypto** - Use established libraries
5. **Constant-time comparisons** - Prevent timing attacks on MACs/signatures
6. **Secure key storage** - Use OS keychain or key management services
7. **Key derivation** - Use PBKDF2, scrypt, or Argon2 for password-based keys
8. **Cryptographically secure randomness** - Always use `crypto/rand`

## CryptoKit Advantages

- Hardware acceleration via Apple's Cryptographic Acceleration
- Secure Enclave integration for private key protection
- Modern, memory-safe API design
- Constant-time operations (timing attack resistant)
- Well-audited implementations
- Native Swift types with strong type safety
- Optimized specifically for Apple platforms

## Common Cryptographic Operations

### Hashing
```go
import "crypto/sha256"

hash := sha256.Sum256([]byte("data"))
```

### HMAC
```go
import (
    "crypto/hmac"
    "crypto/sha256"
)

mac := hmac.New(sha256.New, key)
mac.Write(data)
tag := mac.Sum(nil)
```

### Digital Signatures (Ed25519)
```go
import "crypto/ed25519"

publicKey, privateKey, _ := ed25519.GenerateKey(nil)
signature := ed25519.Sign(privateKey, message)
valid := ed25519.Verify(publicKey, message, signature)
```

## Requirements

- macOS 10.15+ (for CryptoKit framework on macOS)
- iOS 13+ / iPadOS 13+ (for CryptoKit on mobile)
- Go 1.13+ (for Go crypto packages)

## Limitations

This example is educational only and demonstrates:
- What CryptoKit provides conceptually
- Why CryptoKit cannot be used from Go
- How to achieve equivalent functionality using Go's crypto packages

For actual cryptographic operations in Go, use the standard library packages shown above.

## References

- [CryptoKit Framework](https://developer.apple.com/documentation/cryptokit)
- [Go crypto package](https://pkg.go.dev/crypto)
- [Go x/crypto package](https://pkg.go.dev/golang.org/x/crypto)
- [OWASP Cryptographic Storage Cheat Sheet](https://owasp.org/www-project-cryptographic-storage-cheat-sheet/)
- [Go Cryptography Best Practices](https://go.dev/blog/crypto)
