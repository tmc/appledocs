// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserNotification] class.
var UserNotificationClass objc.Class

func init() {
	UserNotificationClass = objc.GetClass("NSUserNotification")
}

type UserNotification struct {
	objc.ID
}

func UserNotificationFrom(ptr unsafe.Pointer) UserNotification {
	return UserNotification{
		ID: objc.ID(ptr),
	}
}



