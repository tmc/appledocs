// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

// Package addressbook provides Go bindings for the AddressBook framework.
//
// Access the centralized database for storing users’ contacts. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AddressBook without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook
package addressbook

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AddressBook.framework/AddressBook"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


