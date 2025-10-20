package main

import (
	"flag"
	"fmt"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	flag.Parse()

	fmt.Println("CryptoKit Framework Overview")
	fmt.Println("=============================")

	// Example 1: Framework Overview
	fmt.Println("\n1. Framework Overview:")
	fmt.Println("   CryptoKit provides secure cryptographic operations:")
	fmt.Println("   - Hashing (SHA-256, SHA-384, SHA-512)")
	fmt.Println("   - Message Authentication Codes (HMAC)")
	fmt.Println("   - Symmetric Encryption (AES-GCM, ChaCha20-Poly1305)")
	fmt.Println("   - Asymmetric Encryption (P-256, P-384, P-521, Curve25519)")
	fmt.Println("   - Digital Signatures (ECDSA, EdDSA)")
	fmt.Println("   - Key Agreement (ECDH, X25519)")

	// Example 2: Hashing
	fmt.Println("\n2. Hashing Algorithms:")
	hashAlgorithms := []string{
		"SHA256 - 256-bit secure hash",
		"SHA384 - 384-bit secure hash",
		"SHA512 - 512-bit secure hash",
		"MD5 (insecure) - Available for compatibility only",
	}
	for i, algo := range hashAlgorithms {
		fmt.Printf("   %d. %s\n", i+1, algo)
	}

	// Example 3: Symmetric Encryption
	fmt.Println("\n3. Symmetric Encryption:")
	symmetricAlgorithms := []string{
		"AES.GCM - AES in Galois/Counter Mode (recommended)",
		"ChaChaPoly - ChaCha20-Poly1305 authenticated encryption",
		"AES.CBC - AES in Cipher Block Chaining mode",
	}
	for i, algo := range symmetricAlgorithms {
		fmt.Printf("   %d. %s\n", i+1, algo)
	}

	// Example 4: Asymmetric Encryption
	fmt.Println("\n4. Asymmetric Encryption Curves:")
	asymmetricCurves := []string{
		"P256 - NIST P-256 elliptic curve",
		"P384 - NIST P-384 elliptic curve",
		"P521 - NIST P-521 elliptic curve",
		"Curve25519 - X25519 for key agreement",
	}
	for i, curve := range asymmetricCurves {
		fmt.Printf("   %d. %s\n", i+1, curve)
	}

	// Example 5: Key Types
	fmt.Println("\n5. Key Types:")
	keyTypes := map[string]string{
		"SymmetricKey":           "Symmetric encryption keys",
		"P256.Signing.PrivateKey": "ECDSA P-256 private signing key",
		"P256.KeyAgreement.PrivateKey": "ECDH P-256 key agreement key",
		"Curve25519.Signing.PrivateKey": "Ed25519 signing key",
		"Curve25519.KeyAgreement.PrivateKey": "X25519 key agreement key",
	}
	for keyType, description := range keyTypes {
		fmt.Printf("   %-38s: %s\n", keyType, description)
	}

	// Example 6: Message Authentication Codes
	fmt.Println("\n6. Message Authentication Codes (MACs):")
	macs := []string{
		"HMAC<SHA256> - HMAC with SHA-256",
		"HMAC<SHA384> - HMAC with SHA-384",
		"HMAC<SHA512> - HMAC with SHA-512",
	}
	for i, mac := range macs {
		fmt.Printf("   %d. %s\n", i+1, mac)
	}

	// Example 7: Digital Signatures
	fmt.Println("\n7. Digital Signature Algorithms:")
	signatures := []string{
		"ECDSA with P-256 - Elliptic Curve Digital Signature Algorithm",
		"ECDSA with P-384 - ECDSA with P-384 curve",
		"ECDSA with P-521 - ECDSA with P-521 curve",
		"EdDSA with Ed25519 - Edwards-curve Digital Signature Algorithm",
	}
	for i, sig := range signatures {
		fmt.Printf("   %d. %s\n", i+1, sig)
	}

	// Example 8: Key Agreement
	fmt.Println("\n8. Key Agreement Protocols:")
	keyAgreement := []string{
		"ECDH with P-256 - Elliptic Curve Diffie-Hellman",
		"ECDH with P-384 - ECDH with P-384 curve",
		"ECDH with P-521 - ECDH with P-521 curve",
		"X25519 - Curve25519 key agreement",
	}
	for i, proto := range keyAgreement {
		fmt.Printf("   %d. %s\n", i+1, proto)
	}

	// Example 9: Common Use Cases
	fmt.Println("\n9. Common Use Cases:")
	useCases := map[string]string{
		"Password Hashing":    "Use SHA-256 or SHA-512 with salt",
		"Data Integrity":      "Use HMAC to verify data hasn't been tampered with",
		"File Encryption":     "Use AES-GCM for authenticated encryption",
		"Secure Communication": "Use ECDH + AES-GCM for perfect forward secrecy",
		"Digital Signatures":  "Use ECDSA or EdDSA to sign documents/data",
		"API Authentication":  "Use HMAC for API request signing",
		"Token Generation":    "Use secure random bytes + hashing",
	}
	for useCase, description := range useCases {
		fmt.Printf("   %-25s: %s\n", useCase, description)
	}

	// Example 10: Security Best Practices
	fmt.Println("\n10. Security Best Practices:")
	bestPractices := []string{
		"Always use authenticated encryption (AES-GCM, ChaCha20-Poly1305)",
		"Use cryptographically secure random number generation",
		"Never reuse nonces with the same key",
		"Use appropriate key sizes (256-bit for symmetric, P-256+ for asymmetric)",
		"Don't implement your own crypto algorithms",
		"Use constant-time comparison for MACs and signatures",
		"Protect private keys with proper access controls",
		"Use key derivation functions (KDF) for password-based keys",
	}
	for i, practice := range bestPractices {
		fmt.Printf("   %d. %s\n", i+1, practice)
	}

	// Example 11: CryptoKit Advantages
	fmt.Println("\n11. CryptoKit Advantages:")
	advantages := []string{
		"Hardware acceleration on Apple Silicon",
		"Secure Enclave integration for key storage",
		"Modern, memory-safe API design",
		"Constant-time operations (timing attack resistant)",
		"Well-audited implementations",
		"Native Swift types with strong type safety",
		"Optimized for Apple platforms",
	}
	for i, advantage := range advantages {
		fmt.Printf("   %d. %s\n", i+1, advantage)
	}

	// Example 12: Limitations and Alternatives
	fmt.Println("\n12. Swift-Only Framework Limitations:")
	fmt.Println("   CryptoKit is a Swift-only framework with no Objective-C bridge.")
	fmt.Println("   ")
	fmt.Println("   Why no Go bindings:")
	fmt.Println("   - CryptoKit uses Swift-specific types (struct, enum, generics)")
	fmt.Println("   - No Objective-C classes to bind via purego")
	fmt.Println("   - Swift ABI is not stable for external language bindings")
	fmt.Println("   ")
	fmt.Println("   Alternatives for Go crypto:")
	fmt.Println("   - Go standard library: crypto/aes, crypto/sha256, crypto/hmac")
	fmt.Println("   - golang.org/x/crypto: Extended cryptographic algorithms")
	fmt.Println("   - crypto/ecdsa, crypto/ed25519: Elliptic curve signatures")
	fmt.Println("   - crypto/rand: Cryptographically secure random numbers")

	// Example 13: Go Crypto Equivalents
	fmt.Println("\n13. CryptoKit → Go crypto/* Equivalents:")
	equivalents := map[string]string{
		"SHA256":           "crypto/sha256",
		"SHA512":           "crypto/sha512",
		"HMAC":             "crypto/hmac",
		"AES-GCM":          "crypto/aes + crypto/cipher.NewGCM",
		"ChaCha20-Poly1305": "golang.org/x/crypto/chacha20poly1305",
		"P256 (ECDSA)":     "crypto/ecdsa + crypto/elliptic.P256",
		"Ed25519":          "crypto/ed25519",
		"X25519":           "golang.org/x/crypto/curve25519",
		"Random bytes":     "crypto/rand",
	}
	for cryptoKit, goEquiv := range equivalents {
		fmt.Printf("   %-20s → %s\n", cryptoKit, goEquiv)
	}

	// Example 14: Example Go Code Structure
	fmt.Println("\n14. Example Go Crypto Code:")
	fmt.Println("   import (")
	fmt.Println("       \"crypto/aes\"")
	fmt.Println("       \"crypto/cipher\"")
	fmt.Println("       \"crypto/rand\"")
	fmt.Println("       \"crypto/sha256\"")
	fmt.Println("   )")
	fmt.Println("   ")
	fmt.Println("   // Hash data")
	fmt.Println("   hash := sha256.Sum256([]byte(\"data\"))")
	fmt.Println("   ")
	fmt.Println("   // Encrypt with AES-GCM")
	fmt.Println("   key := make([]byte, 32) // 256-bit key")
	fmt.Println("   rand.Read(key)")
	fmt.Println("   ")
	fmt.Println("   block, _ := aes.NewCipher(key)")
	fmt.Println("   gcm, _ := cipher.NewGCM(block)")
	fmt.Println("   ")
	fmt.Println("   nonce := make([]byte, gcm.NonceSize())")
	fmt.Println("   rand.Read(nonce)")
	fmt.Println("   ")
	fmt.Println("   ciphertext := gcm.Seal(nil, nonce, plaintext, nil)")

	// Example 15: When to Use Each Approach
	fmt.Println("\n15. When to Use Each Approach:")
	fmt.Println("   Use CryptoKit (Swift):")
	fmt.Println("   - Building native macOS/iOS apps")
	fmt.Println("   - Need Secure Enclave integration")
	fmt.Println("   - Want hardware acceleration")
	fmt.Println("   - Prefer Swift's type safety")
	fmt.Println("   ")
	fmt.Println("   Use Go crypto:")
	fmt.Println("   - Building cross-platform tools")
	fmt.Println("   - Server-side applications")
	fmt.Println("   - Command-line utilities")
	fmt.Println("   - Need cgo-free deployment")

	// Example 16: Resources
	fmt.Println("\n16. Resources:")
	fmt.Println("   CryptoKit Documentation:")
	fmt.Println("   - https://developer.apple.com/documentation/cryptokit")
	fmt.Println("   ")
	fmt.Println("   Go Cryptography:")
	fmt.Println("   - https://pkg.go.dev/crypto")
	fmt.Println("   - https://pkg.go.dev/golang.org/x/crypto")
	fmt.Println("   ")
	fmt.Println("   Security Best Practices:")
	fmt.Println("   - https://owasp.org/www-project-cryptographic-storage-cheat-sheet/")

	fmt.Println("\n✓ CryptoKit framework overview completed!")
	fmt.Println("\nNote: This is an educational overview only.")
	fmt.Println("  CryptoKit is Swift-only and cannot be used from Go via purego.")
	fmt.Println("  For Go cryptography, use the standard library crypto packages.")
}
