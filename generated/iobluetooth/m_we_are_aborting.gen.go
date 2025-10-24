// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mWeAreAborting */


/* debug [class_header]: Header for mWeAreAborting */
// The class instance for the [mWeAreAborting] class.
var (
	MWeAreAbortingClass     _mWeAreAbortingClass
	MWeAreAbortingClassOnce sync.Once
)

func getmWeAreAbortingClass() _mWeAreAbortingClass {
	MWeAreAbortingClassOnce.Do(func() {
		MWeAreAbortingClass = _mWeAreAbortingClass{objc.GetClass("mWeAreAborting")}
	})
	return MWeAreAbortingClass
}

type _mWeAreAbortingClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mWeAreAborting */
// An interface definition for the [mWeAreAborting] class.
type ImWeAreAborting interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mWeAreAborting */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mWeAreAborting */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mWeAreAborting */
// Alloc allocates a new instance without initialization.
func (mc _mWeAreAbortingClass) Alloc() mWeAreAborting {
	rv := objc.Send[mWeAreAborting](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mWeAreAbortingClass) New() mWeAreAborting {
	rv := objc.Send[mWeAreAborting](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mWeAreAborting) Init() mWeAreAborting {
	rv := objc.Send[mWeAreAborting](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mWeAreAborting) Autorelease() mWeAreAborting {
	rv := objc.Send[mWeAreAborting](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmWeAreAborting creates a new mWeAreAborting instance.
func NewmWeAreAborting() mWeAreAborting {
	return getmWeAreAbortingClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mWeAreAborting */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mWeAreAborting
type mWeAreAborting struct {
	objectivec.Object
}

// mWeAreAbortingFrom constructs a [mWeAreAborting] from an unsafe.Pointer.
func mWeAreAbortingFrom(ptr unsafe.Pointer) mWeAreAborting {
	return mWeAreAborting{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mWeAreAborting *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mWeAreAborting */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mWeAreAborting */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mWeAreAborting */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mWeAreAborting */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mWeAreAborting */



