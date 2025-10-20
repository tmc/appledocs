// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNNGroupNormalizationGradient] class.
var (
	CNNGroupNormalizationGradientClass     _CNNGroupNormalizationGradientClass
	CNNGroupNormalizationGradientClassOnce sync.Once
)

func getCNNGroupNormalizationGradientClass() _CNNGroupNormalizationGradientClass {
	CNNGroupNormalizationGradientClassOnce.Do(func() {
		CNNGroupNormalizationGradientClass = _CNNGroupNormalizationGradientClass{objc.GetClass("MPSCNNGroupNormalizationGradient")}
	})
	return CNNGroupNormalizationGradientClass
}

type _CNNGroupNormalizationGradientClass struct {
	class objc.Class
}

// An interface definition for the [CNNGroupNormalizationGradient] class.
type ICNNGroupNormalizationGradient interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradient
type CNNGroupNormalizationGradient struct {
	objectivec.Object
}

// CNNGroupNormalizationGradientFrom constructs a [CNNGroupNormalizationGradient] from an unsafe.Pointer.
func CNNGroupNormalizationGradientFrom(ptr unsafe.Pointer) CNNGroupNormalizationGradient {
	return CNNGroupNormalizationGradient{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNGroupNormalizationGradientClass) Alloc() CNNGroupNormalizationGradient {
	rv := objc.Send[CNNGroupNormalizationGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNGroupNormalizationGradientClass) New() CNNGroupNormalizationGradient {
	rv := objc.Send[CNNGroupNormalizationGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNGroupNormalizationGradient) Init() CNNGroupNormalizationGradient {
	rv := objc.Send[CNNGroupNormalizationGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNGroupNormalizationGradient) Autorelease() CNNGroupNormalizationGradient {
	rv := objc.Send[CNNGroupNormalizationGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNGroupNormalizationGradient creates a new CNNGroupNormalizationGradient instance.
func NewCNNGroupNormalizationGradient() CNNGroupNormalizationGradient {
	return getCNNGroupNormalizationGradientClass().New()
}




