// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSUserScriptTask */


/* debug [class_header]: Header for NSUserScriptTask */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UserScriptTask */
// An interface definition for the [UserScriptTask] class.
type IUserScriptTask interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UserScriptTask */
	// properties:
	ScriptURL() IURL
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UserScriptTask */
	// methods:
	ExecuteWithCompletionHandler(handler UserScriptTaskCompletionHandler /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UserScriptTask */
// Alloc allocates a new instance without initialization.
func (uc _UserScriptTaskClass) Alloc() UserScriptTask {
	rv := objc.Send[UserScriptTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UserScriptTask */
// An object that executes scripts.
//
// The class is able to run all the scripts normally run by the one of its subclasses, however it ignores the results. It is intended to execute user-supplied scripts and will execute them outside of the application’s sandbox, if any. If you need to execute scripts and get the input and output information use the , , and sub classes.


// An object that executes scripts.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UserScriptTask */

// Return a user script task instance given a URL for a script file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserScriptTask/init(url:)
func NewUserScriptTaskWithURLError(url IURL, error_ IError) UserScriptTask {
	instance := getUserScriptTaskClass().Alloc()
	rv := objc.Send[UserScriptTask](instance.ID, objc.Sel("initWithURL:error:"), url, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUserScriptTaskWithURLError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UserScriptTask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UserScriptTask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UserScriptTask */

// Executes the script with no input and ignoring any result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserScriptTask/execute(completionHandler:)
func (u_ UserScriptTask) ExecuteWithCompletionHandler(handler UserScriptTaskCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("executeWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: ExecuteWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UserScriptTask */

// The URL of the script file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserScriptTask/scriptURL
func (u_ UserScriptTask) ScriptURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("scriptURL"))
	return rv
}/* debug [instance_properties/getter]: scriptURL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUserScriptTask */


