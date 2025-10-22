// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKQueryDescriptor] class.
var (
	HKQueryDescriptorClass     _HKQueryDescriptorClass
	HKQueryDescriptorClassOnce sync.Once
)

func getHKQueryDescriptorClass() _HKQueryDescriptorClass {
	HKQueryDescriptorClassOnce.Do(func() {
		HKQueryDescriptorClass = _HKQueryDescriptorClass{objc.GetClass("HKQueryDescriptor")}
	})
	return HKQueryDescriptorClass
}

type _HKQueryDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [HKQueryDescriptor] class.
type IHKQueryDescriptor interface {
	objectivec.IObject
	Predicate() foundation.Predicate
	SetPredicate(value foundation.IPredicate)
	SampleType() HKSampleType
	SetSampleType(value HKSampleType)
}

// A descriptor that specifies a set of samples based on the data type and a predicate.
//
// Use descriptors to create queries that return multiple data types. You can use descriptors when creating , , or instances.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryDescriptor
type HKQueryDescriptor struct {
	objectivec.Object
}

// HKQueryDescriptorFrom constructs a [HKQueryDescriptor] from an unsafe.Pointer.
//
// A descriptor that specifies a set of samples based on the data type and a predicate.
func HKQueryDescriptorFrom(ptr unsafe.Pointer) HKQueryDescriptor {
	return HKQueryDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKQueryDescriptorClass) Alloc() HKQueryDescriptor {
	rv := objc.Send[HKQueryDescriptor](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKQueryDescriptorClass) New() HKQueryDescriptor {
	rv := objc.Send[HKQueryDescriptor](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQueryDescriptor) Init() HKQueryDescriptor {
	rv := objc.Send[HKQueryDescriptor](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQueryDescriptor) Autorelease() HKQueryDescriptor {
	rv := objc.Send[HKQueryDescriptor](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQueryDescriptor creates a new HKQueryDescriptor instance.
func NewHKQueryDescriptor() HKQueryDescriptor {
	return getHKQueryDescriptorClass().New()
}


// The predicate that filters samples matching this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquerydescriptor/predicate
func (h_ HKQueryDescriptor) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](h_.ID, objc.Sel("predicate"))
	return rv
}


// SetPredicate sets the value of the predicate property.
// The predicate that filters samples matching this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquerydescriptor/predicate
func (h_ HKQueryDescriptor) SetPredicate(value foundation.IPredicate) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPredicate:"), value)
}

// The data type of samples that match this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquerydescriptor/sampletype
func (h_ HKQueryDescriptor) SampleType() HKSampleType {
	rv := objc.Send[HKSampleType](h_.ID, objc.Sel("sampleType"))
	return rv
}


// SetSampleType sets the value of the sampleType property.
// The data type of samples that match this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkquerydescriptor/sampletype
func (h_ HKQueryDescriptor) SetSampleType(value HKSampleType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSampleType:"), value)
}



