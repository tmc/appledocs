// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UserAppleScriptTask] class.
var (
	UserAppleScriptTaskClass     _UserAppleScriptTaskClass
	UserAppleScriptTaskClassOnce sync.Once
)

func getUserAppleScriptTaskClass() _UserAppleScriptTaskClass {
	UserAppleScriptTaskClassOnce.Do(func() {
		UserAppleScriptTaskClass = _UserAppleScriptTaskClass{objc.GetClass("NSUserAppleScriptTask")}
	})
	return UserAppleScriptTaskClass
}

type _UserAppleScriptTaskClass struct {
	class objc.Class
}

// An interface definition for the [UserAppleScriptTask] class.
type IUserAppleScriptTask interface {
	IUserScriptTask
	ExecuteWithAppleEventCompletionHandler(event IAppleEventDescriptor, handler unsafe.Pointer)
}

// An object that executes AppleScript scripts.
//
// The class is intended to run AppleScript scripts from your application. It is intended to execute user-supplied scripts and will execute them outside of the application’s sandbox, if any. The class is not intended to execute scripts built into an application; for that, use one of the classes. If the application is sandboxed, then the script must be in the folder. A sandboxed application may read from, but not write to, this folder. If you simply need to execute scripts without regard to input or output, use , which can execute any of the specific types. If you need specific control over the input to or output from the script, use this class.
//
// [Full Topic]
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


// Execute the AppleScript script by sending it the specified Apple event.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserAppleScriptTask/execute(withAppleEvent:completionHandler:)
func (u_ UserAppleScriptTask) ExecuteWithAppleEventCompletionHandler(event IAppleEventDescriptor, handler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("executeWithAppleEvent:completionHandler:"), event, handler)
}



