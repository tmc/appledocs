
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [flushWindow] class.
var flushWindowClass _flushWindowClass

func init() {
	flushWindowClass = _flushWindowClass{objc.GetClass("flushWindow")}
}

type _flushWindowClass struct {
	objc.Class
}

// An interface definition for the [flushWindow] class.
type IflushWindow interface {
	ID() objc.ID
}

type flushWindow struct {
	id objc.ID
}

func flushWindowFrom(ptr unsafe.Pointer) flushWindow {
	return flushWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ flushWindow) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _flushWindowClass) Alloc() flushWindow {
	rv := objc.Send[flushWindow](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _flushWindowClass) New() flushWindow {
	rv := objc.Send[flushWindow](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewflushWindow creates and returns a new initialized instance.
func NewflushWindow() flushWindow {
	return flushWindowClass.New()
}

// Init initializes the instance.
func (f_ flushWindow) Init() flushWindow {
	rv := objc.Send[flushWindow](f_.ID(), selInit)
	return rv
}
