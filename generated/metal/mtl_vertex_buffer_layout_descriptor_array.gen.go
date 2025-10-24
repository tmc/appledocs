// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLVertexBufferLayoutDescriptorArray */


/* debug [class_header]: Header for MTLVertexBufferLayoutDescriptorArray */
// The class instance for the [VertexBufferLayoutDescriptorArray] class.
var (
	VertexBufferLayoutDescriptorArrayClass     _VertexBufferLayoutDescriptorArrayClass
	VertexBufferLayoutDescriptorArrayClassOnce sync.Once
)

func getVertexBufferLayoutDescriptorArrayClass() _VertexBufferLayoutDescriptorArrayClass {
	VertexBufferLayoutDescriptorArrayClassOnce.Do(func() {
		VertexBufferLayoutDescriptorArrayClass = _VertexBufferLayoutDescriptorArrayClass{objc.GetClass("MTLVertexBufferLayoutDescriptorArray")}
	})
	return VertexBufferLayoutDescriptorArrayClass
}

type _VertexBufferLayoutDescriptorArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VertexBufferLayoutDescriptorArray */
// An interface definition for the [VertexBufferLayoutDescriptorArray] class.
type IVertexBufferLayoutDescriptorArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VertexBufferLayoutDescriptorArray */
	// properties:
	MTLBufferLayoutStrideDynamic() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VertexBufferLayoutDescriptorArray */
	// methods:
	SetObjectAtIndexedSubscript(bufferDesc IMTLVertexBufferLayoutDescriptor, index uint)
	ObjectAtIndexedSubscript(index uint) IVertexBufferLayoutDescriptor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VertexBufferLayoutDescriptorArray */
// Alloc allocates a new instance without initialization.
func (vc _VertexBufferLayoutDescriptorArrayClass) Alloc() VertexBufferLayoutDescriptorArray {
	rv := objc.Send[VertexBufferLayoutDescriptorArray](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VertexBufferLayoutDescriptorArrayClass) New() VertexBufferLayoutDescriptorArray {
	rv := objc.Send[VertexBufferLayoutDescriptorArray](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VertexBufferLayoutDescriptorArray) Init() VertexBufferLayoutDescriptorArray {
	rv := objc.Send[VertexBufferLayoutDescriptorArray](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VertexBufferLayoutDescriptorArray) Autorelease() VertexBufferLayoutDescriptorArray {
	rv := objc.Send[VertexBufferLayoutDescriptorArray](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVertexBufferLayoutDescriptorArray creates a new VertexBufferLayoutDescriptorArray instance.
func NewVertexBufferLayoutDescriptorArray() VertexBufferLayoutDescriptorArray {
	return getVertexBufferLayoutDescriptorArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VertexBufferLayoutDescriptorArray */
// An array of vertex buffer layout descriptor instances.
//
// An holds an array of vertex buffer layout states. The methods of set the vertex buffer layout state in the array or retrieve the state from the array.


// An array of vertex buffer layout descriptor instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptorArray
type VertexBufferLayoutDescriptorArray struct {
	objectivec.Object
}

// VertexBufferLayoutDescriptorArrayFrom constructs a [VertexBufferLayoutDescriptorArray] from an unsafe.Pointer.
//
// An array of vertex buffer layout descriptor instances.
func VertexBufferLayoutDescriptorArrayFrom(ptr unsafe.Pointer) VertexBufferLayoutDescriptorArray {
	return VertexBufferLayoutDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VertexBufferLayoutDescriptorArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VertexBufferLayoutDescriptorArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VertexBufferLayoutDescriptorArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VertexBufferLayoutDescriptorArray */

// Sets the state of the specified vertex buffer layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptorArray/setObject:atIndexedSubscript:
func (v_ VertexBufferLayoutDescriptorArray) SetObjectAtIndexedSubscript(bufferDesc IMTLVertexBufferLayoutDescriptor, index uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setObject:atIndexedSubscript:"), bufferDesc, index)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Returns the state of the specified vertex buffer layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptorArray/subscript(_:)
func (v_ VertexBufferLayoutDescriptorArray) ObjectAtIndexedSubscript(index uint) IVertexBufferLayoutDescriptor {
	rv := objc.Send[VertexBufferLayoutDescriptor](v_.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VertexBufferLayoutDescriptorArray */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v_ VertexBufferLayoutDescriptorArray) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v_.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}/* debug [instance_properties/getter]: MTLBufferLayoutStrideDynamic */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLVertexBufferLayoutDescriptorArray */



