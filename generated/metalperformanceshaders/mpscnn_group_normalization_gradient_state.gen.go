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
	Beta() Buffer get /* not a class type */
	SetBeta(value Buffer get /* not a class type */)
	Gamma() Buffer get /* not a class type */
	SetGamma(value Buffer get /* not a class type */)
	GradientForBeta() Buffer get /* not a class type */
	SetGradientForBeta(value Buffer get /* not a class type */)
	GradientForGamma() Buffer get /* not a class type */
	SetGradientForGamma(value Buffer get /* not a class type */)
	GroupNormalization() IMPSCNNGroupNormalization
	SetGroupNormalization(value IMPSCNNGroupNormalization)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNGroupNormalizationGradientStateClass) Alloc() CNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152557-beta
func (c_ CNNGroupNormalizationGradientState) Beta() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("beta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152557-beta
func (c_ CNNGroupNormalizationGradientState) SetBeta(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBeta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152558-gamma
func (c_ CNNGroupNormalizationGradientState) Gamma() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gamma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152558-gamma
func (c_ CNNGroupNormalizationGradientState) SetGamma(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGamma:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152559-gradientforbeta
func (c_ CNNGroupNormalizationGradientState) GradientForBeta() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForBeta"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152559-gradientforbeta
func (c_ CNNGroupNormalizationGradientState) SetGradientForBeta(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForBeta:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152560-gradientforgamma
func (c_ CNNGroupNormalizationGradientState) GradientForGamma() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForGamma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152560-gradientforgamma
func (c_ CNNGroupNormalizationGradientState) SetGradientForGamma(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForGamma:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152561-groupnormalization
func (c_ CNNGroupNormalizationGradientState) GroupNormalization() IMPSCNNGroupNormalization {
	rv := objc.Send[CNNGroupNormalization](c_.ID, objc.Sel("groupNormalization"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152561-groupnormalization
func (c_ CNNGroupNormalizationGradientState) SetGroupNormalization(value IMPSCNNGroupNormalization) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroupNormalization:"), value)
}








