// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNConvolutionGradientState] class.
var (
	CNNConvolutionGradientStateClass     _CNNConvolutionGradientStateClass
	CNNConvolutionGradientStateClassOnce sync.Once
)

func getCNNConvolutionGradientStateClass() _CNNConvolutionGradientStateClass {
	CNNConvolutionGradientStateClassOnce.Do(func() {
		CNNConvolutionGradientStateClass = _CNNConvolutionGradientStateClass{objc.GetClass("MPSCNNConvolutionGradientState")}
	})
	return CNNConvolutionGradientStateClass
}

type _CNNConvolutionGradientStateClass struct {
	class objc.Class
}





// An interface definition for the [CNNConvolutionGradientState] class.
type ICNNConvolutionGradientState interface {
	IGradientState
	

	// properties:
	GradientForBiases() Buffer get /* not a class type */
	SetGradientForBiases(value Buffer get /* not a class type */)
	GradientForWeights() Buffer get /* not a class type */
	SetGradientForWeights(value Buffer get /* not a class type */)
	Convolution() IMPSCNNConvolution
	SetConvolution(value IMPSCNNConvolution)
	GradientForWeightsLayout() CNNConvolutionWeightsLayout get /* not a class type */
	SetGradientForWeightsLayout(value CNNConvolutionWeightsLayout get /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionGradientStateClass) Alloc() CNNConvolutionGradientState {
	rv := objc.Send[CNNConvolutionGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionGradientStateClass) New() CNNConvolutionGradientState {
	rv := objc.Send[CNNConvolutionGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionGradientState) Init() CNNConvolutionGradientState {
	rv := objc.Send[CNNConvolutionGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionGradientState) Autorelease() CNNConvolutionGradientState {
	rv := objc.Send[CNNConvolutionGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionGradientState creates a new CNNConvolutionGradientState instance.
func NewCNNConvolutionGradientState() CNNConvolutionGradientState {
	return getCNNConvolutionGradientStateClass().New()
}





// An object that exposes a gradient convolution kernel’s gradient with respect to weights and biases.


// An object that exposes a gradient convolution kernel’s gradient with respect to weights and biases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionGradientState
type CNNConvolutionGradientState struct {
	GradientState
}

// CNNConvolutionGradientStateFrom constructs a [CNNConvolutionGradientState] from an unsafe.Pointer.
//
// An object that exposes a gradient convolution kernel’s gradient with respect to weights and biases.
func CNNConvolutionGradientStateFrom(ptr unsafe.Pointer) CNNConvolutionGradientState {
	return CNNConvolutionGradientState{
		GradientState: GradientStateFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/2947887-gradientforbiases
func (c_ CNNConvolutionGradientState) GradientForBiases() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForBiases"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/2947887-gradientforbiases
func (c_ CNNConvolutionGradientState) SetGradientForBiases(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForBiases:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/2947889-gradientforweights
func (c_ CNNConvolutionGradientState) GradientForWeights() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/2947889-gradientforweights
func (c_ CNNConvolutionGradientState) SetGradientForWeights(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/2953958-convolution
func (c_ CNNConvolutionGradientState) Convolution() IMPSCNNConvolution {
	rv := objc.Send[CNNConvolution](c_.ID, objc.Sel("convolution"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/2953958-convolution
func (c_ CNNConvolutionGradientState) SetConvolution(value IMPSCNNConvolution) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConvolution:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/3325841-gradientforweightslayout
func (c_ CNNConvolutionGradientState) GradientForWeightsLayout() CNNConvolutionWeightsLayout get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("gradientForWeightsLayout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/3325841-gradientforweightslayout
func (c_ CNNConvolutionGradientState) SetGradientForWeightsLayout(value CNNConvolutionWeightsLayout get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGradientForWeightsLayout:"), value)
}








