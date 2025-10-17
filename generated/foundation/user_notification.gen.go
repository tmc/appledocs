// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserNotification] class.
var userNotificationClass = _UserNotificationClass{objc.GetClass("NSUserNotification")}

type _UserNotificationClass struct {
	class objc.Class
}

// An interface definition for the [UserNotification] class.
type IUserNotification interface {
	objectivec.IObject
}

// A notification that can be scheduled for display in the notification center. [Full Topic]
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



