// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNBoundsPrediction */


/* debug [class_header]: Header for CNBoundsPrediction */
// The class instance for the [CNBoundsPrediction] class.
var (
	CNBoundsPredictionClass     _CNBoundsPredictionClass
	CNBoundsPredictionClassOnce sync.Once
)

func getCNBoundsPredictionClass() _CNBoundsPredictionClass {
	CNBoundsPredictionClassOnce.Do(func() {
		CNBoundsPredictionClass = _CNBoundsPredictionClass{objc.GetClass("CNBoundsPrediction")}
	})
	return CNBoundsPredictionClass
}

type _CNBoundsPredictionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNBoundsPrediction */
// An interface definition for the [CNBoundsPrediction] class.
type ICNBoundsPrediction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNBoundsPrediction */
	// properties:
	Confidence() float32
	SetConfidence(value float32)
	NormalizedBounds() corefoundation.CGRect
	SetNormalizedBounds(value corefoundation.CGRect)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNBoundsPrediction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNBoundsPrediction */
// Alloc allocates a new instance without initialization.
func (cc _CNBoundsPredictionClass) Alloc() CNBoundsPrediction {
	rv := objc.Send[CNBoundsPrediction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNBoundsPredictionClass) New() CNBoundsPrediction {
	rv := objc.Send[CNBoundsPrediction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNBoundsPrediction) Init() CNBoundsPrediction {
	rv := objc.Send[CNBoundsPrediction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNBoundsPrediction) Autorelease() CNBoundsPrediction {
	rv := objc.Send[CNBoundsPrediction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNBoundsPrediction creates a new CNBoundsPrediction instance.
func NewCNBoundsPrediction() CNBoundsPrediction {
	return getCNBoundsPredictionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNBoundsPrediction */
// An object representing the bounds of the predicted subject.


// An object representing the bounds of the predicted subject.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNBoundsPrediction-c.class
type CNBoundsPrediction struct {
	objectivec.Object
}

// CNBoundsPredictionFrom constructs a [CNBoundsPrediction] from an unsafe.Pointer.
//
// An object representing the bounds of the predicted subject.
func CNBoundsPredictionFrom(ptr unsafe.Pointer) CNBoundsPrediction {
	return CNBoundsPrediction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNBoundsPrediction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNBoundsPrediction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNBoundsPrediction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNBoundsPrediction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNBoundsPrediction */

// A number between 0.0 and 1.0 representing the probability that a defined object is within the bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNBoundsPrediction-c.class/confidence
func (c_ CNBoundsPrediction) Confidence() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("confidence"))
	return rv
}/* debug [instance_properties/getter]: confidence */


// A number between 0.0 and 1.0 representing the probability that a defined object is within the bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNBoundsPrediction-c.class/confidence
func (c_ CNBoundsPrediction) SetConfidence(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfidence:"), value)
}/* debug [instance_properties/setter]: confidence */


// The bounds of the detected object in normalized coordinates where (0.0, 0.0) is the upper-left corner, and (1.0, 1.0) is the lower-right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNBoundsPrediction-c.class/normalizedBounds
func (c_ CNBoundsPrediction) NormalizedBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("normalizedBounds"))
	return rv
}/* debug [instance_properties/getter]: normalizedBounds */


// The bounds of the detected object in normalized coordinates where (0.0, 0.0) is the upper-left corner, and (1.0, 1.0) is the lower-right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNBoundsPrediction-c.class/normalizedBounds
func (c_ CNBoundsPrediction) SetNormalizedBounds(value corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNormalizedBounds:"), value)
}/* debug [instance_properties/setter]: normalizedBounds */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNBoundsPrediction */



