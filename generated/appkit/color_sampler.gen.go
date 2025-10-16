
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorSampler] class.
var ColorSamplerClass _ColorSamplerClass

func init() {
	ColorSamplerClass = _ColorSamplerClass{objc.GetClass("NSColorSampler")}
}

type _ColorSamplerClass struct {
	objc.Class
}

// An interface definition for the [ColorSampler] class.
type IColorSampler interface {
	ID() objc.ID
}

type ColorSampler struct {
	id objc.ID
}

func ColorSamplerFrom(ptr unsafe.Pointer) ColorSampler {
	return ColorSampler{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ColorSampler) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ColorSamplerClass) Alloc() ColorSampler {
	rv := objc.Send[ColorSampler](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ColorSamplerClass) New() ColorSampler {
	rv := objc.Send[ColorSampler](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewColorSampler creates and returns a new initialized instance.
func NewColorSampler() ColorSampler {
	return ColorSamplerClass.New()
}

// Init initializes the instance.
func (c_ ColorSampler) Init() ColorSampler {
	rv := objc.Send[ColorSampler](c_.ID(), selInit)
	return rv
}
