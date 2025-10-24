// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class faceAngle */


/* debug [class_header]: Header for faceAngle */
// The class instance for the [faceAngle] class.
var (
	FaceAngleClass     _faceAngleClass
	FaceAngleClassOnce sync.Once
)

func getfaceAngleClass() _faceAngleClass {
	FaceAngleClassOnce.Do(func() {
		FaceAngleClass = _faceAngleClass{objc.GetClass("faceAngle")}
	})
	return FaceAngleClass
}

type _faceAngleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for faceAngle */
// An interface definition for the [faceAngle] class.
type IfaceAngle interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for faceAngle */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for faceAngle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for faceAngle */
// Alloc allocates a new instance without initialization.
func (fc _faceAngleClass) Alloc() faceAngle {
	rv := objc.Send[faceAngle](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _faceAngleClass) New() faceAngle {
	rv := objc.Send[faceAngle](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ faceAngle) Init() faceAngle {
	rv := objc.Send[faceAngle](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ faceAngle) Autorelease() faceAngle {
	rv := objc.Send[faceAngle](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewfaceAngle creates a new faceAngle instance.
func NewfaceAngle() faceAngle {
	return getfaceAngleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for faceAngle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/faceAngle-c.ivar
type faceAngle struct {
	objectivec.Object
}

// faceAngleFrom constructs a [faceAngle] from an unsafe.Pointer.
func faceAngleFrom(ptr unsafe.Pointer) faceAngle {
	return faceAngle{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for faceAngle *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for faceAngle */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for faceAngle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for faceAngle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for faceAngle */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class faceAngle */



