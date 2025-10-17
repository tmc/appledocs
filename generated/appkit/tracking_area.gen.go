
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TrackingArea] class.
var TrackingAreaClass _TrackingAreaClass

func init() {
	TrackingAreaClass = _TrackingAreaClass{objc.GetClass("NSTrackingArea")}
}

type _TrackingAreaClass struct {
	objc.Class
}

// An interface definition for the [TrackingArea] class.
type ITrackingArea interface {
	ID() objc.ID
}

type TrackingArea struct {
	id objc.ID
}

func TrackingAreaFrom(ptr unsafe.Pointer) TrackingArea {
	return TrackingArea{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TrackingArea) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TrackingAreaClass) Alloc() TrackingArea {
	rv := objc.Send[TrackingArea](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TrackingAreaClass) New() TrackingArea {
	rv := objc.Send[TrackingArea](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTrackingArea creates and returns a new initialized instance.
func NewTrackingArea() TrackingArea {
	return TrackingAreaClass.New()
}

// Init initializes the instance.
func (t_ TrackingArea) Init() TrackingArea {
	rv := objc.Send[TrackingArea](t_.ID(), selInit)
	return rv
}
