// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKQuerySubscription */


/* debug [class_header]: Header for CKQuerySubscription */
// The class instance for the [CKQuerySubscription] class.
var (
	CKQuerySubscriptionClass     _CKQuerySubscriptionClass
	CKQuerySubscriptionClassOnce sync.Once
)

func getCKQuerySubscriptionClass() _CKQuerySubscriptionClass {
	CKQuerySubscriptionClassOnce.Do(func() {
		CKQuerySubscriptionClass = _CKQuerySubscriptionClass{objc.GetClass("CKQuerySubscription")}
	})
	return CKQuerySubscriptionClass
}

type _CKQuerySubscriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKQuerySubscription */
// An interface definition for the [CKQuerySubscription] class.
type ICKQuerySubscription interface {
	ICKSubscription
	
/* debug [class_interface_properties]: Properties for CKQuerySubscription */
	// properties:
	Predicate() foundation.Predicate
	QuerySubscriptionOptions() CKQuerySubscriptionOptions
	RecordType() objectivec.IObject
	ZoneID() ICKRecordZoneID
	SetZoneID(value ICKRecordZoneID)
	NotificationInfo() ICKNotificationInfo
	SetNotificationInfo(value ICKNotificationInfo)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKQuerySubscription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKQuerySubscription */
// Alloc allocates a new instance without initialization.
func (cc _CKQuerySubscriptionClass) Alloc() CKQuerySubscription {
	rv := objc.Send[CKQuerySubscription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKQuerySubscriptionClass) New() CKQuerySubscription {
	rv := objc.Send[CKQuerySubscription](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKQuerySubscription) Init() CKQuerySubscription {
	rv := objc.Send[CKQuerySubscription](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKQuerySubscription) Autorelease() CKQuerySubscription {
	rv := objc.Send[CKQuerySubscription](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKQuerySubscription creates a new CKQuerySubscription instance.
func NewCKQuerySubscription() CKQuerySubscription {
	return getCKQuerySubscriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKQuerySubscription */
// A subscription that generates push notifications when CloudKit modifies records that match a predicate.
//
// Subscriptions track the creation, modification, and deletion of records in a database, and are fundamental in keeping data on the user’s device up to date. A subscription applies only to the user that creates it. When a subscription registers a change, such as CloudKit saving a new record, it sends push notifications to the user’s devices to inform your app about the change. You can then fetch the changes and cache them on-device. When appropriate, the server excludes the device where the change originates. Query subscriptions execute whenever a change occurs in a database that matches the predicate and options you specify. You scope a query subscription to an individual record type that you provide during initialization. You can set the subscription’s property to further specialize the subscription to a specific record zone in the database. This limits the scope of the subscription to only track changes in that record zone and reduces the number of notifications it generates. For more information about defining CloudKit-compatible predicates, see . Create any subscriptions on your app’s first launch. After you initialize a subscription, save it to the server using . When the operation completes, record that state on-device (in , for example). You can then check that state on subsequent launches to prevent unnecessary trips to the server. To configure the notification the subscription generates, set the subscription’s property. Because the system coalesces notifications, don’t rely on them for specific changes. CloudKit can omit data to keep the payload size under the APNs size limit. Consider notifications an indication of remote changes and use to fetch the changed records. Create the operation with an instance of that you configure with the same record type and predicate as the subscription. If you limit the subscription to a specific record zone, set the operation’s property to that record zone’s ID. Because doesn’t employ server change tokens, dispose of any records you cache on-device and use the query’s results instead. The example below shows how to create a query subscription in the user’s private database, configure the notifications it generates — in this case, silent push notifications — and then save that subscription to the server:


// A subscription that generates push notifications when CloudKit modifies records that match a predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription
type CKQuerySubscription struct {
	CKSubscription
}

// CKQuerySubscriptionFrom constructs a [CKQuerySubscription] from an unsafe.Pointer.
//
// A subscription that generates push notifications when CloudKit modifies records that match a predicate.
func CKQuerySubscriptionFrom(ptr unsafe.Pointer) CKQuerySubscription {
	return CKQuerySubscription{
		CKSubscription: CKSubscriptionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKQuerySubscription */

// Creates a query-based subscription from a serialized instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/init(coder:)
func NewCKQuerySubscriptionWithCoder(aDecoder foundation.Coder) CKQuerySubscription {
	instance := getCKQuerySubscriptionClass().Alloc()
	rv := objc.Send[CKQuerySubscription](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKQuerySubscriptionWithCoder */


// Creates a query-based subscription that queries records of a specific type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/initWithRecordType:predicate:options:
func NewCKQuerySubscriptionWithRecordTypePredicateOptions(recordType objectivec.IObject, predicate foundation.Predicate, querySubscriptionOptions CKQuerySubscriptionOptions) CKQuerySubscription {
	instance := getCKQuerySubscriptionClass().Alloc()
	rv := objc.Send[CKQuerySubscription](instance.ID, objc.Sel("initWithRecordType:predicate:options:"), recordType, predicate, querySubscriptionOptions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKQuerySubscriptionWithRecordTypePredicateOptions */


// Creates a named query-based subscription that queries records of a specific type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/initWithRecordType:predicate:subscriptionID:options:
func NewCKQuerySubscriptionWithRecordTypePredicateSubscriptionIDOptions(recordType objectivec.IObject, predicate foundation.Predicate, subscriptionID objectivec.IObject, querySubscriptionOptions CKQuerySubscriptionOptions) CKQuerySubscription {
	instance := getCKQuerySubscriptionClass().Alloc()
	rv := objc.Send[CKQuerySubscription](instance.ID, objc.Sel("initWithRecordType:predicate:subscriptionID:options:"), recordType, predicate, subscriptionID, querySubscriptionOptions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKQuerySubscriptionWithRecordTypePredicateSubscriptionIDOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKQuerySubscription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKQuerySubscription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKQuerySubscription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKQuerySubscription */

// The matching criteria to apply to records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/predicate
func (c_ CKQuerySubscription) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](c_.ID, objc.Sel("predicate"))
	return rv
}/* debug [instance_properties/getter]: predicate */


// Options that define the behavior of the subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/querySubscriptionOptions
func (c_ CKQuerySubscription) QuerySubscriptionOptions() CKQuerySubscriptionOptions {
	rv := objc.Send[CKQuerySubscriptionOptions](c_.ID, objc.Sel("querySubscriptionOptions"))
	return rv
}/* debug [instance_properties/getter]: querySubscriptionOptions */


// The type of record that the subscription queries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/recordType-5l3zs
func (c_ CKQuerySubscription) RecordType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("recordType"))
	return rv
}/* debug [instance_properties/getter]: recordType */


// The ID of the record zone that the subscription queries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/zoneID
func (c_ CKQuerySubscription) ZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}/* debug [instance_properties/getter]: zoneID */


// The ID of the record zone that the subscription queries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuerySubscription/zoneID
func (c_ CKQuerySubscription) SetZoneID(value ICKRecordZoneID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZoneID:"), value)
}/* debug [instance_properties/setter]: zoneID */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKQuerySubscription) NotificationInfo() ICKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("notificationInfo"))
	return rv
}/* debug [instance_properties/getter]: notificationInfo */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKQuerySubscription) SetNotificationInfo(value ICKNotificationInfo) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}/* debug [instance_properties/setter]: notificationInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKQuerySubscription */


