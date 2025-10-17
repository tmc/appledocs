// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RunningApplication] class.
var runningApplicationClass = _RunningApplicationClass{objc.GetClass("NSRunningApplication")}

type _RunningApplicationClass struct {
	class objc.Class
}

// An object that can manipulate and provide information for a single instance of an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunningApplication

type RunningApplication struct {
	objectivec.Object
}

// RunningApplicationFrom constructs a [RunningApplication] from an unsafe.Pointer.
//
// An object that can manipulate and provide information for a single instance of an app.
func RunningApplicationFrom(ptr unsafe.Pointer) RunningApplication {
	return RunningApplication{objectivec.Object{objc.ID(ptr)}}
}



