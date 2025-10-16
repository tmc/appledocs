
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [layerContentsRedrawPolicy] class.
var layerContentsRedrawPolicyClass _layerContentsRedrawPolicyClass

func init() {
	layerContentsRedrawPolicyClass = _layerContentsRedrawPolicyClass{objc.GetClass("layerContentsRedrawPolicy")}
}

type _layerContentsRedrawPolicyClass struct {
	objc.Class
}

// An interface definition for the [layerContentsRedrawPolicy] class.
type IlayerContentsRedrawPolicy interface {
	ID() objc.ID
}

type layerContentsRedrawPolicy struct {
	id objc.ID
}

func layerContentsRedrawPolicyFrom(ptr unsafe.Pointer) layerContentsRedrawPolicy {
	return layerContentsRedrawPolicy{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ layerContentsRedrawPolicy) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _layerContentsRedrawPolicyClass) Alloc() layerContentsRedrawPolicy {
	rv := objc.Send[layerContentsRedrawPolicy](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _layerContentsRedrawPolicyClass) New() layerContentsRedrawPolicy {
	rv := objc.Send[layerContentsRedrawPolicy](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlayerContentsRedrawPolicy creates and returns a new initialized instance.
func NewlayerContentsRedrawPolicy() layerContentsRedrawPolicy {
	return layerContentsRedrawPolicyClass.New()
}

// Init initializes the instance.
func (l_ layerContentsRedrawPolicy) Init() layerContentsRedrawPolicy {
	rv := objc.Send[layerContentsRedrawPolicy](l_.ID(), selInit)
	return rv
}
