
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BindingSelectionMarker] class.
var BindingSelectionMarkerClass _BindingSelectionMarkerClass

func init() {
	BindingSelectionMarkerClass = _BindingSelectionMarkerClass{objc.GetClass("NSBindingSelectionMarker")}
}

type _BindingSelectionMarkerClass struct {
	objc.Class
}

// An interface definition for the [BindingSelectionMarker] class.
type IBindingSelectionMarker interface {
	ID() objc.ID
}

type BindingSelectionMarker struct {
	id objc.ID
}

func BindingSelectionMarkerFrom(ptr unsafe.Pointer) BindingSelectionMarker {
	return BindingSelectionMarker{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ BindingSelectionMarker) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _BindingSelectionMarkerClass) Alloc() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _BindingSelectionMarkerClass) New() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewBindingSelectionMarker creates and returns a new initialized instance.
func NewBindingSelectionMarker() BindingSelectionMarker {
	return BindingSelectionMarkerClass.New()
}

// Init initializes the instance.
func (b_ BindingSelectionMarker) Init() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](b_.ID(), selInit)
	return rv
}
