
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [areCursorRectsEnabled] class.
var areCursorRectsEnabledClass _areCursorRectsEnabledClass

func init() {
	areCursorRectsEnabledClass = _areCursorRectsEnabledClass{objc.GetClass("areCursorRectsEnabled")}
}

type _areCursorRectsEnabledClass struct {
	objc.Class
}

// An interface definition for the [areCursorRectsEnabled] class.
type IareCursorRectsEnabled interface {
	ID() objc.ID
}

type areCursorRectsEnabled struct {
	id objc.ID
}

func areCursorRectsEnabledFrom(ptr unsafe.Pointer) areCursorRectsEnabled {
	return areCursorRectsEnabled{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ areCursorRectsEnabled) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _areCursorRectsEnabledClass) Alloc() areCursorRectsEnabled {
	rv := objc.Send[areCursorRectsEnabled](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _areCursorRectsEnabledClass) New() areCursorRectsEnabled {
	rv := objc.Send[areCursorRectsEnabled](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewareCursorRectsEnabled creates and returns a new initialized instance.
func NewareCursorRectsEnabled() areCursorRectsEnabled {
	return areCursorRectsEnabledClass.New()
}

// Init initializes the instance.
func (a_ areCursorRectsEnabled) Init() areCursorRectsEnabled {
	rv := objc.Send[areCursorRectsEnabled](a_.ID(), selInit)
	return rv
}
