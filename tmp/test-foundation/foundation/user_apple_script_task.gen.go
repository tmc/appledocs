// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UserAppleScriptTaskClass _UserAppleScriptTaskClass

func init() {
	UserAppleScriptTaskClass = _UserAppleScriptTaskClass{objc.GetClass("NSUserAppleScriptTask")}
}

type _UserAppleScriptTaskClass struct {
	class objc.Class
}

type UserAppleScriptTask struct {
	objc.ID
}

func UserAppleScriptTaskFrom(ptr unsafe.Pointer) UserAppleScriptTask {
	return UserAppleScriptTask{
		ID: objc.ID(ptr),
	}
}




