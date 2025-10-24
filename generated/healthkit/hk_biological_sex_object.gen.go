// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKBiologicalSexObject */


/* debug [class_header]: Header for HKBiologicalSexObject */
// The class instance for the [HKBiologicalSexObject] class.
var (
	HKBiologicalSexObjectClass     _HKBiologicalSexObjectClass
	HKBiologicalSexObjectClassOnce sync.Once
)

func getHKBiologicalSexObjectClass() _HKBiologicalSexObjectClass {
	HKBiologicalSexObjectClassOnce.Do(func() {
		HKBiologicalSexObjectClass = _HKBiologicalSexObjectClass{objc.GetClass("HKBiologicalSexObject")}
	})
	return HKBiologicalSexObjectClass
}

type _HKBiologicalSexObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKBiologicalSexObject */
// An interface definition for the [HKBiologicalSexObject] class.
type IHKBiologicalSexObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKBiologicalSexObject */
	// properties:
	BiologicalSex() HKBiologicalSex
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKBiologicalSexObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKBiologicalSexObject */
// Alloc allocates a new instance without initialization.
func (hc _HKBiologicalSexObjectClass) Alloc() HKBiologicalSexObject {
	rv := objc.Send[HKBiologicalSexObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKBiologicalSexObjectClass) New() HKBiologicalSexObject {
	rv := objc.Send[HKBiologicalSexObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKBiologicalSexObject) Init() HKBiologicalSexObject {
	rv := objc.Send[HKBiologicalSexObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKBiologicalSexObject) Autorelease() HKBiologicalSexObject {
	rv := objc.Send[HKBiologicalSexObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKBiologicalSexObject creates a new HKBiologicalSexObject instance.
func NewHKBiologicalSexObject() HKBiologicalSexObject {
	return getHKBiologicalSexObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKBiologicalSexObject */
// This class acts as a wrapper for the enumeration.


// This class acts as a wrapper for the enumeration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSexObject
type HKBiologicalSexObject struct {
	objectivec.Object
}

// HKBiologicalSexObjectFrom constructs a [HKBiologicalSexObject] from an unsafe.Pointer.
//
// This class acts as a wrapper for the enumeration.
func HKBiologicalSexObjectFrom(ptr unsafe.Pointer) HKBiologicalSexObject {
	return HKBiologicalSexObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKBiologicalSexObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKBiologicalSexObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKBiologicalSexObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKBiologicalSexObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKBiologicalSexObject */

// The biological sex.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSexObject/biologicalSex
func (h_ HKBiologicalSexObject) BiologicalSex() HKBiologicalSex {
	rv := objc.Send[HKBiologicalSex](h_.ID, objc.Sel("biologicalSex"))
	return rv
}/* debug [instance_properties/getter]: biologicalSex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKBiologicalSexObject */



