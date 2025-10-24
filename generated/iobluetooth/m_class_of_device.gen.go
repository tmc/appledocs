// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mClassOfDevice */


/* debug [class_header]: Header for mClassOfDevice */
// The class instance for the [mClassOfDevice] class.
var (
	MClassOfDeviceClass     _mClassOfDeviceClass
	MClassOfDeviceClassOnce sync.Once
)

func getmClassOfDeviceClass() _mClassOfDeviceClass {
	MClassOfDeviceClassOnce.Do(func() {
		MClassOfDeviceClass = _mClassOfDeviceClass{objc.GetClass("mClassOfDevice")}
	})
	return MClassOfDeviceClass
}

type _mClassOfDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mClassOfDevice */
// An interface definition for the [mClassOfDevice] class.
type ImClassOfDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mClassOfDevice */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mClassOfDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mClassOfDevice */
// Alloc allocates a new instance without initialization.
func (mc _mClassOfDeviceClass) Alloc() mClassOfDevice {
	rv := objc.Send[mClassOfDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mClassOfDeviceClass) New() mClassOfDevice {
	rv := objc.Send[mClassOfDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mClassOfDevice) Init() mClassOfDevice {
	rv := objc.Send[mClassOfDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mClassOfDevice) Autorelease() mClassOfDevice {
	rv := objc.Send[mClassOfDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmClassOfDevice creates a new mClassOfDevice instance.
func NewmClassOfDevice() mClassOfDevice {
	return getmClassOfDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mClassOfDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mClassOfDevice
type mClassOfDevice struct {
	objectivec.Object
}

// mClassOfDeviceFrom constructs a [mClassOfDevice] from an unsafe.Pointer.
func mClassOfDeviceFrom(ptr unsafe.Pointer) mClassOfDevice {
	return mClassOfDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mClassOfDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mClassOfDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mClassOfDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mClassOfDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mClassOfDevice */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mClassOfDevice */



