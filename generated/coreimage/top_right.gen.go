// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class topRight */


/* debug [class_header]: Header for topRight */
// The class instance for the [topRight] class.
var (
	TopRightClass     _topRightClass
	TopRightClassOnce sync.Once
)

func gettopRightClass() _topRightClass {
	TopRightClassOnce.Do(func() {
		TopRightClass = _topRightClass{objc.GetClass("topRight")}
	})
	return TopRightClass
}

type _topRightClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for topRight */
// An interface definition for the [topRight] class.
type ItopRight interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for topRight */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for topRight */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for topRight */
// Alloc allocates a new instance without initialization.
func (tc _topRightClass) Alloc() topRight {
	rv := objc.Send[topRight](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _topRightClass) New() topRight {
	rv := objc.Send[topRight](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ topRight) Init() topRight {
	rv := objc.Send[topRight](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ topRight) Autorelease() topRight {
	rv := objc.Send[topRight](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtopRight creates a new topRight instance.
func NewtopRight() topRight {
	return gettopRightClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for topRight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/topRight-c.ivar
type topRight struct {
	objectivec.Object
}

// topRightFrom constructs a [topRight] from an unsafe.Pointer.
func topRightFrom(ptr unsafe.Pointer) topRight {
	return topRight{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for topRight *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for topRight */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for topRight */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for topRight */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for topRight */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class topRight */



