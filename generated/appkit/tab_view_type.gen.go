
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tabViewType] class.
var tabViewTypeClass _tabViewTypeClass

func init() {
	tabViewTypeClass = _tabViewTypeClass{objc.GetClass("tabViewType")}
}

type _tabViewTypeClass struct {
	objc.Class
}

// An interface definition for the [tabViewType] class.
type ItabViewType interface {
	ID() objc.ID
}

type tabViewType struct {
	id objc.ID
}

func tabViewTypeFrom(ptr unsafe.Pointer) tabViewType {
	return tabViewType{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tabViewType) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tabViewTypeClass) Alloc() tabViewType {
	rv := objc.Send[tabViewType](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tabViewTypeClass) New() tabViewType {
	rv := objc.Send[tabViewType](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtabViewType creates and returns a new initialized instance.
func NewtabViewType() tabViewType {
	return tabViewTypeClass.New()
}

// Init initializes the instance.
func (t_ tabViewType) Init() tabViewType {
	rv := objc.Send[tabViewType](t_.ID(), selInit)
	return rv
}
