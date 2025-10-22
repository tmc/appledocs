// Code generated from Apple documentation for Accounts. DO NOT EDIT.

// Package accounts provides Go bindings for the Accounts framework.
//
// Help users access and manage their external accounts from within your app, without requiring them to enter login credentials.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Accounts without requiring cgo.

// Help users access and manage their external accounts from within your app, without requiring them to enter login credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts

package accounts

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Accounts.framework/Accounts"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

