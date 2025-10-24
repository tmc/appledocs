// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

// Package storekit provides Go bindings for the StoreKit framework.
//
// Support In-App Purchases and interactions with the App Store.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to StoreKit without requiring cgo.
//
// See: https://developer.apple.com/documentation/StoreKit
package storekit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/StoreKit.framework/StoreKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

