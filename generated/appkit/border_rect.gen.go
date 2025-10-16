
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [borderRect] class.
var borderRectClass _borderRectClass

func init() {
	borderRectClass = _borderRectClass{objc.GetClass("borderRect")}
}

type _borderRectClass struct {
	objc.Class
}

// An interface definition for the [borderRect] class.
type IborderRect interface {
	ID() objc.ID
}

type borderRect struct {
	id objc.ID
}

func borderRectFrom(ptr unsafe.Pointer) borderRect {
	return borderRect{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ borderRect) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _borderRectClass) Alloc() borderRect {
	rv := objc.Send[borderRect](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _borderRectClass) New() borderRect {
	rv := objc.Send[borderRect](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewborderRect creates and returns a new initialized instance.
func NewborderRect() borderRect {
	return borderRectClass.New()
}

// Init initializes the instance.
func (b_ borderRect) Init() borderRect {
	rv := objc.Send[borderRect](b_.ID(), selInit)
	return rv
}
