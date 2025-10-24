// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Observation] class.
var (
	ObservationClass     _ObservationClass
	ObservationClassOnce sync.Once
)

func getObservationClass() _ObservationClass {
	ObservationClassOnce.Do(func() {
		ObservationClass = _ObservationClass{objc.GetClass("VNObservation")}
	})
	return ObservationClass
}

type _ObservationClass struct {
	class objc.Class
}





// An interface definition for the [Observation] class.
type IObservation interface {
	objectivec.IObject
	

	// properties:
	Confidence() Confidence /* typedef */
	TimeRange() TimeRange /* not a class type */
	Uuid() foundation.UUID


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (oc _ObservationClass) Alloc() Observation {
	rv := objc.Send[Observation](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _ObservationClass) New() Observation {
	rv := objc.Send[Observation](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Observation) Init() Observation {
	rv := objc.Send[Observation](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Observation) Autorelease() Observation {
	rv := objc.Send[Observation](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewObservation creates a new Observation instance.
func NewObservation() Observation {
	return getObservationClass().New()
}





// The abstract superclass for analysis results.
//
// Observations resulting from Vision image analysis requests inherit from this abstract base class. Don’t use this abstract superclass directly.


// The abstract superclass for analysis results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNObservation
type Observation struct {
	objectivec.Object
}

// ObservationFrom constructs a [Observation] from an unsafe.Pointer.
//
// The abstract superclass for analysis results.
func ObservationFrom(ptr unsafe.Pointer) Observation {
	return Observation{objectivec.Object{objc.ID(ptr)}}
}

























// The level of confidence in the observation’s accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNObservation/confidence
func (o_ Observation) Confidence() Confidence /* typedef */ {
	rv := objc.Send[float32](o_.ID, objc.Sel("confidence"))
	return rv
}


// The time range of the reported observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNObservation/timeRange
func (o_ Observation) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](o_.ID, objc.Sel("timeRange"))
	return rv
}


// A unique identifier assigned to the Vision observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNObservation/uuid
func (o_ Observation) Uuid() foundation.UUID {
	rv := objc.Send[foundation.UUID](o_.ID, objc.Sel("uuid"))
	return rv
}








