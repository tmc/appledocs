// Code generated from Apple documentation for PermissionKit. DO NOT EDIT.

// Package permissionkit provides Go bindings for the PermissionKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PermissionKit without requiring cgo.
package permissionkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PermissionKit.framework/PermissionKit"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
