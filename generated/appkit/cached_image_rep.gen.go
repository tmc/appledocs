
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CachedImageRep] class.
var CachedImageRepClass _CachedImageRepClass

func init() {
	CachedImageRepClass = _CachedImageRepClass{objc.GetClass("NSCachedImageRep")}
}

type _CachedImageRepClass struct {
	objc.Class
}

// An interface definition for the [CachedImageRep] class.
type ICachedImageRep interface {
	ID() objc.ID
}

type CachedImageRep struct {
	id objc.ID
}

func CachedImageRepFrom(ptr unsafe.Pointer) CachedImageRep {
	return CachedImageRep{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ CachedImageRep) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CachedImageRepClass) Alloc() CachedImageRep {
	rv := objc.Send[CachedImageRep](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CachedImageRepClass) New() CachedImageRep {
	rv := objc.Send[CachedImageRep](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCachedImageRep creates and returns a new initialized instance.
func NewCachedImageRep() CachedImageRep {
	return CachedImageRepClass.New()
}

// Init initializes the instance.
func (c_ CachedImageRep) Init() CachedImageRep {
	rv := objc.Send[CachedImageRep](c_.ID(), selInit)
	return rv
}
