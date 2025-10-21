// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A piece of data that can be stored inside the HealthKit store.
//
// The class is an abstract class. You should never instantiate a object directly. Instead, always work with one of its concrete subclasses: , , , or . HealthKit objects are all immutable. With a few exceptions (such as the object’s source revision), the object’s properties are set when the object is first created and they cannot change.
//
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


// The key path for accessing the object’s UUID inside a predicate format string.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathuuid
func (h_ HKObject) HKPredicateKeyPathUUID() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathUUID"))
	return rv
}

// The key path for accessing the object’s metadata dictionary inside a predicate format string.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmetadata
func (h_ HKObject) HKPredicateKeyPathMetadata() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathMetadata"))
	return rv
}

// The device that generated the data for this object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject/device
func (h_ HKObject) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("device"))
	return rv
}

// The metadata for this HealthKit object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject/metadata
func (h_ HKObject) Metadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("metadata"))
	return rv
}

// A HealthKit source, representing the app or device that created this object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject/source
func (h_ HKObject) Source() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("source"))
	return rv
}

// The app or device that created this object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject/sourceRevision
func (h_ HKObject) SourceRevision() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("sourceRevision"))
	return rv
}

// The universally unique identifier (UUID) for this HealthKit object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObject/uuid
func (h_ HKObject) UUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("UUID"))
	return rv
}



