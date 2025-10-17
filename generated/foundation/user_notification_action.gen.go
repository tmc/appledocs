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



