// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [UNNotification] class.
type IUNNotification interface {
	objectivec.IObject
	// properties:
	Date() objc.IObject /* cross-framework: NSDate */
	Request() IUNNotificationRequest
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (uc _UNNotificationClass) Alloc() UNNotification {
	rv := objc.Send[UNNotification](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The delivery date of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotification/date
func (u_ UNNotification) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](u_.ID, objc.Sel("date"))
	return rv
}


// The notification request containing the payload and trigger condition for the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotification/request
func (u_ UNNotification) Request() IUNNotificationRequest {
	rv := objc.Send[UNNotificationRequest](u_.ID, objc.Sel("request"))
	return rv
}



