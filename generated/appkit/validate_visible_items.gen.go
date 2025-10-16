
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [validateVisibleItems] class.
var validateVisibleItemsClass _validateVisibleItemsClass

func init() {
	validateVisibleItemsClass = _validateVisibleItemsClass{objc.GetClass("validateVisibleItems")}
}

type _validateVisibleItemsClass struct {
	objc.Class
}

// An interface definition for the [validateVisibleItems] class.
type IvalidateVisibleItems interface {
	ID() objc.ID
}

type validateVisibleItems struct {
	id objc.ID
}

func validateVisibleItemsFrom(ptr unsafe.Pointer) validateVisibleItems {
	return validateVisibleItems{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ validateVisibleItems) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _validateVisibleItemsClass) Alloc() validateVisibleItems {
	rv := objc.Send[validateVisibleItems](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _validateVisibleItemsClass) New() validateVisibleItems {
	rv := objc.Send[validateVisibleItems](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewvalidateVisibleItems creates and returns a new initialized instance.
func NewvalidateVisibleItems() validateVisibleItems {
	return validateVisibleItemsClass.New()
}

// Init initializes the instance.
func (v_ validateVisibleItems) Init() validateVisibleItems {
	rv := objc.Send[validateVisibleItems](v_.ID(), selInit)
	return rv
}
