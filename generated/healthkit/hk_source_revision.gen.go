// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HKSourceRevision] class.
type IHKSourceRevision interface {
	objectivec.IObject
	// properties:
	SourceRevision() IHKSourceRevision
	SetSourceRevision(value IHKSourceRevision)
	OperatingSystemVersion() objc.IObject /* cross-framework: OperatingSystemVersion */
	SetOperatingSystemVersion(value objc.IObject /* cross-framework: OperatingSystemVersion */)
	ProductType() objc.IObject /* cross-framework: NSString */
	SetProductType(value objc.IObject /* cross-framework: NSString */)
	Source() IHKSource
	SetSource(value IHKSource)
	Version() objc.IObject /* cross-framework: NSString */
	SetVersion(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (hc _HKSourceRevisionClass) Alloc() HKSourceRevision {
	rv := objc.Send[HKSourceRevision](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The app or device that created this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/sourcerevision
func (h_ HKSourceRevision) SourceRevision() IHKSourceRevision {
	rv := objc.Send[HKSourceRevision](h_.ID, objc.Sel("sourceRevision"))
	return rv
}


// The app or device that created this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/sourcerevision
func (h_ HKSourceRevision) SetSourceRevision(value IHKSourceRevision) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSourceRevision:"), value)
}


// A string that identifies the operating system used to save a sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksourcerevision/operatingsystemversion
func (h_ HKSourceRevision) OperatingSystemVersion() objc.IObject /* cross-framework: OperatingSystemVersion */ {
	rv := objc.Send[foundation.OperatingSystemVersion](h_.ID, objc.Sel("operatingSystemVersion"))
	return rv
}


// A string that identifies the operating system used to save a sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksourcerevision/operatingsystemversion
func (h_ HKSourceRevision) SetOperatingSystemVersion(value objc.IObject /* cross-framework: OperatingSystemVersion */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setOperatingSystemVersion:"), value)
}


// A string that identifies the device used to save a sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksourcerevision/producttype
func (h_ HKSourceRevision) ProductType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("productType"))
	return rv
}


// A string that identifies the device used to save a sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksourcerevision/producttype
func (h_ HKSourceRevision) SetProductType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setProductType:"), value)
}


// The source for a sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksourcerevision/source
func (h_ HKSourceRevision) Source() IHKSource {
	rv := objc.Send[HKSource](h_.ID, objc.Sel("source"))
	return rv
}


// The source for a sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksourcerevision/source
func (h_ HKSourceRevision) SetSource(value IHKSource) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSource:"), value)
}


// A string that identifies a particular version of the source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksourcerevision/version
func (h_ HKSourceRevision) Version() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("version"))
	return rv
}


// A string that identifies a particular version of the source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksourcerevision/version
func (h_ HKSourceRevision) SetVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setVersion:"), value)
}



