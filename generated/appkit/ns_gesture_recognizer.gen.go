// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSGestureRecognizer */


/* debug [class_header]: Header for NSGestureRecognizer */
// The class instance for the [GestureRecognizer] class.
var (
	GestureRecognizerClass     _GestureRecognizerClass
	GestureRecognizerClassOnce sync.Once
)

func getGestureRecognizerClass() _GestureRecognizerClass {
	GestureRecognizerClassOnce.Do(func() {
		GestureRecognizerClass = _GestureRecognizerClass{objc.GetClass("NSGestureRecognizer")}
	})
	return GestureRecognizerClass
}

type _GestureRecognizerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GestureRecognizer */
// An interface definition for the [GestureRecognizer] class.
type IGestureRecognizer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GestureRecognizer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GestureRecognizer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GestureRecognizer */
// Alloc allocates a new instance without initialization.
func (gc _GestureRecognizerClass) Alloc() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GestureRecognizerClass) New() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GestureRecognizer) Init() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GestureRecognizer) Autorelease() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGestureRecognizer creates a new GestureRecognizer instance.
func NewGestureRecognizer() GestureRecognizer {
	return getGestureRecognizerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GestureRecognizer */
// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type GestureRecognizer struct {
	objectivec.Object
}

// GestureRecognizerFrom constructs a [GestureRecognizer] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func GestureRecognizerFrom(ptr unsafe.Pointer) GestureRecognizer {
	return GestureRecognizer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GestureRecognizer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GestureRecognizer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GestureRecognizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GestureRecognizer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GestureRecognizer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSGestureRecognizer */



