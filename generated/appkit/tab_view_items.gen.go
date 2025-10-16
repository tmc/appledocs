
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tabViewItems] class.
var tabViewItemsClass _tabViewItemsClass

func init() {
	tabViewItemsClass = _tabViewItemsClass{objc.GetClass("tabViewItems")}
}

type _tabViewItemsClass struct {
	objc.Class
}

// An interface definition for the [tabViewItems] class.
type ItabViewItems interface {
	ID() objc.ID
}

type tabViewItems struct {
	id objc.ID
}

func tabViewItemsFrom(ptr unsafe.Pointer) tabViewItems {
	return tabViewItems{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tabViewItems) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tabViewItemsClass) Alloc() tabViewItems {
	rv := objc.Send[tabViewItems](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tabViewItemsClass) New() tabViewItems {
	rv := objc.Send[tabViewItems](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtabViewItems creates and returns a new initialized instance.
func NewtabViewItems() tabViewItems {
	return tabViewItemsClass.New()
}

// Init initializes the instance.
func (t_ tabViewItems) Init() tabViewItems {
	rv := objc.Send[tabViewItems](t_.ID(), selInit)
	return rv
}
