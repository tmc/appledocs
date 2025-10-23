// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	SubscriptionOwnerUserRecordID() CKRecordID
	AlertActionLocalizationKey() string
	SetAlertActionLocalizationKey(value string)
	AlertBody() string
	SetAlertBody(value string)
	AlertLaunchImage() string
	SetAlertLaunchImage(value string)
	AlertLocalizationArgs() string
	SetAlertLocalizationArgs(value string)
	AlertLocalizationKey() string
	SetAlertLocalizationKey(value string)
	Badge() foundation.Number
	SetBadge(value foundation.Number)
	Category() string
	SetCategory(value string)
	ContainerIdentifier() string
	SetContainerIdentifier(value string)
	IsPruned() bool
	SetIsPruned(value bool)
	NotificationID() unsafe.Pointer
	SetNotificationID(value unsafe.Pointer)
	NotificationType() unsafe.Pointer
	SetNotificationType(value unsafe.Pointer)
	SoundName() string
	SetSoundName(value string)
	SubscriptionID() unsafe.Pointer
	SetSubscriptionID(value unsafe.Pointer)
	Subtitle() string
	SetSubtitle(value string)
	SubtitleLocalizationArgs() string
	SetSubtitleLocalizationArgs(value string)
	SubtitleLocalizationKey() string
	SetSubtitleLocalizationKey(value string)
	Title() string
	SetTitle(value string)
	TitleLocalizationArgs() string
	SetTitleLocalizationArgs(value string)
	TitleLocalizationKey() string
	SetTitleLocalizationKey(value string)
}

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



// The ID of the user record that creates the subscription that generates the push notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotification/subscriptionOwnerUserRecordID
func (c_ CKNotification) SubscriptionOwnerUserRecordID() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("subscriptionOwnerUserRecordID"))
	return rv
}


// The key that identifies the localized string for the notification’s action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/alertactionlocalizationkey
func (c_ CKNotification) AlertActionLocalizationKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("alertActionLocalizationKey"))
	return rv
}


// The key that identifies the localized string for the notification’s action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/alertactionlocalizationkey
func (c_ CKNotification) SetAlertActionLocalizationKey(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertActionLocalizationKey:"), objc.String(value))
}


// The notification’s alert body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/alertbody
func (c_ CKNotification) AlertBody() string {
	rv := objc.Send[string](c_.ID, objc.Sel("alertBody"))
	return rv
}


// The notification’s alert body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/alertbody
func (c_ CKNotification) SetAlertBody(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertBody:"), objc.String(value))
}


// The filename of an image to use as a launch image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/alertlaunchimage
func (c_ CKNotification) AlertLaunchImage() string {
	rv := objc.Send[string](c_.ID, objc.Sel("alertLaunchImage"))
	return rv
}


// The filename of an image to use as a launch image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/alertlaunchimage
func (c_ CKNotification) SetAlertLaunchImage(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertLaunchImage:"), objc.String(value))
}


// The fields for building a notification’s alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/alertlocalizationargs
func (c_ CKNotification) AlertLocalizationArgs() string {
	rv := objc.Send[string](c_.ID, objc.Sel("alertLocalizationArgs"))
	return rv
}


// The fields for building a notification’s alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/alertlocalizationargs
func (c_ CKNotification) SetAlertLocalizationArgs(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertLocalizationArgs:"), objc.String(value))
}


// The key that identifies the localized text for the alert body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/alertlocalizationkey
func (c_ CKNotification) AlertLocalizationKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("alertLocalizationKey"))
	return rv
}


// The key that identifies the localized text for the alert body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/alertlocalizationkey
func (c_ CKNotification) SetAlertLocalizationKey(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertLocalizationKey:"), objc.String(value))
}


// The value that the app icon’s badge displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/badge
func (c_ CKNotification) Badge() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("badge"))
	return rv
}


// The value that the app icon’s badge displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/badge
func (c_ CKNotification) SetBadge(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBadge:"), value)
}


// The name of the action group that corresponds to this notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/category
func (c_ CKNotification) Category() string {
	rv := objc.Send[string](c_.ID, objc.Sel("category"))
	return rv
}


// The name of the action group that corresponds to this notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/category
func (c_ CKNotification) SetCategory(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCategory:"), objc.String(value))
}


// The ID of the container with the content that triggers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/containeridentifier
func (c_ CKNotification) ContainerIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}


// The ID of the container with the content that triggers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/containeridentifier
func (c_ CKNotification) SetContainerIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), objc.String(value))
}


// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKNotification) IsPruned() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPruned"))
	return rv
}


// A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/ispruned
func (c_ CKNotification) SetIsPruned(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPruned:"), value)
}


// The notification’s ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationid
func (c_ CKNotification) NotificationID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("notificationID"))
	return rv
}


// The notification’s ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationid
func (c_ CKNotification) SetNotificationID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationID:"), value)
}


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKNotification) NotificationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("notificationType"))
	return rv
}


// The type of event that generates the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/notificationtype-swift.property
func (c_ CKNotification) SetNotificationType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationType:"), value)
}


// The name of the sound file to play when a notification arrives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/soundname
func (c_ CKNotification) SoundName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("soundName"))
	return rv
}


// The name of the sound file to play when a notification arrives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/soundname
func (c_ CKNotification) SetSoundName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSoundName:"), objc.String(value))
}


// The ID of the subscription that triggers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/subscriptionid-16ygj
func (c_ CKNotification) SubscriptionID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("subscriptionID"))
	return rv
}


// The ID of the subscription that triggers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/subscriptionid-16ygj
func (c_ CKNotification) SetSubscriptionID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubscriptionID:"), value)
}


// The notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/subtitle
func (c_ CKNotification) Subtitle() string {
	rv := objc.Send[string](c_.ID, objc.Sel("subtitle"))
	return rv
}


// The notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/subtitle
func (c_ CKNotification) SetSubtitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubtitle:"), objc.String(value))
}


// The fields for building a notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/subtitlelocalizationargs
func (c_ CKNotification) SubtitleLocalizationArgs() string {
	rv := objc.Send[string](c_.ID, objc.Sel("subtitleLocalizationArgs"))
	return rv
}


// The fields for building a notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/subtitlelocalizationargs
func (c_ CKNotification) SetSubtitleLocalizationArgs(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubtitleLocalizationArgs:"), objc.String(value))
}


// The key that identifies the localized string for the notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/subtitlelocalizationkey
func (c_ CKNotification) SubtitleLocalizationKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("subtitleLocalizationKey"))
	return rv
}


// The key that identifies the localized string for the notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/subtitlelocalizationkey
func (c_ CKNotification) SetSubtitleLocalizationKey(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubtitleLocalizationKey:"), objc.String(value))
}


// The notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/title
func (c_ CKNotification) Title() string {
	rv := objc.Send[string](c_.ID, objc.Sel("title"))
	return rv
}


// The notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/title
func (c_ CKNotification) SetTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The fields for building a notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/titlelocalizationargs
func (c_ CKNotification) TitleLocalizationArgs() string {
	rv := objc.Send[string](c_.ID, objc.Sel("titleLocalizationArgs"))
	return rv
}


// The fields for building a notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/titlelocalizationargs
func (c_ CKNotification) SetTitleLocalizationArgs(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitleLocalizationArgs:"), objc.String(value))
}


// The key that identifies the localized string for the notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/titlelocalizationkey
func (c_ CKNotification) TitleLocalizationKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("titleLocalizationKey"))
	return rv
}


// The key that identifies the localized string for the notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotification/titlelocalizationkey
func (c_ CKNotification) SetTitleLocalizationKey(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitleLocalizationKey:"), objc.String(value))
}



