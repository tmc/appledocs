// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LayoutXAxisAnchor] class.
var layoutXAxisAnchorClass = _LayoutXAxisAnchorClass{objc.GetClass("NSLayoutXAxisAnchor")}

type _LayoutXAxisAnchorClass struct {
	class objc.Class
}

// A factory class for creating horizontal layout constraint objects using a fluent API. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor

type LayoutXAxisAnchor struct {
	LayoutAnchor
}

// LayoutXAxisAnchorFrom constructs a [LayoutXAxisAnchor] from an unsafe.Pointer.
//
// A factory class for creating horizontal layout constraint objects using a fluent API.
func LayoutXAxisAnchorFrom(ptr unsafe.Pointer) LayoutXAxisAnchor {
	return LayoutXAxisAnchor{
		LayoutAnchor: LayoutAnchorFrom(ptr),
	}
}



