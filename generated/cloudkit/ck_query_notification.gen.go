// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKQueryNotification] class.
var (
	CKQueryNotificationClass     _CKQueryNotificationClass
	CKQueryNotificationClassOnce sync.Once
)

func getCKQueryNotificationClass() _CKQueryNotificationClass {
	CKQueryNotificationClassOnce.Do(func() {
		CKQueryNotificationClass = _CKQueryNotificationClass{objc.GetClass("CKQueryNotification")}
	})
	return CKQueryNotificationClass
}

type _CKQueryNotificationClass struct {
	class objc.Class
}

// An interface definition for the [CKQueryNotification] class.
type ICKQueryNotification interface {
	ICKNotification
}

// A notification that triggers when a record that matches the subscription’s predicate changes.
//
// Query subscriptions execute when a record that matches the subscription’s predicate changes, for example, when the user modifies a field’s value in the record. When CloudKit registers the change, it sends push notifications to the user’s devices to inform your app about the change. You can then fetch the changes and cache them on-device. When appropriate, CloudKit excludes the device where the change originates. You configure a subscription’s notifications by setting it’s property. Do this before you save it to the server. A subscription generates either high-priority or medium-priority push notifications. CloudKit delivers medium-priority notifications to your app in the background. High-priority notifications are visual and the system displays them to the user. Visual notifications need the user’s permission. For more information, see . A subscription uses to configure its notifications. For background delivery, set only its property to . If you set any other property, CloudKit treats the notification as high-priority. Don’t rely on push notifications for changes because the system can coalesce them. CloudKit can omit data to keep the notification’s payload size under the APNs size limit. If you use to include extra data in the payload, the server removes that first. A notification’s property is if CloudKit omits data. Consider notifications an indication of remote changes. Use to determine which database contains the changed record. To fetch the changes, configure an instance of to match the subscription and then execute it in the database. CloudKit returns all records that match the predicate, including the changed record. Dispose of any records you cache on-device and use the operation’s results instead. You don’t instantiate this class. Instead, implement in your app delegate. Initialize with the dictionary that CloudKit passes to the method. This returns an instance of the appropriate subclass. Use the property to determine the type. Then cast to that type to access type-specific properties and methods.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryNotification
type CKQueryNotification struct {
	CKNotification
}

// CKQueryNotificationFrom constructs a [CKQueryNotification] from an unsafe.Pointer.
//
// A notification that triggers when a record that matches the subscription’s predicate changes.
func CKQueryNotificationFrom(ptr unsafe.Pointer) CKQueryNotification {
	return CKQueryNotification{
		CKNotification: CKNotificationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKQueryNotificationClass) Alloc() CKQueryNotification {
	rv := objc.Send[CKQueryNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKQueryNotificationClass) New() CKQueryNotification {
	rv := objc.Send[CKQueryNotification](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKQueryNotification) Init() CKQueryNotification {
	rv := objc.Send[CKQueryNotification](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKQueryNotification) Autorelease() CKQueryNotification {
	rv := objc.Send[CKQueryNotification](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKQueryNotification creates a new CKQueryNotification instance.
func NewCKQueryNotification() CKQueryNotification {
	return getCKQueryNotificationClass().New()
}


// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKQueryNotification) IsPruned() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPruned"))
	return rv
}


// SetIsPruned sets the value of the isPruned property.
// A Boolean value that indicates whether the system removes some push notification content before delivery.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKQueryNotification) SetIsPruned(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPruned:"), value)
}

// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/shouldsendcontentavailable
func (c_ CKQueryNotification) ShouldSendContentAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldSendContentAvailable"))
	return rv
}


// SetShouldSendContentAvailable sets the value of the shouldSendContentAvailable property.
// A Boolean value that indicates whether the push notification includes the content available flag.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/shouldsendcontentavailable
func (c_ CKQueryNotification) SetShouldSendContentAvailable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldSendContentAvailable:"), value)
}

// The configuration for a subscription’s push notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKQueryNotification) NotificationInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("notificationInfo"))
	return rv
}


// SetNotificationInfo sets the value of the notificationInfo property.
// The configuration for a subscription’s push notifications.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKQueryNotification) SetNotificationInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}

// A dictionary of fields that have changes.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerynotification/recordfields
func (c_ CKQueryNotification) RecordFields() string {
	rv := objc.Send[string](c_.ID, objc.Sel("recordFields"))
	return rv
}


// SetRecordFields sets the value of the recordFields property.
// A dictionary of fields that have changes.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerynotification/recordfields
func (c_ CKQueryNotification) SetRecordFields(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordFields:"), objc.String(value))
}

// The names of fields to include in the push notification’s payload.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/desiredkeys
func (c_ CKQueryNotification) DesiredKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("desiredKeys"))
	return rv
}


// SetDesiredKeys sets the value of the desiredKeys property.
// The names of fields to include in the push notification’s payload.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/desiredkeys
func (c_ CKQueryNotification) SetDesiredKeys(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredKeys:"), value)
}

// The event that triggers the push notification.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerynotification/querynotificationreason
func (c_ CKQueryNotification) QueryNotificationReason() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("queryNotificationReason"))
	return rv
}


// SetQueryNotificationReason sets the value of the queryNotificationReason property.
// The event that triggers the push notification.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerynotification/querynotificationreason
func (c_ CKQueryNotification) SetQueryNotificationReason(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQueryNotificationReason:"), value)
}

// The type of event that generates the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKQueryNotification) NotificationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("notificationType"))
	return rv
}


// SetNotificationType sets the value of the notificationType property.
// The type of event that generates the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKQueryNotification) SetNotificationType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationType:"), value)
}

// The type of database for the record zone.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryNotification/databaseScope
func (c_ CKQueryNotification) DatabaseScope() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("databaseScope"))
	return rv
}

// The ID of the record that CloudKit creates, updates, or deletes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryNotification/recordID
func (c_ CKQueryNotification) RecordID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordID"))
	return rv
}



