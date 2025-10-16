
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [disableFlushWindow] class.
var disableFlushWindowClass _disableFlushWindowClass

func init() {
	disableFlushWindowClass = _disableFlushWindowClass{objc.GetClass("disableFlushWindow")}
}

type _disableFlushWindowClass struct {
	objc.Class
}

// An interface definition for the [disableFlushWindow] class.
type IdisableFlushWindow interface {
	ID() objc.ID
}

type disableFlushWindow struct {
	id objc.ID
}

func disableFlushWindowFrom(ptr unsafe.Pointer) disableFlushWindow {
	return disableFlushWindow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ disableFlushWindow) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _disableFlushWindowClass) Alloc() disableFlushWindow {
	rv := objc.Send[disableFlushWindow](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _disableFlushWindowClass) New() disableFlushWindow {
	rv := objc.Send[disableFlushWindow](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdisableFlushWindow creates and returns a new initialized instance.
func NewdisableFlushWindow() disableFlushWindow {
	return disableFlushWindowClass.New()
}

// Init initializes the instance.
func (d_ disableFlushWindow) Init() disableFlushWindow {
	rv := objc.Send[disableFlushWindow](d_.ID(), selInit)
	return rv
}
