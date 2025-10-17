// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserScriptTask] class.
var UserScriptTaskClass objc.Class

func init() {
	UserScriptTaskClass = objc.GetClass("NSUserScriptTask")
}

type UserScriptTask struct {
	objc.ID
}

func UserScriptTaskFrom(ptr unsafe.Pointer) UserScriptTask {
	return UserScriptTask{
		ID: objc.ID(ptr),
	}
}



