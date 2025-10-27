// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [BufferLayoutDescriptor] class.
var (
	BufferLayoutDescriptorClass     _BufferLayoutDescriptorClass
	BufferLayoutDescriptorClassOnce sync.Once
)

func getBufferLayoutDescriptorClass() _BufferLayoutDescriptorClass {
	BufferLayoutDescriptorClassOnce.Do(func() {
		BufferLayoutDescriptorClass = _BufferLayoutDescriptorClass{objc.GetClass("MTLBufferLayoutDescriptor")}
	})
	return BufferLayoutDescriptorClass
}

type _BufferLayoutDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [BufferLayoutDescriptor] class.
type IBufferLayoutDescriptor interface {
	objectivec.IObject
	

	// properties:
	StepFunction() StepFunction
	SetStepFunction(value StepFunction)
	StepRate() uint
	SetStepRate(value uint)
	Stride() uint
	SetStride(value uint)
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (bc _BufferLayoutDescriptorClass) Alloc() BufferLayoutDescriptor {
	rv := objc.Send[BufferLayoutDescriptor](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BufferLayoutDescriptorClass) New() BufferLayoutDescriptor {
	rv := objc.Send[BufferLayoutDescriptor](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BufferLayoutDescriptor) Init() BufferLayoutDescriptor {
	rv := objc.Send[BufferLayoutDescriptor](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BufferLayoutDescriptor) Autorelease() BufferLayoutDescriptor {
	rv := objc.Send[BufferLayoutDescriptor](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBufferLayoutDescriptor creates a new BufferLayoutDescriptor instance.
func NewBufferLayoutDescriptor() BufferLayoutDescriptor {
	return getBufferLayoutDescriptorClass().New()
}





// A description of how a compute function fetches input data for an attribute.


// A description of how a compute function fetches input data for an attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor
type BufferLayoutDescriptor struct {
	objectivec.Object
}

// BufferLayoutDescriptorFrom constructs a [BufferLayoutDescriptor] from an unsafe.Pointer.
//
// A description of how a compute function fetches input data for an attribute.
func BufferLayoutDescriptorFrom(ptr unsafe.Pointer) BufferLayoutDescriptor {
	return BufferLayoutDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// Determines how and when compute functions fetch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stepFunction
func (b_ BufferLayoutDescriptor) StepFunction() StepFunction {
	rv := objc.Send[StepFunction](b_.ID, objc.Sel("stepFunction"))
	return rv
}


// Determines how and when compute functions fetch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stepFunction
func (b_ BufferLayoutDescriptor) SetStepFunction(value StepFunction) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setStepFunction:"), value)
}


// How frequently the step function should load data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stepRate
func (b_ BufferLayoutDescriptor) StepRate() uint {
	rv := objc.Send[uint](b_.ID, objc.Sel("stepRate"))
	return rv
}


// How frequently the step function should load data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stepRate
func (b_ BufferLayoutDescriptor) SetStepRate(value uint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setStepRate:"), value)
}


// The number of bytes from one buffer entry to the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stride
func (b_ BufferLayoutDescriptor) Stride() uint {
	rv := objc.Send[uint](b_.ID, objc.Sel("stride"))
	return rv
}


// The number of bytes from one buffer entry to the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stride
func (b_ BufferLayoutDescriptor) SetStride(value uint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setStride:"), value)
}


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (b_ BufferLayoutDescriptor) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](b_.ID, objc.Sel("stageInputDescriptor"))
	return rv
}


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (b_ BufferLayoutDescriptor) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setStageInputDescriptor:"), value)
}








