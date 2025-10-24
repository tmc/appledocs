// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [GRUDescriptor] class.
var (
	GRUDescriptorClass     _GRUDescriptorClass
	GRUDescriptorClassOnce sync.Once
)

func getGRUDescriptorClass() _GRUDescriptorClass {
	GRUDescriptorClassOnce.Do(func() {
		GRUDescriptorClass = _GRUDescriptorClass{objc.GetClass("MPSGRUDescriptor")}
	})
	return GRUDescriptorClass
}

type _GRUDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [GRUDescriptor] class.
type IGRUDescriptor interface {
	IRNNDescriptor
	

	// properties:
	InputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetInputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	RecurrentGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetRecurrentGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	OutputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetOutputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	RecurrentGateInputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetRecurrentGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	OutputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetOutputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	InputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetInputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	GatePnormValue() objectivec.IObject
	SetGatePnormValue(value objectivec.IObject)
	OutputGateInputGateWeights() CNNConvolutionDataSource get set /* not a class type */
	SetOutputGateInputGateWeights(value CNNConvolutionDataSource get set /* not a class type */)
	FlipOutputGates() objectivec.IObject
	SetFlipOutputGates(value objectivec.IObject)
	InputFeatureChannels() int
	SetInputFeatureChannels(value int)
	OutputFeatureChannels() int
	SetOutputFeatureChannels(value int)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (gc _GRUDescriptorClass) Alloc() GRUDescriptor {
	rv := objc.Send[GRUDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GRUDescriptorClass) New() GRUDescriptor {
	rv := objc.Send[GRUDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GRUDescriptor) Init() GRUDescriptor {
	rv := objc.Send[GRUDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GRUDescriptor) Autorelease() GRUDescriptor {
	rv := objc.Send[GRUDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGRUDescriptor creates a new GRUDescriptor instance.
func NewGRUDescriptor() GRUDescriptor {
	return getGRUDescriptorClass().New()
}





// A description of a gated recurrent unit block or layer.
//
// The recurrent neural network (RNN) layer initialized with a transforms the input data (image or matrix) and previous output with a set of filters. Each produces one feature map in the output data according to the gated recurrent unit (GRU) unit formula detailed below. You may provide the GRU unit with a single input or a sequence of inputs. The layer also supports p-norm gating.


// A description of a gated recurrent unit block or layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGRUDescriptor
type GRUDescriptor struct {
	RNNDescriptor
}

// GRUDescriptorFrom constructs a [GRUDescriptor] from an unsafe.Pointer.
//
// A description of a gated recurrent unit block or layer.
func GRUDescriptorFrom(ptr unsafe.Pointer) GRUDescriptor {
	return GRUDescriptor{
		RNNDescriptor: RNNDescriptorFrom(ptr),
	}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865715-creategrudescriptor
func (gc _GRUDescriptorClass) CreateGRUDescriptor() {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("createGRUDescriptor"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865715-creategrudescriptorwithinputfeat
func (gc _GRUDescriptorClass) CreateGRUDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("createGRUDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865690-inputgateinputweights
func (g_ GRUDescriptor) InputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("inputGateInputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865690-inputgateinputweights
func (g_ GRUDescriptor) SetInputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInputGateInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865695-recurrentgaterecurrentweights
func (g_ GRUDescriptor) RecurrentGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("recurrentGateRecurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865695-recurrentgaterecurrentweights
func (g_ GRUDescriptor) SetRecurrentGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRecurrentGateRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865699-outputgaterecurrentweights
func (g_ GRUDescriptor) OutputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("outputGateRecurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865699-outputgaterecurrentweights
func (g_ GRUDescriptor) SetOutputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputGateRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865719-recurrentgateinputweights
func (g_ GRUDescriptor) RecurrentGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("recurrentGateInputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865719-recurrentgateinputweights
func (g_ GRUDescriptor) SetRecurrentGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRecurrentGateInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865722-outputgateinputweights
func (g_ GRUDescriptor) OutputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("outputGateInputWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865722-outputgateinputweights
func (g_ GRUDescriptor) SetOutputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputGateInputWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865724-inputgaterecurrentweights
func (g_ GRUDescriptor) InputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("inputGateRecurrentWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865724-inputgaterecurrentweights
func (g_ GRUDescriptor) SetInputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInputGateRecurrentWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2873332-gatepnormvalue
func (g_ GRUDescriptor) GatePnormValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("gatePnormValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2873332-gatepnormvalue
func (g_ GRUDescriptor) SetGatePnormValue(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGatePnormValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878270-outputgateinputgateweights
func (g_ GRUDescriptor) OutputGateInputGateWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("outputGateInputGateWeights"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878270-outputgateinputgateweights
func (g_ GRUDescriptor) SetOutputGateInputGateWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputGateInputGateWeights:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878271-flipoutputgates
func (g_ GRUDescriptor) FlipOutputGates() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("flipOutputGates"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878271-flipoutputgates
func (g_ GRUDescriptor) SetFlipOutputGates(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setFlipOutputGates:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (g_ GRUDescriptor) InputFeatureChannels() int {
	rv := objc.Send[int](g_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (g_ GRUDescriptor) SetInputFeatureChannels(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (g_ GRUDescriptor) OutputFeatureChannels() int {
	rv := objc.Send[int](g_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (g_ GRUDescriptor) SetOutputFeatureChannels(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}








