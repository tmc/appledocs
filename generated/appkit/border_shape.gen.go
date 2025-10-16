
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [borderShape] class.
var borderShapeClass _borderShapeClass

func init() {
	borderShapeClass = _borderShapeClass{objc.GetClass("borderShape")}
}

type _borderShapeClass struct {
	objc.Class
}

// An interface definition for the [borderShape] class.
type IborderShape interface {
	ID() objc.ID
}

type borderShape struct {
	id objc.ID
}

func borderShapeFrom(ptr unsafe.Pointer) borderShape {
	return borderShape{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ borderShape) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _borderShapeClass) Alloc() borderShape {
	rv := objc.Send[borderShape](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _borderShapeClass) New() borderShape {
	rv := objc.Send[borderShape](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewborderShape creates and returns a new initialized instance.
func NewborderShape() borderShape {
	return borderShapeClass.New()
}

// Init initializes the instance.
func (b_ borderShape) Init() borderShape {
	rv := objc.Send[borderShape](b_.ID(), selInit)
	return rv
}
