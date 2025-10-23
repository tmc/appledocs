// Code generated from Apple documentation for CoreData. DO NOT EDIT.

// Package coredata provides Go bindings for the CoreData framework.
//
// Persist or cache data on a single device, or sync data to multiple devices with CloudKit.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreData without requiring cgo.

// Persist or cache data on a single device, or sync data to multiple devices with CloudKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData
package coredata

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreData.framework/CoreData"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

