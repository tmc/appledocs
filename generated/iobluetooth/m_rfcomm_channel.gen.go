// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mRFCOMMChannel */


/* debug [class_header]: Header for mRFCOMMChannel */
// The class instance for the [mRFCOMMChannel] class.
var (
	MRFCOMMChannelClass     _mRFCOMMChannelClass
	MRFCOMMChannelClassOnce sync.Once
)

func getmRFCOMMChannelClass() _mRFCOMMChannelClass {
	MRFCOMMChannelClassOnce.Do(func() {
		MRFCOMMChannelClass = _mRFCOMMChannelClass{objc.GetClass("mRFCOMMChannel")}
	})
	return MRFCOMMChannelClass
}

type _mRFCOMMChannelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mRFCOMMChannel */
// An interface definition for the [mRFCOMMChannel] class.
type ImRFCOMMChannel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mRFCOMMChannel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mRFCOMMChannel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mRFCOMMChannel */
// Alloc allocates a new instance without initialization.
func (mc _mRFCOMMChannelClass) Alloc() mRFCOMMChannel {
	rv := objc.Send[mRFCOMMChannel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mRFCOMMChannelClass) New() mRFCOMMChannel {
	rv := objc.Send[mRFCOMMChannel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mRFCOMMChannel) Init() mRFCOMMChannel {
	rv := objc.Send[mRFCOMMChannel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mRFCOMMChannel) Autorelease() mRFCOMMChannel {
	rv := objc.Send[mRFCOMMChannel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmRFCOMMChannel creates a new mRFCOMMChannel instance.
func NewmRFCOMMChannel() mRFCOMMChannel {
	return getmRFCOMMChannelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mRFCOMMChannel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/mRFCOMMChannel
type mRFCOMMChannel struct {
	objectivec.Object
}

// mRFCOMMChannelFrom constructs a [mRFCOMMChannel] from an unsafe.Pointer.
func mRFCOMMChannelFrom(ptr unsafe.Pointer) mRFCOMMChannel {
	return mRFCOMMChannel{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mRFCOMMChannel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mRFCOMMChannel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mRFCOMMChannel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mRFCOMMChannel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mRFCOMMChannel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mRFCOMMChannel */



