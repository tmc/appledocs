// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserUnixTask] class.
var UserUnixTaskClass objc.Class

func init() {
	UserUnixTaskClass = objc.GetClass("NSUserUnixTask")
}

type UserUnixTask struct {
	objc.ID
}

func UserUnixTaskFrom(ptr unsafe.Pointer) UserUnixTask {
	return UserUnixTask{
		ID: objc.ID(ptr),
	}
}



