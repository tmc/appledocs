// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKCorrelationType */


/* debug [class_header]: Header for HKCorrelationType */
// The class instance for the [HKCorrelationType] class.
var (
	HKCorrelationTypeClass     _HKCorrelationTypeClass
	HKCorrelationTypeClassOnce sync.Once
)

func getHKCorrelationTypeClass() _HKCorrelationTypeClass {
	HKCorrelationTypeClassOnce.Do(func() {
		HKCorrelationTypeClass = _HKCorrelationTypeClass{objc.GetClass("HKCorrelationType")}
	})
	return HKCorrelationTypeClass
}

type _HKCorrelationTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKCorrelationType */
// An interface definition for the [HKCorrelationType] class.
type IHKCorrelationType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKCorrelationType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKCorrelationType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKCorrelationType */
// Alloc allocates a new instance without initialization.
func (hc _HKCorrelationTypeClass) Alloc() HKCorrelationType {
	rv := objc.Send[HKCorrelationType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKCorrelationTypeClass) New() HKCorrelationType {
	rv := objc.Send[HKCorrelationType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCorrelationType) Init() HKCorrelationType {
	rv := objc.Send[HKCorrelationType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCorrelationType) Autorelease() HKCorrelationType {
	rv := objc.Send[HKCorrelationType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCorrelationType creates a new HKCorrelationType instance.
func NewHKCorrelationType() HKCorrelationType {
	return getHKCorrelationTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKCorrelationType */
// A type that identifies samples that group multiple subsamples.
//
// The class is a concrete subclass of the class. To create a correlation type instance, use the object type’s conveniance method. Use correlation types to: Request permission to read or write matching quantity samples. Create and share matching quantity samples. Query for matching quantity samples. HealthKit provides two correlation types: blood pressure and food.


// A type that identifies samples that group multiple subsamples.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCorrelationType
type HKCorrelationType struct {
	HKSampleType
}

// HKCorrelationTypeFrom constructs a [HKCorrelationType] from an unsafe.Pointer.
//
// A type that identifies samples that group multiple subsamples.
func HKCorrelationTypeFrom(ptr unsafe.Pointer) HKCorrelationType {
	return HKCorrelationType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKCorrelationType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKCorrelationType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKCorrelationType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKCorrelationType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKCorrelationType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKCorrelationType */



