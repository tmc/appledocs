// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLTask */


/* debug [class_header]: Header for MLTask */
// The class instance for the [Task] class.
var (
	TaskClass     _TaskClass
	TaskClassOnce sync.Once
)

func getTaskClass() _TaskClass {
	TaskClassOnce.Do(func() {
		TaskClass = _TaskClass{objc.GetClass("MLTask")}
	})
	return TaskClass
}

type _TaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Task */
// An interface definition for the [Task] class.
type ITask interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Task */
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	State() TaskState
	TaskIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Task */
	// methods:
	Cancel()
	Resume()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Task */
// Alloc allocates a new instance without initialization.
func (tc _TaskClass) Alloc() Task {
	rv := objc.Send[Task](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TaskClass) New() Task {
	rv := objc.Send[Task](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Task) Init() Task {
	rv := objc.Send[Task](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Task) Autorelease() Task {
	rv := objc.Send[Task](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTask creates a new Task instance.
func NewTask() Task {
	return getTaskClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Task */
// An abstract base class for machine learning tasks.
//
// You don’t create use this class directly. Instead, use a class that inherits from this one, such as .


// An abstract base class for machine learning tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTask
type Task struct {
	objectivec.Object
}

// TaskFrom constructs a [Task] from an unsafe.Pointer.
//
// An abstract base class for machine learning tasks.
func TaskFrom(ptr unsafe.Pointer) Task {
	return Task{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Task *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Task */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Task */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Task */

// Cancels a machine learning task before it completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTask/cancel()
func (t_ Task) Cancel() {
	objc.Send[objc.ID](t_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Begins or resumes a machine learning task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTask/resume()
func (t_ Task) Resume() {
	objc.Send[objc.ID](t_.ID, objc.Sel("resume"))
}/* debug [instance_methods/method]: Resume */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Task */

// The underlying error if the task is in a failed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTask/error
func (t_ Task) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](t_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The current state of the machine learning task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTask/state
func (t_ Task) State() TaskState {
	rv := objc.Send[TaskState](t_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// A unique name of the task to distinguish it from all other tasks at runtime.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTask/taskIdentifier
func (t_ Task) TaskIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("taskIdentifier"))
	return rv
}/* debug [instance_properties/getter]: taskIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLTask */



