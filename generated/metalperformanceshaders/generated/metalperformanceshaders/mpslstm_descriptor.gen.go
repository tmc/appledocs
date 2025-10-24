// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSLSTMDescriptor */


/* debug [class_header]: Header for MPSLSTMDescriptor */
// The class instance for the [LSTMDescriptor] class.
var (
	LSTMDescriptorClass     _LSTMDescriptorClass
	LSTMDescriptorClassOnce sync.Once
)

func getLSTMDescriptorClass() _LSTMDescriptorClass {
	LSTMDescriptorClassOnce.Do(func() {
		LSTMDescriptorClass = _LSTMDescriptorClass{objc.GetClass("MPSLSTMDescriptor")}
	})
	return LSTMDescriptorClass
}

type _LSTMDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LSTMDescriptor */
// An interface definition for the [LSTMDescriptor] class.
type ILSTMDescriptor interface {
	IRNNDescriptor
	
/* debug [class_interface_properties]: Properties for LSTMDescriptor */
	// properties:
	CellGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetCellGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	CellGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */
	SetCellGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */)
	InputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetInputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	OutputGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */
	SetOutputGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */)
	ForgetGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */
	SetForgetGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */)
	CellToOutputNeuronParamB() objectivec.IObject
	SetCellToOutputNeuronParamB(value objectivec.IObject)
	OutputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetOutputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	MemoryWeightsAreDiagonal() objectivec.IObject
	SetMemoryWeightsAreDiagonal(value objectivec.IObject)
	InputGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */
	SetInputGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */)
	ForgetGateInputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetForgetGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	ForgetGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetForgetGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	CellToOutputNeuronType() CNNNeuronType get set /* not a class type */
	SetCellToOutputNeuronType(value CNNNeuronType get set /* not a class type */)
	CellGateInputWeights() CNNConvolutionDataSource get set /* not a class type */
	SetCellGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */)
	CellToOutputNeuronParamA() objectivec.IObject
	SetCellToOutputNeuronParamA(value objectivec.IObject)
	InputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetInputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	OutputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */
	SetOutputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */)
	CellToOutputNeuronParamC() objectivec.IObject
	SetCellToOutputNeuronParamC(value objectivec.IObject)
	InputFeatureChannels() int
	SetInputFeatureChannels(value int)
	OutputFeatureChannels() int
	SetOutputFeatureChannels(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LSTMDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LSTMDescriptor */
// Alloc allocates a new instance without initialization.
func (lc _LSTMDescriptorClass) Alloc() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LSTMDescriptorClass) New() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LSTMDescriptor) Init() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LSTMDescriptor) Autorelease() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLSTMDescriptor creates a new LSTMDescriptor instance.
func NewLSTMDescriptor() LSTMDescriptor {
	return getLSTMDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LSTMDescriptor */
// A description of a long short-term memory block or layer.
//
// The recurrent neural network (RNN) layer initialized with transforms the input data (image or matrix), the memory cell data, and previous output with a set of filters. Each produces one feature map in the output data and memory cell according to the long short-term memory (LSTM) formula detailed below. You may provide the LSTM unit with a single input or a sequence of inputs.


// A description of a long short-term memory block or layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor
type LSTMDescriptor struct {
	RNNDescriptor
}

// LSTMDescriptorFrom constructs a [LSTMDescriptor] from an unsafe.Pointer.
//
// A description of a long short-term memory block or layer.
func LSTMDescriptorFrom(ptr unsafe.Pointer) LSTMDescriptor {
	return LSTMDescriptor{
		RNNDescriptor: RNNDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LSTMDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LSTMDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865681-createlstmdescriptor
func (lc _LSTMDescriptorClass) CreateLSTMDescriptor() {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("createLSTMDescriptor"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateLSTMDescriptor) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865681-createlstmdescriptorwithinputfea
func (lc _LSTMDescriptorClass) CreateLSTMDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("createLSTMDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CreateLSTMDescriptorWithInputFeatureChannelsOutputFeatureChannels) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LSTMDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LSTMDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LSTMDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865679-cellgaterecurrentweights
func (l_ LSTMDescriptor) CellGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("cellGateRecurrentWeights"))
	return rv
}/* debug [instance_properties/getter]: cellGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865679-cellgaterecurrentweights
func (l_ LSTMDescriptor) SetCellGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateRecurrentWeights:"), value)
}/* debug [instance_properties/setter]: cellGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865683-cellgatememoryweights
func (l_ LSTMDescriptor) CellGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("cellGateMemoryWeights"))
	return rv
}/* debug [instance_properties/getter]: cellGateMemoryWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865683-cellgatememoryweights
func (l_ LSTMDescriptor) SetCellGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateMemoryWeights:"), value)
}/* debug [instance_properties/setter]: cellGateMemoryWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865684-inputgateinputweights
func (l_ LSTMDescriptor) InputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("inputGateInputWeights"))
	return rv
}/* debug [instance_properties/getter]: inputGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865684-inputgateinputweights
func (l_ LSTMDescriptor) SetInputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateInputWeights:"), value)
}/* debug [instance_properties/setter]: inputGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865688-outputgatememoryweights
func (l_ LSTMDescriptor) OutputGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("outputGateMemoryWeights"))
	return rv
}/* debug [instance_properties/getter]: outputGateMemoryWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865688-outputgatememoryweights
func (l_ LSTMDescriptor) SetOutputGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateMemoryWeights:"), value)
}/* debug [instance_properties/setter]: outputGateMemoryWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865689-forgetgatememoryweights
func (l_ LSTMDescriptor) ForgetGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("forgetGateMemoryWeights"))
	return rv
}/* debug [instance_properties/getter]: forgetGateMemoryWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865689-forgetgatememoryweights
func (l_ LSTMDescriptor) SetForgetGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateMemoryWeights:"), value)
}/* debug [instance_properties/setter]: forgetGateMemoryWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865694-celltooutputneuronparamb
func (l_ LSTMDescriptor) CellToOutputNeuronParamB() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("cellToOutputNeuronParamB"))
	return rv
}/* debug [instance_properties/getter]: cellToOutputNeuronParamB */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865694-celltooutputneuronparamb
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamB(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamB:"), value)
}/* debug [instance_properties/setter]: cellToOutputNeuronParamB */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865701-outputgateinputweights
func (l_ LSTMDescriptor) OutputGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("outputGateInputWeights"))
	return rv
}/* debug [instance_properties/getter]: outputGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865701-outputgateinputweights
func (l_ LSTMDescriptor) SetOutputGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateInputWeights:"), value)
}/* debug [instance_properties/setter]: outputGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865712-memoryweightsarediagonal
func (l_ LSTMDescriptor) MemoryWeightsAreDiagonal() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("memoryWeightsAreDiagonal"))
	return rv
}/* debug [instance_properties/getter]: memoryWeightsAreDiagonal */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865712-memoryweightsarediagonal
func (l_ LSTMDescriptor) SetMemoryWeightsAreDiagonal(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMemoryWeightsAreDiagonal:"), value)
}/* debug [instance_properties/setter]: memoryWeightsAreDiagonal */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865731-inputgatememoryweights
func (l_ LSTMDescriptor) InputGateMemoryWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("inputGateMemoryWeights"))
	return rv
}/* debug [instance_properties/getter]: inputGateMemoryWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865731-inputgatememoryweights
func (l_ LSTMDescriptor) SetInputGateMemoryWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateMemoryWeights:"), value)
}/* debug [instance_properties/setter]: inputGateMemoryWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865734-forgetgateinputweights
func (l_ LSTMDescriptor) ForgetGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("forgetGateInputWeights"))
	return rv
}/* debug [instance_properties/getter]: forgetGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865734-forgetgateinputweights
func (l_ LSTMDescriptor) SetForgetGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateInputWeights:"), value)
}/* debug [instance_properties/setter]: forgetGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865735-forgetgaterecurrentweights
func (l_ LSTMDescriptor) ForgetGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("forgetGateRecurrentWeights"))
	return rv
}/* debug [instance_properties/getter]: forgetGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865735-forgetgaterecurrentweights
func (l_ LSTMDescriptor) SetForgetGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForgetGateRecurrentWeights:"), value)
}/* debug [instance_properties/setter]: forgetGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865736-celltooutputneurontype
func (l_ LSTMDescriptor) CellToOutputNeuronType() CNNNeuronType get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("cellToOutputNeuronType"))
	return rv
}/* debug [instance_properties/getter]: cellToOutputNeuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865736-celltooutputneurontype
func (l_ LSTMDescriptor) SetCellToOutputNeuronType(value CNNNeuronType get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronType:"), value)
}/* debug [instance_properties/setter]: cellToOutputNeuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865741-cellgateinputweights
func (l_ LSTMDescriptor) CellGateInputWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("cellGateInputWeights"))
	return rv
}/* debug [instance_properties/getter]: cellGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865741-cellgateinputweights
func (l_ LSTMDescriptor) SetCellGateInputWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellGateInputWeights:"), value)
}/* debug [instance_properties/setter]: cellGateInputWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865744-celltooutputneuronparama
func (l_ LSTMDescriptor) CellToOutputNeuronParamA() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("cellToOutputNeuronParamA"))
	return rv
}/* debug [instance_properties/getter]: cellToOutputNeuronParamA */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865744-celltooutputneuronparama
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamA(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamA:"), value)
}/* debug [instance_properties/setter]: cellToOutputNeuronParamA */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865747-inputgaterecurrentweights
func (l_ LSTMDescriptor) InputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("inputGateRecurrentWeights"))
	return rv
}/* debug [instance_properties/getter]: inputGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865747-inputgaterecurrentweights
func (l_ LSTMDescriptor) SetInputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputGateRecurrentWeights:"), value)
}/* debug [instance_properties/setter]: inputGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865750-outputgaterecurrentweights
func (l_ LSTMDescriptor) OutputGateRecurrentWeights() CNNConvolutionDataSource get set /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("outputGateRecurrentWeights"))
	return rv
}/* debug [instance_properties/getter]: outputGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865750-outputgaterecurrentweights
func (l_ LSTMDescriptor) SetOutputGateRecurrentWeights(value CNNConvolutionDataSource get set /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputGateRecurrentWeights:"), value)
}/* debug [instance_properties/setter]: outputGateRecurrentWeights */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2935551-celltooutputneuronparamc
func (l_ LSTMDescriptor) CellToOutputNeuronParamC() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("cellToOutputNeuronParamC"))
	return rv
}/* debug [instance_properties/getter]: cellToOutputNeuronParamC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2935551-celltooutputneuronparamc
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamC(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCellToOutputNeuronParamC:"), value)
}/* debug [instance_properties/setter]: cellToOutputNeuronParamC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (l_ LSTMDescriptor) InputFeatureChannels() int {
	rv := objc.Send[int](l_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/inputfeaturechannels
func (l_ LSTMDescriptor) SetInputFeatureChannels(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (l_ LSTMDescriptor) OutputFeatureChannels() int {
	rv := objc.Send[int](l_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/outputfeaturechannels
func (l_ LSTMDescriptor) SetOutputFeatureChannels(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: outputFeatureChannels */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSLSTMDescriptor */



