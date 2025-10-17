// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StatusBar] class.
var statusBarClass = _StatusBarClass{objc.GetClass("NSStatusBar")}

type _StatusBarClass struct {
	class objc.Class
}

// An interface definition for the [StatusBar] class.
type IStatusBar interface {
	objectivec.IObject
}

// An object that manages a collection of status items displayed within the system-wide menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusBar

type StatusBar struct {
	objectivec.Object
}

// StatusBarFrom constructs a [StatusBar] from an unsafe.Pointer.
//
// An object that manages a collection of status items displayed within the system-wide menu bar.
func StatusBarFrom(ptr unsafe.Pointer) StatusBar {
	return StatusBar{objectivec.Object{objc.ID(ptr)}}
}



