// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKNotificationInfo */


/* debug [class_header]: Header for CKNotificationInfo */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKNotificationInfo */
// An interface definition for the [CKNotificationInfo] class.
type ICKNotificationInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKNotificationInfo */
	// properties:
	AlertLocalizationArgs() []string
	SetAlertLocalizationArgs(value []string)
	DesiredKeys() []string
	SetDesiredKeys(value []string)
	SubtitleLocalizationArgs() []string
	SetSubtitleLocalizationArgs(value []string)
	TitleLocalizationArgs() []string
	SetTitleLocalizationArgs(value []string)
	AlertActionLocalizationKey() objc.IObject /* cross-framework: NSString */
	SetAlertActionLocalizationKey(value objc.IObject /* cross-framework: NSString */)
	AlertBody() objc.IObject /* cross-framework: NSString */
	SetAlertBody(value objc.IObject /* cross-framework: NSString */)
	AlertLaunchImage() objc.IObject /* cross-framework: NSString */
	SetAlertLaunchImage(value objc.IObject /* cross-framework: NSString */)
	AlertLocalizationKey() objc.IObject /* cross-framework: NSString */
	SetAlertLocalizationKey(value objc.IObject /* cross-framework: NSString */)
	Category() objc.IObject /* cross-framework: NSString */
	SetCategory(value objc.IObject /* cross-framework: NSString */)
	CollapseIDKey() objc.IObject /* cross-framework: NSString */
	SetCollapseIDKey(value objc.IObject /* cross-framework: NSString */)
	ShouldBadge() bool
	SetShouldBadge(value bool)
	ShouldSendContentAvailable() bool
	SetShouldSendContentAvailable(value bool)
	ShouldSendMutableContent() bool
	SetShouldSendMutableContent(value bool)
	SoundName() objc.IObject /* cross-framework: NSString */
	SetSoundName(value objc.IObject /* cross-framework: NSString */)
	Subtitle() objc.IObject /* cross-framework: NSString */
	SetSubtitle(value objc.IObject /* cross-framework: NSString */)
	SubtitleLocalizationKey() objc.IObject /* cross-framework: NSString */
	SetSubtitleLocalizationKey(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	TitleLocalizationKey() objc.IObject /* cross-framework: NSString */
	SetTitleLocalizationKey(value objc.IObject /* cross-framework: NSString */)
	NotificationInfo() ICKNotificationInfo
	SetNotificationInfo(value ICKNotificationInfo)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKNotificationInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKNotificationInfo */
// Alloc allocates a new instance without initialization.
func (cc _CKNotificationInfoClass) Alloc() CKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKNotificationInfo */
// An object that describes the configuration of a subscription’s push notifications.
//
// When configuring a subscription, use this class to specify the type of push notifications you want to generate when conditions meet the subscription’s trigger. You can provide content that the system displays to the user, describe the sounds to play, and indicate whether the app’s icon has a badge. You can request that the notification include information about the record that triggers it. When your app receives a push notification that a subscription generates, instantiate an instance of using the method and pass the notification’s payload. The object that the method returns contains the data you specify when configuring the subscription. For more information about push notification alerts and how they display to the user, see in .


// An object that describes the configuration of a subscription’s push notifications.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKNotificationInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKNotificationInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKNotificationInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKNotificationInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKNotificationInfo */

// The fields for building a notification’s alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/alertLocalizationArgs
func (c_ CKNotificationInfo) AlertLocalizationArgs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("alertLocalizationArgs"))
	return rv
}/* debug [instance_properties/getter]: alertLocalizationArgs */


// The fields for building a notification’s alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/alertLocalizationArgs
func (c_ CKNotificationInfo) SetAlertLocalizationArgs(value []string) {
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
}/* debug [instance_properties/setter]: alertLocalizationArgs */


// The names of fields to include in the push notification’s payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/desiredKeys
func (c_ CKNotificationInfo) DesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}/* debug [instance_properties/getter]: desiredKeys */


// The names of fields to include in the push notification’s payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/desiredKeys
func (c_ CKNotificationInfo) SetDesiredKeys(value []string) {
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
}/* debug [instance_properties/setter]: desiredKeys */


// The fields for building a notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/subtitleLocalizationArgs
func (c_ CKNotificationInfo) SubtitleLocalizationArgs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("subtitleLocalizationArgs"))
	return rv
}/* debug [instance_properties/getter]: subtitleLocalizationArgs */


// The fields for building a notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/subtitleLocalizationArgs
func (c_ CKNotificationInfo) SetSubtitleLocalizationArgs(value []string) {
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
}/* debug [instance_properties/setter]: subtitleLocalizationArgs */


// The fields for building a notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/titleLocalizationArgs
func (c_ CKNotificationInfo) TitleLocalizationArgs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("titleLocalizationArgs"))
	return rv
}/* debug [instance_properties/getter]: titleLocalizationArgs */


// The fields for building a notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKNotificationInfo/titleLocalizationArgs
func (c_ CKNotificationInfo) SetTitleLocalizationArgs(value []string) {
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
}/* debug [instance_properties/setter]: titleLocalizationArgs */


// The key that identifies the localized string for the notification’s action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertActionLocalizationKey
func (c_ CKNotificationInfo) AlertActionLocalizationKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("alertActionLocalizationKey"))
	return rv
}/* debug [instance_properties/getter]: alertActionLocalizationKey */


// The key that identifies the localized string for the notification’s action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertActionLocalizationKey
func (c_ CKNotificationInfo) SetAlertActionLocalizationKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertActionLocalizationKey:"), value)
}/* debug [instance_properties/setter]: alertActionLocalizationKey */


// The text for the notification’s alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertBody
func (c_ CKNotificationInfo) AlertBody() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("alertBody"))
	return rv
}/* debug [instance_properties/getter]: alertBody */


// The text for the notification’s alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertBody
func (c_ CKNotificationInfo) SetAlertBody(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertBody:"), value)
}/* debug [instance_properties/setter]: alertBody */


// The filename of an image to use as a launch image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertLaunchImage
func (c_ CKNotificationInfo) AlertLaunchImage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("alertLaunchImage"))
	return rv
}/* debug [instance_properties/getter]: alertLaunchImage */


// The filename of an image to use as a launch image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertLaunchImage
func (c_ CKNotificationInfo) SetAlertLaunchImage(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertLaunchImage:"), value)
}/* debug [instance_properties/setter]: alertLaunchImage */


// The key that identifies the localized string for the notification’s alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertLocalizationKey
func (c_ CKNotificationInfo) AlertLocalizationKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("alertLocalizationKey"))
	return rv
}/* debug [instance_properties/getter]: alertLocalizationKey */


// The key that identifies the localized string for the notification’s alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertLocalizationKey
func (c_ CKNotificationInfo) SetAlertLocalizationKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlertLocalizationKey:"), value)
}/* debug [instance_properties/setter]: alertLocalizationKey */


// The name of the action group that corresponds to this notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/category
func (c_ CKNotificationInfo) Category() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("category"))
	return rv
}/* debug [instance_properties/getter]: category */


// The name of the action group that corresponds to this notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/category
func (c_ CKNotificationInfo) SetCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCategory:"), value)
}/* debug [instance_properties/setter]: category */


// A value that the system uses to coalesce unseen push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/collapseIDKey
func (c_ CKNotificationInfo) CollapseIDKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("collapseIDKey"))
	return rv
}/* debug [instance_properties/getter]: collapseIDKey */


// A value that the system uses to coalesce unseen push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/collapseIDKey
func (c_ CKNotificationInfo) SetCollapseIDKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollapseIDKey:"), value)
}/* debug [instance_properties/setter]: collapseIDKey */


// A Boolean value that determines whether an app’s icon badge increments its value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldBadge
func (c_ CKNotificationInfo) ShouldBadge() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldBadge"))
	return rv
}/* debug [instance_properties/getter]: shouldBadge */


// A Boolean value that determines whether an app’s icon badge increments its value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldBadge
func (c_ CKNotificationInfo) SetShouldBadge(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldBadge:"), value)
}/* debug [instance_properties/setter]: shouldBadge */


// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldSendContentAvailable
func (c_ CKNotificationInfo) ShouldSendContentAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldSendContentAvailable"))
	return rv
}/* debug [instance_properties/getter]: shouldSendContentAvailable */


// A Boolean value that indicates whether the push notification includes the content available flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldSendContentAvailable
func (c_ CKNotificationInfo) SetShouldSendContentAvailable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldSendContentAvailable:"), value)
}/* debug [instance_properties/setter]: shouldSendContentAvailable */


// A Boolean value that indicates whether the push notification sets the mutable content flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldSendMutableContent
func (c_ CKNotificationInfo) ShouldSendMutableContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldSendMutableContent"))
	return rv
}/* debug [instance_properties/getter]: shouldSendMutableContent */


// A Boolean value that indicates whether the push notification sets the mutable content flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldSendMutableContent
func (c_ CKNotificationInfo) SetShouldSendMutableContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldSendMutableContent:"), value)
}/* debug [instance_properties/setter]: shouldSendMutableContent */


// The filename of the sound file to play when a notification arrives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/soundName
func (c_ CKNotificationInfo) SoundName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("soundName"))
	return rv
}/* debug [instance_properties/getter]: soundName */


// The filename of the sound file to play when a notification arrives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/soundName
func (c_ CKNotificationInfo) SetSoundName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSoundName:"), value)
}/* debug [instance_properties/setter]: soundName */


// The notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/subtitle
func (c_ CKNotificationInfo) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subtitle"))
	return rv
}/* debug [instance_properties/getter]: subtitle */


// The notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/subtitle
func (c_ CKNotificationInfo) SetSubtitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubtitle:"), value)
}/* debug [instance_properties/setter]: subtitle */


// The key that identifies the localized string for the notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/subtitleLocalizationKey
func (c_ CKNotificationInfo) SubtitleLocalizationKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subtitleLocalizationKey"))
	return rv
}/* debug [instance_properties/getter]: subtitleLocalizationKey */


// The key that identifies the localized string for the notification’s subtitle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/subtitleLocalizationKey
func (c_ CKNotificationInfo) SetSubtitleLocalizationKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubtitleLocalizationKey:"), value)
}/* debug [instance_properties/setter]: subtitleLocalizationKey */


// The notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/title
func (c_ CKNotificationInfo) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/title
func (c_ CKNotificationInfo) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The key that identifies the localized string for the notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/titleLocalizationKey
func (c_ CKNotificationInfo) TitleLocalizationKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("titleLocalizationKey"))
	return rv
}/* debug [instance_properties/getter]: titleLocalizationKey */


// The key that identifies the localized string for the notification’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/titleLocalizationKey
func (c_ CKNotificationInfo) SetTitleLocalizationKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitleLocalizationKey:"), value)
}/* debug [instance_properties/setter]: titleLocalizationKey */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKNotificationInfo) NotificationInfo() ICKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c_.ID, objc.Sel("notificationInfo"))
	return rv
}/* debug [instance_properties/getter]: notificationInfo */


// The configuration for a subscription’s push notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksubscription/notificationinfo-swift.property
func (c_ CKNotificationInfo) SetNotificationInfo(value ICKNotificationInfo) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNotificationInfo:"), value)
}/* debug [instance_properties/setter]: notificationInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKNotificationInfo */



