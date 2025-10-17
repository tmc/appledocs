// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserAppleScriptTask] class.
var UserAppleScriptTaskClass objc.Class

func init() {
	UserAppleScriptTaskClass = objc.GetClass("NSUserAppleScriptTask")
}

type UserAppleScriptTask struct {
	objc.ID
}

func UserAppleScriptTaskFrom(ptr unsafe.Pointer) UserAppleScriptTask {
	return UserAppleScriptTask{
		ID: objc.ID(ptr),
	}
}




