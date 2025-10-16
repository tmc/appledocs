
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isPreviousRowSelected] class.
var isPreviousRowSelectedClass _isPreviousRowSelectedClass

func init() {
	isPreviousRowSelectedClass = _isPreviousRowSelectedClass{objc.GetClass("isPreviousRowSelected")}
}

type _isPreviousRowSelectedClass struct {
	objc.Class
}

// An interface definition for the [isPreviousRowSelected] class.
type IisPreviousRowSelected interface {
	ID() objc.ID
}

type isPreviousRowSelected struct {
	id objc.ID
}

func isPreviousRowSelectedFrom(ptr unsafe.Pointer) isPreviousRowSelected {
	return isPreviousRowSelected{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isPreviousRowSelected) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isPreviousRowSelectedClass) Alloc() isPreviousRowSelected {
	rv := objc.Send[isPreviousRowSelected](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isPreviousRowSelectedClass) New() isPreviousRowSelected {
	rv := objc.Send[isPreviousRowSelected](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisPreviousRowSelected creates and returns a new initialized instance.
func NewisPreviousRowSelected() isPreviousRowSelected {
	return isPreviousRowSelectedClass.New()
}

// Init initializes the instance.
func (i_ isPreviousRowSelected) Init() isPreviousRowSelected {
	rv := objc.Send[isPreviousRowSelected](i_.ID(), selInit)
	return rv
}
