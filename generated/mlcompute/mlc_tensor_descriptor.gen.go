// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CTensorDescriptor] class.
var (
	CTensorDescriptorClass     _CTensorDescriptorClass
	CTensorDescriptorClassOnce sync.Once
)

func getCTensorDescriptorClass() _CTensorDescriptorClass {
	CTensorDescriptorClassOnce.Do(func() {
		CTensorDescriptorClass = _CTensorDescriptorClass{objc.GetClass("MLCTensorDescriptor")}
	})
	return CTensorDescriptorClass
}

type _CTensorDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [CTensorDescriptor] class.
type ICTensorDescriptor interface {
	objectivec.IObject
}

// A configuration object you use to create a tensor.
//
// This class contains the mathematical properties of a tensor, such as data type and shape. It also includes initializers that help you create a tensor descriptor for common use cases, such as convolutional neural networks and recurrent neural networks.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorDescriptor
type CTensorDescriptor struct {
	objectivec.Object
}

// CTensorDescriptorFrom constructs a [CTensorDescriptor] from an unsafe.Pointer.
//
// A configuration object you use to create a tensor.
func CTensorDescriptorFrom(ptr unsafe.Pointer) CTensorDescriptor {
	return CTensorDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CTensorDescriptorClass) Alloc() CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CTensorDescriptorClass) New() CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTensorDescriptor) Init() CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTensorDescriptor) Autorelease() CTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTensorDescriptor creates a new CTensorDescriptor instance.
func NewCTensorDescriptor() CTensorDescriptor {
	return getCTensorDescriptorClass().New()
}




