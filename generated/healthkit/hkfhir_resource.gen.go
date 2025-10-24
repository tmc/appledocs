// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKFHIRResource */


/* debug [class_header]: Header for HKFHIRResource */
// The class instance for the [HKFHIRResource] class.
var (
	HKFHIRResourceClass     _HKFHIRResourceClass
	HKFHIRResourceClassOnce sync.Once
)

func getHKFHIRResourceClass() _HKFHIRResourceClass {
	HKFHIRResourceClassOnce.Do(func() {
		HKFHIRResourceClass = _HKFHIRResourceClass{objc.GetClass("HKFHIRResource")}
	})
	return HKFHIRResourceClass
}

type _HKFHIRResourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKFHIRResource */
// An interface definition for the [HKFHIRResource] class.
type IHKFHIRResource interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKFHIRResource */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	FHIRVersion() IHKFHIRVersion
	Identifier() objc.IObject /* cross-framework: NSString */
	ResourceType() HKFHIRResourceType /* typedef */
	SourceURL() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKFHIRResource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKFHIRResource */
// Alloc allocates a new instance without initialization.
func (hc _HKFHIRResourceClass) Alloc() HKFHIRResource {
	rv := objc.Send[HKFHIRResource](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKFHIRResourceClass) New() HKFHIRResource {
	rv := objc.Send[HKFHIRResource](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKFHIRResource) Init() HKFHIRResource {
	rv := objc.Send[HKFHIRResource](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKFHIRResource) Autorelease() HKFHIRResource {
	rv := objc.Send[HKFHIRResource](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKFHIRResource creates a new HKFHIRResource instance.
func NewHKFHIRResource() HKFHIRResource {
	return getHKFHIRResourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKFHIRResource */
// An object containing Fast Healthcare Interoperability Resources (FHIR) data.


// An object containing Fast Healthcare Interoperability Resources (FHIR) data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRResource
type HKFHIRResource struct {
	objectivec.Object
}

// HKFHIRResourceFrom constructs a [HKFHIRResource] from an unsafe.Pointer.
//
// An object containing Fast Healthcare Interoperability Resources (FHIR) data.
func HKFHIRResourceFrom(ptr unsafe.Pointer) HKFHIRResource {
	return HKFHIRResource{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKFHIRResource *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKFHIRResource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKFHIRResource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKFHIRResource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKFHIRResource */

// The JSON representation of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRResource/data
func (h_ HKFHIRResource) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](h_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The FHIR version used by this resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRResource/fhirVersion
func (h_ HKFHIRResource) FHIRVersion() IHKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](h_.ID, objc.Sel("FHIRVersion"))
	return rv
}/* debug [instance_properties/getter]: FHIRVersion */


// The value from the FHIR resource’s field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRResource/identifier
func (h_ HKFHIRResource) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The value from the FHIR resource’s field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRResource/resourceType
func (h_ HKFHIRResource) ResourceType() HKFHIRResourceType /* typedef */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("resourceType"))
	return rv
}/* debug [instance_properties/getter]: resourceType */


// The full URL for the source of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRResource/sourceURL
func (h_ HKFHIRResource) SourceURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](h_.ID, objc.Sel("sourceURL"))
	return rv
}/* debug [instance_properties/getter]: sourceURL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKFHIRResource */



