// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKStateOfMindType */


/* debug [class_header]: Header for HKStateOfMindType */
// The class instance for the [HKStateOfMindType] class.
var (
	HKStateOfMindTypeClass     _HKStateOfMindTypeClass
	HKStateOfMindTypeClassOnce sync.Once
)

func getHKStateOfMindTypeClass() _HKStateOfMindTypeClass {
	HKStateOfMindTypeClassOnce.Do(func() {
		HKStateOfMindTypeClass = _HKStateOfMindTypeClass{objc.GetClass("HKStateOfMindType")}
	})
	return HKStateOfMindTypeClass
}

type _HKStateOfMindTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKStateOfMindType */
// An interface definition for the [HKStateOfMindType] class.
type IHKStateOfMindType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKStateOfMindType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKStateOfMindType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKStateOfMindType */
// Alloc allocates a new instance without initialization.
func (hc _HKStateOfMindTypeClass) Alloc() HKStateOfMindType {
	rv := objc.Send[HKStateOfMindType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKStateOfMindTypeClass) New() HKStateOfMindType {
	rv := objc.Send[HKStateOfMindType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKStateOfMindType) Init() HKStateOfMindType {
	rv := objc.Send[HKStateOfMindType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKStateOfMindType) Autorelease() HKStateOfMindType {
	rv := objc.Send[HKStateOfMindType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKStateOfMindType creates a new HKStateOfMindType instance.
func NewHKStateOfMindType() HKStateOfMindType {
	return getHKStateOfMindTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKStateOfMindType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMindType
type HKStateOfMindType struct {
	HKSampleType
}

// HKStateOfMindTypeFrom constructs a [HKStateOfMindType] from an unsafe.Pointer.
func HKStateOfMindTypeFrom(ptr unsafe.Pointer) HKStateOfMindType {
	return HKStateOfMindType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKStateOfMindType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKStateOfMindType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKStateOfMindType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKStateOfMindType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKStateOfMindType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKStateOfMindType */



