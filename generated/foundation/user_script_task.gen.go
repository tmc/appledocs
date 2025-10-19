// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserScriptTask] class.
var (
	userScriptTaskClass     _UserScriptTaskClass
	userScriptTaskClassOnce sync.Once
)

func getUserScriptTaskClass() _UserScriptTaskClass {
	userScriptTaskClassOnce.Do(func() {
		userScriptTaskClass = _UserScriptTaskClass{objc.GetClass("NSUserScriptTask")}
	})
	return userScriptTaskClass
}

type _UserScriptTaskClass struct {
	class objc.Class
}

// An interface definition for the [UserScriptTask] class.
type IUserScriptTask interface {
	objectivec.IObject
}

// An object that executes scripts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserScriptTask
type UserScriptTask struct {
	objectivec.Object
}

// UserScriptTaskFrom constructs a [UserScriptTask] from an unsafe.Pointer.
//
// An object that executes scripts.
func UserScriptTaskFrom(ptr unsafe.Pointer) UserScriptTask {
	return UserScriptTask{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UserScriptTaskClass) Alloc() UserScriptTask {
	rv := objc.Send[UserScriptTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserScriptTaskClass) New() UserScriptTask {
	rv := objc.Send[UserScriptTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserScriptTask) Init() UserScriptTask {
	rv := objc.Send[UserScriptTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserScriptTask) Autorelease() UserScriptTask {
	rv := objc.Send[UserScriptTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserScriptTask creates a new UserScriptTask instance.
func NewUserScriptTask() UserScriptTask {
	return getUserScriptTaskClass().New()
}




