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
	UserScriptTaskClass     _UserScriptTaskClass
	UserScriptTaskClassOnce sync.Once
)

func getUserScriptTaskClass() _UserScriptTaskClass {
	UserScriptTaskClassOnce.Do(func() {
		UserScriptTaskClass = _UserScriptTaskClass{objc.GetClass("NSUserScriptTask")}
	})
	return UserScriptTaskClass
}

type _UserScriptTaskClass struct {
	class objc.Class
}

// An interface definition for the [UserScriptTask] class.
type IUserScriptTask interface {
	objectivec.IObject
	ExecuteWithCompletionHandler(handler unsafe.Pointer)
}

// An object that executes scripts.
//
// The class is able to run all the scripts normally run by the one of its subclasses, however it ignores the results. It is intended to execute user-supplied scripts and will execute them outside of the application’s sandbox, if any. If you need to execute scripts and get the input and output information use the , , and sub classes.
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




// Return a user script task instance given a URL for a script file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserScriptTask/init(url:)
func NewUserScriptTaskWithURLError(url IURL, error_ IError) UserScriptTask {
	instance := getUserScriptTaskClass().Alloc()
	rv := objc.Send[UserScriptTask](instance.ID, objc.Sel("initWithURL:error:"), url, error_)
	rv.Autorelease()
	return rv
}


// Executes the script with no input and ignoring any result.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserScriptTask/execute(completionHandler:)
func (u_ UserScriptTask) ExecuteWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("executeWithCompletionHandler:"), handler)
}

// The URL of the script file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserScriptTask/scriptURL
func (u_ UserScriptTask) ScriptURL() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("scriptURL"))
	return rv
}


