// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WindowTab] class.
var windowTabClass = _WindowTabClass{objc.GetClass("NSWindowTab")}

type _WindowTabClass struct {
	class objc.Class
}

// An interface definition for the [WindowTab] class.
type IWindowTab interface {
	objectivec.IObject
}

// A tab associated with a window that is part of a tabbing group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab

type WindowTab struct {
	objectivec.Object
}

// WindowTabFrom constructs a [WindowTab] from an unsafe.Pointer.
//
// A tab associated with a window that is part of a tabbing group.
func WindowTabFrom(ptr unsafe.Pointer) WindowTab {
	return WindowTab{objectivec.Object{objc.ID(ptr)}}
}



