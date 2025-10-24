// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UNUserNotificationCenter */

/* debug [class_header]: Header for UNUserNotificationCenter */
// The class instance for the [UNUserNotificationCenter] class.
var (
	UNUserNotificationCenterClass     _UNUserNotificationCenterClass
	UNUserNotificationCenterClassOnce sync.Once
)

func getUNUserNotificationCenterClass() _UNUserNotificationCenterClass {
	UNUserNotificationCenterClassOnce.Do(func() {
		UNUserNotificationCenterClass = _UNUserNotificationCenterClass{objc.GetClass("UNUserNotificationCenter")}
	})
	return UNUserNotificationCenterClass
}

type _UNUserNotificationCenterClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for UNUserNotificationCenter */
// An interface definition for the [UNUserNotificationCenter] class.
type IUNUserNotificationCenter interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for UNUserNotificationCenter */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	SupportsContentExtensions() bool
	UNErrorDomain() objc.IObject /* cross-framework: NSString */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for UNUserNotificationCenter */
	// methods:
	AddNotificationRequestWithCompletionHandler(request IUNNotificationRequest, completionHandler unsafe.Pointer)
	GetDeliveredNotificationsWithCompletionHandler(completionHandler unsafe.Pointer)
	GetNotificationCategoriesWithCompletionHandler(completionHandler unsafe.Pointer)
	GetNotificationSettingsWithCompletionHandler(completionHandler unsafe.Pointer)
	GetPendingNotificationRequestsWithCompletionHandler(completionHandler unsafe.Pointer)
	RemoveAllDeliveredNotifications()
	RemoveAllPendingNotificationRequests()
	RemoveDeliveredNotificationsWithIdentifiers(identifiers []string)
	RemovePendingNotificationRequestsWithIdentifiers(identifiers []string)
	RequestAuthorizationWithOptionsCompletionHandler(options UNAuthorizationOptions, completionHandler unsafe.Pointer)
	SetBadgeCountWithCompletionHandler(newBadgeCount int, completionHandler unsafe.Pointer)
	SetNotificationCategories(categories unsafe.Pointer)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for UNUserNotificationCenter */
// Alloc allocates a new instance without initialization.
func (uc _UNUserNotificationCenterClass) Alloc() UNUserNotificationCenter {
	rv := objc.Send[UNUserNotificationCenter](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNUserNotificationCenterClass) New() UNUserNotificationCenter {
	rv := objc.Send[UNUserNotificationCenter](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNUserNotificationCenter) Init() UNUserNotificationCenter {
	rv := objc.Send[UNUserNotificationCenter](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNUserNotificationCenter) Autorelease() UNUserNotificationCenter {
	rv := objc.Send[UNUserNotificationCenter](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNUserNotificationCenter creates a new UNUserNotificationCenter instance.
func NewUNUserNotificationCenter() UNUserNotificationCenter {
	return getUNUserNotificationCenterClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for UNUserNotificationCenter */
// The central object for managing notification-related activities for your app or app extension.
//
// Use the shared object to manage all notification-related behaviors in your app or app extension. Specifically, use this object to do the following: Request authorization to interact with the user through alerts, sounds, and icon badges. See . Declare the notification types that your app supports and the custom actions the user may perform when the system delivers those notifications. See . Schedule the delivery of notifications from your app. See . Process the payloads from remote notifications the system delivers by Apple Push Notification service (APNs). See . Manage the already delivered notifications the system displays in Notification Center. See Managing Delivered Notifications. Handle user-selected actions associated with your custom notification types. See . Get the notification-related settings for your app. See Managing Settings and Authorization. To handle incoming notifications and notification-related actions, create an object that adopts the protocol and assign it to the property. Always assign an object to the property before performing any tasks that might interact with that delegate. You may use the shared user notification center object simultaneously from any of your app’s threads. The object processes requests serially in the order that the system initiates them.

// The central object for managing notification-related activities for your app or app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter
type UNUserNotificationCenter struct {
	objectivec.Object
}

// UNUserNotificationCenterFrom constructs a [UNUserNotificationCenter] from an unsafe.Pointer.
//
// The central object for managing notification-related activities for your app or app extension.
func UNUserNotificationCenterFrom(ptr unsafe.Pointer) UNUserNotificationCenter {
	return UNUserNotificationCenter{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for UNUserNotificationCenter */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for UNUserNotificationCenter */

// Returns your app’s notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/current()
func (uc _UNUserNotificationCenterClass) CurrentNotificationCenter() UNUserNotificationCenter {
	rv := objc.Send[UNUserNotificationCenter](objc.ID(uc.class), objc.Sel("currentNotificationCenter"))
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentNotificationCenter) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for UNUserNotificationCenter */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for UNUserNotificationCenter */

// Schedules the delivery of a local notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/add(_:withCompletionHandler:)
func (u_ UNUserNotificationCenter) AddNotificationRequestWithCompletionHandler(request IUNNotificationRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("addNotificationRequest:withCompletionHandler:"), request, completionHandler)
} /* debug [instance_methods/method]: AddNotificationRequestWithCompletionHandler */

// Fetches all of your app’s delivered notifications that are still present in Notification Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/getDeliveredNotifications(completionHandler:)
func (u_ UNUserNotificationCenter) GetDeliveredNotificationsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getDeliveredNotificationsWithCompletionHandler:"), completionHandler)
} /* debug [instance_methods/method]: GetDeliveredNotificationsWithCompletionHandler */

// Fetches your app’s registered notification categories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/getNotificationCategories(completionHandler:)
func (u_ UNUserNotificationCenter) GetNotificationCategoriesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getNotificationCategoriesWithCompletionHandler:"), completionHandler)
} /* debug [instance_methods/method]: GetNotificationCategoriesWithCompletionHandler */

// Retrieves the authorization and feature-related settings for your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/getNotificationSettings(completionHandler:)
func (u_ UNUserNotificationCenter) GetNotificationSettingsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getNotificationSettingsWithCompletionHandler:"), completionHandler)
} /* debug [instance_methods/method]: GetNotificationSettingsWithCompletionHandler */

// Fetches all of your app’s local notifications that are pending delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/getPendingNotificationRequests(completionHandler:)
func (u_ UNUserNotificationCenter) GetPendingNotificationRequestsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getPendingNotificationRequestsWithCompletionHandler:"), completionHandler)
} /* debug [instance_methods/method]: GetPendingNotificationRequestsWithCompletionHandler */

// Removes all of your app’s delivered notifications from Notification Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removeAllDeliveredNotifications()
func (u_ UNUserNotificationCenter) RemoveAllDeliveredNotifications() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllDeliveredNotifications"))
} /* debug [instance_methods/method]: RemoveAllDeliveredNotifications */

// Removes all of your app’s pending local notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removeAllPendingNotificationRequests()
func (u_ UNUserNotificationCenter) RemoveAllPendingNotificationRequests() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllPendingNotificationRequests"))
} /* debug [instance_methods/method]: RemoveAllPendingNotificationRequests */

// Removes your app’s notifications from Notification Center that match the specified identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removeDeliveredNotifications(withIdentifiers:)
func (u_ UNUserNotificationCenter) RemoveDeliveredNotificationsWithIdentifiers(identifiers []string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeDeliveredNotificationsWithIdentifiers:"), identifiers)
} /* debug [instance_methods/method]: RemoveDeliveredNotificationsWithIdentifiers */

// Removes your app’s local notifications that are pending and match the specified identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removePendingNotificationRequests(withIdentifiers:)
func (u_ UNUserNotificationCenter) RemovePendingNotificationRequestsWithIdentifiers(identifiers []string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removePendingNotificationRequestsWithIdentifiers:"), identifiers)
} /* debug [instance_methods/method]: RemovePendingNotificationRequestsWithIdentifiers */

// Requests a person’s authorization to allow local and remote notifications for your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/requestAuthorization(options:completionHandler:)
func (u_ UNUserNotificationCenter) RequestAuthorizationWithOptionsCompletionHandler(options UNAuthorizationOptions, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("requestAuthorizationWithOptions:completionHandler:"), options, completionHandler)
} /* debug [instance_methods/method]: RequestAuthorizationWithOptionsCompletionHandler */

// Updates the badge count for your app’s icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/setBadgeCount(_:withCompletionHandler:)
func (u_ UNUserNotificationCenter) SetBadgeCountWithCompletionHandler(newBadgeCount int, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setBadgeCount:withCompletionHandler:"), newBadgeCount, completionHandler)
} /* debug [instance_methods/method]: SetBadgeCountWithCompletionHandler */

// Registers the notification categories that your app supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/setNotificationCategories(_:)
func (u_ UNUserNotificationCenter) SetNotificationCategories(categories unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNotificationCategories:"), categories)
} /* debug [instance_methods/method]: SetNotificationCategories */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for UNUserNotificationCenter */

// The notification center’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/delegate
func (u_ UNUserNotificationCenter) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("delegate"))
	return rv
} /* debug [instance_properties/getter]: delegate */

// The notification center’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/delegate
func (u_ UNUserNotificationCenter) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDelegate:"), value)
} /* debug [instance_properties/setter]: delegate */

// A Boolean value that indicates whether the device supports notification content extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/supportsContentExtensions
func (u_ UNUserNotificationCenter) SupportsContentExtensions() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("supportsContentExtensions"))
	return rv
} /* debug [instance_properties/getter]: supportsContentExtensions */

// The error domain for notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unerrordomain
func (u_ UNUserNotificationCenter) UNErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("UNErrorDomain"))
	return rv
} /* debug [instance_properties/getter]: UNErrorDomain */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class UNUserNotificationCenter */
