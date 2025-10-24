// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCTensor */


/* debug [class_header]: Header for MLCTensor */
// The class instance for the [CTensor] class.
var (
	CTensorClass     _CTensorClass
	CTensorClassOnce sync.Once
)

func getCTensorClass() _CTensorClass {
	CTensorClassOnce.Do(func() {
		CTensorClass = _CTensorClass{objc.GetClass("MLCTensor")}
	})
	return CTensorClass
}

type _CTensorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CTensor */
// An interface definition for the [CTensor] class.
type ICTensor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CTensor */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	Descriptor() IMLCTensorDescriptor
	Device() IMLCDevice
	HasValidNumerics() bool
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	OptimizerData() []CTensorData
	OptimizerDeviceData() []CTensorOptimizerDeviceData
	TensorID() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CTensor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CTensor */
// Alloc allocates a new instance without initialization.
func (cc _CTensorClass) Alloc() CTensor {
	rv := objc.Send[CTensor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CTensorClass) New() CTensor {
	rv := objc.Send[CTensor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTensor) Init() CTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTensor) Autorelease() CTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTensor creates a new CTensor instance.
func NewCTensor() CTensor {
	return getCTensorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CTensor */
// The data object you use throughout the framework.
//
// Create a tensor with or without data. For example, create a tensor with data for weights used by convolution or mean, variance, beta, and gamma parameters with batch normalization. Create a tensor without data to use as an input tensor when you build a graph.


// The data object you use throughout the framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor
type CTensor struct {
	objectivec.Object
}

// CTensorFrom constructs a [CTensor] from an unsafe.Pointer.
//
// The data object you use throughout the framework.
func CTensorFrom(ptr unsafe.Pointer) CTensor {
	return CTensor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CTensor */

// Creates a tensor without data, using the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:)
func NewCTensorWithDescriptor(tensorDescriptor IMLCTensorDescriptor) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:"), tensorDescriptor)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithDescriptor */


// Creates a tensor with the descriptor and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:data:)
func NewCTensorWithDescriptorData(tensorDescriptor IMLCTensorDescriptor, data IMLCTensorData) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:data:"), tensorDescriptor, data)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithDescriptorData */


// Creates a tensor with the descriptor and scalar value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:fillWithData:)
func NewCTensorWithDescriptorFillWithData(tensorDescriptor IMLCTensorDescriptor, fillData objc.IObject /* cross-framework: NSNumber */) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:fillWithData:"), tensorDescriptor, fillData)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithDescriptorFillWithData */


// Creates a tensor with the descriptor and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:randomInitializerType:)
func NewCTensorWithDescriptorRandomInitializerType(tensorDescriptor IMLCTensorDescriptor, randomInitializerType CRandomInitializerType) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:randomInitializerType:"), tensorDescriptor, randomInitializerType)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithDescriptorRandomInitializerType */


// Creates a tensor without data, with the sequence length, number of feature channels, and batch size you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:)
func NewCTensorWithSequenceLengthFeatureChannelCountBatchSize(sequenceLength uint, featureChannelCount uint, batchSize uint) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:"), sequenceLength, featureChannelCount, batchSize)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithSequenceLengthFeatureChannelCountBatchSize */


// Creates a tensor with the sequence length, number of feature channels, batch size, and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:data:)
func NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeData(sequenceLength uint, featureChannelCount uint, batchSize uint, data IMLCTensorData) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:data:"), sequenceLength, featureChannelCount, batchSize, data)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeData */


// Creates a tensor with the sequence length, number of feature channels, batch size, and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:randomInitializerType:)
func NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType(sequenceLength uint, featureChannelCount uint, batchSize uint, randomInitializerType CRandomInitializerType) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:randomInitializerType:"), sequenceLength, featureChannelCount, batchSize, randomInitializerType)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType */


// Creates a tensor without data, with the sizes and number of feature channels you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSize(width uint, height uint, featureChannelCount uint, batchSize uint) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:"), width, height, featureChannelCount, batchSize)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithWidthHeightFeatureChannelCountBatchSize */


// Creates a tensor with the sizes, number of feature channels, and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeData(width uint, height uint, featureChannelCount uint, batchSize uint, data IMLCTensorData) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:"), width, height, featureChannelCount, batchSize, data)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithWidthHeightFeatureChannelCountBatchSizeData */


// Creates a tensor with the sizes, number of feature channels, data, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:dataType:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, data IMLCTensorData, dataType CDataType) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:dataType:"), width, height, featureChannelCount, batchSize, data, dataType)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType */


// Creates a tensor with the sizes and number of feature channels, and filled with the data and type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:fillWithData:dataType:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, fillData float32, dataType CDataType) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:fillWithData:dataType:"), width, height, featureChannelCount, batchSize, fillData, dataType)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType */


// Creates a tensor with the sizes, number of feature channels, and random data using the random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:randomInitializerType:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType(width uint, height uint, featureChannelCount uint, batchSize uint, randomInitializerType CRandomInitializerType) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:randomInitializerType:"), width, height, featureChannelCount, batchSize, randomInitializerType)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CTensor */

// Creates a tensor without data, using the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:)
func (cc _CTensorClass) TensorWithDescriptor(tensorDescriptor IMLCTensorDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:"), tensorDescriptor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithDescriptor) */


// Creates a tensor with the descriptor and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:data:)
func (cc _CTensorClass) TensorWithDescriptorData(tensorDescriptor IMLCTensorDescriptor, data IMLCTensorData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:data:"), tensorDescriptor, data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithDescriptorData) */


// Creates a tensor with the descriptor and scalar value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:fillWithData:)
func (cc _CTensorClass) TensorWithDescriptorFillWithData(tensorDescriptor IMLCTensorDescriptor, fillData objc.IObject /* cross-framework: NSNumber */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:fillWithData:"), tensorDescriptor, fillData)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithDescriptorFillWithData) */


// Creates a tensor with the descriptor and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:randomInitializerType:)
func (cc _CTensorClass) TensorWithDescriptorRandomInitializerType(tensorDescriptor IMLCTensorDescriptor, randomInitializerType CRandomInitializerType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:randomInitializerType:"), tensorDescriptor, randomInitializerType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithDescriptorRandomInitializerType) */


// Creates a tensor without data, with the sequence length, number of feature channels, and batch size you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:)
func (cc _CTensorClass) TensorWithSequenceLengthFeatureChannelCountBatchSize(sequenceLength uint, featureChannelCount uint, batchSize uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:"), sequenceLength, featureChannelCount, batchSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithSequenceLengthFeatureChannelCountBatchSize) */


// Creates a tensor with the sequence length, number of feature channels, batch size, and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:data:)
func (cc _CTensorClass) TensorWithSequenceLengthFeatureChannelCountBatchSizeData(sequenceLength uint, featureChannelCount uint, batchSize uint, data IMLCTensorData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:data:"), sequenceLength, featureChannelCount, batchSize, data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithSequenceLengthFeatureChannelCountBatchSizeData) */


// Creates a tensor with the sequence length, number of feature channels, batch size, and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:randomInitializerType:)
func (cc _CTensorClass) TensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType(sequenceLength uint, featureChannelCount uint, batchSize uint, randomInitializerType CRandomInitializerType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:randomInitializerType:"), sequenceLength, featureChannelCount, batchSize, randomInitializerType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType) */


// Creates a tensor without data, with the sizes and number of feature channels you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSize(width uint, height uint, featureChannelCount uint, batchSize uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:"), width, height, featureChannelCount, batchSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithWidthHeightFeatureChannelCountBatchSize) */


// Creates a tensor with the sizes, number of feature channels, and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeData(width uint, height uint, featureChannelCount uint, batchSize uint, data IMLCTensorData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:"), width, height, featureChannelCount, batchSize, data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithWidthHeightFeatureChannelCountBatchSizeData) */


// Creates a tensor with the sizes, number of feature channels, data, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:dataType:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, data IMLCTensorData, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:dataType:"), width, height, featureChannelCount, batchSize, data, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType) */


// Creates a tensor with the sizes and number of feature channels, and filled with the data and type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:fillWithData:dataType:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, fillData float32, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:fillWithData:dataType:"), width, height, featureChannelCount, batchSize, fillData, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType) */


// Creates a tensor with the sizes, number of feature channels, and random data using the random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:randomInitializerType:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType(width uint, height uint, featureChannelCount uint, batchSize uint, randomInitializerType CRandomInitializerType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:randomInitializerType:"), width, height, featureChannelCount, batchSize, randomInitializerType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType) */


// Creates a tensor with the sequence lengths, sorting indicator, number of feature channels, batch size, and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:data:
func (cc _CTensorClass) TensorWithSequenceLengthsSortedSequencesFeatureChannelCountBatchSizeData(sequenceLengths []foundation.Number, sortedSequences bool, featureChannelCount uint, batchSize uint, data IMLCTensorData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:data:"), sequenceLengths, sortedSequences, featureChannelCount, batchSize, data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithSequenceLengthsSortedSequencesFeatureChannelCountBatchSizeData) */


// Creates a tensor with the sequence lengths, sorting indicator, number of feature channels, batch size, and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:randomInitializerType:
func (cc _CTensorClass) TensorWithSequenceLengthsSortedSequencesFeatureChannelCountBatchSizeRandomInitializerType(sequenceLengths []foundation.Number, sortedSequences bool, featureChannelCount uint, batchSize uint, randomInitializerType CRandomInitializerType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:randomInitializerType:"), sequenceLengths, sortedSequences, featureChannelCount, batchSize, randomInitializerType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithSequenceLengthsSortedSequencesFeatureChannelCountBatchSizeRandomInitializerType) */


// Creates a tensor without data, with the shape you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:
func (cc _CTensorClass) TensorWithShape(shape []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:"), shape)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithShape) */


// Creates a tensor with the shape, data, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:data:dataType:
func (cc _CTensorClass) TensorWithShapeDataDataType(shape []foundation.Number, data IMLCTensorData, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:data:dataType:"), shape, data, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithShapeDataDataType) */


// Creates a tensor without data, with the shape and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:dataType:
func (cc _CTensorClass) TensorWithShapeDataType(shape []foundation.Number, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:dataType:"), shape, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithShapeDataType) */


// Creates a tensor with the shape, scalar value, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:fillWithData:dataType:
func (cc _CTensorClass) TensorWithShapeFillWithDataDataType(shape []foundation.Number, fillData objc.IObject /* cross-framework: NSNumber */, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:fillWithData:dataType:"), shape, fillData, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithShapeFillWithDataDataType) */


// Creates a tensor with the shape and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:randomInitializerType:
func (cc _CTensorClass) TensorWithShapeRandomInitializerType(shape []foundation.Number, randomInitializerType CRandomInitializerType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:randomInitializerType:"), shape, randomInitializerType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithShapeRandomInitializerType) */


// Creates a tensor with the shape, random initializer, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:randomInitializerType:dataType:
func (cc _CTensorClass) TensorWithShapeRandomInitializerTypeDataType(shape []foundation.Number, randomInitializerType CRandomInitializerType, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:randomInitializerType:dataType:"), shape, randomInitializerType, dataType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TensorWithShapeRandomInitializerTypeDataType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CTensor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CTensor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CTensor */

// The tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/data
func (c_ CTensor) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The configuration object you use to create a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/descriptor
func (c_ CTensor) Descriptor() IMLCTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}/* debug [instance_properties/getter]: descriptor */


// The device associated with this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/device
func (c_ CTensor) Device() IMLCDevice {
	rv := objc.Send[CDevice](c_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// A Boolean that indicates whether a tensor contains NaN or INF values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/hasValidNumerics
func (c_ CTensor) HasValidNumerics() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasValidNumerics"))
	return rv
}/* debug [instance_properties/getter]: hasValidNumerics */


// A string that identifes this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/label
func (c_ CTensor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A string that identifes this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/label
func (c_ CTensor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// An array that contains optimizer buffers you specify when you create a tensor parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/optimizerData
func (c_ CTensor) OptimizerData() []CTensorData {
	rv := objc.Send[[]CTensorData](c_.ID, objc.Sel("optimizerData"))
	return rv
}/* debug [instance_properties/getter]: optimizerData */


// An array that contains the device optimizer buffers you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/optimizerDeviceData
func (c_ CTensor) OptimizerDeviceData() []CTensorOptimizerDeviceData {
	rv := objc.Send[[]CTensorOptimizerDeviceData](c_.ID, objc.Sel("optimizerDeviceData"))
	return rv
}/* debug [instance_properties/getter]: optimizerDeviceData */


// A number that uniquely identifies the tensor, which the framework assigns when it creates a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorID
func (c_ CTensor) TensorID() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("tensorID"))
	return rv
}/* debug [instance_properties/getter]: tensorID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCTensor */


