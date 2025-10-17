// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UserUnixTask] class.
var userUnixTaskClass = _UserUnixTaskClass{objc.GetClass("NSUserUnixTask")}

type _UserUnixTaskClass struct {
	class objc.Class
}

// An object that executes unix applications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserUnixTask

type UserUnixTask struct {
	UserScriptTask
}

// UserUnixTaskFrom constructs a [UserUnixTask] from an unsafe.Pointer.
//
// An object that executes unix applications.
func UserUnixTaskFrom(ptr unsafe.Pointer) UserUnixTask {
	return UserUnixTask{
		UserScriptTask: UserScriptTaskFrom(ptr),
	}
}



