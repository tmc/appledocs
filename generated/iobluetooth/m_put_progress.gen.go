// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mPUTProgress */


/* debug [class_header]: Header for mPUTProgress */
// The class instance for the [mPUTProgress] class.
var (
	MPUTProgressClass     _mPUTProgressClass
	MPUTProgressClassOnce sync.Once
)

func getmPUTProgressClass() _mPUTProgressClass {
	MPUTProgressClassOnce.Do(func() {
		MPUTProgressClass = _mPUTProgressClass{objc.GetClass("mPUTProgress")}
	})
	return MPUTProgressClass
}

type _mPUTProgressClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mPUTProgress */
// An interface definition for the [mPUTProgress] class.
type ImPUTProgress interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mPUTProgress */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mPUTProgress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mPUTProgress */
// Alloc allocates a new instance without initialization.
func (mc _mPUTProgressClass) Alloc() mPUTProgress {
	rv := objc.Send[mPUTProgress](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mPUTProgressClass) New() mPUTProgress {
	rv := objc.Send[mPUTProgress](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPUTProgress) Init() mPUTProgress {
	rv := objc.Send[mPUTProgress](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPUTProgress) Autorelease() mPUTProgress {
	rv := objc.Send[mPUTProgress](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPUTProgress creates a new mPUTProgress instance.
func NewmPUTProgress() mPUTProgress {
	return getmPUTProgressClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mPUTProgress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mPUTProgress
type mPUTProgress struct {
	objectivec.Object
}

// mPUTProgressFrom constructs a [mPUTProgress] from an unsafe.Pointer.
func mPUTProgressFrom(ptr unsafe.Pointer) mPUTProgress {
	return mPUTProgress{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mPUTProgress *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mPUTProgress */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mPUTProgress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mPUTProgress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mPUTProgress */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mPUTProgress */



