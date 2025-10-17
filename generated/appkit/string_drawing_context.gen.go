// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StringDrawingContext] class.
var stringDrawingContextClass = _StringDrawingContextClass{objc.GetClass("NSStringDrawingContext")}

type _StringDrawingContextClass struct {
	class objc.Class
}

// An interface definition for the [StringDrawingContext] class.
type IStringDrawingContext interface {
	objectivec.IObject
}

// An object that manages metrics for drawing attributed strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingContext

type StringDrawingContext struct {
	objectivec.Object
}

// StringDrawingContextFrom constructs a [StringDrawingContext] from an unsafe.Pointer.
//
// An object that manages metrics for drawing attributed strings.
func StringDrawingContextFrom(ptr unsafe.Pointer) StringDrawingContext {
	return StringDrawingContext{objectivec.Object{objc.ID(ptr)}}
}



