// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MTL4MachineLearningPipelineDescriptor] class.
var (
	MTL4MachineLearningPipelineDescriptorClass     _MTL4MachineLearningPipelineDescriptorClass
	MTL4MachineLearningPipelineDescriptorClassOnce sync.Once
)

func getMTL4MachineLearningPipelineDescriptorClass() _MTL4MachineLearningPipelineDescriptorClass {
	MTL4MachineLearningPipelineDescriptorClassOnce.Do(func() {
		MTL4MachineLearningPipelineDescriptorClass = _MTL4MachineLearningPipelineDescriptorClass{objc.GetClass("MTL4MachineLearningPipelineDescriptor")}
	})
	return MTL4MachineLearningPipelineDescriptorClass
}

type _MTL4MachineLearningPipelineDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4MachineLearningPipelineDescriptor] class.
type IMTL4MachineLearningPipelineDescriptor interface {
	IMTL4PipelineDescriptor
	

	// properties:
	Label() foundation.foundation.INSString
	SetLabel(value foundation.foundation.INSString)
	MachineLearningFunctionDescriptor() IMTL4FunctionDescriptor
	SetMachineLearningFunctionDescriptor(value IMTL4FunctionDescriptor)


	

	// methods:
	InputDimensionsAtBufferIndex(bufferIndex int) ITensorExtents
	Reset()
	SetInputDimensionsAtBufferIndex(dimensions IMTLTensorExtents, bufferIndex int)
	SetInputDimensionsWithRange(dimensions []TensorExtents, range_ foundation.Range)


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4MachineLearningPipelineDescriptorClass) Alloc() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4MachineLearningPipelineDescriptorClass) New() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4MachineLearningPipelineDescriptor) Init() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4MachineLearningPipelineDescriptor) Autorelease() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4MachineLearningPipelineDescriptor creates a new MTL4MachineLearningPipelineDescriptor instance.
func NewMTL4MachineLearningPipelineDescriptor() MTL4MachineLearningPipelineDescriptor {
	return getMTL4MachineLearningPipelineDescriptorClass().New()
}





// Description for a machine learning pipeline state.


// Description for a machine learning pipeline state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor
type MTL4MachineLearningPipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4MachineLearningPipelineDescriptorFrom constructs a [MTL4MachineLearningPipelineDescriptor] from an unsafe.Pointer.
//
// Description for a machine learning pipeline state.
func MTL4MachineLearningPipelineDescriptorFrom(ptr unsafe.Pointer) MTL4MachineLearningPipelineDescriptor {
	return MTL4MachineLearningPipelineDescriptor{
		MTL4PipelineDescriptor: MTL4PipelineDescriptorFrom(ptr),
	}
}




















// Obtains the dimensions of the input tensor at if set, otherwise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/inputDimensions(bufferIndex:)
func (m_ MTL4MachineLearningPipelineDescriptor) InputDimensionsAtBufferIndex(bufferIndex int) ITensorExtents {
	rv := objc.Send[TensorExtents](m_.ID, objc.Sel("inputDimensionsAtBufferIndex:"), bufferIndex)
	return rv
}


// Resets the descriptor to its default values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/reset()
func (m_ MTL4MachineLearningPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}


// Sets the dimension of an input tensor at a buffer index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/setInputDimensions(_:bufferIndex:)-34gir
func (m_ MTL4MachineLearningPipelineDescriptor) SetInputDimensionsAtBufferIndex(dimensions IMTLTensorExtents, bufferIndex int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInputDimensions:atBufferIndex:"), dimensions, bufferIndex)
}


// Sets the dimensions of multiple input tensors on a range of buffer bindings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/setInputDimensions:withRange:
func (m_ MTL4MachineLearningPipelineDescriptor) SetInputDimensionsWithRange(dimensions []TensorExtents, range_ foundation.Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInputDimensions:withRange:"), dimensions, range_)
}







// Assigns an optional string that helps identify pipeline states you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/label
func (m_ MTL4MachineLearningPipelineDescriptor) Label() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// Assigns an optional string that helps identify pipeline states you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/label
func (m_ MTL4MachineLearningPipelineDescriptor) SetLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// Assigns the function that the machine learning pipeline you create from this descriptor executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/machineLearningFunctionDescriptor
func (m_ MTL4MachineLearningPipelineDescriptor) MachineLearningFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("machineLearningFunctionDescriptor"))
	return rv
}


// Assigns the function that the machine learning pipeline you create from this descriptor executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor/machineLearningFunctionDescriptor
func (m_ MTL4MachineLearningPipelineDescriptor) SetMachineLearningFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMachineLearningFunctionDescriptor:"), value)
}








