// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [CKQuerySubscription] class.
type ICKQuerySubscription interface {
	ICKSubscription
	ZoneID() ICKRecordZoneID
	SetZoneID(value ICKRecordZoneID)
	Predicate() foundation.Predicate
	SetPredicate(value foundation.Predicate)
	QuerySubscriptionOptions() unsafe.Pointer
	SetQuerySubscriptionOptions(value unsafe.Pointer)
	RecordType() unsafe.Pointer
	SetRecordType(value unsafe.Pointer)
	NotificationInfo() CKNotificationInfo
	SetNotificationInfo(value CKNotificationInfo)
}

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

// Alloc allocates a new instance without initialization.
func (cc _CKQuerySubscriptionClass) Alloc() CKQuerySubscription {
	rv := objc.Send[CKQuerySubscription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The ID of the record zone that contains the records to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/zoneid
func (c_ CKQuerySubscription) ZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}


// The ID of the record zone that contains the records to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/zoneid
func (c_ CKQuerySubscription) SetZoneID(value ICKRecordZoneID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZoneID:"), value)
}


// The matching criteria to apply to records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/predicate
func (c_ CKQuerySubscription) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](c_.ID, objc.Sel("predicate"))
	return rv
}


// The matching criteria to apply to records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/predicate
func (c_ CKQuerySubscription) SetPredicate(value foundation.Predicate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredicate:"), value)
}


// Options that define the behavior of the subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/querysubscriptionoptions
func (c_ CKQuerySubscription) QuerySubscriptionOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("querySubscriptionOptions"))
	return rv
}


// Options that define the behavior of the subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/querysubscriptionoptions
func (c_ CKQuerySubscription) SetQuerySubscriptionOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQuerySubscriptionOptions:"), value)
}


// The type of record that the subscription queries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/recordtype-4qgdo
func (c_ CKQuerySubscription) RecordType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordType"))
	return rv
}


// The type of record that the subscription queries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquerysubscription/recordtype-4qgdo
func (c_ CKQuerySubscription) SetRecordType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordType:"), value)
}


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKQuerySubscription) NotificationInfo() CKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("notificationInfo"))
	return rv
}


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKQuerySubscription) SetNotificationInfo(value CKNotificationInfo) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}



