// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKHealthConceptIdentifier */


/* debug [class_header]: Header for HKHealthConceptIdentifier */
// The class instance for the [HKHealthConceptIdentifier] class.
var (
	HKHealthConceptIdentifierClass     _HKHealthConceptIdentifierClass
	HKHealthConceptIdentifierClassOnce sync.Once
)

func getHKHealthConceptIdentifierClass() _HKHealthConceptIdentifierClass {
	HKHealthConceptIdentifierClassOnce.Do(func() {
		HKHealthConceptIdentifierClass = _HKHealthConceptIdentifierClass{objc.GetClass("HKHealthConceptIdentifier")}
	})
	return HKHealthConceptIdentifierClass
}

type _HKHealthConceptIdentifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKHealthConceptIdentifier */
// An interface definition for the [HKHealthConceptIdentifier] class.
type IHKHealthConceptIdentifier interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKHealthConceptIdentifier */
	// properties:
	Domain() HKHealthConceptDomain /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKHealthConceptIdentifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKHealthConceptIdentifier */
// Alloc allocates a new instance without initialization.
func (hc _HKHealthConceptIdentifierClass) Alloc() HKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKHealthConceptIdentifierClass) New() HKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKHealthConceptIdentifier) Init() HKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKHealthConceptIdentifier) Autorelease() HKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKHealthConceptIdentifier creates a new HKHealthConceptIdentifier instance.
func NewHKHealthConceptIdentifier() HKHealthConceptIdentifier {
	return getHKHealthConceptIdentifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKHealthConceptIdentifier */
// A unique identifier for a specific health concept within a domain.
//
// Each identifier points to one concept inside a domain. For example, within the medication domain, one identifier might represent ibuprofen while another represents insulin.


// A unique identifier for a specific health concept within a domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthConceptIdentifier
type HKHealthConceptIdentifier struct {
	objectivec.Object
}

// HKHealthConceptIdentifierFrom constructs a [HKHealthConceptIdentifier] from an unsafe.Pointer.
//
// A unique identifier for a specific health concept within a domain.
func HKHealthConceptIdentifierFrom(ptr unsafe.Pointer) HKHealthConceptIdentifier {
	return HKHealthConceptIdentifier{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKHealthConceptIdentifier *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKHealthConceptIdentifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKHealthConceptIdentifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKHealthConceptIdentifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKHealthConceptIdentifier */

// The domain this identifier belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKHealthConceptIdentifier/domain
func (h_ HKHealthConceptIdentifier) Domain() HKHealthConceptDomain /* typedef */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("domain"))
	return rv
}/* debug [instance_properties/getter]: domain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKHealthConceptIdentifier */



