// Code generated from Apple documentation for GameSave. DO NOT EDIT.

// Package gamesave provides Go bindings for the GameSave framework.
//
// Store and sync your application’s save files in iCloud. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to GameSave without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/GameSave
package gamesave

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/GameSave.framework/GameSave"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


