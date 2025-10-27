// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MTL4PipelineDescriptor] class.
var (
	MTL4PipelineDescriptorClass     _MTL4PipelineDescriptorClass
	MTL4PipelineDescriptorClassOnce sync.Once
)

func getMTL4PipelineDescriptorClass() _MTL4PipelineDescriptorClass {
	MTL4PipelineDescriptorClassOnce.Do(func() {
		MTL4PipelineDescriptorClass = _MTL4PipelineDescriptorClass{objc.GetClass("MTL4PipelineDescriptor")}
	})
	return MTL4PipelineDescriptorClass
}

type _MTL4PipelineDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4PipelineDescriptor] class.
type IMTL4PipelineDescriptor interface {
	objectivec.IObject
	

	// properties:
	Label() foundation.foundation.INSString
	SetLabel(value foundation.foundation.INSString)
	Options() IMTL4PipelineOptions
	SetOptions(value IMTL4PipelineOptions)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4PipelineDescriptorClass) Alloc() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4PipelineDescriptorClass) New() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4PipelineDescriptor) Init() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4PipelineDescriptor) Autorelease() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PipelineDescriptor creates a new MTL4PipelineDescriptor instance.
func NewMTL4PipelineDescriptor() MTL4PipelineDescriptor {
	return getMTL4PipelineDescriptorClass().New()
}





// Base type for descriptors you use for building pipeline state objects.


// Base type for descriptors you use for building pipeline state objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor
type MTL4PipelineDescriptor struct {
	objectivec.Object
}

// MTL4PipelineDescriptorFrom constructs a [MTL4PipelineDescriptor] from an unsafe.Pointer.
//
// Base type for descriptors you use for building pipeline state objects.
func MTL4PipelineDescriptorFrom(ptr unsafe.Pointer) MTL4PipelineDescriptor {
	return MTL4PipelineDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// Assigns an optional string that uniquely identifies a pipeline descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor/label
func (m_ MTL4PipelineDescriptor) Label() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// Assigns an optional string that uniquely identifies a pipeline descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor/label
func (m_ MTL4PipelineDescriptor) SetLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// Provides compile-time options when you build the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor/options
func (m_ MTL4PipelineDescriptor) Options() IMTL4PipelineOptions {
	rv := objc.Send[MTL4PipelineOptions](m_.ID, objc.Sel("options"))
	return rv
}


// Provides compile-time options when you build the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor/options
func (m_ MTL4PipelineDescriptor) SetOptions(value IMTL4PipelineOptions) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptions:"), value)
}








