// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [StageInputOutputDescriptor] class.
var (
	StageInputOutputDescriptorClass     _StageInputOutputDescriptorClass
	StageInputOutputDescriptorClassOnce sync.Once
)

func getStageInputOutputDescriptorClass() _StageInputOutputDescriptorClass {
	StageInputOutputDescriptorClassOnce.Do(func() {
		StageInputOutputDescriptorClass = _StageInputOutputDescriptorClass{objc.GetClass("MTLStageInputOutputDescriptor")}
	})
	return StageInputOutputDescriptorClass
}

type _StageInputOutputDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [StageInputOutputDescriptor] class.
type IStageInputOutputDescriptor interface {
	objectivec.IObject
	

	// properties:
	Attributes() IMTLAttributeDescriptorArray
	IndexBufferIndex() uint
	SetIndexBufferIndex(value uint)
	IndexType() IndexType
	SetIndexType(value IndexType)
	Layouts() IMTLBufferLayoutDescriptorArray


	

	// methods:
	Reset()


}





// Alloc allocates a new instance without initialization.
func (sc _StageInputOutputDescriptorClass) Alloc() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StageInputOutputDescriptorClass) New() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StageInputOutputDescriptor) Init() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StageInputOutputDescriptor) Autorelease() StageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStageInputOutputDescriptor creates a new StageInputOutputDescriptor instance.
func NewStageInputOutputDescriptor() StageInputOutputDescriptor {
	return getStageInputOutputDescriptorClass().New()
}





// A description of the input and output data of a function.


// A description of the input and output data of a function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor
type StageInputOutputDescriptor struct {
	objectivec.Object
}

// StageInputOutputDescriptorFrom constructs a [StageInputOutputDescriptor] from an unsafe.Pointer.
//
// A description of the input and output data of a function.
func StageInputOutputDescriptorFrom(ptr unsafe.Pointer) StageInputOutputDescriptor {
	return StageInputOutputDescriptor{objectivec.Object{objc.ID(ptr)}}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/stageInputOutputDescriptor
func (sc _StageInputOutputDescriptorClass) StageInputOutputDescriptor() IStageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](objc.ID(sc.class), objc.Sel("stageInputOutputDescriptor"))
	return rv
}












// Resets the default state for the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/reset()
func (s_ StageInputOutputDescriptor) Reset() {
	objc.Send[objc.ID](s_.ID, objc.Sel("reset"))
}







// An array that describes where and how to fetch data for the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/attributes
func (s_ StageInputOutputDescriptor) Attributes() IMTLAttributeDescriptorArray {
	rv := objc.Send[AttributeDescriptorArray](s_.ID, objc.Sel("attributes"))
	return rv
}


// The location of the index buffer for a compute function using indexed thread addressing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/indexBufferIndex
func (s_ StageInputOutputDescriptor) IndexBufferIndex() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("indexBufferIndex"))
	return rv
}


// The location of the index buffer for a compute function using indexed thread addressing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/indexBufferIndex
func (s_ StageInputOutputDescriptor) SetIndexBufferIndex(value uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIndexBufferIndex:"), value)
}


// The data type of the indices stored in the index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/indexType
func (s_ StageInputOutputDescriptor) IndexType() IndexType {
	rv := objc.Send[IndexType](s_.ID, objc.Sel("indexType"))
	return rv
}


// The data type of the indices stored in the index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/indexType
func (s_ StageInputOutputDescriptor) SetIndexType(value IndexType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIndexType:"), value)
}


// An array that describes how the function fetches data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInputOutputDescriptor/layouts
func (s_ StageInputOutputDescriptor) Layouts() IMTLBufferLayoutDescriptorArray {
	rv := objc.Send[BufferLayoutDescriptorArray](s_.ID, objc.Sel("layouts"))
	return rv
}








