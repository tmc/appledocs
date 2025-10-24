// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

// Package cloudkit provides Go bindings for the CloudKit framework.
//
// Store structured app and user data in iCloud containers that all users of your app can share.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CloudKit without requiring cgo.
//
// See: https://developer.apple.com/documentation/CloudKit
package cloudkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CloudKit.framework/CloudKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

