// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

// Package gameplaykit provides Go bindings for the GameplayKit framework.
//
// Architect and organize your game logic. Incorporate common gameplay behaviors such as random number generation, artificial intelligence, pathfinding, and agent behavior. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to GameplayKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit
package gameplaykit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/GameplayKit.framework/GameplayKit"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


