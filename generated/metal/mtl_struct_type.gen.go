// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTLStructType */


/* debug [class_header]: Header for MTLStructType */
// The class instance for the [StructType] class.
var (
	StructTypeClass     _StructTypeClass
	StructTypeClassOnce sync.Once
)

func getStructTypeClass() _StructTypeClass {
	StructTypeClassOnce.Do(func() {
		StructTypeClass = _StructTypeClass{objc.GetClass("MTLStructType")}
	})
	return StructTypeClass
}

type _StructTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StructType */
// An interface definition for the [StructType] class.
type IStructType interface {
	IType
	
/* debug [class_interface_properties]: Properties for StructType */
	// properties:
	Members() []StructMember
	BufferStructType() IMTLStructType
	SetBufferStructType(value IMTLStructType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StructType */
	// methods:
	MemberByName(name objc.IObject /* cross-framework: NSString */) IStructMember
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StructType */
// Alloc allocates a new instance without initialization.
func (sc _StructTypeClass) Alloc() StructType {
	rv := objc.Send[StructType](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StructTypeClass) New() StructType {
	rv := objc.Send[StructType](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StructType) Init() StructType {
	rv := objc.Send[StructType](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StructType) Autorelease() StructType {
	rv := objc.Send[StructType](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStructType creates a new StructType instance.
func NewStructType() StructType {
	return getStructTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StructType */
// A description of a structure.
//
// is part of the reflection API that allows Metal framework code to query details of a struct that is passed as an argument of a Metal shading language function. Don’t create instances directly; instead query the property of an instance, or call the method for an instance. To examine the details of the struct, you can recursively drill down the property of the instance, which contains details about struct members, each of which is represented by an instance.


// A description of a structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructType
type StructType struct {
	Type
}

// StructTypeFrom constructs a [StructType] from an unsafe.Pointer.
//
// A description of a structure.
func StructTypeFrom(ptr unsafe.Pointer) StructType {
	return StructType{
		Type: TypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StructType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StructType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StructType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StructType */

// Provides a representation of a struct member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructType/memberByName(_:)
func (s_ StructType) MemberByName(name objc.IObject /* cross-framework: NSString */) IStructMember {
	rv := objc.Send[StructMember](s_.ID, objc.Sel("memberByName:"), name)
	return rv
}/* debug [instance_methods/method]: MemberByName */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StructType */

// An array of instances that describe the fields in the struct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStructType/members
func (s_ StructType) Members() []StructMember {
	rv := objc.Send[[]StructMember](s_.ID, objc.Sel("members"))
	return rv
}/* debug [instance_properties/getter]: members */


// A description of the structure data of a buffer argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargument/bufferstructtype
func (s_ StructType) BufferStructType() IMTLStructType {
	rv := objc.Send[StructType](s_.ID, objc.Sel("bufferStructType"))
	return rv
}/* debug [instance_properties/getter]: bufferStructType */


// A description of the structure data of a buffer argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargument/bufferstructtype
func (s_ StructType) SetBufferStructType(value IMTLStructType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBufferStructType:"), value)
}/* debug [instance_properties/setter]: bufferStructType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLStructType */



