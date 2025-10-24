// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSGraphFFTDescriptor */


/* debug [class_header]: Header for MPSGraphFFTDescriptor */
// The class instance for the [GraphFFTDescriptor] class.
var (
	GraphFFTDescriptorClass     _GraphFFTDescriptorClass
	GraphFFTDescriptorClassOnce sync.Once
)

func getGraphFFTDescriptorClass() _GraphFFTDescriptorClass {
	GraphFFTDescriptorClassOnce.Do(func() {
		GraphFFTDescriptorClass = _GraphFFTDescriptorClass{objc.GetClass("MPSGraphFFTDescriptor")}
	})
	return GraphFFTDescriptorClass
}

type _GraphFFTDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphFFTDescriptor */
// An interface definition for the [GraphFFTDescriptor] class.
type IGraphFFTDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphFFTDescriptor */
	// properties:
	Inverse() bool
	SetInverse(value bool)
	RoundToOddHermitean() bool
	SetRoundToOddHermitean(value bool)
	ScalingMode() GraphFFTScalingMode
	SetScalingMode(value GraphFFTScalingMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphFFTDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphFFTDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphFFTDescriptorClass) Alloc() GraphFFTDescriptor {
	rv := objc.Send[GraphFFTDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphFFTDescriptorClass) New() GraphFFTDescriptor {
	rv := objc.Send[GraphFFTDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphFFTDescriptor) Init() GraphFFTDescriptor {
	rv := objc.Send[GraphFFTDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphFFTDescriptor) Autorelease() GraphFFTDescriptor {
	rv := objc.Send[GraphFFTDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphFFTDescriptor creates a new GraphFFTDescriptor instance.
func NewGraphFFTDescriptor() GraphFFTDescriptor {
	return getGraphFFTDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphFFTDescriptor */
// The class that defines the parameters for a fast Fourier transform (FFT) operation.
//
// Use this descriptor with , , and methods.


// The class that defines the parameters for a fast Fourier transform (FFT) operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor
type GraphFFTDescriptor struct {
	GraphObject
}

// GraphFFTDescriptorFrom constructs a [GraphFFTDescriptor] from an unsafe.Pointer.
//
// The class that defines the parameters for a fast Fourier transform (FFT) operation.
func GraphFFTDescriptorFrom(ptr unsafe.Pointer) GraphFFTDescriptor {
	return GraphFFTDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphFFTDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphFFTDescriptor */

// Creates a fast Fourier transform descriptor with default parameter values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/descriptor
func (gc _GraphFFTDescriptorClass) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Descriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphFFTDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphFFTDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphFFTDescriptor */

// A Boolean-valued parameter that defines the phase factor sign for Fourier transforms.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/inverse
func (g_ GraphFFTDescriptor) Inverse() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("inverse"))
	return rv
}/* debug [instance_properties/getter]: inverse */


// A Boolean-valued parameter that defines the phase factor sign for Fourier transforms.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/inverse
func (g_ GraphFFTDescriptor) SetInverse(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInverse:"), value)
}/* debug [instance_properties/setter]: inverse */


// A parameter which controls how graph rounds the output tensor size for a Hermitean-to-real Fourier transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/roundToOddHermitean
func (g_ GraphFFTDescriptor) RoundToOddHermitean() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("roundToOddHermitean"))
	return rv
}/* debug [instance_properties/getter]: roundToOddHermitean */


// A parameter which controls how graph rounds the output tensor size for a Hermitean-to-real Fourier transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/roundToOddHermitean
func (g_ GraphFFTDescriptor) SetRoundToOddHermitean(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRoundToOddHermitean:"), value)
}/* debug [instance_properties/setter]: roundToOddHermitean */


// The scaling mode of the fast fourier transform (FFT) operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/scalingMode
func (g_ GraphFFTDescriptor) ScalingMode() GraphFFTScalingMode {
	rv := objc.Send[GraphFFTScalingMode](g_.ID, objc.Sel("scalingMode"))
	return rv
}/* debug [instance_properties/getter]: scalingMode */


// The scaling mode of the fast fourier transform (FFT) operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/scalingMode
func (g_ GraphFFTDescriptor) SetScalingMode(value GraphFFTScalingMode) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScalingMode:"), value)
}/* debug [instance_properties/setter]: scalingMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphFFTDescriptor */



