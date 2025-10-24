// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mIncomingDataQueue */


/* debug [class_header]: Header for mIncomingDataQueue */
// The class instance for the [mIncomingDataQueue] class.
var (
	MIncomingDataQueueClass     _mIncomingDataQueueClass
	MIncomingDataQueueClassOnce sync.Once
)

func getmIncomingDataQueueClass() _mIncomingDataQueueClass {
	MIncomingDataQueueClassOnce.Do(func() {
		MIncomingDataQueueClass = _mIncomingDataQueueClass{objc.GetClass("mIncomingDataQueue")}
	})
	return MIncomingDataQueueClass
}

type _mIncomingDataQueueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mIncomingDataQueue */
// An interface definition for the [mIncomingDataQueue] class.
type ImIncomingDataQueue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mIncomingDataQueue */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mIncomingDataQueue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mIncomingDataQueue */
// Alloc allocates a new instance without initialization.
func (mc _mIncomingDataQueueClass) Alloc() mIncomingDataQueue {
	rv := objc.Send[mIncomingDataQueue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mIncomingDataQueueClass) New() mIncomingDataQueue {
	rv := objc.Send[mIncomingDataQueue](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIncomingDataQueue) Init() mIncomingDataQueue {
	rv := objc.Send[mIncomingDataQueue](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIncomingDataQueue) Autorelease() mIncomingDataQueue {
	rv := objc.Send[mIncomingDataQueue](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIncomingDataQueue creates a new mIncomingDataQueue instance.
func NewmIncomingDataQueue() mIncomingDataQueue {
	return getmIncomingDataQueueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mIncomingDataQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mIncomingDataQueue
type mIncomingDataQueue struct {
	objectivec.Object
}

// mIncomingDataQueueFrom constructs a [mIncomingDataQueue] from an unsafe.Pointer.
func mIncomingDataQueueFrom(ptr unsafe.Pointer) mIncomingDataQueue {
	return mIncomingDataQueue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mIncomingDataQueue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mIncomingDataQueue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mIncomingDataQueue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mIncomingDataQueue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mIncomingDataQueue */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mIncomingDataQueue */



