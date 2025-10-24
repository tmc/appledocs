// Code generated from Apple documentation for FileProviderUI. DO NOT EDIT.

// Package fileproviderui provides Go bindings for the FileProviderUI framework.
//
// Add actions to the document browser’s context menu.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to FileProviderUI without requiring cgo.
//
// See: https://developer.apple.com/documentation/FileProviderUI
package fileproviderui

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/FileProviderUI.framework/FileProviderUI"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

