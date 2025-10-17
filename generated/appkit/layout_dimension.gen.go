// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LayoutDimension] class.
var layoutDimensionClass = _LayoutDimensionClass{objc.GetClass("NSLayoutDimension")}

type _LayoutDimensionClass struct {
	class objc.Class
}

// An interface definition for the [LayoutDimension] class.
type ILayoutDimension interface {
	ILayoutAnchor
}

// A factory class for creating size-based layout constraint objects using a fluent API. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension

type LayoutDimension struct {
	LayoutAnchor
}

// LayoutDimensionFrom constructs a [LayoutDimension] from an unsafe.Pointer.
//
// A factory class for creating size-based layout constraint objects using a fluent API.
func LayoutDimensionFrom(ptr unsafe.Pointer) LayoutDimension {
	return LayoutDimension{
		LayoutAnchor: LayoutAnchorFrom(ptr),
	}
}



