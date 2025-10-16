
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [popoverTouchBar] class.
var popoverTouchBarClass _popoverTouchBarClass

func init() {
	popoverTouchBarClass = _popoverTouchBarClass{objc.GetClass("popoverTouchBar")}
}

type _popoverTouchBarClass struct {
	objc.Class
}

// An interface definition for the [popoverTouchBar] class.
type IpopoverTouchBar interface {
	ID() objc.ID
}

type popoverTouchBar struct {
	id objc.ID
}

func popoverTouchBarFrom(ptr unsafe.Pointer) popoverTouchBar {
	return popoverTouchBar{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ popoverTouchBar) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _popoverTouchBarClass) Alloc() popoverTouchBar {
	rv := objc.Send[popoverTouchBar](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _popoverTouchBarClass) New() popoverTouchBar {
	rv := objc.Send[popoverTouchBar](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpopoverTouchBar creates and returns a new initialized instance.
func NewpopoverTouchBar() popoverTouchBar {
	return popoverTouchBarClass.New()
}

// Init initializes the instance.
func (p_ popoverTouchBar) Init() popoverTouchBar {
	rv := objc.Send[popoverTouchBar](p_.ID(), selInit)
	return rv
}
