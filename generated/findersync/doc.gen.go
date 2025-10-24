
// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

// Package findersync provides Go bindings for the FinderSync framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to FinderSync without requiring cgo.
package findersync

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/FinderSync.framework/FinderSync"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

