// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKDatabaseSubscription */


/* debug [class_header]: Header for CKDatabaseSubscription */
// The class instance for the [CKDatabaseSubscription] class.
var (
	CKDatabaseSubscriptionClass     _CKDatabaseSubscriptionClass
	CKDatabaseSubscriptionClassOnce sync.Once
)

func getCKDatabaseSubscriptionClass() _CKDatabaseSubscriptionClass {
	CKDatabaseSubscriptionClassOnce.Do(func() {
		CKDatabaseSubscriptionClass = _CKDatabaseSubscriptionClass{objc.GetClass("CKDatabaseSubscription")}
	})
	return CKDatabaseSubscriptionClass
}

type _CKDatabaseSubscriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKDatabaseSubscription */
// An interface definition for the [CKDatabaseSubscription] class.
type ICKDatabaseSubscription interface {
	ICKSubscription
	
/* debug [class_interface_properties]: Properties for CKDatabaseSubscription */
	// properties:
	RecordType() objectivec.IObject
	SetRecordType(value objectivec.IObject)
	NotificationInfo() ICKNotificationInfo
	SetNotificationInfo(value ICKNotificationInfo)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKDatabaseSubscription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKDatabaseSubscription */
// Alloc allocates a new instance without initialization.
func (cc _CKDatabaseSubscriptionClass) Alloc() CKDatabaseSubscription {
	rv := objc.Send[CKDatabaseSubscription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKDatabaseSubscriptionClass) New() CKDatabaseSubscription {
	rv := objc.Send[CKDatabaseSubscription](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKDatabaseSubscription) Init() CKDatabaseSubscription {
	rv := objc.Send[CKDatabaseSubscription](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKDatabaseSubscription) Autorelease() CKDatabaseSubscription {
	rv := objc.Send[CKDatabaseSubscription](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDatabaseSubscription creates a new CKDatabaseSubscription instance.
func NewCKDatabaseSubscription() CKDatabaseSubscription {
	return getCKDatabaseSubscriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKDatabaseSubscription */
// A subscription that generates push notifications when CloudKit modifies records in a database.
//
// Subscriptions track the creation, modification, and deletion of records in a database, and are fundamental in keeping data on the user’s device up to date. A subscription applies only to the user that creates it. When a subscription registers a change, such as CloudKit saving a new record, it sends push notifications to the user’s devices to inform your app about the change. You can then fetch the changes and cache them on-device. When appropriate, the server excludes the device where the change originates. A database subscription executes whenever a change occurs in a custom record zone that resides in the database where you save the subscription. This is important for the shared database because you don’t know what record zones exist in advance. The only exception to this is the default record zone in the user’s private database, which doesn’t participate in database subscriptions. You can further specialize a database subscription by setting its property to a specific record type. This limits the scope of the subscription to only track changes to records of that type and reduces the number of notifications it generates. Create any subscriptions on your app’s first launch. After you initialize a subscription, save it to the server using . After the operation completes, record that state on-device (in , for example). You can then check that state on subsequent launches to prevent unnecessary trips to the server. To configure the notification that the subscription generates, set the subscription’s property. Because the system coalesces notifications, don’t rely on them for specific changes. CloudKit can omit data to keep the payload size under the APNs size limit. Consider notifications an indication of remote changes, and use to fetch the record zones that contain those changes. After you have the record zones, use to fetch the changed records in each zone. Server change tokens allow you to limit the fetch results to just the changes since your previous fetch. The example below shows how to create a database subscription in the user’s private database, configure the notifications it generates — in this case, silent push notifications — and then save that subscription to the server:


// A subscription that generates push notifications when CloudKit modifies records in a database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabaseSubscription
type CKDatabaseSubscription struct {
	CKSubscription
}

// CKDatabaseSubscriptionFrom constructs a [CKDatabaseSubscription] from an unsafe.Pointer.
//
// A subscription that generates push notifications when CloudKit modifies records in a database.
func CKDatabaseSubscriptionFrom(ptr unsafe.Pointer) CKDatabaseSubscription {
	return CKDatabaseSubscription{
		CKSubscription: CKSubscriptionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKDatabaseSubscription */

// Creates a database subscription from a serialized instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabaseSubscription/init(coder:)
func NewCKDatabaseSubscriptionWithCoder(aDecoder foundation.Coder) CKDatabaseSubscription {
	instance := getCKDatabaseSubscriptionClass().Alloc()
	rv := objc.Send[CKDatabaseSubscription](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKDatabaseSubscriptionWithCoder */


// Creates a named subscription for all records in a database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabaseSubscription/initWithSubscriptionID:
func NewCKDatabaseSubscriptionWithSubscriptionID(subscriptionID objectivec.IObject) CKDatabaseSubscription {
	instance := getCKDatabaseSubscriptionClass().Alloc()
	rv := objc.Send[CKDatabaseSubscription](instance.ID, objc.Sel("initWithSubscriptionID:"), subscriptionID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKDatabaseSubscriptionWithSubscriptionID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKDatabaseSubscription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKDatabaseSubscription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKDatabaseSubscription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKDatabaseSubscription */

// The type of record that the subscription queries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabaseSubscription/recordType-1y7dv
func (c_ CKDatabaseSubscription) RecordType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("recordType"))
	return rv
}/* debug [instance_properties/getter]: recordType */


// The type of record that the subscription queries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKDatabaseSubscription/recordType-1y7dv
func (c_ CKDatabaseSubscription) SetRecordType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordType:"), value)
}/* debug [instance_properties/setter]: recordType */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKDatabaseSubscription) NotificationInfo() ICKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("notificationInfo"))
	return rv
}/* debug [instance_properties/getter]: notificationInfo */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKDatabaseSubscription) SetNotificationInfo(value ICKNotificationInfo) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}/* debug [instance_properties/setter]: notificationInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKDatabaseSubscription */


