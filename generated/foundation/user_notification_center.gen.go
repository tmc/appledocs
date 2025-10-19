// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserNotificationCenter] class.
var (
	userNotificationCenterClass     _UserNotificationCenterClass
	userNotificationCenterClassOnce sync.Once
)

func getUserNotificationCenterClass() _UserNotificationCenterClass {
	userNotificationCenterClassOnce.Do(func() {
		userNotificationCenterClass = _UserNotificationCenterClass{objc.GetClass("NSUserNotificationCenter")}
	})
	return userNotificationCenterClass
}

type _UserNotificationCenterClass struct {
	class objc.Class
}

// An interface definition for the [UserNotificationCenter] class.
type IUserNotificationCenter interface {
	objectivec.IObject
}

// An object that delivers notifications from apps to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotificationCenter
type UserNotificationCenter struct {
	objectivec.Object
}

// UserNotificationCenterFrom constructs a [UserNotificationCenter] from an unsafe.Pointer.
//
// An object that delivers notifications from apps to the user.
func UserNotificationCenterFrom(ptr unsafe.Pointer) UserNotificationCenter {
	return UserNotificationCenter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UserNotificationCenterClass) Alloc() UserNotificationCenter {
	rv := objc.Send[UserNotificationCenter](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserNotificationCenterClass) New() UserNotificationCenter {
	rv := objc.Send[UserNotificationCenter](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserNotificationCenter) Init() UserNotificationCenter {
	rv := objc.Send[UserNotificationCenter](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserNotificationCenter) Autorelease() UserNotificationCenter {
	rv := objc.Send[UserNotificationCenter](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserNotificationCenter creates a new UserNotificationCenter instance.
func NewUserNotificationCenter() UserNotificationCenter {
	return getUserNotificationCenterClass().New()
}




