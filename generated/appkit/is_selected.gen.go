
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isSelected] class.
var isSelectedClass _isSelectedClass

func init() {
	isSelectedClass = _isSelectedClass{objc.GetClass("isSelected")}
}

type _isSelectedClass struct {
	objc.Class
}

// An interface definition for the [isSelected] class.
type IisSelected interface {
	ID() objc.ID
}

type isSelected struct {
	id objc.ID
}

func isSelectedFrom(ptr unsafe.Pointer) isSelected {
	return isSelected{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isSelected) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isSelectedClass) Alloc() isSelected {
	rv := objc.Send[isSelected](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isSelectedClass) New() isSelected {
	rv := objc.Send[isSelected](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisSelected creates and returns a new initialized instance.
func NewisSelected() isSelected {
	return isSelectedClass.New()
}

// Init initializes the instance.
func (i_ isSelected) Init() isSelected {
	rv := objc.Send[isSelected](i_.ID(), selInit)
	return rv
}
