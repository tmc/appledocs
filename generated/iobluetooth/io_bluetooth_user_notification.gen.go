// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [BluetoothUserNotification] class.
type IBluetoothUserNotification interface {
	objectivec.IObject
	// properties:
	// methods:
	Unregister()
}

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

// Alloc allocates a new instance without initialization.
func (bc _BluetoothUserNotificationClass) Alloc() BluetoothUserNotification {
	rv := objc.Send[BluetoothUserNotification](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Called to unregister the target notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothUserNotification/unregister()
func (b_ BluetoothUserNotification) Unregister() {
	objc.Send[objc.ID](b_.ID, objc.Sel("unregister"))
}



