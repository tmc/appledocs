// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

// Package spritekit provides Go bindings for the SpriteKit framework.
//
// Add high-performance 2D content with smooth animations to your app, or create a game with a high-level set of 2D game-based tools. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to SpriteKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit
package spritekit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/SpriteKit.framework/SpriteKit"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


