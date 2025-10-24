// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSubscription */


/* debug [class_header]: Header for CKSubscription */
// The class instance for the [CKSubscription] class.
var (
	CKSubscriptionClass     _CKSubscriptionClass
	CKSubscriptionClassOnce sync.Once
)

func getCKSubscriptionClass() _CKSubscriptionClass {
	CKSubscriptionClassOnce.Do(func() {
		CKSubscriptionClass = _CKSubscriptionClass{objc.GetClass("CKSubscription")}
	})
	return CKSubscriptionClass
}

type _CKSubscriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSubscription */
// An interface definition for the [CKSubscription] class.
type ICKSubscription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSubscription */
	// properties:
	NotificationInfo() ICKNotificationInfo
	SetNotificationInfo(value ICKNotificationInfo)
	SubscriptionID() objectivec.IObject
	SubscriptionType() CKSubscriptionType
	DesiredKeys() objectivec.IObject
	SetDesiredKeys(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSubscription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSubscription */
// Alloc allocates a new instance without initialization.
func (cc _CKSubscriptionClass) Alloc() CKSubscription {
	rv := objc.Send[CKSubscription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSubscriptionClass) New() CKSubscription {
	rv := objc.Send[CKSubscription](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSubscription) Init() CKSubscription {
	rv := objc.Send[CKSubscription](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSubscription) Autorelease() CKSubscription {
	rv := objc.Send[CKSubscription](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSubscription creates a new CKSubscription instance.
func NewCKSubscription() CKSubscription {
	return getCKSubscriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSubscription */
// An abstract base class for subscriptions.
//
// A subscription acts like a persistent query on the server that can track the creation, deletion, and modification of records. When changes occur, they trigger the delivery of push notifications so that your app can respond appropriately. Subscriptions don’t become active until you save them to the server and the server has time to index them. To save a subscription, use an instance of or the method of . To cancel a subscription, delete the corresponding subscription from the server. Most of a subscription’s configuration happens at initialization time. You must, however, specify how to deliver push notifications to the user’s device. Use the property to configure the delivery options. You must save the subscription before the changes take effect.


// An abstract base class for subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription
type CKSubscription struct {
	objectivec.Object
}

// CKSubscriptionFrom constructs a [CKSubscription] from an unsafe.Pointer.
//
// An abstract base class for subscriptions.
func CKSubscriptionFrom(ptr unsafe.Pointer) CKSubscription {
	return CKSubscription{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSubscription *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSubscription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSubscription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSubscription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSubscription */

// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/notificationInfo-swift.property
func (c_ CKSubscription) NotificationInfo() ICKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("notificationInfo"))
	return rv
}/* debug [instance_properties/getter]: notificationInfo */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/notificationInfo-swift.property
func (c_ CKSubscription) SetNotificationInfo(value ICKNotificationInfo) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}/* debug [instance_properties/setter]: notificationInfo */


// The subscription’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/subscriptionID-12vxy
func (c_ CKSubscription) SubscriptionID() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("subscriptionID"))
	return rv
}/* debug [instance_properties/getter]: subscriptionID */


// The behavior that a subscription provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/subscriptionType-swift.property
func (c_ CKSubscription) SubscriptionType() CKSubscriptionType {
	rv := objc.Send[CKSubscriptionType](c_.ID, objc.Sel("subscriptionType"))
	return rv
}/* debug [instance_properties/getter]: subscriptionType */


// The names of fields to include in the push notification’s payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/desiredkeys
func (c_ CKSubscription) DesiredKeys() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("desiredKeys"))
	return rv
}/* debug [instance_properties/getter]: desiredKeys */


// The names of fields to include in the push notification’s payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/desiredkeys
func (c_ CKSubscription) SetDesiredKeys(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredKeys:"), value)
}/* debug [instance_properties/setter]: desiredKeys */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSubscription */



