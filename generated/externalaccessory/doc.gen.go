
// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

// Package externalaccessory provides Go bindings for the ExternalAccessory framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ExternalAccessory without requiring cgo.
package externalaccessory

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ExternalAccessory.framework/ExternalAccessory"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

