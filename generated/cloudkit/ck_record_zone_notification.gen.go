// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CKRecordZoneNotification */


/* debug [class_header]: Header for CKRecordZoneNotification */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKRecordZoneNotification */
// An interface definition for the [CKRecordZoneNotification] class.
type ICKRecordZoneNotification interface {
	ICKNotification
	
/* debug [class_interface_properties]: Properties for CKRecordZoneNotification */
	// properties:
	DatabaseScope() CKDatabaseScope
	RecordZoneID() ICKRecordZoneID
	IsPruned() bool
	SetIsPruned(value bool)
	NotificationType() objectivec.IObject
	SetNotificationType(value objectivec.IObject)
	ShouldSendContentAvailable() bool
	SetShouldSendContentAvailable(value bool)
	NotificationInfo() ICKNotificationInfo
	SetNotificationInfo(value ICKNotificationInfo)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKRecordZoneNotification */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKRecordZoneNotification */
// Alloc allocates a new instance without initialization.
func (cc _CKRecordZoneNotificationClass) Alloc() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKRecordZoneNotification */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKRecordZoneNotification *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKRecordZoneNotification */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKRecordZoneNotification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKRecordZoneNotification */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKRecordZoneNotification */

// The type of database for the record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneNotification/databaseScope
func (c_ CKRecordZoneNotification) DatabaseScope() CKDatabaseScope {
	rv := objc.Send[CKDatabaseScope](c_.ID, objc.Sel("databaseScope"))
	return rv
}/* debug [instance_properties/getter]: databaseScope */


// The ID of the record zone that has changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneNotification/recordZoneID
func (c_ CKRecordZoneNotification) RecordZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("recordZoneID"))
	return rv
}/* debug [instance_properties/getter]: recordZoneID */


// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKRecordZoneNotification) IsPruned() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPruned"))
	return rv
}/* debug [instance_properties/getter]: isPruned */


// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKRecordZoneNotification) SetIsPruned(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPruned:"), value)
}/* debug [instance_properties/setter]: isPruned */


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKRecordZoneNotification) NotificationType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("notificationType"))
	return rv
}/* debug [instance_properties/getter]: notificationType */


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKRecordZoneNotification) SetNotificationType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationType:"), value)
}/* debug [instance_properties/setter]: notificationType */


// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/shouldsendcontentavailable
func (c_ CKRecordZoneNotification) ShouldSendContentAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldSendContentAvailable"))
	return rv
}/* debug [instance_properties/getter]: shouldSendContentAvailable */


// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/shouldsendcontentavailable
func (c_ CKRecordZoneNotification) SetShouldSendContentAvailable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldSendContentAvailable:"), value)
}/* debug [instance_properties/setter]: shouldSendContentAvailable */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKRecordZoneNotification) NotificationInfo() ICKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("notificationInfo"))
	return rv
}/* debug [instance_properties/getter]: notificationInfo */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKRecordZoneNotification) SetNotificationInfo(value ICKNotificationInfo) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}/* debug [instance_properties/setter]: notificationInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKRecordZoneNotification */



