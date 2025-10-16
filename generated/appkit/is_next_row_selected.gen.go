
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isNextRowSelected] class.
var isNextRowSelectedClass _isNextRowSelectedClass

func init() {
	isNextRowSelectedClass = _isNextRowSelectedClass{objc.GetClass("isNextRowSelected")}
}

type _isNextRowSelectedClass struct {
	objc.Class
}

// An interface definition for the [isNextRowSelected] class.
type IisNextRowSelected interface {
	ID() objc.ID
}

type isNextRowSelected struct {
	id objc.ID
}

func isNextRowSelectedFrom(ptr unsafe.Pointer) isNextRowSelected {
	return isNextRowSelected{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isNextRowSelected) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isNextRowSelectedClass) Alloc() isNextRowSelected {
	rv := objc.Send[isNextRowSelected](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isNextRowSelectedClass) New() isNextRowSelected {
	rv := objc.Send[isNextRowSelected](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisNextRowSelected creates and returns a new initialized instance.
func NewisNextRowSelected() isNextRowSelected {
	return isNextRowSelectedClass.New()
}

// Init initializes the instance.
func (i_ isNextRowSelected) Init() isNextRowSelected {
	rv := objc.Send[isNextRowSelected](i_.ID(), selInit)
	return rv
}
