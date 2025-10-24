// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class bounds */


/* debug [class_header]: Header for bounds */
// The class instance for the [bounds] class.
var (
	BoundsClass     _boundsClass
	BoundsClassOnce sync.Once
)

func getboundsClass() _boundsClass {
	BoundsClassOnce.Do(func() {
		BoundsClass = _boundsClass{objc.GetClass("bounds")}
	})
	return BoundsClass
}

type _boundsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for bounds */
// An interface definition for the [bounds] class.
type Ibounds interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for bounds */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for bounds */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for bounds */
// Alloc allocates a new instance without initialization.
func (bc _boundsClass) Alloc() bounds {
	rv := objc.Send[bounds](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _boundsClass) New() bounds {
	rv := objc.Send[bounds](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ bounds) Init() bounds {
	rv := objc.Send[bounds](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ bounds) Autorelease() bounds {
	rv := objc.Send[bounds](b_.ID, objc.Sel("autorelease"))
	return rv
}

// Newbounds creates a new bounds instance.
func Newbounds() bounds {
	return getboundsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for bounds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/bounds-c.ivar
type bounds struct {
	objectivec.Object
}

// boundsFrom constructs a [bounds] from an unsafe.Pointer.
func boundsFrom(ptr unsafe.Pointer) bounds {
	return bounds{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for bounds *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for bounds */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for bounds */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for bounds */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for bounds */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class bounds */



