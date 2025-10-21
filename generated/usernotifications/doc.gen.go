// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

// Package usernotifications provides Go bindings for the UserNotifications framework.
//
// Push user-facing notifications to the user’s device from a server, or generate them locally from your app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to UserNotifications without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications
package usernotifications

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/UserNotifications.framework/UserNotifications"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

