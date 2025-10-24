// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [SaliencyImageObservation] class.
var (
	SaliencyImageObservationClass     _SaliencyImageObservationClass
	SaliencyImageObservationClassOnce sync.Once
)

func getSaliencyImageObservationClass() _SaliencyImageObservationClass {
	SaliencyImageObservationClassOnce.Do(func() {
		SaliencyImageObservationClass = _SaliencyImageObservationClass{objc.GetClass("VNSaliencyImageObservation")}
	})
	return SaliencyImageObservationClass
}

type _SaliencyImageObservationClass struct {
	class objc.Class
}





// An interface definition for the [SaliencyImageObservation] class.
type ISaliencyImageObservation interface {
	IPixelBufferObservation
	

	// properties:
	SalientObjects() []RectangleObservation


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SaliencyImageObservationClass) Alloc() SaliencyImageObservation {
	rv := objc.Send[SaliencyImageObservation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SaliencyImageObservationClass) New() SaliencyImageObservation {
	rv := objc.Send[SaliencyImageObservation](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SaliencyImageObservation) Init() SaliencyImageObservation {
	rv := objc.Send[SaliencyImageObservation](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SaliencyImageObservation) Autorelease() SaliencyImageObservation {
	rv := objc.Send[SaliencyImageObservation](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSaliencyImageObservation creates a new SaliencyImageObservation instance.
func NewSaliencyImageObservation() SaliencyImageObservation {
	return getSaliencyImageObservationClass().New()
}





// An observation that contains a grayscale heat map of important areas across an image.
//
// The heat map is a in a one-component floating-point pixel format. Its dimensions are 64 x 64 when fetched in real time, or 68 x 68 when requested in its deferred form.


// An observation that contains a grayscale heat map of important areas across an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSaliencyImageObservation
type SaliencyImageObservation struct {
	PixelBufferObservation
}

// SaliencyImageObservationFrom constructs a [SaliencyImageObservation] from an unsafe.Pointer.
//
// An observation that contains a grayscale heat map of important areas across an image.
func SaliencyImageObservationFrom(ptr unsafe.Pointer) SaliencyImageObservation {
	return SaliencyImageObservation{
		PixelBufferObservation: PixelBufferObservationFrom(ptr),
	}
}

























// A collection of objects describing the distinct areas of the saliency heat map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSaliencyImageObservation/salientObjects
func (s_ SaliencyImageObservation) SalientObjects() []RectangleObservation {
	rv := objc.Send[[]RectangleObservation](s_.ID, objc.Sel("salientObjects"))
	return rv
}








