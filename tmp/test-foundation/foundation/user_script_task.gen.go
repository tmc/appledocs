// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UserScriptTaskClass _UserScriptTaskClass

func init() {
	UserScriptTaskClass = _UserScriptTaskClass{objc.GetClass("NSUserScriptTask")}
}

type _UserScriptTaskClass struct {
	class objc.Class
}

type UserScriptTask struct {
	objc.ID
}

func UserScriptTaskFrom(ptr unsafe.Pointer) UserScriptTask {
	return UserScriptTask{
		ID: objc.ID(ptr),
	}
}




