// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UNNotificationRequest] class.
var (
	UNNotificationRequestClass     _UNNotificationRequestClass
	UNNotificationRequestClassOnce sync.Once
)

func getUNNotificationRequestClass() _UNNotificationRequestClass {
	UNNotificationRequestClassOnce.Do(func() {
		UNNotificationRequestClass = _UNNotificationRequestClass{objc.GetClass("UNNotificationRequest")}
	})
	return UNNotificationRequestClass
}

type _UNNotificationRequestClass struct {
	class objc.Class
}

// An interface definition for the [UNNotificationRequest] class.
type IUNNotificationRequest interface {
	objectivec.IObject
}

// A request to schedule a local notification, which includes the content of the notification and the trigger conditions for delivery.
//
// Create a object when you want to schedule the delivery of a local notification. A notification request object contains a object with the payload and the object with the conditions that trigger the delivery of the notification. To schedule the delivery of your notification, pass your request object to the method of the shared user notification center object. After scheduling a request, you interact with objects in the following ways: View your app’s pending notifications by calling the method of your shared user notification center object. When the system delivers a notification to your app, the provided object contains a object that you can inspect to get the notification details. Use the request’s to remove delivered notifications from Notification Center. When receiving a local or remote notification, use the provided object to fetch details about the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest
type UNNotificationRequest struct {
	objectivec.Object
}

// UNNotificationRequestFrom constructs a [UNNotificationRequest] from an unsafe.Pointer.
//
// A request to schedule a local notification, which includes the content of the notification and the trigger conditions for delivery.
func UNNotificationRequestFrom(ptr unsafe.Pointer) UNNotificationRequest {
	return UNNotificationRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UNNotificationRequestClass) Alloc() UNNotificationRequest {
	rv := objc.Send[UNNotificationRequest](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UNNotificationRequestClass) New() UNNotificationRequest {
	rv := objc.Send[UNNotificationRequest](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationRequest) Init() UNNotificationRequest {
	rv := objc.Send[UNNotificationRequest](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationRequest) Autorelease() UNNotificationRequest {
	rv := objc.Send[UNNotificationRequest](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationRequest creates a new UNNotificationRequest instance.
func NewUNNotificationRequest() UNNotificationRequest {
	return getUNNotificationRequestClass().New()
}




// Creates a notification request object that you use to schedule a notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest/init(identifier:content:trigger:)
func NewUNNotificationRequestWithIdentifierContentTrigger(identifier appkit.string, content IUNNotificationContent, trigger IUNNotificationTrigger) UNNotificationRequest {
	rv := objc.Send[UNNotificationRequest](objc.ID(getUNNotificationRequestClass().class), objc.Sel("requestWithIdentifier:content:trigger:"), identifier, content, trigger)
	return rv
}


// Creates a notification request object that you use to schedule a notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest/init(identifier:content:trigger:)
func (uc _UNNotificationRequestClass) RequestWithIdentifierContentTrigger(identifier appkit.string, content IUNNotificationContent, trigger IUNNotificationTrigger) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("requestWithIdentifier:content:trigger:"), identifier, content, trigger)
	return rv
}

// The content associated with the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest/content
func (u_ UNNotificationRequest) Content() UNNotificationContent {
	rv := objc.Send[UNNotificationContent](u_.ID, objc.Sel("content"))
	return rv
}

// The unique identifier for this notification request.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest/identifier
func (u_ UNNotificationRequest) Identifier() appkit.string {
	rv := objc.Send[appkit.string](u_.ID, objc.Sel("identifier"))
	return rv
}

// The conditions that trigger the delivery of the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest/trigger
func (u_ UNNotificationRequest) Trigger() UNNotificationTrigger {
	rv := objc.Send[UNNotificationTrigger](u_.ID, objc.Sel("trigger"))
	return rv
}


