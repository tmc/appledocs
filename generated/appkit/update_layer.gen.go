
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [updateLayer] class.
var updateLayerClass _updateLayerClass

func init() {
	updateLayerClass = _updateLayerClass{objc.GetClass("updateLayer")}
}

type _updateLayerClass struct {
	objc.Class
}

// An interface definition for the [updateLayer] class.
type IupdateLayer interface {
	ID() objc.ID
}

type updateLayer struct {
	id objc.ID
}

func updateLayerFrom(ptr unsafe.Pointer) updateLayer {
	return updateLayer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ updateLayer) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _updateLayerClass) Alloc() updateLayer {
	rv := objc.Send[updateLayer](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _updateLayerClass) New() updateLayer {
	rv := objc.Send[updateLayer](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewupdateLayer creates and returns a new initialized instance.
func NewupdateLayer() updateLayer {
	return updateLayerClass.New()
}

// Init initializes the instance.
func (u_ updateLayer) Init() updateLayer {
	rv := objc.Send[updateLayer](u_.ID(), selInit)
	return rv
}
