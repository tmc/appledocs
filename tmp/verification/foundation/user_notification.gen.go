// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var userNotificationClass _UserNotificationClass

func init() {
	userNotificationClass = _UserNotificationClass{objc.GetClass("NSUserNotification")}
}

type _UserNotificationClass struct {
	class objc.Class
}

type UserNotification struct {
	objc.ID
}

func UserNotificationFrom(ptr unsafe.Pointer) UserNotification {
	return UserNotification{
		ID: objc.ID(ptr),
	}
}




