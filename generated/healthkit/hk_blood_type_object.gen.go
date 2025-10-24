// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKBloodTypeObject */


/* debug [class_header]: Header for HKBloodTypeObject */
// The class instance for the [HKBloodTypeObject] class.
var (
	HKBloodTypeObjectClass     _HKBloodTypeObjectClass
	HKBloodTypeObjectClassOnce sync.Once
)

func getHKBloodTypeObjectClass() _HKBloodTypeObjectClass {
	HKBloodTypeObjectClassOnce.Do(func() {
		HKBloodTypeObjectClass = _HKBloodTypeObjectClass{objc.GetClass("HKBloodTypeObject")}
	})
	return HKBloodTypeObjectClass
}

type _HKBloodTypeObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKBloodTypeObject */
// An interface definition for the [HKBloodTypeObject] class.
type IHKBloodTypeObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKBloodTypeObject */
	// properties:
	BloodType() HKBloodType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKBloodTypeObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKBloodTypeObject */
// Alloc allocates a new instance without initialization.
func (hc _HKBloodTypeObjectClass) Alloc() HKBloodTypeObject {
	rv := objc.Send[HKBloodTypeObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKBloodTypeObjectClass) New() HKBloodTypeObject {
	rv := objc.Send[HKBloodTypeObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKBloodTypeObject) Init() HKBloodTypeObject {
	rv := objc.Send[HKBloodTypeObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKBloodTypeObject) Autorelease() HKBloodTypeObject {
	rv := objc.Send[HKBloodTypeObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKBloodTypeObject creates a new HKBloodTypeObject instance.
func NewHKBloodTypeObject() HKBloodTypeObject {
	return getHKBloodTypeObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKBloodTypeObject */
// This class acts as a wrapper for the enumeration.


// This class acts as a wrapper for the enumeration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodTypeObject
type HKBloodTypeObject struct {
	objectivec.Object
}

// HKBloodTypeObjectFrom constructs a [HKBloodTypeObject] from an unsafe.Pointer.
//
// This class acts as a wrapper for the enumeration.
func HKBloodTypeObjectFrom(ptr unsafe.Pointer) HKBloodTypeObject {
	return HKBloodTypeObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKBloodTypeObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKBloodTypeObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKBloodTypeObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKBloodTypeObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKBloodTypeObject */

// The blood type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodTypeObject/bloodType
func (h_ HKBloodTypeObject) BloodType() HKBloodType {
	rv := objc.Send[HKBloodType](h_.ID, objc.Sel("bloodType"))
	return rv
}/* debug [instance_properties/getter]: bloodType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKBloodTypeObject */



