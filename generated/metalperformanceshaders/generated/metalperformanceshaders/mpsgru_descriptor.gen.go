// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSGRUDescriptor */


/* debug [class_header]: Header for MPSGRUDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GRUDescriptor */
// An interface definition for the [GRUDescriptor] class.
type IGRUDescriptor interface {
	IRNNDescriptor
	
/* debug [class_interface_properties]: Properties for GRUDescriptor */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GRUDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GRUDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GRUDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GRUDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GRUDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865715-creategrudescriptor
func (gc _GRUDescriptorClass) CreateGRUDescriptor() {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("createGRUDescriptor"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateGRUDescriptor) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865715-creategrudescriptorwithinputfeat
func (gc _GRUDescriptorClass) CreateGRUDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(gc.class), objc.Sel("createGRUDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateGRUDescriptorWithInputFeatureChannelsOutputFeatureChannels) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GRUDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GRUDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GRUDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865690-inputgateinputweights
func (g_ GRUDescriptor) InputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("inputGateInputWeights"))
	return rv
}/* debug [instance_properties/getter]: inputGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865690-inputgateinputweights
func (g_ GRUDescriptor) SetInputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInputGateInputWeights:"), value)
}/* debug [instance_properties/setter]: inputGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865695-recurrentgaterecurrentweights
func (g_ GRUDescriptor) RecurrentGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("recurrentGateRecurrentWeights"))
	return rv
}/* debug [instance_properties/getter]: recurrentGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865695-recurrentgaterecurrentweights
func (g_ GRUDescriptor) SetRecurrentGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRecurrentGateRecurrentWeights:"), value)
}/* debug [instance_properties/setter]: recurrentGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865699-outputgaterecurrentweights
func (g_ GRUDescriptor) OutputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("outputGateRecurrentWeights"))
	return rv
}/* debug [instance_properties/getter]: outputGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865699-outputgaterecurrentweights
func (g_ GRUDescriptor) SetOutputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputGateRecurrentWeights:"), value)
}/* debug [instance_properties/setter]: outputGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865719-recurrentgateinputweights
func (g_ GRUDescriptor) RecurrentGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("recurrentGateInputWeights"))
	return rv
}/* debug [instance_properties/getter]: recurrentGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865719-recurrentgateinputweights
func (g_ GRUDescriptor) SetRecurrentGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRecurrentGateInputWeights:"), value)
}/* debug [instance_properties/setter]: recurrentGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865722-outputgateinputweights
func (g_ GRUDescriptor) OutputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("outputGateInputWeights"))
	return rv
}/* debug [instance_properties/getter]: outputGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865722-outputgateinputweights
func (g_ GRUDescriptor) SetOutputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputGateInputWeights:"), value)
}/* debug [instance_properties/setter]: outputGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865724-inputgaterecurrentweights
func (g_ GRUDescriptor) InputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("inputGateRecurrentWeights"))
	return rv
}/* debug [instance_properties/getter]: inputGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865724-inputgaterecurrentweights
func (g_ GRUDescriptor) SetInputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInputGateRecurrentWeights:"), value)
}/* debug [instance_properties/setter]: inputGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2873332-gatepnormvalue
func (g_ GRUDescriptor) GatePnormValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("gatePnormValue"))
	return rv
}/* debug [instance_properties/getter]: gatePnormValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2873332-gatepnormvalue
func (g_ GRUDescriptor) SetGatePnormValue(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGatePnormValue:"), value)
}/* debug [instance_properties/setter]: gatePnormValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878270-outputgateinputgateweights
func (g_ GRUDescriptor) OutputGateInputGateWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("outputGateInputGateWeights"))
	return rv
}/* debug [instance_properties/getter]: outputGateInputGateWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878270-outputgateinputgateweights
func (g_ GRUDescriptor) SetOutputGateInputGateWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputGateInputGateWeights:"), value)
}/* debug [instance_properties/setter]: outputGateInputGateWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878271-flipoutputgates
func (g_ GRUDescriptor) FlipOutputGates() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](g_.ID, objc.Sel("flipOutputGates"))
	return rv
}/* debug [instance_properties/getter]: flipOutputGates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878271-flipoutputgates
func (g_ GRUDescriptor) SetFlipOutputGates(value objectivec.IObject) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setFlipOutputGates:"), value)
}/* debug [instance_properties/setter]: flipOutputGates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (g_ GRUDescriptor) InputFeatureChannels() int {
	rv := objc.Send[int](g_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (g_ GRUDescriptor) SetInputFeatureChannels(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (g_ GRUDescriptor) OutputFeatureChannels() int {
	rv := objc.Send[int](g_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (g_ GRUDescriptor) SetOutputFeatureChannels(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: outputFeatureChannels */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGRUDescriptor */



