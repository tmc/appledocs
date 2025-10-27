// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [VisibleFunctionTableDescriptor] class.
var (
	VisibleFunctionTableDescriptorClass     _VisibleFunctionTableDescriptorClass
	VisibleFunctionTableDescriptorClassOnce sync.Once
)

func getVisibleFunctionTableDescriptorClass() _VisibleFunctionTableDescriptorClass {
	VisibleFunctionTableDescriptorClassOnce.Do(func() {
		VisibleFunctionTableDescriptorClass = _VisibleFunctionTableDescriptorClass{objc.GetClass("MTLVisibleFunctionTableDescriptor")}
	})
	return VisibleFunctionTableDescriptorClass
}

type _VisibleFunctionTableDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [VisibleFunctionTableDescriptor] class.
type IVisibleFunctionTableDescriptor interface {
	objectivec.IObject
	

	// properties:
	FunctionCount() uint
	SetFunctionCount(value uint)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _VisibleFunctionTableDescriptorClass) Alloc() VisibleFunctionTableDescriptor {
	rv := objc.Send[VisibleFunctionTableDescriptor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VisibleFunctionTableDescriptorClass) New() VisibleFunctionTableDescriptor {
	rv := objc.Send[VisibleFunctionTableDescriptor](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VisibleFunctionTableDescriptor) Init() VisibleFunctionTableDescriptor {
	rv := objc.Send[VisibleFunctionTableDescriptor](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VisibleFunctionTableDescriptor) Autorelease() VisibleFunctionTableDescriptor {
	rv := objc.Send[VisibleFunctionTableDescriptor](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVisibleFunctionTableDescriptor creates a new VisibleFunctionTableDescriptor instance.
func NewVisibleFunctionTableDescriptor() VisibleFunctionTableDescriptor {
	return getVisibleFunctionTableDescriptorClass().New()
}





// A specification of how to create a visible function table.


// A specification of how to create a visible function table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor
type VisibleFunctionTableDescriptor struct {
	objectivec.Object
}

// VisibleFunctionTableDescriptorFrom constructs a [VisibleFunctionTableDescriptor] from an unsafe.Pointer.
//
// A specification of how to create a visible function table.
func VisibleFunctionTableDescriptorFrom(ptr unsafe.Pointer) VisibleFunctionTableDescriptor {
	return VisibleFunctionTableDescriptor{objectivec.Object{objc.ID(ptr)}}
}










// Creates a default visible function table descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor/visibleFunctionTableDescriptor
func (vc _VisibleFunctionTableDescriptorClass) VisibleFunctionTableDescriptor() IVisibleFunctionTableDescriptor {
	rv := objc.Send[VisibleFunctionTableDescriptor](objc.ID(vc.class), objc.Sel("visibleFunctionTableDescriptor"))
	return rv
}

















// The number of entries in the function table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor/functionCount
func (v_ VisibleFunctionTableDescriptor) FunctionCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("functionCount"))
	return rv
}


// The number of entries in the function table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor/functionCount
func (v_ VisibleFunctionTableDescriptor) SetFunctionCount(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFunctionCount:"), value)
}










