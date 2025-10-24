// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class UNPushNotificationTrigger */


/* debug [class_header]: Header for UNPushNotificationTrigger */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UNPushNotificationTrigger */
// An interface definition for the [UNPushNotificationTrigger] class.
type IUNPushNotificationTrigger interface {
	IUNNotificationTrigger
	
/* debug [class_interface_properties]: Properties for UNPushNotificationTrigger */
	// properties:
	Trigger() IUNNotificationTrigger
	SetTrigger(value IUNNotificationTrigger)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UNPushNotificationTrigger */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UNPushNotificationTrigger */
// Alloc allocates a new instance without initialization.
func (uc _UNPushNotificationTriggerClass) Alloc() UNPushNotificationTrigger {
	rv := objc.Send[UNPushNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UNPushNotificationTrigger */
// A trigger condition that indicates Apple Push Notification Service (APNs) has sent the notification.
//
// You don’t create instances of this class yourself. The system creates objects and associates them with requests that originated from Apple Push Notification service. You encounter instances of this class when managing your app’s delivered notification requests, which store an object of this type in their property.


// A trigger condition that indicates Apple Push Notification Service (APNs) has sent the notification.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UNPushNotificationTrigger *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UNPushNotificationTrigger */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UNPushNotificationTrigger */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UNPushNotificationTrigger */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UNPushNotificationTrigger */

// The conditions that trigger the delivery of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unnotificationrequest/trigger
func (u_ UNPushNotificationTrigger) Trigger() IUNNotificationTrigger {
	rv := objc.Send[UNNotificationTrigger](u_.ID, objc.Sel("trigger"))
	return rv
}/* debug [instance_properties/getter]: trigger */


// The conditions that trigger the delivery of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unnotificationrequest/trigger
func (u_ UNPushNotificationTrigger) SetTrigger(value IUNNotificationTrigger) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTrigger:"), value)
}/* debug [instance_properties/setter]: trigger */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UNPushNotificationTrigger */



