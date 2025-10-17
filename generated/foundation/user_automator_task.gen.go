// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UserAutomatorTask] class.
var userAutomatorTaskClass = _UserAutomatorTaskClass{objc.GetClass("NSUserAutomatorTask")}

type _UserAutomatorTaskClass struct {
	class objc.Class
}

// An interface definition for the [UserAutomatorTask] class.
type IUserAutomatorTask interface {
	IUserScriptTask
}

// An object that executes Automator workflows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserAutomatorTask

type UserAutomatorTask struct {
	UserScriptTask
}

// UserAutomatorTaskFrom constructs a [UserAutomatorTask] from an unsafe.Pointer.
//
// An object that executes Automator workflows.
func UserAutomatorTaskFrom(ptr unsafe.Pointer) UserAutomatorTask {
	return UserAutomatorTask{
		UserScriptTask: UserScriptTaskFrom(ptr),
	}
}



