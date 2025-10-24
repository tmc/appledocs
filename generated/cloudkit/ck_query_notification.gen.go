// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKQueryNotification */


/* debug [class_header]: Header for CKQueryNotification */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKQueryNotification */
// An interface definition for the [CKQueryNotification] class.
type ICKQueryNotification interface {
	ICKNotification
	
/* debug [class_interface_properties]: Properties for CKQueryNotification */
	// properties:
	DatabaseScope() CKDatabaseScope
	QueryNotificationReason() CKQueryNotificationReason
	RecordFields() foundation.IDictionary
	RecordID() ICKRecordID
	IsPruned() bool
	SetIsPruned(value bool)
	NotificationType() objectivec.IObject
	SetNotificationType(value objectivec.IObject)
	DesiredKeys() objectivec.IObject
	SetDesiredKeys(value objectivec.IObject)
	ShouldSendContentAvailable() bool
	SetShouldSendContentAvailable(value bool)
	NotificationInfo() ICKNotificationInfo
	SetNotificationInfo(value ICKNotificationInfo)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKQueryNotification */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKQueryNotification */
// Alloc allocates a new instance without initialization.
func (cc _CKQueryNotificationClass) Alloc() CKQueryNotification {
	rv := objc.Send[CKQueryNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKQueryNotification */
// A notification that triggers when a record that matches the subscription’s predicate changes.
//
// Query subscriptions execute when a record that matches the subscription’s predicate changes, for example, when the user modifies a field’s value in the record. When CloudKit registers the change, it sends push notifications to the user’s devices to inform your app about the change. You can then fetch the changes and cache them on-device. When appropriate, CloudKit excludes the device where the change originates. You configure a subscription’s notifications by setting it’s property. Do this before you save it to the server. A subscription generates either high-priority or medium-priority push notifications. CloudKit delivers medium-priority notifications to your app in the background. High-priority notifications are visual and the system displays them to the user. Visual notifications need the user’s permission. For more information, see . A subscription uses to configure its notifications. For background delivery, set only its property to . If you set any other property, CloudKit treats the notification as high-priority. Don’t rely on push notifications for changes because the system can coalesce them. CloudKit can omit data to keep the notification’s payload size under the APNs size limit. If you use to include extra data in the payload, the server removes that first. A notification’s property is if CloudKit omits data. Consider notifications an indication of remote changes. Use to determine which database contains the changed record. To fetch the changes, configure an instance of to match the subscription and then execute it in the database. CloudKit returns all records that match the predicate, including the changed record. Dispose of any records you cache on-device and use the operation’s results instead. You don’t instantiate this class. Instead, implement in your app delegate. Initialize with the dictionary that CloudKit passes to the method. This returns an instance of the appropriate subclass. Use the property to determine the type. Then cast to that type to access type-specific properties and methods.


// A notification that triggers when a record that matches the subscription’s predicate changes.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKQueryNotification *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKQueryNotification */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKQueryNotification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKQueryNotification */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKQueryNotification */

// The type of database for the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryNotification/databaseScope
func (c_ CKQueryNotification) DatabaseScope() CKDatabaseScope {
	rv := objc.Send[CKDatabaseScope](c_.ID, objc.Sel("databaseScope"))
	return rv
}/* debug [instance_properties/getter]: databaseScope */


// The event that triggers the push notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryNotification/queryNotificationReason
func (c_ CKQueryNotification) QueryNotificationReason() CKQueryNotificationReason {
	rv := objc.Send[CKQueryNotificationReason](c_.ID, objc.Sel("queryNotificationReason"))
	return rv
}/* debug [instance_properties/getter]: queryNotificationReason */


// A dictionary of fields that have changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryNotification/recordFields
func (c_ CKQueryNotification) RecordFields() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("recordFields"))
	return rv
}/* debug [instance_properties/getter]: recordFields */


// The ID of the record that CloudKit creates, updates, or deletes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryNotification/recordID
func (c_ CKQueryNotification) RecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordID"))
	return rv
}/* debug [instance_properties/getter]: recordID */


// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKQueryNotification) IsPruned() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPruned"))
	return rv
}/* debug [instance_properties/getter]: isPruned */


// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKQueryNotification) SetIsPruned(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPruned:"), value)
}/* debug [instance_properties/setter]: isPruned */


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKQueryNotification) NotificationType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("notificationType"))
	return rv
}/* debug [instance_properties/getter]: notificationType */


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKQueryNotification) SetNotificationType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationType:"), value)
}/* debug [instance_properties/setter]: notificationType */


// The names of fields to include in the push notification’s payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/desiredkeys
func (c_ CKQueryNotification) DesiredKeys() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("desiredKeys"))
	return rv
}/* debug [instance_properties/getter]: desiredKeys */


// The names of fields to include in the push notification’s payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/desiredkeys
func (c_ CKQueryNotification) SetDesiredKeys(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredKeys:"), value)
}/* debug [instance_properties/setter]: desiredKeys */


// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/shouldsendcontentavailable
func (c_ CKQueryNotification) ShouldSendContentAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldSendContentAvailable"))
	return rv
}/* debug [instance_properties/getter]: shouldSendContentAvailable */


// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/shouldsendcontentavailable
func (c_ CKQueryNotification) SetShouldSendContentAvailable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldSendContentAvailable:"), value)
}/* debug [instance_properties/setter]: shouldSendContentAvailable */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKQueryNotification) NotificationInfo() ICKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("notificationInfo"))
	return rv
}/* debug [instance_properties/getter]: notificationInfo */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKQueryNotification) SetNotificationInfo(value ICKNotificationInfo) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}/* debug [instance_properties/setter]: notificationInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKQueryNotification */



