// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKFHIRVersion */


/* debug [class_header]: Header for HKFHIRVersion */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKFHIRVersion */
// An interface definition for the [HKFHIRVersion] class.
type IHKFHIRVersion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKFHIRVersion */
	// properties:
	FHIRRelease() HKFHIRRelease /* typedef */
	MajorVersion() int
	MinorVersion() int
	PatchVersion() int
	StringRepresentation() objc.IObject /* cross-framework: NSString */
	Data() foundation.Data
	SetData(value foundation.Data)
	FhirVersion() IHKFHIRVersion
	SetFhirVersion(value IHKFHIRVersion)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	ResourceType() HKFHIRResourceType /* typedef */
	SetResourceType(value HKFHIRResourceType /* typedef */)
	SourceURL() foundation.URL
	SetSourceURL(value foundation.URL)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKFHIRVersion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKFHIRVersion */
// Alloc allocates a new instance without initialization.
func (hc _HKFHIRVersionClass) Alloc() HKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKFHIRVersion */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKFHIRVersion */

// Creates an FHIR version object from a string representation of the version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRVersion/init(fromVersionString:)
func NewHKFHIRVersionFromVersionStringError(versionString objc.IObject /* cross-framework: NSString */, errorOut objectivec.IObject) HKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](objc.ID(getHKFHIRVersionClass().class), objc.Sel("versionFromVersionString:error:"), versionString, errorOut)
	return rv
}/* debug [class_init_methods/constructor]: NewHKFHIRVersionFromVersionStringError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKFHIRVersion */

// Creates an FHIR version object from a string representation of the version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRVersion/init(fromVersionString:)
func (hc _HKFHIRVersionClass) VersionFromVersionStringError(versionString objc.IObject /* cross-framework: NSString */, errorOut objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("versionFromVersionString:error:"), versionString, errorOut)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VersionFromVersionStringError) */


// Returns the primary Second Draft Standard for Trial Use (DSTU2) version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRVersion/primaryDSTU2()
func (hc _HKFHIRVersionClass) PrimaryDSTU2Version() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("primaryDSTU2Version"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrimaryDSTU2Version) */


// Returns the primary Release 4 (R4) version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRVersion/primaryR4()
func (hc _HKFHIRVersionClass) PrimaryR4Version() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("primaryR4Version"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrimaryR4Version) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKFHIRVersion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKFHIRVersion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKFHIRVersion */

// An official release of the FHIR specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRVersion/fhirRelease
func (h_ HKFHIRVersion) FHIRRelease() HKFHIRRelease /* typedef */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("FHIRRelease"))
	return rv
}/* debug [instance_properties/getter]: FHIRRelease */


// The standard’s major version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRVersion/majorVersion
func (h_ HKFHIRVersion) MajorVersion() int {
	rv := objc.Send[int](h_.ID, objc.Sel("majorVersion"))
	return rv
}/* debug [instance_properties/getter]: majorVersion */


// The standard’s minor version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRVersion/minorVersion
func (h_ HKFHIRVersion) MinorVersion() int {
	rv := objc.Send[int](h_.ID, objc.Sel("minorVersion"))
	return rv
}/* debug [instance_properties/getter]: minorVersion */


// The standard’s patch version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRVersion/patchVersion
func (h_ HKFHIRVersion) PatchVersion() int {
	rv := objc.Send[int](h_.ID, objc.Sel("patchVersion"))
	return rv
}/* debug [instance_properties/getter]: patchVersion */


// A string representation of the version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRVersion/stringRepresentation
func (h_ HKFHIRVersion) StringRepresentation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("stringRepresentation"))
	return rv
}/* debug [instance_properties/getter]: stringRepresentation */


// The JSON representation of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/data
func (h_ HKFHIRVersion) Data() foundation.Data {
	rv := objc.Send[foundation.Data](h_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The JSON representation of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/data
func (h_ HKFHIRVersion) SetData(value foundation.Data) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// The FHIR version used by this resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/fhirversion
func (h_ HKFHIRVersion) FhirVersion() IHKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](h_.ID, objc.Sel("fhirVersion"))
	return rv
}/* debug [instance_properties/getter]: fhirVersion */


// The FHIR version used by this resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/fhirversion
func (h_ HKFHIRVersion) SetFhirVersion(value IHKFHIRVersion) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setFhirVersion:"), value)
}/* debug [instance_properties/setter]: fhirVersion */


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/identifier
func (h_ HKFHIRVersion) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/identifier
func (h_ HKFHIRVersion) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/resourcetype
func (h_ HKFHIRVersion) ResourceType() HKFHIRResourceType /* typedef */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("resourceType"))
	return rv
}/* debug [instance_properties/getter]: resourceType */


// The value from the FHIR resource’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/resourcetype
func (h_ HKFHIRVersion) SetResourceType(value HKFHIRResourceType /* typedef */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setResourceType:"), value)
}/* debug [instance_properties/setter]: resourceType */


// The full URL for the source of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/sourceurl
func (h_ HKFHIRVersion) SourceURL() foundation.URL {
	rv := objc.Send[foundation.URL](h_.ID, objc.Sel("sourceURL"))
	return rv
}/* debug [instance_properties/getter]: sourceURL */


// The full URL for the source of the FHIR resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkfhirresource/sourceurl
func (h_ HKFHIRVersion) SetSourceURL(value foundation.URL) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSourceURL:"), value)
}/* debug [instance_properties/setter]: sourceURL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKFHIRVersion */


