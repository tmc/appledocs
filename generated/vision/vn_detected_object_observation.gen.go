// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DetectedObjectObservation] class.
var (
	DetectedObjectObservationClass     _DetectedObjectObservationClass
	DetectedObjectObservationClassOnce sync.Once
)

func getDetectedObjectObservationClass() _DetectedObjectObservationClass {
	DetectedObjectObservationClassOnce.Do(func() {
		DetectedObjectObservationClass = _DetectedObjectObservationClass{objc.GetClass("VNDetectedObjectObservation")}
	})
	return DetectedObjectObservationClass
}

type _DetectedObjectObservationClass struct {
	class objc.Class
}

// An interface definition for the [DetectedObjectObservation] class.
type IDetectedObjectObservation interface {
	IObservation
}

// An observation that provides the position and extent of an image feature that an image- analysis request detects.
//
// This class is the observation type that generates. It represents an object that the Vision request detects and tracks.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedObjectObservation
type DetectedObjectObservation struct {
	Observation
}

// DetectedObjectObservationFrom constructs a [DetectedObjectObservation] from an unsafe.Pointer.
//
// An observation that provides the position and extent of an image feature that an image- analysis request detects.
func DetectedObjectObservationFrom(ptr unsafe.Pointer) DetectedObjectObservation {
	return DetectedObjectObservation{
		Observation: ObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DetectedObjectObservationClass) Alloc() DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DetectedObjectObservationClass) New() DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectedObjectObservation) Init() DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectedObjectObservation) Autorelease() DetectedObjectObservation {
	rv := objc.Send[DetectedObjectObservation](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectedObjectObservation creates a new DetectedObjectObservation instance.
func NewDetectedObjectObservation() DetectedObjectObservation {
	return getDetectedObjectObservationClass().New()
}




