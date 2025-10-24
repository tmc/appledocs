// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLArgumentDescriptor */


/* debug [class_header]: Header for MTLArgumentDescriptor */
// The class instance for the [ArgumentDescriptor] class.
var (
	ArgumentDescriptorClass     _ArgumentDescriptorClass
	ArgumentDescriptorClassOnce sync.Once
)

func getArgumentDescriptorClass() _ArgumentDescriptorClass {
	ArgumentDescriptorClassOnce.Do(func() {
		ArgumentDescriptorClass = _ArgumentDescriptorClass{objc.GetClass("MTLArgumentDescriptor")}
	})
	return ArgumentDescriptorClass
}

type _ArgumentDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ArgumentDescriptor */
// An interface definition for the [ArgumentDescriptor] class.
type IArgumentDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ArgumentDescriptor */
	// properties:
	Access() BindingAccess
	SetAccess(value BindingAccess)
	ArrayLength() uint
	SetArrayLength(value uint)
	ConstantBlockAlignment() uint
	SetConstantBlockAlignment(value uint)
	DataType() DataType
	SetDataType(value DataType)
	Index() uint
	SetIndex(value uint)
	TextureType() TextureType
	SetTextureType(value TextureType)
	MTLAttributeStrideStatic() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ArgumentDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ArgumentDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _ArgumentDescriptorClass) Alloc() ArgumentDescriptor {
	rv := objc.Send[ArgumentDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ArgumentDescriptorClass) New() ArgumentDescriptor {
	rv := objc.Send[ArgumentDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ArgumentDescriptor) Init() ArgumentDescriptor {
	rv := objc.Send[ArgumentDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ArgumentDescriptor) Autorelease() ArgumentDescriptor {
	rv := objc.Send[ArgumentDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArgumentDescriptor creates a new ArgumentDescriptor instance.
func NewArgumentDescriptor() ArgumentDescriptor {
	return getArgumentDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ArgumentDescriptor */
// A representation of an argument within an argument buffer.
//
// This descriptor can represent arguments within flat structures only. It can represent arrays of allowed argument buffer data types, but it cannot represent arguments within nested structures. Argument buffers with simple, flat structures can be represented by an array of instances. You can then use this array to create an instance by calling the method. Argument buffers with complex, nested structures must define their structure in Metal shading language code, which can then be directly assigned to a specific buffer index of a function. You can then use this buffer index to call the method and create an instance.


// A representation of an argument within an argument buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor
type ArgumentDescriptor struct {
	objectivec.Object
}

// ArgumentDescriptorFrom constructs a [ArgumentDescriptor] from an unsafe.Pointer.
//
// A representation of an argument within an argument buffer.
func ArgumentDescriptorFrom(ptr unsafe.Pointer) ArgumentDescriptor {
	return ArgumentDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ArgumentDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ArgumentDescriptor */

// Creates an empty argument descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/argumentDescriptor
func (ac _ArgumentDescriptorClass) ArgumentDescriptor() IArgumentDescriptor {
	rv := objc.Send[ArgumentDescriptor](objc.ID(ac.class), objc.Sel("argumentDescriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ArgumentDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ArgumentDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ArgumentDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ArgumentDescriptor */

// The access permissions of the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/access
func (a_ ArgumentDescriptor) Access() BindingAccess {
	rv := objc.Send[BindingAccess](a_.ID, objc.Sel("access"))
	return rv
}/* debug [instance_properties/getter]: access */


// The access permissions of the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/access
func (a_ ArgumentDescriptor) SetAccess(value BindingAccess) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccess:"), value)
}/* debug [instance_properties/setter]: access */


// The length of an array argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/arrayLength
func (a_ ArgumentDescriptor) ArrayLength() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("arrayLength"))
	return rv
}/* debug [instance_properties/getter]: arrayLength */


// The length of an array argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/arrayLength
func (a_ ArgumentDescriptor) SetArrayLength(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setArrayLength:"), value)
}/* debug [instance_properties/setter]: arrayLength */


// The alignment of the constant block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/constantBlockAlignment
func (a_ ArgumentDescriptor) ConstantBlockAlignment() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("constantBlockAlignment"))
	return rv
}/* debug [instance_properties/getter]: constantBlockAlignment */


// The alignment of the constant block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/constantBlockAlignment
func (a_ ArgumentDescriptor) SetConstantBlockAlignment(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConstantBlockAlignment:"), value)
}/* debug [instance_properties/setter]: constantBlockAlignment */


// The data type of the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/dataType
func (a_ ArgumentDescriptor) DataType() DataType {
	rv := objc.Send[DataType](a_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The data type of the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/dataType
func (a_ ArgumentDescriptor) SetDataType(value DataType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// The index ID of the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/index
func (a_ ArgumentDescriptor) Index() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("index"))
	return rv
}/* debug [instance_properties/getter]: index */


// The index ID of the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/index
func (a_ ArgumentDescriptor) SetIndex(value uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIndex:"), value)
}/* debug [instance_properties/setter]: index */


// The texture type of a texture argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/textureType
func (a_ ArgumentDescriptor) TextureType() TextureType {
	rv := objc.Send[TextureType](a_.ID, objc.Sel("textureType"))
	return rv
}/* debug [instance_properties/getter]: textureType */


// The texture type of a texture argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgumentDescriptor/textureType
func (a_ ArgumentDescriptor) SetTextureType(value TextureType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTextureType:"), value)
}/* debug [instance_properties/setter]: textureType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributestridestatic
func (a_ ArgumentDescriptor) MTLAttributeStrideStatic() int {
	rv := objc.Send[int](a_.ID, objc.Sel("MTLAttributeStrideStatic"))
	return rv
}/* debug [instance_properties/getter]: MTLAttributeStrideStatic */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLArgumentDescriptor */



