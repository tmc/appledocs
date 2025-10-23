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
	// properties:
	Event() UpdateProgressEvent /* not a class type */
	SetEvent(value UpdateProgressEvent /* not a class type */)
	Metrics() objc.IObject /* cross-framework: MetricKey */
	SetMetrics(value objc.IObject /* cross-framework: MetricKey */)
	Model() Writable /* not a class type */
	SetModel(value Writable /* not a class type */)
	Parameters() IMLParameterKey
	SetParameters(value IMLParameterKey)
	Task() IMLUpdateTask
	SetTask(value IMLUpdateTask)
	// methods:
}

// The context an update task provides to your app’s completion and update progress handlers.


// The context an update task provides to your app’s completion and update progress handlers.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/event
func (u_ UpdateContext) Event() UpdateProgressEvent /* not a class type */ {
	rv := objc.Send[UpdateProgressEvent](u_.ID, objc.Sel("event"))
	return rv
}


// The event type that triggered an update task to notify your app’s completion and update progress handlers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/event
func (u_ UpdateContext) SetEvent(value UpdateProgressEvent /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEvent:"), value)
}


// The training metrics of the model for the update task, contained in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/metrics
func (u_ UpdateContext) Metrics() objc.IObject /* cross-framework: MetricKey */ {
	rv := objc.Send[MetricKey](u_.ID, objc.Sel("metrics"))
	return rv
}


// The training metrics of the model for the update task, contained in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/metrics
func (u_ UpdateContext) SetMetrics(value objc.IObject /* cross-framework: MetricKey */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMetrics:"), value)
}


// The underlying Core ML model stored in memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/model
func (u_ UpdateContext) Model() Writable /* not a class type */ {
	rv := objc.Send[Writable](u_.ID, objc.Sel("model"))
	return rv
}


// The underlying Core ML model stored in memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/model
func (u_ UpdateContext) SetModel(value Writable /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setModel:"), value)
}


// The parameters for the update task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/parameters
func (u_ UpdateContext) Parameters() IMLParameterKey {
	rv := objc.Send[ParameterKey](u_.ID, objc.Sel("parameters"))
	return rv
}


// The parameters for the update task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/parameters
func (u_ UpdateContext) SetParameters(value IMLParameterKey) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setParameters:"), value)
}


// The update task that generated the update context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/task
func (u_ UpdateContext) Task() IMLUpdateTask {
	rv := objc.Send[UpdateTask](u_.ID, objc.Sel("task"))
	return rv
}


// The update task that generated the update context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/task
func (u_ UpdateContext) SetTask(value IMLUpdateTask) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTask:"), value)
}



