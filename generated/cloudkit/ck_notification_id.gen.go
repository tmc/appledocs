// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKNotificationID */


/* debug [class_header]: Header for CKNotificationID */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKNotificationID */
// An interface definition for the [CKNotificationID] class.
type ICKNotificationID interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKNotificationID */
	// properties:
	ContainerIdentifier() objc.IObject /* cross-framework: NSString */
	SetContainerIdentifier(value objc.IObject /* cross-framework: NSString */)
	NotificationID() ICKNotificationID
	SetNotificationID(value ICKNotificationID)
	NotificationType() objectivec.IObject
	SetNotificationType(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKNotificationID */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKNotificationID */
// Alloc allocates a new instance without initialization.
func (cc _CKNotificationIDClass) Alloc() CKNotificationID {
	rv := objc.Send[CKNotificationID](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKNotificationID */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKNotificationID *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKNotificationID */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKNotificationID */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKNotificationID */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKNotificationID */

// The ID of the container with the content that triggers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/containeridentifier
func (c_ CKNotificationID) ContainerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}/* debug [instance_properties/getter]: containerIdentifier */


// The ID of the container with the content that triggers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/containeridentifier
func (c_ CKNotificationID) SetContainerIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), value)
}/* debug [instance_properties/setter]: containerIdentifier */


// The notification’s ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationid
func (c_ CKNotificationID) NotificationID() ICKNotificationID {
	rv := objc.Send[CKNotificationID](c_.ID, objc.Sel("notificationID"))
	return rv
}/* debug [instance_properties/getter]: notificationID */


// The notification’s ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationid
func (c_ CKNotificationID) SetNotificationID(value ICKNotificationID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationID:"), value)
}/* debug [instance_properties/setter]: notificationID */


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKNotificationID) NotificationType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("notificationType"))
	return rv
}/* debug [instance_properties/getter]: notificationType */


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKNotificationID) SetNotificationType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationType:"), value)
}/* debug [instance_properties/setter]: notificationType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKNotificationID */



