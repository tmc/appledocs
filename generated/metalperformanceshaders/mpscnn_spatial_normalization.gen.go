// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNNSpatialNormalization] class.
var (
	CNNSpatialNormalizationClass     _CNNSpatialNormalizationClass
	CNNSpatialNormalizationClassOnce sync.Once
)

func getCNNSpatialNormalizationClass() _CNNSpatialNormalizationClass {
	CNNSpatialNormalizationClassOnce.Do(func() {
		CNNSpatialNormalizationClass = _CNNSpatialNormalizationClass{objc.GetClass("MPSCNNSpatialNormalization")}
	})
	return CNNSpatialNormalizationClass
}

type _CNNSpatialNormalizationClass struct {
	class objc.Class
}

// An interface definition for the [CNNSpatialNormalization] class.
type ICNNSpatialNormalization interface {
	objectivec.IObject
	Alpha() float32
	SetAlpha(value float32)
	Beta() float32
	SetBeta(value float32)
	Delta() float32
	SetDelta(value float32)
}

// A spatial normalization kernel.
//
// The spatial normalization for a feature channel applies the kernel over local regions which extend spatially, but are in separate feature channels (i.e., they have the shape ). For each feature channel, the function computes the sum of squares of inside each rectangle, . It then divides each element of as follows: Where and are the values of the and properties, respectively. It is your responsibility to ensure that the combination of the values of the and properties does not result in a situation where the denominator becomes zero (in such situations the resulting pixel-value is undefined).
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSpatialNormalization
type CNNSpatialNormalization struct {
	objectivec.Object
}

// CNNSpatialNormalizationFrom constructs a [CNNSpatialNormalization] from an unsafe.Pointer.
//
// A spatial normalization kernel.
func CNNSpatialNormalizationFrom(ptr unsafe.Pointer) CNNSpatialNormalization {
	return CNNSpatialNormalization{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNSpatialNormalizationClass) Alloc() CNNSpatialNormalization {
	rv := objc.Send[CNNSpatialNormalization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNSpatialNormalizationClass) New() CNNSpatialNormalization {
	rv := objc.Send[CNNSpatialNormalization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSpatialNormalization) Init() CNNSpatialNormalization {
	rv := objc.Send[CNNSpatialNormalization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSpatialNormalization) Autorelease() CNNSpatialNormalization {
	rv := objc.Send[CNNSpatialNormalization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSpatialNormalization creates a new CNNSpatialNormalization instance.
func NewCNNSpatialNormalization() CNNSpatialNormalization {
	return getCNNSpatialNormalizationClass().New()
}


// The “alpha” variable of the kernel function.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/alpha
func (c_ CNNSpatialNormalization) Alpha() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("alpha"))
	return rv
}


// SetAlpha sets the value of the alpha property.
// The “alpha” variable of the kernel function.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/alpha
func (c_ CNNSpatialNormalization) SetAlpha(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlpha:"), value)
}

// The “beta” variable of the kernel function.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/beta
func (c_ CNNSpatialNormalization) Beta() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("beta"))
	return rv
}


// SetBeta sets the value of the beta property.
// The “beta” variable of the kernel function.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/beta
func (c_ CNNSpatialNormalization) SetBeta(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}

// The “delta” variable of the kernel function.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/delta
func (c_ CNNSpatialNormalization) Delta() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("delta"))
	return rv
}


// SetDelta sets the value of the delta property.
// The “delta” variable of the kernel function.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/delta
func (c_ CNNSpatialNormalization) SetDelta(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}



