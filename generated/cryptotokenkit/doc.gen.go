// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

// Package cryptotokenkit provides Go bindings for the CryptoTokenKit framework.
//
// Access security tokens and the cryptographic assets they store.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CryptoTokenKit without requiring cgo.

// Access security tokens and the cryptographic assets they store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit

package cryptotokenkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CryptoTokenKit.framework/CryptoTokenKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

