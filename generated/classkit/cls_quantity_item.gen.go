// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CLSQuantityItem */


/* debug [class_header]: Header for CLSQuantityItem */
// The class instance for the [SQuantityItem] class.
var (
	SQuantityItemClass     _SQuantityItemClass
	SQuantityItemClassOnce sync.Once
)

func getSQuantityItemClass() _SQuantityItemClass {
	SQuantityItemClassOnce.Do(func() {
		SQuantityItemClass = _SQuantityItemClass{objc.GetClass("CLSQuantityItem")}
	})
	return SQuantityItemClass
}

type _SQuantityItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SQuantityItem */
// An interface definition for the [SQuantityItem] class.
type ISQuantityItem interface {
	ISActivityItem
	
/* debug [class_interface_properties]: Properties for SQuantityItem */
	// properties:
	Quantity() float64
	SetQuantity(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SQuantityItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SQuantityItem */
// Alloc allocates a new instance without initialization.
func (sc _SQuantityItemClass) Alloc() SQuantityItem {
	rv := objc.Send[SQuantityItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SQuantityItemClass) New() SQuantityItem {
	rv := objc.Send[SQuantityItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SQuantityItem) Init() SQuantityItem {
	rv := objc.Send[SQuantityItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SQuantityItem) Autorelease() SQuantityItem {
	rv := objc.Send[SQuantityItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSQuantityItem creates a new SQuantityItem instance.
func NewSQuantityItem() SQuantityItem {
	return getSQuantityItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SQuantityItem */
// Activity information that signifies a quantity.
//
// Use an activity item of this type to associate a discrete value with a task. For example, you might use it to indicate how many times the user requested a hint while taking a quiz.


// Activity information that signifies a quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSQuantityItem
type SQuantityItem struct {
	SActivityItem
}

// SQuantityItemFrom constructs a [SQuantityItem] from an unsafe.Pointer.
//
// Activity information that signifies a quantity.
func SQuantityItemFrom(ptr unsafe.Pointer) SQuantityItem {
	return SQuantityItem{
		SActivityItem: SActivityItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SQuantityItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SQuantityItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SQuantityItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SQuantityItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SQuantityItem */

// A quantity associated with the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsquantityitem/quantity
func (s_ SQuantityItem) Quantity() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("quantity"))
	return rv
}/* debug [instance_properties/getter]: quantity */


// A quantity associated with the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsquantityitem/quantity
func (s_ SQuantityItem) SetQuantity(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setQuantity:"), value)
}/* debug [instance_properties/setter]: quantity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CLSQuantityItem */



