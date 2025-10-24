// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKNotificationID] class.
var (
	CKNotificationIDClass     _CKNotificationIDClass
	CKNotificationIDClassOnce sync.Once
)

func getCKNotificationIDClass() _CKNotificationIDClass {
	CKNotificationIDClassOnce.Do(func() {
		CKNotificationIDClass = _CKNotificationIDClass{objc.GetClass("CKNotificationID")}
	})
	return CKNotificationIDClass
}

type _CKNotificationIDClass struct {
	class objc.Class
}

// An interface definition for the [CKNotificationID] class.
type ICKNotificationID interface {
	objectivec.IObject
	// properties:
	ContainerIdentifier() objc.IObject /* cross-framework: NSString */
	SetContainerIdentifier(value objc.IObject /* cross-framework: NSString */)
	NotificationID() ICKNotificationID
	SetNotificationID(value ICKNotificationID)
	NotificationType() unsafe.Pointer
	SetNotificationType(value unsafe.Pointer)
	// methods:
}

// An object that uniquely identifies a push notification that a container sends.
//
// You don’t create notification IDs directly. The server creates them when it creates instances of that correspond to the push notifications that CloudKit sends to your app. You can compare two IDs using the method to determine whether two notifications are the same. This class defines no methods or properties.


// An object that uniquely identifies a push notification that a container sends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/ID
type CKNotificationID struct {
	objectivec.Object
}

// CKNotificationIDFrom constructs a [CKNotificationID] from an unsafe.Pointer.
//
// An object that uniquely identifies a push notification that a container sends.
func CKNotificationIDFrom(ptr unsafe.Pointer) CKNotificationID {
	return CKNotificationID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKNotificationIDClass) Alloc() CKNotificationID {
	rv := objc.Send[CKNotificationID](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKNotificationIDClass) New() CKNotificationID {
	rv := objc.Send[CKNotificationID](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKNotificationID) Init() CKNotificationID {
	rv := objc.Send[CKNotificationID](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKNotificationID) Autorelease() CKNotificationID {
	rv := objc.Send[CKNotificationID](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKNotificationID creates a new CKNotificationID instance.
func NewCKNotificationID() CKNotificationID {
	return getCKNotificationIDClass().New()
}



// The ID of the container with the content that triggers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/containeridentifier
func (c_ CKNotificationID) ContainerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}


// The ID of the container with the content that triggers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/containeridentifier
func (c_ CKNotificationID) SetContainerIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), value)
}


// The notification’s ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationid
func (c_ CKNotificationID) NotificationID() ICKNotificationID {
	rv := objc.Send[CKNotificationID](c_.ID, objc.Sel("notificationID"))
	return rv
}


// The notification’s ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationid
func (c_ CKNotificationID) SetNotificationID(value ICKNotificationID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationID:"), value)
}


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKNotificationID) NotificationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("notificationType"))
	return rv
}


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKNotificationID) SetNotificationType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationType:"), value)
}



