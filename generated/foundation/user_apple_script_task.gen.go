// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UserAppleScriptTask] class.
var userAppleScriptTaskClass = _UserAppleScriptTaskClass{objc.GetClass("NSUserAppleScriptTask")}

type _UserAppleScriptTaskClass struct {
	class objc.Class
}

// An object that executes AppleScript scripts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserAppleScriptTask

type UserAppleScriptTask struct {
	UserScriptTask
}

// UserAppleScriptTaskFrom constructs a [UserAppleScriptTask] from an unsafe.Pointer.
//
// An object that executes AppleScript scripts.
func UserAppleScriptTaskFrom(ptr unsafe.Pointer) UserAppleScriptTask {
	return UserAppleScriptTask{
		UserScriptTask: UserScriptTaskFrom(ptr),
	}
}



