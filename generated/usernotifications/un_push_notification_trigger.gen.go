// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UNPushNotificationTrigger] class.
var (
	UNPushNotificationTriggerClass     _UNPushNotificationTriggerClass
	UNPushNotificationTriggerClassOnce sync.Once
)

func getUNPushNotificationTriggerClass() _UNPushNotificationTriggerClass {
	UNPushNotificationTriggerClassOnce.Do(func() {
		UNPushNotificationTriggerClass = _UNPushNotificationTriggerClass{objc.GetClass("UNPushNotificationTrigger")}
	})
	return UNPushNotificationTriggerClass
}

type _UNPushNotificationTriggerClass struct {
	class objc.Class
}

// An interface definition for the [UNPushNotificationTrigger] class.
type IUNPushNotificationTrigger interface {
	IUNNotificationTrigger
}

// A trigger condition that indicates Apple Push Notification Service (APNs) has sent the notification.
//
// You don’t create instances of this class yourself. The system creates objects and associates them with requests that originated from Apple Push Notification service. You encounter instances of this class when managing your app’s delivered notification requests, which store an object of this type in their property.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNPushNotificationTrigger
type UNPushNotificationTrigger struct {
	UNNotificationTrigger
}

// UNPushNotificationTriggerFrom constructs a [UNPushNotificationTrigger] from an unsafe.Pointer.
//
// A trigger condition that indicates Apple Push Notification Service (APNs) has sent the notification.
func UNPushNotificationTriggerFrom(ptr unsafe.Pointer) UNPushNotificationTrigger {
	return UNPushNotificationTrigger{
		UNNotificationTrigger: UNNotificationTriggerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UNPushNotificationTriggerClass) Alloc() UNPushNotificationTrigger {
	rv := objc.Send[UNPushNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UNPushNotificationTriggerClass) New() UNPushNotificationTrigger {
	rv := objc.Send[UNPushNotificationTrigger](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNPushNotificationTrigger) Init() UNPushNotificationTrigger {
	rv := objc.Send[UNPushNotificationTrigger](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNPushNotificationTrigger) Autorelease() UNPushNotificationTrigger {
	rv := objc.Send[UNPushNotificationTrigger](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNPushNotificationTrigger creates a new UNPushNotificationTrigger instance.
func NewUNPushNotificationTrigger() UNPushNotificationTrigger {
	return getUNPushNotificationTriggerClass().New()
}


// The conditions that trigger the delivery of the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unnotificationrequest/trigger
func (u_ UNPushNotificationTrigger) Trigger() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("trigger"))
	return rv
}


// SetTrigger sets the value of the trigger property.
// The conditions that trigger the delivery of the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unnotificationrequest/trigger
func (u_ UNPushNotificationTrigger) SetTrigger(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTrigger:"), value)
}



