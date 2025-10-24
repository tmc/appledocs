// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKActivitySummaryType */


/* debug [class_header]: Header for HKActivitySummaryType */
// The class instance for the [HKActivitySummaryType] class.
var (
	HKActivitySummaryTypeClass     _HKActivitySummaryTypeClass
	HKActivitySummaryTypeClassOnce sync.Once
)

func getHKActivitySummaryTypeClass() _HKActivitySummaryTypeClass {
	HKActivitySummaryTypeClassOnce.Do(func() {
		HKActivitySummaryTypeClass = _HKActivitySummaryTypeClass{objc.GetClass("HKActivitySummaryType")}
	})
	return HKActivitySummaryTypeClass
}

type _HKActivitySummaryTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKActivitySummaryType */
// An interface definition for the [HKActivitySummaryType] class.
type IHKActivitySummaryType interface {
	IHKObjectType
	
/* debug [class_interface_properties]: Properties for HKActivitySummaryType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKActivitySummaryType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKActivitySummaryType */
// Alloc allocates a new instance without initialization.
func (hc _HKActivitySummaryTypeClass) Alloc() HKActivitySummaryType {
	rv := objc.Send[HKActivitySummaryType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKActivitySummaryTypeClass) New() HKActivitySummaryType {
	rv := objc.Send[HKActivitySummaryType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKActivitySummaryType) Init() HKActivitySummaryType {
	rv := objc.Send[HKActivitySummaryType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKActivitySummaryType) Autorelease() HKActivitySummaryType {
	rv := objc.Send[HKActivitySummaryType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKActivitySummaryType creates a new HKActivitySummaryType instance.
func NewHKActivitySummaryType() HKActivitySummaryType {
	return getHKActivitySummaryTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKActivitySummaryType */
// A type that identifies activity summary objects.
//
// Use the activity summary type to request permission to read objects from the HealthKit store. To create an activity summary type, use the class’s convenience method. The class is a concrete subclass of the class. Like many HealthKit classes, activity summary types aren’t extensible and you shouldn’t subclass them.


// A type that identifies activity summary objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummaryType
type HKActivitySummaryType struct {
	HKObjectType
}

// HKActivitySummaryTypeFrom constructs a [HKActivitySummaryType] from an unsafe.Pointer.
//
// A type that identifies activity summary objects.
func HKActivitySummaryTypeFrom(ptr unsafe.Pointer) HKActivitySummaryType {
	return HKActivitySummaryType{
		HKObjectType: HKObjectTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKActivitySummaryType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKActivitySummaryType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKActivitySummaryType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKActivitySummaryType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKActivitySummaryType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKActivitySummaryType */



