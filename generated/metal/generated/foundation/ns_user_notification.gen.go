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
	// properties:
	ActionButtonTitle() IString
	SetActionButtonTitle(value IString)
	ActivationType() unsafe.Pointer
	SetActivationType(value unsafe.Pointer)
	ActualDeliveryDate() IDate
	SetActualDeliveryDate(value IDate)
	AdditionalActions() IUserNotificationAction
	SetAdditionalActions(value IUserNotificationAction)
	AdditionalActivationAction() IUserNotificationAction
	SetAdditionalActivationAction(value IUserNotificationAction)
	ContentImage() Image /* not a class type */
	SetContentImage(value Image /* not a class type */)
	DeliveryDate() IDate
	SetDeliveryDate(value IDate)
	DeliveryRepeatInterval() IDateComponents
	SetDeliveryRepeatInterval(value IDateComponents)
	DeliveryTimeZone() ITimeZone
	SetDeliveryTimeZone(value ITimeZone)
	HasActionButton() bool
	SetHasActionButton(value bool)
	HasReplyButton() bool
	SetHasReplyButton(value bool)
	Identifier() IString
	SetIdentifier(value IString)
	InformativeText() IString
	SetInformativeText(value IString)
	IsPresented() bool
	SetIsPresented(value bool)
	IsRemote() bool
	SetIsRemote(value bool)
	OtherButtonTitle() IString
	SetOtherButtonTitle(value IString)
	Response() IAttributedString
	SetResponse(value IAttributedString)
	ResponsePlaceholder() IString
	SetResponsePlaceholder(value IString)
	SoundName() IString
	SetSoundName(value IString)
	Subtitle() IString
	SetSubtitle(value IString)
	Title() IString
	SetTitle(value IString)
	UserInfo() IString
	SetUserInfo(value IString)
	NSUserNotificationDefaultSoundName() IString
	// methods:
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
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actionbuttontitle
func (u_ UserNotification) ActionButtonTitle() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("actionButtonTitle"))
	return rv
}


// Specifies the title of the action button displayed in the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actionbuttontitle
func (u_ UserNotification) SetActionButtonTitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionButtonTitle:"), value)
}


// Specifies what caused a user notification to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/activationtype-swift.property
func (u_ UserNotification) ActivationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("activationType"))
	return rv
}


// Specifies what caused a user notification to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/activationtype-swift.property
func (u_ UserNotification) SetActivationType(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActivationType:"), value)
}


// The date this notification was actually delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actualdeliverydate
func (u_ UserNotification) ActualDeliveryDate() IDate {
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


// The actions that can be taken on a notification in addition to the default action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/additionalactions
func (u_ UserNotification) AdditionalActions() IUserNotificationAction {
	rv := objc.Send[UserNotificationAction](u_.ID, objc.Sel("additionalActions"))
	return rv
}


// The actions that can be taken on a notification in addition to the default action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/additionalactions
func (u_ UserNotification) SetAdditionalActions(value IUserNotificationAction) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAdditionalActions:"), value)
}


// An additional action selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/additionalactivationaction
func (u_ UserNotification) AdditionalActivationAction() IUserNotificationAction {
	rv := objc.Send[UserNotificationAction](u_.ID, objc.Sel("additionalActivationAction"))
	return rv
}


// An additional action selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/additionalactivationaction
func (u_ UserNotification) SetAdditionalActivationAction(value IUserNotificationAction) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAdditionalActivationAction:"), value)
}


// Image shown in the content of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/contentimage
func (u_ UserNotification) ContentImage() Image /* not a class type */ {
	rv := objc.Send[Image](u_.ID, objc.Sel("contentImage"))
	return rv
}


// Image shown in the content of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/contentimage
func (u_ UserNotification) SetContentImage(value Image /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setContentImage:"), value)
}


// Specifies when the notification should be delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverydate
func (u_ UserNotification) DeliveryDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("deliveryDate"))
	return rv
}


// Specifies when the notification should be delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverydate
func (u_ UserNotification) SetDeliveryDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryDate:"), value)
}


// Specifies the date components that control how often a user notification is repeated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliveryrepeatinterval
func (u_ UserNotification) DeliveryRepeatInterval() IDateComponents {
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
func (u_ UserNotification) DeliveryTimeZone() ITimeZone {
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
func (u_ UserNotification) Identifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("identifier"))
	return rv
}


// A string that uniquely identifies a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/identifier
func (u_ UserNotification) SetIdentifier(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIdentifier:"), value)
}


// The body text of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/informativetext
func (u_ UserNotification) InformativeText() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("informativeText"))
	return rv
}


// The body text of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/informativetext
func (u_ UserNotification) SetInformativeText(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInformativeText:"), value)
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


// Specifies a custom title for the close button in an alert-style notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/otherbuttontitle
func (u_ UserNotification) OtherButtonTitle() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("otherButtonTitle"))
	return rv
}


// Specifies a custom title for the close button in an alert-style notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/otherbuttontitle
func (u_ UserNotification) SetOtherButtonTitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOtherButtonTitle:"), value)
}


// The response with which the user responded to a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/response
func (u_ UserNotification) Response() IAttributedString {
	rv := objc.Send[AttributedString](u_.ID, objc.Sel("response"))
	return rv
}


// The response with which the user responded to a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/response
func (u_ UserNotification) SetResponse(value IAttributedString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResponse:"), value)
}


// Optional placeholder string for inline reply field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/responseplaceholder
func (u_ UserNotification) ResponsePlaceholder() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("responsePlaceholder"))
	return rv
}


// Optional placeholder string for inline reply field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/responseplaceholder
func (u_ UserNotification) SetResponsePlaceholder(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResponsePlaceholder:"), value)
}


// Specifies the name of the sound to play when the notification is delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/soundname
func (u_ UserNotification) SoundName() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("soundName"))
	return rv
}


// Specifies the name of the sound to play when the notification is delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/soundname
func (u_ UserNotification) SetSoundName(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSoundName:"), value)
}


// Specifies the subtitle of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/subtitle
func (u_ UserNotification) Subtitle() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("subtitle"))
	return rv
}


// Specifies the subtitle of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/subtitle
func (u_ UserNotification) SetSubtitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSubtitle:"), value)
}


// Specifies the title of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/title
func (u_ UserNotification) Title() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("title"))
	return rv
}


// Specifies the title of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/title
func (u_ UserNotification) SetTitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), value)
}


// Application-specific user info that can be attached to the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/userinfo
func (u_ UserNotification) UserInfo() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("userInfo"))
	return rv
}


// Application-specific user info that can be attached to the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/userinfo
func (u_ UserNotification) SetUserInfo(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserInfo:"), value)
}


// The default notification sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationdefaultsoundname
func (u_ UserNotification) NSUserNotificationDefaultSoundName() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("NSUserNotificationDefaultSoundName"))
	return rv
}



