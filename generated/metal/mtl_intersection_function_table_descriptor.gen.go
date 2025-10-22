// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [IntersectionFunctionTableDescriptor] class.
type IIntersectionFunctionTableDescriptor interface {
	objectivec.IObject
	FunctionCount() int
	SetFunctionCount(value int)
}

// A specification of how to create an intersection function table.
//
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

// Alloc allocates a new instance without initialization.
func (ic _IntersectionFunctionTableDescriptorClass) Alloc() IntersectionFunctionTableDescriptor {
	rv := objc.Send[IntersectionFunctionTableDescriptor](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The number of entries in the intersection function table.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontabledescriptor/functioncount
func (i_ IntersectionFunctionTableDescriptor) FunctionCount() int {
	rv := objc.Send[int](i_.ID, objc.Sel("functionCount"))
	return rv
}


// SetFunctionCount sets the value of the functionCount property.
// The number of entries in the intersection function table.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctiontabledescriptor/functioncount
func (i_ IntersectionFunctionTableDescriptor) SetFunctionCount(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFunctionCount:"), value)
}



