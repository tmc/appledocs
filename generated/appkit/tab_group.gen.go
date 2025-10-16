
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tabGroup] class.
var tabGroupClass _tabGroupClass

func init() {
	tabGroupClass = _tabGroupClass{objc.GetClass("tabGroup")}
}

type _tabGroupClass struct {
	objc.Class
}

// An interface definition for the [tabGroup] class.
type ItabGroup interface {
	ID() objc.ID
}

type tabGroup struct {
	id objc.ID
}

func tabGroupFrom(ptr unsafe.Pointer) tabGroup {
	return tabGroup{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tabGroup) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tabGroupClass) Alloc() tabGroup {
	rv := objc.Send[tabGroup](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tabGroupClass) New() tabGroup {
	rv := objc.Send[tabGroup](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtabGroup creates and returns a new initialized instance.
func NewtabGroup() tabGroup {
	return tabGroupClass.New()
}

// Init initializes the instance.
func (t_ tabGroup) Init() tabGroup {
	rv := objc.Send[tabGroup](t_.ID(), selInit)
	return rv
}
