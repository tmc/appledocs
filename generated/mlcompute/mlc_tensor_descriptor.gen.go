// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CTensorDescriptor] class.
type ICTensorDescriptor interface {
	objectivec.IObject
}

// A configuration object you use to create a tensor.
//
// This class contains the mathematical properties of a tensor, such as data type and shape. It also includes initializers that help you create a tensor descriptor for common use cases, such as convolutional neural networks and recurrent neural networks.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CTensorDescriptorClass) Alloc() CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The batch size for each sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/batchsizepersequencestep-6iz59
func (c_ CTensorDescriptor) BatchSizePerSequenceStep() int {
	rv := objc.Send[int](c_.ID, objc.Sel("batchSizePerSequenceStep"))
	return rv
}


// SetBatchSizePerSequenceStep sets the value of the batchSizePerSequenceStep property.
// The batch size for each sequence.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/batchsizepersequencestep-6iz59
func (c_ CTensorDescriptor) SetBatchSizePerSequenceStep(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBatchSizePerSequenceStep:"), value)
}

// The tensor data type.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/datatype
func (c_ CTensorDescriptor) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dataType"))
	return rv
}


// SetDataType sets the value of the dataType property.
// The tensor data type.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/datatype
func (c_ CTensorDescriptor) SetDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataType:"), value)
}

// The number of dimensions in the tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/dimensioncount
func (c_ CTensorDescriptor) DimensionCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("dimensionCount"))
	return rv
}


// SetDimensionCount sets the value of the dimensionCount property.
// The number of dimensions in the tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/dimensioncount
func (c_ CTensorDescriptor) SetDimensionCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDimensionCount:"), value)
}

// An array that contains the variable lengths of sequences stored in the tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/sequencelengths-3jdab
func (c_ CTensorDescriptor) SequenceLengths() int {
	rv := objc.Send[int](c_.ID, objc.Sel("sequenceLengths"))
	return rv
}


// SetSequenceLengths sets the value of the sequenceLengths property.
// An array that contains the variable lengths of sequences stored in the tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/sequencelengths-3jdab
func (c_ CTensorDescriptor) SetSequenceLengths(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSequenceLengths:"), value)
}

// An array that contains the size in each dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/shape-7i1rw
func (c_ CTensorDescriptor) Shape() int {
	rv := objc.Send[int](c_.ID, objc.Sel("shape"))
	return rv
}


// SetShape sets the value of the shape property.
// An array that contains the size in each dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/shape-7i1rw
func (c_ CTensorDescriptor) SetShape(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShape:"), value)
}

// A Boolean that indicates whether you provided the sequence lengths sorted in descending order.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/sortedsequences
func (c_ CTensorDescriptor) SortedSequences() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sortedSequences"))
	return rv
}


// SetSortedSequences sets the value of the sortedSequences property.
// A Boolean that indicates whether you provided the sequence lengths sorted in descending order.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/sortedsequences
func (c_ CTensorDescriptor) SetSortedSequences(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSortedSequences:"), value)
}

// An array that contains the stride, in bytes, in each dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/stride-5mzlt
func (c_ CTensorDescriptor) Stride() int {
	rv := objc.Send[int](c_.ID, objc.Sel("stride"))
	return rv
}


// SetStride sets the value of the stride property.
// An array that contains the stride, in bytes, in each dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/stride-5mzlt
func (c_ CTensorDescriptor) SetStride(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStride:"), value)
}

// The allocation size, in bytes, for a tensor.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/tensorallocationsizeinbytes
func (c_ CTensorDescriptor) TensorAllocationSizeInBytes() int {
	rv := objc.Send[int](c_.ID, objc.Sel("tensorAllocationSizeInBytes"))
	return rv
}


// SetTensorAllocationSizeInBytes sets the value of the tensorAllocationSizeInBytes property.
// The allocation size, in bytes, for a tensor.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordescriptor/tensorallocationsizeinbytes
func (c_ CTensorDescriptor) SetTensorAllocationSizeInBytes(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTensorAllocationSizeInBytes:"), value)
}



