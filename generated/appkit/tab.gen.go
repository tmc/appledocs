
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tab] class.
var tabClass _tabClass

func init() {
	tabClass = _tabClass{objc.GetClass("tab")}
}

type _tabClass struct {
	objc.Class
}

// An interface definition for the [tab] class.
type Itab interface {
	ID() objc.ID
}

type tab struct {
	id objc.ID
}

func tabFrom(ptr unsafe.Pointer) tab {
	return tab{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tab) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tabClass) Alloc() tab {
	rv := objc.Send[tab](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tabClass) New() tab {
	rv := objc.Send[tab](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newtab creates and returns a new initialized instance.
func Newtab() tab {
	return tabClass.New()
}

// Init initializes the instance.
func (t_ tab) Init() tab {
	rv := objc.Send[tab](t_.ID(), selInit)
	return rv
}
