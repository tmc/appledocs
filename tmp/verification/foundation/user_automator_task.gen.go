// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var userAutomatorTaskClass _UserAutomatorTaskClass

func init() {
	userAutomatorTaskClass = _UserAutomatorTaskClass{objc.GetClass("NSUserAutomatorTask")}
}

type _UserAutomatorTaskClass struct {
	class objc.Class
}

type UserAutomatorTask struct {
	objc.ID
}

func UserAutomatorTaskFrom(ptr unsafe.Pointer) UserAutomatorTask {
	return UserAutomatorTask{
		ID: objc.ID(ptr),
	}
}




