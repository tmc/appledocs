
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [fillColor] class.
var fillColorClass _fillColorClass

func init() {
	fillColorClass = _fillColorClass{objc.GetClass("fillColor")}
}

type _fillColorClass struct {
	objc.Class
}

// An interface definition for the [fillColor] class.
type IfillColor interface {
	ID() objc.ID
}

type fillColor struct {
	id objc.ID
}

func fillColorFrom(ptr unsafe.Pointer) fillColor {
	return fillColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ fillColor) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _fillColorClass) Alloc() fillColor {
	rv := objc.Send[fillColor](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _fillColorClass) New() fillColor {
	rv := objc.Send[fillColor](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfillColor creates and returns a new initialized instance.
func NewfillColor() fillColor {
	return fillColorClass.New()
}

// Init initializes the instance.
func (f_ fillColor) Init() fillColor {
	rv := objc.Send[fillColor](f_.ID(), selInit)
	return rv
}
