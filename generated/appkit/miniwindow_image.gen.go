
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [miniwindowImage] class.
var miniwindowImageClass _miniwindowImageClass

func init() {
	miniwindowImageClass = _miniwindowImageClass{objc.GetClass("miniwindowImage")}
}

type _miniwindowImageClass struct {
	objc.Class
}

// An interface definition for the [miniwindowImage] class.
type IminiwindowImage interface {
	ID() objc.ID
}

type miniwindowImage struct {
	id objc.ID
}

func miniwindowImageFrom(ptr unsafe.Pointer) miniwindowImage {
	return miniwindowImage{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ miniwindowImage) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _miniwindowImageClass) Alloc() miniwindowImage {
	rv := objc.Send[miniwindowImage](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _miniwindowImageClass) New() miniwindowImage {
	rv := objc.Send[miniwindowImage](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewminiwindowImage creates and returns a new initialized instance.
func NewminiwindowImage() miniwindowImage {
	return miniwindowImageClass.New()
}

// Init initializes the instance.
func (m_ miniwindowImage) Init() miniwindowImage {
	rv := objc.Send[miniwindowImage](m_.ID(), selInit)
	return rv
}
