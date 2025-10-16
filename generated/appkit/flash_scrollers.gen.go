
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [flashScrollers] class.
var flashScrollersClass _flashScrollersClass

func init() {
	flashScrollersClass = _flashScrollersClass{objc.GetClass("flashScrollers")}
}

type _flashScrollersClass struct {
	objc.Class
}

// An interface definition for the [flashScrollers] class.
type IflashScrollers interface {
	ID() objc.ID
}

type flashScrollers struct {
	id objc.ID
}

func flashScrollersFrom(ptr unsafe.Pointer) flashScrollers {
	return flashScrollers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ flashScrollers) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _flashScrollersClass) Alloc() flashScrollers {
	rv := objc.Send[flashScrollers](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _flashScrollersClass) New() flashScrollers {
	rv := objc.Send[flashScrollers](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewflashScrollers creates and returns a new initialized instance.
func NewflashScrollers() flashScrollers {
	return flashScrollersClass.New()
}

// Init initializes the instance.
func (f_ flashScrollers) Init() flashScrollers {
	rv := objc.Send[flashScrollers](f_.ID(), selInit)
	return rv
}
