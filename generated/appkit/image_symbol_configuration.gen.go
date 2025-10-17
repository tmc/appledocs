
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ImageSymbolConfiguration] class.
var ImageSymbolConfigurationClass _ImageSymbolConfigurationClass

func init() {
	ImageSymbolConfigurationClass = _ImageSymbolConfigurationClass{objc.GetClass("NSImageSymbolConfiguration")}
}

type _ImageSymbolConfigurationClass struct {
	objc.Class
}

// An interface definition for the [ImageSymbolConfiguration] class.
type IImageSymbolConfiguration interface {
	ID() objc.ID
}

type ImageSymbolConfiguration struct {
	id objc.ID
}

func ImageSymbolConfigurationFrom(ptr unsafe.Pointer) ImageSymbolConfiguration {
	return ImageSymbolConfiguration{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ ImageSymbolConfiguration) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _ImageSymbolConfigurationClass) Alloc() ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _ImageSymbolConfigurationClass) New() ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewImageSymbolConfiguration creates and returns a new initialized instance.
func NewImageSymbolConfiguration() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.New()
}

// Init initializes the instance.
func (i_ ImageSymbolConfiguration) Init() ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](i_.ID(), selInit)
	return rv
}
