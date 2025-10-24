// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLUpdateProgressHandlers */


/* debug [class_header]: Header for MLUpdateProgressHandlers */
// The class instance for the [UpdateProgressHandlers] class.
var (
	UpdateProgressHandlersClass     _UpdateProgressHandlersClass
	UpdateProgressHandlersClassOnce sync.Once
)

func getUpdateProgressHandlersClass() _UpdateProgressHandlersClass {
	UpdateProgressHandlersClassOnce.Do(func() {
		UpdateProgressHandlersClass = _UpdateProgressHandlersClass{objc.GetClass("MLUpdateProgressHandlers")}
	})
	return UpdateProgressHandlersClass
}

type _UpdateProgressHandlersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UpdateProgressHandlers */
// An interface definition for the [UpdateProgressHandlers] class.
type IUpdateProgressHandlers interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UpdateProgressHandlers */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UpdateProgressHandlers */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UpdateProgressHandlers */
// Alloc allocates a new instance without initialization.
func (uc _UpdateProgressHandlersClass) Alloc() UpdateProgressHandlers {
	rv := objc.Send[UpdateProgressHandlers](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UpdateProgressHandlersClass) New() UpdateProgressHandlers {
	rv := objc.Send[UpdateProgressHandlers](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UpdateProgressHandlers) Init() UpdateProgressHandlers {
	rv := objc.Send[UpdateProgressHandlers](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UpdateProgressHandlers) Autorelease() UpdateProgressHandlers {
	rv := objc.Send[UpdateProgressHandlers](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUpdateProgressHandlers creates a new UpdateProgressHandlers instance.
func NewUpdateProgressHandlers() UpdateProgressHandlers {
	return getUpdateProgressHandlersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UpdateProgressHandlers */
// A collection of closures an update task uses to notify your app of its progress.


// A collection of closures an update task uses to notify your app of its progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers
type UpdateProgressHandlers struct {
	objectivec.Object
}

// UpdateProgressHandlersFrom constructs a [UpdateProgressHandlers] from an unsafe.Pointer.
//
// A collection of closures an update task uses to notify your app of its progress.
func UpdateProgressHandlersFrom(ptr unsafe.Pointer) UpdateProgressHandlers {
	return UpdateProgressHandlers{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UpdateProgressHandlers */

// Creates the collection of closures an update task uses to notify your app of its progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlers/init(forEvents:progressHandler:completionHandler:)
func NewUpdateProgressHandlersForEventsProgressHandlerCompletionHandler(interestedEvents UpdateProgressEvent, progressHandler unsafe.Pointer, completionHandler unsafe.Pointer) UpdateProgressHandlers {
	instance := getUpdateProgressHandlersClass().Alloc()
	rv := objc.Send[UpdateProgressHandlers](instance.ID, objc.Sel("initForEvents:progressHandler:completionHandler:"), interestedEvents, progressHandler, completionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUpdateProgressHandlersForEventsProgressHandlerCompletionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UpdateProgressHandlers */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UpdateProgressHandlers */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UpdateProgressHandlers */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UpdateProgressHandlers */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLUpdateProgressHandlers */


