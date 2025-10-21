// Code generated from Apple documentation for AppleArchive. DO NOT EDIT.

// Package applearchive provides Go bindings for the AppleArchive framework.
//
// Perform multithreaded lossless compression of directories, files, and data. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AppleArchive without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive
package applearchive

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AppleArchive.framework/AppleArchive"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


