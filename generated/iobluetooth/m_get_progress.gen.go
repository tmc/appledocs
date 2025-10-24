// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mGETProgress */


/* debug [class_header]: Header for mGETProgress */
// The class instance for the [mGETProgress] class.
var (
	MGETProgressClass     _mGETProgressClass
	MGETProgressClassOnce sync.Once
)

func getmGETProgressClass() _mGETProgressClass {
	MGETProgressClassOnce.Do(func() {
		MGETProgressClass = _mGETProgressClass{objc.GetClass("mGETProgress")}
	})
	return MGETProgressClass
}

type _mGETProgressClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mGETProgress */
// An interface definition for the [mGETProgress] class.
type ImGETProgress interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mGETProgress */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mGETProgress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mGETProgress */
// Alloc allocates a new instance without initialization.
func (mc _mGETProgressClass) Alloc() mGETProgress {
	rv := objc.Send[mGETProgress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mGETProgressClass) New() mGETProgress {
	rv := objc.Send[mGETProgress](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mGETProgress) Init() mGETProgress {
	rv := objc.Send[mGETProgress](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mGETProgress) Autorelease() mGETProgress {
	rv := objc.Send[mGETProgress](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmGETProgress creates a new mGETProgress instance.
func NewmGETProgress() mGETProgress {
	return getmGETProgressClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mGETProgress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mGETProgress
type mGETProgress struct {
	objectivec.Object
}

// mGETProgressFrom constructs a [mGETProgress] from an unsafe.Pointer.
func mGETProgressFrom(ptr unsafe.Pointer) mGETProgress {
	return mGETProgress{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mGETProgress *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mGETProgress */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mGETProgress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mGETProgress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mGETProgress */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mGETProgress */



