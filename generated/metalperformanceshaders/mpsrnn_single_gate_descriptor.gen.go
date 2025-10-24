// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [RNNSingleGateDescriptor] class.
var (
	RNNSingleGateDescriptorClass     _RNNSingleGateDescriptorClass
	RNNSingleGateDescriptorClassOnce sync.Once
)

func getRNNSingleGateDescriptorClass() _RNNSingleGateDescriptorClass {
	RNNSingleGateDescriptorClassOnce.Do(func() {
		RNNSingleGateDescriptorClass = _RNNSingleGateDescriptorClass{objc.GetClass("MPSRNNSingleGateDescriptor")}
	})
	return RNNSingleGateDescriptorClass
}

type _RNNSingleGateDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [RNNSingleGateDescriptor] class.
type IRNNSingleGateDescriptor interface {
	IRNNDescriptor
	

	// properties:
	RecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	InputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	InputFeatureChannels() int
	SetInputFeatureChannels(value int)
	OutputFeatureChannels() int
	SetOutputFeatureChannels(value int)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _RNNSingleGateDescriptorClass) Alloc() RNNSingleGateDescriptor {
	rv := objc.Send[RNNSingleGateDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RNNSingleGateDescriptorClass) New() RNNSingleGateDescriptor {
	rv := objc.Send[RNNSingleGateDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNSingleGateDescriptor) Init() RNNSingleGateDescriptor {
	rv := objc.Send[RNNSingleGateDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNSingleGateDescriptor) Autorelease() RNNSingleGateDescriptor {
	rv := objc.Send[RNNSingleGateDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNSingleGateDescriptor creates a new RNNSingleGateDescriptor instance.
func NewRNNSingleGateDescriptor() RNNSingleGateDescriptor {
	return getRNNSingleGateDescriptorClass().New()
}





// A description of a simple recurrent block or layer.
//
// The recurrent neural network (RNN) layer initialized with a transforms the input data (image or matrix) and previous output with a set of filters. Each produces one feature map in the new output data. You may provide the RNN unit with a single input or a sequence of inputs.


// A description of a simple recurrent block or layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNSingleGateDescriptor
type RNNSingleGateDescriptor struct {
	RNNDescriptor
}

// RNNSingleGateDescriptorFrom constructs a [RNNSingleGateDescriptor] from an unsafe.Pointer.
//
// A description of a simple recurrent block or layer.
func RNNSingleGateDescriptorFrom(ptr unsafe.Pointer) RNNSingleGateDescriptor {
	return RNNSingleGateDescriptor{
		RNNDescriptor: RNNDescriptorFrom(ptr),
	}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865703-creaternnsinglegatedescriptor
func (rc _RNNSingleGateDescriptorClass) CreateRNNSingleGateDescriptor() {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("createRNNSingleGateDescriptor"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865703-creaternnsinglegatedescriptorwit
func (rc _RNNSingleGateDescriptorClass) CreateRNNSingleGateDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("createRNNSingleGateDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865686-recurrentweights
func (r_ RNNSingleGateDescriptor) RecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("recurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865686-recurrentweights
func (r_ RNNSingleGateDescriptor) SetRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865723-inputweights
func (r_ RNNSingleGateDescriptor) InputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("inputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865723-inputweights
func (r_ RNNSingleGateDescriptor) SetInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (r_ RNNSingleGateDescriptor) InputFeatureChannels() int {
	rv := objc.Send[int](r_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (r_ RNNSingleGateDescriptor) SetInputFeatureChannels(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (r_ RNNSingleGateDescriptor) OutputFeatureChannels() int {
	rv := objc.Send[int](r_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (r_ RNNSingleGateDescriptor) SetOutputFeatureChannels(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}








