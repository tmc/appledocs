
// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

// Package fileprovider provides Go bindings for the FileProvider framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to FileProvider without requiring cgo.
package fileprovider

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/FileProvider.framework/FileProvider"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

