
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [lastBaselineAnchor] class.
var lastBaselineAnchorClass _lastBaselineAnchorClass

func init() {
	lastBaselineAnchorClass = _lastBaselineAnchorClass{objc.GetClass("lastBaselineAnchor")}
}

type _lastBaselineAnchorClass struct {
	objc.Class
}

// An interface definition for the [lastBaselineAnchor] class.
type IlastBaselineAnchor interface {
	ID() objc.ID
}

type lastBaselineAnchor struct {
	id objc.ID
}

func lastBaselineAnchorFrom(ptr unsafe.Pointer) lastBaselineAnchor {
	return lastBaselineAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ lastBaselineAnchor) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _lastBaselineAnchorClass) Alloc() lastBaselineAnchor {
	rv := objc.Send[lastBaselineAnchor](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _lastBaselineAnchorClass) New() lastBaselineAnchor {
	rv := objc.Send[lastBaselineAnchor](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlastBaselineAnchor creates and returns a new initialized instance.
func NewlastBaselineAnchor() lastBaselineAnchor {
	return lastBaselineAnchorClass.New()
}

// Init initializes the instance.
func (l_ lastBaselineAnchor) Init() lastBaselineAnchor {
	rv := objc.Send[lastBaselineAnchor](l_.ID(), selInit)
	return rv
}
