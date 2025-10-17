
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Image] class.
var ImageClass _ImageClass

func init() {
	ImageClass = _ImageClass{objc.GetClass("NSImage")}
}

type _ImageClass struct {
	objc.Class
}

// An interface definition for the [Image] class.
type IImage interface {
	ID() objc.ID
	BestRepresentationForDevice(deviceDescription unsafe.Pointer) unsafe.Pointer
}

type Image struct {
	id objc.ID
}

func ImageFrom(ptr unsafe.Pointer) Image {
	return Image{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ Image) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _ImageClass) Alloc() Image {
	rv := objc.Send[Image](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _ImageClass) New() Image {
	rv := objc.Send[Image](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewImage creates and returns a new initialized instance.
func NewImage() Image {
	return ImageClass.New()
}

// Init initializes the instance.
func (i_ Image) Init() Image {
	rv := objc.Send[Image](i_.ID(), selInit)
	return rv
}
// Returns the best representation for the device with the specified characteristics. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSImage/bestRepresentationForDevice:
func (i_ Image) BestRepresentationForDevice(deviceDescription unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID(), objc.RegisterName("bestRepresentationForDevice:"), deviceDescription)
	return rv
}
