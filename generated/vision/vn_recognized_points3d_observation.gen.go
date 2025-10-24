// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	AvailableGroupKeys() []string
	AvailableKeys() []string


	

	// methods:
	RecognizedPointForKeyError(pointKey RecognizedPointKey /* typedef */, error_ objectivec.IObject) IRecognizedPoint3D
	RecognizedPointsForGroupKeyError(groupKey RecognizedPointGroupKey /* typedef */, error_ objectivec.IObject) foundation.IDictionary


}





// Alloc allocates a new instance without initialization.
func (rc _RecognizedPoints3DObservationClass) Alloc() RecognizedPoints3DObservation {
	rv := objc.Send[RecognizedPoints3DObservation](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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




















// Returns a point for a key you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoints3DObservation/recognizedPoint(forKey:)
func (r_ RecognizedPoints3DObservation) RecognizedPointForKeyError(pointKey RecognizedPointKey /* typedef */, error_ objectivec.IObject) IRecognizedPoint3D {
	rv := objc.Send[RecognizedPoint3D](r_.ID, objc.Sel("recognizedPointForKey:error:"), pointKey, error_)
	return rv
}


// Returns a point for a group key you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoints3DObservation/recognizedPoints(forGroupKey:)
func (r_ RecognizedPoints3DObservation) RecognizedPointsForGroupKeyError(groupKey RecognizedPointGroupKey /* typedef */, error_ objectivec.IObject) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](r_.ID, objc.Sel("recognizedPointsForGroupKey:error:"), groupKey, error_)
	return rv
}







// The available point group keys in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoints3DObservation/availableGroupKeys
func (r_ RecognizedPoints3DObservation) AvailableGroupKeys() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("availableGroupKeys"))
	return rv
}


// The available point keys in the observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoints3DObservation/availableKeys
func (r_ RecognizedPoints3DObservation) AvailableKeys() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("availableKeys"))
	return rv
}








