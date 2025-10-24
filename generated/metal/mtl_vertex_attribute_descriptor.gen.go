// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLVertexAttributeDescriptor */


/* debug [class_header]: Header for MTLVertexAttributeDescriptor */
// The class instance for the [VertexAttributeDescriptor] class.
var (
	VertexAttributeDescriptorClass     _VertexAttributeDescriptorClass
	VertexAttributeDescriptorClassOnce sync.Once
)

func getVertexAttributeDescriptorClass() _VertexAttributeDescriptorClass {
	VertexAttributeDescriptorClassOnce.Do(func() {
		VertexAttributeDescriptorClass = _VertexAttributeDescriptorClass{objc.GetClass("MTLVertexAttributeDescriptor")}
	})
	return VertexAttributeDescriptorClass
}

type _VertexAttributeDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VertexAttributeDescriptor */
// An interface definition for the [VertexAttributeDescriptor] class.
type IVertexAttributeDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VertexAttributeDescriptor */
	// properties:
	BufferIndex() uint
	SetBufferIndex(value uint)
	Format() VertexFormat
	SetFormat(value VertexFormat)
	Offset() uint
	SetOffset(value uint)
	MTLBufferLayoutStrideDynamic() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VertexAttributeDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VertexAttributeDescriptor */
// Alloc allocates a new instance without initialization.
func (vc _VertexAttributeDescriptorClass) Alloc() VertexAttributeDescriptor {
	rv := objc.Send[VertexAttributeDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VertexAttributeDescriptorClass) New() VertexAttributeDescriptor {
	rv := objc.Send[VertexAttributeDescriptor](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VertexAttributeDescriptor) Init() VertexAttributeDescriptor {
	rv := objc.Send[VertexAttributeDescriptor](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VertexAttributeDescriptor) Autorelease() VertexAttributeDescriptor {
	rv := objc.Send[VertexAttributeDescriptor](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVertexAttributeDescriptor creates a new VertexAttributeDescriptor instance.
func NewVertexAttributeDescriptor() VertexAttributeDescriptor {
	return getVertexAttributeDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VertexAttributeDescriptor */
// An object that determines how to store attribute data in memory and map it to the arguments of a vertex function.
//
// A vertex attribute descriptor provides organization information so a vertex shader function can locate and load data into its arguments. The descriptor maps memory locations to attribute locations. It supports access to multiple attributes (such as vertex coordinates, surface normals, and texture coordinates) that are interleaved within the same buffer.


// An object that determines how to store attribute data in memory and map it to the arguments of a vertex function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor
type VertexAttributeDescriptor struct {
	objectivec.Object
}

// VertexAttributeDescriptorFrom constructs a [VertexAttributeDescriptor] from an unsafe.Pointer.
//
// An object that determines how to store attribute data in memory and map it to the arguments of a vertex function.
func VertexAttributeDescriptorFrom(ptr unsafe.Pointer) VertexAttributeDescriptor {
	return VertexAttributeDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VertexAttributeDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VertexAttributeDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VertexAttributeDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VertexAttributeDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VertexAttributeDescriptor */

// The index in the argument table for the associated vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/bufferIndex
func (v_ VertexAttributeDescriptor) BufferIndex() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("bufferIndex"))
	return rv
}/* debug [instance_properties/getter]: bufferIndex */


// The index in the argument table for the associated vertex buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/bufferIndex
func (v_ VertexAttributeDescriptor) SetBufferIndex(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBufferIndex:"), value)
}/* debug [instance_properties/setter]: bufferIndex */


// The format of the vertex attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/format
func (v_ VertexAttributeDescriptor) Format() VertexFormat {
	rv := objc.Send[VertexFormat](v_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// The format of the vertex attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/format
func (v_ VertexAttributeDescriptor) SetFormat(value VertexFormat) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFormat:"), value)
}/* debug [instance_properties/setter]: format */


// The location of an attribute in vertex data, determined by the byte offset from the start of the vertex data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/offset
func (v_ VertexAttributeDescriptor) Offset() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// The location of an attribute in vertex data, determined by the byte offset from the start of the vertex data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptor/offset
func (v_ VertexAttributeDescriptor) SetOffset(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v_ VertexAttributeDescriptor) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v_.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}/* debug [instance_properties/getter]: MTLBufferLayoutStrideDynamic */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLVertexAttributeDescriptor */



