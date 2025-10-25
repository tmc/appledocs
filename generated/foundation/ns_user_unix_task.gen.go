// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUserUnixTask */


/* debug [class_header]: Header for NSUserUnixTask */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UserUnixTask */
// An interface definition for the [UserUnixTask] class.
type IUserUnixTask interface {
	IUserScriptTask
	
/* debug [class_interface_properties]: Properties for UserUnixTask */
	// properties:
	StandardError() IFileHandle
	SetStandardError(value IFileHandle)
	StandardInput() IFileHandle
	SetStandardInput(value IFileHandle)
	StandardOutput() IFileHandle
	SetStandardOutput(value IFileHandle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UserUnixTask */
	// methods:
	ExecuteWithArgumentsCompletionHandler(arguments []string, handler UserUnixTaskCompletionHandler /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UserUnixTask */
// Alloc allocates a new instance without initialization.
func (uc _UserUnixTaskClass) Alloc() UserUnixTask {
	rv := objc.Send[UserUnixTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UserUnixTask */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UserUnixTask *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UserUnixTask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UserUnixTask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UserUnixTask */

// Execute the unix script with the specified arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/execute(withArguments:completionHandler:)
func (u_ UserUnixTask) ExecuteWithArgumentsCompletionHandler(arguments []string, handler UserUnixTaskCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("executeWithArguments:completionHandler:"), arguments, handler)
}/* debug [instance_methods/method]: ExecuteWithArgumentsCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UserUnixTask */

// The standard error stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/standardError
func (u_ UserUnixTask) StandardError() IFileHandle {
	rv := objc.Send[FileHandle](u_.ID, objc.Sel("standardError"))
	return rv
}/* debug [instance_properties/getter]: standardError */


// The standard error stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/standardError
func (u_ UserUnixTask) SetStandardError(value IFileHandle) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setStandardError:"), value)
}/* debug [instance_properties/setter]: standardError */


// The standard input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/standardInput
func (u_ UserUnixTask) StandardInput() IFileHandle {
	rv := objc.Send[FileHandle](u_.ID, objc.Sel("standardInput"))
	return rv
}/* debug [instance_properties/getter]: standardInput */


// The standard input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/standardInput
func (u_ UserUnixTask) SetStandardInput(value IFileHandle) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setStandardInput:"), value)
}/* debug [instance_properties/setter]: standardInput */


// The standard output stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/standardOutput
func (u_ UserUnixTask) StandardOutput() IFileHandle {
	rv := objc.Send[FileHandle](u_.ID, objc.Sel("standardOutput"))
	return rv
}/* debug [instance_properties/getter]: standardOutput */


// The standard output stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/standardOutput
func (u_ UserUnixTask) SetStandardOutput(value IFileHandle) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setStandardOutput:"), value)
}/* debug [instance_properties/setter]: standardOutput */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUserUnixTask */



