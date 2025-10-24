// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
)

// The class instance for the [RecognizedPointsObservation] class.
var (
	RecognizedPointsObservationClass     _RecognizedPointsObservationClass
	RecognizedPointsObservationClassOnce sync.Once
)

func getRecognizedPointsObservationClass() _RecognizedPointsObservationClass {
	RecognizedPointsObservationClassOnce.Do(func() {
		RecognizedPointsObservationClass = _RecognizedPointsObservationClass{objc.GetClass("VNRecognizedPointsObservation")}
	})
	return RecognizedPointsObservationClass
}

type _RecognizedPointsObservationClass struct {
	class objc.Class
}

// An interface definition for the [RecognizedPointsObservation] class.
type IRecognizedPointsObservation interface {
	IObservation
	// properties:
	AvailableGroupKeys() objc.IObject /* cross-framework: RecognizedPointGroupKey */
	SetAvailableGroupKeys(value objc.IObject /* cross-framework: RecognizedPointGroupKey */)
	AvailableKeys() objc.IObject /* cross-framework: RecognizedPointKey */
	SetAvailableKeys(value objc.IObject /* cross-framework: RecognizedPointKey */)
	// methods:
	KeypointsMultiArrayAndReturnError(error_ unsafe.Pointer) objc.IObject /* cross-framework: MultiArray */
}

// An observation that provides the points the analysis recognized.


// An observation that provides the points the analysis recognized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPointsObservation
type RecognizedPointsObservation struct {
	Observation
}

// RecognizedPointsObservationFrom constructs a [RecognizedPointsObservation] from an unsafe.Pointer.
//
// An observation that provides the points the analysis recognized.
func RecognizedPointsObservationFrom(ptr unsafe.Pointer) RecognizedPointsObservation {
	return RecognizedPointsObservation{
		Observation: ObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RecognizedPointsObservationClass) Alloc() RecognizedPointsObservation {
	rv := objc.Send[RecognizedPointsObservation](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RecognizedPointsObservationClass) New() RecognizedPointsObservation {
	rv := objc.Send[RecognizedPointsObservation](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedPointsObservation) Init() RecognizedPointsObservation {
	rv := objc.Send[RecognizedPointsObservation](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedPointsObservation) Autorelease() RecognizedPointsObservation {
	rv := objc.Send[RecognizedPointsObservation](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedPointsObservation creates a new RecognizedPointsObservation instance.
func NewRecognizedPointsObservation() RecognizedPointsObservation {
	return getRecognizedPointsObservationClass().New()
}



// Retrieves the grouping of normalized point coordinates and confidence scores in a format compatible with Core ML.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPointsObservation/keypointsMultiArray()
func (r_ RecognizedPointsObservation) KeypointsMultiArrayAndReturnError(error_ unsafe.Pointer) objc.IObject /* cross-framework: MultiArray */ {
	rv := objc.Send[coreml.MultiArray](r_.ID, objc.Sel("keypointsMultiArrayAndReturnError:"), error_)
	return rv
}


// The available point group keys in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointsobservation/availablegroupkeys
func (r_ RecognizedPointsObservation) AvailableGroupKeys() objc.IObject /* cross-framework: RecognizedPointGroupKey */ {
	rv := objc.Send[RecognizedPointGroupKey](r_.ID, objc.Sel("availableGroupKeys"))
	return rv
}


// The available point group keys in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointsobservation/availablegroupkeys
func (r_ RecognizedPointsObservation) SetAvailableGroupKeys(value objc.IObject /* cross-framework: RecognizedPointGroupKey */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAvailableGroupKeys:"), value)
}


// The available point keys in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointsobservation/availablekeys
func (r_ RecognizedPointsObservation) AvailableKeys() objc.IObject /* cross-framework: RecognizedPointKey */ {
	rv := objc.Send[RecognizedPointKey](r_.ID, objc.Sel("availableKeys"))
	return rv
}


// The available point keys in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizedpointsobservation/availablekeys
func (r_ RecognizedPointsObservation) SetAvailableKeys(value objc.IObject /* cross-framework: RecognizedPointKey */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAvailableKeys:"), value)
}



