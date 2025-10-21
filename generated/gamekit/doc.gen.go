// Code generated from Apple documentation for GameKit. DO NOT EDIT.

// Package gamekit provides Go bindings for the GameKit framework.
//
// Enable players to interact with friends, compare leaderboard ranks, earn achievements, and participate in multiplayer games. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to GameKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit
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

