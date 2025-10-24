// Code generated from Apple documentation for WebKit. DO NOT EDIT.

// Package webkit provides Go bindings for the WebKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to WebKit without requiring cgo.
package webkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/WebKit.framework/WebKit"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
