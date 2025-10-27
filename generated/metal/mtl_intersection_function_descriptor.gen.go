// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [IntersectionFunctionDescriptor] class.
type IIntersectionFunctionDescriptor interface {
	IFunctionDescriptor
	

	// properties:


	

	// methods:


}





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































