// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [UNUserNotificationCenter] class.
type IUNUserNotificationCenter interface {
	objectivec.IObject
	AddNotificationRequestWithCompletionHandler(request unsafe.Pointer, completionHandler unsafe.Pointer)
	GetDeliveredNotificationsWithCompletionHandler(completionHandler unsafe.Pointer)
	GetNotificationCategoriesWithCompletionHandler(completionHandler unsafe.Pointer)
	GetNotificationSettingsWithCompletionHandler(completionHandler unsafe.Pointer)
	GetPendingNotificationRequestsWithCompletionHandler(completionHandler unsafe.Pointer)
	RemoveAllDeliveredNotifications()
	RemoveAllPendingNotificationRequests()
	RemoveDeliveredNotificationsWithIdentifiers(identifiers unsafe.Pointer)
	RemovePendingNotificationRequestsWithIdentifiers(identifiers unsafe.Pointer)
	RequestAuthorizationWithOptionsCompletionHandler(options unsafe.Pointer, completionHandler unsafe.Pointer)
	SetBadgeCountWithCompletionHandler(newBadgeCount int, completionHandler unsafe.Pointer)
	SetNotificationCategories(categories unsafe.Pointer)
}

// The central object for managing notification-related activities for your app or app extension.
//
// Use the shared object to manage all notification-related behaviors in your app or app extension. Specifically, use this object to do the following: Request authorization to interact with the user through alerts, sounds, and icon badges. See . Declare the notification types that your app supports and the custom actions the user may perform when the system delivers those notifications. See . Schedule the delivery of notifications from your app. See . Process the payloads from remote notifications the system delivers by Apple Push Notification service (APNs). See . Manage the already delivered notifications the system displays in Notification Center. See Managing Delivered Notifications. Handle user-selected actions associated with your custom notification types. See . Get the notification-related settings for your app. See Managing Settings and Authorization. To handle incoming notifications and notification-related actions, create an object that adopts the protocol and assign it to the property. Always assign an object to the property before performing any tasks that might interact with that delegate. You may use the shared user notification center object simultaneously from any of your app’s threads. The object processes requests serially in the order that the system initiates them.
//
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

// Alloc allocates a new instance without initialization.
func (uc _UNUserNotificationCenterClass) Alloc() UNUserNotificationCenter {
	rv := objc.Send[UNUserNotificationCenter](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns your app’s notification center.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/current()
func (uc _UNUserNotificationCenterClass) CurrentNotificationCenter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("currentNotificationCenter"))
	return rv
}

// Schedules the delivery of a local notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/add(_:withCompletionHandler:)
func (u_ UNUserNotificationCenter) AddNotificationRequestWithCompletionHandler(request unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("addNotificationRequest:withCompletionHandler:"), request, completionHandler)
}

// Fetches all of your app’s delivered notifications that are still present in Notification Center.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/getDeliveredNotifications(completionHandler:)
func (u_ UNUserNotificationCenter) GetDeliveredNotificationsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getDeliveredNotificationsWithCompletionHandler:"), completionHandler)
}

// Fetches your app’s registered notification categories.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/getNotificationCategories(completionHandler:)
func (u_ UNUserNotificationCenter) GetNotificationCategoriesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getNotificationCategoriesWithCompletionHandler:"), completionHandler)
}

// Retrieves the authorization and feature-related settings for your app.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/getNotificationSettings(completionHandler:)
func (u_ UNUserNotificationCenter) GetNotificationSettingsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getNotificationSettingsWithCompletionHandler:"), completionHandler)
}

// Fetches all of your app’s local notifications that are pending delivery.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/getPendingNotificationRequests(completionHandler:)
func (u_ UNUserNotificationCenter) GetPendingNotificationRequestsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getPendingNotificationRequestsWithCompletionHandler:"), completionHandler)
}

// Removes all of your app’s delivered notifications from Notification Center.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removeAllDeliveredNotifications()
func (u_ UNUserNotificationCenter) RemoveAllDeliveredNotifications() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllDeliveredNotifications"))
}

// Removes all of your app’s pending local notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removeAllPendingNotificationRequests()
func (u_ UNUserNotificationCenter) RemoveAllPendingNotificationRequests() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllPendingNotificationRequests"))
}

// Removes your app’s notifications from Notification Center that match the specified identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removeDeliveredNotifications(withIdentifiers:)
func (u_ UNUserNotificationCenter) RemoveDeliveredNotificationsWithIdentifiers(identifiers unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeDeliveredNotificationsWithIdentifiers:"), identifiers)
}

// Removes your app’s local notifications that are pending and match the specified identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removePendingNotificationRequests(withIdentifiers:)
func (u_ UNUserNotificationCenter) RemovePendingNotificationRequestsWithIdentifiers(identifiers unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removePendingNotificationRequestsWithIdentifiers:"), identifiers)
}

// Requests a person’s authorization to allow local and remote notifications for your app.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/requestAuthorization(options:completionHandler:)
func (u_ UNUserNotificationCenter) RequestAuthorizationWithOptionsCompletionHandler(options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("requestAuthorizationWithOptions:completionHandler:"), options, completionHandler)
}

// Updates the badge count for your app’s icon.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/setBadgeCount(_:withCompletionHandler:)
func (u_ UNUserNotificationCenter) SetBadgeCountWithCompletionHandler(newBadgeCount int, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setBadgeCount:withCompletionHandler:"), newBadgeCount, completionHandler)
}

// Registers the notification categories that your app supports.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/setNotificationCategories(_:)
func (u_ UNUserNotificationCenter) SetNotificationCategories(categories unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNotificationCategories:"), categories)
}

// The notification center’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/delegate
func (u_ UNUserNotificationCenter) Delegate() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The notification center’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/delegate
func (u_ UNUserNotificationCenter) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the device supports notification content extensions.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/supportsContentExtensions
func (u_ UNUserNotificationCenter) SupportsContentExtensions() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("supportsContentExtensions"))
	return rv
}

// The error domain for notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unerrordomain
func (u_ UNUserNotificationCenter) UNErrorDomain() string {
	rv := objc.Send[string](u_.ID, objc.Sel("UNErrorDomain"))
	return rv
}




