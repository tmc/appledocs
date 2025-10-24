
// Code generated from Apple documentation for GameKit. DO NOT EDIT.

// Package gamekit provides Go bindings for the GameKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to GameKit without requiring cgo.
package gamekit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/GameKit.framework/GameKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

