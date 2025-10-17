// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TouchBarItem] class.
var touchBarItemClass = _TouchBarItemClass{objc.GetClass("NSTouchBarItem")}

type _TouchBarItemClass struct {
	class objc.Class
}

// A UI control shown in the Touch Bar on supported models of MacBook Pro. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem

type TouchBarItem struct {
	objectivec.Object
}

// TouchBarItemFrom constructs a [TouchBarItem] from an unsafe.Pointer.
//
// A UI control shown in the Touch Bar on supported models of MacBook Pro.
func TouchBarItemFrom(ptr unsafe.Pointer) TouchBarItem {
	return TouchBarItem{objectivec.Object{objc.ID(ptr)}}
}



