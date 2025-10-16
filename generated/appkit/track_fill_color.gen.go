
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [trackFillColor] class.
var trackFillColorClass _trackFillColorClass

func init() {
	trackFillColorClass = _trackFillColorClass{objc.GetClass("trackFillColor")}
}

type _trackFillColorClass struct {
	objc.Class
}

// An interface definition for the [trackFillColor] class.
type ItrackFillColor interface {
	ID() objc.ID
}

type trackFillColor struct {
	id objc.ID
}

func trackFillColorFrom(ptr unsafe.Pointer) trackFillColor {
	return trackFillColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ trackFillColor) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _trackFillColorClass) Alloc() trackFillColor {
	rv := objc.Send[trackFillColor](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _trackFillColorClass) New() trackFillColor {
	rv := objc.Send[trackFillColor](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtrackFillColor creates and returns a new initialized instance.
func NewtrackFillColor() trackFillColor {
	return trackFillColorClass.New()
}

// Init initializes the instance.
func (t_ trackFillColor) Init() trackFillColor {
	rv := objc.Send[trackFillColor](t_.ID(), selInit)
	return rv
}
