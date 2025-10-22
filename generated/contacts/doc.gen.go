// Code generated from Apple documentation for Contacts. DO NOT EDIT.

// Package contacts provides Go bindings for the Contacts framework.
//
// Access the user’s contacts, and format and localize contact information.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Contacts without requiring cgo.

// Access the user’s contacts, and format and localize contact information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts

package contacts

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Contacts.framework/Contacts"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

