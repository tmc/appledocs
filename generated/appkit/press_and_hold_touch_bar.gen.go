
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [pressAndHoldTouchBar] class.
var pressAndHoldTouchBarClass _pressAndHoldTouchBarClass

func init() {
	pressAndHoldTouchBarClass = _pressAndHoldTouchBarClass{objc.GetClass("pressAndHoldTouchBar")}
}

type _pressAndHoldTouchBarClass struct {
	objc.Class
}

// An interface definition for the [pressAndHoldTouchBar] class.
type IpressAndHoldTouchBar interface {
	ID() objc.ID
}

type pressAndHoldTouchBar struct {
	id objc.ID
}

func pressAndHoldTouchBarFrom(ptr unsafe.Pointer) pressAndHoldTouchBar {
	return pressAndHoldTouchBar{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ pressAndHoldTouchBar) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _pressAndHoldTouchBarClass) Alloc() pressAndHoldTouchBar {
	rv := objc.Send[pressAndHoldTouchBar](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _pressAndHoldTouchBarClass) New() pressAndHoldTouchBar {
	rv := objc.Send[pressAndHoldTouchBar](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpressAndHoldTouchBar creates and returns a new initialized instance.
func NewpressAndHoldTouchBar() pressAndHoldTouchBar {
	return pressAndHoldTouchBarClass.New()
}

// Init initializes the instance.
func (p_ pressAndHoldTouchBar) Init() pressAndHoldTouchBar {
	rv := objc.Send[pressAndHoldTouchBar](p_.ID(), selInit)
	return rv
}
