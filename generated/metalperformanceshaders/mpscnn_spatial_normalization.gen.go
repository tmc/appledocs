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
