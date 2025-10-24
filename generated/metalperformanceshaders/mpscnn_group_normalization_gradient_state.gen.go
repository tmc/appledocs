// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNNGroupNormalizationGradientState] class.
var (
	CNNGroupNormalizationGradientStateClass     _CNNGroupNormalizationGradientStateClass
	CNNGroupNormalizationGradientStateClassOnce sync.Once
)

func getCNNGroupNormalizationGradientStateClass() _CNNGroupNormalizationGradientStateClass {
	CNNGroupNormalizationGradientStateClassOnce.Do(func() {
		CNNGroupNormalizationGradientStateClass = _CNNGroupNormalizationGradientStateClass{objc.GetClass("MPSCNNGroupNormalizationGradientState")}
	})
	return CNNGroupNormalizationGradientStateClass
}

type _CNNGroupNormalizationGradientStateClass struct {
	class objc.Class
}

// An interface definition for the [CNNGroupNormalizationGradientState] class.
type ICNNGroupNormalizationGradientState interface {
	IGradientState
	// properties:
	Beta() objc.ID
	GradientForBeta() objc.ID
	Gamma() Buffer /* not a class type */
	SetGamma(value Buffer /* not a class type */)
	GradientForGamma() Buffer /* not a class type */
	SetGradientForGamma(value Buffer /* not a class type */)
	GroupNormalization() IMPSCNNGroupNormalization
	SetGroupNormalization(value IMPSCNNGroupNormalization)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientState
type CNNGroupNormalizationGradientState struct {
	GradientState
}

// CNNGroupNormalizationGradientStateFrom constructs a [CNNGroupNormalizationGradientState] from an unsafe.Pointer.
func CNNGroupNormalizationGradientStateFrom(ptr unsafe.Pointer) CNNGroupNormalizationGradientState {
	return CNNGroupNormalizationGradientState{
		GradientState: GradientStateFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNNGroupNormalizationGradientStateClass) Alloc() CNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNNGroupNormalizationGradientStateClass) New() CNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNGroupNormalizationGradientState) Init() CNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNGroupNormalizationGradientState) Autorelease() CNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNGroupNormalizationGradientState creates a new CNNGroupNormalizationGradientState instance.
func NewCNNGroupNormalizationGradientState() CNNGroupNormalizationGradientState {
	return getCNNGroupNormalizationGradientStateClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientState/beta
func (c_ CNNGroupNormalizationGradientState) Beta() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("beta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalizationGradientState/gradientForBeta
func (c_ CNNGroupNormalizationGradientState) GradientForBeta() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForBeta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/gamma
func (c_ CNNGroupNormalizationGradientState) Gamma() Buffer /* not a class type */ {
	rv := objc.Send[Buffer](c_.ID, objc.Sel("gamma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/gamma
func (c_ CNNGroupNormalizationGradientState) SetGamma(value Buffer /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/gradientforgamma
func (c_ CNNGroupNormalizationGradientState) GradientForGamma() Buffer /* not a class type */ {
	rv := objc.Send[Buffer](c_.ID, objc.Sel("gradientForGamma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/gradientforgamma
func (c_ CNNGroupNormalizationGradientState) SetGradientForGamma(value Buffer /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForGamma:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/groupnormalization
func (c_ CNNGroupNormalizationGradientState) GroupNormalization() IMPSCNNGroupNormalization {
	rv := objc.Send[CNNGroupNormalization](c_.ID, objc.Sel("groupNormalization"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/groupnormalization
func (c_ CNNGroupNormalizationGradientState) SetGroupNormalization(value IMPSCNNGroupNormalization) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroupNormalization:"), value)
}



