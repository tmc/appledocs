// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [SVGF] class.
var (
	SVGFClass     _SVGFClass
	SVGFClassOnce sync.Once
)

func getSVGFClass() _SVGFClass {
	SVGFClassOnce.Do(func() {
		SVGFClass = _SVGFClass{objc.GetClass("MPSSVGF")}
	})
	return SVGFClass
}

type _SVGFClass struct {
	class objc.Class
}





// An interface definition for the [SVGF] class.
type ISVGF interface {
	IKernel
	

	// properties:
	BilateralFilterRadius() objectivec.IObject
	SetBilateralFilterRadius(value objectivec.IObject)
	BilateralFilterSigma() objectivec.IObject
	SetBilateralFilterSigma(value objectivec.IObject)
	ChannelCount() objectivec.IObject
	SetChannelCount(value objectivec.IObject)
	ChannelCount2() objectivec.IObject
	SetChannelCount2(value objectivec.IObject)
	DepthWeight() objectivec.IObject
	SetDepthWeight(value objectivec.IObject)
	LuminanceWeight() objectivec.IObject
	SetLuminanceWeight(value objectivec.IObject)
	MinimumFramesForVarianceEstimation() objectivec.IObject
	SetMinimumFramesForVarianceEstimation(value objectivec.IObject)
	NormalWeight() objectivec.IObject
	SetNormalWeight(value objectivec.IObject)
	ReprojectionThreshold() objectivec.IObject
	SetReprojectionThreshold(value objectivec.IObject)
	TemporalReprojectionBlendFactor() objectivec.IObject
	SetTemporalReprojectionBlendFactor(value objectivec.IObject)
	TemporalWeighting() TemporalWeighting get set /* not a class type */
	SetTemporalWeighting(value TemporalWeighting get set /* not a class type */)
	VarianceEstimationRadius() objectivec.IObject
	SetVarianceEstimationRadius(value objectivec.IObject)
	VarianceEstimationSigma() objectivec.IObject
	SetVarianceEstimationSigma(value objectivec.IObject)
	VariancePrefilterRadius() objectivec.IObject
	SetVariancePrefilterRadius(value objectivec.IObject)
	VariancePrefilterSigma() objectivec.IObject
	SetVariancePrefilterSigma(value objectivec.IObject)


	

	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	EncodeBilateralFilter()
	EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureSourceTexture2DestinationTexture2DepthNormalTexture(commandBuffer unsafe.Pointer, stepDistance uint, sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer, sourceTexture2 unsafe.Pointer, destinationTexture2 unsafe.Pointer, depthNormalTexture unsafe.Pointer)
	EncodeReprojection()
	EncodeReprojectionToCommandBufferSourceTexturePreviousTextureDestinationTexturePreviousLuminanceMomentsTextureDestinationLuminanceMomentsTextureSourceTexture2PreviousTexture2DestinationTexture2PreviousLuminanceMomentsTexture2DestinationLuminanceMomentsTexture2PreviousFrameCountTextureDestinationFrameCountTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, previousTexture unsafe.Pointer, destinationTexture unsafe.Pointer, previousLuminanceMomentsTexture unsafe.Pointer, destinationLuminanceMomentsTexture unsafe.Pointer, sourceTexture2 unsafe.Pointer, previousTexture2 unsafe.Pointer, destinationTexture2 unsafe.Pointer, previousLuminanceMomentsTexture2 unsafe.Pointer, destinationLuminanceMomentsTexture2 unsafe.Pointer, previousFrameCountTexture unsafe.Pointer, destinationFrameCountTexture unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer, previousDepthNormalTexture unsafe.Pointer)
	EncodeVarianceEstimation()
	EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureSourceTexture2LuminanceMomentsTexture2DestinationTexture2FrameCountTextureDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, luminanceMomentsTexture unsafe.Pointer, destinationTexture unsafe.Pointer, sourceTexture2 unsafe.Pointer, luminanceMomentsTexture2 unsafe.Pointer, destinationTexture2 unsafe.Pointer, frameCountTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer)
	Encode()
	EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureDepthNormalTexture(commandBuffer unsafe.Pointer, stepDistance uint, sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer)
	EncodeReprojectionToCommandBufferSourceTexturePreviousTextureDestinationTexturePreviousLuminanceMomentsTextureDestinationLuminanceMomentsTexturePreviousFrameCountTextureDestinationFrameCountTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, previousTexture unsafe.Pointer, destinationTexture unsafe.Pointer, previousLuminanceMomentsTexture unsafe.Pointer, destinationLuminanceMomentsTexture unsafe.Pointer, previousFrameCountTexture unsafe.Pointer, destinationFrameCountTexture unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer, previousDepthNormalTexture unsafe.Pointer)
	EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureFrameCountTextureDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, luminanceMomentsTexture unsafe.Pointer, destinationTexture unsafe.Pointer, frameCountTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (sc _SVGFClass) Alloc() SVGF {
	rv := objc.Send[SVGF](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SVGFClass) New() SVGF {
	rv := objc.Send[SVGF](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SVGF) Init() SVGF {
	rv := objc.Send[SVGF](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SVGF) Autorelease() SVGF {
	rv := objc.Send[SVGF](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSVGF creates a new SVGF instance.
func NewSVGF() SVGF {
	return getSVGFClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF
type SVGF struct {
	Kernel
}

// SVGFFrom constructs a [SVGF] from an unsafe.Pointer.
func SVGFFrom(ptr unsafe.Pointer) SVGF {
	return SVGF{
		Kernel: KernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143565-initwithcoder
func NewSVGFWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) SVGF {
	instance := getSVGFClass().Alloc()
	rv := objc.Send[SVGF](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143566-initwithdevice
func NewSVGFWithDevice(device unsafe.Pointer) SVGF {
	instance := getSVGFClass().Alloc()
	rv := objc.Send[SVGF](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143559-copywithzone
func (s_ SVGF) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143561-encodebilateralfilter
func (s_ SVGF) EncodeBilateralFilter() {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeBilateralFilter"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143561-encodebilateralfiltertocommandbu
func (s_ SVGF) EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureSourceTexture2DestinationTexture2DepthNormalTexture(commandBuffer unsafe.Pointer, stepDistance uint, sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer, sourceTexture2 unsafe.Pointer, destinationTexture2 unsafe.Pointer, depthNormalTexture unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeBilateralFilterToCommandBuffer:stepDistance:sourceTexture:destinationTexture:sourceTexture2:destinationTexture2:depthNormalTexture:"), commandBuffer, stepDistance, sourceTexture, destinationTexture, sourceTexture2, destinationTexture2, depthNormalTexture)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143562-encodereprojection
func (s_ SVGF) EncodeReprojection() {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeReprojection"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143562-encodereprojectiontocommandbuffe
func (s_ SVGF) EncodeReprojectionToCommandBufferSourceTexturePreviousTextureDestinationTexturePreviousLuminanceMomentsTextureDestinationLuminanceMomentsTextureSourceTexture2PreviousTexture2DestinationTexture2PreviousLuminanceMomentsTexture2DestinationLuminanceMomentsTexture2PreviousFrameCountTextureDestinationFrameCountTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, previousTexture unsafe.Pointer, destinationTexture unsafe.Pointer, previousLuminanceMomentsTexture unsafe.Pointer, destinationLuminanceMomentsTexture unsafe.Pointer, sourceTexture2 unsafe.Pointer, previousTexture2 unsafe.Pointer, destinationTexture2 unsafe.Pointer, previousLuminanceMomentsTexture2 unsafe.Pointer, destinationLuminanceMomentsTexture2 unsafe.Pointer, previousFrameCountTexture unsafe.Pointer, destinationFrameCountTexture unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer, previousDepthNormalTexture unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeReprojectionToCommandBuffer:sourceTexture:previousTexture:destinationTexture:previousLuminanceMomentsTexture:destinationLuminanceMomentsTexture:sourceTexture2:previousTexture2:destinationTexture2:previousLuminanceMomentsTexture2:destinationLuminanceMomentsTexture2:previousFrameCountTexture:destinationFrameCountTexture:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), commandBuffer, sourceTexture, previousTexture, destinationTexture, previousLuminanceMomentsTexture, destinationLuminanceMomentsTexture, sourceTexture2, previousTexture2, destinationTexture2, previousLuminanceMomentsTexture2, destinationLuminanceMomentsTexture2, previousFrameCountTexture, destinationFrameCountTexture, motionVectorTexture, depthNormalTexture, previousDepthNormalTexture)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143563-encodevarianceestimation
func (s_ SVGF) EncodeVarianceEstimation() {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeVarianceEstimation"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143563-encodevarianceestimationtocomman
func (s_ SVGF) EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureSourceTexture2LuminanceMomentsTexture2DestinationTexture2FrameCountTextureDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, luminanceMomentsTexture unsafe.Pointer, destinationTexture unsafe.Pointer, sourceTexture2 unsafe.Pointer, luminanceMomentsTexture2 unsafe.Pointer, destinationTexture2 unsafe.Pointer, frameCountTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeVarianceEstimationToCommandBuffer:sourceTexture:luminanceMomentsTexture:destinationTexture:sourceTexture2:luminanceMomentsTexture2:destinationTexture2:frameCountTexture:depthNormalTexture:"), commandBuffer, sourceTexture, luminanceMomentsTexture, destinationTexture, sourceTexture2, luminanceMomentsTexture2, destinationTexture2, frameCountTexture, depthNormalTexture)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143564-encode
func (s_ SVGF) Encode() {
	objc.Send[objc.ID](s_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143564-encodewithcoder
func (s_ SVGF) EncodeWithCoder(coder foundation.Coder) {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeWithCoder:"), coder)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3242891-encodebilateralfiltertocommandbu
func (s_ SVGF) EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureDepthNormalTexture(commandBuffer unsafe.Pointer, stepDistance uint, sourceTexture unsafe.Pointer, destinationTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeBilateralFilterToCommandBuffer:stepDistance:sourceTexture:destinationTexture:depthNormalTexture:"), commandBuffer, stepDistance, sourceTexture, destinationTexture, depthNormalTexture)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3242892-encodereprojectiontocommandbuffe
func (s_ SVGF) EncodeReprojectionToCommandBufferSourceTexturePreviousTextureDestinationTexturePreviousLuminanceMomentsTextureDestinationLuminanceMomentsTexturePreviousFrameCountTextureDestinationFrameCountTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, previousTexture unsafe.Pointer, destinationTexture unsafe.Pointer, previousLuminanceMomentsTexture unsafe.Pointer, destinationLuminanceMomentsTexture unsafe.Pointer, previousFrameCountTexture unsafe.Pointer, destinationFrameCountTexture unsafe.Pointer, motionVectorTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer, previousDepthNormalTexture unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeReprojectionToCommandBuffer:sourceTexture:previousTexture:destinationTexture:previousLuminanceMomentsTexture:destinationLuminanceMomentsTexture:previousFrameCountTexture:destinationFrameCountTexture:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), commandBuffer, sourceTexture, previousTexture, destinationTexture, previousLuminanceMomentsTexture, destinationLuminanceMomentsTexture, previousFrameCountTexture, destinationFrameCountTexture, motionVectorTexture, depthNormalTexture, previousDepthNormalTexture)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3242893-encodevarianceestimationtocomman
func (s_ SVGF) EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureFrameCountTextureDepthNormalTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, luminanceMomentsTexture unsafe.Pointer, destinationTexture unsafe.Pointer, frameCountTexture unsafe.Pointer, depthNormalTexture unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("encodeVarianceEstimationToCommandBuffer:sourceTexture:luminanceMomentsTexture:destinationTexture:frameCountTexture:depthNormalTexture:"), commandBuffer, sourceTexture, luminanceMomentsTexture, destinationTexture, frameCountTexture, depthNormalTexture)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143555-bilateralfilterradius
func (s_ SVGF) BilateralFilterRadius() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("bilateralFilterRadius"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143555-bilateralfilterradius
func (s_ SVGF) SetBilateralFilterRadius(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBilateralFilterRadius:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143556-bilateralfiltersigma
func (s_ SVGF) BilateralFilterSigma() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("bilateralFilterSigma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143556-bilateralfiltersigma
func (s_ SVGF) SetBilateralFilterSigma(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBilateralFilterSigma:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143557-channelcount
func (s_ SVGF) ChannelCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("channelCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143557-channelcount
func (s_ SVGF) SetChannelCount(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setChannelCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143558-channelcount2
func (s_ SVGF) ChannelCount2() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("channelCount2"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143558-channelcount2
func (s_ SVGF) SetChannelCount2(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setChannelCount2:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143560-depthweight
func (s_ SVGF) DepthWeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("depthWeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143560-depthweight
func (s_ SVGF) SetDepthWeight(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDepthWeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143567-luminanceweight
func (s_ SVGF) LuminanceWeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("luminanceWeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143567-luminanceweight
func (s_ SVGF) SetLuminanceWeight(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLuminanceWeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143568-minimumframesforvarianceestimati
func (s_ SVGF) MinimumFramesForVarianceEstimation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("minimumFramesForVarianceEstimation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143568-minimumframesforvarianceestimati
func (s_ SVGF) SetMinimumFramesForVarianceEstimation(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumFramesForVarianceEstimation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143569-normalweight
func (s_ SVGF) NormalWeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("normalWeight"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143569-normalweight
func (s_ SVGF) SetNormalWeight(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNormalWeight:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143570-reprojectionthreshold
func (s_ SVGF) ReprojectionThreshold() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("reprojectionThreshold"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143570-reprojectionthreshold
func (s_ SVGF) SetReprojectionThreshold(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReprojectionThreshold:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143571-temporalreprojectionblendfactor
func (s_ SVGF) TemporalReprojectionBlendFactor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("temporalReprojectionBlendFactor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143571-temporalreprojectionblendfactor
func (s_ SVGF) SetTemporalReprojectionBlendFactor(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTemporalReprojectionBlendFactor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143572-temporalweighting
func (s_ SVGF) TemporalWeighting() TemporalWeighting get set /* not a class type */ {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("temporalWeighting"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143572-temporalweighting
func (s_ SVGF) SetTemporalWeighting(value TemporalWeighting get set /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTemporalWeighting:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143573-varianceestimationradius
func (s_ SVGF) VarianceEstimationRadius() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("varianceEstimationRadius"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143573-varianceestimationradius
func (s_ SVGF) SetVarianceEstimationRadius(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVarianceEstimationRadius:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143574-varianceestimationsigma
func (s_ SVGF) VarianceEstimationSigma() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("varianceEstimationSigma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143574-varianceestimationsigma
func (s_ SVGF) SetVarianceEstimationSigma(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVarianceEstimationSigma:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143575-varianceprefilterradius
func (s_ SVGF) VariancePrefilterRadius() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("variancePrefilterRadius"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143575-varianceprefilterradius
func (s_ SVGF) SetVariancePrefilterRadius(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVariancePrefilterRadius:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143576-varianceprefiltersigma
func (s_ SVGF) VariancePrefilterSigma() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("variancePrefilterSigma"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143576-varianceprefiltersigma
func (s_ SVGF) SetVariancePrefilterSigma(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVariancePrefilterSigma:"), value)
}







