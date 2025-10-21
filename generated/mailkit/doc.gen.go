// Code generated from Apple documentation for MailKit. DO NOT EDIT.

// Package mailkit provides Go bindings for the MailKit framework.
//
// Secure, customize, and act on email messages that users send and receive. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MailKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/MailKit
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


