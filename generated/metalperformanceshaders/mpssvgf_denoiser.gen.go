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
	EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer objectivec.IObject, sourceTexture objectivec.IObject, destinationTexture objectivec.IObject, sourceTexture2 objectivec.IObject, destinationTexture2 objectivec.IObject, motionVectorTexture objectivec.IObject, depthNormalTexture objectivec.IObject, previousDepthNormalTexture objectivec.IObject)
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
func (s_ SVGFDenoiser) EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer objectivec.IObject, sourceTexture objectivec.IObject, destinationTexture objectivec.IObject, sourceTexture2 objectivec.IObject, destinationTexture2 objectivec.IObject, motionVectorTexture objectivec.IObject, depthNormalTexture objectivec.IObject, previousDepthNormalTexture objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:sourceTexture2:destinationTexture2:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), commandBuffer, sourceTexture, destinationTexture, sourceTexture2, destinationTexture2, motionVectorTexture, depthNormalTexture, previousDepthNormalTexture)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/bilateralfilteriterations
func (s_ SVGFDenoiser) BilateralFilterIterations() int {
	rv := objc.Send[int](s_.ID, objc.Sel("bilateralFilterIterations"))
	return rv
}


// SetBilateralFilterIterations sets the value of the bilateralFilterIterations property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/bilateralfilteriterations
func (s_ SVGFDenoiser) SetBilateralFilterIterations(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBilateralFilterIterations:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/svgf
func (s_ SVGFDenoiser) Svgf() MPSSVGF {
	rv := objc.Send[MPSSVGF](s_.ID, objc.Sel("svgf"))
	return rv
}


// SetSvgf sets the value of the svgf property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/svgf
func (s_ SVGFDenoiser) SetSvgf(value IMPSSVGF) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSvgf:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/textureallocator
func (s_ SVGFDenoiser) TextureAllocator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("textureAllocator"))
	return rv
}


// SetTextureAllocator sets the value of the textureAllocator property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/textureallocator
func (s_ SVGFDenoiser) SetTextureAllocator(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTextureAllocator:"), value)
}



