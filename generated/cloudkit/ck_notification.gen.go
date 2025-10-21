// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKNotification] class.
var (
	CKNotificationClass     _CKNotificationClass
	CKNotificationClassOnce sync.Once
)

func getCKNotificationClass() _CKNotificationClass {
	CKNotificationClassOnce.Do(func() {
		CKNotificationClass = _CKNotificationClass{objc.GetClass("CKNotification")}
	})
	return CKNotificationClass
}

type _CKNotificationClass struct {
	class objc.Class
}

// An interface definition for the [CKNotification] class.
type ICKNotification interface {
	objectivec.IObject
}

// The abstract base class for CloudKit notifications.
//
// Use subclasses of to extract data from push notifications that the system receives, or to fetch a container’s previous push notifications. In both cases, the object indicates the changed data. is an abstract class. When you create a notification from a payload dictionary, the method returns an instance of the appropriate subclass. Similarly, when you fetch notifications from a container, you receive instances of a concrete subclass. provides information about the push notification and its method of delivery. Subclasses contain specific data that provides the changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification
type CKNotification struct {
	objectivec.Object
}

// CKNotificationFrom constructs a [CKNotification] from an unsafe.Pointer.
//
// The abstract base class for CloudKit notifications.
func CKNotificationFrom(ptr unsafe.Pointer) CKNotification {
	return CKNotification{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKNotificationClass) Alloc() CKNotification {
	rv := objc.Send[CKNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKNotificationClass) New() CKNotification {
	rv := objc.Send[CKNotification](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKNotification) Init() CKNotification {
	rv := objc.Send[CKNotification](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKNotification) Autorelease() CKNotification {
	rv := objc.Send[CKNotification](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKNotification creates a new CKNotification instance.
func NewCKNotification() CKNotification {
	return getCKNotificationClass().New()
}




// Creates a new notification using the specified payload data.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/init(fromRemoteNotificationDictionary:)
func NewCKNotificationFromRemoteNotificationDictionary(notificationDictionary objc.ID) CKNotification {
	rv := objc.Send[CKNotification](objc.ID(getCKNotificationClass().class), objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return rv
}


// Creates a new notification using the specified payload data.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/init(fromRemoteNotificationDictionary:)
func (cc _CKNotificationClass) NotificationFromRemoteNotificationDictionary(notificationDictionary objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return rv
}

// The ID of the container with the content that triggers the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/containerIdentifier
func (c_ CKNotification) ContainerIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}

// The notification’s ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/notificationID
func (c_ CKNotification) NotificationID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("notificationID"))
	return rv
}

// The type of event that generates the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/notificationType-swift.property
func (c_ CKNotification) NotificationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("notificationType"))
	return rv
}

// The ID of the user record that creates the subscription that generates the push notification.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/subscriptionOwnerUserRecordID
func (c_ CKNotification) SubscriptionOwnerUserRecordID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("subscriptionOwnerUserRecordID"))
	return rv
}


