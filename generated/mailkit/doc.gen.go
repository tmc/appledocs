
// Code generated from Apple documentation for MailKit. DO NOT EDIT.

// Package mailkit provides Go bindings for the MailKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MailKit without requiring cgo.
package mailkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MailKit.framework/MailKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

