// Code generated from Apple documentation for AppleNewsFormat. DO NOT EDIT.

// Package applenewsformat provides Go bindings for the AppleNewsFormat framework.
//
// Get Apple News Format reference information, and create signature content for Apple News. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AppleNewsFormat without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/AppleNewsFormat
package applenewsformat

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AppleNewsFormat.framework/AppleNewsFormat"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


