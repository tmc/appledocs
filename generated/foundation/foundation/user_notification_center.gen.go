// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserNotificationCenter] class.
var UserNotificationCenterClass objc.Class

func init() {
	UserNotificationCenterClass = objc.GetClass("NSUserNotificationCenter")
}

type UserNotificationCenter struct {
	objc.ID
}

func UserNotificationCenterFrom(ptr unsafe.Pointer) UserNotificationCenter {
	return UserNotificationCenter{
		ID: objc.ID(ptr),
	}
}




