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
	// properties:
	Associations() unsafe.Pointer
	SetAssociations(value unsafe.Pointer)
	Kind() unsafe.Pointer
	SetKind(value unsafe.Pointer)
	Labels() unsafe.Pointer
	SetLabels(value unsafe.Pointer)
	Valence() float64
	SetValence(value float64)
	ValenceClassification() unsafe.Pointer
	SetValenceClassification(value unsafe.Pointer)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstateofmind/associations-7gwps
func (h_ HKStateOfMind) Associations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("associations"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstateofmind/associations-7gwps
func (h_ HKStateOfMind) SetAssociations(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAssociations:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstateofmind/kind-swift.property
func (h_ HKStateOfMind) Kind() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("kind"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstateofmind/kind-swift.property
func (h_ HKStateOfMind) SetKind(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setKind:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstateofmind/labels-994n4
func (h_ HKStateOfMind) Labels() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("labels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstateofmind/labels-994n4
func (h_ HKStateOfMind) SetLabels(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLabels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstateofmind/valence
func (h_ HKStateOfMind) Valence() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("valence"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstateofmind/valence
func (h_ HKStateOfMind) SetValence(value float64) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setValence:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstateofmind/valenceclassification-swift.property
func (h_ HKStateOfMind) ValenceClassification() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("valenceClassification"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkstateofmind/valenceclassification-swift.property
func (h_ HKStateOfMind) SetValenceClassification(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setValenceClassification:"), value)
}



