// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKVisionPrism] class.
var (
	HKVisionPrismClass     _HKVisionPrismClass
	HKVisionPrismClassOnce sync.Once
)

func getHKVisionPrismClass() _HKVisionPrismClass {
	HKVisionPrismClassOnce.Do(func() {
		HKVisionPrismClass = _HKVisionPrismClass{objc.GetClass("HKVisionPrism")}
	})
	return HKVisionPrismClass
}

type _HKVisionPrismClass struct {
	class objc.Class
}

// An interface definition for the [HKVisionPrism] class.
type IHKVisionPrism interface {
	objectivec.IObject
}

// Prescription data for eye alignment.
//
// To include prism information in a glasses prescription, start by creating an object. Then, pass this value to the ’s initializer. Finally, create the glasses prescription and save it to the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism
type HKVisionPrism struct {
	objectivec.Object
}

// HKVisionPrismFrom constructs a [HKVisionPrism] from an unsafe.Pointer.
//
// Prescription data for eye alignment.
func HKVisionPrismFrom(ptr unsafe.Pointer) HKVisionPrism {
	return HKVisionPrism{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKVisionPrismClass) Alloc() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKVisionPrismClass) New() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKVisionPrism) Init() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKVisionPrism) Autorelease() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKVisionPrism creates a new HKVisionPrism instance.
func NewHKVisionPrism() HKVisionPrism {
	return getHKVisionPrismClass().New()
}


// The strength of the correction.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/amount
func (h_ HKVisionPrism) Amount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("amount"))
	return rv
}


// SetAmount sets the value of the amount property.
// The strength of the correction.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/amount
func (h_ HKVisionPrism) SetAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAmount:"), value)
}

// The orientation of the adjustment.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/angle
func (h_ HKVisionPrism) Angle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("angle"))
	return rv
}


// SetAngle sets the value of the angle property.
// The orientation of the adjustment.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/angle
func (h_ HKVisionPrism) SetAngle(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAngle:"), value)
}

// A value indicating which eye the correction applies to.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/eye
func (h_ HKVisionPrism) Eye() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("eye"))
	return rv
}


// SetEye sets the value of the eye property.
// A value indicating which eye the correction applies to.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/eye
func (h_ HKVisionPrism) SetEye(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setEye:"), value)
}

// The strength of the horizontal correction.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/horizontalamount
func (h_ HKVisionPrism) HorizontalAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("horizontalAmount"))
	return rv
}


// SetHorizontalAmount sets the value of the horizontalAmount property.
// The strength of the horizontal correction.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/horizontalamount
func (h_ HKVisionPrism) SetHorizontalAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setHorizontalAmount:"), value)
}

// The orientation of the horizontal portion of the correction.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/horizontalbase
func (h_ HKVisionPrism) HorizontalBase() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("horizontalBase"))
	return rv
}


// SetHorizontalBase sets the value of the horizontalBase property.
// The orientation of the horizontal portion of the correction.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/horizontalbase
func (h_ HKVisionPrism) SetHorizontalBase(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setHorizontalBase:"), value)
}

// The strength of the vertical correction.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/verticalamount
func (h_ HKVisionPrism) VerticalAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("verticalAmount"))
	return rv
}


// SetVerticalAmount sets the value of the verticalAmount property.
// The strength of the vertical correction.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/verticalamount
func (h_ HKVisionPrism) SetVerticalAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setVerticalAmount:"), value)
}

// The orientation of the vertical portion of the correction.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/verticalbase
func (h_ HKVisionPrism) VerticalBase() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("verticalBase"))
	return rv
}


// SetVerticalBase sets the value of the verticalBase property.
// The orientation of the vertical portion of the correction.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprism/verticalbase
func (h_ HKVisionPrism) SetVerticalBase(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setVerticalBase:"), value)
}



