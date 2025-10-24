// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKClinicalCoding */


/* debug [class_header]: Header for HKClinicalCoding */
// The class instance for the [HKClinicalCoding] class.
var (
	HKClinicalCodingClass     _HKClinicalCodingClass
	HKClinicalCodingClassOnce sync.Once
)

func getHKClinicalCodingClass() _HKClinicalCodingClass {
	HKClinicalCodingClassOnce.Do(func() {
		HKClinicalCodingClass = _HKClinicalCodingClass{objc.GetClass("HKClinicalCoding")}
	})
	return HKClinicalCodingClass
}

type _HKClinicalCodingClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKClinicalCoding */
// An interface definition for the [HKClinicalCoding] class.
type IHKClinicalCoding interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKClinicalCoding */
	// properties:
	Code() objc.IObject /* cross-framework: NSString */
	System() objc.IObject /* cross-framework: NSString */
	Version() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKClinicalCoding */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKClinicalCoding */
// Alloc allocates a new instance without initialization.
func (hc _HKClinicalCodingClass) Alloc() HKClinicalCoding {
	rv := objc.Send[HKClinicalCoding](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKClinicalCodingClass) New() HKClinicalCoding {
	rv := objc.Send[HKClinicalCoding](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKClinicalCoding) Init() HKClinicalCoding {
	rv := objc.Send[HKClinicalCoding](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKClinicalCoding) Autorelease() HKClinicalCoding {
	rv := objc.Send[HKClinicalCoding](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKClinicalCoding creates a new HKClinicalCoding instance.
func NewHKClinicalCoding() HKClinicalCoding {
	return getHKClinicalCodingClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKClinicalCoding */
// A clinical coding that represents a medical concept using a standardized coding system.
//
// A clinical coding pairs a , an optional , and a which identify a medical concept. This model is closely related to the .


// A clinical coding that represents a medical concept using a standardized coding system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalCoding
type HKClinicalCoding struct {
	objectivec.Object
}

// HKClinicalCodingFrom constructs a [HKClinicalCoding] from an unsafe.Pointer.
//
// A clinical coding that represents a medical concept using a standardized coding system.
func HKClinicalCodingFrom(ptr unsafe.Pointer) HKClinicalCoding {
	return HKClinicalCoding{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKClinicalCoding */

// Creates a clinical coding with the specified system, version, and code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalCoding/init(system:version:code:)
func NewHKClinicalCodingWithSystemVersionCode(system objc.IObject /* cross-framework: NSString */, version objc.IObject /* cross-framework: NSString */, code objc.IObject /* cross-framework: NSString */) HKClinicalCoding {
	instance := getHKClinicalCodingClass().Alloc()
	rv := objc.Send[HKClinicalCoding](instance.ID, objc.Sel("initWithSystem:version:code:"), system, version, code)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKClinicalCodingWithSystemVersionCode */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKClinicalCoding */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKClinicalCoding */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKClinicalCoding */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKClinicalCoding */

// The clinical code that represents a medical concept inside the coding system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalCoding/code
func (h_ HKClinicalCoding) Code() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("code"))
	return rv
}/* debug [instance_properties/getter]: code */


// The string that identifies the coding system that defines this clinical code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalCoding/system
func (h_ HKClinicalCoding) System() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("system"))
	return rv
}/* debug [instance_properties/getter]: system */


// The version of the coding system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalCoding/version
func (h_ HKClinicalCoding) Version() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("version"))
	return rv
}/* debug [instance_properties/getter]: version */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKClinicalCoding */


