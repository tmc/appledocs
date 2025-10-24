// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKSourceRevision */


/* debug [class_header]: Header for HKSourceRevision */
// The class instance for the [HKSourceRevision] class.
var (
	HKSourceRevisionClass     _HKSourceRevisionClass
	HKSourceRevisionClassOnce sync.Once
)

func getHKSourceRevisionClass() _HKSourceRevisionClass {
	HKSourceRevisionClassOnce.Do(func() {
		HKSourceRevisionClass = _HKSourceRevisionClass{objc.GetClass("HKSourceRevision")}
	})
	return HKSourceRevisionClass
}

type _HKSourceRevisionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKSourceRevision */
// An interface definition for the [HKSourceRevision] class.
type IHKSourceRevision interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKSourceRevision */
	// properties:
	OperatingSystemVersion() foundation.OperatingSystemVersion
	ProductType() objc.IObject /* cross-framework: NSString */
	Source() IHKSource
	Version() objc.IObject /* cross-framework: NSString */
	SourceRevision() IHKSourceRevision
	SetSourceRevision(value IHKSourceRevision)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKSourceRevision */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKSourceRevision */
// Alloc allocates a new instance without initialization.
func (hc _HKSourceRevisionClass) Alloc() HKSourceRevision {
	rv := objc.Send[HKSourceRevision](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKSourceRevisionClass) New() HKSourceRevision {
	rv := objc.Send[HKSourceRevision](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSourceRevision) Init() HKSourceRevision {
	rv := objc.Send[HKSourceRevision](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSourceRevision) Autorelease() HKSourceRevision {
	rv := objc.Send[HKSourceRevision](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSourceRevision creates a new HKSourceRevision instance.
func NewHKSourceRevision() HKSourceRevision {
	return getHKSourceRevisionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKSourceRevision */
// An object indicating the source of a HealthKit sample.
//
// The class acts as a wrapper for the class, adding information about the source’s version, operating system, and product type. Source revision objects are immutable: you set the source revision’s properties when you create the object, and they cannot change. When an instance is created, its property is set to . When the object is saved to the HealthKit store, HealthKit assigns a new source revision to the object’s property. The source revision can be accessed only on objects retrieved from the HealthKit store.


// An object indicating the source of a HealthKit sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSourceRevision
type HKSourceRevision struct {
	objectivec.Object
}

// HKSourceRevisionFrom constructs a [HKSourceRevision] from an unsafe.Pointer.
//
// An object indicating the source of a HealthKit sample.
func HKSourceRevisionFrom(ptr unsafe.Pointer) HKSourceRevision {
	return HKSourceRevision{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKSourceRevision */

// Initializes a new source revision object with the provided source and version information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSourceRevision/init(source:version:)
func NewHKSourceRevisionWithSourceVersion(source IHKSource, version objc.IObject /* cross-framework: NSString */) HKSourceRevision {
	instance := getHKSourceRevisionClass().Alloc()
	rv := objc.Send[HKSourceRevision](instance.ID, objc.Sel("initWithSource:version:"), source, version)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKSourceRevisionWithSourceVersion */


// Initializes a new source revision object with the provided source, version, product type, and operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSourceRevision/init(source:version:productType:operatingSystemVersion:)
func NewHKSourceRevisionWithSourceVersionProductTypeOperatingSystemVersion(source IHKSource, version objc.IObject /* cross-framework: NSString */, productType objc.IObject /* cross-framework: NSString */, operatingSystemVersion foundation.OperatingSystemVersion) HKSourceRevision {
	instance := getHKSourceRevisionClass().Alloc()
	rv := objc.Send[HKSourceRevision](instance.ID, objc.Sel("initWithSource:version:productType:operatingSystemVersion:"), source, version, productType, operatingSystemVersion)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKSourceRevisionWithSourceVersionProductTypeOperatingSystemVersion */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKSourceRevision */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKSourceRevision */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKSourceRevision */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKSourceRevision */

// A string that identifies the operating system used to save a sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSourceRevision/operatingSystemVersion
func (h_ HKSourceRevision) OperatingSystemVersion() foundation.OperatingSystemVersion {
	rv := objc.Send[foundation.OperatingSystemVersion](h_.ID, objc.Sel("operatingSystemVersion"))
	return rv
}/* debug [instance_properties/getter]: operatingSystemVersion */


// A string that identifies the device used to save a sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSourceRevision/productType
func (h_ HKSourceRevision) ProductType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("productType"))
	return rv
}/* debug [instance_properties/getter]: productType */


// The source for a sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSourceRevision/source
func (h_ HKSourceRevision) Source() IHKSource {
	rv := objc.Send[HKSource](h_.ID, objc.Sel("source"))
	return rv
}/* debug [instance_properties/getter]: source */


// A string that identifies a particular version of the source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSourceRevision/version
func (h_ HKSourceRevision) Version() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("version"))
	return rv
}/* debug [instance_properties/getter]: version */


// The app or device that created this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/sourcerevision
func (h_ HKSourceRevision) SourceRevision() IHKSourceRevision {
	rv := objc.Send[HKSourceRevision](h_.ID, objc.Sel("sourceRevision"))
	return rv
}/* debug [instance_properties/getter]: sourceRevision */


// The app or device that created this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/sourcerevision
func (h_ HKSourceRevision) SetSourceRevision(value IHKSourceRevision) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSourceRevision:"), value)
}/* debug [instance_properties/setter]: sourceRevision */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKSourceRevision */


