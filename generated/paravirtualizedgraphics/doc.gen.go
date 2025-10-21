// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

// Package paravirtualizedgraphics provides Go bindings for the ParavirtualizedGraphics framework.
//
// Add graphics acceleration to your guest driver stack. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ParavirtualizedGraphics without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics
package paravirtualizedgraphics

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ParavirtualizedGraphics.framework/ParavirtualizedGraphics"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

