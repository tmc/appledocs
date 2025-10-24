// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLVertexBufferLayoutDescriptor */


/* debug [class_header]: Header for MTLVertexBufferLayoutDescriptor */
// The class instance for the [VertexBufferLayoutDescriptor] class.
var (
	VertexBufferLayoutDescriptorClass     _VertexBufferLayoutDescriptorClass
	VertexBufferLayoutDescriptorClassOnce sync.Once
)

func getVertexBufferLayoutDescriptorClass() _VertexBufferLayoutDescriptorClass {
	VertexBufferLayoutDescriptorClassOnce.Do(func() {
		VertexBufferLayoutDescriptorClass = _VertexBufferLayoutDescriptorClass{objc.GetClass("MTLVertexBufferLayoutDescriptor")}
	})
	return VertexBufferLayoutDescriptorClass
}

type _VertexBufferLayoutDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VertexBufferLayoutDescriptor */
// An interface definition for the [VertexBufferLayoutDescriptor] class.
type IVertexBufferLayoutDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VertexBufferLayoutDescriptor */
	// properties:
	StepFunction() VertexStepFunction
	SetStepFunction(value VertexStepFunction)
	StepRate() uint
	SetStepRate(value uint)
	Stride() uint
	SetStride(value uint)
	MTLBufferLayoutStrideDynamic() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VertexBufferLayoutDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VertexBufferLayoutDescriptor */
// Alloc allocates a new instance without initialization.
func (vc _VertexBufferLayoutDescriptorClass) Alloc() VertexBufferLayoutDescriptor {
	rv := objc.Send[VertexBufferLayoutDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VertexBufferLayoutDescriptorClass) New() VertexBufferLayoutDescriptor {
	rv := objc.Send[VertexBufferLayoutDescriptor](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VertexBufferLayoutDescriptor) Init() VertexBufferLayoutDescriptor {
	rv := objc.Send[VertexBufferLayoutDescriptor](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VertexBufferLayoutDescriptor) Autorelease() VertexBufferLayoutDescriptor {
	rv := objc.Send[VertexBufferLayoutDescriptor](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVertexBufferLayoutDescriptor creates a new VertexBufferLayoutDescriptor instance.
func NewVertexBufferLayoutDescriptor() VertexBufferLayoutDescriptor {
	return getVertexBufferLayoutDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VertexBufferLayoutDescriptor */
// An object that configures how a render pipeline fetches data to send to the vertex function.


// An object that configures how a render pipeline fetches data to send to the vertex function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor
type VertexBufferLayoutDescriptor struct {
	objectivec.Object
}

// VertexBufferLayoutDescriptorFrom constructs a [VertexBufferLayoutDescriptor] from an unsafe.Pointer.
//
// An object that configures how a render pipeline fetches data to send to the vertex function.
func VertexBufferLayoutDescriptorFrom(ptr unsafe.Pointer) VertexBufferLayoutDescriptor {
	return VertexBufferLayoutDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VertexBufferLayoutDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VertexBufferLayoutDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VertexBufferLayoutDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VertexBufferLayoutDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VertexBufferLayoutDescriptor */

// The circumstances under which the vertex and its attributes are presented to the vertex function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stepFunction
func (v_ VertexBufferLayoutDescriptor) StepFunction() VertexStepFunction {
	rv := objc.Send[VertexStepFunction](v_.ID, objc.Sel("stepFunction"))
	return rv
}/* debug [instance_properties/getter]: stepFunction */


// The circumstances under which the vertex and its attributes are presented to the vertex function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stepFunction
func (v_ VertexBufferLayoutDescriptor) SetStepFunction(value VertexStepFunction) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStepFunction:"), value)
}/* debug [instance_properties/setter]: stepFunction */


// The interval at which the vertex and its attributes are presented to the vertex function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stepRate
func (v_ VertexBufferLayoutDescriptor) StepRate() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("stepRate"))
	return rv
}/* debug [instance_properties/getter]: stepRate */


// The interval at which the vertex and its attributes are presented to the vertex function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stepRate
func (v_ VertexBufferLayoutDescriptor) SetStepRate(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStepRate:"), value)
}/* debug [instance_properties/setter]: stepRate */


// The number of bytes between the first byte of two consecutive vertices in a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stride
func (v_ VertexBufferLayoutDescriptor) Stride() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("stride"))
	return rv
}/* debug [instance_properties/getter]: stride */


// The number of bytes between the first byte of two consecutive vertices in a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexBufferLayoutDescriptor/stride
func (v_ VertexBufferLayoutDescriptor) SetStride(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setStride:"), value)
}/* debug [instance_properties/setter]: stride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutstridedynamic
func (v_ VertexBufferLayoutDescriptor) MTLBufferLayoutStrideDynamic() int {
	rv := objc.Send[int](v_.ID, objc.Sel("MTLBufferLayoutStrideDynamic"))
	return rv
}/* debug [instance_properties/getter]: MTLBufferLayoutStrideDynamic */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLVertexBufferLayoutDescriptor */



