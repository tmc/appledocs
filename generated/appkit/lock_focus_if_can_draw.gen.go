
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [lockFocusIfCanDraw] class.
var lockFocusIfCanDrawClass _lockFocusIfCanDrawClass

func init() {
	lockFocusIfCanDrawClass = _lockFocusIfCanDrawClass{objc.GetClass("lockFocusIfCanDraw")}
}

type _lockFocusIfCanDrawClass struct {
	objc.Class
}

// An interface definition for the [lockFocusIfCanDraw] class.
type IlockFocusIfCanDraw interface {
	ID() objc.ID
}

type lockFocusIfCanDraw struct {
	id objc.ID
}

func lockFocusIfCanDrawFrom(ptr unsafe.Pointer) lockFocusIfCanDraw {
	return lockFocusIfCanDraw{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ lockFocusIfCanDraw) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _lockFocusIfCanDrawClass) Alloc() lockFocusIfCanDraw {
	rv := objc.Send[lockFocusIfCanDraw](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _lockFocusIfCanDrawClass) New() lockFocusIfCanDraw {
	rv := objc.Send[lockFocusIfCanDraw](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlockFocusIfCanDraw creates and returns a new initialized instance.
func NewlockFocusIfCanDraw() lockFocusIfCanDraw {
	return lockFocusIfCanDrawClass.New()
}

// Init initializes the instance.
func (l_ lockFocusIfCanDraw) Init() lockFocusIfCanDraw {
	rv := objc.Send[lockFocusIfCanDraw](l_.ID(), selInit)
	return rv
}
