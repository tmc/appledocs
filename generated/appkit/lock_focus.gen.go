
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [lockFocus] class.
var lockFocusClass _lockFocusClass

func init() {
	lockFocusClass = _lockFocusClass{objc.GetClass("lockFocus")}
}

type _lockFocusClass struct {
	objc.Class
}

// An interface definition for the [lockFocus] class.
type IlockFocus interface {
	ID() objc.ID
}

type lockFocus struct {
	id objc.ID
}

func lockFocusFrom(ptr unsafe.Pointer) lockFocus {
	return lockFocus{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ lockFocus) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _lockFocusClass) Alloc() lockFocus {
	rv := objc.Send[lockFocus](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _lockFocusClass) New() lockFocus {
	rv := objc.Send[lockFocus](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlockFocus creates and returns a new initialized instance.
func NewlockFocus() lockFocus {
	return lockFocusClass.New()
}

// Init initializes the instance.
func (l_ lockFocus) Init() lockFocus {
	rv := objc.Send[lockFocus](l_.ID(), selInit)
	return rv
}
