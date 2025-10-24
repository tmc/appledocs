// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNRecognizedObjectObservation */


/* debug [class_header]: Header for VNRecognizedObjectObservation */
// The class instance for the [RecognizedObjectObservation] class.
var (
	RecognizedObjectObservationClass     _RecognizedObjectObservationClass
	RecognizedObjectObservationClassOnce sync.Once
)

func getRecognizedObjectObservationClass() _RecognizedObjectObservationClass {
	RecognizedObjectObservationClassOnce.Do(func() {
		RecognizedObjectObservationClass = _RecognizedObjectObservationClass{objc.GetClass("VNRecognizedObjectObservation")}
	})
	return RecognizedObjectObservationClass
}

type _RecognizedObjectObservationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecognizedObjectObservation */
// An interface definition for the [RecognizedObjectObservation] class.
type IRecognizedObjectObservation interface {
	IDetectedObjectObservation
	
/* debug [class_interface_properties]: Properties for RecognizedObjectObservation */
	// properties:
	Labels() []ClassificationObservation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecognizedObjectObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecognizedObjectObservation */
// Alloc allocates a new instance without initialization.
func (rc _RecognizedObjectObservationClass) Alloc() RecognizedObjectObservation {
	rv := objc.Send[RecognizedObjectObservation](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecognizedObjectObservationClass) New() RecognizedObjectObservation {
	rv := objc.Send[RecognizedObjectObservation](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedObjectObservation) Init() RecognizedObjectObservation {
	rv := objc.Send[RecognizedObjectObservation](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedObjectObservation) Autorelease() RecognizedObjectObservation {
	rv := objc.Send[RecognizedObjectObservation](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedObjectObservation creates a new RecognizedObjectObservation instance.
func NewRecognizedObjectObservation() RecognizedObjectObservation {
	return getRecognizedObjectObservationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecognizedObjectObservation */
// A detected object observation with an array of classification labels that classify the recognized object.
//
// The confidence of the classifications sum up to Multiply the classification confidence with the confidence of this observation.


// A detected object observation with an array of classification labels that classify the recognized object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedObjectObservation
type RecognizedObjectObservation struct {
	DetectedObjectObservation
}

// RecognizedObjectObservationFrom constructs a [RecognizedObjectObservation] from an unsafe.Pointer.
//
// A detected object observation with an array of classification labels that classify the recognized object.
func RecognizedObjectObservationFrom(ptr unsafe.Pointer) RecognizedObjectObservation {
	return RecognizedObjectObservation{
		DetectedObjectObservation: DetectedObjectObservationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecognizedObjectObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecognizedObjectObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecognizedObjectObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecognizedObjectObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecognizedObjectObservation */

// An array of observations that classify the recognized object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedObjectObservation/labels
func (r_ RecognizedObjectObservation) Labels() []ClassificationObservation {
	rv := objc.Send[[]ClassificationObservation](r_.ID, objc.Sel("labels"))
	return rv
}/* debug [instance_properties/getter]: labels */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNRecognizedObjectObservation */



