
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ToolbarItemGroup] class.
var ToolbarItemGroupClass _ToolbarItemGroupClass

func init() {
	ToolbarItemGroupClass = _ToolbarItemGroupClass{objc.GetClass("NSToolbarItemGroup")}
}

type _ToolbarItemGroupClass struct {
	objc.Class
}

// An interface definition for the [ToolbarItemGroup] class.
type IToolbarItemGroup interface {
	ID() objc.ID
}

type ToolbarItemGroup struct {
	id objc.ID
}

func ToolbarItemGroupFrom(ptr unsafe.Pointer) ToolbarItemGroup {
	return ToolbarItemGroup{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ ToolbarItemGroup) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _ToolbarItemGroupClass) Alloc() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _ToolbarItemGroupClass) New() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewToolbarItemGroup creates and returns a new initialized instance.
func NewToolbarItemGroup() ToolbarItemGroup {
	return ToolbarItemGroupClass.New()
}

// Init initializes the instance.
func (t_ ToolbarItemGroup) Init() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](t_.ID(), selInit)
	return rv
}
