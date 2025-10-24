// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class vec */


/* debug [class_header]: Header for vec */
// The class instance for the [vec] class.
var (
	VecClass     _vecClass
	VecClassOnce sync.Once
)

func getvecClass() _vecClass {
	VecClassOnce.Do(func() {
		VecClass = _vecClass{objc.GetClass("vec")}
	})
	return VecClass
}

type _vecClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for vec */
// An interface definition for the [vec] class.
type Ivec interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for vec */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for vec */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for vec */
// Alloc allocates a new instance without initialization.
func (vc _vecClass) Alloc() vec {
	rv := objc.Send[vec](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _vecClass) New() vec {
	rv := objc.Send[vec](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ vec) Init() vec {
	rv := objc.Send[vec](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ vec) Autorelease() vec {
	rv := objc.Send[vec](v_.ID, objc.Sel("autorelease"))
	return rv
}

// Newvec creates a new vec instance.
func Newvec() vec {
	return getvecClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for vec */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/union_(unnamed)/vec
type vec struct {
	objectivec.Object
}

// vecFrom constructs a [vec] from an unsafe.Pointer.
func vecFrom(ptr unsafe.Pointer) vec {
	return vec{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for vec *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for vec */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for vec */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for vec */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for vec */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class vec */



