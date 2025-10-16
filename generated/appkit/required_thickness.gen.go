
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [requiredThickness] class.
var requiredThicknessClass _requiredThicknessClass

func init() {
	requiredThicknessClass = _requiredThicknessClass{objc.GetClass("requiredThickness")}
}

type _requiredThicknessClass struct {
	objc.Class
}

// An interface definition for the [requiredThickness] class.
type IrequiredThickness interface {
	ID() objc.ID
}

type requiredThickness struct {
	id objc.ID
}

func requiredThicknessFrom(ptr unsafe.Pointer) requiredThickness {
	return requiredThickness{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ requiredThickness) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _requiredThicknessClass) Alloc() requiredThickness {
	rv := objc.Send[requiredThickness](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _requiredThicknessClass) New() requiredThickness {
	rv := objc.Send[requiredThickness](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrequiredThickness creates and returns a new initialized instance.
func NewrequiredThickness() requiredThickness {
	return requiredThicknessClass.New()
}

// Init initializes the instance.
func (r_ requiredThickness) Init() requiredThickness {
	rv := objc.Send[requiredThickness](r_.ID(), selInit)
	return rv
}
