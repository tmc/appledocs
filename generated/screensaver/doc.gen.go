// Code generated from Apple documentation for ScreenSaver. DO NOT EDIT.

// Package screensaver provides Go bindings for the ScreenSaver framework.
//
// Animate screen savers, and interact with the screen saver infrastructure. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ScreenSaver without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver
package screensaver

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ScreenSaver.framework/ScreenSaver"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


