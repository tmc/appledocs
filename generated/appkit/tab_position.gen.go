
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tabPosition] class.
var tabPositionClass _tabPositionClass

func init() {
	tabPositionClass = _tabPositionClass{objc.GetClass("tabPosition")}
}

type _tabPositionClass struct {
	objc.Class
}

// An interface definition for the [tabPosition] class.
type ItabPosition interface {
	ID() objc.ID
}

type tabPosition struct {
	id objc.ID
}

func tabPositionFrom(ptr unsafe.Pointer) tabPosition {
	return tabPosition{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tabPosition) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tabPositionClass) Alloc() tabPosition {
	rv := objc.Send[tabPosition](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tabPositionClass) New() tabPosition {
	rv := objc.Send[tabPosition](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtabPosition creates and returns a new initialized instance.
func NewtabPosition() tabPosition {
	return tabPositionClass.New()
}

// Init initializes the instance.
func (t_ tabPosition) Init() tabPosition {
	rv := objc.Send[tabPosition](t_.ID(), selInit)
	return rv
}
