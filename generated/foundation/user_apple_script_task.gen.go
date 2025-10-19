// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UserAppleScriptTask] class.
var (
	userAppleScriptTaskClass     _UserAppleScriptTaskClass
	userAppleScriptTaskClassOnce sync.Once
)

func getUserAppleScriptTaskClass() _UserAppleScriptTaskClass {
	userAppleScriptTaskClassOnce.Do(func() {
		userAppleScriptTaskClass = _UserAppleScriptTaskClass{objc.GetClass("NSUserAppleScriptTask")}
	})
	return userAppleScriptTaskClass
}

type _UserAppleScriptTaskClass struct {
	class objc.Class
}

// An interface definition for the [UserAppleScriptTask] class.
type IUserAppleScriptTask interface {
	IUserScriptTask
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

// Alloc allocates a new instance without initialization.
func (uc _UserAppleScriptTaskClass) Alloc() UserAppleScriptTask {
	rv := objc.Send[UserAppleScriptTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserAppleScriptTaskClass) New() UserAppleScriptTask {
	rv := objc.Send[UserAppleScriptTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserAppleScriptTask) Init() UserAppleScriptTask {
	rv := objc.Send[UserAppleScriptTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserAppleScriptTask) Autorelease() UserAppleScriptTask {
	rv := objc.Send[UserAppleScriptTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserAppleScriptTask creates a new UserAppleScriptTask instance.
func NewUserAppleScriptTask() UserAppleScriptTask {
	return getUserAppleScriptTaskClass().New()
}




