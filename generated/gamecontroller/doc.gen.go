// Code generated from Apple documentation for GameController. DO NOT EDIT.

// Package gamecontroller provides Go bindings for the GameController framework.
//
// Support hardware game controllers in your game. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to GameController without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController
package gamecontroller

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/GameController.framework/GameController"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

