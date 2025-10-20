// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HumanBodyPose3DObservation] class.
var (
	HumanBodyPose3DObservationClass     _HumanBodyPose3DObservationClass
	HumanBodyPose3DObservationClassOnce sync.Once
)

func getHumanBodyPose3DObservationClass() _HumanBodyPose3DObservationClass {
	HumanBodyPose3DObservationClassOnce.Do(func() {
		HumanBodyPose3DObservationClass = _HumanBodyPose3DObservationClass{objc.GetClass("VNHumanBodyPose3DObservation")}
	})
	return HumanBodyPose3DObservationClass
}

type _HumanBodyPose3DObservationClass struct {
	class objc.Class
}

// An interface definition for the [HumanBodyPose3DObservation] class.
type IHumanBodyPose3DObservation interface {
	IRecognizedPoints3DObservation
}

// An observation that provides the 3D body points the request recognizes.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation
type HumanBodyPose3DObservation struct {
	RecognizedPoints3DObservation
}

// HumanBodyPose3DObservationFrom constructs a [HumanBodyPose3DObservation] from an unsafe.Pointer.
//
// An observation that provides the 3D body points the request recognizes.
func HumanBodyPose3DObservationFrom(ptr unsafe.Pointer) HumanBodyPose3DObservation {
	return HumanBodyPose3DObservation{
		RecognizedPoints3DObservation: RecognizedPoints3DObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HumanBodyPose3DObservationClass) Alloc() HumanBodyPose3DObservation {
	rv := objc.Send[HumanBodyPose3DObservation](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HumanBodyPose3DObservationClass) New() HumanBodyPose3DObservation {
	rv := objc.Send[HumanBodyPose3DObservation](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HumanBodyPose3DObservation) Init() HumanBodyPose3DObservation {
	rv := objc.Send[HumanBodyPose3DObservation](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HumanBodyPose3DObservation) Autorelease() HumanBodyPose3DObservation {
	rv := objc.Send[HumanBodyPose3DObservation](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHumanBodyPose3DObservation creates a new HumanBodyPose3DObservation instance.
func NewHumanBodyPose3DObservation() HumanBodyPose3DObservation {
	return getHumanBodyPose3DObservationClass().New()
}




