// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNDepthWiseConvolutionDescriptor */


/* debug [class_header]: Header for MPSCNNDepthWiseConvolutionDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNDepthWiseConvolutionDescriptor */
// An interface definition for the [CNNDepthWiseConvolutionDescriptor] class.
type ICNNDepthWiseConvolutionDescriptor interface {
	ICNNConvolutionDescriptor
	
/* debug [class_interface_properties]: Properties for CNNDepthWiseConvolutionDescriptor */
	// properties:
	ChannelMultiplier() objectivec.IObject
	SetChannelMultiplier(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNDepthWiseConvolutionDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNDepthWiseConvolutionDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNDepthWiseConvolutionDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNDepthWiseConvolutionDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNDepthWiseConvolutionDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNDepthWiseConvolutionDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNDepthWiseConvolutionDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNDepthWiseConvolutionDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndepthwiseconvolutiondescriptor/2919731-channelmultiplier
func (c_ CNNDepthWiseConvolutionDescriptor) ChannelMultiplier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("channelMultiplier"))
	return rv
}/* debug [instance_properties/getter]: channelMultiplier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndepthwiseconvolutiondescriptor/2919731-channelmultiplier
func (c_ CNNDepthWiseConvolutionDescriptor) SetChannelMultiplier(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setChannelMultiplier:"), value)
}/* debug [instance_properties/setter]: channelMultiplier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNDepthWiseConvolutionDescriptor */



