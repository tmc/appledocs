// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNSubPixelConvolutionDescriptor] class.
var (
	CNNSubPixelConvolutionDescriptorClass     _CNNSubPixelConvolutionDescriptorClass
	CNNSubPixelConvolutionDescriptorClassOnce sync.Once
)

func getCNNSubPixelConvolutionDescriptorClass() _CNNSubPixelConvolutionDescriptorClass {
	CNNSubPixelConvolutionDescriptorClassOnce.Do(func() {
		CNNSubPixelConvolutionDescriptorClass = _CNNSubPixelConvolutionDescriptorClass{objc.GetClass("MPSCNNSubPixelConvolutionDescriptor")}
	})
	return CNNSubPixelConvolutionDescriptorClass
}

type _CNNSubPixelConvolutionDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [CNNSubPixelConvolutionDescriptor] class.
type ICNNSubPixelConvolutionDescriptor interface {
	ICNNConvolutionDescriptor
	

	// properties:
	SubPixelScaleFactor() objectivec.IObject
	SetSubPixelScaleFactor(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNSubPixelConvolutionDescriptorClass) Alloc() CNNSubPixelConvolutionDescriptor {
	rv := objc.Send[CNNSubPixelConvolutionDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSubPixelConvolutionDescriptorClass) New() CNNSubPixelConvolutionDescriptor {
	rv := objc.Send[CNNSubPixelConvolutionDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSubPixelConvolutionDescriptor) Init() CNNSubPixelConvolutionDescriptor {
	rv := objc.Send[CNNSubPixelConvolutionDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSubPixelConvolutionDescriptor) Autorelease() CNNSubPixelConvolutionDescriptor {
	rv := objc.Send[CNNSubPixelConvolutionDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSubPixelConvolutionDescriptor creates a new CNNSubPixelConvolutionDescriptor instance.
func NewCNNSubPixelConvolutionDescriptor() CNNSubPixelConvolutionDescriptor {
	return getCNNSubPixelConvolutionDescriptorClass().New()
}





// A description of a convolution object that does subpixel upsampling and reshaping.


// A description of a convolution object that does subpixel upsampling and reshaping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubPixelConvolutionDescriptor
type CNNSubPixelConvolutionDescriptor struct {
	CNNConvolutionDescriptor
}

// CNNSubPixelConvolutionDescriptorFrom constructs a [CNNSubPixelConvolutionDescriptor] from an unsafe.Pointer.
//
// A description of a convolution object that does subpixel upsampling and reshaping.
func CNNSubPixelConvolutionDescriptorFrom(ptr unsafe.Pointer) CNNSubPixelConvolutionDescriptor {
	return CNNSubPixelConvolutionDescriptor{
		CNNConvolutionDescriptor: CNNConvolutionDescriptorFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubpixelconvolutiondescriptor/2875156-subpixelscalefactor
func (c_ CNNSubPixelConvolutionDescriptor) SubPixelScaleFactor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("subPixelScaleFactor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubpixelconvolutiondescriptor/2875156-subpixelscalefactor
func (c_ CNNSubPixelConvolutionDescriptor) SetSubPixelScaleFactor(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubPixelScaleFactor:"), value)
}








