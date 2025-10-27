// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MTL4FunctionDescriptor] class.
var (
	MTL4FunctionDescriptorClass     _MTL4FunctionDescriptorClass
	MTL4FunctionDescriptorClassOnce sync.Once
)

func getMTL4FunctionDescriptorClass() _MTL4FunctionDescriptorClass {
	MTL4FunctionDescriptorClassOnce.Do(func() {
		MTL4FunctionDescriptorClass = _MTL4FunctionDescriptorClass{objc.GetClass("MTL4FunctionDescriptor")}
	})
	return MTL4FunctionDescriptorClass
}

type _MTL4FunctionDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4FunctionDescriptor] class.
type IMTL4FunctionDescriptor interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4FunctionDescriptorClass) Alloc() MTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4FunctionDescriptorClass) New() MTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4FunctionDescriptor) Init() MTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4FunctionDescriptor) Autorelease() MTL4FunctionDescriptor {
	rv := objc.Send[MTL4FunctionDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4FunctionDescriptor creates a new MTL4FunctionDescriptor instance.
func NewMTL4FunctionDescriptor() MTL4FunctionDescriptor {
	return getMTL4FunctionDescriptorClass().New()
}





// Base interface for describing a Metal 4 shader function.


// Base interface for describing a Metal 4 shader function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4FunctionDescriptor
type MTL4FunctionDescriptor struct {
	objectivec.Object
}

// MTL4FunctionDescriptorFrom constructs a [MTL4FunctionDescriptor] from an unsafe.Pointer.
//
// Base interface for describing a Metal 4 shader function.
func MTL4FunctionDescriptorFrom(ptr unsafe.Pointer) MTL4FunctionDescriptor {
	return MTL4FunctionDescriptor{objectivec.Object{objc.ID(ptr)}}
}































