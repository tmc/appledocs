// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

//
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

// Alloc allocates a new instance without initialization.
func (sc _SVGFClass) Alloc() SVGF {
	rv := objc.Send[SVGF](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/variancePrefilterRadius
func (s_ SVGF) VariancePrefilterRadius() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("variancePrefilterRadius"))
	return rv
}


// SetVariancePrefilterRadius sets the value of the variancePrefilterRadius property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGF/variancePrefilterRadius
func (s_ SVGF) SetVariancePrefilterRadius(value uint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVariancePrefilterRadius:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/bilateralfilterradius
func (s_ SVGF) BilateralFilterRadius() int {
	rv := objc.Send[int](s_.ID, objc.Sel("bilateralFilterRadius"))
	return rv
}


// SetBilateralFilterRadius sets the value of the bilateralFilterRadius property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/bilateralfilterradius
func (s_ SVGF) SetBilateralFilterRadius(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBilateralFilterRadius:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/bilateralfiltersigma
func (s_ SVGF) BilateralFilterSigma() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("bilateralFilterSigma"))
	return rv
}


// SetBilateralFilterSigma sets the value of the bilateralFilterSigma property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/bilateralfiltersigma
func (s_ SVGF) SetBilateralFilterSigma(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBilateralFilterSigma:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/channelcount
func (s_ SVGF) ChannelCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("channelCount"))
	return rv
}


// SetChannelCount sets the value of the channelCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/channelcount
func (s_ SVGF) SetChannelCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setChannelCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/channelcount2
func (s_ SVGF) ChannelCount2() int {
	rv := objc.Send[int](s_.ID, objc.Sel("channelCount2"))
	return rv
}


// SetChannelCount2 sets the value of the channelCount2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/channelcount2
func (s_ SVGF) SetChannelCount2(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setChannelCount2:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/depthweight
func (s_ SVGF) DepthWeight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("depthWeight"))
	return rv
}


// SetDepthWeight sets the value of the depthWeight property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/depthweight
func (s_ SVGF) SetDepthWeight(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDepthWeight:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/luminanceweight
func (s_ SVGF) LuminanceWeight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("luminanceWeight"))
	return rv
}


// SetLuminanceWeight sets the value of the luminanceWeight property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/luminanceweight
func (s_ SVGF) SetLuminanceWeight(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLuminanceWeight:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/minimumframesforvarianceestimation
func (s_ SVGF) MinimumFramesForVarianceEstimation() int {
	rv := objc.Send[int](s_.ID, objc.Sel("minimumFramesForVarianceEstimation"))
	return rv
}


// SetMinimumFramesForVarianceEstimation sets the value of the minimumFramesForVarianceEstimation property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/minimumframesforvarianceestimation
func (s_ SVGF) SetMinimumFramesForVarianceEstimation(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumFramesForVarianceEstimation:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/normalweight
func (s_ SVGF) NormalWeight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("normalWeight"))
	return rv
}


// SetNormalWeight sets the value of the normalWeight property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/normalweight
func (s_ SVGF) SetNormalWeight(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNormalWeight:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/reprojectionthreshold
func (s_ SVGF) ReprojectionThreshold() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("reprojectionThreshold"))
	return rv
}


// SetReprojectionThreshold sets the value of the reprojectionThreshold property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/reprojectionthreshold
func (s_ SVGF) SetReprojectionThreshold(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReprojectionThreshold:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/temporalreprojectionblendfactor
func (s_ SVGF) TemporalReprojectionBlendFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("temporalReprojectionBlendFactor"))
	return rv
}


// SetTemporalReprojectionBlendFactor sets the value of the temporalReprojectionBlendFactor property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/temporalreprojectionblendfactor
func (s_ SVGF) SetTemporalReprojectionBlendFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTemporalReprojectionBlendFactor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/temporalweighting
func (s_ SVGF) TemporalWeighting() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("temporalWeighting"))
	return rv
}


// SetTemporalWeighting sets the value of the temporalWeighting property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/temporalweighting
func (s_ SVGF) SetTemporalWeighting(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTemporalWeighting:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/varianceestimationradius
func (s_ SVGF) VarianceEstimationRadius() int {
	rv := objc.Send[int](s_.ID, objc.Sel("varianceEstimationRadius"))
	return rv
}


// SetVarianceEstimationRadius sets the value of the varianceEstimationRadius property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/varianceestimationradius
func (s_ SVGF) SetVarianceEstimationRadius(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVarianceEstimationRadius:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/varianceestimationsigma
func (s_ SVGF) VarianceEstimationSigma() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("varianceEstimationSigma"))
	return rv
}


// SetVarianceEstimationSigma sets the value of the varianceEstimationSigma property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/varianceestimationsigma
func (s_ SVGF) SetVarianceEstimationSigma(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVarianceEstimationSigma:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/varianceprefiltersigma
func (s_ SVGF) VariancePrefilterSigma() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("variancePrefilterSigma"))
	return rv
}


// SetVariancePrefilterSigma sets the value of the variancePrefilterSigma property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/varianceprefiltersigma
func (s_ SVGF) SetVariancePrefilterSigma(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVariancePrefilterSigma:"), value)
}



