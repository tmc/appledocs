// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AttributeDescriptor] class.
var (
	AttributeDescriptorClass     _AttributeDescriptorClass
	AttributeDescriptorClassOnce sync.Once
)

func getAttributeDescriptorClass() _AttributeDescriptorClass {
	AttributeDescriptorClassOnce.Do(func() {
		AttributeDescriptorClass = _AttributeDescriptorClass{objc.GetClass("MTLAttributeDescriptor")}
	})
	return AttributeDescriptorClass
}

type _AttributeDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [AttributeDescriptor] class.
type IAttributeDescriptor interface {
	objectivec.IObject
	

	// properties:
	BufferIndex() uint
	SetBufferIndex(value uint)
	Format() AttributeFormat
	SetFormat(value AttributeFormat)
	Offset() uint
	SetOffset(value uint)
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AttributeDescriptorClass) Alloc() AttributeDescriptor {
	rv := objc.Send[AttributeDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AttributeDescriptorClass) New() AttributeDescriptor {
	rv := objc.Send[AttributeDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AttributeDescriptor) Init() AttributeDescriptor {
	rv := objc.Send[AttributeDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AttributeDescriptor) Autorelease() AttributeDescriptor {
	rv := objc.Send[AttributeDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttributeDescriptor creates a new AttributeDescriptor instance.
func NewAttributeDescriptor() AttributeDescriptor {
	return getAttributeDescriptorClass().New()
}





// A descriptor of an argument’s format and where its data is in memory.
//
// Attribute descriptors are part of an or instance to provide layout information about a function’s arguments. Each descriptor is for a single argument, containing information about the attached data, offset and stride, and data type.


// A descriptor of an argument’s format and where its data is in memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor
type AttributeDescriptor struct {
	objectivec.Object
}

// AttributeDescriptorFrom constructs a [AttributeDescriptor] from an unsafe.Pointer.
//
// A descriptor of an argument’s format and where its data is in memory.
func AttributeDescriptorFrom(ptr unsafe.Pointer) AttributeDescriptor {
	return AttributeDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// The index in the buffer argument table for the buffer that contains the data for this attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/bufferIndex
func (a_ AttributeDescriptor) BufferIndex() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("bufferIndex"))
	return rv
}


// The index in the buffer argument table for the buffer that contains the data for this attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/bufferIndex
func (a_ AttributeDescriptor) SetBufferIndex(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBufferIndex:"), value)
}


// The format of the attribute’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/format
func (a_ AttributeDescriptor) Format() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("format"))
	return rv
}


// The format of the attribute’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/format
func (a_ AttributeDescriptor) SetFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFormat:"), value)
}


// The offset, in bytes, from the start of the buffer containing the attribute data to the start of the data itself.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/offset
func (a_ AttributeDescriptor) Offset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("offset"))
	return rv
}


// The offset, in bytes, from the start of the buffer containing the attribute data to the start of the data itself.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/offset
func (a_ AttributeDescriptor) SetOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOffset:"), value)
}


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (a_ AttributeDescriptor) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](a_.ID, objc.Sel("stageInputDescriptor"))
	return rv
}


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (a_ AttributeDescriptor) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStageInputDescriptor:"), value)
}








