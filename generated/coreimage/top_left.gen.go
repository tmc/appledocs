// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class topLeft */


/* debug [class_header]: Header for topLeft */
// The class instance for the [topLeft] class.
var (
	TopLeftClass     _topLeftClass
	TopLeftClassOnce sync.Once
)

func gettopLeftClass() _topLeftClass {
	TopLeftClassOnce.Do(func() {
		TopLeftClass = _topLeftClass{objc.GetClass("topLeft")}
	})
	return TopLeftClass
}

type _topLeftClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for topLeft */
// An interface definition for the [topLeft] class.
type ItopLeft interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for topLeft */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for topLeft */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for topLeft */
// Alloc allocates a new instance without initialization.
func (tc _topLeftClass) Alloc() topLeft {
	rv := objc.Send[topLeft](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _topLeftClass) New() topLeft {
	rv := objc.Send[topLeft](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ topLeft) Init() topLeft {
	rv := objc.Send[topLeft](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ topLeft) Autorelease() topLeft {
	rv := objc.Send[topLeft](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtopLeft creates a new topLeft instance.
func NewtopLeft() topLeft {
	return gettopLeftClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for topLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/topLeft-c.ivar
type topLeft struct {
	objectivec.Object
}

// topLeftFrom constructs a [topLeft] from an unsafe.Pointer.
func topLeftFrom(ptr unsafe.Pointer) topLeft {
	return topLeft{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for topLeft *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for topLeft */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for topLeft */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for topLeft */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for topLeft */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class topLeft */



