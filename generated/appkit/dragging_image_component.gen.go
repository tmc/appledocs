// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DraggingImageComponent] class.
var draggingImageComponentClass = _DraggingImageComponentClass{objc.GetClass("NSDraggingImageComponent")}

type _DraggingImageComponentClass struct {
	class objc.Class
}

// An interface definition for the [DraggingImageComponent] class.
type IDraggingImageComponent interface {
	objectivec.IObject
}

// A single object in a dragging item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent

type DraggingImageComponent struct {
	objectivec.Object
}

// DraggingImageComponentFrom constructs a [DraggingImageComponent] from an unsafe.Pointer.
//
// A single object in a dragging item.
func DraggingImageComponentFrom(ptr unsafe.Pointer) DraggingImageComponent {
	return DraggingImageComponent{objectivec.Object{objc.ID(ptr)}}
}



