// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLArgument */


/* debug [class_header]: Header for MTLArgument */
// The class instance for the [Argument] class.
var (
	ArgumentClass     _ArgumentClass
	ArgumentClassOnce sync.Once
)

func getArgumentClass() _ArgumentClass {
	ArgumentClassOnce.Do(func() {
		ArgumentClass = _ArgumentClass{objc.GetClass("MTLArgument")}
	})
	return ArgumentClass
}

type _ArgumentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Argument */
// An interface definition for the [Argument] class.
type IArgument interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Argument */
	// properties:
	Access() BindingAccess
	ArrayLength() uint
	BufferAlignment() uint
	BufferDataSize() uint
	BufferDataType() DataType
	BufferPointerType() IMTLPointerType
	BufferStructType() IMTLStructType
	Index() uint
	Active() bool
	IsDepthTexture() bool
	Name() objc.IObject /* cross-framework: NSString */
	TextureDataType() DataType
	TextureType() TextureType
	ThreadgroupMemoryAlignment() uint
	ThreadgroupMemoryDataSize() uint
	Type() ArgumentType
	IsActive() bool
	SetIsActive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Argument */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Argument */
// Alloc allocates a new instance without initialization.
func (ac _ArgumentClass) Alloc() Argument {
	rv := objc.Send[Argument](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ArgumentClass) New() Argument {
	rv := objc.Send[Argument](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Argument) Init() Argument {
	rv := objc.Send[Argument](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Argument) Autorelease() Argument {
	rv := objc.Send[Argument](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArgument creates a new Argument instance.
func NewArgument() Argument {
	return getArgumentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Argument */
// Information about an argument of a graphics or compute function.
//
// An instance describes a single argument to a Metal function. Your app uses the properties to read details about a function argument as it was defined in the Metal Shading Language. You can determine the argument’s data type, access restrictions, and its associated resource type. For buffer, texture, and threadgroup memory arguments, additional properties can be read to determine more details about the argument. Your app does not create an instance directly. Creating an or instance can generate a reflection instance ( or ) that contains instances.


// Information about an argument of a graphics or compute function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument
type Argument struct {
	objectivec.Object
}

// ArgumentFrom constructs a [Argument] from an unsafe.Pointer.
//
// Information about an argument of a graphics or compute function.
func ArgumentFrom(ptr unsafe.Pointer) Argument {
	return Argument{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Argument *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Argument */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Argument */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Argument */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Argument */

// The argument’s read and/or write access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/access
func (a_ Argument) Access() BindingAccess {
	rv := objc.Send[BindingAccess](a_.ID, objc.Sel("access"))
	return rv
}/* debug [instance_properties/getter]: access */


// The number of elements, if the argument is an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/arrayLength
func (a_ Argument) ArrayLength() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("arrayLength"))
	return rv
}/* debug [instance_properties/getter]: arrayLength */


// The required byte alignment in memory for the buffer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/bufferAlignment
func (a_ Argument) BufferAlignment() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("bufferAlignment"))
	return rv
}/* debug [instance_properties/getter]: bufferAlignment */


// The size, in bytes, of the buffer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/bufferDataSize
func (a_ Argument) BufferDataSize() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("bufferDataSize"))
	return rv
}/* debug [instance_properties/getter]: bufferDataSize */


// The data type of the buffer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/bufferDataType
func (a_ Argument) BufferDataType() DataType {
	rv := objc.Send[DataType](a_.ID, objc.Sel("bufferDataType"))
	return rv
}/* debug [instance_properties/getter]: bufferDataType */


// A description of the pointer to a buffer argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/bufferPointerType
func (a_ Argument) BufferPointerType() IMTLPointerType {
	rv := objc.Send[PointerType](a_.ID, objc.Sel("bufferPointerType"))
	return rv
}/* debug [instance_properties/getter]: bufferPointerType */


// A description of the structure data of a buffer argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/bufferStructType
func (a_ Argument) BufferStructType() IMTLStructType {
	rv := objc.Send[StructType](a_.ID, objc.Sel("bufferStructType"))
	return rv
}/* debug [instance_properties/getter]: bufferStructType */


// The index in the argument table that corresponds to the function argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/index
func (a_ Argument) Index() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("index"))
	return rv
}/* debug [instance_properties/getter]: index */


// A Boolean that indicates whether the compiled function uses the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/isActive
func (a_ Argument) Active() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// A Boolean value that indicates whether the texture is a depth texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/isDepthTexture
func (a_ Argument) IsDepthTexture() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isDepthTexture"))
	return rv
}/* debug [instance_properties/getter]: isDepthTexture */


// The name of the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/name
func (a_ Argument) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The data type of a texture argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/textureDataType
func (a_ Argument) TextureDataType() DataType {
	rv := objc.Send[DataType](a_.ID, objc.Sel("textureDataType"))
	return rv
}/* debug [instance_properties/getter]: textureDataType */


// The texture type of a texture argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/textureType
func (a_ Argument) TextureType() TextureType {
	rv := objc.Send[TextureType](a_.ID, objc.Sel("textureType"))
	return rv
}/* debug [instance_properties/getter]: textureType */


// The required byte alignment in memory for the threadgroup data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/threadgroupMemoryAlignment
func (a_ Argument) ThreadgroupMemoryAlignment() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("threadgroupMemoryAlignment"))
	return rv
}/* debug [instance_properties/getter]: threadgroupMemoryAlignment */


// The size, in bytes, of the threadgroup data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/threadgroupMemoryDataSize
func (a_ Argument) ThreadgroupMemoryDataSize() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("threadgroupMemoryDataSize"))
	return rv
}/* debug [instance_properties/getter]: threadgroupMemoryDataSize */


// The argument’s resource type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/type
func (a_ Argument) Type() ArgumentType {
	rv := objc.Send[ArgumentType](a_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// A Boolean that indicates whether the compiled function uses the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargument/isactive
func (a_ Argument) IsActive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean that indicates whether the compiled function uses the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargument/isactive
func (a_ Argument) SetIsActive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLArgument */



