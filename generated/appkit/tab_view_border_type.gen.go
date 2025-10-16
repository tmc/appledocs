
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tabViewBorderType] class.
var tabViewBorderTypeClass _tabViewBorderTypeClass

func init() {
	tabViewBorderTypeClass = _tabViewBorderTypeClass{objc.GetClass("tabViewBorderType")}
}

type _tabViewBorderTypeClass struct {
	objc.Class
}

// An interface definition for the [tabViewBorderType] class.
type ItabViewBorderType interface {
	ID() objc.ID
}

type tabViewBorderType struct {
	id objc.ID
}

func tabViewBorderTypeFrom(ptr unsafe.Pointer) tabViewBorderType {
	return tabViewBorderType{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tabViewBorderType) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tabViewBorderTypeClass) Alloc() tabViewBorderType {
	rv := objc.Send[tabViewBorderType](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tabViewBorderTypeClass) New() tabViewBorderType {
	rv := objc.Send[tabViewBorderType](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtabViewBorderType creates and returns a new initialized instance.
func NewtabViewBorderType() tabViewBorderType {
	return tabViewBorderTypeClass.New()
}

// Init initializes the instance.
func (t_ tabViewBorderType) Init() tabViewBorderType {
	rv := objc.Send[tabViewBorderType](t_.ID(), selInit)
	return rv
}
