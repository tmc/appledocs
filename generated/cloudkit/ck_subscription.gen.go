// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CKSubscription] class.
type ICKSubscription interface {
	objectivec.IObject
	NotificationInfo() CKNotificationInfo
	SetNotificationInfo(value ICKNotificationInfo)
	SubscriptionID() unsafe.Pointer
	SubscriptionType() unsafe.Pointer
	DesiredKeys() unsafe.Pointer
	SetDesiredKeys(value unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (cc _CKSubscriptionClass) Alloc() CKSubscription {
	rv := objc.Send[CKSubscription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/notificationInfo-swift.property

func (c_ CKSubscription) NotificationInfo() CKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("notificationInfo"))
	return rv
}


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/notificationInfo-swift.property

func (c_ CKSubscription) SetNotificationInfo(value ICKNotificationInfo) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}


// The subscription’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/subscriptionID-12vxy

func (c_ CKSubscription) SubscriptionID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("subscriptionID"))
	return rv
}


// The behavior that a subscription provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/subscriptionType-swift.property

func (c_ CKSubscription) SubscriptionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("subscriptionType"))
	return rv
}


// The names of fields to include in the push notification’s payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/desiredkeys

func (c_ CKSubscription) DesiredKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("desiredKeys"))
	return rv
}


// The names of fields to include in the push notification’s payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.class/desiredkeys

func (c_ CKSubscription) SetDesiredKeys(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredKeys:"), value)
}



