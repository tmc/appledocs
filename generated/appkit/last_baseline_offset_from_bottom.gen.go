
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [lastBaselineOffsetFromBottom] class.
var lastBaselineOffsetFromBottomClass _lastBaselineOffsetFromBottomClass

func init() {
	lastBaselineOffsetFromBottomClass = _lastBaselineOffsetFromBottomClass{objc.GetClass("lastBaselineOffsetFromBottom")}
}

type _lastBaselineOffsetFromBottomClass struct {
	objc.Class
}

// An interface definition for the [lastBaselineOffsetFromBottom] class.
type IlastBaselineOffsetFromBottom interface {
	ID() objc.ID
}

type lastBaselineOffsetFromBottom struct {
	id objc.ID
}

func lastBaselineOffsetFromBottomFrom(ptr unsafe.Pointer) lastBaselineOffsetFromBottom {
	return lastBaselineOffsetFromBottom{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ lastBaselineOffsetFromBottom) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _lastBaselineOffsetFromBottomClass) Alloc() lastBaselineOffsetFromBottom {
	rv := objc.Send[lastBaselineOffsetFromBottom](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _lastBaselineOffsetFromBottomClass) New() lastBaselineOffsetFromBottom {
	rv := objc.Send[lastBaselineOffsetFromBottom](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlastBaselineOffsetFromBottom creates and returns a new initialized instance.
func NewlastBaselineOffsetFromBottom() lastBaselineOffsetFromBottom {
	return lastBaselineOffsetFromBottomClass.New()
}

// Init initializes the instance.
func (l_ lastBaselineOffsetFromBottom) Init() lastBaselineOffsetFromBottom {
	rv := objc.Send[lastBaselineOffsetFromBottom](l_.ID(), selInit)
	return rv
}
