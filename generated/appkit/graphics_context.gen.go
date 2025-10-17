// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GraphicsContext] class.
var graphicsContextClass = _GraphicsContextClass{objc.GetClass("NSGraphicsContext")}

type _GraphicsContextClass struct {
	class objc.Class
}

// An object that represents a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext

type GraphicsContext struct {
	objectivec.Object
}

// GraphicsContextFrom constructs a [GraphicsContext] from an unsafe.Pointer.
//
// An object that represents a graphics context.
func GraphicsContextFrom(ptr unsafe.Pointer) GraphicsContext {
	return GraphicsContext{objectivec.Object{objc.ID(ptr)}}
}



