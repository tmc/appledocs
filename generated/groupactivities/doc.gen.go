// Code generated from Apple documentation for GroupActivities. DO NOT EDIT.

// Package groupactivities provides Go bindings for the GroupActivities framework.
//
// Create app-specific activities your users can share and experience together.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to GroupActivities without requiring cgo.

// Create app-specific activities your users can share and experience together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GroupActivities
package groupactivities

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/GroupActivities.framework/GroupActivities"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

