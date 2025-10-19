// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Image] class.
var (
	imageClass     _ImageClass
	imageClassOnce sync.Once
)

func getImageClass() _ImageClass {
	imageClassOnce.Do(func() {
		imageClass = _ImageClass{objc.GetClass("NSImage")}
	})
	return imageClass
}

type _ImageClass struct {
	class objc.Class
}

// An interface definition for the [Image] class.
type IImage interface {
	objectivec.IObject
	BestRepresentationForDevice(deviceDescription unsafe.Pointer) unsafe.Pointer
}

// A high-level interface for manipulating image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage

type Image struct {
	objectivec.Object
}

// ImageFrom constructs a [Image] from an unsafe.Pointer.
//
// A high-level interface for manipulating image data.
func ImageFrom(ptr unsafe.Pointer) Image {
	return Image{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ic _ImageClass) Alloc() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ic _ImageClass) New() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ Image) Init() Image {
	rv := objc.Send[Image](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ Image) Autorelease() Image {
	rv := objc.Send[Image](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImage creates a new Image instance.
func NewImage() Image {
	return getImageClass().New()
}


// Returns the best representation for the device with the specified characteristics. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/bestRepresentationForDevice:
func (i_ Image) BestRepresentationForDevice(deviceDescription unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("bestRepresentationForDevice:"), deviceDescription)
	return rv
}


