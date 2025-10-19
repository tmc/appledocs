// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UserUnixTask] class.
var (
	userUnixTaskClass     _UserUnixTaskClass
	userUnixTaskClassOnce sync.Once
)

func getUserUnixTaskClass() _UserUnixTaskClass {
	userUnixTaskClassOnce.Do(func() {
		userUnixTaskClass = _UserUnixTaskClass{objc.GetClass("NSUserUnixTask")}
	})
	return userUnixTaskClass
}

type _UserUnixTaskClass struct {
	class objc.Class
}

// An interface definition for the [UserUnixTask] class.
type IUserUnixTask interface {
	IUserScriptTask
}

// An object that executes unix applications.
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

// Alloc allocates a new instance without initialization.
func (uc _UserUnixTaskClass) Alloc() UserUnixTask {
	rv := objc.Send[UserUnixTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserUnixTaskClass) New() UserUnixTask {
	rv := objc.Send[UserUnixTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserUnixTask) Init() UserUnixTask {
	rv := objc.Send[UserUnixTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserUnixTask) Autorelease() UserUnixTask {
	rv := objc.Send[UserUnixTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserUnixTask creates a new UserUnixTask instance.
func NewUserUnixTask() UserUnixTask {
	return getUserUnixTaskClass().New()
}




