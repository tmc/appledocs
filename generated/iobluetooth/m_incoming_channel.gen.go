// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mIncomingChannel */


/* debug [class_header]: Header for mIncomingChannel */
// The class instance for the [mIncomingChannel] class.
var (
	MIncomingChannelClass     _mIncomingChannelClass
	MIncomingChannelClassOnce sync.Once
)

func getmIncomingChannelClass() _mIncomingChannelClass {
	MIncomingChannelClassOnce.Do(func() {
		MIncomingChannelClass = _mIncomingChannelClass{objc.GetClass("mIncomingChannel")}
	})
	return MIncomingChannelClass
}

type _mIncomingChannelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mIncomingChannel */
// An interface definition for the [mIncomingChannel] class.
type ImIncomingChannel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mIncomingChannel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mIncomingChannel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mIncomingChannel */
// Alloc allocates a new instance without initialization.
func (mc _mIncomingChannelClass) Alloc() mIncomingChannel {
	rv := objc.Send[mIncomingChannel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mIncomingChannelClass) New() mIncomingChannel {
	rv := objc.Send[mIncomingChannel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIncomingChannel) Init() mIncomingChannel {
	rv := objc.Send[mIncomingChannel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIncomingChannel) Autorelease() mIncomingChannel {
	rv := objc.Send[mIncomingChannel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIncomingChannel creates a new mIncomingChannel instance.
func NewmIncomingChannel() mIncomingChannel {
	return getmIncomingChannelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mIncomingChannel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mIncomingChannel
type mIncomingChannel struct {
	objectivec.Object
}

// mIncomingChannelFrom constructs a [mIncomingChannel] from an unsafe.Pointer.
func mIncomingChannelFrom(ptr unsafe.Pointer) mIncomingChannel {
	return mIncomingChannel{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mIncomingChannel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mIncomingChannel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mIncomingChannel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mIncomingChannel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mIncomingChannel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mIncomingChannel */



