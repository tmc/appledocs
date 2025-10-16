
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CustomImageRep] class.
var CustomImageRepClass _CustomImageRepClass

func init() {
	CustomImageRepClass = _CustomImageRepClass{objc.GetClass("NSCustomImageRep")}
}

type _CustomImageRepClass struct {
	objc.Class
}

// An interface definition for the [CustomImageRep] class.
type ICustomImageRep interface {
	ID() objc.ID
}

type CustomImageRep struct {
	id objc.ID
}

func CustomImageRepFrom(ptr unsafe.Pointer) CustomImageRep {
	return CustomImageRep{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ CustomImageRep) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CustomImageRepClass) Alloc() CustomImageRep {
	rv := objc.Send[CustomImageRep](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CustomImageRepClass) New() CustomImageRep {
	rv := objc.Send[CustomImageRep](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCustomImageRep creates and returns a new initialized instance.
func NewCustomImageRep() CustomImageRep {
	return CustomImageRepClass.New()
}

// Init initializes the instance.
func (c_ CustomImageRep) Init() CustomImageRep {
	rv := objc.Send[CustomImageRep](c_.ID(), selInit)
	return rv
}
