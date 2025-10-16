
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [layer] class.
var layerClass _layerClass

func init() {
	layerClass = _layerClass{objc.GetClass("layer")}
}

type _layerClass struct {
	objc.Class
}

// An interface definition for the [layer] class.
type Ilayer interface {
	ID() objc.ID
}

type layer struct {
	id objc.ID
}

func layerFrom(ptr unsafe.Pointer) layer {
	return layer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ layer) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _layerClass) Alloc() layer {
	rv := objc.Send[layer](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _layerClass) New() layer {
	rv := objc.Send[layer](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newlayer creates and returns a new initialized instance.
func Newlayer() layer {
	return layerClass.New()
}

// Init initializes the instance.
func (l_ layer) Init() layer {
	rv := objc.Send[layer](l_.ID(), selInit)
	return rv
}
