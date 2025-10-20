// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
}

// A filter subclass that produces an image by manipulating RAW image sensor data from a digital camera or scanner.
//
// Use this class to generate a   object based on the configuration parameters you provide. You can use this object in conjunction with other Core Image classes—such as and —to take advantage of the built-in Core Image filters when processing images or writing custom filters. You can also query this object to find out about the supported camera models, decoders, and filters.
//
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
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(cvPixelBuffer:properties:)
func NewRAWFilterWithCVPixelBufferProperties(buffer unsafe.Pointer, properties objc.ID) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(getRAWFilterClass().class), objc.Sel("filterWithCVPixelBuffer:properties:"), buffer, properties)
	return rv
}

// Creates a RAW filter from the image data and type hint that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageData:identifierHint:)
func NewRAWFilterWithImageDataIdentifierHint(data unsafe.Pointer, identifierHint string) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(getRAWFilterClass().class), objc.Sel("filterWithImageData:identifierHint:"), data, objc.String(identifierHint))
	return rv
}

// Creates a RAW filter from the image at the URL location that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageURL:)
func NewRAWFilterWithImageURL(url unsafe.Pointer) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(getRAWFilterClass().class), objc.Sel("filterWithImageURL:"), url)
	return rv
}


// Creates a RAW filter from the pixel buffer and its properties that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(cvPixelBuffer:properties:)
func (rc _RAWFilterClass) FilterWithCVPixelBufferProperties(buffer unsafe.Pointer, properties objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("filterWithCVPixelBuffer:properties:"), buffer, properties)
	return rv
}

// Creates a RAW filter from the image data and type hint that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageData:identifierHint:)
func (rc _RAWFilterClass) FilterWithImageDataIdentifierHint(data unsafe.Pointer, identifierHint string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("filterWithImageData:identifierHint:"), data, objc.String(identifierHint))
	return rv
}

// Creates a RAW filter from the image at the URL location that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageURL:)
func (rc _RAWFilterClass) FilterWithImageURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("filterWithImageURL:"), url)
	return rv
}

// A value that indicates the baseline exposure to apply to the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/baselineExposure
func (r_ RAWFilter) BaselineExposure() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("baselineExposure"))
	return rv
}


// SetBaselineExposure sets the value of the baselineExposure property.
// A value that indicates the baseline exposure to apply to the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/baselineExposure
func (r_ RAWFilter) SetBaselineExposure(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBaselineExposure:"), value)
}
// A value that indicates the amount of global tone curve to apply to the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostAmount
func (r_ RAWFilter) BoostAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("boostAmount"))
	return rv
}


// SetBoostAmount sets the value of the boostAmount property.
// A value that indicates the amount of global tone curve to apply to the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostAmount
func (r_ RAWFilter) SetBoostAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBoostAmount:"), value)
}
// A value that indicates the amount to boost the shadow areas of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostShadowAmount
func (r_ RAWFilter) BoostShadowAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("boostShadowAmount"))
	return rv
}


// SetBoostShadowAmount sets the value of the boostShadowAmount property.
// A value that indicates the amount to boost the shadow areas of the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostShadowAmount
func (r_ RAWFilter) SetBoostShadowAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBoostShadowAmount:"), value)
}
// A value that indicates the amount of chroma noise reduction to apply to the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/colorNoiseReductionAmount
func (r_ RAWFilter) ColorNoiseReductionAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("colorNoiseReductionAmount"))
	return rv
}


// SetColorNoiseReductionAmount sets the value of the colorNoiseReductionAmount property.
// A value that indicates the amount of chroma noise reduction to apply to the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/colorNoiseReductionAmount
func (r_ RAWFilter) SetColorNoiseReductionAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setColorNoiseReductionAmount:"), value)
}
// A value that indicates the amount of local contrast to apply to the edges of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/contrastAmount
func (r_ RAWFilter) ContrastAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("contrastAmount"))
	return rv
}


// SetContrastAmount sets the value of the contrastAmount property.
// A value that indicates the amount of local contrast to apply to the edges of the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/contrastAmount
func (r_ RAWFilter) SetContrastAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setContrastAmount:"), value)
}
// A value that indicates the decoder version to use.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/decoderVersion
func (r_ RAWFilter) DecoderVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("decoderVersion"))
	return rv
}


// SetDecoderVersion sets the value of the decoderVersion property.
// A value that indicates the decoder version to use.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/decoderVersion
func (r_ RAWFilter) SetDecoderVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDecoderVersion:"), value)
}
// A value that indicates the amount of detail enhancement to apply to the edges of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/detailAmount
func (r_ RAWFilter) DetailAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("detailAmount"))
	return rv
}


// SetDetailAmount sets the value of the detailAmount property.
// A value that indicates the amount of detail enhancement to apply to the edges of the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/detailAmount
func (r_ RAWFilter) SetDetailAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDetailAmount:"), value)
}
// A value that indicates the amount of exposure to apply to the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/exposure
func (r_ RAWFilter) Exposure() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("exposure"))
	return rv
}


// SetExposure sets the value of the exposure property.
// A value that indicates the amount of exposure to apply to the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/exposure
func (r_ RAWFilter) SetExposure(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setExposure:"), value)
}
// A value that indicates the amount of extended dynamic range (EDR) to apply to the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/extendedDynamicRangeAmount
func (r_ RAWFilter) ExtendedDynamicRangeAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("extendedDynamicRangeAmount"))
	return rv
}


// SetExtendedDynamicRangeAmount sets the value of the extendedDynamicRangeAmount property.
// A value that indicates the amount of extended dynamic range (EDR) to apply to the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/extendedDynamicRangeAmount
func (r_ RAWFilter) SetExtendedDynamicRangeAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setExtendedDynamicRangeAmount:"), value)
}
// A Boolean that indicates if the current image supports color noise reduction adjustments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isColorNoiseReductionSupported
func (r_ RAWFilter) ColorNoiseReductionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("colorNoiseReductionSupported"))
	return rv
}

// A Boolean that indicates if the current image supports contrast adjustments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isContrastSupported
func (r_ RAWFilter) ContrastSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("contrastSupported"))
	return rv
}

// A Boolean that indicates if the current image supports detail enhancement adjustments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isDetailSupported
func (r_ RAWFilter) DetailSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("detailSupported"))
	return rv
}

// A Boolean that indicates whether to enable draft mode.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isDraftModeEnabled
func (r_ RAWFilter) DraftModeEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("draftModeEnabled"))
	return rv
}


// SetDraftModeEnabled sets the value of the draftModeEnabled property.
// A Boolean that indicates whether to enable draft mode.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isDraftModeEnabled
func (r_ RAWFilter) SetDraftModeEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDraftModeEnabled:"), value)
}
// A Boolean that indicates whether to enable gamut mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isGamutMappingEnabled
func (r_ RAWFilter) GamutMappingEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("gamutMappingEnabled"))
	return rv
}


// SetGamutMappingEnabled sets the value of the gamutMappingEnabled property.
// A Boolean that indicates whether to enable gamut mapping.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isGamutMappingEnabled
func (r_ RAWFilter) SetGamutMappingEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setGamutMappingEnabled:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isHighlightRecoveryEnabled
func (r_ RAWFilter) HighlightRecoveryEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("highlightRecoveryEnabled"))
	return rv
}


// SetHighlightRecoveryEnabled sets the value of the highlightRecoveryEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isHighlightRecoveryEnabled
func (r_ RAWFilter) SetHighlightRecoveryEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHighlightRecoveryEnabled:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isHighlightRecoverySupported
func (r_ RAWFilter) HighlightRecoverySupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("highlightRecoverySupported"))
	return rv
}

// A Boolean that indicates whether to enable lens correction.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLensCorrectionEnabled
func (r_ RAWFilter) LensCorrectionEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("lensCorrectionEnabled"))
	return rv
}


// SetLensCorrectionEnabled sets the value of the lensCorrectionEnabled property.
// A Boolean that indicates whether to enable lens correction.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLensCorrectionEnabled
func (r_ RAWFilter) SetLensCorrectionEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLensCorrectionEnabled:"), value)
}
// A Boolean that indicates if you can enable lens correction for the current image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLensCorrectionSupported
func (r_ RAWFilter) LensCorrectionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("lensCorrectionSupported"))
	return rv
}

// A Boolean that indicates if the current image supports local tone curve adjustments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLocalToneMapSupported
func (r_ RAWFilter) LocalToneMapSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("localToneMapSupported"))
	return rv
}

// A Boolean that indicates if the current image supports luminance noise reduction adjustments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLuminanceNoiseReductionSupported
func (r_ RAWFilter) LuminanceNoiseReductionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("luminanceNoiseReductionSupported"))
	return rv
}

// A Boolean that indicates if the current image supports moire artifact reduction adjustments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isMoireReductionSupported
func (r_ RAWFilter) MoireReductionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("moireReductionSupported"))
	return rv
}

// A Boolean that indicates if the current image supports sharpness adjustments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isSharpnessSupported
func (r_ RAWFilter) SharpnessSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("sharpnessSupported"))
	return rv
}

// An optional filter you can apply to the RAW image while it’s in linear space.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/linearSpaceFilter
func (r_ RAWFilter) LinearSpaceFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("linearSpaceFilter"))
	return rv
}


// SetLinearSpaceFilter sets the value of the linearSpaceFilter property.
// An optional filter you can apply to the RAW image while it’s in linear space.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/linearSpaceFilter
func (r_ RAWFilter) SetLinearSpaceFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLinearSpaceFilter:"), value)
}
// A value that indicates the amount of local tone curve to apply to the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/localToneMapAmount
func (r_ RAWFilter) LocalToneMapAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("localToneMapAmount"))
	return rv
}


// SetLocalToneMapAmount sets the value of the localToneMapAmount property.
// A value that indicates the amount of local tone curve to apply to the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/localToneMapAmount
func (r_ RAWFilter) SetLocalToneMapAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLocalToneMapAmount:"), value)
}
// A value that indicates the amount of luminance noise reduction to apply to the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/luminanceNoiseReductionAmount
func (r_ RAWFilter) LuminanceNoiseReductionAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("luminanceNoiseReductionAmount"))
	return rv
}


// SetLuminanceNoiseReductionAmount sets the value of the luminanceNoiseReductionAmount property.
// A value that indicates the amount of luminance noise reduction to apply to the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/luminanceNoiseReductionAmount
func (r_ RAWFilter) SetLuminanceNoiseReductionAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLuminanceNoiseReductionAmount:"), value)
}
// A value that indicates the amount of moire artifact reduction to apply to high frequency areas of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/moireReductionAmount
func (r_ RAWFilter) MoireReductionAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("moireReductionAmount"))
	return rv
}


// SetMoireReductionAmount sets the value of the moireReductionAmount property.
// A value that indicates the amount of moire artifact reduction to apply to high frequency areas of the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/moireReductionAmount
func (r_ RAWFilter) SetMoireReductionAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMoireReductionAmount:"), value)
}
// The full native size of the unscaled image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/nativeSize
func (r_ RAWFilter) NativeSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](r_.ID, objc.Sel("nativeSize"))
	return rv
}

// A value that indicates the amount of white balance based on chromaticity values to apply to the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralChromaticity
func (r_ RAWFilter) NeutralChromaticity() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("neutralChromaticity"))
	return rv
}


// SetNeutralChromaticity sets the value of the neutralChromaticity property.
// A value that indicates the amount of white balance based on chromaticity values to apply to the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralChromaticity
func (r_ RAWFilter) SetNeutralChromaticity(value coregraphics.CGPoint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralChromaticity:"), value)
}
// A value that indicates the amount of white balance based on pixel coordinates to apply to the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralLocation
func (r_ RAWFilter) NeutralLocation() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("neutralLocation"))
	return rv
}


// SetNeutralLocation sets the value of the neutralLocation property.
// A value that indicates the amount of white balance based on pixel coordinates to apply to the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralLocation
func (r_ RAWFilter) SetNeutralLocation(value coregraphics.CGPoint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralLocation:"), value)
}
// A value that indicates the amount of white balance based on temperature values to apply to the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTemperature
func (r_ RAWFilter) NeutralTemperature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("neutralTemperature"))
	return rv
}


// SetNeutralTemperature sets the value of the neutralTemperature property.
// A value that indicates the amount of white balance based on temperature values to apply to the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTemperature
func (r_ RAWFilter) SetNeutralTemperature(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralTemperature:"), value)
}
// A value that indicates the amount of white balance based on tint values to apply to the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTint
func (r_ RAWFilter) NeutralTint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("neutralTint"))
	return rv
}


// SetNeutralTint sets the value of the neutralTint property.
// A value that indicates the amount of white balance based on tint values to apply to the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTint
func (r_ RAWFilter) SetNeutralTint(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralTint:"), value)
}
// A value that indicates the orientation of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/orientation
func (r_ RAWFilter) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("orientation"))
	return rv
}


// SetOrientation sets the value of the orientation property.
// A value that indicates the orientation of the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/orientation
func (r_ RAWFilter) SetOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOrientation:"), value)
}
// An optional auxiliary image that represents the portrait effects matte of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/portraitEffectsMatte
func (r_ RAWFilter) PortraitEffectsMatte() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("portraitEffectsMatte"))
	return rv
}

// An optional auxiliary image that represents a preview of the original image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/previewImage
func (r_ RAWFilter) PreviewImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("previewImage"))
	return rv
}

// A dictionary that contains properties of the image source.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/properties
func (r_ RAWFilter) Properties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("properties"))
	return rv
}

// A value that indicates the desired scale factor to draw the output image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/scaleFactor
func (r_ RAWFilter) ScaleFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("scaleFactor"))
	return rv
}


// SetScaleFactor sets the value of the scaleFactor property.
// A value that indicates the desired scale factor to draw the output image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/scaleFactor
func (r_ RAWFilter) SetScaleFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setScaleFactor:"), value)
}
// An optional auxiliary image that represents the semantic segmentation glasses matte of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationGlassesMatte
func (r_ RAWFilter) SemanticSegmentationGlassesMatte() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("semanticSegmentationGlassesMatte"))
	return rv
}

// An optional auxiliary image that represents the semantic segmentation hair matte of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationHairMatte
func (r_ RAWFilter) SemanticSegmentationHairMatte() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("semanticSegmentationHairMatte"))
	return rv
}

// An optional auxiliary image that represents the semantic segmentation skin matte of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationSkinMatte
func (r_ RAWFilter) SemanticSegmentationSkinMatte() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("semanticSegmentationSkinMatte"))
	return rv
}

// An optional auxiliary image that represents the semantic segmentation sky matte of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationSkyMatte
func (r_ RAWFilter) SemanticSegmentationSkyMatte() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("semanticSegmentationSkyMatte"))
	return rv
}

// An optional auxiliary image that represents the semantic segmentation teeth matte of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationTeethMatte
func (r_ RAWFilter) SemanticSegmentationTeethMatte() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("semanticSegmentationTeethMatte"))
	return rv
}

// A value that indicates the amount to subtract from the shadows in the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/shadowBias
func (r_ RAWFilter) ShadowBias() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("shadowBias"))
	return rv
}


// SetShadowBias sets the value of the shadowBias property.
// A value that indicates the amount to subtract from the shadows in the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/shadowBias
func (r_ RAWFilter) SetShadowBias(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setShadowBias:"), value)
}
// A value that indicates the amount of sharpness to apply to the edges of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/sharpnessAmount
func (r_ RAWFilter) SharpnessAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("sharpnessAmount"))
	return rv
}


// SetSharpnessAmount sets the value of the sharpnessAmount property.
// A value that indicates the amount of sharpness to apply to the edges of the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/sharpnessAmount
func (r_ RAWFilter) SetSharpnessAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSharpnessAmount:"), value)
}
// An array of all supported decoder versions for the given image type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/supportedDecoderVersions
func (r_ RAWFilter) SupportedDecoderVersions() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("supportedDecoderVersions"))
	return rv
}


