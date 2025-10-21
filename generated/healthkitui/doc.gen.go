// Code generated from Apple documentation for HealthKitUI. DO NOT EDIT.

// Package healthkitui provides Go bindings for the HealthKitUI framework.
//
// Display user interface that enables a person to view and interact with their health data. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to HealthKitUI without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKitUI
package healthkitui

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/HealthKitUI.framework/HealthKitUI"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


