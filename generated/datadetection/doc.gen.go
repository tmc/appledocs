// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

// Package datadetection provides Go bindings for the DataDetection framework.
//
// Access and utilize common types of data that the data detection system matches.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to DataDetection without requiring cgo.

// Access and utilize common types of data that the data detection system matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection

package datadetection

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/DataDetection.framework/DataDetection"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

