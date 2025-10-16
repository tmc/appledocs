
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [updateTrackingAreas] class.
var updateTrackingAreasClass _updateTrackingAreasClass

func init() {
	updateTrackingAreasClass = _updateTrackingAreasClass{objc.GetClass("updateTrackingAreas")}
}

type _updateTrackingAreasClass struct {
	objc.Class
}

// An interface definition for the [updateTrackingAreas] class.
type IupdateTrackingAreas interface {
	ID() objc.ID
}

type updateTrackingAreas struct {
	id objc.ID
}

func updateTrackingAreasFrom(ptr unsafe.Pointer) updateTrackingAreas {
	return updateTrackingAreas{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ updateTrackingAreas) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _updateTrackingAreasClass) Alloc() updateTrackingAreas {
	rv := objc.Send[updateTrackingAreas](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _updateTrackingAreasClass) New() updateTrackingAreas {
	rv := objc.Send[updateTrackingAreas](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewupdateTrackingAreas creates and returns a new initialized instance.
func NewupdateTrackingAreas() updateTrackingAreas {
	return updateTrackingAreasClass.New()
}

// Init initializes the instance.
func (u_ updateTrackingAreas) Init() updateTrackingAreas {
	rv := objc.Send[updateTrackingAreas](u_.ID(), selInit)
	return rv
}
