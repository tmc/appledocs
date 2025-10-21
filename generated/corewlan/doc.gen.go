// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

// Package corewlan provides Go bindings for the CoreWLAN framework.
//
// Query AirPort interfaces and choose wireless networks. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreWLAN without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN
package corewlan

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreWLAN.framework/CoreWLAN"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

