// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SVGFDenoiser] class.
var (
	SVGFDenoiserClass     _SVGFDenoiserClass
	SVGFDenoiserClassOnce sync.Once
)

func getSVGFDenoiserClass() _SVGFDenoiserClass {
	SVGFDenoiserClassOnce.Do(func() {
		SVGFDenoiserClass = _SVGFDenoiserClass{objc.GetClass("MPSSVGFDenoiser")}
	})
	return SVGFDenoiserClass
}

type _SVGFDenoiserClass struct {
	class objc.Class
}

// An interface definition for the [SVGFDenoiser] class.
type ISVGFDenoiser interface {
	objectivec.IObject
	EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer objc.ID, sourceTexture objc.ID, destinationTexture objc.ID, sourceTexture2 objc.ID, destinationTexture2 objc.ID, motionVectorTexture objc.ID, depthNormalTexture objc.ID, previousDepthNormalTexture objc.ID)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser
type SVGFDenoiser struct {
	objectivec.Object
}

// SVGFDenoiserFrom constructs a [SVGFDenoiser] from an unsafe.Pointer.
func SVGFDenoiserFrom(ptr unsafe.Pointer) SVGFDenoiser {
	return SVGFDenoiser{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SVGFDenoiserClass) Alloc() SVGFDenoiser {
	rv := objc.Send[SVGFDenoiser](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SVGFDenoiserClass) New() SVGFDenoiser {
	rv := objc.Send[SVGFDenoiser](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SVGFDenoiser) Init() SVGFDenoiser {
	rv := objc.Send[SVGFDenoiser](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SVGFDenoiser) Autorelease() SVGFDenoiser {
	rv := objc.Send[SVGFDenoiser](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSVGFDenoiser creates a new SVGFDenoiser instance.
func NewSVGFDenoiser() SVGFDenoiser {
	return getSVGFDenoiserClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser/encode(commandBuffer:sourceTexture:destinationTexture:sourceTexture2:destinationTexture2:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:)
func (s_ SVGFDenoiser) EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer objc.ID, sourceTexture objc.ID, destinationTexture objc.ID, sourceTexture2 objc.ID, destinationTexture2 objc.ID, motionVectorTexture objc.ID, depthNormalTexture objc.ID, previousDepthNormalTexture objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:sourceTexture2:destinationTexture2:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), commandBuffer, sourceTexture, destinationTexture, sourceTexture2, destinationTexture2, motionVectorTexture, depthNormalTexture, previousDepthNormalTexture)
}
