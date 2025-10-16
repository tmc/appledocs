
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [baselineOffsetFromBottom] class.
var baselineOffsetFromBottomClass _baselineOffsetFromBottomClass

func init() {
	baselineOffsetFromBottomClass = _baselineOffsetFromBottomClass{objc.GetClass("baselineOffsetFromBottom")}
}

type _baselineOffsetFromBottomClass struct {
	objc.Class
}

// An interface definition for the [baselineOffsetFromBottom] class.
type IbaselineOffsetFromBottom interface {
	ID() objc.ID
}

type baselineOffsetFromBottom struct {
	id objc.ID
}

func baselineOffsetFromBottomFrom(ptr unsafe.Pointer) baselineOffsetFromBottom {
	return baselineOffsetFromBottom{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ baselineOffsetFromBottom) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _baselineOffsetFromBottomClass) Alloc() baselineOffsetFromBottom {
	rv := objc.Send[baselineOffsetFromBottom](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _baselineOffsetFromBottomClass) New() baselineOffsetFromBottom {
	rv := objc.Send[baselineOffsetFromBottom](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbaselineOffsetFromBottom creates and returns a new initialized instance.
func NewbaselineOffsetFromBottom() baselineOffsetFromBottom {
	return baselineOffsetFromBottomClass.New()
}

// Init initializes the instance.
func (b_ baselineOffsetFromBottom) Init() baselineOffsetFromBottom {
	rv := objc.Send[baselineOffsetFromBottom](b_.ID(), selInit)
	return rv
}
