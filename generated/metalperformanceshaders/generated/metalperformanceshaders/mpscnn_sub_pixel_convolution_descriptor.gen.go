// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNSubPixelConvolutionDescriptor */


/* debug [class_header]: Header for MPSCNNSubPixelConvolutionDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNSubPixelConvolutionDescriptor */
// An interface definition for the [CNNSubPixelConvolutionDescriptor] class.
type ICNNSubPixelConvolutionDescriptor interface {
	ICNNConvolutionDescriptor
	
/* debug [class_interface_properties]: Properties for CNNSubPixelConvolutionDescriptor */
	// properties:
	SubPixelScaleFactor() objectivec.IObject
	SetSubPixelScaleFactor(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNSubPixelConvolutionDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNSubPixelConvolutionDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNSubPixelConvolutionDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNSubPixelConvolutionDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNSubPixelConvolutionDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNSubPixelConvolutionDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNSubPixelConvolutionDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNSubPixelConvolutionDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubpixelconvolutiondescriptor/2875156-subpixelscalefactor
func (c_ CNNSubPixelConvolutionDescriptor) SubPixelScaleFactor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("subPixelScaleFactor"))
	return rv
}/* debug [instance_properties/getter]: subPixelScaleFactor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubpixelconvolutiondescriptor/2875156-subpixelscalefactor
func (c_ CNNSubPixelConvolutionDescriptor) SetSubPixelScaleFactor(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubPixelScaleFactor:"), value)
}/* debug [instance_properties/setter]: subPixelScaleFactor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNSubPixelConvolutionDescriptor */



