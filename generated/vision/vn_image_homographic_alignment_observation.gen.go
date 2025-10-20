// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageHomographicAlignmentObservation] class.
var (
	ImageHomographicAlignmentObservationClass     _ImageHomographicAlignmentObservationClass
	ImageHomographicAlignmentObservationClassOnce sync.Once
)

func getImageHomographicAlignmentObservationClass() _ImageHomographicAlignmentObservationClass {
	ImageHomographicAlignmentObservationClassOnce.Do(func() {
		ImageHomographicAlignmentObservationClass = _ImageHomographicAlignmentObservationClass{objc.GetClass("VNImageHomographicAlignmentObservation")}
	})
	return ImageHomographicAlignmentObservationClass
}

type _ImageHomographicAlignmentObservationClass struct {
	class objc.Class
}

// An interface definition for the [ImageHomographicAlignmentObservation] class.
type IImageHomographicAlignmentObservation interface {
	IImageAlignmentObservation
}

// An object that represents a perspective warp transformation.
//
// This type of observation results from a , informing the performed to align the input images.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageHomographicAlignmentObservation
type ImageHomographicAlignmentObservation struct {
	ImageAlignmentObservation
}

// ImageHomographicAlignmentObservationFrom constructs a [ImageHomographicAlignmentObservation] from an unsafe.Pointer.
//
// An object that represents a perspective warp transformation.
func ImageHomographicAlignmentObservationFrom(ptr unsafe.Pointer) ImageHomographicAlignmentObservation {
	return ImageHomographicAlignmentObservation{
		ImageAlignmentObservation: ImageAlignmentObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageHomographicAlignmentObservationClass) Alloc() ImageHomographicAlignmentObservation {
	rv := objc.Send[ImageHomographicAlignmentObservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageHomographicAlignmentObservationClass) New() ImageHomographicAlignmentObservation {
	rv := objc.Send[ImageHomographicAlignmentObservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageHomographicAlignmentObservation) Init() ImageHomographicAlignmentObservation {
	rv := objc.Send[ImageHomographicAlignmentObservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageHomographicAlignmentObservation) Autorelease() ImageHomographicAlignmentObservation {
	rv := objc.Send[ImageHomographicAlignmentObservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageHomographicAlignmentObservation creates a new ImageHomographicAlignmentObservation instance.
func NewImageHomographicAlignmentObservation() ImageHomographicAlignmentObservation {
	return getImageHomographicAlignmentObservationClass().New()
}


// The warp transform matrix to morph the floating image into the reference image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageHomographicAlignmentObservation/warpTransform
func (i_ ImageHomographicAlignmentObservation) WarpTransform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("warpTransform"))
	return rv
}



