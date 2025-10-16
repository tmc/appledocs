
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GroupTouchBarItem] class.
var GroupTouchBarItemClass _GroupTouchBarItemClass

func init() {
	GroupTouchBarItemClass = _GroupTouchBarItemClass{objc.GetClass("NSGroupTouchBarItem")}
}

type _GroupTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [GroupTouchBarItem] class.
type IGroupTouchBarItem interface {
	ID() objc.ID
}

type GroupTouchBarItem struct {
	id objc.ID
}

func GroupTouchBarItemFrom(ptr unsafe.Pointer) GroupTouchBarItem {
	return GroupTouchBarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ GroupTouchBarItem) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _GroupTouchBarItemClass) Alloc() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _GroupTouchBarItemClass) New() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewGroupTouchBarItem creates and returns a new initialized instance.
func NewGroupTouchBarItem() GroupTouchBarItem {
	return GroupTouchBarItemClass.New()
}

// Init initializes the instance.
func (g_ GroupTouchBarItem) Init() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](g_.ID(), selInit)
	return rv
}
// A bar that holds this group’s items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGroupTouchBarItem/groupTouchBar
func (g_ GroupTouchBarItem) GroupTouchBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("groupTouchBar"))
	return rv
}
// SetGroupTouchBar sets the value of the groupTouchBar property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGroupTouchBarItem/groupTouchBar
func (g_ GroupTouchBarItem) SetGroupTouchBar(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setGroupTouchBar:"), value)
}
