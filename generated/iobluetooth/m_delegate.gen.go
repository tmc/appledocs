// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mDelegate */


/* debug [class_header]: Header for mDelegate */
// The class instance for the [mDelegate] class.
var (
	MDelegateClass     _mDelegateClass
	MDelegateClassOnce sync.Once
)

func getmDelegateClass() _mDelegateClass {
	MDelegateClassOnce.Do(func() {
		MDelegateClass = _mDelegateClass{objc.GetClass("mDelegate")}
	})
	return MDelegateClass
}

type _mDelegateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mDelegate */
// An interface definition for the [mDelegate] class.
type ImDelegate interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mDelegate */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mDelegate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mDelegate */
// Alloc allocates a new instance without initialization.
func (mc _mDelegateClass) Alloc() mDelegate {
	rv := objc.Send[mDelegate](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mDelegateClass) New() mDelegate {
	rv := objc.Send[mDelegate](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mDelegate) Init() mDelegate {
	rv := objc.Send[mDelegate](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mDelegate) Autorelease() mDelegate {
	rv := objc.Send[mDelegate](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmDelegate creates a new mDelegate instance.
func NewmDelegate() mDelegate {
	return getmDelegateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mDelegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mDelegate
type mDelegate struct {
	objectivec.Object
}

// mDelegateFrom constructs a [mDelegate] from an unsafe.Pointer.
func mDelegateFrom(ptr unsafe.Pointer) mDelegate {
	return mDelegate{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mDelegate *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mDelegate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mDelegate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mDelegate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mDelegate */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mDelegate */



