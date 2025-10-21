// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	BindAndWriteDataToDevice(data unsafe.Pointer, device unsafe.Pointer) bool
	BindOptimizerDataDeviceData(data unsafe.Pointer, deviceData unsafe.Pointer) bool
	CopyDataFromDeviceMemoryToBytesLengthSynchronizeWithDevice(bytes unsafe.Pointer, length uint, synchronizeWithDevice bool) bool
	TensorByDequantizingToTypeScaleBiasAxis(type_ unsafe.Pointer, scale unsafe.Pointer, bias unsafe.Pointer, axis int) unsafe.Pointer
	TensorByDequantizingToTypeScaleBias(type_ unsafe.Pointer, scale unsafe.Pointer, bias unsafe.Pointer) unsafe.Pointer
	TensorByQuantizingToTypeScaleBias(type_ unsafe.Pointer, scale unsafe.Pointer, bias int) unsafe.Pointer
	TensorByQuantizingToTypeScaleBiasAxis(type_ unsafe.Pointer, scale unsafe.Pointer, bias unsafe.Pointer, axis int) unsafe.Pointer
	SynchronizeData() bool
	SynchronizeOptimizerData() bool
}

// The data object you use throughout the framework.
//
// Create a tensor with or without data. For example, create a tensor with data for weights used by convolution or mean, variance, beta, and gamma parameters with batch normalization. Create a tensor without data to use as an input tensor when you build a graph.
//
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
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:)
func NewCTensorWithDescriptor(tensorDescriptor unsafe.Pointer) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:"), tensorDescriptor)
	return rv
}



// Creates a tensor with the descriptor and data you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:data:)
func NewCTensorWithDescriptorData(tensorDescriptor unsafe.Pointer, data unsafe.Pointer) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:data:"), tensorDescriptor, data)
	return rv
}



// Creates a tensor with the descriptor and scalar value you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:fillWithData:)
func NewCTensorWithDescriptorFillWithData(tensorDescriptor unsafe.Pointer, fillData unsafe.Pointer) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:fillWithData:"), tensorDescriptor, fillData)
	return rv
}



// Creates a tensor with the descriptor and random initializer type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:randomInitializerType:)
func NewCTensorWithDescriptorRandomInitializerType(tensorDescriptor unsafe.Pointer, randomInitializerType unsafe.Pointer) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithDescriptor:randomInitializerType:"), tensorDescriptor, randomInitializerType)
	return rv
}



// Creates a tensor without data, with the sequence length, number of feature channels, and batch size you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:)
func NewCTensorWithSequenceLengthFeatureChannelCountBatchSize(sequenceLength uint, featureChannelCount uint, batchSize uint) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:"), sequenceLength, featureChannelCount, batchSize)
	return rv
}



// Creates a tensor with the sequence length, number of feature channels, batch size, and data you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:data:)
func NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeData(sequenceLength uint, featureChannelCount uint, batchSize uint, data unsafe.Pointer) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:data:"), sequenceLength, featureChannelCount, batchSize, data)
	return rv
}



// Creates a tensor with the sequence length, number of feature channels, batch size, and random initializer type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:randomInitializerType:)
func NewCTensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType(sequenceLength uint, featureChannelCount uint, batchSize uint, randomInitializerType unsafe.Pointer) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:randomInitializerType:"), sequenceLength, featureChannelCount, batchSize, randomInitializerType)
	return rv
}



// Creates a tensor without data, with the sizes and number of feature channels you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSize(width uint, height uint, featureChannelCount uint, batchSize uint) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:"), width, height, featureChannelCount, batchSize)
	return rv
}



// Creates a tensor with the sizes, number of feature channels, and data you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeData(width uint, height uint, featureChannelCount uint, batchSize uint, data unsafe.Pointer) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:"), width, height, featureChannelCount, batchSize, data)
	return rv
}



// Creates a tensor with the sizes, number of feature channels, data, and data type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:dataType:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, data unsafe.Pointer, dataType unsafe.Pointer) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:dataType:"), width, height, featureChannelCount, batchSize, data, dataType)
	return rv
}



// Creates a tensor with the sizes and number of feature channels, and filled with the data and type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:fillWithData:dataType:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, fillData unsafe.Pointer, dataType unsafe.Pointer) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:fillWithData:dataType:"), width, height, featureChannelCount, batchSize, fillData, dataType)
	return rv
}



// Creates a tensor with the sizes, number of feature channels, and random data using the random initializer type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:randomInitializerType:)
func NewCTensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType(width uint, height uint, featureChannelCount uint, batchSize uint, randomInitializerType unsafe.Pointer) CTensor {
	rv := objc.Send[CTensor](objc.ID(getCTensorClass().class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:randomInitializerType:"), width, height, featureChannelCount, batchSize, randomInitializerType)
	return rv
}


// Creates a tensor without data, using the descriptor you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:)
func (cc _CTensorClass) TensorWithDescriptor(tensorDescriptor unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:"), tensorDescriptor)
	return rv
}

// Creates a tensor with the descriptor and data you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:data:)
func (cc _CTensorClass) TensorWithDescriptorData(tensorDescriptor unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:data:"), tensorDescriptor, data)
	return rv
}

// Creates a tensor with the descriptor and scalar value you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:fillWithData:)
func (cc _CTensorClass) TensorWithDescriptorFillWithData(tensorDescriptor unsafe.Pointer, fillData unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:fillWithData:"), tensorDescriptor, fillData)
	return rv
}

// Creates a tensor with the descriptor and random initializer type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(descriptor:randomInitializerType:)
func (cc _CTensorClass) TensorWithDescriptorRandomInitializerType(tensorDescriptor unsafe.Pointer, randomInitializerType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithDescriptor:randomInitializerType:"), tensorDescriptor, randomInitializerType)
	return rv
}

// Creates a tensor without data, with the sequence length, number of feature channels, and batch size you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:)
func (cc _CTensorClass) TensorWithSequenceLengthFeatureChannelCountBatchSize(sequenceLength uint, featureChannelCount uint, batchSize uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:"), sequenceLength, featureChannelCount, batchSize)
	return rv
}

// Creates a tensor with the sequence length, number of feature channels, batch size, and data you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:data:)
func (cc _CTensorClass) TensorWithSequenceLengthFeatureChannelCountBatchSizeData(sequenceLength uint, featureChannelCount uint, batchSize uint, data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:data:"), sequenceLength, featureChannelCount, batchSize, data)
	return rv
}

// Creates a tensor with the sequence length, number of feature channels, batch size, and random initializer type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(sequenceLength:featureChannelCount:batchSize:randomInitializerType:)
func (cc _CTensorClass) TensorWithSequenceLengthFeatureChannelCountBatchSizeRandomInitializerType(sequenceLength uint, featureChannelCount uint, batchSize uint, randomInitializerType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLength:featureChannelCount:batchSize:randomInitializerType:"), sequenceLength, featureChannelCount, batchSize, randomInitializerType)
	return rv
}

// Creates a tensor without data, with the sizes and number of feature channels you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSize(width uint, height uint, featureChannelCount uint, batchSize uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:"), width, height, featureChannelCount, batchSize)
	return rv
}

// Creates a tensor with the sizes, number of feature channels, and data you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeData(width uint, height uint, featureChannelCount uint, batchSize uint, data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:"), width, height, featureChannelCount, batchSize, data)
	return rv
}

// Creates a tensor with the sizes, number of feature channels, data, and data type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:data:dataType:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, data unsafe.Pointer, dataType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:data:dataType:"), width, height, featureChannelCount, batchSize, data, dataType)
	return rv
}

// Creates a tensor with the sizes and number of feature channels, and filled with the data and type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:fillWithData:dataType:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeFillWithDataDataType(width uint, height uint, featureChannelCount uint, batchSize uint, fillData unsafe.Pointer, dataType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:fillWithData:dataType:"), width, height, featureChannelCount, batchSize, fillData, dataType)
	return rv
}

// Creates a tensor with the sizes, number of feature channels, and random data using the random initializer type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/init(width:height:featureChannelCount:batchSize:randomInitializerType:)
func (cc _CTensorClass) TensorWithWidthHeightFeatureChannelCountBatchSizeRandomInitializerType(width uint, height uint, featureChannelCount uint, batchSize uint, randomInitializerType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithWidth:height:featureChannelCount:batchSize:randomInitializerType:"), width, height, featureChannelCount, batchSize, randomInitializerType)
	return rv
}

// Creates a tensor with the sequence lengths, sorting indicator, number of feature channels, batch size, and data you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:data:
func (cc _CTensorClass) TensorWithSequenceLengthsSortedSequencesFeatureChannelCountBatchSizeData(sequenceLengths unsafe.Pointer, sortedSequences bool, featureChannelCount uint, batchSize uint, data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:data:"), sequenceLengths, sortedSequences, featureChannelCount, batchSize, data)
	return rv
}

// Creates a tensor with the sequence lengths, sorting indicator, number of feature channels, batch size, and random initializer type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:randomInitializerType:
func (cc _CTensorClass) TensorWithSequenceLengthsSortedSequencesFeatureChannelCountBatchSizeRandomInitializerType(sequenceLengths unsafe.Pointer, sortedSequences bool, featureChannelCount uint, batchSize uint, randomInitializerType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithSequenceLengths:sortedSequences:featureChannelCount:batchSize:randomInitializerType:"), sequenceLengths, sortedSequences, featureChannelCount, batchSize, randomInitializerType)
	return rv
}

// Creates a tensor without data, with the shape you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:
func (cc _CTensorClass) TensorWithShape(shape unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:"), shape)
	return rv
}

// Creates a tensor with the shape, data, and data type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:data:dataType:
func (cc _CTensorClass) TensorWithShapeDataDataType(shape unsafe.Pointer, data unsafe.Pointer, dataType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:data:dataType:"), shape, data, dataType)
	return rv
}

// Creates a tensor without data, with the shape and data type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:dataType:
func (cc _CTensorClass) TensorWithShapeDataType(shape unsafe.Pointer, dataType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:dataType:"), shape, dataType)
	return rv
}

// Creates a tensor with the shape, scalar value, and data type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:fillWithData:dataType:
func (cc _CTensorClass) TensorWithShapeFillWithDataDataType(shape unsafe.Pointer, fillData unsafe.Pointer, dataType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:fillWithData:dataType:"), shape, fillData, dataType)
	return rv
}

// Creates a tensor with the shape and random initializer type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:randomInitializerType:
func (cc _CTensorClass) TensorWithShapeRandomInitializerType(shape unsafe.Pointer, randomInitializerType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:randomInitializerType:"), shape, randomInitializerType)
	return rv
}

// Creates a tensor with the shape, random initializer, and data type you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorWithShape:randomInitializerType:dataType:
func (cc _CTensorClass) TensorWithShapeRandomInitializerTypeDataType(shape unsafe.Pointer, randomInitializerType unsafe.Pointer, dataType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("tensorWithShape:randomInitializerType:dataType:"), shape, randomInitializerType, dataType)
	return rv
}

// Associates the given data to the tensor, and if the device is a GPU, also copies the data to the device memory.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/bindAndWriteData(_:to:)
func (c_ CTensor) BindAndWriteDataToDevice(data unsafe.Pointer, device unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bindAndWriteData:toDevice:"), data, device)
	return rv
}

// Associates the optimizer and device data buffers you specify to the tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/bindOptimizerData(_:deviceData:)
func (c_ CTensor) BindOptimizerDataDeviceData(data unsafe.Pointer, deviceData unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bindOptimizerData:deviceData:"), data, deviceData)
	return rv
}

// Copies tensor data from device memory to user-specified memory.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/copyDataFromDeviceMemory(toBytes:length:synchronizeWithDevice:)
func (c_ CTensor) CopyDataFromDeviceMemoryToBytesLengthSynchronizeWithDevice(bytes unsafe.Pointer, length uint, synchronizeWithDevice bool) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("copyDataFromDeviceMemoryToBytes:length:synchronizeWithDevice:"), bytes, length, synchronizeWithDevice)
	return rv
}

// Converts a tensor you quantize to a 32-bit floating-point tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/dequantized(to:scale:bias:axis:)
func (c_ CTensor) TensorByDequantizingToTypeScaleBiasAxis(type_ unsafe.Pointer, scale unsafe.Pointer, bias unsafe.Pointer, axis int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("tensorByDequantizingToType:scale:bias:axis:"), type_, scale, bias, axis)
	return rv
}

// Converts a tensor you quantize to a 32-bit floating-point tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/dequantized(to:scale:zeroPoint:)
func (c_ CTensor) TensorByDequantizingToTypeScaleBias(type_ unsafe.Pointer, scale unsafe.Pointer, bias unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("tensorByDequantizingToType:scale:bias:"), type_, scale, bias)
	return rv
}

// Converts a 32-bit floating-point tensor with the scale and bias you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/quantized(to:scale:bias:)
func (c_ CTensor) TensorByQuantizingToTypeScaleBias(type_ unsafe.Pointer, scale unsafe.Pointer, bias int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("tensorByQuantizingToType:scale:bias:"), type_, scale, bias)
	return rv
}

// Converts a 32-bit floating-point tensor with the scale and bias you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/quantized(to:scale:bias:axis:)
func (c_ CTensor) TensorByQuantizingToTypeScaleBiasAxis(type_ unsafe.Pointer, scale unsafe.Pointer, bias unsafe.Pointer, axis int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("tensorByQuantizingToType:scale:bias:axis:"), type_, scale, bias, axis)
	return rv
}

// Synchronizes the data in host memory.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/synchronizeData()
func (c_ CTensor) SynchronizeData() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("synchronizeData"))
	return rv
}

// Synchronizes the optimizer data in host memory.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/synchronizeOptimizerData()
func (c_ CTensor) SynchronizeOptimizerData() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("synchronizeOptimizerData"))
	return rv
}

// The tensor data.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/data
func (c_ CTensor) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("data"))
	return rv
}

// The configuration object you use to create a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/descriptor
func (c_ CTensor) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("descriptor"))
	return rv
}

// The device associated with this tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/device
func (c_ CTensor) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("device"))
	return rv
}

// A Boolean that indicates whether a tensor contains NaN or INF values.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/hasValidNumerics
func (c_ CTensor) HasValidNumerics() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasValidNumerics"))
	return rv
}

// A string that identifes this tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/label
func (c_ CTensor) Label() string {
	rv := objc.Send[string](c_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A string that identifes this tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/label
func (c_ CTensor) SetLabel(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// An array that contains optimizer buffers you specify when you create a tensor parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/optimizerData
func (c_ CTensor) OptimizerData() []CTensorData {
	rv := objc.Send[[]CTensorData](c_.ID, objc.Sel("optimizerData"))
	return rv
}

// An array that contains the device optimizer buffers you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/optimizerDeviceData
func (c_ CTensor) OptimizerDeviceData() []CTensorOptimizerDeviceData {
	rv := objc.Send[[]CTensorOptimizerDeviceData](c_.ID, objc.Sel("optimizerDeviceData"))
	return rv
}

// A number that uniquely identifies the tensor, which the framework assigns when it creates a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensor/tensorID
func (c_ CTensor) TensorID() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("tensorID"))
	return rv
}


