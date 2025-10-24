// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLStructMember */


/* debug [class_header]: Header for MTLStructMember */
// The class instance for the [StructMember] class.
var (
	StructMemberClass     _StructMemberClass
	StructMemberClassOnce sync.Once
)

func getStructMemberClass() _StructMemberClass {
	StructMemberClassOnce.Do(func() {
		StructMemberClass = _StructMemberClass{objc.GetClass("MTLStructMember")}
	})
	return StructMemberClass
}

type _StructMemberClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StructMember */
// An interface definition for the [StructMember] class.
type IStructMember interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StructMember */
	// properties:
	ArgumentIndex() uint
	DataType() DataType
	Name() objc.IObject /* cross-framework: NSString */
	Offset() uint
	Members() IMTLStructMember
	SetMembers(value IMTLStructMember)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StructMember */
	// methods:
	ArrayType() IArrayType
	PointerType() IPointerType
	StructType() IStructType
	TensorReferenceType() ITensorReferenceType
	TextureReferenceType() ITextureReferenceType
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StructMember */
// Alloc allocates a new instance without initialization.
func (sc _StructMemberClass) Alloc() StructMember {
	rv := objc.Send[StructMember](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StructMemberClass) New() StructMember {
	rv := objc.Send[StructMember](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StructMember) Init() StructMember {
	rv := objc.Send[StructMember](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StructMember) Autorelease() StructMember {
	rv := objc.Send[StructMember](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStructMember creates a new StructMember instance.
func NewStructMember() StructMember {
	return getStructMemberClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StructMember */
// An instance that provides information about a field in a structure.
//
// is part of the reflection API that allows Metal framework code to query details about an argument of a Metal shading language function. An instance describes the data type of one field in a struct that is passed as an argument, which is represented by . Don’t create instances directly. You obtain an instance from either the property or the method of an instance. The property of the instance tells you what kind of data is stored in the member. Recursively drill down every struct member until you reach a data type that is neither a struct nor an array.


// An instance that provides information about a field in a structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructMember
type StructMember struct {
	objectivec.Object
}

// StructMemberFrom constructs a [StructMember] from an unsafe.Pointer.
//
// An instance that provides information about a field in a structure.
func StructMemberFrom(ptr unsafe.Pointer) StructMember {
	return StructMember{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StructMember *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StructMember */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StructMember */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StructMember */

// Provides a description of the underlying array when the struct member holds an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructMember/arrayType()
func (s_ StructMember) ArrayType() IArrayType {
	rv := objc.Send[ArrayType](s_.ID, objc.Sel("arrayType"))
	return rv
}/* debug [instance_methods/method]: ArrayType */


// Provides a description of the underlying pointer when the struct member holds a pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructMember/pointerType()
func (s_ StructMember) PointerType() IPointerType {
	rv := objc.Send[PointerType](s_.ID, objc.Sel("pointerType"))
	return rv
}/* debug [instance_methods/method]: PointerType */


// Provides a description of the underlying struct when the struct member holds a struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructMember/structType()
func (s_ StructMember) StructType() IStructType {
	rv := objc.Send[StructType](s_.ID, objc.Sel("structType"))
	return rv
}/* debug [instance_methods/method]: StructType */


// Provides a description of the underlying tensor type when this struct member holds a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructMember/tensorReferenceType()
func (s_ StructMember) TensorReferenceType() ITensorReferenceType {
	rv := objc.Send[TensorReferenceType](s_.ID, objc.Sel("tensorReferenceType"))
	return rv
}/* debug [instance_methods/method]: TensorReferenceType */


// Provides a description of the underlying texture when the struct member holds a texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructMember/textureReferenceType()
func (s_ StructMember) TextureReferenceType() ITextureReferenceType {
	rv := objc.Send[TextureReferenceType](s_.ID, objc.Sel("textureReferenceType"))
	return rv
}/* debug [instance_methods/method]: TextureReferenceType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StructMember */

// The index in the argument table that corresponds to the struct member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructMember/argumentIndex
func (s_ StructMember) ArgumentIndex() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("argumentIndex"))
	return rv
}/* debug [instance_properties/getter]: argumentIndex */


// The data type of the struct member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructMember/dataType
func (s_ StructMember) DataType() DataType {
	rv := objc.Send[DataType](s_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The name of the struct member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructMember/name
func (s_ StructMember) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The location of this member relative to the start of its struct, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructMember/offset
func (s_ StructMember) Offset() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// An array of instances that describe the fields in the struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructtype/members
func (s_ StructMember) Members() IMTLStructMember {
	rv := objc.Send[StructMember](s_.ID, objc.Sel("members"))
	return rv
}/* debug [instance_properties/getter]: members */


// An array of instances that describe the fields in the struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructtype/members
func (s_ StructMember) SetMembers(value IMTLStructMember) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMembers:"), value)
}/* debug [instance_properties/setter]: members */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLStructMember */



