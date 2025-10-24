// Code generated from Apple documentation for AdSupport. DO NOT EDIT.

// Package adsupport provides Go bindings for the AdSupport framework.
//
// Provide apps with access to an advertising identifier.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AdSupport without requiring cgo.
//
// See: https://developer.apple.com/documentation/AdSupport
package adsupport

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AdSupport.framework/AdSupport"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

