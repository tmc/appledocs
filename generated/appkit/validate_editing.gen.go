
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [validateEditing] class.
var validateEditingClass _validateEditingClass

func init() {
	validateEditingClass = _validateEditingClass{objc.GetClass("validateEditing")}
}

type _validateEditingClass struct {
	objc.Class
}

// An interface definition for the [validateEditing] class.
type IvalidateEditing interface {
	ID() objc.ID
}

type validateEditing struct {
	id objc.ID
}

func validateEditingFrom(ptr unsafe.Pointer) validateEditing {
	return validateEditing{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ validateEditing) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _validateEditingClass) Alloc() validateEditing {
	rv := objc.Send[validateEditing](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _validateEditingClass) New() validateEditing {
	rv := objc.Send[validateEditing](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewvalidateEditing creates and returns a new initialized instance.
func NewvalidateEditing() validateEditing {
	return validateEditingClass.New()
}

// Init initializes the instance.
func (v_ validateEditing) Init() validateEditing {
	rv := objc.Send[validateEditing](v_.ID(), selInit)
	return rv
}
