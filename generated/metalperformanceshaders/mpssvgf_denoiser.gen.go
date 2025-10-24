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
	

	// properties:
	BilateralFilterIterations() objectivec.IObject
	SetBilateralFilterIterations(value objectivec.IObject)
	Svgf() IMPSSVGF
	SetSvgf(value IMPSSVGF)
	TextureAllocator() SVGFTextureAllocator get /* not a class type */
	SetTextureAllocator(value SVGFTextureAllocator get /* not a class type */)


	

	// methods:
	ClearTemporalHistory()
	ReleaseTemporaryTextures()
	Encode()
	EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer, sourceTexture2 unsafe.Pointer, destinationTexture2 unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer, previousDepthNormalTexture unsafe.Pointer)
	EncodeToCommandBufferSourceTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer, previousDepthNormalTexture unsafe.Pointer) unsafe.Pointer


}





// Alloc allocates a new instance without initialization.
func (sc _SVGFDenoiserClass) Alloc() SVGFDenoiser {
	rv := objc.Send[SVGFDenoiser](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDenoiser
type SVGFDenoiser struct {
	objectivec.Object
}

// SVGFDenoiserFrom constructs a [SVGFDenoiser] from an unsafe.Pointer.
func SVGFDenoiserFrom(ptr unsafe.Pointer) SVGFDenoiser {
	return SVGFDenoiser{objectivec.Object{objc.ID(ptr)}}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3353094-initwithdevice
func NewSVGFDenoiserWithDevice(device unsafe.Pointer) SVGFDenoiser {
	instance := getSVGFDenoiserClass().Alloc()
	rv := objc.Send[SVGFDenoiser](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242908-initwithsvgf
func NewSVGFDenoiserWithSVGFTextureAllocator(svgf ISVGF, textureAllocator unsafe.Pointer) SVGFDenoiser {
	instance := getSVGFDenoiserClass().Alloc()
	rv := objc.Send[SVGFDenoiser](instance.ID, objc.Sel("initWithSVGF:textureAllocator:"), svgf, textureAllocator)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242903-cleartemporalhistory
func (s_ SVGFDenoiser) ClearTemporalHistory() {
	objc.Send[objc.ID](s_.ID, objc.Sel("clearTemporalHistory"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242911-releasetemporarytextures
func (s_ SVGFDenoiser) ReleaseTemporaryTextures() {
	objc.Send[objc.ID](s_.ID, objc.Sel("releaseTemporaryTextures"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3353092-encode
func (s_ SVGFDenoiser) Encode() {
	objc.Send[objc.ID](s_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3353092-encodetocommandbuffer
func (s_ SVGFDenoiser) EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer, sourceTexture2 unsafe.Pointer, destinationTexture2 unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer, previousDepthNormalTexture unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:sourceTexture2:destinationTexture2:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), commandBuffer, sourceTexture, destinationTexture, sourceTexture2, destinationTexture2, motionVectorTexture, depthNormalTexture, previousDepthNormalTexture)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3353093-encodetocommandbuffer
func (s_ SVGFDenoiser) EncodeToCommandBufferSourceTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer, previousDepthNormalTexture unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), commandBuffer, sourceTexture, motionVectorTexture, depthNormalTexture, previousDepthNormalTexture)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242902-bilateralfilteriterations
func (s_ SVGFDenoiser) BilateralFilterIterations() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("bilateralFilterIterations"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242902-bilateralfilteriterations
func (s_ SVGFDenoiser) SetBilateralFilterIterations(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBilateralFilterIterations:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242914-svgf
func (s_ SVGFDenoiser) Svgf() IMPSSVGF {
	rv := objc.Send[SVGF](s_.ID, objc.Sel("svgf"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242914-svgf
func (s_ SVGFDenoiser) SetSvgf(value IMPSSVGF) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSvgf:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242915-textureallocator
func (s_ SVGFDenoiser) TextureAllocator() SVGFTextureAllocator get /* not a class type */ {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("textureAllocator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242915-textureallocator
func (s_ SVGFDenoiser) SetTextureAllocator(value SVGFTextureAllocator get /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTextureAllocator:"), value)
}







