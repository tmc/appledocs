// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


// The feature print data.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfeatureprintobservation/data
func (f_ FeaturePrintObservation) Data() foundation.Data {
	rv := objc.Send[foundation.Data](f_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
// The feature print data.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfeatureprintobservation/data
func (f_ FeaturePrintObservation) SetData(value foundation.IData) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setData:"), value)
}

// The total number of elements in the data.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfeatureprintobservation/elementcount
func (f_ FeaturePrintObservation) ElementCount() int {
	rv := objc.Send[int](f_.ID, objc.Sel("elementCount"))
	return rv
}


// SetElementCount sets the value of the elementCount property.
// The total number of elements in the data.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfeatureprintobservation/elementcount
func (f_ FeaturePrintObservation) SetElementCount(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setElementCount:"), value)
}

// The type of each element in the data.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfeatureprintobservation/elementtype
func (f_ FeaturePrintObservation) ElementType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("elementType"))
	return rv
}


// SetElementType sets the value of the elementType property.
// The type of each element in the data.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfeatureprintobservation/elementtype
func (f_ FeaturePrintObservation) SetElementType(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setElementType:"), value)
}



