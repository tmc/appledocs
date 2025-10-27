// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [BufferLayoutDescriptorArray] class.
var (
	BufferLayoutDescriptorArrayClass     _BufferLayoutDescriptorArrayClass
	BufferLayoutDescriptorArrayClassOnce sync.Once
)

func getBufferLayoutDescriptorArrayClass() _BufferLayoutDescriptorArrayClass {
	BufferLayoutDescriptorArrayClassOnce.Do(func() {
		BufferLayoutDescriptorArrayClass = _BufferLayoutDescriptorArrayClass{objc.GetClass("MTLBufferLayoutDescriptorArray")}
	})
	return BufferLayoutDescriptorArrayClass
}

type _BufferLayoutDescriptorArrayClass struct {
	class objc.Class
}





// An interface definition for the [BufferLayoutDescriptorArray] class.
type IBufferLayoutDescriptorArray interface {
	objectivec.IObject
	

	// properties:
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)


	

	// methods:
	SetObjectAtIndexedSubscript(bufferDesc IMTLBufferLayoutDescriptor, index uint)
	ObjectAtIndexedSubscript(index uint) IBufferLayoutDescriptor


}





// Alloc allocates a new instance without initialization.
func (bc _BufferLayoutDescriptorArrayClass) Alloc() BufferLayoutDescriptorArray {
	rv := objc.Send[BufferLayoutDescriptorArray](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BufferLayoutDescriptorArrayClass) New() BufferLayoutDescriptorArray {
	rv := objc.Send[BufferLayoutDescriptorArray](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BufferLayoutDescriptorArray) Init() BufferLayoutDescriptorArray {
	rv := objc.Send[BufferLayoutDescriptorArray](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BufferLayoutDescriptorArray) Autorelease() BufferLayoutDescriptorArray {
	rv := objc.Send[BufferLayoutDescriptorArray](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBufferLayoutDescriptorArray creates a new BufferLayoutDescriptorArray instance.
func NewBufferLayoutDescriptorArray() BufferLayoutDescriptorArray {
	return getBufferLayoutDescriptorArrayClass().New()
}





// An array of buffer layout descriptor objects.
//
// An defines the data layout and loading for compute data, using instances.


// An array of buffer layout descriptor objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptorArray
type BufferLayoutDescriptorArray struct {
	objectivec.Object
}

// BufferLayoutDescriptorArrayFrom constructs a [BufferLayoutDescriptorArray] from an unsafe.Pointer.
//
// An array of buffer layout descriptor objects.
func BufferLayoutDescriptorArrayFrom(ptr unsafe.Pointer) BufferLayoutDescriptorArray {
	return BufferLayoutDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}




















// Sets the state of the specified buffer layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptorArray/setObject:atIndexedSubscript:
func (b_ BufferLayoutDescriptorArray) SetObjectAtIndexedSubscript(bufferDesc IMTLBufferLayoutDescriptor, index uint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setObject:atIndexedSubscript:"), bufferDesc, index)
}


// Returns the state of the specified buffer layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptorArray/subscript(_:)
func (b_ BufferLayoutDescriptorArray) ObjectAtIndexedSubscript(index uint) IBufferLayoutDescriptor {
	rv := objc.Send[BufferLayoutDescriptor](b_.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}







// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (b_ BufferLayoutDescriptorArray) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](b_.ID, objc.Sel("stageInputDescriptor"))
	return rv
}


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (b_ BufferLayoutDescriptorArray) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setStageInputDescriptor:"), value)
}








