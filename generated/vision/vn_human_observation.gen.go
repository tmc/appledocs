// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [HumanObservation] class.
var (
	HumanObservationClass     _HumanObservationClass
	HumanObservationClassOnce sync.Once
)

func getHumanObservationClass() _HumanObservationClass {
	HumanObservationClassOnce.Do(func() {
		HumanObservationClass = _HumanObservationClass{objc.GetClass("VNHumanObservation")}
	})
	return HumanObservationClass
}

type _HumanObservationClass struct {
	class objc.Class
}





// An interface definition for the [HumanObservation] class.
type IHumanObservation interface {
	IDetectedObjectObservation
	

	// properties:
	UpperBodyOnly() bool


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (hc _HumanObservationClass) Alloc() HumanObservation {
	rv := objc.Send[HumanObservation](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HumanObservationClass) New() HumanObservation {
	rv := objc.Send[HumanObservation](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HumanObservation) Init() HumanObservation {
	rv := objc.Send[HumanObservation](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HumanObservation) Autorelease() HumanObservation {
	rv := objc.Send[HumanObservation](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHumanObservation creates a new HumanObservation instance.
func NewHumanObservation() HumanObservation {
	return getHumanObservationClass().New()
}





// An object that represents a person that the request detects.


// An object that represents a person that the request detects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanObservation
type HumanObservation struct {
	DetectedObjectObservation
}

// HumanObservationFrom constructs a [HumanObservation] from an unsafe.Pointer.
//
// An object that represents a person that the request detects.
func HumanObservationFrom(ptr unsafe.Pointer) HumanObservation {
	return HumanObservation{
		DetectedObjectObservation: DetectedObjectObservationFrom(ptr),
	}
}

























// A Boolean value that indicates whether the observation represents an upper-body or full-body rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanObservation/upperBodyOnly
func (h_ HumanObservation) UpperBodyOnly() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("upperBodyOnly"))
	return rv
}








