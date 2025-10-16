
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [autohidesScrollers] class.
var autohidesScrollersClass _autohidesScrollersClass

func init() {
	autohidesScrollersClass = _autohidesScrollersClass{objc.GetClass("autohidesScrollers")}
}

type _autohidesScrollersClass struct {
	objc.Class
}

// An interface definition for the [autohidesScrollers] class.
type IautohidesScrollers interface {
	ID() objc.ID
}

type autohidesScrollers struct {
	id objc.ID
}

func autohidesScrollersFrom(ptr unsafe.Pointer) autohidesScrollers {
	return autohidesScrollers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ autohidesScrollers) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _autohidesScrollersClass) Alloc() autohidesScrollers {
	rv := objc.Send[autohidesScrollers](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _autohidesScrollersClass) New() autohidesScrollers {
	rv := objc.Send[autohidesScrollers](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewautohidesScrollers creates and returns a new initialized instance.
func NewautohidesScrollers() autohidesScrollers {
	return autohidesScrollersClass.New()
}

// Init initializes the instance.
func (a_ autohidesScrollers) Init() autohidesScrollers {
	rv := objc.Send[autohidesScrollers](a_.ID(), selInit)
	return rv
}
