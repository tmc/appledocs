// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [Argument] class.
type IArgument interface {
	objectivec.IObject
	

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
	Name() foundation.foundation.INSString
	TextureDataType() DataType
	TextureType() TextureType
	ThreadgroupMemoryAlignment() uint
	ThreadgroupMemoryDataSize() uint
	Type() ArgumentType
	IsActive() bool
	SetIsActive(value bool)


	

	// methods:


}





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

























// The argument’s read and/or write access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/access
func (a_ Argument) Access() BindingAccess {
	rv := objc.Send[BindingAccess](a_.ID, objc.Sel("access"))
	return rv
}


// The number of elements, if the argument is an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/arrayLength
func (a_ Argument) ArrayLength() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("arrayLength"))
	return rv
}


// The required byte alignment in memory for the buffer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/bufferAlignment
func (a_ Argument) BufferAlignment() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("bufferAlignment"))
	return rv
}


// The size, in bytes, of the buffer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/bufferDataSize
func (a_ Argument) BufferDataSize() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("bufferDataSize"))
	return rv
}


// The data type of the buffer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/bufferDataType
func (a_ Argument) BufferDataType() DataType {
	rv := objc.Send[DataType](a_.ID, objc.Sel("bufferDataType"))
	return rv
}


// A description of the pointer to a buffer argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/bufferPointerType
func (a_ Argument) BufferPointerType() IMTLPointerType {
	rv := objc.Send[PointerType](a_.ID, objc.Sel("bufferPointerType"))
	return rv
}


// A description of the structure data of a buffer argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/bufferStructType
func (a_ Argument) BufferStructType() IMTLStructType {
	rv := objc.Send[StructType](a_.ID, objc.Sel("bufferStructType"))
	return rv
}


// The index in the argument table that corresponds to the function argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/index
func (a_ Argument) Index() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("index"))
	return rv
}


// A Boolean that indicates whether the compiled function uses the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/isActive
func (a_ Argument) Active() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("active"))
	return rv
}


// A Boolean value that indicates whether the texture is a depth texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/isDepthTexture
func (a_ Argument) IsDepthTexture() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isDepthTexture"))
	return rv
}


// The name of the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/name
func (a_ Argument) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}


// The data type of a texture argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/textureDataType
func (a_ Argument) TextureDataType() DataType {
	rv := objc.Send[DataType](a_.ID, objc.Sel("textureDataType"))
	return rv
}


// The texture type of a texture argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/textureType
func (a_ Argument) TextureType() TextureType {
	rv := objc.Send[TextureType](a_.ID, objc.Sel("textureType"))
	return rv
}


// The required byte alignment in memory for the threadgroup data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/threadgroupMemoryAlignment
func (a_ Argument) ThreadgroupMemoryAlignment() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("threadgroupMemoryAlignment"))
	return rv
}


// The size, in bytes, of the threadgroup data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/threadgroupMemoryDataSize
func (a_ Argument) ThreadgroupMemoryDataSize() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("threadgroupMemoryDataSize"))
	return rv
}


// The argument’s resource type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLArgument/type
func (a_ Argument) Type() ArgumentType {
	rv := objc.Send[ArgumentType](a_.ID, objc.Sel("type"))
	return rv
}


// A Boolean that indicates whether the compiled function uses the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargument/isactive
func (a_ Argument) IsActive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isActive"))
	return rv
}


// A Boolean that indicates whether the compiled function uses the argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargument/isactive
func (a_ Argument) SetIsActive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsActive:"), value)
}








