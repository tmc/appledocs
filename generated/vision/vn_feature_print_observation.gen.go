// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FeaturePrintObservation] class.
var (
	FeaturePrintObservationClass     _FeaturePrintObservationClass
	FeaturePrintObservationClassOnce sync.Once
)

func getFeaturePrintObservationClass() _FeaturePrintObservationClass {
	FeaturePrintObservationClassOnce.Do(func() {
		FeaturePrintObservationClass = _FeaturePrintObservationClass{objc.GetClass("VNFeaturePrintObservation")}
	})
	return FeaturePrintObservationClass
}

type _FeaturePrintObservationClass struct {
	class objc.Class
}

// An interface definition for the [FeaturePrintObservation] class.
type IFeaturePrintObservation interface {
	IObservation
}

// An observation that provides the recognized feature print.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNFeaturePrintObservation
type FeaturePrintObservation struct {
	Observation
}

// FeaturePrintObservationFrom constructs a [FeaturePrintObservation] from an unsafe.Pointer.
//
// An observation that provides the recognized feature print.
func FeaturePrintObservationFrom(ptr unsafe.Pointer) FeaturePrintObservation {
	return FeaturePrintObservation{
		Observation: ObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FeaturePrintObservationClass) Alloc() FeaturePrintObservation {
	rv := objc.Send[FeaturePrintObservation](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FeaturePrintObservationClass) New() FeaturePrintObservation {
	rv := objc.Send[FeaturePrintObservation](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FeaturePrintObservation) Init() FeaturePrintObservation {
	rv := objc.Send[FeaturePrintObservation](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FeaturePrintObservation) Autorelease() FeaturePrintObservation {
	rv := objc.Send[FeaturePrintObservation](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFeaturePrintObservation creates a new FeaturePrintObservation instance.
func NewFeaturePrintObservation() FeaturePrintObservation {
	return getFeaturePrintObservationClass().New()
}




