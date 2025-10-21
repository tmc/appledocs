// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

// Package fileprovider provides Go bindings for the FileProvider framework.
//
// An extension other apps use to access files and folders managed by your app and synced with a remote storage. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to FileProvider without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/FileProvider
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

