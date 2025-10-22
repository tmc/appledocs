// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserNotification] class.
var (
	UserNotificationClass     _UserNotificationClass
	UserNotificationClassOnce sync.Once
)

func getUserNotificationClass() _UserNotificationClass {
	UserNotificationClassOnce.Do(func() {
		UserNotificationClass = _UserNotificationClass{objc.GetClass("NSUserNotification")}
	})
	return UserNotificationClass
}

type _UserNotificationClass struct {
	class objc.Class
}

// An interface definition for the [UserNotification] class.
type IUserNotification interface {
	objectivec.IObject
	ActionButtonTitle() string
	SetActionButtonTitle(value string)
	ActivationType() UserNotificationActivationType
	AdditionalActions() []UserNotificationAction
	SetAdditionalActions(value []UserNotificationAction)
	AdditionalActivationAction() NSUserNotificationAction
	DeliveryDate() NSDate
	SetDeliveryDate(value IDate)
	InformativeText() string
	SetInformativeText(value string)
	Presented() bool
	Remote() bool
	OtherButtonTitle() string
	SetOtherButtonTitle(value string)
	Response() NSAttributedString
	Title() string
	SetTitle(value string)
	ActualDeliveryDate() Date
	SetActualDeliveryDate(value IDate)
	DeliveryRepeatInterval() DateComponents
	SetDeliveryRepeatInterval(value IDateComponents)
	DeliveryTimeZone() TimeZone
	SetDeliveryTimeZone(value ITimeZone)
	HasActionButton() bool
	SetHasActionButton(value bool)
	HasReplyButton() bool
	SetHasReplyButton(value bool)
	Identifier() string
	SetIdentifier(value string)
	IsPresented() bool
	SetIsPresented(value bool)
	IsRemote() bool
	SetIsRemote(value bool)
	ResponsePlaceholder() string
	SetResponsePlaceholder(value string)
	SoundName() string
	SetSoundName(value string)
	Subtitle() string
	SetSubtitle(value string)
	UserInfo() string
	SetUserInfo(value string)
	NSUserNotificationDefaultSoundName() string
}

// A notification that can be scheduled for display in the notification center.
//
// When the system delivers a notification, information about when the notification was actually presented to the user (if at all) and other details are provided in the notification object. User applications can create objects and register them with the object to notify the user when an application requires attention.


// A notification that can be scheduled for display in the notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification

type UserNotification struct {
	objectivec.Object
}

// UserNotificationFrom constructs a [UserNotification] from an unsafe.Pointer.
//
// A notification that can be scheduled for display in the notification center.
func UserNotificationFrom(ptr unsafe.Pointer) UserNotification {
	return UserNotification{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UserNotificationClass) Alloc() UserNotification {
	rv := objc.Send[UserNotification](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserNotificationClass) New() UserNotification {
	rv := objc.Send[UserNotification](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserNotification) Init() UserNotification {
	rv := objc.Send[UserNotification](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserNotification) Autorelease() UserNotification {
	rv := objc.Send[UserNotification](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserNotification creates a new UserNotification instance.
func NewUserNotification() UserNotification {
	return getUserNotificationClass().New()
}



// Specifies the title of the action button displayed in the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/actionButtonTitle

func (u_ UserNotification) ActionButtonTitle() string {
	rv := objc.Send[string](u_.ID, objc.Sel("actionButtonTitle"))
	return rv
}


// Specifies the title of the action button displayed in the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/actionButtonTitle

func (u_ UserNotification) SetActionButtonTitle(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionButtonTitle:"), objc.String(value))
}


// Specifies what caused a user notification to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/activationType-swift.property

func (u_ UserNotification) ActivationType() UserNotificationActivationType {
	rv := objc.Send[UserNotificationActivationType](u_.ID, objc.Sel("activationType"))
	return rv
}


// The actions that can be taken on a notification in addition to the default action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/additionalActions

func (u_ UserNotification) AdditionalActions() []UserNotificationAction {
	rv := objc.Send[[]UserNotificationAction](u_.ID, objc.Sel("additionalActions"))
	return rv
}


// The actions that can be taken on a notification in addition to the default action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/additionalActions

func (u_ UserNotification) SetAdditionalActions(value []UserNotificationAction) {
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
	objc.Send[objc.ID](u_.ID, objc.Sel("setAdditionalActions:"), nsArray)
}


// An additional action selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/additionalActivationAction

func (u_ UserNotification) AdditionalActivationAction() NSUserNotificationAction {
	rv := objc.Send[NSUserNotificationAction](u_.ID, objc.Sel("additionalActivationAction"))
	return rv
}


// Specifies when the notification should be delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/deliveryDate

func (u_ UserNotification) DeliveryDate() NSDate {
	rv := objc.Send[NSDate](u_.ID, objc.Sel("deliveryDate"))
	return rv
}


// Specifies when the notification should be delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/deliveryDate

func (u_ UserNotification) SetDeliveryDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryDate:"), value)
}


// The body text of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/informativeText

func (u_ UserNotification) InformativeText() string {
	rv := objc.Send[string](u_.ID, objc.Sel("informativeText"))
	return rv
}


// The body text of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/informativeText

func (u_ UserNotification) SetInformativeText(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInformativeText:"), objc.String(value))
}


// Specifies whether the user notification has been presented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/isPresented

func (u_ UserNotification) Presented() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("presented"))
	return rv
}


// Specifies whether the remote was generated by a push notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/isRemote

func (u_ UserNotification) Remote() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("remote"))
	return rv
}


// Specifies a custom title for the close button in an alert-style notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/otherButtonTitle

func (u_ UserNotification) OtherButtonTitle() string {
	rv := objc.Send[string](u_.ID, objc.Sel("otherButtonTitle"))
	return rv
}


// Specifies a custom title for the close button in an alert-style notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/otherButtonTitle

func (u_ UserNotification) SetOtherButtonTitle(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOtherButtonTitle:"), objc.String(value))
}


// The response with which the user responded to a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/response

func (u_ UserNotification) Response() NSAttributedString {
	rv := objc.Send[NSAttributedString](u_.ID, objc.Sel("response"))
	return rv
}


// Specifies the title of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/title

func (u_ UserNotification) Title() string {
	rv := objc.Send[string](u_.ID, objc.Sel("title"))
	return rv
}


// Specifies the title of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/title

func (u_ UserNotification) SetTitle(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The date this notification was actually delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actualdeliverydate

func (u_ UserNotification) ActualDeliveryDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("actualDeliveryDate"))
	return rv
}


// The date this notification was actually delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actualdeliverydate

func (u_ UserNotification) SetActualDeliveryDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActualDeliveryDate:"), value)
}


// Specifies the date components that control how often a user notification is repeated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliveryrepeatinterval

func (u_ UserNotification) DeliveryRepeatInterval() DateComponents {
	rv := objc.Send[DateComponents](u_.ID, objc.Sel("deliveryRepeatInterval"))
	return rv
}


// Specifies the date components that control how often a user notification is repeated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliveryrepeatinterval

func (u_ UserNotification) SetDeliveryRepeatInterval(value IDateComponents) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryRepeatInterval:"), value)
}


// Specify the time zone to interpret the delivery date in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverytimezone

func (u_ UserNotification) DeliveryTimeZone() TimeZone {
	rv := objc.Send[TimeZone](u_.ID, objc.Sel("deliveryTimeZone"))
	return rv
}


// Specify the time zone to interpret the delivery date in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverytimezone

func (u_ UserNotification) SetDeliveryTimeZone(value ITimeZone) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryTimeZone:"), value)
}


// A Boolean value that specifies whether the notification displays an action button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/hasactionbutton

func (u_ UserNotification) HasActionButton() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasActionButton"))
	return rv
}


// A Boolean value that specifies whether the notification displays an action button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/hasactionbutton

func (u_ UserNotification) SetHasActionButton(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHasActionButton:"), value)
}


// A Boolean value that specifies whether the notification displays a reply button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/hasreplybutton

func (u_ UserNotification) HasReplyButton() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasReplyButton"))
	return rv
}


// A Boolean value that specifies whether the notification displays a reply button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/hasreplybutton

func (u_ UserNotification) SetHasReplyButton(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHasReplyButton:"), value)
}


// A string that uniquely identifies a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/identifier

func (u_ UserNotification) Identifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("identifier"))
	return rv
}


// A string that uniquely identifies a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/identifier

func (u_ UserNotification) SetIdentifier(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// Specifies whether the user notification has been presented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/ispresented

func (u_ UserNotification) IsPresented() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isPresented"))
	return rv
}


// Specifies whether the user notification has been presented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/ispresented

func (u_ UserNotification) SetIsPresented(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsPresented:"), value)
}


// Specifies whether the remote was generated by a push notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/isremote

func (u_ UserNotification) IsRemote() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isRemote"))
	return rv
}


// Specifies whether the remote was generated by a push notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/isremote

func (u_ UserNotification) SetIsRemote(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsRemote:"), value)
}


// Optional placeholder string for inline reply field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/responseplaceholder

func (u_ UserNotification) ResponsePlaceholder() string {
	rv := objc.Send[string](u_.ID, objc.Sel("responsePlaceholder"))
	return rv
}


// Optional placeholder string for inline reply field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/responseplaceholder

func (u_ UserNotification) SetResponsePlaceholder(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResponsePlaceholder:"), objc.String(value))
}


// Specifies the name of the sound to play when the notification is delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/soundname

func (u_ UserNotification) SoundName() string {
	rv := objc.Send[string](u_.ID, objc.Sel("soundName"))
	return rv
}


// Specifies the name of the sound to play when the notification is delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/soundname

func (u_ UserNotification) SetSoundName(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSoundName:"), objc.String(value))
}


// Specifies the subtitle of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/subtitle

func (u_ UserNotification) Subtitle() string {
	rv := objc.Send[string](u_.ID, objc.Sel("subtitle"))
	return rv
}


// Specifies the subtitle of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/subtitle

func (u_ UserNotification) SetSubtitle(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSubtitle:"), objc.String(value))
}


// Application-specific user info that can be attached to the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/userinfo

func (u_ UserNotification) UserInfo() string {
	rv := objc.Send[string](u_.ID, objc.Sel("userInfo"))
	return rv
}


// Application-specific user info that can be attached to the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/userinfo

func (u_ UserNotification) SetUserInfo(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserInfo:"), objc.String(value))
}


// The default notification sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationdefaultsoundname

func (u_ UserNotification) NSUserNotificationDefaultSoundName() string {
	rv := objc.Send[string](u_.ID, objc.Sel("NSUserNotificationDefaultSoundName"))
	return rv
}



