
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RulerMarker] class.
var RulerMarkerClass _RulerMarkerClass

func init() {
	RulerMarkerClass = _RulerMarkerClass{objc.GetClass("NSRulerMarker")}
}

type _RulerMarkerClass struct {
	objc.Class
}

// An interface definition for the [RulerMarker] class.
type IRulerMarker interface {
	ID() objc.ID
}

type RulerMarker struct {
	id objc.ID
}

func RulerMarkerFrom(ptr unsafe.Pointer) RulerMarker {
	return RulerMarker{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ RulerMarker) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _RulerMarkerClass) Alloc() RulerMarker {
	rv := objc.Send[RulerMarker](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _RulerMarkerClass) New() RulerMarker {
	rv := objc.Send[RulerMarker](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewRulerMarker creates and returns a new initialized instance.
func NewRulerMarker() RulerMarker {
	return RulerMarkerClass.New()
}

// Init initializes the instance.
func (r_ RulerMarker) Init() RulerMarker {
	rv := objc.Send[RulerMarker](r_.ID(), selInit)
	return rv
}
