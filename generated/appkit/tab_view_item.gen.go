
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TabViewItem] class.
var TabViewItemClass _TabViewItemClass

func init() {
	TabViewItemClass = _TabViewItemClass{objc.GetClass("NSTabViewItem")}
}

type _TabViewItemClass struct {
	objc.Class
}

// An interface definition for the [TabViewItem] class.
type ITabViewItem interface {
	ID() objc.ID
}

type TabViewItem struct {
	id objc.ID
}

func TabViewItemFrom(ptr unsafe.Pointer) TabViewItem {
	return TabViewItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TabViewItem) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TabViewItemClass) Alloc() TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TabViewItemClass) New() TabViewItem {
	rv := objc.Send[TabViewItem](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTabViewItem creates and returns a new initialized instance.
func NewTabViewItem() TabViewItem {
	return TabViewItemClass.New()
}

// Init initializes the instance.
func (t_ TabViewItem) Init() TabViewItem {
	rv := objc.Send[TabViewItem](t_.ID(), selInit)
	return rv
}
