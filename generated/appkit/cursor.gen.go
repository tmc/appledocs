
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Cursor] class.
var CursorClass _CursorClass

func init() {
	CursorClass = _CursorClass{objc.GetClass("NSCursor")}
}

type _CursorClass struct {
	objc.Class
}

// An interface definition for the [Cursor] class.
type ICursor interface {
	ID() objc.ID
}

type Cursor struct {
	id objc.ID
}

func CursorFrom(ptr unsafe.Pointer) Cursor {
	return Cursor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ Cursor) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CursorClass) Alloc() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CursorClass) New() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCursor creates and returns a new initialized instance.
func NewCursor() Cursor {
	return CursorClass.New()
}

// Init initializes the instance.
func (c_ Cursor) Init() Cursor {
	rv := objc.Send[Cursor](c_.ID(), selInit)
	return rv
}
