// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UserNotificationAction] class.
var UserNotificationActionClass = _UserNotificationActionClass{objc.GetClass("NSUserNotificationAction")}

type _UserNotificationActionClass struct {
	class objc.Class
}

type UserNotificationAction struct {
	objc.ID
}

func UserNotificationActionFrom(ptr unsafe.Pointer) UserNotificationAction {
	return UserNotificationAction{
		ID: objc.ID(ptr),
	}
}




