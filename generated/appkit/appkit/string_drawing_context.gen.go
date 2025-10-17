// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StringDrawingContext] class.
var StringDrawingContextClass objc.Class

func init() {
	StringDrawingContextClass = objc.GetClass("NSStringDrawingContext")
}

type StringDrawingContext struct {
	objc.ID
}

func StringDrawingContextFrom(ptr unsafe.Pointer) StringDrawingContext {
	return StringDrawingContext{
		ID: objc.ID(ptr),
	}
}



