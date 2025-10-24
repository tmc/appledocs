// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mL2Channel */


/* debug [class_header]: Header for mL2Channel */
// The class instance for the [mL2Channel] class.
var (
	ML2ChannelClass     _mL2ChannelClass
	ML2ChannelClassOnce sync.Once
)

func getmL2ChannelClass() _mL2ChannelClass {
	ML2ChannelClassOnce.Do(func() {
		ML2ChannelClass = _mL2ChannelClass{objc.GetClass("mL2Channel")}
	})
	return ML2ChannelClass
}

type _mL2ChannelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mL2Channel */
// An interface definition for the [mL2Channel] class.
type ImL2Channel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mL2Channel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mL2Channel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mL2Channel */
// Alloc allocates a new instance without initialization.
func (mc _mL2ChannelClass) Alloc() mL2Channel {
	rv := objc.Send[mL2Channel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mL2ChannelClass) New() mL2Channel {
	rv := objc.Send[mL2Channel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mL2Channel) Init() mL2Channel {
	rv := objc.Send[mL2Channel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mL2Channel) Autorelease() mL2Channel {
	rv := objc.Send[mL2Channel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmL2Channel creates a new mL2Channel instance.
func NewmL2Channel() mL2Channel {
	return getmL2ChannelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mL2Channel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/mL2Channel
type mL2Channel struct {
	objectivec.Object
}

// mL2ChannelFrom constructs a [mL2Channel] from an unsafe.Pointer.
func mL2ChannelFrom(ptr unsafe.Pointer) mL2Channel {
	return mL2Channel{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mL2Channel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mL2Channel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mL2Channel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mL2Channel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mL2Channel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mL2Channel */



