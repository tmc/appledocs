// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UpdateTask] class.
var (
	UpdateTaskClass     _UpdateTaskClass
	UpdateTaskClassOnce sync.Once
)

func getUpdateTaskClass() _UpdateTaskClass {
	UpdateTaskClassOnce.Do(func() {
		UpdateTaskClass = _UpdateTaskClass{objc.GetClass("MLUpdateTask")}
	})
	return UpdateTaskClass
}

type _UpdateTaskClass struct {
	class objc.Class
}

// An interface definition for the [UpdateTask] class.
type IUpdateTask interface {
	ITask
}

// A task that updates a model with additional training data.
//
// Use an to update a machine learning model on a user’s device.


// A task that updates a model with additional training data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask
type UpdateTask struct {
	Task
}

// UpdateTaskFrom constructs a [UpdateTask] from an unsafe.Pointer.
//
// A task that updates a model with additional training data.
func UpdateTaskFrom(ptr unsafe.Pointer) UpdateTask {
	return UpdateTask{
		Task: TaskFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UpdateTaskClass) Alloc() UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UpdateTaskClass) New() UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UpdateTask) Init() UpdateTask {
	rv := objc.Send[UpdateTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UpdateTask) Autorelease() UpdateTask {
	rv := objc.Send[UpdateTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUpdateTask creates a new UpdateTask instance.
func NewUpdateTask() UpdateTask {
	return getUpdateTaskClass().New()
}





