// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RNNDescriptor] class.
var (
	RNNDescriptorClass     _RNNDescriptorClass
	RNNDescriptorClassOnce sync.Once
)

func getRNNDescriptorClass() _RNNDescriptorClass {
	RNNDescriptorClassOnce.Do(func() {
		RNNDescriptorClass = _RNNDescriptorClass{objc.GetClass("MPSRNNDescriptor")}
	})
	return RNNDescriptorClass
}

type _RNNDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [RNNDescriptor] class.
type IRNNDescriptor interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other MetalPerformanceShaders classes.


// A parent class referenced by other MetalPerformanceShaders classes. [Full Topic]
type RNNDescriptor struct {
	objectivec.Object
}

// RNNDescriptorFrom constructs a [RNNDescriptor] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
func RNNDescriptorFrom(ptr unsafe.Pointer) RNNDescriptor {
	return RNNDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RNNDescriptorClass) Alloc() RNNDescriptor {
	rv := objc.Send[RNNDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RNNDescriptorClass) New() RNNDescriptor {
	rv := objc.Send[RNNDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNDescriptor) Init() RNNDescriptor {
	rv := objc.Send[RNNDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNDescriptor) Autorelease() RNNDescriptor {
	rv := objc.Send[RNNDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNDescriptor creates a new RNNDescriptor instance.
func NewRNNDescriptor() RNNDescriptor {
	return getRNNDescriptorClass().New()
}




