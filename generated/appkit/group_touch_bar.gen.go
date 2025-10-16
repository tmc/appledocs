
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [groupTouchBar] class.
var groupTouchBarClass _groupTouchBarClass

func init() {
	groupTouchBarClass = _groupTouchBarClass{objc.GetClass("groupTouchBar")}
}

type _groupTouchBarClass struct {
	objc.Class
}

// An interface definition for the [groupTouchBar] class.
type IgroupTouchBar interface {
	ID() objc.ID
}

type groupTouchBar struct {
	id objc.ID
}

func groupTouchBarFrom(ptr unsafe.Pointer) groupTouchBar {
	return groupTouchBar{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ groupTouchBar) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _groupTouchBarClass) Alloc() groupTouchBar {
	rv := objc.Send[groupTouchBar](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _groupTouchBarClass) New() groupTouchBar {
	rv := objc.Send[groupTouchBar](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewgroupTouchBar creates and returns a new initialized instance.
func NewgroupTouchBar() groupTouchBar {
	return groupTouchBarClass.New()
}

// Init initializes the instance.
func (g_ groupTouchBar) Init() groupTouchBar {
	rv := objc.Send[groupTouchBar](g_.ID(), selInit)
	return rv
}
