// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNRecognizedPoint */


/* debug [class_header]: Header for VNRecognizedPoint */
// The class instance for the [RecognizedPoint] class.
var (
	RecognizedPointClass     _RecognizedPointClass
	RecognizedPointClassOnce sync.Once
)

func getRecognizedPointClass() _RecognizedPointClass {
	RecognizedPointClassOnce.Do(func() {
		RecognizedPointClass = _RecognizedPointClass{objc.GetClass("VNRecognizedPoint")}
	})
	return RecognizedPointClass
}

type _RecognizedPointClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecognizedPoint */
// An interface definition for the [RecognizedPoint] class.
type IRecognizedPoint interface {
	IDetectedPoint
	
/* debug [class_interface_properties]: Properties for RecognizedPoint */
	// properties:
	Identifier() RecognizedPointKey /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecognizedPoint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecognizedPoint */
// Alloc allocates a new instance without initialization.
func (rc _RecognizedPointClass) Alloc() RecognizedPoint {
	rv := objc.Send[RecognizedPoint](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecognizedPointClass) New() RecognizedPoint {
	rv := objc.Send[RecognizedPoint](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedPoint) Init() RecognizedPoint {
	rv := objc.Send[RecognizedPoint](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedPoint) Autorelease() RecognizedPoint {
	rv := objc.Send[RecognizedPoint](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedPoint creates a new RecognizedPoint instance.
func NewRecognizedPoint() RecognizedPoint {
	return getRecognizedPointClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecognizedPoint */
// An object that represents a normalized point in an image, along with an identifier label and a confidence value.


// An object that represents a normalized point in an image, along with an identifier label and a confidence value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoint
type RecognizedPoint struct {
	DetectedPoint
}

// RecognizedPointFrom constructs a [RecognizedPoint] from an unsafe.Pointer.
//
// An object that represents a normalized point in an image, along with an identifier label and a confidence value.
func RecognizedPointFrom(ptr unsafe.Pointer) RecognizedPoint {
	return RecognizedPoint{
		DetectedPoint: DetectedPointFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecognizedPoint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecognizedPoint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecognizedPoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecognizedPoint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecognizedPoint */

// The point’s identifier label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoint/identifier
func (r_ RecognizedPoint) Identifier() RecognizedPointKey /* typedef */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNRecognizedPoint */



