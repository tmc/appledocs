// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKNotification */


/* debug [class_header]: Header for CKNotification */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKNotification */
// An interface definition for the [CKNotification] class.
type ICKNotification interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKNotification */
	// properties:
	AlertActionLocalizationKey() objc.IObject /* cross-framework: NSString */
	AlertBody() objc.IObject /* cross-framework: NSString */
	AlertLaunchImage() objc.IObject /* cross-framework: NSString */
	AlertLocalizationArgs() []string
	AlertLocalizationKey() objc.IObject /* cross-framework: NSString */
	Badge() objc.IObject /* cross-framework: NSNumber */
	Category() objc.IObject /* cross-framework: NSString */
	ContainerIdentifier() objc.IObject /* cross-framework: NSString */
	IsPruned() bool
	NotificationID() ICKNotificationID
	NotificationType() CKNotificationType
	SoundName() objc.IObject /* cross-framework: NSString */
	SubscriptionID() objectivec.IObject
	SubscriptionOwnerUserRecordID() ICKRecordID
	Subtitle() objc.IObject /* cross-framework: NSString */
	SubtitleLocalizationArgs() []string
	SubtitleLocalizationKey() objc.IObject /* cross-framework: NSString */
	Title() objc.IObject /* cross-framework: NSString */
	TitleLocalizationArgs() []string
	TitleLocalizationKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKNotification */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKNotification */
// Alloc allocates a new instance without initialization.
func (cc _CKNotificationClass) Alloc() CKNotification {
	rv := objc.Send[CKNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKNotification */
// The abstract base class for CloudKit notifications.
//
// Use subclasses of to extract data from push notifications that the system receives, or to fetch a container’s previous push notifications. In both cases, the object indicates the changed data. is an abstract class. When you create a notification from a payload dictionary, the method returns an instance of the appropriate subclass. Similarly, when you fetch notifications from a container, you receive instances of a concrete subclass. provides information about the push notification and its method of delivery. Subclasses contain specific data that provides the changes.


// The abstract base class for CloudKit notifications.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKNotification */

// Creates a new notification using the specified payload data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/init(fromRemoteNotificationDictionary:)
func NewCKNotificationFromRemoteNotificationDictionary(notificationDictionary objc.IObject /* cross-framework: NSDictionary */) CKNotification {
	rv := objc.Send[CKNotification](objc.ID(getCKNotificationClass().class), objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return rv
}/* debug [class_init_methods/constructor]: NewCKNotificationFromRemoteNotificationDictionary */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKNotification */

// Creates a new notification using the specified payload data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/init(fromRemoteNotificationDictionary:)
func (cc _CKNotificationClass) NotificationFromRemoteNotificationDictionary(notificationDictionary objc.IObject /* cross-framework: NSDictionary */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NotificationFromRemoteNotificationDictionary) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKNotification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKNotification */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKNotification */

// The key that identifies the localized string for the notification’s action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/alertActionLocalizationKey
func (c_ CKNotification) AlertActionLocalizationKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("alertActionLocalizationKey"))
	return rv
}/* debug [instance_properties/getter]: alertActionLocalizationKey */


// The notification’s alert body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/alertBody
func (c_ CKNotification) AlertBody() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("alertBody"))
	return rv
}/* debug [instance_properties/getter]: alertBody */


// The filename of an image to use as a launch image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/alertLaunchImage
func (c_ CKNotification) AlertLaunchImage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("alertLaunchImage"))
	return rv
}/* debug [instance_properties/getter]: alertLaunchImage */


// The fields for building a notification’s alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/alertLocalizationArgs
func (c_ CKNotification) AlertLocalizationArgs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("alertLocalizationArgs"))
	return rv
}/* debug [instance_properties/getter]: alertLocalizationArgs */


// The key that identifies the localized text for the alert body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/alertLocalizationKey
func (c_ CKNotification) AlertLocalizationKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("alertLocalizationKey"))
	return rv
}/* debug [instance_properties/getter]: alertLocalizationKey */


// The value that the app icon’s badge displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/badge
func (c_ CKNotification) Badge() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("badge"))
	return rv
}/* debug [instance_properties/getter]: badge */


// The name of the action group that corresponds to this notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/category
func (c_ CKNotification) Category() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("category"))
	return rv
}/* debug [instance_properties/getter]: category */


// The ID of the container with the content that triggers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/containerIdentifier
func (c_ CKNotification) ContainerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}/* debug [instance_properties/getter]: containerIdentifier */


// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/isPruned
func (c_ CKNotification) IsPruned() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPruned"))
	return rv
}/* debug [instance_properties/getter]: isPruned */


// The notification’s ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/notificationID
func (c_ CKNotification) NotificationID() ICKNotificationID {
	rv := objc.Send[CKNotificationID](c_.ID, objc.Sel("notificationID"))
	return rv
}/* debug [instance_properties/getter]: notificationID */


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/notificationType-swift.property
func (c_ CKNotification) NotificationType() CKNotificationType {
	rv := objc.Send[CKNotificationType](c_.ID, objc.Sel("notificationType"))
	return rv
}/* debug [instance_properties/getter]: notificationType */


// The name of the sound file to play when a notification arrives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/soundName
func (c_ CKNotification) SoundName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("soundName"))
	return rv
}/* debug [instance_properties/getter]: soundName */


// The ID of the subscription that triggers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/subscriptionID-90zhj
func (c_ CKNotification) SubscriptionID() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("subscriptionID"))
	return rv
}/* debug [instance_properties/getter]: subscriptionID */


// The ID of the user record that creates the subscription that generates the push notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/subscriptionOwnerUserRecordID
func (c_ CKNotification) SubscriptionOwnerUserRecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("subscriptionOwnerUserRecordID"))
	return rv
}/* debug [instance_properties/getter]: subscriptionOwnerUserRecordID */


// The notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/subtitle
func (c_ CKNotification) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subtitle"))
	return rv
}/* debug [instance_properties/getter]: subtitle */


// The fields for building a notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/subtitleLocalizationArgs
func (c_ CKNotification) SubtitleLocalizationArgs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("subtitleLocalizationArgs"))
	return rv
}/* debug [instance_properties/getter]: subtitleLocalizationArgs */


// The key that identifies the localized string for the notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/subtitleLocalizationKey
func (c_ CKNotification) SubtitleLocalizationKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subtitleLocalizationKey"))
	return rv
}/* debug [instance_properties/getter]: subtitleLocalizationKey */


// The notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/title
func (c_ CKNotification) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The fields for building a notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/titleLocalizationArgs
func (c_ CKNotification) TitleLocalizationArgs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("titleLocalizationArgs"))
	return rv
}/* debug [instance_properties/getter]: titleLocalizationArgs */


// The key that identifies the localized string for the notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/titleLocalizationKey
func (c_ CKNotification) TitleLocalizationKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("titleLocalizationKey"))
	return rv
}/* debug [instance_properties/getter]: titleLocalizationKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKNotification */


