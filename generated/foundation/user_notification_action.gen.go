// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserNotificationAction] class.
var userNotificationActionClass = _UserNotificationActionClass{objc.GetClass("NSUserNotificationAction")}

type _UserNotificationActionClass struct {
	class objc.Class
}

// An interface definition for the [UserNotificationAction] class.
type IUserNotificationAction interface {
	objectivec.IObject
}

// An action that the user can take in response to receiving a notification. [Full Topic]
//
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
// Alloc allocates a new instance without initialization.
func (uc _UserNotificationActionClass) Alloc() UserNotificationAction {
	rv := objc.Send[UserNotificationAction](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return userNotificationActionClass.New()
}




