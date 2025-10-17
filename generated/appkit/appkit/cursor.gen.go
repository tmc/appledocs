// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Cursor] class.
var CursorClass objc.Class

func init() {
	CursorClass = objc.GetClass("NSCursor")
}

type Cursor struct {
	objc.ID
}

func CursorFrom(ptr unsafe.Pointer) Cursor {
	return Cursor{
		ID: objc.ID(ptr),
	}
}



