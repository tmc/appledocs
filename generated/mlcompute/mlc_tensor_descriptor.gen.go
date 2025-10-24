// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCTensorDescriptor */


/* debug [class_header]: Header for MLCTensorDescriptor */
// The class instance for the [CTensorDescriptor] class.
var (
	CTensorDescriptorClass     _CTensorDescriptorClass
	CTensorDescriptorClassOnce sync.Once
)

func getCTensorDescriptorClass() _CTensorDescriptorClass {
	CTensorDescriptorClassOnce.Do(func() {
		CTensorDescriptorClass = _CTensorDescriptorClass{objc.GetClass("MLCTensorDescriptor")}
	})
	return CTensorDescriptorClass
}

type _CTensorDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CTensorDescriptor */
// An interface definition for the [CTensorDescriptor] class.
type ICTensorDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CTensorDescriptor */
	// properties:
	BatchSizePerSequenceStep() []foundation.Number
	DataType() CDataType
	DimensionCount() uint
	SequenceLengths() []foundation.Number
	Shape() []foundation.Number
	SortedSequences() bool
	Stride() []foundation.Number
	TensorAllocationSizeInBytes() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CTensorDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CTensorDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CTensorDescriptorClass) Alloc() CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CTensorDescriptorClass) New() CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTensorDescriptor) Init() CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTensorDescriptor) Autorelease() CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTensorDescriptor creates a new CTensorDescriptor instance.
func NewCTensorDescriptor() CTensorDescriptor {
	return getCTensorDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CTensorDescriptor */
// A configuration object you use to create a tensor.
//
// This class contains the mathematical properties of a tensor, such as data type and shape. It also includes initializers that help you create a tensor descriptor for common use cases, such as convolutional neural networks and recurrent neural networks.


// A configuration object you use to create a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor
type CTensorDescriptor struct {
	objectivec.Object
}

// CTensorDescriptorFrom constructs a [CTensorDescriptor] from an unsafe.Pointer.
//
// A configuration object you use to create a tensor.
func CTensorDescriptorFrom(ptr unsafe.Pointer) CTensorDescriptor {
	return CTensorDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CTensorDescriptor */

// Creates a tensor descriptor with the number of feature channels and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/init(convolutionBiasesWithFeatureChannelCount:dataType:)
func NewCTensorDescriptorConvolutionBiasesDescriptorWithFeatureChannelCountDataType(featureChannelCount uint, dataType CDataType) CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](objc.ID(getCTensorDescriptorClass().class), objc.Sel("convolutionBiasesDescriptorWithFeatureChannelCount:dataType:"), featureChannelCount, dataType)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorDescriptorConvolutionBiasesDescriptorWithFeatureChannelCountDataType */


// Creates a tensor descriptor with the number of feature channels and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/init(convolutionWeightsWithInputFeatureChannelCount:outputFeatureChannelCount:dataType:)
func NewCTensorDescriptorConvolutionWeightsDescriptorWithInputFeatureChannelCountOutputFeatureChannelCountDataType(inputFeatureChannelCount uint, outputFeatureChannelCount uint, dataType CDataType) CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](objc.ID(getCTensorDescriptorClass().class), objc.Sel("convolutionWeightsDescriptorWithInputFeatureChannelCount:outputFeatureChannelCount:dataType:"), inputFeatureChannelCount, outputFeatureChannelCount, dataType)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorDescriptorConvolutionWeightsDescriptorWithInputFeatureChannelCountOutputFeatureChannelCountDataType */


// Creates a tensor descriptor with the sizing, number of feature channels, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/init(convolutionWeightsWithWidth:height:inputFeatureChannelCount:outputFeatureChannelCount:dataType:)
func NewCTensorDescriptorConvolutionWeightsDescriptorWithWidthHeightInputFeatureChannelCountOutputFeatureChannelCountDataType(width uint, height uint, inputFeatureChannelCount uint, outputFeatureChannelCount uint, dataType CDataType) CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](objc.ID(getCTensorDescriptorClass().class), objc.Sel("convolutionWeightsDescriptorWithWidth:height:inputFeatureChannelCount:outputFeatureChannelCount:dataType:"), width, height, inputFeatureChannelCount, outputFeatureChannelCount, dataType)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorDescriptorConvolutionWeightsDescriptorWithWidthHeightInputFeatureChannelCountOutputFeatureChannelCountDataType */


// Creates a tensor descriptor with the width and height, number of feature channels, and batch size you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/init(width:height:featureChannelCount:batchSize:)
func NewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSize(width uint, height uint, featureChannels uint, batchSize uint) CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](objc.ID(getCTensorDescriptorClass().class), objc.Sel("descriptorWithWidth:height:featureChannelCount:batchSize:"), width, height, featureChannels, batchSize)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSize */


// Creates a tensor descriptor with the width and height, number of feature channels, batch size, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/init(width:height:featureChannelCount:batchSize:dataType:)
func NewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSizeDataType(width uint, height uint, featureChannelCount uint, batchSize uint, dataType CDataType) CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](objc.ID(getCTensorDescriptorClass().class), objc.Sel("descriptorWithWidth:height:featureChannelCount:batchSize:dataType:"), width, height, featureChannelCount, batchSize, dataType)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorDescriptorWithWidthHeightFeatureChannelCountBatchSizeDataType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CTensorDescriptor */

// Creates a tensor descriptor with the shape and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/descriptorWithShape:dataType:
func (cc _CTensorDescriptorClass) DescriptorWithShapeDataType(shape []foundation.Number, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithShape:dataType:"), shape, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithShapeDataType) */


// Creates a tensor descriptor with the shape, variable sequence lengths, sorting indicator, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/descriptorWithShape:sequenceLengths:sortedSequences:dataType:
func (cc _CTensorDescriptorClass) DescriptorWithShapeSequenceLengthsSortedSequencesDataType(shape []foundation.Number, sequenceLengths []foundation.Number, sortedSequences bool, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithShape:sequenceLengths:sortedSequences:dataType:"), shape, sequenceLengths, sortedSequences, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithShapeSequenceLengthsSortedSequencesDataType) */


// Creates a tensor descriptor with the number of feature channels and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/init(convolutionBiasesWithFeatureChannelCount:dataType:)
func (cc _CTensorDescriptorClass) ConvolutionBiasesDescriptorWithFeatureChannelCountDataType(featureChannelCount uint, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("convolutionBiasesDescriptorWithFeatureChannelCount:dataType:"), featureChannelCount, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConvolutionBiasesDescriptorWithFeatureChannelCountDataType) */


// Creates a tensor descriptor with the number of feature channels and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/init(convolutionWeightsWithInputFeatureChannelCount:outputFeatureChannelCount:dataType:)
func (cc _CTensorDescriptorClass) ConvolutionWeightsDescriptorWithInputFeatureChannelCountOutputFeatureChannelCountDataType(inputFeatureChannelCount uint, outputFeatureChannelCount uint, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("convolutionWeightsDescriptorWithInputFeatureChannelCount:outputFeatureChannelCount:dataType:"), inputFeatureChannelCount, outputFeatureChannelCount, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConvolutionWeightsDescriptorWithInputFeatureChannelCountOutputFeatureChannelCountDataType) */


// Creates a tensor descriptor with the sizing, number of feature channels, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/init(convolutionWeightsWithWidth:height:inputFeatureChannelCount:outputFeatureChannelCount:dataType:)
func (cc _CTensorDescriptorClass) ConvolutionWeightsDescriptorWithWidthHeightInputFeatureChannelCountOutputFeatureChannelCountDataType(width uint, height uint, inputFeatureChannelCount uint, outputFeatureChannelCount uint, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("convolutionWeightsDescriptorWithWidth:height:inputFeatureChannelCount:outputFeatureChannelCount:dataType:"), width, height, inputFeatureChannelCount, outputFeatureChannelCount, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConvolutionWeightsDescriptorWithWidthHeightInputFeatureChannelCountOutputFeatureChannelCountDataType) */


// Creates a tensor descriptor with the width and height, number of feature channels, and batch size you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/init(width:height:featureChannelCount:batchSize:)
func (cc _CTensorDescriptorClass) DescriptorWithWidthHeightFeatureChannelCountBatchSize(width uint, height uint, featureChannels uint, batchSize uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithWidth:height:featureChannelCount:batchSize:"), width, height, featureChannels, batchSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithWidthHeightFeatureChannelCountBatchSize) */


// Creates a tensor descriptor with the width and height, number of feature channels, batch size, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/init(width:height:featureChannelCount:batchSize:dataType:)
func (cc _CTensorDescriptorClass) DescriptorWithWidthHeightFeatureChannelCountBatchSizeDataType(width uint, height uint, featureChannelCount uint, batchSize uint, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithWidth:height:featureChannelCount:batchSize:dataType:"), width, height, featureChannelCount, batchSize, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithWidthHeightFeatureChannelCountBatchSizeDataType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CTensorDescriptor */

// The maximum number of tensor dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/maxTensorDimensions
func (cc _CTensorDescriptorClass) MaxTensorDimensions() uint {
	rv := objc.Send[uint](objc.ID(cc.class), objc.Sel("maxTensorDimensions"))
	return rv
}/* debug [class_properties_class/property]: maxTensorDimensions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CTensorDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CTensorDescriptor */

// The batch size for each sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/batchSizePerSequenceStep-55mp8
func (c_ CTensorDescriptor) BatchSizePerSequenceStep() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("batchSizePerSequenceStep"))
	return rv
}/* debug [instance_properties/getter]: batchSizePerSequenceStep */


// The tensor data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/dataType
func (c_ CTensorDescriptor) DataType() CDataType {
	rv := objc.Send[CDataType](c_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The number of dimensions in the tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/dimensionCount
func (c_ CTensorDescriptor) DimensionCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dimensionCount"))
	return rv
}/* debug [instance_properties/getter]: dimensionCount */


// The maximum number of tensor dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/maxTensorDimensions
func (c_ CTensorDescriptor) MaxTensorDimensions() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxTensorDimensions"))
	return rv
}/* debug [instance_properties/getter]: maxTensorDimensions */


// An array that contains the variable lengths of sequences stored in the tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/sequenceLengths-3ntsa
func (c_ CTensorDescriptor) SequenceLengths() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("sequenceLengths"))
	return rv
}/* debug [instance_properties/getter]: sequenceLengths */


// An array that contains the size in each dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/shape-91vng
func (c_ CTensorDescriptor) Shape() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */


// A Boolean that indicates whether you provided the sequence lengths sorted in descending order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/sortedSequences
func (c_ CTensorDescriptor) SortedSequences() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sortedSequences"))
	return rv
}/* debug [instance_properties/getter]: sortedSequences */


// An array that contains the stride, in bytes, in each dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/stride-3ydik
func (c_ CTensorDescriptor) Stride() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("stride"))
	return rv
}/* debug [instance_properties/getter]: stride */


// The allocation size, in bytes, for a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor/tensorAllocationSizeInBytes
func (c_ CTensorDescriptor) TensorAllocationSizeInBytes() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("tensorAllocationSizeInBytes"))
	return rv
}/* debug [instance_properties/getter]: tensorAllocationSizeInBytes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCTensorDescriptor */


