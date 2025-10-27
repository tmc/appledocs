// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [HorizonObservation] class.
var (
	HorizonObservationClass     _HorizonObservationClass
	HorizonObservationClassOnce sync.Once
)

func getHorizonObservationClass() _HorizonObservationClass {
	HorizonObservationClassOnce.Do(func() {
		HorizonObservationClass = _HorizonObservationClass{objc.GetClass("VNHorizonObservation")}
	})
	return HorizonObservationClass
}

type _HorizonObservationClass struct {
	class objc.Class
}





// An interface definition for the [HorizonObservation] class.
type IHorizonObservation interface {
	IObservation
	

	// properties:
	Angle() float64
	Transform() corefoundation.CGAffineTransform


	

	// methods:
	TransformForImageWidthHeight(width uintptr /* not a class type */, height uintptr /* not a class type */) corefoundation.CGAffineTransform


}





// Alloc allocates a new instance without initialization.
func (hc _HorizonObservationClass) Alloc() HorizonObservation {
	rv := objc.Send[HorizonObservation](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HorizonObservationClass) New() HorizonObservation {
	rv := objc.Send[HorizonObservation](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HorizonObservation) Init() HorizonObservation {
	rv := objc.Send[HorizonObservation](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HorizonObservation) Autorelease() HorizonObservation {
	rv := objc.Send[HorizonObservation](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHorizonObservation creates a new HorizonObservation instance.
func NewHorizonObservation() HorizonObservation {
	return getHorizonObservationClass().New()
}





// The horizon angle information that an image-analysis request detects.
//
// Instances of this class result from invoking a , and report the and of the horizon in an image.


// The horizon angle information that an image-analysis request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHorizonObservation
type HorizonObservation struct {
	Observation
}

// HorizonObservationFrom constructs a [HorizonObservation] from an unsafe.Pointer.
//
// The horizon angle information that an image-analysis request detects.
func HorizonObservationFrom(ptr unsafe.Pointer) HorizonObservation {
	return HorizonObservation{
		Observation: ObservationFrom(ptr),
	}
}




















// Creates an affine transform for the specified image width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHorizonObservation/transform(forImageWidth:height:)
func (h_ HorizonObservation) TransformForImageWidthHeight(width uintptr /* not a class type */, height uintptr /* not a class type */) corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](h_.ID, objc.Sel("transformForImageWidth:height:"), width, height)
	return rv
}







// The angle of the observed horizon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHorizonObservation/angle
func (h_ HorizonObservation) Angle() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("angle"))
	return rv
}


// The transform to apply to the detected horizon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHorizonObservation/transform
func (h_ HorizonObservation) Transform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](h_.ID, objc.Sel("transform"))
	return rv
}








