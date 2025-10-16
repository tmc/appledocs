
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [enableFlushWindow] class.
var enableFlushWindowClass _enableFlushWindowClass

func init() {
	enableFlushWindowClass = _enableFlushWindowClass{objc.GetClass("enableFlushWindow")}
}

type _enableFlushWindowClass struct {
	objc.Class
}

// An interface definition for the [enableFlushWindow] class.
type IenableFlushWindow interface {
	ID() objc.ID
}

type enableFlushWindow struct {
	id objc.ID
}

func enableFlushWindowFrom(ptr unsafe.Pointer) enableFlushWindow {
	return enableFlushWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ enableFlushWindow) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _enableFlushWindowClass) Alloc() enableFlushWindow {
	rv := objc.Send[enableFlushWindow](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _enableFlushWindowClass) New() enableFlushWindow {
	rv := objc.Send[enableFlushWindow](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewenableFlushWindow creates and returns a new initialized instance.
func NewenableFlushWindow() enableFlushWindow {
	return enableFlushWindowClass.New()
}

// Init initializes the instance.
func (e_ enableFlushWindow) Init() enableFlushWindow {
	rv := objc.Send[enableFlushWindow](e_.ID(), selInit)
	return rv
}
