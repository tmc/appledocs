
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canDrawSubviewsIntoLayer] class.
var canDrawSubviewsIntoLayerClass _canDrawSubviewsIntoLayerClass

func init() {
	canDrawSubviewsIntoLayerClass = _canDrawSubviewsIntoLayerClass{objc.GetClass("canDrawSubviewsIntoLayer")}
}

type _canDrawSubviewsIntoLayerClass struct {
	objc.Class
}

// An interface definition for the [canDrawSubviewsIntoLayer] class.
type IcanDrawSubviewsIntoLayer interface {
	ID() objc.ID
}

type canDrawSubviewsIntoLayer struct {
	id objc.ID
}

func canDrawSubviewsIntoLayerFrom(ptr unsafe.Pointer) canDrawSubviewsIntoLayer {
	return canDrawSubviewsIntoLayer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canDrawSubviewsIntoLayer) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canDrawSubviewsIntoLayerClass) Alloc() canDrawSubviewsIntoLayer {
	rv := objc.Send[canDrawSubviewsIntoLayer](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canDrawSubviewsIntoLayerClass) New() canDrawSubviewsIntoLayer {
	rv := objc.Send[canDrawSubviewsIntoLayer](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanDrawSubviewsIntoLayer creates and returns a new initialized instance.
func NewcanDrawSubviewsIntoLayer() canDrawSubviewsIntoLayer {
	return canDrawSubviewsIntoLayerClass.New()
}

// Init initializes the instance.
func (c_ canDrawSubviewsIntoLayer) Init() canDrawSubviewsIntoLayer {
	rv := objc.Send[canDrawSubviewsIntoLayer](c_.ID(), selInit)
	return rv
}
