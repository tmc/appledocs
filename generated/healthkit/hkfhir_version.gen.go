// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKFHIRVersion] class.
var (
	HKFHIRVersionClass     _HKFHIRVersionClass
	HKFHIRVersionClassOnce sync.Once
)

func getHKFHIRVersionClass() _HKFHIRVersionClass {
	HKFHIRVersionClassOnce.Do(func() {
		HKFHIRVersionClass = _HKFHIRVersionClass{objc.GetClass("HKFHIRVersion")}
	})
	return HKFHIRVersionClass
}

type _HKFHIRVersionClass struct {
	class objc.Class
}

// An interface definition for the [HKFHIRVersion] class.
type IHKFHIRVersion interface {
	objectivec.IObject
	Data() foundation.Data
	SetData(value foundation.IData)
	FhirVersion() HKFHIRVersion
	SetFhirVersion(value IHKFHIRVersion)
	Identifier() string
	SetIdentifier(value string)
	ResourceType() HKFHIRResourceType
	SetResourceType(value HKFHIRResourceType)
	SourceURL() foundation.URL
	SetSourceURL(value foundation.IURL)
	FhirRelease() unsafe.Pointer
	SetFhirRelease(value unsafe.Pointer)
	MajorVersion() int
	SetMajorVersion(value int)
	MinorVersion() int
	SetMinorVersion(value int)
	PatchVersion() int
	SetPatchVersion(value int)
	StringRepresentation() string
	SetStringRepresentation(value string)
}

// The FHIR version.
//
// Use an instance to represent the version of the Fast Healthcare Interoperability Resources (FHIR) standard used to create a sample.


// The FHIR version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRVersion

type HKFHIRVersion struct {
	objectivec.Object
}

// HKFHIRVersionFrom constructs a [HKFHIRVersion] from an unsafe.Pointer.
//
// The FHIR version.
func HKFHIRVersionFrom(ptr unsafe.Pointer) HKFHIRVersion {
	return HKFHIRVersion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKFHIRVersionClass) Alloc() HKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKFHIRVersionClass) New() HKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKFHIRVersion) Init() HKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKFHIRVersion) Autorelease() HKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKFHIRVersion creates a new HKFHIRVersion instance.
func NewHKFHIRVersion() HKFHIRVersion {
	return getHKFHIRVersionClass().New()
}



// The JSON representation of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/data

func (h_ HKFHIRVersion) Data() foundation.Data {
	rv := objc.Send[foundation.Data](h_.ID, objc.Sel("data"))
	return rv
}


// The JSON representation of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/data

func (h_ HKFHIRVersion) SetData(value foundation.IData) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setData:"), value)
}


// The FHIR version used by this resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/fhirversion

func (h_ HKFHIRVersion) FhirVersion() HKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](h_.ID, objc.Sel("fhirVersion"))
	return rv
}


// The FHIR version used by this resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/fhirversion

func (h_ HKFHIRVersion) SetFhirVersion(value IHKFHIRVersion) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setFhirVersion:"), value)
}


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/identifier

func (h_ HKFHIRVersion) Identifier() string {
	rv := objc.Send[string](h_.ID, objc.Sel("identifier"))
	return rv
}


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/identifier

func (h_ HKFHIRVersion) SetIdentifier(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/resourcetype

func (h_ HKFHIRVersion) ResourceType() HKFHIRResourceType {
	rv := objc.Send[HKFHIRResourceType](h_.ID, objc.Sel("resourceType"))
	return rv
}


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/resourcetype

func (h_ HKFHIRVersion) SetResourceType(value HKFHIRResourceType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setResourceType:"), value)
}


// The full URL for the source of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/sourceurl

func (h_ HKFHIRVersion) SourceURL() foundation.URL {
	rv := objc.Send[foundation.URL](h_.ID, objc.Sel("sourceURL"))
	return rv
}


// The full URL for the source of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/sourceurl

func (h_ HKFHIRVersion) SetSourceURL(value foundation.IURL) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSourceURL:"), value)
}


// An official release of the FHIR specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirversion/fhirrelease

func (h_ HKFHIRVersion) FhirRelease() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("fhirRelease"))
	return rv
}


// An official release of the FHIR specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirversion/fhirrelease

func (h_ HKFHIRVersion) SetFhirRelease(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setFhirRelease:"), value)
}


// The standard’s major version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirversion/majorversion

func (h_ HKFHIRVersion) MajorVersion() int {
	rv := objc.Send[int](h_.ID, objc.Sel("majorVersion"))
	return rv
}


// The standard’s major version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirversion/majorversion

func (h_ HKFHIRVersion) SetMajorVersion(value int) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMajorVersion:"), value)
}


// The standard’s minor version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirversion/minorversion

func (h_ HKFHIRVersion) MinorVersion() int {
	rv := objc.Send[int](h_.ID, objc.Sel("minorVersion"))
	return rv
}


// The standard’s minor version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirversion/minorversion

func (h_ HKFHIRVersion) SetMinorVersion(value int) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMinorVersion:"), value)
}


// The standard’s patch version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirversion/patchversion

func (h_ HKFHIRVersion) PatchVersion() int {
	rv := objc.Send[int](h_.ID, objc.Sel("patchVersion"))
	return rv
}


// The standard’s patch version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirversion/patchversion

func (h_ HKFHIRVersion) SetPatchVersion(value int) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPatchVersion:"), value)
}


// A string representation of the version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirversion/stringrepresentation

func (h_ HKFHIRVersion) StringRepresentation() string {
	rv := objc.Send[string](h_.ID, objc.Sel("stringRepresentation"))
	return rv
}


// A string representation of the version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirversion/stringrepresentation

func (h_ HKFHIRVersion) SetStringRepresentation(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStringRepresentation:"), objc.String(value))
}



