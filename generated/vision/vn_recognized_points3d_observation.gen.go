// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RecognizedPoints3DObservation] class.
var (
	RecognizedPoints3DObservationClass     _RecognizedPoints3DObservationClass
	RecognizedPoints3DObservationClassOnce sync.Once
)

func getRecognizedPoints3DObservationClass() _RecognizedPoints3DObservationClass {
	RecognizedPoints3DObservationClassOnce.Do(func() {
		RecognizedPoints3DObservationClass = _RecognizedPoints3DObservationClass{objc.GetClass("VNRecognizedPoints3DObservation")}
	})
	return RecognizedPoints3DObservationClass
}

type _RecognizedPoints3DObservationClass struct {
	class objc.Class
}

// An interface definition for the [RecognizedPoints3DObservation] class.
type IRecognizedPoints3DObservation interface {
	IObservation
	// properties:
	AvailableGroupKeys() objc.IObject /* cross-framework: RecognizedPointGroupKey */
	SetAvailableGroupKeys(value objc.IObject /* cross-framework: RecognizedPointGroupKey */)
	AvailableKeys() objc.IObject /* cross-framework: RecognizedPointKey */
	SetAvailableKeys(value objc.IObject /* cross-framework: RecognizedPointKey */)
	// methods:
}

// An observation that provides the 3D points for a request.


// An observation that provides the 3D points for a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoints3DObservation
type RecognizedPoints3DObservation struct {
	Observation
}

// RecognizedPoints3DObservationFrom constructs a [RecognizedPoints3DObservation] from an unsafe.Pointer.
//
// An observation that provides the 3D points for a request.
func RecognizedPoints3DObservationFrom(ptr unsafe.Pointer) RecognizedPoints3DObservation {
	return RecognizedPoints3DObservation{
		Observation: ObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RecognizedPoints3DObservationClass) Alloc() RecognizedPoints3DObservation {
	rv := objc.Send[RecognizedPoints3DObservation](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RecognizedPoints3DObservationClass) New() RecognizedPoints3DObservation {
	rv := objc.Send[RecognizedPoints3DObservation](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedPoints3DObservation) Init() RecognizedPoints3DObservation {
	rv := objc.Send[RecognizedPoints3DObservation](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedPoints3DObservation) Autorelease() RecognizedPoints3DObservation {
	rv := objc.Send[RecognizedPoints3DObservation](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedPoints3DObservation creates a new RecognizedPoints3DObservation instance.
func NewRecognizedPoints3DObservation() RecognizedPoints3DObservation {
	return getRecognizedPoints3DObservationClass().New()
}



// The available point group keys in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpoints3dobservation/availablegroupkeys
func (r_ RecognizedPoints3DObservation) AvailableGroupKeys() objc.IObject /* cross-framework: RecognizedPointGroupKey */ {
	rv := objc.Send[RecognizedPointGroupKey](r_.ID, objc.Sel("availableGroupKeys"))
	return rv
}


// The available point group keys in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpoints3dobservation/availablegroupkeys
func (r_ RecognizedPoints3DObservation) SetAvailableGroupKeys(value objc.IObject /* cross-framework: RecognizedPointGroupKey */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAvailableGroupKeys:"), value)
}


// The available point keys in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpoints3dobservation/availablekeys
func (r_ RecognizedPoints3DObservation) AvailableKeys() objc.IObject /* cross-framework: RecognizedPointKey */ {
	rv := objc.Send[RecognizedPointKey](r_.ID, objc.Sel("availableKeys"))
	return rv
}


// The available point keys in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpoints3dobservation/availablekeys
func (r_ RecognizedPoints3DObservation) SetAvailableKeys(value objc.IObject /* cross-framework: RecognizedPointKey */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAvailableKeys:"), value)
}



