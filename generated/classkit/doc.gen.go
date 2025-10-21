// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

// Package classkit provides Go bindings for the ClassKit framework.
//
// Enable teachers to assign activities from your app’s content and to view student progress. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ClassKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit
package classkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ClassKit.framework/ClassKit"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


