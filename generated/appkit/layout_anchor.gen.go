// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LayoutAnchor] class.
var layoutAnchorClass = _LayoutAnchorClass{objc.GetClass("NSLayoutAnchor")}

type _LayoutAnchorClass struct {
	class objc.Class
}

// A factory class for creating layout constraint objects using a fluent API. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor

type LayoutAnchor struct {
	objectivec.Object
}

// LayoutAnchorFrom constructs a [LayoutAnchor] from an unsafe.Pointer.
//
// A factory class for creating layout constraint objects using a fluent API.
func LayoutAnchorFrom(ptr unsafe.Pointer) LayoutAnchor {
	return LayoutAnchor{objectivec.Object{objc.ID(ptr)}}
}



