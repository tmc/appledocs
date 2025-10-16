
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [flushWindowIfNeeded] class.
var flushWindowIfNeededClass _flushWindowIfNeededClass

func init() {
	flushWindowIfNeededClass = _flushWindowIfNeededClass{objc.GetClass("flushWindowIfNeeded")}
}

type _flushWindowIfNeededClass struct {
	objc.Class
}

// An interface definition for the [flushWindowIfNeeded] class.
type IflushWindowIfNeeded interface {
	ID() objc.ID
}

type flushWindowIfNeeded struct {
	id objc.ID
}

func flushWindowIfNeededFrom(ptr unsafe.Pointer) flushWindowIfNeeded {
	return flushWindowIfNeeded{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ flushWindowIfNeeded) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _flushWindowIfNeededClass) Alloc() flushWindowIfNeeded {
	rv := objc.Send[flushWindowIfNeeded](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _flushWindowIfNeededClass) New() flushWindowIfNeeded {
	rv := objc.Send[flushWindowIfNeeded](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewflushWindowIfNeeded creates and returns a new initialized instance.
func NewflushWindowIfNeeded() flushWindowIfNeeded {
	return flushWindowIfNeededClass.New()
}

// Init initializes the instance.
func (f_ flushWindowIfNeeded) Init() flushWindowIfNeeded {
	rv := objc.Send[flushWindowIfNeeded](f_.ID(), selInit)
	return rv
}
