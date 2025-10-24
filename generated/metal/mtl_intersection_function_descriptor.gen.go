// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTLIntersectionFunctionDescriptor */


/* debug [class_header]: Header for MTLIntersectionFunctionDescriptor */
// The class instance for the [IntersectionFunctionDescriptor] class.
var (
	IntersectionFunctionDescriptorClass     _IntersectionFunctionDescriptorClass
	IntersectionFunctionDescriptorClassOnce sync.Once
)

func getIntersectionFunctionDescriptorClass() _IntersectionFunctionDescriptorClass {
	IntersectionFunctionDescriptorClassOnce.Do(func() {
		IntersectionFunctionDescriptorClass = _IntersectionFunctionDescriptorClass{objc.GetClass("MTLIntersectionFunctionDescriptor")}
	})
	return IntersectionFunctionDescriptorClass
}

type _IntersectionFunctionDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IntersectionFunctionDescriptor */
// An interface definition for the [IntersectionFunctionDescriptor] class.
type IIntersectionFunctionDescriptor interface {
	IFunctionDescriptor
	
/* debug [class_interface_properties]: Properties for IntersectionFunctionDescriptor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IntersectionFunctionDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IntersectionFunctionDescriptor */
// Alloc allocates a new instance without initialization.
func (ic _IntersectionFunctionDescriptorClass) Alloc() IntersectionFunctionDescriptor {
	rv := objc.Send[IntersectionFunctionDescriptor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IntersectionFunctionDescriptorClass) New() IntersectionFunctionDescriptor {
	rv := objc.Send[IntersectionFunctionDescriptor](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IntersectionFunctionDescriptor) Init() IntersectionFunctionDescriptor {
	rv := objc.Send[IntersectionFunctionDescriptor](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IntersectionFunctionDescriptor) Autorelease() IntersectionFunctionDescriptor {
	rv := objc.Send[IntersectionFunctionDescriptor](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIntersectionFunctionDescriptor creates a new IntersectionFunctionDescriptor instance.
func NewIntersectionFunctionDescriptor() IntersectionFunctionDescriptor {
	return getIntersectionFunctionDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IntersectionFunctionDescriptor */
// A description of an intersection function that performs an intersection test.
//
// This class doesn’t add any additional API over its parent class.


// A description of an intersection function that performs an intersection test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionDescriptor
type IntersectionFunctionDescriptor struct {
	FunctionDescriptor
}

// IntersectionFunctionDescriptorFrom constructs a [IntersectionFunctionDescriptor] from an unsafe.Pointer.
//
// A description of an intersection function that performs an intersection test.
func IntersectionFunctionDescriptorFrom(ptr unsafe.Pointer) IntersectionFunctionDescriptor {
	return IntersectionFunctionDescriptor{
		FunctionDescriptor: FunctionDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IntersectionFunctionDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IntersectionFunctionDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IntersectionFunctionDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IntersectionFunctionDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IntersectionFunctionDescriptor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLIntersectionFunctionDescriptor */



