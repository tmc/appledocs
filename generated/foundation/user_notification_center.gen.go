// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserNotificationCenter] class.
var userNotificationCenterClass = _UserNotificationCenterClass{objc.GetClass("NSUserNotificationCenter")}

type _UserNotificationCenterClass struct {
	class objc.Class
}

// An object that delivers notifications from apps to the user. [Full Topic]
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



