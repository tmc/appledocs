// Code generated from Apple documentation for OSLog. DO NOT EDIT.

// Package oslog provides Go bindings for the OSLog framework.
//
// A unified logging system for the reading of historical data. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to OSLog without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog
package oslog

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/OSLog.framework/OSLog"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


