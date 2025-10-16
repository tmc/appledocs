
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [aspectRatio] class.
var aspectRatioClass _aspectRatioClass

func init() {
	aspectRatioClass = _aspectRatioClass{objc.GetClass("aspectRatio")}
}

type _aspectRatioClass struct {
	objc.Class
}

// An interface definition for the [aspectRatio] class.
type IaspectRatio interface {
	ID() objc.ID
}

type aspectRatio struct {
	id objc.ID
}

func aspectRatioFrom(ptr unsafe.Pointer) aspectRatio {
	return aspectRatio{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ aspectRatio) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _aspectRatioClass) Alloc() aspectRatio {
	rv := objc.Send[aspectRatio](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _aspectRatioClass) New() aspectRatio {
	rv := objc.Send[aspectRatio](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewaspectRatio creates and returns a new initialized instance.
func NewaspectRatio() aspectRatio {
	return aspectRatioClass.New()
}

// Init initializes the instance.
func (a_ aspectRatio) Init() aspectRatio {
	rv := objc.Send[aspectRatio](a_.ID(), selInit)
	return rv
}
