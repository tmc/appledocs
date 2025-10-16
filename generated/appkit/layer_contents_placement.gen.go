
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [layerContentsPlacement] class.
var layerContentsPlacementClass _layerContentsPlacementClass

func init() {
	layerContentsPlacementClass = _layerContentsPlacementClass{objc.GetClass("layerContentsPlacement")}
}

type _layerContentsPlacementClass struct {
	objc.Class
}

// An interface definition for the [layerContentsPlacement] class.
type IlayerContentsPlacement interface {
	ID() objc.ID
}

type layerContentsPlacement struct {
	id objc.ID
}

func layerContentsPlacementFrom(ptr unsafe.Pointer) layerContentsPlacement {
	return layerContentsPlacement{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ layerContentsPlacement) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _layerContentsPlacementClass) Alloc() layerContentsPlacement {
	rv := objc.Send[layerContentsPlacement](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _layerContentsPlacementClass) New() layerContentsPlacement {
	rv := objc.Send[layerContentsPlacement](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlayerContentsPlacement creates and returns a new initialized instance.
func NewlayerContentsPlacement() layerContentsPlacement {
	return layerContentsPlacementClass.New()
}

// Init initializes the instance.
func (l_ layerContentsPlacement) Init() layerContentsPlacement {
	rv := objc.Send[layerContentsPlacement](l_.ID(), selInit)
	return rv
}
