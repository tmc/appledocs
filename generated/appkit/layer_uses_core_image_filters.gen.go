
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [layerUsesCoreImageFilters] class.
var layerUsesCoreImageFiltersClass _layerUsesCoreImageFiltersClass

func init() {
	layerUsesCoreImageFiltersClass = _layerUsesCoreImageFiltersClass{objc.GetClass("layerUsesCoreImageFilters")}
}

type _layerUsesCoreImageFiltersClass struct {
	objc.Class
}

// An interface definition for the [layerUsesCoreImageFilters] class.
type IlayerUsesCoreImageFilters interface {
	ID() objc.ID
}

type layerUsesCoreImageFilters struct {
	id objc.ID
}

func layerUsesCoreImageFiltersFrom(ptr unsafe.Pointer) layerUsesCoreImageFilters {
	return layerUsesCoreImageFilters{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ layerUsesCoreImageFilters) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _layerUsesCoreImageFiltersClass) Alloc() layerUsesCoreImageFilters {
	rv := objc.Send[layerUsesCoreImageFilters](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _layerUsesCoreImageFiltersClass) New() layerUsesCoreImageFilters {
	rv := objc.Send[layerUsesCoreImageFilters](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlayerUsesCoreImageFilters creates and returns a new initialized instance.
func NewlayerUsesCoreImageFilters() layerUsesCoreImageFilters {
	return layerUsesCoreImageFiltersClass.New()
}

// Init initializes the instance.
func (l_ layerUsesCoreImageFilters) Init() layerUsesCoreImageFilters {
	rv := objc.Send[layerUsesCoreImageFilters](l_.ID(), selInit)
	return rv
}
