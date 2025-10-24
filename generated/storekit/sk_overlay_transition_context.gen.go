// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKOverlayTransitionContext */


/* debug [class_header]: Header for SKOverlayTransitionContext */
// The class instance for the [OverlayTransitionContext] class.
var (
	OverlayTransitionContextClass     _OverlayTransitionContextClass
	OverlayTransitionContextClassOnce sync.Once
)

func getOverlayTransitionContextClass() _OverlayTransitionContextClass {
	OverlayTransitionContextClassOnce.Do(func() {
		OverlayTransitionContextClass = _OverlayTransitionContextClass{objc.GetClass("SKOverlayTransitionContext")}
	})
	return OverlayTransitionContextClass
}

type _OverlayTransitionContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OverlayTransitionContext */
// An interface definition for the [OverlayTransitionContext] class.
type IOverlayTransitionContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OverlayTransitionContext */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OverlayTransitionContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OverlayTransitionContext */
// Alloc allocates a new instance without initialization.
func (oc _OverlayTransitionContextClass) Alloc() OverlayTransitionContext {
	rv := objc.Send[OverlayTransitionContext](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OverlayTransitionContextClass) New() OverlayTransitionContext {
	rv := objc.Send[OverlayTransitionContext](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OverlayTransitionContext) Init() OverlayTransitionContext {
	rv := objc.Send[OverlayTransitionContext](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OverlayTransitionContext) Autorelease() OverlayTransitionContext {
	rv := objc.Send[OverlayTransitionContext](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOverlayTransitionContext creates a new OverlayTransitionContext instance.
func NewOverlayTransitionContext() OverlayTransitionContext {
	return getOverlayTransitionContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OverlayTransitionContext */
// A context object you can use to animate UI changes while the platform presents or dismisses an overlay.
//
// For more information on animating UI changes while the system presents or dismisses an overlay, see and .


// A context object you can use to animate UI changes while the platform presents or dismisses an overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/TransitionContext
type OverlayTransitionContext struct {
	objectivec.Object
}

// OverlayTransitionContextFrom constructs a [OverlayTransitionContext] from an unsafe.Pointer.
//
// A context object you can use to animate UI changes while the platform presents or dismisses an overlay.
func OverlayTransitionContextFrom(ptr unsafe.Pointer) OverlayTransitionContext {
	return OverlayTransitionContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OverlayTransitionContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OverlayTransitionContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OverlayTransitionContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OverlayTransitionContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OverlayTransitionContext */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKOverlayTransitionContext */


