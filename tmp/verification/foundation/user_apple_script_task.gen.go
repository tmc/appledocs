// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var userAppleScriptTaskClass _UserAppleScriptTaskClass

func init() {
	userAppleScriptTaskClass = _UserAppleScriptTaskClass{objc.GetClass("NSUserAppleScriptTask")}
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




