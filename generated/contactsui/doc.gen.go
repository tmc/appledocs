
// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

// Package contactsui provides Go bindings for the ContactsUI framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ContactsUI without requiring cgo.
package contactsui

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ContactsUI.framework/ContactsUI"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

