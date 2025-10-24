// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [DetectedPoint] class.
var (
	DetectedPointClass     _DetectedPointClass
	DetectedPointClassOnce sync.Once
)

func getDetectedPointClass() _DetectedPointClass {
	DetectedPointClassOnce.Do(func() {
		DetectedPointClass = _DetectedPointClass{objc.GetClass("VNDetectedPoint")}
	})
	return DetectedPointClass
}

type _DetectedPointClass struct {
	class objc.Class
}





// An interface definition for the [DetectedPoint] class.
type IDetectedPoint interface {
	IPoint
	

	// properties:
	Confidence() Confidence /* typedef */


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DetectedPointClass) Alloc() DetectedPoint {
	rv := objc.Send[DetectedPoint](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectedPointClass) New() DetectedPoint {
	rv := objc.Send[DetectedPoint](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectedPoint) Init() DetectedPoint {
	rv := objc.Send[DetectedPoint](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectedPoint) Autorelease() DetectedPoint {
	rv := objc.Send[DetectedPoint](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectedPoint creates a new DetectedPoint instance.
func NewDetectedPoint() DetectedPoint {
	return getDetectedPointClass().New()
}





// An object that represents a normalized point in an image, along with a confidence value.


// An object that represents a normalized point in an image, along with a confidence value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedPoint
type DetectedPoint struct {
	Point
}

// DetectedPointFrom constructs a [DetectedPoint] from an unsafe.Pointer.
//
// An object that represents a normalized point in an image, along with a confidence value.
func DetectedPointFrom(ptr unsafe.Pointer) DetectedPoint {
	return DetectedPoint{
		Point: PointFrom(ptr),
	}
}

























// A confidence score that indicates the detected point’s accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedPoint/confidence
func (d_ DetectedPoint) Confidence() Confidence /* typedef */ {
	rv := objc.Send[float32](d_.ID, objc.Sel("confidence"))
	return rv
}








