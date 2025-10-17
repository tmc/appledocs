// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LayoutGuide] class.
var layoutGuideClass = _LayoutGuideClass{objc.GetClass("NSLayoutGuide")}

type _LayoutGuideClass struct {
	class objc.Class
}

// A rectangular area that can interact with Auto Layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide

type LayoutGuide struct {
	objectivec.Object
}

// LayoutGuideFrom constructs a [LayoutGuide] from an unsafe.Pointer.
//
// A rectangular area that can interact with Auto Layout.
func LayoutGuideFrom(ptr unsafe.Pointer) LayoutGuide {
	return LayoutGuide{objectivec.Object{objc.ID(ptr)}}
}



