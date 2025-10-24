// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLIntersectionFunctionTableDescriptor */


/* debug [class_header]: Header for MTLIntersectionFunctionTableDescriptor */
// The class instance for the [IntersectionFunctionTableDescriptor] class.
var (
	IntersectionFunctionTableDescriptorClass     _IntersectionFunctionTableDescriptorClass
	IntersectionFunctionTableDescriptorClassOnce sync.Once
)

func getIntersectionFunctionTableDescriptorClass() _IntersectionFunctionTableDescriptorClass {
	IntersectionFunctionTableDescriptorClassOnce.Do(func() {
		IntersectionFunctionTableDescriptorClass = _IntersectionFunctionTableDescriptorClass{objc.GetClass("MTLIntersectionFunctionTableDescriptor")}
	})
	return IntersectionFunctionTableDescriptorClass
}

type _IntersectionFunctionTableDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IntersectionFunctionTableDescriptor */
// An interface definition for the [IntersectionFunctionTableDescriptor] class.
type IIntersectionFunctionTableDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for IntersectionFunctionTableDescriptor */
	// properties:
	FunctionCount() uint
	SetFunctionCount(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IntersectionFunctionTableDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IntersectionFunctionTableDescriptor */
// Alloc allocates a new instance without initialization.
func (ic _IntersectionFunctionTableDescriptorClass) Alloc() IntersectionFunctionTableDescriptor {
	rv := objc.Send[IntersectionFunctionTableDescriptor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IntersectionFunctionTableDescriptorClass) New() IntersectionFunctionTableDescriptor {
	rv := objc.Send[IntersectionFunctionTableDescriptor](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IntersectionFunctionTableDescriptor) Init() IntersectionFunctionTableDescriptor {
	rv := objc.Send[IntersectionFunctionTableDescriptor](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IntersectionFunctionTableDescriptor) Autorelease() IntersectionFunctionTableDescriptor {
	rv := objc.Send[IntersectionFunctionTableDescriptor](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIntersectionFunctionTableDescriptor creates a new IntersectionFunctionTableDescriptor instance.
func NewIntersectionFunctionTableDescriptor() IntersectionFunctionTableDescriptor {
	return getIntersectionFunctionTableDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IntersectionFunctionTableDescriptor */
// A specification of how to create an intersection function table.


// A specification of how to create an intersection function table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTableDescriptor
type IntersectionFunctionTableDescriptor struct {
	objectivec.Object
}

// IntersectionFunctionTableDescriptorFrom constructs a [IntersectionFunctionTableDescriptor] from an unsafe.Pointer.
//
// A specification of how to create an intersection function table.
func IntersectionFunctionTableDescriptorFrom(ptr unsafe.Pointer) IntersectionFunctionTableDescriptor {
	return IntersectionFunctionTableDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IntersectionFunctionTableDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IntersectionFunctionTableDescriptor */

// Creates an intersection function table descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTableDescriptor/intersectionFunctionTableDescriptor
func (ic _IntersectionFunctionTableDescriptorClass) IntersectionFunctionTableDescriptor() IIntersectionFunctionTableDescriptor {
	rv := objc.Send[IntersectionFunctionTableDescriptor](objc.ID(ic.class), objc.Sel("intersectionFunctionTableDescriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IntersectionFunctionTableDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IntersectionFunctionTableDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IntersectionFunctionTableDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IntersectionFunctionTableDescriptor */

// The number of entries in the intersection function table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTableDescriptor/functionCount
func (i_ IntersectionFunctionTableDescriptor) FunctionCount() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("functionCount"))
	return rv
}/* debug [instance_properties/getter]: functionCount */


// The number of entries in the intersection function table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTableDescriptor/functionCount
func (i_ IntersectionFunctionTableDescriptor) SetFunctionCount(value uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFunctionCount:"), value)
}/* debug [instance_properties/setter]: functionCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLIntersectionFunctionTableDescriptor */



