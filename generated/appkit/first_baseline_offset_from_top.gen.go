
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [firstBaselineOffsetFromTop] class.
var firstBaselineOffsetFromTopClass _firstBaselineOffsetFromTopClass

func init() {
	firstBaselineOffsetFromTopClass = _firstBaselineOffsetFromTopClass{objc.GetClass("firstBaselineOffsetFromTop")}
}

type _firstBaselineOffsetFromTopClass struct {
	objc.Class
}

// An interface definition for the [firstBaselineOffsetFromTop] class.
type IfirstBaselineOffsetFromTop interface {
	ID() objc.ID
}

type firstBaselineOffsetFromTop struct {
	id objc.ID
}

func firstBaselineOffsetFromTopFrom(ptr unsafe.Pointer) firstBaselineOffsetFromTop {
	return firstBaselineOffsetFromTop{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ firstBaselineOffsetFromTop) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _firstBaselineOffsetFromTopClass) Alloc() firstBaselineOffsetFromTop {
	rv := objc.Send[firstBaselineOffsetFromTop](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _firstBaselineOffsetFromTopClass) New() firstBaselineOffsetFromTop {
	rv := objc.Send[firstBaselineOffsetFromTop](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfirstBaselineOffsetFromTop creates and returns a new initialized instance.
func NewfirstBaselineOffsetFromTop() firstBaselineOffsetFromTop {
	return firstBaselineOffsetFromTopClass.New()
}

// Init initializes the instance.
func (f_ firstBaselineOffsetFromTop) Init() firstBaselineOffsetFromTop {
	rv := objc.Send[firstBaselineOffsetFromTop](f_.ID(), selInit)
	return rv
}
