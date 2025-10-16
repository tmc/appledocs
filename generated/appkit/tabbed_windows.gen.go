
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tabbedWindows] class.
var tabbedWindowsClass _tabbedWindowsClass

func init() {
	tabbedWindowsClass = _tabbedWindowsClass{objc.GetClass("tabbedWindows")}
}

type _tabbedWindowsClass struct {
	objc.Class
}

// An interface definition for the [tabbedWindows] class.
type ItabbedWindows interface {
	ID() objc.ID
}

type tabbedWindows struct {
	id objc.ID
}

func tabbedWindowsFrom(ptr unsafe.Pointer) tabbedWindows {
	return tabbedWindows{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tabbedWindows) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tabbedWindowsClass) Alloc() tabbedWindows {
	rv := objc.Send[tabbedWindows](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tabbedWindowsClass) New() tabbedWindows {
	rv := objc.Send[tabbedWindows](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtabbedWindows creates and returns a new initialized instance.
func NewtabbedWindows() tabbedWindows {
	return tabbedWindowsClass.New()
}

// Init initializes the instance.
func (t_ tabbedWindows) Init() tabbedWindows {
	rv := objc.Send[tabbedWindows](t_.ID(), selInit)
	return rv
}
