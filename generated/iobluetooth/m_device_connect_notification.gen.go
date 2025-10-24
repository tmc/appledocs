// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mDeviceConnectNotification */


/* debug [class_header]: Header for mDeviceConnectNotification */
// The class instance for the [mDeviceConnectNotification] class.
var (
	MDeviceConnectNotificationClass     _mDeviceConnectNotificationClass
	MDeviceConnectNotificationClassOnce sync.Once
)

func getmDeviceConnectNotificationClass() _mDeviceConnectNotificationClass {
	MDeviceConnectNotificationClassOnce.Do(func() {
		MDeviceConnectNotificationClass = _mDeviceConnectNotificationClass{objc.GetClass("mDeviceConnectNotification")}
	})
	return MDeviceConnectNotificationClass
}

type _mDeviceConnectNotificationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mDeviceConnectNotification */
// An interface definition for the [mDeviceConnectNotification] class.
type ImDeviceConnectNotification interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mDeviceConnectNotification */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mDeviceConnectNotification */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mDeviceConnectNotification */
// Alloc allocates a new instance without initialization.
func (mc _mDeviceConnectNotificationClass) Alloc() mDeviceConnectNotification {
	rv := objc.Send[mDeviceConnectNotification](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mDeviceConnectNotificationClass) New() mDeviceConnectNotification {
	rv := objc.Send[mDeviceConnectNotification](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mDeviceConnectNotification) Init() mDeviceConnectNotification {
	rv := objc.Send[mDeviceConnectNotification](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mDeviceConnectNotification) Autorelease() mDeviceConnectNotification {
	rv := objc.Send[mDeviceConnectNotification](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmDeviceConnectNotification creates a new mDeviceConnectNotification instance.
func NewmDeviceConnectNotification() mDeviceConnectNotification {
	return getmDeviceConnectNotificationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mDeviceConnectNotification */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mDeviceConnectNotification
type mDeviceConnectNotification struct {
	objectivec.Object
}

// mDeviceConnectNotificationFrom constructs a [mDeviceConnectNotification] from an unsafe.Pointer.
func mDeviceConnectNotificationFrom(ptr unsafe.Pointer) mDeviceConnectNotification {
	return mDeviceConnectNotification{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mDeviceConnectNotification *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mDeviceConnectNotification */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mDeviceConnectNotification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mDeviceConnectNotification */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mDeviceConnectNotification */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mDeviceConnectNotification */



