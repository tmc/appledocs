// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKStateOfMind] class.
var (
	HKStateOfMindClass     _HKStateOfMindClass
	HKStateOfMindClassOnce sync.Once
)

func getHKStateOfMindClass() _HKStateOfMindClass {
	HKStateOfMindClassOnce.Do(func() {
		HKStateOfMindClass = _HKStateOfMindClass{objc.GetClass("HKStateOfMind")}
	})
	return HKStateOfMindClass
}

type _HKStateOfMindClass struct {
	class objc.Class
}

// An interface definition for the [HKStateOfMind] class.
type IHKStateOfMind interface {
	IHKSample
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind
type HKStateOfMind struct {
	HKSample
}

// HKStateOfMindFrom constructs a [HKStateOfMind] from an unsafe.Pointer.
func HKStateOfMindFrom(ptr unsafe.Pointer) HKStateOfMind {
	return HKStateOfMind{
		HKSample: HKSampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKStateOfMindClass) Alloc() HKStateOfMind {
	rv := objc.Send[HKStateOfMind](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKStateOfMindClass) New() HKStateOfMind {
	rv := objc.Send[HKStateOfMind](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKStateOfMind) Init() HKStateOfMind {
	rv := objc.Send[HKStateOfMind](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKStateOfMind) Autorelease() HKStateOfMind {
	rv := objc.Send[HKStateOfMind](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKStateOfMind creates a new HKStateOfMind instance.
func NewHKStateOfMind() HKStateOfMind {
	return getHKStateOfMindClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/stateOfMindWithDate:kind:valence:labels:associations:
func (hc _HKStateOfMindClass) StateOfMindWithDateKindValenceLabelsAssociations(date unsafe.Pointer, kind unsafe.Pointer, valence unsafe.Pointer, labels unsafe.Pointer, associations unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("stateOfMindWithDate:kind:valence:labels:associations:"), date, kind, valence, labels, associations)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/stateOfMindWithDate:kind:valence:labels:associations:metadata:
func (hc _HKStateOfMindClass) StateOfMindWithDateKindValenceLabelsAssociationsMetadata(date unsafe.Pointer, kind unsafe.Pointer, valence unsafe.Pointer, labels unsafe.Pointer, associations unsafe.Pointer, metadata unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("stateOfMindWithDate:kind:valence:labels:associations:metadata:"), date, kind, valence, labels, associations, metadata)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/associations-5vfw8
func (h_ HKStateOfMind) Associations() []NSNumber {
	rv := objc.Send[[]NSNumber](h_.ID, objc.Sel("associations"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/kind-swift.property
func (h_ HKStateOfMind) Kind() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("kind"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/labels-11jl3
func (h_ HKStateOfMind) Labels() []NSNumber {
	rv := objc.Send[[]NSNumber](h_.ID, objc.Sel("labels"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/valence
func (h_ HKStateOfMind) Valence() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("valence"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMind/valenceClassification-swift.property
func (h_ HKStateOfMind) ValenceClassification() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("valenceClassification"))
	return rv
}



