// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UserUnixTaskClass _UserUnixTaskClass

func init() {
	UserUnixTaskClass = _UserUnixTaskClass{objc.GetClass("NSUserUnixTask")}
}

type _UserUnixTaskClass struct {
	class objc.Class
}

type UserUnixTask struct {
	objc.ID
}

func UserUnixTaskFrom(ptr unsafe.Pointer) UserUnixTask {
	return UserUnixTask{
		ID: objc.ID(ptr),
	}
}




