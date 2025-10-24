// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKObjectType] class.
var (
	HKObjectTypeClass     _HKObjectTypeClass
	HKObjectTypeClassOnce sync.Once
)

func getHKObjectTypeClass() _HKObjectTypeClass {
	HKObjectTypeClassOnce.Do(func() {
		HKObjectTypeClass = _HKObjectTypeClass{objc.GetClass("HKObjectType")}
	})
	return HKObjectTypeClass
}

type _HKObjectTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKObjectType] class.
type IHKObjectType interface {
	objectivec.IObject
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// An abstract superclass with subclasses that identify a specific type of data for the HealthKit store.
//
// The class is an abstract class. Don’t instantiate an object directly. Instead, instantiate one of the following concrete subclasses: The class provides a convenience method to create each of these subclasses.


// An abstract superclass with subclasses that identify a specific type of data for the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType
type HKObjectType struct {
	objectivec.Object
}

// HKObjectTypeFrom constructs a [HKObjectType] from an unsafe.Pointer.
//
// An abstract superclass with subclasses that identify a specific type of data for the HealthKit store.
func HKObjectTypeFrom(ptr unsafe.Pointer) HKObjectType {
	return HKObjectType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKObjectTypeClass) Alloc() HKObjectType {
	rv := objc.Send[HKObjectType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKObjectTypeClass) New() HKObjectType {
	rv := objc.Send[HKObjectType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKObjectType) Init() HKObjectType {
	rv := objc.Send[HKObjectType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKObjectType) Autorelease() HKObjectType {
	rv := objc.Send[HKObjectType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKObjectType creates a new HKObjectType instance.
func NewHKObjectType() HKObjectType {
	return getHKObjectTypeClass().New()
}



// Returns the shared electrocardiogram type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/electrocardiogramType()
func (hc _HKObjectTypeClass) ElectrocardiogramType() IHKElectrocardiogramType {
	rv := objc.Send[HKElectrocardiogramType](objc.ID(hc.class), objc.Sel("electrocardiogramType"))
	return rv
}


// Returns the shared quantity type for the provided identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/quantityType(forIdentifier:)
func (hc _HKObjectTypeClass) QuantityTypeForIdentifier(identifier HKQuantityTypeIdentifier /* typedef */) IHKQuantityType {
	rv := objc.Send[HKQuantityType](objc.ID(hc.class), objc.Sel("quantityTypeForIdentifier:"), identifier)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKObjectType/stateOfMindType()
func (hc _HKObjectTypeClass) StateOfMindType() IHKStateOfMindType {
	rv := objc.Send[HKStateOfMindType](objc.ID(hc.class), objc.Sel("stateOfMindType"))
	return rv
}


// A unique string identifying the HealthKit object type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobjecttype/identifier
func (h_ HKObjectType) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("identifier"))
	return rv
}


// A unique string identifying the HealthKit object type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkobjecttype/identifier
func (h_ HKObjectType) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIdentifier:"), value)
}



