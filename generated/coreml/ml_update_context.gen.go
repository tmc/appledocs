// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLUpdateContext */


/* debug [class_header]: Header for MLUpdateContext */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UpdateContext */
// An interface definition for the [UpdateContext] class.
type IUpdateContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UpdateContext */
	// properties:
	Event() UpdateProgressEvent
	Metrics() foundation.IDictionary
	Model() unsafe.Pointer
	Parameters() foundation.IDictionary
	Task() IMLUpdateTask
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UpdateContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UpdateContext */
// Alloc allocates a new instance without initialization.
func (uc _UpdateContextClass) Alloc() UpdateContext {
	rv := objc.Send[UpdateContext](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UpdateContext */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UpdateContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UpdateContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UpdateContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UpdateContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UpdateContext */

// The event type that triggered an update task to notify your app’s completion and update progress handlers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateContext/event
func (u_ UpdateContext) Event() UpdateProgressEvent {
	rv := objc.Send[UpdateProgressEvent](u_.ID, objc.Sel("event"))
	return rv
}/* debug [instance_properties/getter]: event */


// The training metrics of the model for the update task, contained in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateContext/metrics
func (u_ UpdateContext) Metrics() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](u_.ID, objc.Sel("metrics"))
	return rv
}/* debug [instance_properties/getter]: metrics */


// The underlying Core ML model stored in memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateContext/model
func (u_ UpdateContext) Model() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("model"))
	return rv
}/* debug [instance_properties/getter]: model */


// The parameters for the update task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateContext/parameters
func (u_ UpdateContext) Parameters() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](u_.ID, objc.Sel("parameters"))
	return rv
}/* debug [instance_properties/getter]: parameters */


// The update task that generated the update context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateContext/task
func (u_ UpdateContext) Task() IMLUpdateTask {
	rv := objc.Send[UpdateTask](u_.ID, objc.Sel("task"))
	return rv
}/* debug [instance_properties/getter]: task */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLUpdateContext */



