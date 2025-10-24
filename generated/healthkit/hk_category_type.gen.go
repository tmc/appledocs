// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKCategoryType */


/* debug [class_header]: Header for HKCategoryType */
// The class instance for the [HKCategoryType] class.
var (
	HKCategoryTypeClass     _HKCategoryTypeClass
	HKCategoryTypeClassOnce sync.Once
)

func getHKCategoryTypeClass() _HKCategoryTypeClass {
	HKCategoryTypeClassOnce.Do(func() {
		HKCategoryTypeClass = _HKCategoryTypeClass{objc.GetClass("HKCategoryType")}
	})
	return HKCategoryTypeClass
}

type _HKCategoryTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKCategoryType */
// An interface definition for the [HKCategoryType] class.
type IHKCategoryType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKCategoryType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKCategoryType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKCategoryType */
// Alloc allocates a new instance without initialization.
func (hc _HKCategoryTypeClass) Alloc() HKCategoryType {
	rv := objc.Send[HKCategoryType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKCategoryTypeClass) New() HKCategoryType {
	rv := objc.Send[HKCategoryType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCategoryType) Init() HKCategoryType {
	rv := objc.Send[HKCategoryType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCategoryType) Autorelease() HKCategoryType {
	rv := objc.Send[HKCategoryType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCategoryType creates a new HKCategoryType instance.
func NewHKCategoryType() HKCategoryType {
	return getHKCategoryTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKCategoryType */
// A type that identifies samples that contain a value from a small set of possible values.
//
// The class is a concrete subclass of the HKObjectType class. To create a category type instance, use the convenience method. For example, the following code creates a category sample type for handwashing events. Use category types to: Request permission to read or write matching category samples. Create and share matching category samples. Query for matching category samples. For a complete list of category types, refer to .


// A type that identifies samples that contain a value from a small set of possible values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryType
type HKCategoryType struct {
	HKSampleType
}

// HKCategoryTypeFrom constructs a [HKCategoryType] from an unsafe.Pointer.
//
// A type that identifies samples that contain a value from a small set of possible values.
func HKCategoryTypeFrom(ptr unsafe.Pointer) HKCategoryType {
	return HKCategoryType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKCategoryType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKCategoryType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKCategoryType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKCategoryType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKCategoryType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKCategoryType */



