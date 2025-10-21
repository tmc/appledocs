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
}

// A notification that can be scheduled for display in the notification center.
//
// When the system delivers a notification, information about when the notification was actually presented to the user (if at all) and other details are provided in the notification object. User applications can create objects and register them with the object to notify the user when an application requires attention.
//
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


// The date this notification was actually delivered.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actualdeliverydate
func (u_ UserNotification) ActualDeliveryDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("actualDeliveryDate"))
	return rv
}


// SetActualDeliveryDate sets the value of the actualDeliveryDate property.
// The date this notification was actually delivered.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actualdeliverydate
func (u_ UserNotification) SetActualDeliveryDate(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActualDeliveryDate:"), value)
}

// Specifies the date components that control how often a user notification is repeated.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliveryrepeatinterval
func (u_ UserNotification) DeliveryRepeatInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("deliveryRepeatInterval"))
	return rv
}


// SetDeliveryRepeatInterval sets the value of the deliveryRepeatInterval property.
// Specifies the date components that control how often a user notification is repeated.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliveryrepeatinterval
func (u_ UserNotification) SetDeliveryRepeatInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryRepeatInterval:"), value)
}

// Specify the time zone to interpret the delivery date in.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverytimezone
func (u_ UserNotification) DeliveryTimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("deliveryTimeZone"))
	return rv
}


// SetDeliveryTimeZone sets the value of the deliveryTimeZone property.
// Specify the time zone to interpret the delivery date in.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverytimezone
func (u_ UserNotification) SetDeliveryTimeZone(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryTimeZone:"), value)
}

// A Boolean value that specifies whether the notification displays an action button.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/hasactionbutton
func (u_ UserNotification) HasActionButton() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasActionButton"))
	return rv
}


// SetHasActionButton sets the value of the hasActionButton property.
// A Boolean value that specifies whether the notification displays an action button.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/hasactionbutton
func (u_ UserNotification) SetHasActionButton(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHasActionButton:"), value)
}

// A Boolean value that specifies whether the notification displays a reply button.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/hasreplybutton
func (u_ UserNotification) HasReplyButton() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasReplyButton"))
	return rv
}


// SetHasReplyButton sets the value of the hasReplyButton property.
// A Boolean value that specifies whether the notification displays a reply button.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/hasreplybutton
func (u_ UserNotification) SetHasReplyButton(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHasReplyButton:"), value)
}

// A string that uniquely identifies a notification.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/identifier
func (u_ UserNotification) Identifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// A string that uniquely identifies a notification.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/identifier
func (u_ UserNotification) SetIdentifier(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// Specifies whether the user notification has been presented.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/ispresented
func (u_ UserNotification) IsPresented() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isPresented"))
	return rv
}


// SetIsPresented sets the value of the isPresented property.
// Specifies whether the user notification has been presented.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/ispresented
func (u_ UserNotification) SetIsPresented(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsPresented:"), value)
}

// Specifies whether the remote was generated by a push notification.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/isremote
func (u_ UserNotification) IsRemote() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isRemote"))
	return rv
}


// SetIsRemote sets the value of the isRemote property.
// Specifies whether the remote was generated by a push notification.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/isremote
func (u_ UserNotification) SetIsRemote(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsRemote:"), value)
}

// Optional placeholder string for inline reply field.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/responseplaceholder
func (u_ UserNotification) ResponsePlaceholder() string {
	rv := objc.Send[string](u_.ID, objc.Sel("responsePlaceholder"))
	return rv
}


// SetResponsePlaceholder sets the value of the responsePlaceholder property.
// Optional placeholder string for inline reply field.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/responseplaceholder
func (u_ UserNotification) SetResponsePlaceholder(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResponsePlaceholder:"), objc.String(value))
}

// Specifies the name of the sound to play when the notification is delivered.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/soundname
func (u_ UserNotification) SoundName() string {
	rv := objc.Send[string](u_.ID, objc.Sel("soundName"))
	return rv
}


// SetSoundName sets the value of the soundName property.
// Specifies the name of the sound to play when the notification is delivered.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/soundname
func (u_ UserNotification) SetSoundName(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSoundName:"), objc.String(value))
}

// Specifies the subtitle of the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/subtitle
func (u_ UserNotification) Subtitle() string {
	rv := objc.Send[string](u_.ID, objc.Sel("subtitle"))
	return rv
}


// SetSubtitle sets the value of the subtitle property.
// Specifies the subtitle of the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/subtitle
func (u_ UserNotification) SetSubtitle(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSubtitle:"), objc.String(value))
}

// Application-specific user info that can be attached to the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/userinfo
func (u_ UserNotification) UserInfo() string {
	rv := objc.Send[string](u_.ID, objc.Sel("userInfo"))
	return rv
}


// SetUserInfo sets the value of the userInfo property.
// Application-specific user info that can be attached to the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/userinfo
func (u_ UserNotification) SetUserInfo(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserInfo:"), objc.String(value))
}

// The default notification sound.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationdefaultsoundname
func (u_ UserNotification) NSUserNotificationDefaultSoundName() string {
	rv := objc.Send[string](u_.ID, objc.Sel("NSUserNotificationDefaultSoundName"))
	return rv
}

// Specifies the title of the action button displayed in the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/actionButtonTitle
func (u_ UserNotification) ActionButtonTitle() string {
	rv := objc.Send[string](u_.ID, objc.Sel("actionButtonTitle"))
	return rv
}


// SetActionButtonTitle sets the value of the actionButtonTitle property.
// Specifies the title of the action button displayed in the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/actionButtonTitle
func (u_ UserNotification) SetActionButtonTitle(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionButtonTitle:"), objc.String(value))
}

// Specifies what caused a user notification to occur.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/activationType-swift.property
func (u_ UserNotification) ActivationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("activationType"))
	return rv
}

// The actions that can be taken on a notification in addition to the default action.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/additionalActions
func (u_ UserNotification) AdditionalActions() []UserNotificationAction {
	rv := objc.Send[[]UserNotificationAction](u_.ID, objc.Sel("additionalActions"))
	return rv
}


// SetAdditionalActions sets the value of the additionalActions property.
// The actions that can be taken on a notification in addition to the default action.

//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/additionalActivationAction
func (u_ UserNotification) AdditionalActivationAction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("additionalActivationAction"))
	return rv
}

// Image shown in the content of the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/contentImage
func (u_ UserNotification) ContentImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("contentImage"))
	return rv
}


// SetContentImage sets the value of the contentImage property.
// Image shown in the content of the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/contentImage
func (u_ UserNotification) SetContentImage(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setContentImage:"), value)
}

// Specifies when the notification should be delivered.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/deliveryDate
func (u_ UserNotification) DeliveryDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("deliveryDate"))
	return rv
}


// SetDeliveryDate sets the value of the deliveryDate property.
// Specifies when the notification should be delivered.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/deliveryDate
func (u_ UserNotification) SetDeliveryDate(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryDate:"), value)
}

// The body text of the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/informativeText
func (u_ UserNotification) InformativeText() string {
	rv := objc.Send[string](u_.ID, objc.Sel("informativeText"))
	return rv
}


// SetInformativeText sets the value of the informativeText property.
// The body text of the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/informativeText
func (u_ UserNotification) SetInformativeText(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInformativeText:"), objc.String(value))
}

// Specifies whether the user notification has been presented.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/isPresented
func (u_ UserNotification) Presented() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("presented"))
	return rv
}

// Specifies whether the remote was generated by a push notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/isRemote
func (u_ UserNotification) Remote() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("remote"))
	return rv
}

// Specifies a custom title for the close button in an alert-style notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/otherButtonTitle
func (u_ UserNotification) OtherButtonTitle() string {
	rv := objc.Send[string](u_.ID, objc.Sel("otherButtonTitle"))
	return rv
}


// SetOtherButtonTitle sets the value of the otherButtonTitle property.
// Specifies a custom title for the close button in an alert-style notification.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/otherButtonTitle
func (u_ UserNotification) SetOtherButtonTitle(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOtherButtonTitle:"), objc.String(value))
}

// The response with which the user responded to a notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/response
func (u_ UserNotification) Response() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("response"))
	return rv
}

// Specifies the title of the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/title
func (u_ UserNotification) Title() string {
	rv := objc.Send[string](u_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// Specifies the title of the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/title
func (u_ UserNotification) SetTitle(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), objc.String(value))
}



