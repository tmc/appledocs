
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [wantsLayer] class.
var wantsLayerClass _wantsLayerClass

func init() {
	wantsLayerClass = _wantsLayerClass{objc.GetClass("wantsLayer")}
}

type _wantsLayerClass struct {
	objc.Class
}

// An interface definition for the [wantsLayer] class.
type IwantsLayer interface {
	ID() objc.ID
}

type wantsLayer struct {
	id objc.ID
}

func wantsLayerFrom(ptr unsafe.Pointer) wantsLayer {
	return wantsLayer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ wantsLayer) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _wantsLayerClass) Alloc() wantsLayer {
	rv := objc.Send[wantsLayer](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _wantsLayerClass) New() wantsLayer {
	rv := objc.Send[wantsLayer](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwantsLayer creates and returns a new initialized instance.
func NewwantsLayer() wantsLayer {
	return wantsLayerClass.New()
}

// Init initializes the instance.
func (w_ wantsLayer) Init() wantsLayer {
	rv := objc.Send[wantsLayer](w_.ID(), selInit)
	return rv
}
