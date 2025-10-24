// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mDevice */


/* debug [class_header]: Header for mDevice */
// The class instance for the [mDevice] class.
var (
	MDeviceClass     _mDeviceClass
	MDeviceClassOnce sync.Once
)

func getmDeviceClass() _mDeviceClass {
	MDeviceClassOnce.Do(func() {
		MDeviceClass = _mDeviceClass{objc.GetClass("mDevice")}
	})
	return MDeviceClass
}

type _mDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mDevice */
// An interface definition for the [mDevice] class.
type ImDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mDevice */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mDevice */
// Alloc allocates a new instance without initialization.
func (mc _mDeviceClass) Alloc() mDevice {
	rv := objc.Send[mDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mDeviceClass) New() mDevice {
	rv := objc.Send[mDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mDevice) Init() mDevice {
	rv := objc.Send[mDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mDevice) Autorelease() mDevice {
	rv := objc.Send[mDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmDevice creates a new mDevice instance.
func NewmDevice() mDevice {
	return getmDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mDevice
type mDevice struct {
	objectivec.Object
}

// mDeviceFrom constructs a [mDevice] from an unsafe.Pointer.
func mDeviceFrom(ptr unsafe.Pointer) mDevice {
	return mDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mDevice */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mDevice */



