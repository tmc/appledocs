
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [trackingAreas] class.
var trackingAreasClass _trackingAreasClass

func init() {
	trackingAreasClass = _trackingAreasClass{objc.GetClass("trackingAreas")}
}

type _trackingAreasClass struct {
	objc.Class
}

// An interface definition for the [trackingAreas] class.
type ItrackingAreas interface {
	ID() objc.ID
}

type trackingAreas struct {
	id objc.ID
}

func trackingAreasFrom(ptr unsafe.Pointer) trackingAreas {
	return trackingAreas{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ trackingAreas) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _trackingAreasClass) Alloc() trackingAreas {
	rv := objc.Send[trackingAreas](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _trackingAreasClass) New() trackingAreas {
	rv := objc.Send[trackingAreas](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtrackingAreas creates and returns a new initialized instance.
func NewtrackingAreas() trackingAreas {
	return trackingAreasClass.New()
}

// Init initializes the instance.
func (t_ trackingAreas) Init() trackingAreas {
	rv := objc.Send[trackingAreas](t_.ID(), selInit)
	return rv
}
