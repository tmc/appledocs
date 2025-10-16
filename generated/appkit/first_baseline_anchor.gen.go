
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [firstBaselineAnchor] class.
var firstBaselineAnchorClass _firstBaselineAnchorClass

func init() {
	firstBaselineAnchorClass = _firstBaselineAnchorClass{objc.GetClass("firstBaselineAnchor")}
}

type _firstBaselineAnchorClass struct {
	objc.Class
}

// An interface definition for the [firstBaselineAnchor] class.
type IfirstBaselineAnchor interface {
	ID() objc.ID
}

type firstBaselineAnchor struct {
	id objc.ID
}

func firstBaselineAnchorFrom(ptr unsafe.Pointer) firstBaselineAnchor {
	return firstBaselineAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ firstBaselineAnchor) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _firstBaselineAnchorClass) Alloc() firstBaselineAnchor {
	rv := objc.Send[firstBaselineAnchor](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _firstBaselineAnchorClass) New() firstBaselineAnchor {
	rv := objc.Send[firstBaselineAnchor](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfirstBaselineAnchor creates and returns a new initialized instance.
func NewfirstBaselineAnchor() firstBaselineAnchor {
	return firstBaselineAnchorClass.New()
}

// Init initializes the instance.
func (f_ firstBaselineAnchor) Init() firstBaselineAnchor {
	rv := objc.Send[firstBaselineAnchor](f_.ID(), selInit)
	return rv
}
