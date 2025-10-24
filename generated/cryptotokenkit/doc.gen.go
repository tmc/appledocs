
// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

// Package cryptotokenkit provides Go bindings for the CryptoTokenKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CryptoTokenKit without requiring cgo.
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

