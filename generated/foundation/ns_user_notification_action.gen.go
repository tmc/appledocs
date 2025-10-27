// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [UserNotificationAction] class.
var (
	UserNotificationActionClass     _UserNotificationActionClass
	UserNotificationActionClassOnce sync.Once
)

func getUserNotificationActionClass() _UserNotificationActionClass {
	UserNotificationActionClassOnce.Do(func() {
		UserNotificationActionClass = _UserNotificationActionClass{objc.GetClass("NSUserNotificationAction")}
	})
	return UserNotificationActionClass
}

type _UserNotificationActionClass struct {
	class objc.Class
}





// An interface definition for the [UserNotificationAction] class.
type IUserNotificationAction interface {
	objectivec.IObject
	

	// properties:
	AdditionalActions() IUserNotificationAction
	SetAdditionalActions(value IUserNotificationAction)
	AdditionalActivationAction() IUserNotificationAction
	SetAdditionalActivationAction(value IUserNotificationAction)
	Identifier() IString
	SetIdentifier(value IString)
	Title() IString
	SetTitle(value IString)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (uc _UserNotificationActionClass) Alloc() UserNotificationAction {
	rv := objc.Send[UserNotificationAction](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UserNotificationActionClass) New() UserNotificationAction {
	rv := objc.Send[UserNotificationAction](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserNotificationAction) Init() UserNotificationAction {
	rv := objc.Send[UserNotificationAction](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserNotificationAction) Autorelease() UserNotificationAction {
	rv := objc.Send[UserNotificationAction](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserNotificationAction creates a new UserNotificationAction instance.
func NewUserNotificationAction() UserNotificationAction {
	return getUserNotificationActionClass().New()
}





// An action that the user can take in response to receiving a notification.
//
// User notifications can specify one or more actions to show to the user by using the or properties. objects contain the localized title shown to the user and an identifier used to differentiate between presented actions.


// An action that the user can take in response to receiving a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotificationAction
type UserNotificationAction struct {
	objectivec.Object
}

// UserNotificationActionFrom constructs a [UserNotificationAction] from an unsafe.Pointer.
//
// An action that the user can take in response to receiving a notification.
func UserNotificationActionFrom(ptr unsafe.Pointer) UserNotificationAction {
	return UserNotificationAction{objectivec.Object{objc.ID(ptr)}}
}

























// The actions that can be taken on a notification in addition to the default action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/additionalactions
func (u_ UserNotificationAction) AdditionalActions() IUserNotificationAction {
	rv := objc.Send[UserNotificationAction](u_.ID, objc.Sel("additionalActions"))
	return rv
}


// The actions that can be taken on a notification in addition to the default action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/additionalactions
func (u_ UserNotificationAction) SetAdditionalActions(value IUserNotificationAction) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAdditionalActions:"), value)
}


// An additional action selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/additionalactivationaction
func (u_ UserNotificationAction) AdditionalActivationAction() IUserNotificationAction {
	rv := objc.Send[UserNotificationAction](u_.ID, objc.Sel("additionalActivationAction"))
	return rv
}


// An additional action selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/additionalactivationaction
func (u_ UserNotificationAction) SetAdditionalActivationAction(value IUserNotificationAction) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAdditionalActivationAction:"), value)
}


// The identifier for the user notification action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationaction/identifier
func (u_ UserNotificationAction) Identifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("identifier"))
	return rv
}


// The identifier for the user notification action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationaction/identifier
func (u_ UserNotificationAction) SetIdentifier(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIdentifier:"), value)
}


// The localized title shown to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationaction/title
func (u_ UserNotificationAction) Title() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("title"))
	return rv
}


// The localized title shown to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationaction/title
func (u_ UserNotificationAction) SetTitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), value)
}








