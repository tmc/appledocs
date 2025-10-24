// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNRecognizedPoint3D */


/* debug [class_header]: Header for VNRecognizedPoint3D */
// The class instance for the [RecognizedPoint3D] class.
var (
	RecognizedPoint3DClass     _RecognizedPoint3DClass
	RecognizedPoint3DClassOnce sync.Once
)

func getRecognizedPoint3DClass() _RecognizedPoint3DClass {
	RecognizedPoint3DClassOnce.Do(func() {
		RecognizedPoint3DClass = _RecognizedPoint3DClass{objc.GetClass("VNRecognizedPoint3D")}
	})
	return RecognizedPoint3DClass
}

type _RecognizedPoint3DClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecognizedPoint3D */
// An interface definition for the [RecognizedPoint3D] class.
type IRecognizedPoint3D interface {
	IPoint3D
	
/* debug [class_interface_properties]: Properties for RecognizedPoint3D */
	// properties:
	Identifier() RecognizedPointKey /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecognizedPoint3D */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecognizedPoint3D */
// Alloc allocates a new instance without initialization.
func (rc _RecognizedPoint3DClass) Alloc() RecognizedPoint3D {
	rv := objc.Send[RecognizedPoint3D](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecognizedPoint3DClass) New() RecognizedPoint3D {
	rv := objc.Send[RecognizedPoint3D](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizedPoint3D) Init() RecognizedPoint3D {
	rv := objc.Send[RecognizedPoint3D](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizedPoint3D) Autorelease() RecognizedPoint3D {
	rv := objc.Send[RecognizedPoint3D](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizedPoint3D creates a new RecognizedPoint3D instance.
func NewRecognizedPoint3D() RecognizedPoint3D {
	return getRecognizedPoint3DClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecognizedPoint3D */
// A 3D point that includes an identifier to the point.


// A 3D point that includes an identifier to the point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoint3D
type RecognizedPoint3D struct {
	Point3D
}

// RecognizedPoint3DFrom constructs a [RecognizedPoint3D] from an unsafe.Pointer.
//
// A 3D point that includes an identifier to the point.
func RecognizedPoint3DFrom(ptr unsafe.Pointer) RecognizedPoint3D {
	return RecognizedPoint3D{
		Point3D: Point3DFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecognizedPoint3D *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecognizedPoint3D */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecognizedPoint3D */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecognizedPoint3D */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecognizedPoint3D */

// The identifier that provides context about what kind of point the request recognizes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizedPoint3D/identifier
func (r_ RecognizedPoint3D) Identifier() RecognizedPointKey /* typedef */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNRecognizedPoint3D */



