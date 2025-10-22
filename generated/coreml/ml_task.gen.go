// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Task] class.
type ITask interface {
	objectivec.IObject
	Cancel()
	Resume()
	Error() foundation.Error
	State() TaskState
	TaskIdentifier() string
}

// An abstract base class for machine learning tasks.
//
// You don’t create use this class directly. Instead, use a class that inherits from this one, such as .
//
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

// Alloc allocates a new instance without initialization.
func (tc _TaskClass) Alloc() Task {
	rv := objc.Send[Task](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Cancels a machine learning task before it completes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTask/cancel()
func (t_ Task) Cancel() {
	objc.Send[objc.ID](t_.ID, objc.Sel("cancel"))
}

// Begins or resumes a machine learning task.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTask/resume()
func (t_ Task) Resume() {
	objc.Send[objc.ID](t_.ID, objc.Sel("resume"))
}

// The underlying error if the task is in a failed state.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTask/error
func (t_ Task) Error() foundation.Error {
	rv := objc.Send[foundation.Error](t_.ID, objc.Sel("error"))
	return rv
}

// The current state of the machine learning task.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTask/state
func (t_ Task) State() TaskState {
	rv := objc.Send[TaskState](t_.ID, objc.Sel("state"))
	return rv
}

// A unique name of the task to distinguish it from all other tasks at runtime.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLTask/taskIdentifier
func (t_ Task) TaskIdentifier() string {
	rv := objc.Send[string](t_.ID, objc.Sel("taskIdentifier"))
	return rv
}



