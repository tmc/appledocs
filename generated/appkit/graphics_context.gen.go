// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GraphicsContext] class.
var GraphicsContextClass objc.Class

func init() {
	GraphicsContextClass = objc.GetClass("NSGraphicsContext")
}

type GraphicsContext struct {
	objc.ID
}

func GraphicsContextFrom(ptr unsafe.Pointer) GraphicsContext {
	return GraphicsContext{
		ID: objc.ID(ptr),
	}
}



