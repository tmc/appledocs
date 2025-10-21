// Code generated from Apple documentation for Charts. DO NOT EDIT.

// Package charts provides Go bindings for the Charts framework.
//
// Construct and customize charts on every Apple platform. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Charts without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/Charts
package charts

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Charts.framework/Charts"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


