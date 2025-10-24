
// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

// Package inputmethodkit provides Go bindings for the InputMethodKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to InputMethodKit without requiring cgo.
package inputmethodkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/InputMethodKit.framework/InputMethodKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

