
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [dividerThickness] class.
var dividerThicknessClass _dividerThicknessClass

func init() {
	dividerThicknessClass = _dividerThicknessClass{objc.GetClass("dividerThickness")}
}

type _dividerThicknessClass struct {
	objc.Class
}

// An interface definition for the [dividerThickness] class.
type IdividerThickness interface {
	ID() objc.ID
}

type dividerThickness struct {
	id objc.ID
}

func dividerThicknessFrom(ptr unsafe.Pointer) dividerThickness {
	return dividerThickness{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ dividerThickness) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _dividerThicknessClass) Alloc() dividerThickness {
	rv := objc.Send[dividerThickness](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _dividerThicknessClass) New() dividerThickness {
	rv := objc.Send[dividerThickness](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdividerThickness creates and returns a new initialized instance.
func NewdividerThickness() dividerThickness {
	return dividerThicknessClass.New()
}

// Init initializes the instance.
func (d_ dividerThickness) Init() dividerThickness {
	rv := objc.Send[dividerThickness](d_.ID(), selInit)
	return rv
}
