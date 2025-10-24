// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class hasFaceAngle */


/* debug [class_header]: Header for hasFaceAngle */
// The class instance for the [hasFaceAngle] class.
var (
	HasFaceAngleClass     _hasFaceAngleClass
	HasFaceAngleClassOnce sync.Once
)

func gethasFaceAngleClass() _hasFaceAngleClass {
	HasFaceAngleClassOnce.Do(func() {
		HasFaceAngleClass = _hasFaceAngleClass{objc.GetClass("hasFaceAngle")}
	})
	return HasFaceAngleClass
}

type _hasFaceAngleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for hasFaceAngle */
// An interface definition for the [hasFaceAngle] class.
type IhasFaceAngle interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for hasFaceAngle */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for hasFaceAngle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for hasFaceAngle */
// Alloc allocates a new instance without initialization.
func (hc _hasFaceAngleClass) Alloc() hasFaceAngle {
	rv := objc.Send[hasFaceAngle](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _hasFaceAngleClass) New() hasFaceAngle {
	rv := objc.Send[hasFaceAngle](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasFaceAngle) Init() hasFaceAngle {
	rv := objc.Send[hasFaceAngle](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasFaceAngle) Autorelease() hasFaceAngle {
	rv := objc.Send[hasFaceAngle](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasFaceAngle creates a new hasFaceAngle instance.
func NewhasFaceAngle() hasFaceAngle {
	return gethasFaceAngleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for hasFaceAngle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasFaceAngle-c.ivar
type hasFaceAngle struct {
	objectivec.Object
}

// hasFaceAngleFrom constructs a [hasFaceAngle] from an unsafe.Pointer.
func hasFaceAngleFrom(ptr unsafe.Pointer) hasFaceAngle {
	return hasFaceAngle{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for hasFaceAngle *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for hasFaceAngle */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for hasFaceAngle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for hasFaceAngle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for hasFaceAngle */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class hasFaceAngle */



