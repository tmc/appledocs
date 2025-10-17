// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserAutomatorTask] class.
var UserAutomatorTaskClass objc.Class

func init() {
	UserAutomatorTaskClass = objc.GetClass("NSUserAutomatorTask")
}

type UserAutomatorTask struct {
	objc.ID
}

func UserAutomatorTaskFrom(ptr unsafe.Pointer) UserAutomatorTask {
	return UserAutomatorTask{
		ID: objc.ID(ptr),
	}
}



