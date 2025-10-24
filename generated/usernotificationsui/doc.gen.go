
// Code generated from Apple documentation for UserNotificationsUI. DO NOT EDIT.

// Package usernotificationsui provides Go bindings for the UserNotificationsUI framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to UserNotificationsUI without requiring cgo.
package usernotificationsui

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/UserNotificationsUI.framework/UserNotificationsUI"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

