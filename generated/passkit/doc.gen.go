// Code generated from Apple documentation for PassKit. DO NOT EDIT.

// Package passkit provides Go bindings for the PassKit framework.
//
// Process Apple Pay payments in your app, and create and distribute passes for the Wallet app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PassKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit
package passkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PassKit.framework/PassKit"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


