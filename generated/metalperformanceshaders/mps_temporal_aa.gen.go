// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TemporalAA] class.
var (
	TemporalAAClass     _TemporalAAClass
	TemporalAAClassOnce sync.Once
)

func getTemporalAAClass() _TemporalAAClass {
	TemporalAAClassOnce.Do(func() {
		TemporalAAClass = _TemporalAAClass{objc.GetClass("MPSTemporalAA")}
	})
	return TemporalAAClass
}

type _TemporalAAClass struct {
	class objc.Class
}

// An interface definition for the [TemporalAA] class.
type ITemporalAA interface {
	IKernel
	EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture(commandBuffer objectivec.IObject, sourceTexture objectivec.IObject, previousTexture objectivec.IObject, destinationTexture objectivec.IObject, motionVectorTexture objectivec.IObject, depthTexture objectivec.IObject)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalAA
type TemporalAA struct {
	Kernel
}

// TemporalAAFrom constructs a [TemporalAA] from an unsafe.Pointer.
func TemporalAAFrom(ptr unsafe.Pointer) TemporalAA {
	return TemporalAA{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TemporalAAClass) Alloc() TemporalAA {
	rv := objc.Send[TemporalAA](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TemporalAAClass) New() TemporalAA {
	rv := objc.Send[TemporalAA](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TemporalAA) Init() TemporalAA {
	rv := objc.Send[TemporalAA](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TemporalAA) Autorelease() TemporalAA {
	rv := objc.Send[TemporalAA](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTemporalAA creates a new TemporalAA instance.
func NewTemporalAA() TemporalAA {
	return getTemporalAAClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTemporalAA/encode(to:sourceTexture:previousTexture:destinationTexture:motionVectorTexture:depthTexture:)
func (t_ TemporalAA) EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture(commandBuffer objectivec.IObject, sourceTexture objectivec.IObject, previousTexture objectivec.IObject, destinationTexture objectivec.IObject, motionVectorTexture objectivec.IObject, depthTexture objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:previousTexture:destinationTexture:motionVectorTexture:depthTexture:"), commandBuffer, sourceTexture, previousTexture, destinationTexture, motionVectorTexture, depthTexture)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/blendfactor
func (t_ TemporalAA) BlendFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("blendFactor"))
	return rv
}


// SetBlendFactor sets the value of the blendFactor property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/blendfactor
func (t_ TemporalAA) SetBlendFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBlendFactor:"), value)
}



