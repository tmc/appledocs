// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNDepthWiseConvolutionDescriptor] class.
var (
	CNNDepthWiseConvolutionDescriptorClass     _CNNDepthWiseConvolutionDescriptorClass
	CNNDepthWiseConvolutionDescriptorClassOnce sync.Once
)

func getCNNDepthWiseConvolutionDescriptorClass() _CNNDepthWiseConvolutionDescriptorClass {
	CNNDepthWiseConvolutionDescriptorClassOnce.Do(func() {
		CNNDepthWiseConvolutionDescriptorClass = _CNNDepthWiseConvolutionDescriptorClass{objc.GetClass("MPSCNNDepthWiseConvolutionDescriptor")}
	})
	return CNNDepthWiseConvolutionDescriptorClass
}

type _CNNDepthWiseConvolutionDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [CNNDepthWiseConvolutionDescriptor] class.
type ICNNDepthWiseConvolutionDescriptor interface {
	ICNNConvolutionDescriptor
	

	// properties:
	ChannelMultiplier() objectivec.IObject
	SetChannelMultiplier(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNDepthWiseConvolutionDescriptorClass) Alloc() CNNDepthWiseConvolutionDescriptor {
	rv := objc.Send[CNNDepthWiseConvolutionDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDepthWiseConvolutionDescriptorClass) New() CNNDepthWiseConvolutionDescriptor {
	rv := objc.Send[CNNDepthWiseConvolutionDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDepthWiseConvolutionDescriptor) Init() CNNDepthWiseConvolutionDescriptor {
	rv := objc.Send[CNNDepthWiseConvolutionDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDepthWiseConvolutionDescriptor) Autorelease() CNNDepthWiseConvolutionDescriptor {
	rv := objc.Send[CNNDepthWiseConvolutionDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDepthWiseConvolutionDescriptor creates a new CNNDepthWiseConvolutionDescriptor instance.
func NewCNNDepthWiseConvolutionDescriptor() CNNDepthWiseConvolutionDescriptor {
	return getCNNDepthWiseConvolutionDescriptorClass().New()
}





// A description of a convolution object that does depthwise convolution.


// A description of a convolution object that does depthwise convolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDepthWiseConvolutionDescriptor
type CNNDepthWiseConvolutionDescriptor struct {
	CNNConvolutionDescriptor
}

// CNNDepthWiseConvolutionDescriptorFrom constructs a [CNNDepthWiseConvolutionDescriptor] from an unsafe.Pointer.
//
// A description of a convolution object that does depthwise convolution.
func CNNDepthWiseConvolutionDescriptorFrom(ptr unsafe.Pointer) CNNDepthWiseConvolutionDescriptor {
	return CNNDepthWiseConvolutionDescriptor{
		CNNConvolutionDescriptor: CNNConvolutionDescriptorFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndepthwiseconvolutiondescriptor/2919731-channelmultiplier
func (c_ CNNDepthWiseConvolutionDescriptor) ChannelMultiplier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("channelMultiplier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndepthwiseconvolutiondescriptor/2919731-channelmultiplier
func (c_ CNNDepthWiseConvolutionDescriptor) SetChannelMultiplier(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setChannelMultiplier:"), value)
}








