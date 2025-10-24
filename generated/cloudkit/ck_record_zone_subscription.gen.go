// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKRecordZoneSubscription */


/* debug [class_header]: Header for CKRecordZoneSubscription */
// The class instance for the [CKRecordZoneSubscription] class.
var (
	CKRecordZoneSubscriptionClass     _CKRecordZoneSubscriptionClass
	CKRecordZoneSubscriptionClassOnce sync.Once
)

func getCKRecordZoneSubscriptionClass() _CKRecordZoneSubscriptionClass {
	CKRecordZoneSubscriptionClassOnce.Do(func() {
		CKRecordZoneSubscriptionClass = _CKRecordZoneSubscriptionClass{objc.GetClass("CKRecordZoneSubscription")}
	})
	return CKRecordZoneSubscriptionClass
}

type _CKRecordZoneSubscriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKRecordZoneSubscription */
// An interface definition for the [CKRecordZoneSubscription] class.
type ICKRecordZoneSubscription interface {
	ICKSubscription
	
/* debug [class_interface_properties]: Properties for CKRecordZoneSubscription */
	// properties:
	RecordType() objectivec.IObject
	SetRecordType(value objectivec.IObject)
	ZoneID() ICKRecordZoneID
	NotificationInfo() ICKNotificationInfo
	SetNotificationInfo(value ICKNotificationInfo)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKRecordZoneSubscription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKRecordZoneSubscription */
// Alloc allocates a new instance without initialization.
func (cc _CKRecordZoneSubscriptionClass) Alloc() CKRecordZoneSubscription {
	rv := objc.Send[CKRecordZoneSubscription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKRecordZoneSubscriptionClass) New() CKRecordZoneSubscription {
	rv := objc.Send[CKRecordZoneSubscription](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKRecordZoneSubscription) Init() CKRecordZoneSubscription {
	rv := objc.Send[CKRecordZoneSubscription](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKRecordZoneSubscription) Autorelease() CKRecordZoneSubscription {
	rv := objc.Send[CKRecordZoneSubscription](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecordZoneSubscription creates a new CKRecordZoneSubscription instance.
func NewCKRecordZoneSubscription() CKRecordZoneSubscription {
	return getCKRecordZoneSubscriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKRecordZoneSubscription */
// A subscription that generates push notifications when CloudKit modifies records in a specific record zone.
//
// Subscriptions track the creation, modification, and deletion of records in a database, and are fundamental in keeping data on the user’s device up to date. A subscription applies only to the user that creates it. When a subscription registers a change, such as CloudKit saving a new record, it sends push notifications to the user’s devices to inform your app about the change. You can then fetch the changes and cache them on-device. When appropriate, the server excludes the device where the change originates. Record zone subscriptions execute whenever a change happens in the record zone you specify when you create the subscription. You can further specialize the subscription by setting its property to a specific record type. This limits the scope of the subscription to only track changes to records of that type and reduces the number of notifications it generates. Create any subscriptions on your app’s first launch. After you initialize a subscription, save it to the server using . When the operation completes, record that state on-device (in , for example). You can then check that state on subsequent launches to prevent unnecessary trips to the server. To configure the notification that the subscription generates, set the subscription’s property. Because the system coalesces notifications, don’t rely on them for specific changes. CloudKit can omit data to keep the payload size under the APNs size limit. Consider notifications an indication of remote changes and use to fetch the changed records. Server change tokens allow you to limit the fetch results to just the changes since your previous fetch. The example below shows how to create a record zone subscription in the user’s private database, configure the notifications it generates — in this case, silent push notifications — and then save that subscription to the server:


// A subscription that generates push notifications when CloudKit modifies records in a specific record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneSubscription
type CKRecordZoneSubscription struct {
	CKSubscription
}

// CKRecordZoneSubscriptionFrom constructs a [CKRecordZoneSubscription] from an unsafe.Pointer.
//
// A subscription that generates push notifications when CloudKit modifies records in a specific record zone.
func CKRecordZoneSubscriptionFrom(ptr unsafe.Pointer) CKRecordZoneSubscription {
	return CKRecordZoneSubscription{
		CKSubscription: CKSubscriptionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKRecordZoneSubscription */

// Creates a zone-based subscription from a serialized instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneSubscription/init(coder:)
func NewCKRecordZoneSubscriptionWithCoder(aDecoder foundation.Coder) CKRecordZoneSubscription {
	instance := getCKRecordZoneSubscriptionClass().Alloc()
	rv := objc.Send[CKRecordZoneSubscription](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKRecordZoneSubscriptionWithCoder */


// Creates a subscription for all records in the specified record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneSubscription/init(zoneID:)
func NewCKRecordZoneSubscriptionWithZoneID(zoneID ICKRecordZoneID) CKRecordZoneSubscription {
	instance := getCKRecordZoneSubscriptionClass().Alloc()
	rv := objc.Send[CKRecordZoneSubscription](instance.ID, objc.Sel("initWithZoneID:"), zoneID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKRecordZoneSubscriptionWithZoneID */


// Creates a named subscription for all records in the specified record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneSubscription/initWithZoneID:subscriptionID:
func NewCKRecordZoneSubscriptionWithZoneIDSubscriptionID(zoneID ICKRecordZoneID, subscriptionID objectivec.IObject) CKRecordZoneSubscription {
	instance := getCKRecordZoneSubscriptionClass().Alloc()
	rv := objc.Send[CKRecordZoneSubscription](instance.ID, objc.Sel("initWithZoneID:subscriptionID:"), zoneID, subscriptionID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKRecordZoneSubscriptionWithZoneIDSubscriptionID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKRecordZoneSubscription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKRecordZoneSubscription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKRecordZoneSubscription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKRecordZoneSubscription */

// The type of record that the subscription queries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneSubscription/recordType-1kt07
func (c_ CKRecordZoneSubscription) RecordType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("recordType"))
	return rv
}/* debug [instance_properties/getter]: recordType */


// The type of record that the subscription queries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneSubscription/recordType-1kt07
func (c_ CKRecordZoneSubscription) SetRecordType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordType:"), value)
}/* debug [instance_properties/setter]: recordType */


// The ID of the record zone that the subscription queries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneSubscription/zoneID
func (c_ CKRecordZoneSubscription) ZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}/* debug [instance_properties/getter]: zoneID */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKRecordZoneSubscription) NotificationInfo() ICKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("notificationInfo"))
	return rv
}/* debug [instance_properties/getter]: notificationInfo */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKRecordZoneSubscription) SetNotificationInfo(value ICKNotificationInfo) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}/* debug [instance_properties/setter]: notificationInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKRecordZoneSubscription */


