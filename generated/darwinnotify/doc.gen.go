
// Code generated from Apple documentation for DarwinNotify. DO NOT EDIT.

// Package darwinnotify provides Go bindings for the DarwinNotify framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to DarwinNotify without requiring cgo.
package darwinnotify

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/DarwinNotify.framework/DarwinNotify"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

