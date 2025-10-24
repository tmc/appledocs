// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCGameControllerActivationContext */


/* debug [class_header]: Header for GCGameControllerActivationContext */
// The class instance for the [GCGameControllerActivationContext] class.
var (
	GCGameControllerActivationContextClass     _GCGameControllerActivationContextClass
	GCGameControllerActivationContextClassOnce sync.Once
)

func getGCGameControllerActivationContextClass() _GCGameControllerActivationContextClass {
	GCGameControllerActivationContextClassOnce.Do(func() {
		GCGameControllerActivationContextClass = _GCGameControllerActivationContextClass{objc.GetClass("GCGameControllerActivationContext")}
	})
	return GCGameControllerActivationContextClass
}

type _GCGameControllerActivationContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCGameControllerActivationContext */
// An interface definition for the [GCGameControllerActivationContext] class.
type IGCGameControllerActivationContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCGameControllerActivationContext */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCGameControllerActivationContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCGameControllerActivationContext */
// Alloc allocates a new instance without initialization.
func (gc _GCGameControllerActivationContextClass) Alloc() GCGameControllerActivationContext {
	rv := objc.Send[GCGameControllerActivationContext](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCGameControllerActivationContextClass) New() GCGameControllerActivationContext {
	rv := objc.Send[GCGameControllerActivationContext](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCGameControllerActivationContext) Init() GCGameControllerActivationContext {
	rv := objc.Send[GCGameControllerActivationContext](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCGameControllerActivationContext) Autorelease() GCGameControllerActivationContext {
	rv := objc.Send[GCGameControllerActivationContext](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCGameControllerActivationContext creates a new GCGameControllerActivationContext instance.
func NewGCGameControllerActivationContext() GCGameControllerActivationContext {
	return getGCGameControllerActivationContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCGameControllerActivationContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGameControllerActivationContext
type GCGameControllerActivationContext struct {
	objectivec.Object
}

// GCGameControllerActivationContextFrom constructs a [GCGameControllerActivationContext] from an unsafe.Pointer.
func GCGameControllerActivationContextFrom(ptr unsafe.Pointer) GCGameControllerActivationContext {
	return GCGameControllerActivationContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCGameControllerActivationContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCGameControllerActivationContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCGameControllerActivationContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCGameControllerActivationContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCGameControllerActivationContext */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCGameControllerActivationContext */


