// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCGearShifterElement */


/* debug [class_header]: Header for GCGearShifterElement */
// The class instance for the [GCGearShifterElement] class.
var (
	GCGearShifterElementClass     _GCGearShifterElementClass
	GCGearShifterElementClassOnce sync.Once
)

func getGCGearShifterElementClass() _GCGearShifterElementClass {
	GCGearShifterElementClassOnce.Do(func() {
		GCGearShifterElementClass = _GCGearShifterElementClass{objc.GetClass("GCGearShifterElement")}
	})
	return GCGearShifterElementClass
}

type _GCGearShifterElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCGearShifterElement */
// An interface definition for the [GCGearShifterElement] class.
type IGCGearShifterElement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCGearShifterElement */
	// properties:
	PatternInput() unsafe.Pointer
	SequentialInput() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCGearShifterElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCGearShifterElement */
// Alloc allocates a new instance without initialization.
func (gc _GCGearShifterElementClass) Alloc() GCGearShifterElement {
	rv := objc.Send[GCGearShifterElement](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCGearShifterElementClass) New() GCGearShifterElement {
	rv := objc.Send[GCGearShifterElement](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCGearShifterElement) Init() GCGearShifterElement {
	rv := objc.Send[GCGearShifterElement](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCGearShifterElement) Autorelease() GCGearShifterElement {
	rv := objc.Send[GCGearShifterElement](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCGearShifterElement creates a new GCGearShifterElement instance.
func NewGCGearShifterElement() GCGearShifterElement {
	return getGCGearShifterElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCGearShifterElement */
// An element that represents either a pattern or a sequential gear shift.


// An element that represents either a pattern or a sequential gear shift.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGearShifterElement
type GCGearShifterElement struct {
	objectivec.Object
}

// GCGearShifterElementFrom constructs a [GCGearShifterElement] from an unsafe.Pointer.
//
// An element that represents either a pattern or a sequential gear shift.
func GCGearShifterElementFrom(ptr unsafe.Pointer) GCGearShifterElement {
	return GCGearShifterElement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCGearShifterElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCGearShifterElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCGearShifterElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCGearShifterElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCGearShifterElement */

// The input object for a pattern gear shift.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGearShifterElement/patternInput
func (g_ GCGearShifterElement) PatternInput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("patternInput"))
	return rv
}/* debug [instance_properties/getter]: patternInput */


// The input object for a sequential gear shift.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGearShifterElement/sequentialInput
func (g_ GCGearShifterElement) SequentialInput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("sequentialInput"))
	return rv
}/* debug [instance_properties/getter]: sequentialInput */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCGearShifterElement */



