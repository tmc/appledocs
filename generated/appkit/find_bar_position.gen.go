
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [findBarPosition] class.
var findBarPositionClass _findBarPositionClass

func init() {
	findBarPositionClass = _findBarPositionClass{objc.GetClass("findBarPosition")}
}

type _findBarPositionClass struct {
	objc.Class
}

// An interface definition for the [findBarPosition] class.
type IfindBarPosition interface {
	ID() objc.ID
}

type findBarPosition struct {
	id objc.ID
}

func findBarPositionFrom(ptr unsafe.Pointer) findBarPosition {
	return findBarPosition{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ findBarPosition) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _findBarPositionClass) Alloc() findBarPosition {
	rv := objc.Send[findBarPosition](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _findBarPositionClass) New() findBarPosition {
	rv := objc.Send[findBarPosition](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfindBarPosition creates and returns a new initialized instance.
func NewfindBarPosition() findBarPosition {
	return findBarPositionClass.New()
}

// Init initializes the instance.
func (f_ findBarPosition) Init() findBarPosition {
	rv := objc.Send[findBarPosition](f_.ID(), selInit)
	return rv
}
