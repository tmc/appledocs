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


// Specifies the title of the action button displayed in the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/actionButtonTitle
func (u_ UserNotification) ActionButtonTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("actionButtonTitle"))
	return rv
}


// SetActionButtonTitle sets the value of the actionButtonTitle property.
// Specifies the title of the action button displayed in the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/actionButtonTitle
func (u_ UserNotification) SetActionButtonTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionButtonTitle:"), value)
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
	objc.Send[objc.ID](u_.ID, objc.Sel("setAdditionalActions:"), value)
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
func (u_ UserNotification) InformativeText() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("informativeText"))
	return rv
}


// SetInformativeText sets the value of the informativeText property.
// The body text of the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/informativeText
func (u_ UserNotification) SetInformativeText(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInformativeText:"), value)
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
func (u_ UserNotification) OtherButtonTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("otherButtonTitle"))
	return rv
}


// SetOtherButtonTitle sets the value of the otherButtonTitle property.
// Specifies a custom title for the close button in an alert-style notification.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/otherButtonTitle
func (u_ UserNotification) SetOtherButtonTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOtherButtonTitle:"), value)
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
func (u_ UserNotification) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// Specifies the title of the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotification/title
func (u_ UserNotification) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), value)
}


