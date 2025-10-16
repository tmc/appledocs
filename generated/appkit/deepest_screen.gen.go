
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [deepestScreen] class.
var deepestScreenClass _deepestScreenClass

func init() {
	deepestScreenClass = _deepestScreenClass{objc.GetClass("deepestScreen")}
}

type _deepestScreenClass struct {
	objc.Class
}

// An interface definition for the [deepestScreen] class.
type IdeepestScreen interface {
	ID() objc.ID
}

type deepestScreen struct {
	id objc.ID
}

func deepestScreenFrom(ptr unsafe.Pointer) deepestScreen {
	return deepestScreen{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ deepestScreen) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _deepestScreenClass) Alloc() deepestScreen {
	rv := objc.Send[deepestScreen](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _deepestScreenClass) New() deepestScreen {
	rv := objc.Send[deepestScreen](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdeepestScreen creates and returns a new initialized instance.
func NewdeepestScreen() deepestScreen {
	return deepestScreenClass.New()
}

// Init initializes the instance.
func (d_ deepestScreen) Init() deepestScreen {
	rv := objc.Send[deepestScreen](d_.ID(), selInit)
	return rv
}
