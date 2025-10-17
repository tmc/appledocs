// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StatusItem] class.
var statusItemClass = _StatusItemClass{objc.GetClass("NSStatusItem")}

type _StatusItemClass struct {
	class objc.Class
}

// An interface definition for the [StatusItem] class.
type IStatusItem interface {
	objectivec.IObject
}

// An individual element displayed in the system menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem

type StatusItem struct {
	objectivec.Object
}

// StatusItemFrom constructs a [StatusItem] from an unsafe.Pointer.
//
// An individual element displayed in the system menu bar.
func StatusItemFrom(ptr unsafe.Pointer) StatusItem {
	return StatusItem{objectivec.Object{objc.ID(ptr)}}
}



