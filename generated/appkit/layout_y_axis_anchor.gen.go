// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LayoutYAxisAnchor] class.
var layoutYAxisAnchorClass = _LayoutYAxisAnchorClass{objc.GetClass("NSLayoutYAxisAnchor")}

type _LayoutYAxisAnchorClass struct {
	class objc.Class
}

// An interface definition for the [LayoutYAxisAnchor] class.
type ILayoutYAxisAnchor interface {
	ILayoutAnchor
}

// A factory class for creating vertical layout constraint objects using a fluent API. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor

type LayoutYAxisAnchor struct {
	LayoutAnchor
}

// LayoutYAxisAnchorFrom constructs a [LayoutYAxisAnchor] from an unsafe.Pointer.
//
// A factory class for creating vertical layout constraint objects using a fluent API.
func LayoutYAxisAnchorFrom(ptr unsafe.Pointer) LayoutYAxisAnchor {
	return LayoutYAxisAnchor{
		LayoutAnchor: LayoutAnchorFrom(ptr),
	}
}



