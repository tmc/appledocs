// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSUserNotification */


/* debug [class_header]: Header for NSUserNotification */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UserNotification */
// An interface definition for the [UserNotification] class.
type IUserNotification interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UserNotification */
	// properties:
	AdditionalActivationAction() IUserNotificationAction
	HasReplyButton() bool
	SetHasReplyButton(value bool)
	ActionButtonTitle() IString
	SetActionButtonTitle(value IString)
	ActivationType() objectivec.IObject
	SetActivationType(value objectivec.IObject)
	ActualDeliveryDate() IDate
	SetActualDeliveryDate(value IDate)
	AdditionalActions() IUserNotificationAction
	SetAdditionalActions(value IUserNotificationAction)
	ContentImage() objectivec.IObject
	SetContentImage(value objectivec.IObject)
	DeliveryDate() IDate
	SetDeliveryDate(value IDate)
	DeliveryRepeatInterval() IDateComponents
	SetDeliveryRepeatInterval(value IDateComponents)
	DeliveryTimeZone() ITimeZone
	SetDeliveryTimeZone(value ITimeZone)
	HasActionButton() bool
	SetHasActionButton(value bool)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UserNotification */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UserNotification */
// Alloc allocates a new instance without initialization.
func (uc _UserNotificationClass) Alloc() UserNotification {
	rv := objc.Send[UserNotification](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UserNotification */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UserNotification *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UserNotification */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UserNotification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UserNotification */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UserNotification */

// An additional action selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/additionalActivationAction
func (u_ UserNotification) AdditionalActivationAction() IUserNotificationAction {
	rv := objc.Send[UserNotificationAction](u_.ID, objc.Sel("additionalActivationAction"))
	return rv
}/* debug [instance_properties/getter]: additionalActivationAction */


// A Boolean value that specifies whether the notification displays a reply button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/hasReplyButton
func (u_ UserNotification) HasReplyButton() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasReplyButton"))
	return rv
}/* debug [instance_properties/getter]: hasReplyButton */


// A Boolean value that specifies whether the notification displays a reply button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/hasReplyButton
func (u_ UserNotification) SetHasReplyButton(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHasReplyButton:"), value)
}/* debug [instance_properties/setter]: hasReplyButton */


// Specifies the title of the action button displayed in the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actionbuttontitle
func (u_ UserNotification) ActionButtonTitle() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("actionButtonTitle"))
	return rv
}/* debug [instance_properties/getter]: actionButtonTitle */


// Specifies the title of the action button displayed in the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actionbuttontitle
func (u_ UserNotification) SetActionButtonTitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionButtonTitle:"), value)
}/* debug [instance_properties/setter]: actionButtonTitle */


// Specifies what caused a user notification to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/activationtype-swift.property
func (u_ UserNotification) ActivationType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("activationType"))
	return rv
}/* debug [instance_properties/getter]: activationType */


// Specifies what caused a user notification to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/activationtype-swift.property
func (u_ UserNotification) SetActivationType(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActivationType:"), value)
}/* debug [instance_properties/setter]: activationType */


// The date this notification was actually delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actualdeliverydate
func (u_ UserNotification) ActualDeliveryDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("actualDeliveryDate"))
	return rv
}/* debug [instance_properties/getter]: actualDeliveryDate */


// The date this notification was actually delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actualdeliverydate
func (u_ UserNotification) SetActualDeliveryDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActualDeliveryDate:"), value)
}/* debug [instance_properties/setter]: actualDeliveryDate */


// The actions that can be taken on a notification in addition to the default action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/additionalactions
func (u_ UserNotification) AdditionalActions() IUserNotificationAction {
	rv := objc.Send[UserNotificationAction](u_.ID, objc.Sel("additionalActions"))
	return rv
}/* debug [instance_properties/getter]: additionalActions */


// The actions that can be taken on a notification in addition to the default action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/additionalactions
func (u_ UserNotification) SetAdditionalActions(value IUserNotificationAction) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAdditionalActions:"), value)
}/* debug [instance_properties/setter]: additionalActions */


// Image shown in the content of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/contentimage
func (u_ UserNotification) ContentImage() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("contentImage"))
	return rv
}/* debug [instance_properties/getter]: contentImage */


// Image shown in the content of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/contentimage
func (u_ UserNotification) SetContentImage(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setContentImage:"), value)
}/* debug [instance_properties/setter]: contentImage */


// Specifies when the notification should be delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverydate
func (u_ UserNotification) DeliveryDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("deliveryDate"))
	return rv
}/* debug [instance_properties/getter]: deliveryDate */


// Specifies when the notification should be delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverydate
func (u_ UserNotification) SetDeliveryDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryDate:"), value)
}/* debug [instance_properties/setter]: deliveryDate */


// Specifies the date components that control how often a user notification is repeated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliveryrepeatinterval
func (u_ UserNotification) DeliveryRepeatInterval() IDateComponents {
	rv := objc.Send[DateComponents](u_.ID, objc.Sel("deliveryRepeatInterval"))
	return rv
}/* debug [instance_properties/getter]: deliveryRepeatInterval */


// Specifies the date components that control how often a user notification is repeated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliveryrepeatinterval
func (u_ UserNotification) SetDeliveryRepeatInterval(value IDateComponents) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryRepeatInterval:"), value)
}/* debug [instance_properties/setter]: deliveryRepeatInterval */


// Specify the time zone to interpret the delivery date in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverytimezone
func (u_ UserNotification) DeliveryTimeZone() ITimeZone {
	rv := objc.Send[TimeZone](u_.ID, objc.Sel("deliveryTimeZone"))
	return rv
}/* debug [instance_properties/getter]: deliveryTimeZone */


// Specify the time zone to interpret the delivery date in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverytimezone
func (u_ UserNotification) SetDeliveryTimeZone(value ITimeZone) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryTimeZone:"), value)
}/* debug [instance_properties/setter]: deliveryTimeZone */


// A Boolean value that specifies whether the notification displays an action button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/hasactionbutton
func (u_ UserNotification) HasActionButton() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasActionButton"))
	return rv
}/* debug [instance_properties/getter]: hasActionButton */


// A Boolean value that specifies whether the notification displays an action button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/hasactionbutton
func (u_ UserNotification) SetHasActionButton(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHasActionButton:"), value)
}/* debug [instance_properties/setter]: hasActionButton */


// A string that uniquely identifies a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/identifier
func (u_ UserNotification) Identifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A string that uniquely identifies a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/identifier
func (u_ UserNotification) SetIdentifier(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// The body text of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/informativetext
func (u_ UserNotification) InformativeText() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("informativeText"))
	return rv
}/* debug [instance_properties/getter]: informativeText */


// The body text of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/informativetext
func (u_ UserNotification) SetInformativeText(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInformativeText:"), value)
}/* debug [instance_properties/setter]: informativeText */


// Specifies whether the user notification has been presented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/ispresented
func (u_ UserNotification) IsPresented() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isPresented"))
	return rv
}/* debug [instance_properties/getter]: isPresented */


// Specifies whether the user notification has been presented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/ispresented
func (u_ UserNotification) SetIsPresented(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsPresented:"), value)
}/* debug [instance_properties/setter]: isPresented */


// Specifies whether the remote was generated by a push notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/isremote
func (u_ UserNotification) IsRemote() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isRemote"))
	return rv
}/* debug [instance_properties/getter]: isRemote */


// Specifies whether the remote was generated by a push notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/isremote
func (u_ UserNotification) SetIsRemote(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsRemote:"), value)
}/* debug [instance_properties/setter]: isRemote */


// Specifies a custom title for the close button in an alert-style notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/otherbuttontitle
func (u_ UserNotification) OtherButtonTitle() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("otherButtonTitle"))
	return rv
}/* debug [instance_properties/getter]: otherButtonTitle */


// Specifies a custom title for the close button in an alert-style notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/otherbuttontitle
func (u_ UserNotification) SetOtherButtonTitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOtherButtonTitle:"), value)
}/* debug [instance_properties/setter]: otherButtonTitle */


// The response with which the user responded to a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/response
func (u_ UserNotification) Response() IAttributedString {
	rv := objc.Send[AttributedString](u_.ID, objc.Sel("response"))
	return rv
}/* debug [instance_properties/getter]: response */


// The response with which the user responded to a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/response
func (u_ UserNotification) SetResponse(value IAttributedString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResponse:"), value)
}/* debug [instance_properties/setter]: response */


// Optional placeholder string for inline reply field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/responseplaceholder
func (u_ UserNotification) ResponsePlaceholder() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("responsePlaceholder"))
	return rv
}/* debug [instance_properties/getter]: responsePlaceholder */


// Optional placeholder string for inline reply field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/responseplaceholder
func (u_ UserNotification) SetResponsePlaceholder(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResponsePlaceholder:"), value)
}/* debug [instance_properties/setter]: responsePlaceholder */


// Specifies the name of the sound to play when the notification is delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/soundname
func (u_ UserNotification) SoundName() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("soundName"))
	return rv
}/* debug [instance_properties/getter]: soundName */


// Specifies the name of the sound to play when the notification is delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/soundname
func (u_ UserNotification) SetSoundName(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSoundName:"), value)
}/* debug [instance_properties/setter]: soundName */


// Specifies the subtitle of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/subtitle
func (u_ UserNotification) Subtitle() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("subtitle"))
	return rv
}/* debug [instance_properties/getter]: subtitle */


// Specifies the subtitle of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/subtitle
func (u_ UserNotification) SetSubtitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSubtitle:"), value)
}/* debug [instance_properties/setter]: subtitle */


// Specifies the title of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/title
func (u_ UserNotification) Title() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// Specifies the title of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/title
func (u_ UserNotification) SetTitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// Application-specific user info that can be attached to the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/userinfo
func (u_ UserNotification) UserInfo() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// Application-specific user info that can be attached to the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/userinfo
func (u_ UserNotification) SetUserInfo(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */


// The default notification sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationdefaultsoundname
func (u_ UserNotification) NSUserNotificationDefaultSoundName() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("NSUserNotificationDefaultSoundName"))
	return rv
}/* debug [instance_properties/getter]: NSUserNotificationDefaultSoundName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUserNotification */



