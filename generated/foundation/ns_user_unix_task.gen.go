// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UserUnixTask] class.
var (
	UserUnixTaskClass     _UserUnixTaskClass
	UserUnixTaskClassOnce sync.Once
)

func getUserUnixTaskClass() _UserUnixTaskClass {
	UserUnixTaskClassOnce.Do(func() {
		UserUnixTaskClass = _UserUnixTaskClass{objc.GetClass("NSUserUnixTask")}
	})
	return UserUnixTaskClass
}

type _UserUnixTaskClass struct {
	class objc.Class
}

// An interface definition for the [UserUnixTask] class.
type IUserUnixTask interface {
	IUserScriptTask
	// properties:
	// methods:
}

// An object that executes unix applications.
//
// The class is intended to run unix applications, typically a shell script, from your application. It is intended to execute user-supplied scripts, and will execute them outside of the application’s sandbox, if any. The class is not intended to execute scripts built into an application; for that, use one of the , , or classes. If the application is sandboxed, then the script must be in the folder. A sandboxed application may read from, but not write to, this folder. If you simply need to execute unix scripts without regard to input or output, use , which can execute any of the specific types. If you need specific control over the input to, or output from, or the error stream of the script, use this class.


// An object that executes unix applications.
//
// [Full Topic]
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




