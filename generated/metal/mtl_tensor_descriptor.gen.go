// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TensorDescriptor] class.
var (
	TensorDescriptorClass     _TensorDescriptorClass
	TensorDescriptorClassOnce sync.Once
)

func getTensorDescriptorClass() _TensorDescriptorClass {
	TensorDescriptorClassOnce.Do(func() {
		TensorDescriptorClass = _TensorDescriptorClass{objc.GetClass("MTLTensorDescriptor")}
	})
	return TensorDescriptorClass
}

type _TensorDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [TensorDescriptor] class.
type ITensorDescriptor interface {
	objectivec.IObject
}

// A configuration type for creating new tensor instances.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor
type TensorDescriptor struct {
	objectivec.Object
}

// TensorDescriptorFrom constructs a [TensorDescriptor] from an unsafe.Pointer.
//
// A configuration type for creating new tensor instances.
func TensorDescriptorFrom(ptr unsafe.Pointer) TensorDescriptor {
	return TensorDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TensorDescriptorClass) Alloc() TensorDescriptor {
	rv := objc.Send[TensorDescriptor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TensorDescriptorClass) New() TensorDescriptor {
	rv := objc.Send[TensorDescriptor](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TensorDescriptor) Init() TensorDescriptor {
	rv := objc.Send[TensorDescriptor](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TensorDescriptor) Autorelease() TensorDescriptor {
	rv := objc.Send[TensorDescriptor](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTensorDescriptor creates a new TensorDescriptor instance.
func NewTensorDescriptor() TensorDescriptor {
	return getTensorDescriptorClass().New()
}


// An array of sizes, in elements, one for each dimension of the tensors you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/dimensions
func (t_ TensorDescriptor) Dimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("dimensions"))
	return rv
}


// SetDimensions sets the value of the dimensions property.
// An array of sizes, in elements, one for each dimension of the tensors you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTensorDescriptor/dimensions
func (t_ TensorDescriptor) SetDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDimensions:"), value)
}


