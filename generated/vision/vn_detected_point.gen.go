// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNDetectedPoint */


/* debug [class_header]: Header for VNDetectedPoint */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectedPoint */
// An interface definition for the [DetectedPoint] class.
type IDetectedPoint interface {
	IPoint
	
/* debug [class_interface_properties]: Properties for DetectedPoint */
	// properties:
	Confidence() Confidence /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectedPoint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectedPoint */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectedPoint */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectedPoint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectedPoint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectedPoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectedPoint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectedPoint */

// A confidence score that indicates the detected point’s accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectedPoint/confidence
func (d_ DetectedPoint) Confidence() Confidence /* typedef */ {
	rv := objc.Send[float32](d_.ID, objc.Sel("confidence"))
	return rv
}/* debug [instance_properties/getter]: confidence */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectedPoint */



