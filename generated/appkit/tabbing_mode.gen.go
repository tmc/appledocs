
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tabbingMode] class.
var tabbingModeClass _tabbingModeClass

func init() {
	tabbingModeClass = _tabbingModeClass{objc.GetClass("tabbingMode")}
}

type _tabbingModeClass struct {
	objc.Class
}

// An interface definition for the [tabbingMode] class.
type ItabbingMode interface {
	ID() objc.ID
}

type tabbingMode struct {
	id objc.ID
}

func tabbingModeFrom(ptr unsafe.Pointer) tabbingMode {
	return tabbingMode{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tabbingMode) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tabbingModeClass) Alloc() tabbingMode {
	rv := objc.Send[tabbingMode](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tabbingModeClass) New() tabbingMode {
	rv := objc.Send[tabbingMode](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtabbingMode creates and returns a new initialized instance.
func NewtabbingMode() tabbingMode {
	return tabbingModeClass.New()
}

// Init initializes the instance.
func (t_ tabbingMode) Init() tabbingMode {
	rv := objc.Send[tabbingMode](t_.ID(), selInit)
	return rv
}
