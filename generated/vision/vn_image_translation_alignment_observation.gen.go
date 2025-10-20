// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageTranslationAlignmentObservation] class.
var (
	ImageTranslationAlignmentObservationClass     _ImageTranslationAlignmentObservationClass
	ImageTranslationAlignmentObservationClassOnce sync.Once
)

func getImageTranslationAlignmentObservationClass() _ImageTranslationAlignmentObservationClass {
	ImageTranslationAlignmentObservationClassOnce.Do(func() {
		ImageTranslationAlignmentObservationClass = _ImageTranslationAlignmentObservationClass{objc.GetClass("VNImageTranslationAlignmentObservation")}
	})
	return ImageTranslationAlignmentObservationClass
}

type _ImageTranslationAlignmentObservationClass struct {
	class objc.Class
}

// An interface definition for the [ImageTranslationAlignmentObservation] class.
type IImageTranslationAlignmentObservation interface {
	IImageAlignmentObservation
}

// Affine transform information that an image-alignment request produces.
//
// This type of observation results from a , informing the performed to align the input images.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageTranslationAlignmentObservation
type ImageTranslationAlignmentObservation struct {
	ImageAlignmentObservation
}

// ImageTranslationAlignmentObservationFrom constructs a [ImageTranslationAlignmentObservation] from an unsafe.Pointer.
//
// Affine transform information that an image-alignment request produces.
func ImageTranslationAlignmentObservationFrom(ptr unsafe.Pointer) ImageTranslationAlignmentObservation {
	return ImageTranslationAlignmentObservation{
		ImageAlignmentObservation: ImageAlignmentObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageTranslationAlignmentObservationClass) Alloc() ImageTranslationAlignmentObservation {
	rv := objc.Send[ImageTranslationAlignmentObservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageTranslationAlignmentObservationClass) New() ImageTranslationAlignmentObservation {
	rv := objc.Send[ImageTranslationAlignmentObservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageTranslationAlignmentObservation) Init() ImageTranslationAlignmentObservation {
	rv := objc.Send[ImageTranslationAlignmentObservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageTranslationAlignmentObservation) Autorelease() ImageTranslationAlignmentObservation {
	rv := objc.Send[ImageTranslationAlignmentObservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageTranslationAlignmentObservation creates a new ImageTranslationAlignmentObservation instance.
func NewImageTranslationAlignmentObservation() ImageTranslationAlignmentObservation {
	return getImageTranslationAlignmentObservationClass().New()
}




