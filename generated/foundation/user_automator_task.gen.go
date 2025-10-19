// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UserAutomatorTask] class.
var (
	userAutomatorTaskClass     _UserAutomatorTaskClass
	userAutomatorTaskClassOnce sync.Once
)

func getUserAutomatorTaskClass() _UserAutomatorTaskClass {
	userAutomatorTaskClassOnce.Do(func() {
		userAutomatorTaskClass = _UserAutomatorTaskClass{objc.GetClass("NSUserAutomatorTask")}
	})
	return userAutomatorTaskClass
}

type _UserAutomatorTaskClass struct {
	class objc.Class
}

// An interface definition for the [UserAutomatorTask] class.
type IUserAutomatorTask interface {
	IUserScriptTask
}

// An object that executes Automator workflows.
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

// Alloc allocates a new instance without initialization.
func (uc _UserAutomatorTaskClass) Alloc() UserAutomatorTask {
	rv := objc.Send[UserAutomatorTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserAutomatorTaskClass) New() UserAutomatorTask {
	rv := objc.Send[UserAutomatorTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserAutomatorTask) Init() UserAutomatorTask {
	rv := objc.Send[UserAutomatorTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserAutomatorTask) Autorelease() UserAutomatorTask {
	rv := objc.Send[UserAutomatorTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserAutomatorTask creates a new UserAutomatorTask instance.
func NewUserAutomatorTask() UserAutomatorTask {
	return getUserAutomatorTaskClass().New()
}




