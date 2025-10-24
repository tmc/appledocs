// Code generated from Apple documentation for GSS. DO NOT EDIT.

// Package gss provides Go bindings for the GSS framework.
//
// Conduct secure, authenticated network transactions.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to GSS without requiring cgo.
//
// See: https://developer.apple.com/documentation/GSS
package gss

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/GSS.framework/GSS"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

