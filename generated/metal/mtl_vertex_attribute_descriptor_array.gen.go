// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLVertexAttributeDescriptorArray */


/* debug [class_header]: Header for MTLVertexAttributeDescriptorArray */
// The class instance for the [VertexAttributeDescriptorArray] class.
var (
	VertexAttributeDescriptorArrayClass     _VertexAttributeDescriptorArrayClass
	VertexAttributeDescriptorArrayClassOnce sync.Once
)

func getVertexAttributeDescriptorArrayClass() _VertexAttributeDescriptorArrayClass {
	VertexAttributeDescriptorArrayClassOnce.Do(func() {
		VertexAttributeDescriptorArrayClass = _VertexAttributeDescriptorArrayClass{objc.GetClass("MTLVertexAttributeDescriptorArray")}
	})
	return VertexAttributeDescriptorArrayClass
}

type _VertexAttributeDescriptorArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VertexAttributeDescriptorArray */
// An interface definition for the [VertexAttributeDescriptorArray] class.
type IVertexAttributeDescriptorArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VertexAttributeDescriptorArray */
	// properties:
	MTLBufferLayoutStrideDynamic() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VertexAttributeDescriptorArray */
	// methods:
	SetObjectAtIndexedSubscript(attributeDesc IMTLVertexAttributeDescriptor, index uint)
	ObjectAtIndexedSubscript(index uint) IVertexAttributeDescriptor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VertexAttributeDescriptorArray */
// Alloc allocates a new instance without initialization.
func (vc _VertexAttributeDescriptorArrayClass) Alloc() VertexAttributeDescriptorArray {
	rv := objc.Send[VertexAttributeDescriptorArray](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VertexAttributeDescriptorArrayClass) New() VertexAttributeDescriptorArray {
	rv := objc.Send[VertexAttributeDescriptorArray](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VertexAttributeDescriptorArray) Init() VertexAttributeDescriptorArray {
	rv := objc.Send[VertexAttributeDescriptorArray](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VertexAttributeDescriptorArray) Autorelease() VertexAttributeDescriptorArray {
	rv := objc.Send[VertexAttributeDescriptorArray](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVertexAttributeDescriptorArray creates a new VertexAttributeDescriptorArray instance.
func NewVertexAttributeDescriptorArray() VertexAttributeDescriptorArray {
	return getVertexAttributeDescriptorArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VertexAttributeDescriptorArray */
// An array of vertex attribute descriptor instances.
//
// An instance is an array of instances that defines how vertex attribute data is formatted and assigned to an index in the attribute argument table. The methods of set or retrieve the attribute formatting information from the array.


// An array of vertex attribute descriptor instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptorArray
type VertexAttributeDescriptorArray struct {
	objectivec.Object
}

// VertexAttributeDescriptorArrayFrom constructs a [VertexAttributeDescriptorArray] from an unsafe.Pointer.
//
// An array of vertex attribute descriptor instances.
func VertexAttributeDescriptorArrayFrom(ptr unsafe.Pointer) VertexAttributeDescriptorArray {
	return VertexAttributeDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VertexAttributeDescriptorArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VertexAttributeDescriptorArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VertexAttributeDescriptorArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VertexAttributeDescriptorArray */

// Sets state for the specified vertex attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptorArray/setObject:atIndexedSubscript:
func (v_ VertexAttributeDescriptorArray) SetObjectAtIndexedSubscript(attributeDesc IMTLVertexAttributeDescriptor, index uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setObject:atIndexedSubscript:"), attributeDesc, index)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Returns the state of the specified vertex attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAttributeDescriptorArray/subscript(_:)
func (v_ VertexAttributeDescriptorArray) ObjectAtIndexedSubscript(index uint) IVertexAttributeDescriptor {
	rv := objc.Send[VertexAttributeDescriptor](v_.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VertexAttributeDescriptorArray */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v_ VertexAttributeDescriptorArray) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v_.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}/* debug [instance_properties/getter]: MTLBufferLayoutStrideDynamic */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLVertexAttributeDescriptorArray */



