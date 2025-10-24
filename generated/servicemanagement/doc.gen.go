// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

// Package servicemanagement provides Go bindings for the ServiceManagement framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ServiceManagement without requiring cgo.
package servicemanagement

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ServiceManagement.framework/ServiceManagement"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
