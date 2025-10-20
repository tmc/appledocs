// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UpdateContext] class.
var (
	UpdateContextClass     _UpdateContextClass
	UpdateContextClassOnce sync.Once
)

func getUpdateContextClass() _UpdateContextClass {
	UpdateContextClassOnce.Do(func() {
		UpdateContextClass = _UpdateContextClass{objc.GetClass("MLUpdateContext")}
	})
	return UpdateContextClass
}

type _UpdateContextClass struct {
	class objc.Class
}

// An interface definition for the [UpdateContext] class.
type IUpdateContext interface {
	objectivec.IObject
}

// The context an update task provides to your app’s completion and update progress handlers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateContext
type UpdateContext struct {
	objectivec.Object
}

// UpdateContextFrom constructs a [UpdateContext] from an unsafe.Pointer.
//
// The context an update task provides to your app’s completion and update progress handlers.
func UpdateContextFrom(ptr unsafe.Pointer) UpdateContext {
	return UpdateContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UpdateContextClass) Alloc() UpdateContext {
	rv := objc.Send[UpdateContext](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UpdateContextClass) New() UpdateContext {
	rv := objc.Send[UpdateContext](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UpdateContext) Init() UpdateContext {
	rv := objc.Send[UpdateContext](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UpdateContext) Autorelease() UpdateContext {
	rv := objc.Send[UpdateContext](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUpdateContext creates a new UpdateContext instance.
func NewUpdateContext() UpdateContext {
	return getUpdateContextClass().New()
}


// The event type that triggered an update task to notify your app’s completion and update progress handlers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateContext/event
func (u_ UpdateContext) Event() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("event"))
	return rv
}

// The training metrics of the model for the update task, contained in a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateContext/metrics
func (u_ UpdateContext) Metrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("metrics"))
	return rv
}

// The underlying Core ML model stored in memory.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateContext/model
func (u_ UpdateContext) Model() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("model"))
	return rv
}

// The parameters for the update task.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateContext/parameters
func (u_ UpdateContext) Parameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("parameters"))
	return rv
}

// The update task that generated the update context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateContext/task
func (u_ UpdateContext) Task() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("task"))
	return rv
}



