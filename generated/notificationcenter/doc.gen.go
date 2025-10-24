// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

// Package notificationcenter provides Go bindings for the NotificationCenter framework.
//
// Create and manage widgets for the Today view.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to NotificationCenter without requiring cgo.
//
// See: https://developer.apple.com/documentation/NotificationCenter
package notificationcenter

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/NotificationCenter.framework/NotificationCenter"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

