// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mChannelPSM */


/* debug [class_header]: Header for mChannelPSM */
// The class instance for the [mChannelPSM] class.
var (
	MChannelPSMClass     _mChannelPSMClass
	MChannelPSMClassOnce sync.Once
)

func getmChannelPSMClass() _mChannelPSMClass {
	MChannelPSMClassOnce.Do(func() {
		MChannelPSMClass = _mChannelPSMClass{objc.GetClass("mChannelPSM")}
	})
	return MChannelPSMClass
}

type _mChannelPSMClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mChannelPSM */
// An interface definition for the [mChannelPSM] class.
type ImChannelPSM interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mChannelPSM */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mChannelPSM */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mChannelPSM */
// Alloc allocates a new instance without initialization.
func (mc _mChannelPSMClass) Alloc() mChannelPSM {
	rv := objc.Send[mChannelPSM](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mChannelPSMClass) New() mChannelPSM {
	rv := objc.Send[mChannelPSM](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mChannelPSM) Init() mChannelPSM {
	rv := objc.Send[mChannelPSM](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mChannelPSM) Autorelease() mChannelPSM {
	rv := objc.Send[mChannelPSM](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmChannelPSM creates a new mChannelPSM instance.
func NewmChannelPSM() mChannelPSM {
	return getmChannelPSMClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mChannelPSM */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mChannelPSM
type mChannelPSM struct {
	objectivec.Object
}

// mChannelPSMFrom constructs a [mChannelPSM] from an unsafe.Pointer.
func mChannelPSMFrom(ptr unsafe.Pointer) mChannelPSM {
	return mChannelPSM{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mChannelPSM *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mChannelPSM */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mChannelPSM */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mChannelPSM */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mChannelPSM */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mChannelPSM */



