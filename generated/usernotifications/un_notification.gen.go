// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UNNotification */

/* debug [class_header]: Header for UNNotification */
// The class instance for the [UNNotification] class.
var (
	UNNotificationClass     _UNNotificationClass
	UNNotificationClassOnce sync.Once
)

func getUNNotificationClass() _UNNotificationClass {
	UNNotificationClassOnce.Do(func() {
		UNNotificationClass = _UNNotificationClass{objc.GetClass("UNNotification")}
	})
	return UNNotificationClass
}

type _UNNotificationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for UNNotification */
// An interface definition for the [UNNotification] class.
type IUNNotification interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for UNNotification */
	// properties:
	Date() objc.IObject /* cross-framework: NSDate */
	Request() IUNNotificationRequest
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for UNNotification */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for UNNotification */
// Alloc allocates a new instance without initialization.
func (uc _UNNotificationClass) Alloc() UNNotification {
	rv := objc.Send[UNNotification](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNNotificationClass) New() UNNotification {
	rv := objc.Send[UNNotification](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotification) Init() UNNotification {
	rv := objc.Send[UNNotification](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotification) Autorelease() UNNotification {
	rv := objc.Send[UNNotification](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotification creates a new UNNotification instance.
func NewUNNotification() UNNotification {
	return getUNNotificationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for UNNotification */
// The data for a local or remote notification the system delivers to your app.
//
// A object contains the initial notification request, which contains the notification’s payload, and the date that the system delivered the notification. Don’t create notification objects directly. When handling notifications, the system delivers notification objects to your object. The object also maintains the list of notifications that the system delivers, and you use the method to retrieve those objects.

// The data for a local or remote notification the system delivers to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotification
type UNNotification struct {
	objectivec.Object
}

// UNNotificationFrom constructs a [UNNotification] from an unsafe.Pointer.
//
// The data for a local or remote notification the system delivers to your app.
func UNNotificationFrom(ptr unsafe.Pointer) UNNotification {
	return UNNotification{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for UNNotification */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for UNNotification */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for UNNotification */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for UNNotification */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for UNNotification */

// The delivery date of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotification/date
func (u_ UNNotification) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](u_.ID, objc.Sel("date"))
	return rv
} /* debug [instance_properties/getter]: date */

// The notification request containing the payload and trigger condition for the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotification/request
func (u_ UNNotification) Request() IUNNotificationRequest {
	rv := objc.Send[UNNotificationRequest](u_.ID, objc.Sel("request"))
	return rv
} /* debug [instance_properties/getter]: request */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class UNNotification */
