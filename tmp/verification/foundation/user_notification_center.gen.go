// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var userNotificationCenterClass _UserNotificationCenterClass

func init() {
	userNotificationCenterClass = _UserNotificationCenterClass{objc.GetClass("NSUserNotificationCenter")}
}

type _UserNotificationCenterClass struct {
	class objc.Class
}

type UserNotificationCenter struct {
	objc.ID
}

func UserNotificationCenterFrom(ptr unsafe.Pointer) UserNotificationCenter {
	return UserNotificationCenter{
		ID: objc.ID(ptr),
	}
}




