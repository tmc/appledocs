// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RAWFilter] class.
var (
	RAWFilterClass     _RAWFilterClass
	RAWFilterClassOnce sync.Once
)

func getRAWFilterClass() _RAWFilterClass {
	RAWFilterClassOnce.Do(func() {
		RAWFilterClass = _RAWFilterClass{objc.GetClass("CIRAWFilter")}
	})
	return RAWFilterClass
}

type _RAWFilterClass struct {
	class objc.Class
}

// An interface definition for the [RAWFilter] class.
type IRAWFilter interface {
	IFilter
	// properties:
	BaselineExposure() float32 /* primitive/slice/pointer. */
	SetBaselineExposure(value float32 /* primitive/slice/pointer. */)
	BoostAmount() float32 /* primitive/slice/pointer. */
	SetBoostAmount(value float32 /* primitive/slice/pointer. */)
	BoostShadowAmount() float32 /* primitive/slice/pointer. */
	SetBoostShadowAmount(value float32 /* primitive/slice/pointer. */)
	ColorNoiseReductionAmount() float32 /* primitive/slice/pointer. */
	SetColorNoiseReductionAmount(value float32 /* primitive/slice/pointer. */)
	ContrastAmount() float32 /* primitive/slice/pointer. */
	SetContrastAmount(value float32 /* primitive/slice/pointer. */)
	DecoderVersion() objc.IObject /* cross-framework: RAWDecoderVersion */
	SetDecoderVersion(value objc.IObject /* cross-framework: RAWDecoderVersion */)
	DetailAmount() float32 /* primitive/slice/pointer. */
	SetDetailAmount(value float32 /* primitive/slice/pointer. */)
	Exposure() float32 /* primitive/slice/pointer. */
	SetExposure(value float32 /* primitive/slice/pointer. */)
	ExtendedDynamicRangeAmount() float32 /* primitive/slice/pointer. */
	SetExtendedDynamicRangeAmount(value float32 /* primitive/slice/pointer. */)
	ColorNoiseReductionSupported() bool /* primitive/slice/pointer. */
	ContrastSupported() bool /* primitive/slice/pointer. */
	DetailSupported() bool /* primitive/slice/pointer. */
	DraftModeEnabled() bool /* primitive/slice/pointer. */
	SetDraftModeEnabled(value bool /* primitive/slice/pointer. */)
	GamutMappingEnabled() bool /* primitive/slice/pointer. */
	SetGamutMappingEnabled(value bool /* primitive/slice/pointer. */)
	HighlightRecoveryEnabled() bool /* primitive/slice/pointer. */
	SetHighlightRecoveryEnabled(value bool /* primitive/slice/pointer. */)
	HighlightRecoverySupported() bool /* primitive/slice/pointer. */
	LensCorrectionEnabled() bool /* primitive/slice/pointer. */
	SetLensCorrectionEnabled(value bool /* primitive/slice/pointer. */)
	LensCorrectionSupported() bool /* primitive/slice/pointer. */
	LocalToneMapSupported() bool /* primitive/slice/pointer. */
	LuminanceNoiseReductionSupported() bool /* primitive/slice/pointer. */
	MoireReductionSupported() bool /* primitive/slice/pointer. */
	SharpnessSupported() bool /* primitive/slice/pointer. */
	LinearSpaceFilter() ICIFilter
	SetLinearSpaceFilter(value ICIFilter)
	LocalToneMapAmount() float32 /* primitive/slice/pointer. */
	SetLocalToneMapAmount(value float32 /* primitive/slice/pointer. */)
	LuminanceNoiseReductionAmount() float32 /* primitive/slice/pointer. */
	SetLuminanceNoiseReductionAmount(value float32 /* primitive/slice/pointer. */)
	MoireReductionAmount() float32 /* primitive/slice/pointer. */
	SetMoireReductionAmount(value float32 /* primitive/slice/pointer. */)
	NativeSize() coregraphics.CGSize
	NeutralChromaticity() coregraphics.CGPoint
	SetNeutralChromaticity(value coregraphics.CGPoint)
	NeutralLocation() coregraphics.CGPoint
	SetNeutralLocation(value coregraphics.CGPoint)
	NeutralTemperature() float32 /* primitive/slice/pointer. */
	SetNeutralTemperature(value float32 /* primitive/slice/pointer. */)
	NeutralTint() float32 /* primitive/slice/pointer. */
	SetNeutralTint(value float32 /* primitive/slice/pointer. */)
	Orientation() ImagePropertyOrientation /* not a class type */
	SetOrientation(value ImagePropertyOrientation /* not a class type */)
	PortraitEffectsMatte() ICIImage
	PreviewImage() ICIImage
	Properties() objc.ID
	ScaleFactor() float32 /* primitive/slice/pointer. */
	SetScaleFactor(value float32 /* primitive/slice/pointer. */)
	SemanticSegmentationGlassesMatte() ICIImage
	SemanticSegmentationHairMatte() ICIImage
	SemanticSegmentationSkinMatte() ICIImage
	SemanticSegmentationSkyMatte() ICIImage
	SemanticSegmentationTeethMatte() ICIImage
	ShadowBias() float32 /* primitive/slice/pointer. */
	SetShadowBias(value float32 /* primitive/slice/pointer. */)
	SharpnessAmount() float32 /* primitive/slice/pointer. */
	SetSharpnessAmount(value float32 /* primitive/slice/pointer. */)
	SupportedDecoderVersions() []string /* primitive/slice/pointer. */
	IsColorNoiseReductionSupported() bool /* primitive/slice/pointer. */
	SetIsColorNoiseReductionSupported(value bool /* primitive/slice/pointer. */)
	IsContrastSupported() bool /* primitive/slice/pointer. */
	SetIsContrastSupported(value bool /* primitive/slice/pointer. */)
	IsDetailSupported() bool /* primitive/slice/pointer. */
	SetIsDetailSupported(value bool /* primitive/slice/pointer. */)
	IsDraftModeEnabled() bool /* primitive/slice/pointer. */
	SetIsDraftModeEnabled(value bool /* primitive/slice/pointer. */)
	IsGamutMappingEnabled() bool /* primitive/slice/pointer. */
	SetIsGamutMappingEnabled(value bool /* primitive/slice/pointer. */)
	IsHighlightRecoveryEnabled() bool /* primitive/slice/pointer. */
	SetIsHighlightRecoveryEnabled(value bool /* primitive/slice/pointer. */)
	IsHighlightRecoverySupported() bool /* primitive/slice/pointer. */
	SetIsHighlightRecoverySupported(value bool /* primitive/slice/pointer. */)
	IsLensCorrectionEnabled() bool /* primitive/slice/pointer. */
	SetIsLensCorrectionEnabled(value bool /* primitive/slice/pointer. */)
	IsLensCorrectionSupported() bool /* primitive/slice/pointer. */
	SetIsLensCorrectionSupported(value bool /* primitive/slice/pointer. */)
	IsLocalToneMapSupported() bool /* primitive/slice/pointer. */
	SetIsLocalToneMapSupported(value bool /* primitive/slice/pointer. */)
	IsLuminanceNoiseReductionSupported() bool /* primitive/slice/pointer. */
	SetIsLuminanceNoiseReductionSupported(value bool /* primitive/slice/pointer. */)
	IsMoireReductionSupported() bool /* primitive/slice/pointer. */
	SetIsMoireReductionSupported(value bool /* primitive/slice/pointer. */)
	IsSharpnessSupported() bool /* primitive/slice/pointer. */
	SetIsSharpnessSupported(value bool /* primitive/slice/pointer. */)
	// methods:
}

// A filter subclass that produces an image by manipulating RAW image sensor data from a digital camera or scanner.
//
// Use this class to generate a   object based on the configuration parameters you provide. You can use this object in conjunction with other Core Image classes—such as and —to take advantage of the built-in Core Image filters when processing images or writing custom filters. You can also query this object to find out about the supported camera models, decoders, and filters.


// A filter subclass that produces an image by manipulating RAW image sensor data from a digital camera or scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter
type RAWFilter struct {
	Filter
}

// RAWFilterFrom constructs a [RAWFilter] from an unsafe.Pointer.
//
// A filter subclass that produces an image by manipulating RAW image sensor data from a digital camera or scanner.
func RAWFilterFrom(ptr unsafe.Pointer) RAWFilter {
	return RAWFilter{
		Filter: FilterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RAWFilterClass) Alloc() RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RAWFilterClass) New() RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RAWFilter) Init() RAWFilter {
	rv := objc.Send[RAWFilter](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RAWFilter) Autorelease() RAWFilter {
	rv := objc.Send[RAWFilter](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRAWFilter creates a new RAWFilter instance.
func NewRAWFilter() RAWFilter {
	return getRAWFilterClass().New()
}



// Creates a RAW filter from the pixel buffer and its properties that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(cvPixelBuffer:properties:)
func NewRAWFilterWithCVPixelBufferProperties(buffer PixelBufferRef /* not a class type */, properties objectivec.IObject) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(getRAWFilterClass().class), objc.Sel("filterWithCVPixelBuffer:properties:"), buffer, properties)
	return rv
}


// Creates a RAW filter from the image data and type hint that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageData:identifierHint:)
func NewRAWFilterWithImageDataIdentifierHint(data foundation.objc.IObject /* cross-framework NSData */, identifierHint string /* primitive/slice/pointer. */) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(getRAWFilterClass().class), objc.Sel("filterWithImageData:identifierHint:"), data, objc.String(identifierHint))
	return rv
}


// Creates a RAW filter from the image at the URL location that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageURL:)
func NewRAWFilterWithImageURL(url foundation.objc.IObject /* cross-framework URL */) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(getRAWFilterClass().class), objc.Sel("filterWithImageURL:"), url)
	return rv
}



// Creates a RAW filter from the pixel buffer and its properties that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(cvPixelBuffer:properties:)
func (rc _RAWFilterClass) FilterWithCVPixelBufferProperties(buffer PixelBufferRef /* not a class type */, properties objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("filterWithCVPixelBuffer:properties:"), buffer, properties)
	return rv
}


// Creates a RAW filter from the image data and type hint that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageData:identifierHint:)
func (rc _RAWFilterClass) FilterWithImageDataIdentifierHint(data foundation.objc.IObject /* cross-framework NSData */, identifierHint string /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("filterWithImageData:identifierHint:"), data, objc.String(identifierHint))
	return rv
}


// Creates a RAW filter from the image at the URL location that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageURL:)
func (rc _RAWFilterClass) FilterWithImageURL(url foundation.objc.IObject /* cross-framework URL */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("filterWithImageURL:"), url)
	return rv
}


// An array containing the names of all supported camera models.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/supportedCameraModels
func (rc _RAWFilterClass) SupportedCameraModels() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(rc.class), objc.Sel("supportedCameraModels"))
	return rv
}

// A value that indicates the baseline exposure to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/baselineExposure
func (r_ RAWFilter) BaselineExposure() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("baselineExposure"))
	return rv
}


// A value that indicates the baseline exposure to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/baselineExposure
func (r_ RAWFilter) SetBaselineExposure(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBaselineExposure:"), value)
}


// A value that indicates the amount of global tone curve to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostAmount
func (r_ RAWFilter) BoostAmount() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("boostAmount"))
	return rv
}


// A value that indicates the amount of global tone curve to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostAmount
func (r_ RAWFilter) SetBoostAmount(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBoostAmount:"), value)
}


// A value that indicates the amount to boost the shadow areas of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostShadowAmount
func (r_ RAWFilter) BoostShadowAmount() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("boostShadowAmount"))
	return rv
}


// A value that indicates the amount to boost the shadow areas of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostShadowAmount
func (r_ RAWFilter) SetBoostShadowAmount(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBoostShadowAmount:"), value)
}


// A value that indicates the amount of chroma noise reduction to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/colorNoiseReductionAmount
func (r_ RAWFilter) ColorNoiseReductionAmount() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("colorNoiseReductionAmount"))
	return rv
}


// A value that indicates the amount of chroma noise reduction to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/colorNoiseReductionAmount
func (r_ RAWFilter) SetColorNoiseReductionAmount(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setColorNoiseReductionAmount:"), value)
}


// A value that indicates the amount of local contrast to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/contrastAmount
func (r_ RAWFilter) ContrastAmount() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("contrastAmount"))
	return rv
}


// A value that indicates the amount of local contrast to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/contrastAmount
func (r_ RAWFilter) SetContrastAmount(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setContrastAmount:"), value)
}


// A value that indicates the decoder version to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/decoderVersion
func (r_ RAWFilter) DecoderVersion() objc.IObject /* cross-framework: RAWDecoderVersion */ {
	rv := objc.Send[RAWDecoderVersion](r_.ID, objc.Sel("decoderVersion"))
	return rv
}


// A value that indicates the decoder version to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/decoderVersion
func (r_ RAWFilter) SetDecoderVersion(value objc.IObject /* cross-framework: RAWDecoderVersion */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDecoderVersion:"), value)
}


// A value that indicates the amount of detail enhancement to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/detailAmount
func (r_ RAWFilter) DetailAmount() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("detailAmount"))
	return rv
}


// A value that indicates the amount of detail enhancement to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/detailAmount
func (r_ RAWFilter) SetDetailAmount(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDetailAmount:"), value)
}


// A value that indicates the amount of exposure to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/exposure
func (r_ RAWFilter) Exposure() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("exposure"))
	return rv
}


// A value that indicates the amount of exposure to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/exposure
func (r_ RAWFilter) SetExposure(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setExposure:"), value)
}


// A value that indicates the amount of extended dynamic range (EDR) to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/extendedDynamicRangeAmount
func (r_ RAWFilter) ExtendedDynamicRangeAmount() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("extendedDynamicRangeAmount"))
	return rv
}


// A value that indicates the amount of extended dynamic range (EDR) to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/extendedDynamicRangeAmount
func (r_ RAWFilter) SetExtendedDynamicRangeAmount(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setExtendedDynamicRangeAmount:"), value)
}


// A Boolean that indicates if the current image supports color noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isColorNoiseReductionSupported
func (r_ RAWFilter) ColorNoiseReductionSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("colorNoiseReductionSupported"))
	return rv
}


// A Boolean that indicates if the current image supports contrast adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isContrastSupported
func (r_ RAWFilter) ContrastSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("contrastSupported"))
	return rv
}


// A Boolean that indicates if the current image supports detail enhancement adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isDetailSupported
func (r_ RAWFilter) DetailSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("detailSupported"))
	return rv
}


// A Boolean that indicates whether to enable draft mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isDraftModeEnabled
func (r_ RAWFilter) DraftModeEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("draftModeEnabled"))
	return rv
}


// A Boolean that indicates whether to enable draft mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isDraftModeEnabled
func (r_ RAWFilter) SetDraftModeEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDraftModeEnabled:"), value)
}


// A Boolean that indicates whether to enable gamut mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isGamutMappingEnabled
func (r_ RAWFilter) GamutMappingEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("gamutMappingEnabled"))
	return rv
}


// A Boolean that indicates whether to enable gamut mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isGamutMappingEnabled
func (r_ RAWFilter) SetGamutMappingEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setGamutMappingEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isHighlightRecoveryEnabled
func (r_ RAWFilter) HighlightRecoveryEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("highlightRecoveryEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isHighlightRecoveryEnabled
func (r_ RAWFilter) SetHighlightRecoveryEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHighlightRecoveryEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isHighlightRecoverySupported
func (r_ RAWFilter) HighlightRecoverySupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("highlightRecoverySupported"))
	return rv
}


// A Boolean that indicates whether to enable lens correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLensCorrectionEnabled
func (r_ RAWFilter) LensCorrectionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("lensCorrectionEnabled"))
	return rv
}


// A Boolean that indicates whether to enable lens correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLensCorrectionEnabled
func (r_ RAWFilter) SetLensCorrectionEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLensCorrectionEnabled:"), value)
}


// A Boolean that indicates if you can enable lens correction for the current image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLensCorrectionSupported
func (r_ RAWFilter) LensCorrectionSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("lensCorrectionSupported"))
	return rv
}


// A Boolean that indicates if the current image supports local tone curve adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLocalToneMapSupported
func (r_ RAWFilter) LocalToneMapSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("localToneMapSupported"))
	return rv
}


// A Boolean that indicates if the current image supports luminance noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLuminanceNoiseReductionSupported
func (r_ RAWFilter) LuminanceNoiseReductionSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("luminanceNoiseReductionSupported"))
	return rv
}


// A Boolean that indicates if the current image supports moire artifact reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isMoireReductionSupported
func (r_ RAWFilter) MoireReductionSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("moireReductionSupported"))
	return rv
}


// A Boolean that indicates if the current image supports sharpness adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isSharpnessSupported
func (r_ RAWFilter) SharpnessSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("sharpnessSupported"))
	return rv
}


// An optional filter you can apply to the RAW image while it’s in linear space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/linearSpaceFilter
func (r_ RAWFilter) LinearSpaceFilter() ICIFilter {
	rv := objc.Send[Filter](r_.ID, objc.Sel("linearSpaceFilter"))
	return rv
}


// An optional filter you can apply to the RAW image while it’s in linear space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/linearSpaceFilter
func (r_ RAWFilter) SetLinearSpaceFilter(value ICIFilter) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLinearSpaceFilter:"), value)
}


// A value that indicates the amount of local tone curve to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/localToneMapAmount
func (r_ RAWFilter) LocalToneMapAmount() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("localToneMapAmount"))
	return rv
}


// A value that indicates the amount of local tone curve to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/localToneMapAmount
func (r_ RAWFilter) SetLocalToneMapAmount(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLocalToneMapAmount:"), value)
}


// A value that indicates the amount of luminance noise reduction to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/luminanceNoiseReductionAmount
func (r_ RAWFilter) LuminanceNoiseReductionAmount() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("luminanceNoiseReductionAmount"))
	return rv
}


// A value that indicates the amount of luminance noise reduction to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/luminanceNoiseReductionAmount
func (r_ RAWFilter) SetLuminanceNoiseReductionAmount(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLuminanceNoiseReductionAmount:"), value)
}


// A value that indicates the amount of moire artifact reduction to apply to high frequency areas of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/moireReductionAmount
func (r_ RAWFilter) MoireReductionAmount() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("moireReductionAmount"))
	return rv
}


// A value that indicates the amount of moire artifact reduction to apply to high frequency areas of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/moireReductionAmount
func (r_ RAWFilter) SetMoireReductionAmount(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMoireReductionAmount:"), value)
}


// The full native size of the unscaled image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/nativeSize
func (r_ RAWFilter) NativeSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](r_.ID, objc.Sel("nativeSize"))
	return rv
}


// A value that indicates the amount of white balance based on chromaticity values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralChromaticity
func (r_ RAWFilter) NeutralChromaticity() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("neutralChromaticity"))
	return rv
}


// A value that indicates the amount of white balance based on chromaticity values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralChromaticity
func (r_ RAWFilter) SetNeutralChromaticity(value coregraphics.CGPoint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralChromaticity:"), value)
}


// A value that indicates the amount of white balance based on pixel coordinates to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralLocation
func (r_ RAWFilter) NeutralLocation() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("neutralLocation"))
	return rv
}


// A value that indicates the amount of white balance based on pixel coordinates to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralLocation
func (r_ RAWFilter) SetNeutralLocation(value coregraphics.CGPoint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralLocation:"), value)
}


// A value that indicates the amount of white balance based on temperature values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTemperature
func (r_ RAWFilter) NeutralTemperature() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("neutralTemperature"))
	return rv
}


// A value that indicates the amount of white balance based on temperature values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTemperature
func (r_ RAWFilter) SetNeutralTemperature(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralTemperature:"), value)
}


// A value that indicates the amount of white balance based on tint values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTint
func (r_ RAWFilter) NeutralTint() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("neutralTint"))
	return rv
}


// A value that indicates the amount of white balance based on tint values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTint
func (r_ RAWFilter) SetNeutralTint(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralTint:"), value)
}


// A value that indicates the orientation of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/orientation
func (r_ RAWFilter) Orientation() ImagePropertyOrientation /* not a class type */ {
	rv := objc.Send[ImagePropertyOrientation](r_.ID, objc.Sel("orientation"))
	return rv
}


// A value that indicates the orientation of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/orientation
func (r_ RAWFilter) SetOrientation(value ImagePropertyOrientation /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOrientation:"), value)
}


// An optional auxiliary image that represents the portrait effects matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/portraitEffectsMatte
func (r_ RAWFilter) PortraitEffectsMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("portraitEffectsMatte"))
	return rv
}


// An optional auxiliary image that represents a preview of the original image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/previewImage
func (r_ RAWFilter) PreviewImage() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("previewImage"))
	return rv
}


// A dictionary that contains properties of the image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/properties
func (r_ RAWFilter) Properties() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("properties"))
	return rv
}


// A value that indicates the desired scale factor to draw the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/scaleFactor
func (r_ RAWFilter) ScaleFactor() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("scaleFactor"))
	return rv
}


// A value that indicates the desired scale factor to draw the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/scaleFactor
func (r_ RAWFilter) SetScaleFactor(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setScaleFactor:"), value)
}


// An optional auxiliary image that represents the semantic segmentation glasses matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationGlassesMatte
func (r_ RAWFilter) SemanticSegmentationGlassesMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("semanticSegmentationGlassesMatte"))
	return rv
}


// An optional auxiliary image that represents the semantic segmentation hair matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationHairMatte
func (r_ RAWFilter) SemanticSegmentationHairMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("semanticSegmentationHairMatte"))
	return rv
}


// An optional auxiliary image that represents the semantic segmentation skin matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationSkinMatte
func (r_ RAWFilter) SemanticSegmentationSkinMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("semanticSegmentationSkinMatte"))
	return rv
}


// An optional auxiliary image that represents the semantic segmentation sky matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationSkyMatte
func (r_ RAWFilter) SemanticSegmentationSkyMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("semanticSegmentationSkyMatte"))
	return rv
}


// An optional auxiliary image that represents the semantic segmentation teeth matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationTeethMatte
func (r_ RAWFilter) SemanticSegmentationTeethMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("semanticSegmentationTeethMatte"))
	return rv
}


// A value that indicates the amount to subtract from the shadows in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/shadowBias
func (r_ RAWFilter) ShadowBias() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("shadowBias"))
	return rv
}


// A value that indicates the amount to subtract from the shadows in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/shadowBias
func (r_ RAWFilter) SetShadowBias(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setShadowBias:"), value)
}


// A value that indicates the amount of sharpness to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/sharpnessAmount
func (r_ RAWFilter) SharpnessAmount() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](r_.ID, objc.Sel("sharpnessAmount"))
	return rv
}


// A value that indicates the amount of sharpness to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/sharpnessAmount
func (r_ RAWFilter) SetSharpnessAmount(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSharpnessAmount:"), value)
}


// An array containing the names of all supported camera models.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/supportedCameraModels
func (r_ RAWFilter) SupportedCameraModels() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](r_.ID, objc.Sel("supportedCameraModels"))
	return rv
}


// An array of all supported decoder versions for the given image type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/supportedDecoderVersions
func (r_ RAWFilter) SupportedDecoderVersions() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](r_.ID, objc.Sel("supportedDecoderVersions"))
	return rv
}


// A Boolean that indicates if the current image supports color noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/iscolornoisereductionsupported
func (r_ RAWFilter) IsColorNoiseReductionSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isColorNoiseReductionSupported"))
	return rv
}


// A Boolean that indicates if the current image supports color noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/iscolornoisereductionsupported
func (r_ RAWFilter) SetIsColorNoiseReductionSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsColorNoiseReductionSupported:"), value)
}


// A Boolean that indicates if the current image supports contrast adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/iscontrastsupported
func (r_ RAWFilter) IsContrastSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isContrastSupported"))
	return rv
}


// A Boolean that indicates if the current image supports contrast adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/iscontrastsupported
func (r_ RAWFilter) SetIsContrastSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsContrastSupported:"), value)
}


// A Boolean that indicates if the current image supports detail enhancement adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isdetailsupported
func (r_ RAWFilter) IsDetailSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isDetailSupported"))
	return rv
}


// A Boolean that indicates if the current image supports detail enhancement adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isdetailsupported
func (r_ RAWFilter) SetIsDetailSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsDetailSupported:"), value)
}


// A Boolean that indicates whether to enable draft mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isdraftmodeenabled
func (r_ RAWFilter) IsDraftModeEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isDraftModeEnabled"))
	return rv
}


// A Boolean that indicates whether to enable draft mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isdraftmodeenabled
func (r_ RAWFilter) SetIsDraftModeEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsDraftModeEnabled:"), value)
}


// A Boolean that indicates whether to enable gamut mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isgamutmappingenabled
func (r_ RAWFilter) IsGamutMappingEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isGamutMappingEnabled"))
	return rv
}


// A Boolean that indicates whether to enable gamut mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isgamutmappingenabled
func (r_ RAWFilter) SetIsGamutMappingEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsGamutMappingEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ishighlightrecoveryenabled
func (r_ RAWFilter) IsHighlightRecoveryEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isHighlightRecoveryEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ishighlightrecoveryenabled
func (r_ RAWFilter) SetIsHighlightRecoveryEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsHighlightRecoveryEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ishighlightrecoverysupported
func (r_ RAWFilter) IsHighlightRecoverySupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isHighlightRecoverySupported"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ishighlightrecoverysupported
func (r_ RAWFilter) SetIsHighlightRecoverySupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsHighlightRecoverySupported:"), value)
}


// A Boolean that indicates whether to enable lens correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islenscorrectionenabled
func (r_ RAWFilter) IsLensCorrectionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isLensCorrectionEnabled"))
	return rv
}


// A Boolean that indicates whether to enable lens correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islenscorrectionenabled
func (r_ RAWFilter) SetIsLensCorrectionEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsLensCorrectionEnabled:"), value)
}


// A Boolean that indicates if you can enable lens correction for the current image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islenscorrectionsupported
func (r_ RAWFilter) IsLensCorrectionSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isLensCorrectionSupported"))
	return rv
}


// A Boolean that indicates if you can enable lens correction for the current image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islenscorrectionsupported
func (r_ RAWFilter) SetIsLensCorrectionSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsLensCorrectionSupported:"), value)
}


// A Boolean that indicates if the current image supports local tone curve adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islocaltonemapsupported
func (r_ RAWFilter) IsLocalToneMapSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isLocalToneMapSupported"))
	return rv
}


// A Boolean that indicates if the current image supports local tone curve adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islocaltonemapsupported
func (r_ RAWFilter) SetIsLocalToneMapSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsLocalToneMapSupported:"), value)
}


// A Boolean that indicates if the current image supports luminance noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isluminancenoisereductionsupported
func (r_ RAWFilter) IsLuminanceNoiseReductionSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isLuminanceNoiseReductionSupported"))
	return rv
}


// A Boolean that indicates if the current image supports luminance noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isluminancenoisereductionsupported
func (r_ RAWFilter) SetIsLuminanceNoiseReductionSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsLuminanceNoiseReductionSupported:"), value)
}


// A Boolean that indicates if the current image supports moire artifact reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ismoirereductionsupported
func (r_ RAWFilter) IsMoireReductionSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isMoireReductionSupported"))
	return rv
}


// A Boolean that indicates if the current image supports moire artifact reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ismoirereductionsupported
func (r_ RAWFilter) SetIsMoireReductionSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsMoireReductionSupported:"), value)
}


// A Boolean that indicates if the current image supports sharpness adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/issharpnesssupported
func (r_ RAWFilter) IsSharpnessSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isSharpnessSupported"))
	return rv
}


// A Boolean that indicates if the current image supports sharpness adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/issharpnesssupported
func (r_ RAWFilter) SetIsSharpnessSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsSharpnessSupported:"), value)
}


