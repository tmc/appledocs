// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RunningApplication] class.
var RunningApplicationClass objc.Class

func init() {
	RunningApplicationClass = objc.GetClass("NSRunningApplication")
}

type RunningApplication struct {
	objc.ID
}

func RunningApplicationFrom(ptr unsafe.Pointer) RunningApplication {
	return RunningApplication{
		ID: objc.ID(ptr),
	}
}



