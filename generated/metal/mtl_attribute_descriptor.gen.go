// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLAttributeDescriptor */


/* debug [class_header]: Header for MTLAttributeDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AttributeDescriptor */
// An interface definition for the [AttributeDescriptor] class.
type IAttributeDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AttributeDescriptor */
	// properties:
	BufferIndex() uint
	SetBufferIndex(value uint)
	Format() AttributeFormat
	SetFormat(value AttributeFormat)
	Offset() uint
	SetOffset(value uint)
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AttributeDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AttributeDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AttributeDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AttributeDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AttributeDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AttributeDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AttributeDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AttributeDescriptor */

// The index in the buffer argument table for the buffer that contains the data for this attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/bufferIndex
func (a_ AttributeDescriptor) BufferIndex() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("bufferIndex"))
	return rv
}/* debug [instance_properties/getter]: bufferIndex */


// The index in the buffer argument table for the buffer that contains the data for this attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/bufferIndex
func (a_ AttributeDescriptor) SetBufferIndex(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBufferIndex:"), value)
}/* debug [instance_properties/setter]: bufferIndex */


// The format of the attribute’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/format
func (a_ AttributeDescriptor) Format() AttributeFormat {
	rv := objc.Send[AttributeFormat](a_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// The format of the attribute’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/format
func (a_ AttributeDescriptor) SetFormat(value AttributeFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFormat:"), value)
}/* debug [instance_properties/setter]: format */


// The offset, in bytes, from the start of the buffer containing the attribute data to the start of the data itself.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/offset
func (a_ AttributeDescriptor) Offset() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// The offset, in bytes, from the start of the buffer containing the attribute data to the start of the data itself.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptor/offset
func (a_ AttributeDescriptor) SetOffset(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (a_ AttributeDescriptor) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](a_.ID, objc.Sel("stageInputDescriptor"))
	return rv
}/* debug [instance_properties/getter]: stageInputDescriptor */


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (a_ AttributeDescriptor) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStageInputDescriptor:"), value)
}/* debug [instance_properties/setter]: stageInputDescriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLAttributeDescriptor */



