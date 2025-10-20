// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageAlignmentObservation] class.
var (
	ImageAlignmentObservationClass     _ImageAlignmentObservationClass
	ImageAlignmentObservationClassOnce sync.Once
)

func getImageAlignmentObservationClass() _ImageAlignmentObservationClass {
	ImageAlignmentObservationClassOnce.Do(func() {
		ImageAlignmentObservationClass = _ImageAlignmentObservationClass{objc.GetClass("VNImageAlignmentObservation")}
	})
	return ImageAlignmentObservationClass
}

type _ImageAlignmentObservationClass struct {
	class objc.Class
}

// An interface definition for the [ImageAlignmentObservation] class.
type IImageAlignmentObservation interface {
	IObservation
}

// The abstract superclass for image-analysis results that describe the relative alignment of two images.
//
// This abstract superclass forms the basis of image alignment or registration output. You receive its subclasses, such as and , by performing specific registration requests. Don’t create one of these classes yourself.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageAlignmentObservation
type ImageAlignmentObservation struct {
	Observation
}

// ImageAlignmentObservationFrom constructs a [ImageAlignmentObservation] from an unsafe.Pointer.
//
// The abstract superclass for image-analysis results that describe the relative alignment of two images.
func ImageAlignmentObservationFrom(ptr unsafe.Pointer) ImageAlignmentObservation {
	return ImageAlignmentObservation{
		Observation: ObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageAlignmentObservationClass) Alloc() ImageAlignmentObservation {
	rv := objc.Send[ImageAlignmentObservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageAlignmentObservationClass) New() ImageAlignmentObservation {
	rv := objc.Send[ImageAlignmentObservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageAlignmentObservation) Init() ImageAlignmentObservation {
	rv := objc.Send[ImageAlignmentObservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageAlignmentObservation) Autorelease() ImageAlignmentObservation {
	rv := objc.Send[ImageAlignmentObservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageAlignmentObservation creates a new ImageAlignmentObservation instance.
func NewImageAlignmentObservation() ImageAlignmentObservation {
	return getImageAlignmentObservationClass().New()
}




