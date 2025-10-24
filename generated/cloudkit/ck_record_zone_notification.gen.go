// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKRecordZoneNotification] class.
var (
	CKRecordZoneNotificationClass     _CKRecordZoneNotificationClass
	CKRecordZoneNotificationClassOnce sync.Once
)

func getCKRecordZoneNotificationClass() _CKRecordZoneNotificationClass {
	CKRecordZoneNotificationClassOnce.Do(func() {
		CKRecordZoneNotificationClass = _CKRecordZoneNotificationClass{objc.GetClass("CKRecordZoneNotification")}
	})
	return CKRecordZoneNotificationClass
}

type _CKRecordZoneNotificationClass struct {
	class objc.Class
}

// An interface definition for the [CKRecordZoneNotification] class.
type ICKRecordZoneNotification interface {
	ICKNotification
	// properties:
	IsPruned() bool
	SetIsPruned(value bool)
	NotificationType() unsafe.Pointer
	SetNotificationType(value unsafe.Pointer)
	DatabaseScope() unsafe.Pointer
	SetDatabaseScope(value unsafe.Pointer)
	RecordZoneID() ICKRecordZoneID
	SetRecordZoneID(value ICKRecordZoneID)
	ShouldSendContentAvailable() bool
	SetShouldSendContentAvailable(value bool)
	NotificationInfo() objc.IObject /* cross-framework: CKNotificationInfo */
	SetNotificationInfo(value objc.IObject /* cross-framework: CKNotificationInfo */)
	// methods:
}

// A notification that triggers when the contents of a record zone change.
//
// A record zone subscription executes when a user, or in certain scenarios, CloudKit, modifies a record in that zone, for example, when a field’s value changes in a record. When CloudKit registers the change, it sends push notifications to the user’s devices to inform your app about the change. You can then fetch the changes and cache them on-device. When appropriate, CloudKit excludes the device where the change originates. You configure a subscription’s notifications by setting it’s property. Do this before you save it to the server. A subscription generates either high-priority or medium-priority push notifications. CloudKit delivers medium-priority notifications to your app in the background. High-priority notifications are visual and the system displays them to the user. Visual notifications need the user’s permission. For more information, see . A subscription uses to configure its notifications. For background delivery, set only its property to . If you set any other property, CloudKit treats the notification as high-priority. Don’t rely on push notifications for specific changes to records because the system can coalesce them. CloudKit can omit data to keep the notification’s payload size under the APNs size limit. Consider notifications an indication of remote changes. Use to determine which database contains the changed record zone, and to determine which zone contains changed records. You can then fetch just those changes using . A notification’s property is if CloudKit omits data. You don’t instantiate this class. Instead, implement in your app delegate. Initialize with the dictionary that CloudKit passes to the method. This returns an instance of the appropriate subclass. Use the property to determine the type. Then cast to that type to access type-specific properties and methods.


// A notification that triggers when the contents of a record zone change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneNotification
type CKRecordZoneNotification struct {
	CKNotification
}

// CKRecordZoneNotificationFrom constructs a [CKRecordZoneNotification] from an unsafe.Pointer.
//
// A notification that triggers when the contents of a record zone change.
func CKRecordZoneNotificationFrom(ptr unsafe.Pointer) CKRecordZoneNotification {
	return CKRecordZoneNotification{
		CKNotification: CKNotificationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKRecordZoneNotificationClass) Alloc() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKRecordZoneNotificationClass) New() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKRecordZoneNotification) Init() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKRecordZoneNotification) Autorelease() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecordZoneNotification creates a new CKRecordZoneNotification instance.
func NewCKRecordZoneNotification() CKRecordZoneNotification {
	return getCKRecordZoneNotificationClass().New()
}



// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKRecordZoneNotification) IsPruned() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPruned"))
	return rv
}


// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKRecordZoneNotification) SetIsPruned(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPruned:"), value)
}


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKRecordZoneNotification) NotificationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("notificationType"))
	return rv
}


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKRecordZoneNotification) SetNotificationType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationType:"), value)
}


// The type of database for the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonenotification/databasescope
func (c_ CKRecordZoneNotification) DatabaseScope() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("databaseScope"))
	return rv
}


// The type of database for the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonenotification/databasescope
func (c_ CKRecordZoneNotification) SetDatabaseScope(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDatabaseScope:"), value)
}


// The ID of the record zone that has changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonenotification/recordzoneid
func (c_ CKRecordZoneNotification) RecordZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("recordZoneID"))
	return rv
}


// The ID of the record zone that has changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonenotification/recordzoneid
func (c_ CKRecordZoneNotification) SetRecordZoneID(value ICKRecordZoneID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneID:"), value)
}


// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/shouldsendcontentavailable
func (c_ CKRecordZoneNotification) ShouldSendContentAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldSendContentAvailable"))
	return rv
}


// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/shouldsendcontentavailable
func (c_ CKRecordZoneNotification) SetShouldSendContentAvailable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldSendContentAvailable:"), value)
}


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKRecordZoneNotification) NotificationInfo() objc.IObject /* cross-framework: CKNotificationInfo */ {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("notificationInfo"))
	return rv
}


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKRecordZoneNotification) SetNotificationInfo(value objc.IObject /* cross-framework: CKNotificationInfo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}



