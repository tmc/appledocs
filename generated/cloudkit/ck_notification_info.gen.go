// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKNotificationInfo] class.
var (
	CKNotificationInfoClass     _CKNotificationInfoClass
	CKNotificationInfoClassOnce sync.Once
)

func getCKNotificationInfoClass() _CKNotificationInfoClass {
	CKNotificationInfoClassOnce.Do(func() {
		CKNotificationInfoClass = _CKNotificationInfoClass{objc.GetClass("CKNotificationInfo")}
	})
	return CKNotificationInfoClass
}

type _CKNotificationInfoClass struct {
	class objc.Class
}

// An interface definition for the [CKNotificationInfo] class.
type ICKNotificationInfo interface {
	objectivec.IObject
}

// An object that describes the configuration of a subscription’s push notifications.
//
// When configuring a subscription, use this class to specify the type of push notifications you want to generate when conditions meet the subscription’s trigger. You can provide content that the system displays to the user, describe the sounds to play, and indicate whether the app’s icon has a badge. You can request that the notification include information about the record that triggers it. When your app receives a push notification that a subscription generates, instantiate an instance of using the method and pass the notification’s payload. The object that the method returns contains the data you specify when configuring the subscription. For more information about push notification alerts and how they display to the user, see in .
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class
type CKNotificationInfo struct {
	objectivec.Object
}

// CKNotificationInfoFrom constructs a [CKNotificationInfo] from an unsafe.Pointer.
//
// An object that describes the configuration of a subscription’s push notifications.
func CKNotificationInfoFrom(ptr unsafe.Pointer) CKNotificationInfo {
	return CKNotificationInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKNotificationInfoClass) Alloc() CKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKNotificationInfoClass) New() CKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKNotificationInfo) Init() CKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKNotificationInfo) Autorelease() CKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKNotificationInfo creates a new CKNotificationInfo instance.
func NewCKNotificationInfo() CKNotificationInfo {
	return getCKNotificationInfoClass().New()
}


// The fields for building a notification’s alert.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/alertLocalizationArgs
func (c_ CKNotificationInfo) AlertLocalizationArgs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("alertLocalizationArgs"))
	return rv
}


// SetAlertLocalizationArgs sets the value of the alertLocalizationArgs property.
// The fields for building a notification’s alert.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/alertLocalizationArgs
func (c_ CKNotificationInfo) SetAlertLocalizationArgs(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertLocalizationArgs:"), nsArray)
}

// The names of fields to include in the push notification’s payload.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/desiredKeys
func (c_ CKNotificationInfo) DesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}


// SetDesiredKeys sets the value of the desiredKeys property.
// The names of fields to include in the push notification’s payload.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/desiredKeys
func (c_ CKNotificationInfo) SetDesiredKeys(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredKeys:"), nsArray)
}

// The fields for building a notification’s subtitle.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/subtitleLocalizationArgs
func (c_ CKNotificationInfo) SubtitleLocalizationArgs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("subtitleLocalizationArgs"))
	return rv
}


// SetSubtitleLocalizationArgs sets the value of the subtitleLocalizationArgs property.
// The fields for building a notification’s subtitle.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/subtitleLocalizationArgs
func (c_ CKNotificationInfo) SetSubtitleLocalizationArgs(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubtitleLocalizationArgs:"), nsArray)
}

// The fields for building a notification’s title.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/titleLocalizationArgs
func (c_ CKNotificationInfo) TitleLocalizationArgs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("titleLocalizationArgs"))
	return rv
}


// SetTitleLocalizationArgs sets the value of the titleLocalizationArgs property.
// The fields for building a notification’s title.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/titleLocalizationArgs
func (c_ CKNotificationInfo) SetTitleLocalizationArgs(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitleLocalizationArgs:"), nsArray)
}

// The key that identifies the localized string for the notification’s action.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertActionLocalizationKey
func (c_ CKNotificationInfo) AlertActionLocalizationKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("alertActionLocalizationKey"))
	return rv
}


// SetAlertActionLocalizationKey sets the value of the alertActionLocalizationKey property.
// The key that identifies the localized string for the notification’s action.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertActionLocalizationKey
func (c_ CKNotificationInfo) SetAlertActionLocalizationKey(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertActionLocalizationKey:"), objc.String(value))
}

// The text for the notification’s alert.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertBody
func (c_ CKNotificationInfo) AlertBody() string {
	rv := objc.Send[string](c_.ID, objc.Sel("alertBody"))
	return rv
}


// SetAlertBody sets the value of the alertBody property.
// The text for the notification’s alert.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertBody
func (c_ CKNotificationInfo) SetAlertBody(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertBody:"), objc.String(value))
}

// The filename of an image to use as a launch image.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertLaunchImage
func (c_ CKNotificationInfo) AlertLaunchImage() string {
	rv := objc.Send[string](c_.ID, objc.Sel("alertLaunchImage"))
	return rv
}


// SetAlertLaunchImage sets the value of the alertLaunchImage property.
// The filename of an image to use as a launch image.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertLaunchImage
func (c_ CKNotificationInfo) SetAlertLaunchImage(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertLaunchImage:"), objc.String(value))
}

// The key that identifies the localized string for the notification’s alert.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertLocalizationKey
func (c_ CKNotificationInfo) AlertLocalizationKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("alertLocalizationKey"))
	return rv
}


// SetAlertLocalizationKey sets the value of the alertLocalizationKey property.
// The key that identifies the localized string for the notification’s alert.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertLocalizationKey
func (c_ CKNotificationInfo) SetAlertLocalizationKey(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertLocalizationKey:"), objc.String(value))
}

// The name of the action group that corresponds to this notification.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/category
func (c_ CKNotificationInfo) Category() string {
	rv := objc.Send[string](c_.ID, objc.Sel("category"))
	return rv
}


// SetCategory sets the value of the category property.
// The name of the action group that corresponds to this notification.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/category
func (c_ CKNotificationInfo) SetCategory(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCategory:"), objc.String(value))
}

// A value that the system uses to coalesce unseen push notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/collapseIDKey
func (c_ CKNotificationInfo) CollapseIDKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("collapseIDKey"))
	return rv
}


// SetCollapseIDKey sets the value of the collapseIDKey property.
// A value that the system uses to coalesce unseen push notifications.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/collapseIDKey
func (c_ CKNotificationInfo) SetCollapseIDKey(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollapseIDKey:"), objc.String(value))
}

// A Boolean value that determines whether an app’s icon badge increments its value.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldBadge
func (c_ CKNotificationInfo) ShouldBadge() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldBadge"))
	return rv
}


// SetShouldBadge sets the value of the shouldBadge property.
// A Boolean value that determines whether an app’s icon badge increments its value.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldBadge
func (c_ CKNotificationInfo) SetShouldBadge(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldBadge:"), value)
}

// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldSendContentAvailable
func (c_ CKNotificationInfo) ShouldSendContentAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldSendContentAvailable"))
	return rv
}


// SetShouldSendContentAvailable sets the value of the shouldSendContentAvailable property.
// A Boolean value that indicates whether the push notification includes the content available flag.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldSendContentAvailable
func (c_ CKNotificationInfo) SetShouldSendContentAvailable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldSendContentAvailable:"), value)
}

// A Boolean value that indicates whether the push notification sets the mutable content flag.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldSendMutableContent
func (c_ CKNotificationInfo) ShouldSendMutableContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldSendMutableContent"))
	return rv
}


// SetShouldSendMutableContent sets the value of the shouldSendMutableContent property.
// A Boolean value that indicates whether the push notification sets the mutable content flag.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldSendMutableContent
func (c_ CKNotificationInfo) SetShouldSendMutableContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldSendMutableContent:"), value)
}

// The filename of the sound file to play when a notification arrives.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/soundName
func (c_ CKNotificationInfo) SoundName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("soundName"))
	return rv
}


// SetSoundName sets the value of the soundName property.
// The filename of the sound file to play when a notification arrives.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/soundName
func (c_ CKNotificationInfo) SetSoundName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSoundName:"), objc.String(value))
}

// The notification’s subtitle.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/subtitle
func (c_ CKNotificationInfo) Subtitle() string {
	rv := objc.Send[string](c_.ID, objc.Sel("subtitle"))
	return rv
}


// SetSubtitle sets the value of the subtitle property.
// The notification’s subtitle.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/subtitle
func (c_ CKNotificationInfo) SetSubtitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubtitle:"), objc.String(value))
}

// The key that identifies the localized string for the notification’s subtitle.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/subtitleLocalizationKey
func (c_ CKNotificationInfo) SubtitleLocalizationKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("subtitleLocalizationKey"))
	return rv
}


// SetSubtitleLocalizationKey sets the value of the subtitleLocalizationKey property.
// The key that identifies the localized string for the notification’s subtitle.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/subtitleLocalizationKey
func (c_ CKNotificationInfo) SetSubtitleLocalizationKey(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubtitleLocalizationKey:"), objc.String(value))
}

// The notification’s title.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/title
func (c_ CKNotificationInfo) Title() string {
	rv := objc.Send[string](c_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The notification’s title.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/title
func (c_ CKNotificationInfo) SetTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The key that identifies the localized string for the notification’s title.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/titleLocalizationKey
func (c_ CKNotificationInfo) TitleLocalizationKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("titleLocalizationKey"))
	return rv
}


// SetTitleLocalizationKey sets the value of the titleLocalizationKey property.
// The key that identifies the localized string for the notification’s title.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/titleLocalizationKey
func (c_ CKNotificationInfo) SetTitleLocalizationKey(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitleLocalizationKey:"), objc.String(value))
}

// The configuration for a subscription’s push notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKNotificationInfo) NotificationInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("notificationInfo"))
	return rv
}


// SetNotificationInfo sets the value of the notificationInfo property.
// The configuration for a subscription’s push notifications.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKNotificationInfo) SetNotificationInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}



