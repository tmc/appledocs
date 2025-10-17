// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Cursor] class.
var cursorClass = _CursorClass{objc.GetClass("NSCursor")}

type _CursorClass struct {
	class objc.Class
}

// A pointer (also called a cursor). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor

type Cursor struct {
	objectivec.Object
}

// CursorFrom constructs a [Cursor] from an unsafe.Pointer.
//
// A pointer (also called a cursor).
func CursorFrom(ptr unsafe.Pointer) Cursor {
	return Cursor{objectivec.Object{objc.ID(ptr)}}
}



