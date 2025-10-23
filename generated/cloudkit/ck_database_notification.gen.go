// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKDatabaseNotification] class.
var (
	CKDatabaseNotificationClass     _CKDatabaseNotificationClass
	CKDatabaseNotificationClassOnce sync.Once
)

func getCKDatabaseNotificationClass() _CKDatabaseNotificationClass {
	CKDatabaseNotificationClassOnce.Do(func() {
		CKDatabaseNotificationClass = _CKDatabaseNotificationClass{objc.GetClass("CKDatabaseNotification")}
	})
	return CKDatabaseNotificationClass
}

type _CKDatabaseNotificationClass struct {
	class objc.Class
}

// An interface definition for the [CKDatabaseNotification] class.
type ICKDatabaseNotification interface {
	ICKNotification
	DatabaseScope() unsafe.Pointer
	SetDatabaseScope(value unsafe.Pointer)
	IsPruned() bool
	SetIsPruned(value bool)
	NotificationType() unsafe.Pointer
	SetNotificationType(value unsafe.Pointer)
	ShouldSendContentAvailable() bool
	SetShouldSendContentAvailable(value bool)
	NotificationInfo() CKNotificationInfo
	SetNotificationInfo(value ICKNotificationInfo)
}

// A notification that triggers when the contents of a database change.
//
// Database subscriptions execute when changes happen in any of a database’s record zones, for example, when CloudKit saves a new record. When the subscription registers a change, it sends push notifications to the user’s devices to inform your app about the change. You can then fetch the changes and cache them on-device. When appropriate, CloudKit excludes the device where the change originates. You configure a subscription’s notifications by setting it’s property. Do this before you save it to the server. A subscription generates either high-priority or medium-priority push notifications. CloudKit delivers medium-priority notifications to your app in the background. High-priority notifications are visual and the system displays them to the user. Visual notifications need the user’s permission. For more information, see . A subscription uses to configure its notifications. For background delivery, set only its property to . If you set any other property, CloudKit treats the notification as high-priority. Don’t rely on push notifications for specific changes because the system can coalesce them. CloudKit can omit data to keep the notification’s payload size under the APNs size limit. Consider notifications an indication of remote changes. Use to determine which database has changes, and then to fetch those changes. A notification’s property is if CloudKit omits data. You don’t instantiate this class. Instead, implement in your app delegate. Initialize with the dictionary that CloudKit passes to the method. This returns an instance of the appropriate subclass. Use the property to determine the type. Then cast to that type to access type-specific properties and methods.


// A notification that triggers when the contents of a database change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabaseNotification
type CKDatabaseNotification struct {
	CKNotification
}

// CKDatabaseNotificationFrom constructs a [CKDatabaseNotification] from an unsafe.Pointer.
//
// A notification that triggers when the contents of a database change.
func CKDatabaseNotificationFrom(ptr unsafe.Pointer) CKDatabaseNotification {
	return CKDatabaseNotification{
		CKNotification: CKNotificationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKDatabaseNotificationClass) Alloc() CKDatabaseNotification {
	rv := objc.Send[CKDatabaseNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKDatabaseNotificationClass) New() CKDatabaseNotification {
	rv := objc.Send[CKDatabaseNotification](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKDatabaseNotification) Init() CKDatabaseNotification {
	rv := objc.Send[CKDatabaseNotification](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKDatabaseNotification) Autorelease() CKDatabaseNotification {
	rv := objc.Send[CKDatabaseNotification](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDatabaseNotification creates a new CKDatabaseNotification instance.
func NewCKDatabaseNotification() CKDatabaseNotification {
	return getCKDatabaseNotificationClass().New()
}



// The type of database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabasenotification/databasescope
func (c_ CKDatabaseNotification) DatabaseScope() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("databaseScope"))
	return rv
}


// The type of database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabasenotification/databasescope
func (c_ CKDatabaseNotification) SetDatabaseScope(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDatabaseScope:"), value)
}


// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKDatabaseNotification) IsPruned() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPruned"))
	return rv
}


// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKDatabaseNotification) SetIsPruned(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPruned:"), value)
}


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKDatabaseNotification) NotificationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("notificationType"))
	return rv
}


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKDatabaseNotification) SetNotificationType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationType:"), value)
}


// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/shouldsendcontentavailable
func (c_ CKDatabaseNotification) ShouldSendContentAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldSendContentAvailable"))
	return rv
}


// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/shouldsendcontentavailable
func (c_ CKDatabaseNotification) SetShouldSendContentAvailable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldSendContentAvailable:"), value)
}


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKDatabaseNotification) NotificationInfo() CKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("notificationInfo"))
	return rv
}


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKDatabaseNotification) SetNotificationInfo(value ICKNotificationInfo) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}



