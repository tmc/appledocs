// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

// Package coretelephony provides Go bindings for the CoreTelephony framework.
//
// Access information about a user’s cellular service provider, such as its unique identifier and whether the carrier allows VoIP.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreTelephony without requiring cgo.
//
// See: https://developer.apple.com/documentation/CoreTelephony
package coretelephony

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreTelephony.framework/CoreTelephony"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

