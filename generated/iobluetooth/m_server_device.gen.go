// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mServerDevice */


/* debug [class_header]: Header for mServerDevice */
// The class instance for the [mServerDevice] class.
var (
	MServerDeviceClass     _mServerDeviceClass
	MServerDeviceClassOnce sync.Once
)

func getmServerDeviceClass() _mServerDeviceClass {
	MServerDeviceClassOnce.Do(func() {
		MServerDeviceClass = _mServerDeviceClass{objc.GetClass("mServerDevice")}
	})
	return MServerDeviceClass
}

type _mServerDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mServerDevice */
// An interface definition for the [mServerDevice] class.
type ImServerDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mServerDevice */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mServerDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mServerDevice */
// Alloc allocates a new instance without initialization.
func (mc _mServerDeviceClass) Alloc() mServerDevice {
	rv := objc.Send[mServerDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mServerDeviceClass) New() mServerDevice {
	rv := objc.Send[mServerDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mServerDevice) Init() mServerDevice {
	rv := objc.Send[mServerDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mServerDevice) Autorelease() mServerDevice {
	rv := objc.Send[mServerDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmServerDevice creates a new mServerDevice instance.
func NewmServerDevice() mServerDevice {
	return getmServerDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mServerDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mServerDevice
type mServerDevice struct {
	objectivec.Object
}

// mServerDeviceFrom constructs a [mServerDevice] from an unsafe.Pointer.
func mServerDeviceFrom(ptr unsafe.Pointer) mServerDevice {
	return mServerDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mServerDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mServerDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mServerDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mServerDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mServerDevice */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mServerDevice */



