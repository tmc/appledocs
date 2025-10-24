// Code generated from Apple documentation for CryptoKit. DO NOT EDIT.

// Package cryptokit provides Go bindings for the CryptoKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CryptoKit without requiring cgo.
package cryptokit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CryptoKit.framework/CryptoKit"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
