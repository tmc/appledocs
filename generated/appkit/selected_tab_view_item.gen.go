
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [selectedTabViewItem] class.
var selectedTabViewItemClass _selectedTabViewItemClass

func init() {
	selectedTabViewItemClass = _selectedTabViewItemClass{objc.GetClass("selectedTabViewItem")}
}

type _selectedTabViewItemClass struct {
	objc.Class
}

// An interface definition for the [selectedTabViewItem] class.
type IselectedTabViewItem interface {
	ID() objc.ID
}

type selectedTabViewItem struct {
	id objc.ID
}

func selectedTabViewItemFrom(ptr unsafe.Pointer) selectedTabViewItem {
	return selectedTabViewItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ selectedTabViewItem) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _selectedTabViewItemClass) Alloc() selectedTabViewItem {
	rv := objc.Send[selectedTabViewItem](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _selectedTabViewItemClass) New() selectedTabViewItem {
	rv := objc.Send[selectedTabViewItem](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewselectedTabViewItem creates and returns a new initialized instance.
func NewselectedTabViewItem() selectedTabViewItem {
	return selectedTabViewItemClass.New()
}

// Init initializes the instance.
func (s_ selectedTabViewItem) Init() selectedTabViewItem {
	rv := objc.Send[selectedTabViewItem](s_.ID(), selInit)
	return rv
}
