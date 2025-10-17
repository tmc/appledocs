// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserNotificationAction] class.
var UserNotificationActionClass objc.Class

func init() {
	UserNotificationActionClass = objc.GetClass("NSUserNotificationAction")
}

type UserNotificationAction struct {
	objc.ID
}

func UserNotificationActionFrom(ptr unsafe.Pointer) UserNotificationAction {
	return UserNotificationAction{
		ID: objc.ID(ptr),
	}
}



