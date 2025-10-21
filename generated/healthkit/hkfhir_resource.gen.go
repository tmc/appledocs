// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HKFHIRResource] class.
type IHKFHIRResource interface {
	objectivec.IObject
}

// An object containing Fast Healthcare Interoperability Resources (FHIR) data.
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKFHIRResourceClass) Alloc() HKFHIRResource {
	rv := objc.Send[HKFHIRResource](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The JSON representation of the FHIR resource.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRResource/data
func (h_ HKFHIRResource) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("data"))
	return rv
}

// The FHIR version used by this resource.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRResource/fhirVersion
func (h_ HKFHIRResource) FHIRVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("FHIRVersion"))
	return rv
}

// The value from the FHIR resource’s field.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRResource/identifier
func (h_ HKFHIRResource) Identifier() string {
	rv := objc.Send[string](h_.ID, objc.Sel("identifier"))
	return rv
}

// The value from the FHIR resource’s field.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRResource/resourceType
func (h_ HKFHIRResource) ResourceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("resourceType"))
	return rv
}

// The full URL for the source of the FHIR resource.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRResource/sourceURL
func (h_ HKFHIRResource) SourceURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("sourceURL"))
	return rv
}



