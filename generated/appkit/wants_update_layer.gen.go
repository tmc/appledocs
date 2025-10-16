
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [wantsUpdateLayer] class.
var wantsUpdateLayerClass _wantsUpdateLayerClass

func init() {
	wantsUpdateLayerClass = _wantsUpdateLayerClass{objc.GetClass("wantsUpdateLayer")}
}

type _wantsUpdateLayerClass struct {
	objc.Class
}

// An interface definition for the [wantsUpdateLayer] class.
type IwantsUpdateLayer interface {
	ID() objc.ID
}

type wantsUpdateLayer struct {
	id objc.ID
}

func wantsUpdateLayerFrom(ptr unsafe.Pointer) wantsUpdateLayer {
	return wantsUpdateLayer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ wantsUpdateLayer) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _wantsUpdateLayerClass) Alloc() wantsUpdateLayer {
	rv := objc.Send[wantsUpdateLayer](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _wantsUpdateLayerClass) New() wantsUpdateLayer {
	rv := objc.Send[wantsUpdateLayer](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwantsUpdateLayer creates and returns a new initialized instance.
func NewwantsUpdateLayer() wantsUpdateLayer {
	return wantsUpdateLayerClass.New()
}

// Init initializes the instance.
func (w_ wantsUpdateLayer) Init() wantsUpdateLayer {
	rv := objc.Send[wantsUpdateLayer](w_.ID(), selInit)
	return rv
}
