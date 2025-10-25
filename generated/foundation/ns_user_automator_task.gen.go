// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUserAutomatorTask */


/* debug [class_header]: Header for NSUserAutomatorTask */
// The class instance for the [UserAutomatorTask] class.
var (
	UserAutomatorTaskClass     _UserAutomatorTaskClass
	UserAutomatorTaskClassOnce sync.Once
)

func getUserAutomatorTaskClass() _UserAutomatorTaskClass {
	UserAutomatorTaskClassOnce.Do(func() {
		UserAutomatorTaskClass = _UserAutomatorTaskClass{objc.GetClass("NSUserAutomatorTask")}
	})
	return UserAutomatorTaskClass
}

type _UserAutomatorTaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UserAutomatorTask */
// An interface definition for the [UserAutomatorTask] class.
type IUserAutomatorTask interface {
	IUserScriptTask
	
/* debug [class_interface_properties]: Properties for UserAutomatorTask */
	// properties:
	Variables() IDictionary
	SetVariables(value IDictionary)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UserAutomatorTask */
	// methods:
	ExecuteWithInputCompletionHandler(input unsafe.Pointer, handler UserAutomatorTaskCompletionHandler /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UserAutomatorTask */
// Alloc allocates a new instance without initialization.
func (uc _UserAutomatorTaskClass) Alloc() UserAutomatorTask {
	rv := objc.Send[UserAutomatorTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UserAutomatorTask */
// An object that executes Automator workflows.
//
// The class is intended to run Automator workflows from your application. It is intended to execute user-supplied workflows, and will execute them outside of the application’s sandbox, if any. The class is not intended to execute scripts built into an application; for that, use one of the or classes. If the application is sandboxed, then the script must be in the folder. A sandboxed application may read from, but not write to, this folder. If you simply need to execute scripts without regard to input or output, use , which can execute any of the specific types. If you need specific control over the input to or output from the workflow, use this class.


// An object that executes Automator workflows.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UserAutomatorTask *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UserAutomatorTask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UserAutomatorTask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UserAutomatorTask */

// Execute the Automator workflow by providing it as securely coded input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserAutomatorTask/execute(withInput:completionHandler:)
func (u_ UserAutomatorTask) ExecuteWithInputCompletionHandler(input unsafe.Pointer, handler UserAutomatorTaskCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("executeWithInput:completionHandler:"), input, handler)
}/* debug [instance_methods/method]: ExecuteWithInputCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UserAutomatorTask */

// The variables required by the Automator workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserAutomatorTask/variables
func (u_ UserAutomatorTask) Variables() IDictionary {
	rv := objc.Send[Dictionary](u_.ID, objc.Sel("variables"))
	return rv
}/* debug [instance_properties/getter]: variables */


// The variables required by the Automator workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserAutomatorTask/variables
func (u_ UserAutomatorTask) SetVariables(value IDictionary) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setVariables:"), value)
}/* debug [instance_properties/setter]: variables */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUserAutomatorTask */



