// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Data() foundation.objc.IObject /* cross-framework: Data */
	SetData(value foundation.objc.IObject /* cross-framework: Data */)
	FhirVersion() objc.IObject /* cross-framework: HKFHIRVersion */
	SetFhirVersion(value objc.IObject /* cross-framework: HKFHIRVersion */)
	Identifier() string /* primitive/slice/pointer. */
	SetIdentifier(value string /* primitive/slice/pointer. */)
	ResourceType() unsafe.Pointer
	SetResourceType(value unsafe.Pointer)
	SourceURL() foundation.objc.IObject /* cross-framework: URL */
	SetSourceURL(value foundation.objc.IObject /* cross-framework: URL */)
	// methods:
}

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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/data
func (h_ HKFHIRResource) Data() foundation.objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](h_.ID, objc.Sel("data"))
	return rv
}


// The JSON representation of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/data
func (h_ HKFHIRResource) SetData(value foundation.objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setData:"), value)
}


// The FHIR version used by this resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/fhirversion
func (h_ HKFHIRResource) FhirVersion() objc.IObject /* cross-framework: HKFHIRVersion */ {
	rv := objc.Send[HKFHIRVersion](h_.ID, objc.Sel("fhirVersion"))
	return rv
}


// The FHIR version used by this resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/fhirversion
func (h_ HKFHIRResource) SetFhirVersion(value objc.IObject /* cross-framework: HKFHIRVersion */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setFhirVersion:"), value)
}


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/identifier
func (h_ HKFHIRResource) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("identifier"))
	return rv
}


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/identifier
func (h_ HKFHIRResource) SetIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/resourcetype
func (h_ HKFHIRResource) ResourceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("resourceType"))
	return rv
}


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/resourcetype
func (h_ HKFHIRResource) SetResourceType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setResourceType:"), value)
}


// The full URL for the source of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/sourceurl
func (h_ HKFHIRResource) SourceURL() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](h_.ID, objc.Sel("sourceURL"))
	return rv
}


// The full URL for the source of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/sourceurl
func (h_ HKFHIRResource) SetSourceURL(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSourceURL:"), value)
}



