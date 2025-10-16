
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [makeTouchBar] class.
var makeTouchBarClass _makeTouchBarClass

func init() {
	makeTouchBarClass = _makeTouchBarClass{objc.GetClass("makeTouchBar")}
}

type _makeTouchBarClass struct {
	objc.Class
}

// An interface definition for the [makeTouchBar] class.
type ImakeTouchBar interface {
	ID() objc.ID
}

type makeTouchBar struct {
	id objc.ID
}

func makeTouchBarFrom(ptr unsafe.Pointer) makeTouchBar {
	return makeTouchBar{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ makeTouchBar) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _makeTouchBarClass) Alloc() makeTouchBar {
	rv := objc.Send[makeTouchBar](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _makeTouchBarClass) New() makeTouchBar {
	rv := objc.Send[makeTouchBar](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmakeTouchBar creates and returns a new initialized instance.
func NewmakeTouchBar() makeTouchBar {
	return makeTouchBarClass.New()
}

// Init initializes the instance.
func (m_ makeTouchBar) Init() makeTouchBar {
	rv := objc.Send[makeTouchBar](m_.ID(), selInit)
	return rv
}
