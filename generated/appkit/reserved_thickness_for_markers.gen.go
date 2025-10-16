
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [reservedThicknessForMarkers] class.
var reservedThicknessForMarkersClass _reservedThicknessForMarkersClass

func init() {
	reservedThicknessForMarkersClass = _reservedThicknessForMarkersClass{objc.GetClass("reservedThicknessForMarkers")}
}

type _reservedThicknessForMarkersClass struct {
	objc.Class
}

// An interface definition for the [reservedThicknessForMarkers] class.
type IreservedThicknessForMarkers interface {
	ID() objc.ID
}

type reservedThicknessForMarkers struct {
	id objc.ID
}

func reservedThicknessForMarkersFrom(ptr unsafe.Pointer) reservedThicknessForMarkers {
	return reservedThicknessForMarkers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ reservedThicknessForMarkers) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _reservedThicknessForMarkersClass) Alloc() reservedThicknessForMarkers {
	rv := objc.Send[reservedThicknessForMarkers](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _reservedThicknessForMarkersClass) New() reservedThicknessForMarkers {
	rv := objc.Send[reservedThicknessForMarkers](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewreservedThicknessForMarkers creates and returns a new initialized instance.
func NewreservedThicknessForMarkers() reservedThicknessForMarkers {
	return reservedThicknessForMarkersClass.New()
}

// Init initializes the instance.
func (r_ reservedThicknessForMarkers) Init() reservedThicknessForMarkers {
	rv := objc.Send[reservedThicknessForMarkers](r_.ID(), selInit)
	return rv
}
