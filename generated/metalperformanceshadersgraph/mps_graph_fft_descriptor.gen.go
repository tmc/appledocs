// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GraphFFTDescriptor] class.
type IGraphFFTDescriptor interface {
	IGraphObject
}

// The class that defines the parameters for a fast Fourier transform (FFT) operation.
//
// Use this descriptor with , , and methods.
//
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

// Alloc allocates a new instance without initialization.
func (gc _GraphFFTDescriptorClass) Alloc() GraphFFTDescriptor {
	rv := objc.Send[GraphFFTDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a fast Fourier transform descriptor with default parameter values.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/descriptor
func (gc _GraphFFTDescriptorClass) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("descriptor"))
	return rv
}

// A Boolean-valued parameter that defines the phase factor sign for Fourier transforms.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphfftdescriptor/inverse
func (g_ GraphFFTDescriptor) Inverse() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("inverse"))
	return rv
}


// SetInverse sets the value of the inverse property.
// A Boolean-valued parameter that defines the phase factor sign for Fourier transforms.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphfftdescriptor/inverse
func (g_ GraphFFTDescriptor) SetInverse(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInverse:"), value)
}

// A parameter which controls how graph rounds the output tensor size for a Hermitean-to-real Fourier transform.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphfftdescriptor/roundtooddhermitean
func (g_ GraphFFTDescriptor) RoundToOddHermitean() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("roundToOddHermitean"))
	return rv
}


// SetRoundToOddHermitean sets the value of the roundToOddHermitean property.
// A parameter which controls how graph rounds the output tensor size for a Hermitean-to-real Fourier transform.

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphfftdescriptor/roundtooddhermitean
func (g_ GraphFFTDescriptor) SetRoundToOddHermitean(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRoundToOddHermitean:"), value)
}

// The scaling mode of the fast fourier transform (FFT) operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/scalingMode
func (g_ GraphFFTDescriptor) ScalingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("scalingMode"))
	return rv
}


// SetScalingMode sets the value of the scalingMode property.
// The scaling mode of the fast fourier transform (FFT) operation.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphFFTDescriptor/scalingMode
func (g_ GraphFFTDescriptor) SetScalingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setScalingMode:"), value)
}



