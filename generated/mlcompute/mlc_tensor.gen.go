// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CTensor] class.
type ICTensor interface {
	objectivec.IObject
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	Descriptor() IMLCTensorDescriptor
	Device() IMLCDevice
	HasValidNumerics() bool
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	OptimizerData() []ICTensorData
	OptimizerDeviceData() []ICTensorOptimizerDeviceData
	TensorID() uint
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (cc _CTensorClass) Alloc() CTensor {
	rv := objc.Send[CTensor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a tensor without data, using the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:)
func NewCTensorWithDescriptor(tensorDescriptor IMLCTensorDescriptor) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:"), tensorDescriptor)
	return rv
}


// Creates a tensor with the descriptor and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:data:)
func NewCTensorWithDescriptorData(tensorDescriptor IMLCTensorDescriptor, data IMLCTensorData) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:data:"), tensorDescriptor, data)
	return rv
}


// Creates a tensor with the descriptor and scalar value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:fillWithData:)
func NewCTensorWithDescriptorFillWithData(tensorDescriptor IMLCTensorDescriptor, fillData objc.IObject /* cross-framework: NSNumber */) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:fillWithData:"), tensorDescriptor, fillData)
	return rv
}


// Creates a tensor with the descriptor and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:randomInitializerType:)
func NewCTensorWithDescriptorRandomInitializerType(tensorDescriptor IMLCTensorDescriptor, randomInitializerType CRandomInitializerType) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:randomInitializerType:"), tensorDescriptor, randomInitializerType)
	return rv
}


// Creates a tensor without data, with the sequence length, number of feature channels, and batch size you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:)
func NewCTensorWithSequenceLengthFeatureChannelCountBatchSize(sequenceLength uint, featureChannelCount uint, batchSize uint) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:"), sequenceLength, featureChannelCount, batchSize)
	return rv
}


// Creates a tensor with the sequence length, number of feature channels, batch size, and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:data:)
func NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeData(sequenceLength uint, featureChannelCount uint, batchSize uint, data IMLCTensorData) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:data:"), sequenceLength, featureChannelCount, batchSize, data)
	return rv
}


// Creates a tensor with the sequence length, number of feature channels, batch size, and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:randomInitializerType:)
func NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType(sequenceLength uint, featureChannelCount uint, batchSize uint, randomInitializerType CRandomInitializerType) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:randomInitializerType:"), sequenceLength, featureChannelCount, batchSize, randomInitializerType)
	return rv
}


// Creates a tensor without data, with the sizes and number of feature channels you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSize(width uint, height uint, featureChannelCount uint, batchSize uint) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:"), width, height, featureChannelCount, batchSize)
	return rv
}


// Creates a tensor with the sizes, number of feature channels, and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeData(width uint, height uint, featureChannelCount uint, batchSize uint, data IMLCTensorData) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:"), width, height, featureChannelCount, batchSize, data)
	return rv
}


// Creates a tensor with the sizes, number of feature channels, data, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:dataType:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, data IMLCTensorData, dataType CDataType) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:dataType:"), width, height, featureChannelCount, batchSize, data, dataType)
	return rv
}


// Creates a tensor with the sizes and number of feature channels, and filled with the data and type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:fillWithData:dataType:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, fillData float32, dataType CDataType) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:fillWithData:dataType:"), width, height, featureChannelCount, batchSize, fillData, dataType)
	return rv
}


// Creates a tensor with the sizes, number of feature channels, and random data using the random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:randomInitializerType:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType(width uint, height uint, featureChannelCount uint, batchSize uint, randomInitializerType CRandomInitializerType) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:randomInitializerType:"), width, height, featureChannelCount, batchSize, randomInitializerType)
	return rv
}



// Creates a tensor without data, using the descriptor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:)
func (cc _CTensorClass) TensorWithDescriptor(tensorDescriptor IMLCTensorDescriptor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:"), tensorDescriptor)
	return rv
}


// Creates a tensor with the descriptor and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:data:)
func (cc _CTensorClass) TensorWithDescriptorData(tensorDescriptor IMLCTensorDescriptor, data IMLCTensorData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:data:"), tensorDescriptor, data)
	return rv
}


// Creates a tensor with the descriptor and scalar value you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:fillWithData:)
func (cc _CTensorClass) TensorWithDescriptorFillWithData(tensorDescriptor IMLCTensorDescriptor, fillData objc.IObject /* cross-framework: NSNumber */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:fillWithData:"), tensorDescriptor, fillData)
	return rv
}


// Creates a tensor with the descriptor and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:randomInitializerType:)
func (cc _CTensorClass) TensorWithDescriptorRandomInitializerType(tensorDescriptor IMLCTensorDescriptor, randomInitializerType CRandomInitializerType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:randomInitializerType:"), tensorDescriptor, randomInitializerType)
	return rv
}


// Creates a tensor without data, with the sequence length, number of feature channels, and batch size you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:)
func (cc _CTensorClass) TensorWithSequenceLengthFeatureChannelCountBatchSize(sequenceLength uint, featureChannelCount uint, batchSize uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:"), sequenceLength, featureChannelCount, batchSize)
	return rv
}


// Creates a tensor with the sequence length, number of feature channels, batch size, and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:data:)
func (cc _CTensorClass) TensorWithSequenceLengthFeatureChannelCountBatchSizeData(sequenceLength uint, featureChannelCount uint, batchSize uint, data IMLCTensorData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:data:"), sequenceLength, featureChannelCount, batchSize, data)
	return rv
}


// Creates a tensor with the sequence length, number of feature channels, batch size, and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:randomInitializerType:)
func (cc _CTensorClass) TensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType(sequenceLength uint, featureChannelCount uint, batchSize uint, randomInitializerType CRandomInitializerType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:randomInitializerType:"), sequenceLength, featureChannelCount, batchSize, randomInitializerType)
	return rv
}


// Creates a tensor without data, with the sizes and number of feature channels you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSize(width uint, height uint, featureChannelCount uint, batchSize uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:"), width, height, featureChannelCount, batchSize)
	return rv
}


// Creates a tensor with the sizes, number of feature channels, and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeData(width uint, height uint, featureChannelCount uint, batchSize uint, data IMLCTensorData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:"), width, height, featureChannelCount, batchSize, data)
	return rv
}


// Creates a tensor with the sizes, number of feature channels, data, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:dataType:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, data IMLCTensorData, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:dataType:"), width, height, featureChannelCount, batchSize, data, dataType)
	return rv
}


// Creates a tensor with the sizes and number of feature channels, and filled with the data and type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:fillWithData:dataType:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, fillData float32, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:fillWithData:dataType:"), width, height, featureChannelCount, batchSize, fillData, dataType)
	return rv
}


// Creates a tensor with the sizes, number of feature channels, and random data using the random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:randomInitializerType:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType(width uint, height uint, featureChannelCount uint, batchSize uint, randomInitializerType CRandomInitializerType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:randomInitializerType:"), width, height, featureChannelCount, batchSize, randomInitializerType)
	return rv
}


// Creates a tensor with the sequence lengths, sorting indicator, number of feature channels, batch size, and data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:data:
func (cc _CTensorClass) TensorWithSequenceLengthsSortedSequencesFeatureChannelCountBatchSizeData(sequenceLengths []objc.IObject /* cross-framework: Number */, sortedSequences bool, featureChannelCount uint, batchSize uint, data IMLCTensorData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:data:"), sequenceLengths, sortedSequences, featureChannelCount, batchSize, data)
	return rv
}


// Creates a tensor with the sequence lengths, sorting indicator, number of feature channels, batch size, and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:randomInitializerType:
func (cc _CTensorClass) TensorWithSequenceLengthsSortedSequencesFeatureChannelCountBatchSizeRandomInitializerType(sequenceLengths []objc.IObject /* cross-framework: Number */, sortedSequences bool, featureChannelCount uint, batchSize uint, randomInitializerType CRandomInitializerType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:randomInitializerType:"), sequenceLengths, sortedSequences, featureChannelCount, batchSize, randomInitializerType)
	return rv
}


// Creates a tensor without data, with the shape you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:
func (cc _CTensorClass) TensorWithShape(shape []objc.IObject /* cross-framework: Number */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:"), shape)
	return rv
}


// Creates a tensor with the shape, data, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:data:dataType:
func (cc _CTensorClass) TensorWithShapeDataDataType(shape []objc.IObject /* cross-framework: Number */, data IMLCTensorData, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:data:dataType:"), shape, data, dataType)
	return rv
}


// Creates a tensor without data, with the shape and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:dataType:
func (cc _CTensorClass) TensorWithShapeDataType(shape []objc.IObject /* cross-framework: Number */, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:dataType:"), shape, dataType)
	return rv
}


// Creates a tensor with the shape, scalar value, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:fillWithData:dataType:
func (cc _CTensorClass) TensorWithShapeFillWithDataDataType(shape []objc.IObject /* cross-framework: Number */, fillData objc.IObject /* cross-framework: NSNumber */, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:fillWithData:dataType:"), shape, fillData, dataType)
	return rv
}


// Creates a tensor with the shape and random initializer type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:randomInitializerType:
func (cc _CTensorClass) TensorWithShapeRandomInitializerType(shape []objc.IObject /* cross-framework: Number */, randomInitializerType CRandomInitializerType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:randomInitializerType:"), shape, randomInitializerType)
	return rv
}


// Creates a tensor with the shape, random initializer, and data type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:randomInitializerType:dataType:
func (cc _CTensorClass) TensorWithShapeRandomInitializerTypeDataType(shape []objc.IObject /* cross-framework: Number */, randomInitializerType CRandomInitializerType, dataType CDataType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:randomInitializerType:dataType:"), shape, randomInitializerType, dataType)
	return rv
}


// The tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/data
func (c_ CTensor) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("data"))
	return rv
}


// The configuration object you use to create a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/descriptor
func (c_ CTensor) Descriptor() IMLCTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}


// The device associated with this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/device
func (c_ CTensor) Device() IMLCDevice {
	rv := objc.Send[CDevice](c_.ID, objc.Sel("device"))
	return rv
}


// A Boolean that indicates whether a tensor contains NaN or INF values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/hasValidNumerics
func (c_ CTensor) HasValidNumerics() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasValidNumerics"))
	return rv
}


// A string that identifes this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/label
func (c_ CTensor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}


// A string that identifes this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/label
func (c_ CTensor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}


// An array that contains optimizer buffers you specify when you create a tensor parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/optimizerData
func (c_ CTensor) OptimizerData() []ICTensorData {
	rv := objc.Send[[]CTensorData](c_.ID, objc.Sel("optimizerData"))
	return rv
}


// An array that contains the device optimizer buffers you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/optimizerDeviceData
func (c_ CTensor) OptimizerDeviceData() []ICTensorOptimizerDeviceData {
	rv := objc.Send[[]CTensorOptimizerDeviceData](c_.ID, objc.Sel("optimizerDeviceData"))
	return rv
}


// A number that uniquely identifies the tensor, which the framework assigns when it creates a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorID
func (c_ CTensor) TensorID() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("tensorID"))
	return rv
}


