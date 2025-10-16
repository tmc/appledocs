
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [makeBackingLayer] class.
var makeBackingLayerClass _makeBackingLayerClass

func init() {
	makeBackingLayerClass = _makeBackingLayerClass{objc.GetClass("makeBackingLayer")}
}

type _makeBackingLayerClass struct {
	objc.Class
}

// An interface definition for the [makeBackingLayer] class.
type ImakeBackingLayer interface {
	ID() objc.ID
}

type makeBackingLayer struct {
	id objc.ID
}

func makeBackingLayerFrom(ptr unsafe.Pointer) makeBackingLayer {
	return makeBackingLayer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ makeBackingLayer) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _makeBackingLayerClass) Alloc() makeBackingLayer {
	rv := objc.Send[makeBackingLayer](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _makeBackingLayerClass) New() makeBackingLayer {
	rv := objc.Send[makeBackingLayer](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmakeBackingLayer creates and returns a new initialized instance.
func NewmakeBackingLayer() makeBackingLayer {
	return makeBackingLayerClass.New()
}

// Init initializes the instance.
func (m_ makeBackingLayer) Init() makeBackingLayer {
	rv := objc.Send[makeBackingLayer](m_.ID(), selInit)
	return rv
}
