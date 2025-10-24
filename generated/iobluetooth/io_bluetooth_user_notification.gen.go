// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothUserNotification */


/* debug [class_header]: Header for IOBluetoothUserNotification */
// The class instance for the [BluetoothUserNotification] class.
var (
	BluetoothUserNotificationClass     _BluetoothUserNotificationClass
	BluetoothUserNotificationClassOnce sync.Once
)

func getBluetoothUserNotificationClass() _BluetoothUserNotificationClass {
	BluetoothUserNotificationClassOnce.Do(func() {
		BluetoothUserNotificationClass = _BluetoothUserNotificationClass{objc.GetClass("IOBluetoothUserNotification")}
	})
	return BluetoothUserNotificationClass
}

type _BluetoothUserNotificationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothUserNotification */
// An interface definition for the [BluetoothUserNotification] class.
type IBluetoothUserNotification interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BluetoothUserNotification */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothUserNotification */
	// methods:
	Unregister()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothUserNotification */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothUserNotificationClass) Alloc() BluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BluetoothUserNotificationClass) New() BluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothUserNotification) Init() BluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothUserNotification) Autorelease() BluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothUserNotification creates a new BluetoothUserNotification instance.
func NewBluetoothUserNotification() BluetoothUserNotification {
	return getBluetoothUserNotificationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothUserNotification */
// Represents a registered notification.
//
// When registering for various notifications in the system, an IOBluetoothUserNotification object is returned. To unregister from the notification, call -unregister on the IOBluetoothUserNotification object. Once -unregister is called, the object will no longer be valid.


// Represents a registered notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUserNotification
type BluetoothUserNotification struct {
	objectivec.Object
}

// BluetoothUserNotificationFrom constructs a [BluetoothUserNotification] from an unsafe.Pointer.
//
// Represents a registered notification.
func BluetoothUserNotificationFrom(ptr unsafe.Pointer) BluetoothUserNotification {
	return BluetoothUserNotification{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothUserNotification *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothUserNotification */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothUserNotification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothUserNotification */

// Called to unregister the target notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUserNotification/unregister()
func (b_ BluetoothUserNotification) Unregister() {
	objc.Send[objc.ID](b_.ID, objc.Sel("unregister"))
}/* debug [instance_methods/method]: Unregister */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothUserNotification */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothUserNotification */



