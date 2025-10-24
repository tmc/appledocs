// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKObject] class.
var (
	HKObjectClass     _HKObjectClass
	HKObjectClassOnce sync.Once
)

func getHKObjectClass() _HKObjectClass {
	HKObjectClassOnce.Do(func() {
		HKObjectClass = _HKObjectClass{objc.GetClass("HKObject")}
	})
	return HKObjectClass
}

type _HKObjectClass struct {
	class objc.Class
}

// An interface definition for the [HKObject] class.
type IHKObject interface {
	objectivec.IObject
	// properties:
	Device() IHKDevice
	SetDevice(value IHKDevice)
	Metadata() objc.IObject /* cross-framework: NSString */
	SetMetadata(value objc.IObject /* cross-framework: NSString */)
	Source() IHKSource
	SetSource(value IHKSource)
	SourceRevision() IHKSourceRevision
	SetSourceRevision(value IHKSourceRevision)
	Uuid() objc.IObject /* cross-framework: UUID */
	SetUuid(value objc.IObject /* cross-framework: UUID */)
	HKPredicateKeyPathMetadata() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathUUID() objc.IObject /* cross-framework: NSString */
	// methods:
}

// A piece of data that can be stored inside the HealthKit store.
//
// The class is an abstract class. You should never instantiate a object directly. Instead, always work with one of its concrete subclasses: , , , or . HealthKit objects are all immutable. With a few exceptions (such as the object’s source revision), the object’s properties are set when the object is first created and they cannot change.


// A piece of data that can be stored inside the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject
type HKObject struct {
	objectivec.Object
}

// HKObjectFrom constructs a [HKObject] from an unsafe.Pointer.
//
// A piece of data that can be stored inside the HealthKit store.
func HKObjectFrom(ptr unsafe.Pointer) HKObject {
	return HKObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKObjectClass) Alloc() HKObject {
	rv := objc.Send[HKObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKObjectClass) New() HKObject {
	rv := objc.Send[HKObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKObject) Init() HKObject {
	rv := objc.Send[HKObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKObject) Autorelease() HKObject {
	rv := objc.Send[HKObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKObject creates a new HKObject instance.
func NewHKObject() HKObject {
	return getHKObjectClass().New()
}



// The device that generated the data for this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/device
func (h_ HKObject) Device() IHKDevice {
	rv := objc.Send[HKDevice](h_.ID, objc.Sel("device"))
	return rv
}


// The device that generated the data for this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/device
func (h_ HKObject) SetDevice(value IHKDevice) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDevice:"), value)
}


// The metadata for this HealthKit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/metadata
func (h_ HKObject) Metadata() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("metadata"))
	return rv
}


// The metadata for this HealthKit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/metadata
func (h_ HKObject) SetMetadata(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMetadata:"), value)
}


// A HealthKit source, representing the app or device that created this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/source
func (h_ HKObject) Source() IHKSource {
	rv := objc.Send[HKSource](h_.ID, objc.Sel("source"))
	return rv
}


// A HealthKit source, representing the app or device that created this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/source
func (h_ HKObject) SetSource(value IHKSource) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSource:"), value)
}


// The app or device that created this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/sourcerevision
func (h_ HKObject) SourceRevision() IHKSourceRevision {
	rv := objc.Send[HKSourceRevision](h_.ID, objc.Sel("sourceRevision"))
	return rv
}


// The app or device that created this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/sourcerevision
func (h_ HKObject) SetSourceRevision(value IHKSourceRevision) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSourceRevision:"), value)
}


// The universally unique identifier (UUID) for this HealthKit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/uuid
func (h_ HKObject) Uuid() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](h_.ID, objc.Sel("uuid"))
	return rv
}


// The universally unique identifier (UUID) for this HealthKit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobject/uuid
func (h_ HKObject) SetUuid(value objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setUuid:"), value)
}


// The key path for accessing the object’s metadata dictionary inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmetadata
func (h_ HKObject) HKPredicateKeyPathMetadata() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMetadata"))
	return rv
}


// The key path for accessing the object’s UUID inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathuuid
func (h_ HKObject) HKPredicateKeyPathUUID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathUUID"))
	return rv
}



