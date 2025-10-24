// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIRAWFilter */


/* debug [class_header]: Header for CIRAWFilter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RAWFilter */
// An interface definition for the [RAWFilter] class.
type IRAWFilter interface {
	IFilter
	
/* debug [class_interface_properties]: Properties for RAWFilter */
	// properties:
	BaselineExposure() float32
	SetBaselineExposure(value float32)
	BoostAmount() float32
	SetBoostAmount(value float32)
	BoostShadowAmount() float32
	SetBoostShadowAmount(value float32)
	ColorNoiseReductionAmount() float32
	SetColorNoiseReductionAmount(value float32)
	ContrastAmount() float32
	SetContrastAmount(value float32)
	DecoderVersion() RAWDecoderVersion /* typedef */
	SetDecoderVersion(value RAWDecoderVersion /* typedef */)
	DetailAmount() float32
	SetDetailAmount(value float32)
	Exposure() float32
	SetExposure(value float32)
	ExtendedDynamicRangeAmount() float32
	SetExtendedDynamicRangeAmount(value float32)
	ColorNoiseReductionSupported() bool
	ContrastSupported() bool
	DetailSupported() bool
	DraftModeEnabled() bool
	SetDraftModeEnabled(value bool)
	GamutMappingEnabled() bool
	SetGamutMappingEnabled(value bool)
	HighlightRecoveryEnabled() bool
	SetHighlightRecoveryEnabled(value bool)
	HighlightRecoverySupported() bool
	LensCorrectionEnabled() bool
	SetLensCorrectionEnabled(value bool)
	LensCorrectionSupported() bool
	LocalToneMapSupported() bool
	LuminanceNoiseReductionSupported() bool
	MoireReductionSupported() bool
	SharpnessSupported() bool
	LinearSpaceFilter() ICIFilter
	SetLinearSpaceFilter(value ICIFilter)
	LocalToneMapAmount() float32
	SetLocalToneMapAmount(value float32)
	LuminanceNoiseReductionAmount() float32
	SetLuminanceNoiseReductionAmount(value float32)
	MoireReductionAmount() float32
	SetMoireReductionAmount(value float32)
	NativeSize() corefoundation.CGSize
	NeutralChromaticity() corefoundation.CGPoint
	SetNeutralChromaticity(value corefoundation.CGPoint)
	NeutralLocation() corefoundation.CGPoint
	SetNeutralLocation(value corefoundation.CGPoint)
	NeutralTemperature() float32
	SetNeutralTemperature(value float32)
	NeutralTint() float32
	SetNeutralTint(value float32)
	Orientation() ImagePropertyOrientation /* not a class type */
	SetOrientation(value ImagePropertyOrientation /* not a class type */)
	PortraitEffectsMatte() ICIImage
	PreviewImage() ICIImage
	Properties() objc.IObject /* cross-framework: NSDictionary */
	ScaleFactor() float32
	SetScaleFactor(value float32)
	SemanticSegmentationGlassesMatte() ICIImage
	SemanticSegmentationHairMatte() ICIImage
	SemanticSegmentationSkinMatte() ICIImage
	SemanticSegmentationSkyMatte() ICIImage
	SemanticSegmentationTeethMatte() ICIImage
	ShadowBias() float32
	SetShadowBias(value float32)
	SharpnessAmount() float32
	SetSharpnessAmount(value float32)
	SupportedDecoderVersions() []string
	IsColorNoiseReductionSupported() bool
	SetIsColorNoiseReductionSupported(value bool)
	IsContrastSupported() bool
	SetIsContrastSupported(value bool)
	IsDetailSupported() bool
	SetIsDetailSupported(value bool)
	IsDraftModeEnabled() bool
	SetIsDraftModeEnabled(value bool)
	IsGamutMappingEnabled() bool
	SetIsGamutMappingEnabled(value bool)
	IsHighlightRecoveryEnabled() bool
	SetIsHighlightRecoveryEnabled(value bool)
	IsHighlightRecoverySupported() bool
	SetIsHighlightRecoverySupported(value bool)
	IsLensCorrectionEnabled() bool
	SetIsLensCorrectionEnabled(value bool)
	IsLensCorrectionSupported() bool
	SetIsLensCorrectionSupported(value bool)
	IsLocalToneMapSupported() bool
	SetIsLocalToneMapSupported(value bool)
	IsLuminanceNoiseReductionSupported() bool
	SetIsLuminanceNoiseReductionSupported(value bool)
	IsMoireReductionSupported() bool
	SetIsMoireReductionSupported(value bool)
	IsSharpnessSupported() bool
	SetIsSharpnessSupported(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RAWFilter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RAWFilter */
// Alloc allocates a new instance without initialization.
func (rc _RAWFilterClass) Alloc() RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RAWFilter */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RAWFilter */

// Creates a RAW filter from the pixel buffer and its properties that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(cvPixelBuffer:properties:)
func NewRAWFilterWithCVPixelBufferProperties(buffer PixelBufferRef /* not a class type */, properties objc.IObject /* cross-framework: NSDictionary */) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(getRAWFilterClass().class), objc.Sel("filterWithCVPixelBuffer:properties:"), buffer, properties)
	return rv
}/* debug [class_init_methods/constructor]: NewRAWFilterWithCVPixelBufferProperties */


// Creates a RAW filter from the image data and type hint that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageData:identifierHint:)
func NewRAWFilterWithImageDataIdentifierHint(data objc.IObject /* cross-framework: NSData */, identifierHint objc.IObject /* cross-framework: NSString */) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(getRAWFilterClass().class), objc.Sel("filterWithImageData:identifierHint:"), data, identifierHint)
	return rv
}/* debug [class_init_methods/constructor]: NewRAWFilterWithImageDataIdentifierHint */


// Creates a RAW filter from the image at the URL location that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageURL:)
func NewRAWFilterWithImageURL(url objc.IObject /* cross-framework: NSURL */) RAWFilter {
	rv := objc.Send[RAWFilter](objc.ID(getRAWFilterClass().class), objc.Sel("filterWithImageURL:"), url)
	return rv
}/* debug [class_init_methods/constructor]: NewRAWFilterWithImageURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RAWFilter */

// Creates a RAW filter from the pixel buffer and its properties that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(cvPixelBuffer:properties:)
func (rc _RAWFilterClass) FilterWithCVPixelBufferProperties(buffer PixelBufferRef /* not a class type */, properties objc.IObject /* cross-framework: NSDictionary */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("filterWithCVPixelBuffer:properties:"), buffer, properties)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FilterWithCVPixelBufferProperties) */


// Creates a RAW filter from the image data and type hint that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageData:identifierHint:)
func (rc _RAWFilterClass) FilterWithImageDataIdentifierHint(data objc.IObject /* cross-framework: NSData */, identifierHint objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("filterWithImageData:identifierHint:"), data, identifierHint)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FilterWithImageDataIdentifierHint) */


// Creates a RAW filter from the image at the URL location that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/init(imageURL:)
func (rc _RAWFilterClass) FilterWithImageURL(url objc.IObject /* cross-framework: NSURL */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("filterWithImageURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FilterWithImageURL) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RAWFilter */

// An array containing the names of all supported camera models.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/supportedCameraModels
func (rc _RAWFilterClass) SupportedCameraModels() []string {
	rv := objc.Send[[]string](objc.ID(rc.class), objc.Sel("supportedCameraModels"))
	return rv
}/* debug [class_properties_class/property]: supportedCameraModels */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RAWFilter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RAWFilter */

// A value that indicates the baseline exposure to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/baselineExposure
func (r_ RAWFilter) BaselineExposure() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("baselineExposure"))
	return rv
}/* debug [instance_properties/getter]: baselineExposure */


// A value that indicates the baseline exposure to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/baselineExposure
func (r_ RAWFilter) SetBaselineExposure(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBaselineExposure:"), value)
}/* debug [instance_properties/setter]: baselineExposure */


// A value that indicates the amount of global tone curve to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostAmount
func (r_ RAWFilter) BoostAmount() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("boostAmount"))
	return rv
}/* debug [instance_properties/getter]: boostAmount */


// A value that indicates the amount of global tone curve to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostAmount
func (r_ RAWFilter) SetBoostAmount(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBoostAmount:"), value)
}/* debug [instance_properties/setter]: boostAmount */


// A value that indicates the amount to boost the shadow areas of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostShadowAmount
func (r_ RAWFilter) BoostShadowAmount() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("boostShadowAmount"))
	return rv
}/* debug [instance_properties/getter]: boostShadowAmount */


// A value that indicates the amount to boost the shadow areas of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/boostShadowAmount
func (r_ RAWFilter) SetBoostShadowAmount(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBoostShadowAmount:"), value)
}/* debug [instance_properties/setter]: boostShadowAmount */


// A value that indicates the amount of chroma noise reduction to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/colorNoiseReductionAmount
func (r_ RAWFilter) ColorNoiseReductionAmount() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("colorNoiseReductionAmount"))
	return rv
}/* debug [instance_properties/getter]: colorNoiseReductionAmount */


// A value that indicates the amount of chroma noise reduction to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/colorNoiseReductionAmount
func (r_ RAWFilter) SetColorNoiseReductionAmount(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setColorNoiseReductionAmount:"), value)
}/* debug [instance_properties/setter]: colorNoiseReductionAmount */


// A value that indicates the amount of local contrast to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/contrastAmount
func (r_ RAWFilter) ContrastAmount() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("contrastAmount"))
	return rv
}/* debug [instance_properties/getter]: contrastAmount */


// A value that indicates the amount of local contrast to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/contrastAmount
func (r_ RAWFilter) SetContrastAmount(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setContrastAmount:"), value)
}/* debug [instance_properties/setter]: contrastAmount */


// A value that indicates the decoder version to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/decoderVersion
func (r_ RAWFilter) DecoderVersion() RAWDecoderVersion /* typedef */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("decoderVersion"))
	return rv
}/* debug [instance_properties/getter]: decoderVersion */


// A value that indicates the decoder version to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/decoderVersion
func (r_ RAWFilter) SetDecoderVersion(value RAWDecoderVersion /* typedef */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDecoderVersion:"), value)
}/* debug [instance_properties/setter]: decoderVersion */


// A value that indicates the amount of detail enhancement to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/detailAmount
func (r_ RAWFilter) DetailAmount() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("detailAmount"))
	return rv
}/* debug [instance_properties/getter]: detailAmount */


// A value that indicates the amount of detail enhancement to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/detailAmount
func (r_ RAWFilter) SetDetailAmount(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDetailAmount:"), value)
}/* debug [instance_properties/setter]: detailAmount */


// A value that indicates the amount of exposure to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/exposure
func (r_ RAWFilter) Exposure() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("exposure"))
	return rv
}/* debug [instance_properties/getter]: exposure */


// A value that indicates the amount of exposure to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/exposure
func (r_ RAWFilter) SetExposure(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setExposure:"), value)
}/* debug [instance_properties/setter]: exposure */


// A value that indicates the amount of extended dynamic range (EDR) to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/extendedDynamicRangeAmount
func (r_ RAWFilter) ExtendedDynamicRangeAmount() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("extendedDynamicRangeAmount"))
	return rv
}/* debug [instance_properties/getter]: extendedDynamicRangeAmount */


// A value that indicates the amount of extended dynamic range (EDR) to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/extendedDynamicRangeAmount
func (r_ RAWFilter) SetExtendedDynamicRangeAmount(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setExtendedDynamicRangeAmount:"), value)
}/* debug [instance_properties/setter]: extendedDynamicRangeAmount */


// A Boolean that indicates if the current image supports color noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isColorNoiseReductionSupported
func (r_ RAWFilter) ColorNoiseReductionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("colorNoiseReductionSupported"))
	return rv
}/* debug [instance_properties/getter]: colorNoiseReductionSupported */


// A Boolean that indicates if the current image supports contrast adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isContrastSupported
func (r_ RAWFilter) ContrastSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("contrastSupported"))
	return rv
}/* debug [instance_properties/getter]: contrastSupported */


// A Boolean that indicates if the current image supports detail enhancement adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isDetailSupported
func (r_ RAWFilter) DetailSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("detailSupported"))
	return rv
}/* debug [instance_properties/getter]: detailSupported */


// A Boolean that indicates whether to enable draft mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isDraftModeEnabled
func (r_ RAWFilter) DraftModeEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("draftModeEnabled"))
	return rv
}/* debug [instance_properties/getter]: draftModeEnabled */


// A Boolean that indicates whether to enable draft mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isDraftModeEnabled
func (r_ RAWFilter) SetDraftModeEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDraftModeEnabled:"), value)
}/* debug [instance_properties/setter]: draftModeEnabled */


// A Boolean that indicates whether to enable gamut mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isGamutMappingEnabled
func (r_ RAWFilter) GamutMappingEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("gamutMappingEnabled"))
	return rv
}/* debug [instance_properties/getter]: gamutMappingEnabled */


// A Boolean that indicates whether to enable gamut mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isGamutMappingEnabled
func (r_ RAWFilter) SetGamutMappingEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setGamutMappingEnabled:"), value)
}/* debug [instance_properties/setter]: gamutMappingEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isHighlightRecoveryEnabled
func (r_ RAWFilter) HighlightRecoveryEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("highlightRecoveryEnabled"))
	return rv
}/* debug [instance_properties/getter]: highlightRecoveryEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isHighlightRecoveryEnabled
func (r_ RAWFilter) SetHighlightRecoveryEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHighlightRecoveryEnabled:"), value)
}/* debug [instance_properties/setter]: highlightRecoveryEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isHighlightRecoverySupported
func (r_ RAWFilter) HighlightRecoverySupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("highlightRecoverySupported"))
	return rv
}/* debug [instance_properties/getter]: highlightRecoverySupported */


// A Boolean that indicates whether to enable lens correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLensCorrectionEnabled
func (r_ RAWFilter) LensCorrectionEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("lensCorrectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: lensCorrectionEnabled */


// A Boolean that indicates whether to enable lens correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLensCorrectionEnabled
func (r_ RAWFilter) SetLensCorrectionEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLensCorrectionEnabled:"), value)
}/* debug [instance_properties/setter]: lensCorrectionEnabled */


// A Boolean that indicates if you can enable lens correction for the current image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLensCorrectionSupported
func (r_ RAWFilter) LensCorrectionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("lensCorrectionSupported"))
	return rv
}/* debug [instance_properties/getter]: lensCorrectionSupported */


// A Boolean that indicates if the current image supports local tone curve adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLocalToneMapSupported
func (r_ RAWFilter) LocalToneMapSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("localToneMapSupported"))
	return rv
}/* debug [instance_properties/getter]: localToneMapSupported */


// A Boolean that indicates if the current image supports luminance noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isLuminanceNoiseReductionSupported
func (r_ RAWFilter) LuminanceNoiseReductionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("luminanceNoiseReductionSupported"))
	return rv
}/* debug [instance_properties/getter]: luminanceNoiseReductionSupported */


// A Boolean that indicates if the current image supports moire artifact reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isMoireReductionSupported
func (r_ RAWFilter) MoireReductionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("moireReductionSupported"))
	return rv
}/* debug [instance_properties/getter]: moireReductionSupported */


// A Boolean that indicates if the current image supports sharpness adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/isSharpnessSupported
func (r_ RAWFilter) SharpnessSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("sharpnessSupported"))
	return rv
}/* debug [instance_properties/getter]: sharpnessSupported */


// An optional filter you can apply to the RAW image while it’s in linear space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/linearSpaceFilter
func (r_ RAWFilter) LinearSpaceFilter() ICIFilter {
	rv := objc.Send[Filter](r_.ID, objc.Sel("linearSpaceFilter"))
	return rv
}/* debug [instance_properties/getter]: linearSpaceFilter */


// An optional filter you can apply to the RAW image while it’s in linear space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/linearSpaceFilter
func (r_ RAWFilter) SetLinearSpaceFilter(value ICIFilter) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLinearSpaceFilter:"), value)
}/* debug [instance_properties/setter]: linearSpaceFilter */


// A value that indicates the amount of local tone curve to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/localToneMapAmount
func (r_ RAWFilter) LocalToneMapAmount() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("localToneMapAmount"))
	return rv
}/* debug [instance_properties/getter]: localToneMapAmount */


// A value that indicates the amount of local tone curve to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/localToneMapAmount
func (r_ RAWFilter) SetLocalToneMapAmount(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLocalToneMapAmount:"), value)
}/* debug [instance_properties/setter]: localToneMapAmount */


// A value that indicates the amount of luminance noise reduction to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/luminanceNoiseReductionAmount
func (r_ RAWFilter) LuminanceNoiseReductionAmount() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("luminanceNoiseReductionAmount"))
	return rv
}/* debug [instance_properties/getter]: luminanceNoiseReductionAmount */


// A value that indicates the amount of luminance noise reduction to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/luminanceNoiseReductionAmount
func (r_ RAWFilter) SetLuminanceNoiseReductionAmount(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLuminanceNoiseReductionAmount:"), value)
}/* debug [instance_properties/setter]: luminanceNoiseReductionAmount */


// A value that indicates the amount of moire artifact reduction to apply to high frequency areas of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/moireReductionAmount
func (r_ RAWFilter) MoireReductionAmount() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("moireReductionAmount"))
	return rv
}/* debug [instance_properties/getter]: moireReductionAmount */


// A value that indicates the amount of moire artifact reduction to apply to high frequency areas of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/moireReductionAmount
func (r_ RAWFilter) SetMoireReductionAmount(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMoireReductionAmount:"), value)
}/* debug [instance_properties/setter]: moireReductionAmount */


// The full native size of the unscaled image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/nativeSize
func (r_ RAWFilter) NativeSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](r_.ID, objc.Sel("nativeSize"))
	return rv
}/* debug [instance_properties/getter]: nativeSize */


// A value that indicates the amount of white balance based on chromaticity values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralChromaticity
func (r_ RAWFilter) NeutralChromaticity() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r_.ID, objc.Sel("neutralChromaticity"))
	return rv
}/* debug [instance_properties/getter]: neutralChromaticity */


// A value that indicates the amount of white balance based on chromaticity values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralChromaticity
func (r_ RAWFilter) SetNeutralChromaticity(value corefoundation.CGPoint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralChromaticity:"), value)
}/* debug [instance_properties/setter]: neutralChromaticity */


// A value that indicates the amount of white balance based on pixel coordinates to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralLocation
func (r_ RAWFilter) NeutralLocation() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r_.ID, objc.Sel("neutralLocation"))
	return rv
}/* debug [instance_properties/getter]: neutralLocation */


// A value that indicates the amount of white balance based on pixel coordinates to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralLocation
func (r_ RAWFilter) SetNeutralLocation(value corefoundation.CGPoint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralLocation:"), value)
}/* debug [instance_properties/setter]: neutralLocation */


// A value that indicates the amount of white balance based on temperature values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTemperature
func (r_ RAWFilter) NeutralTemperature() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("neutralTemperature"))
	return rv
}/* debug [instance_properties/getter]: neutralTemperature */


// A value that indicates the amount of white balance based on temperature values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTemperature
func (r_ RAWFilter) SetNeutralTemperature(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralTemperature:"), value)
}/* debug [instance_properties/setter]: neutralTemperature */


// A value that indicates the amount of white balance based on tint values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTint
func (r_ RAWFilter) NeutralTint() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("neutralTint"))
	return rv
}/* debug [instance_properties/getter]: neutralTint */


// A value that indicates the amount of white balance based on tint values to apply to the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/neutralTint
func (r_ RAWFilter) SetNeutralTint(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNeutralTint:"), value)
}/* debug [instance_properties/setter]: neutralTint */


// A value that indicates the orientation of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/orientation
func (r_ RAWFilter) Orientation() ImagePropertyOrientation /* not a class type */ {
	rv := objc.Send[ImagePropertyOrientation](r_.ID, objc.Sel("orientation"))
	return rv
}/* debug [instance_properties/getter]: orientation */


// A value that indicates the orientation of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/orientation
func (r_ RAWFilter) SetOrientation(value ImagePropertyOrientation /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOrientation:"), value)
}/* debug [instance_properties/setter]: orientation */


// An optional auxiliary image that represents the portrait effects matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/portraitEffectsMatte
func (r_ RAWFilter) PortraitEffectsMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("portraitEffectsMatte"))
	return rv
}/* debug [instance_properties/getter]: portraitEffectsMatte */


// An optional auxiliary image that represents a preview of the original image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/previewImage
func (r_ RAWFilter) PreviewImage() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("previewImage"))
	return rv
}/* debug [instance_properties/getter]: previewImage */


// A dictionary that contains properties of the image source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/properties
func (r_ RAWFilter) Properties() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](r_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_properties/getter]: properties */


// A value that indicates the desired scale factor to draw the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/scaleFactor
func (r_ RAWFilter) ScaleFactor() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("scaleFactor"))
	return rv
}/* debug [instance_properties/getter]: scaleFactor */


// A value that indicates the desired scale factor to draw the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/scaleFactor
func (r_ RAWFilter) SetScaleFactor(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setScaleFactor:"), value)
}/* debug [instance_properties/setter]: scaleFactor */


// An optional auxiliary image that represents the semantic segmentation glasses matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationGlassesMatte
func (r_ RAWFilter) SemanticSegmentationGlassesMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("semanticSegmentationGlassesMatte"))
	return rv
}/* debug [instance_properties/getter]: semanticSegmentationGlassesMatte */


// An optional auxiliary image that represents the semantic segmentation hair matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationHairMatte
func (r_ RAWFilter) SemanticSegmentationHairMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("semanticSegmentationHairMatte"))
	return rv
}/* debug [instance_properties/getter]: semanticSegmentationHairMatte */


// An optional auxiliary image that represents the semantic segmentation skin matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationSkinMatte
func (r_ RAWFilter) SemanticSegmentationSkinMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("semanticSegmentationSkinMatte"))
	return rv
}/* debug [instance_properties/getter]: semanticSegmentationSkinMatte */


// An optional auxiliary image that represents the semantic segmentation sky matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationSkyMatte
func (r_ RAWFilter) SemanticSegmentationSkyMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("semanticSegmentationSkyMatte"))
	return rv
}/* debug [instance_properties/getter]: semanticSegmentationSkyMatte */


// An optional auxiliary image that represents the semantic segmentation teeth matte of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/semanticSegmentationTeethMatte
func (r_ RAWFilter) SemanticSegmentationTeethMatte() ICIImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("semanticSegmentationTeethMatte"))
	return rv
}/* debug [instance_properties/getter]: semanticSegmentationTeethMatte */


// A value that indicates the amount to subtract from the shadows in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/shadowBias
func (r_ RAWFilter) ShadowBias() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("shadowBias"))
	return rv
}/* debug [instance_properties/getter]: shadowBias */


// A value that indicates the amount to subtract from the shadows in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/shadowBias
func (r_ RAWFilter) SetShadowBias(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setShadowBias:"), value)
}/* debug [instance_properties/setter]: shadowBias */


// A value that indicates the amount of sharpness to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/sharpnessAmount
func (r_ RAWFilter) SharpnessAmount() float32 {
	rv := objc.Send[float32](r_.ID, objc.Sel("sharpnessAmount"))
	return rv
}/* debug [instance_properties/getter]: sharpnessAmount */


// A value that indicates the amount of sharpness to apply to the edges of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/sharpnessAmount
func (r_ RAWFilter) SetSharpnessAmount(value float32) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSharpnessAmount:"), value)
}/* debug [instance_properties/setter]: sharpnessAmount */


// An array containing the names of all supported camera models.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/supportedCameraModels
func (r_ RAWFilter) SupportedCameraModels() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("supportedCameraModels"))
	return rv
}/* debug [instance_properties/getter]: supportedCameraModels */


// An array of all supported decoder versions for the given image type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRAWFilter/supportedDecoderVersions
func (r_ RAWFilter) SupportedDecoderVersions() []string {
	rv := objc.Send[[]string](r_.ID, objc.Sel("supportedDecoderVersions"))
	return rv
}/* debug [instance_properties/getter]: supportedDecoderVersions */


// A Boolean that indicates if the current image supports color noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/iscolornoisereductionsupported
func (r_ RAWFilter) IsColorNoiseReductionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isColorNoiseReductionSupported"))
	return rv
}/* debug [instance_properties/getter]: isColorNoiseReductionSupported */


// A Boolean that indicates if the current image supports color noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/iscolornoisereductionsupported
func (r_ RAWFilter) SetIsColorNoiseReductionSupported(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsColorNoiseReductionSupported:"), value)
}/* debug [instance_properties/setter]: isColorNoiseReductionSupported */


// A Boolean that indicates if the current image supports contrast adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/iscontrastsupported
func (r_ RAWFilter) IsContrastSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isContrastSupported"))
	return rv
}/* debug [instance_properties/getter]: isContrastSupported */


// A Boolean that indicates if the current image supports contrast adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/iscontrastsupported
func (r_ RAWFilter) SetIsContrastSupported(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsContrastSupported:"), value)
}/* debug [instance_properties/setter]: isContrastSupported */


// A Boolean that indicates if the current image supports detail enhancement adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isdetailsupported
func (r_ RAWFilter) IsDetailSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isDetailSupported"))
	return rv
}/* debug [instance_properties/getter]: isDetailSupported */


// A Boolean that indicates if the current image supports detail enhancement adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isdetailsupported
func (r_ RAWFilter) SetIsDetailSupported(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsDetailSupported:"), value)
}/* debug [instance_properties/setter]: isDetailSupported */


// A Boolean that indicates whether to enable draft mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isdraftmodeenabled
func (r_ RAWFilter) IsDraftModeEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isDraftModeEnabled"))
	return rv
}/* debug [instance_properties/getter]: isDraftModeEnabled */


// A Boolean that indicates whether to enable draft mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isdraftmodeenabled
func (r_ RAWFilter) SetIsDraftModeEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsDraftModeEnabled:"), value)
}/* debug [instance_properties/setter]: isDraftModeEnabled */


// A Boolean that indicates whether to enable gamut mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isgamutmappingenabled
func (r_ RAWFilter) IsGamutMappingEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isGamutMappingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isGamutMappingEnabled */


// A Boolean that indicates whether to enable gamut mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isgamutmappingenabled
func (r_ RAWFilter) SetIsGamutMappingEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsGamutMappingEnabled:"), value)
}/* debug [instance_properties/setter]: isGamutMappingEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ishighlightrecoveryenabled
func (r_ RAWFilter) IsHighlightRecoveryEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isHighlightRecoveryEnabled"))
	return rv
}/* debug [instance_properties/getter]: isHighlightRecoveryEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ishighlightrecoveryenabled
func (r_ RAWFilter) SetIsHighlightRecoveryEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsHighlightRecoveryEnabled:"), value)
}/* debug [instance_properties/setter]: isHighlightRecoveryEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ishighlightrecoverysupported
func (r_ RAWFilter) IsHighlightRecoverySupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isHighlightRecoverySupported"))
	return rv
}/* debug [instance_properties/getter]: isHighlightRecoverySupported */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ishighlightrecoverysupported
func (r_ RAWFilter) SetIsHighlightRecoverySupported(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsHighlightRecoverySupported:"), value)
}/* debug [instance_properties/setter]: isHighlightRecoverySupported */


// A Boolean that indicates whether to enable lens correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islenscorrectionenabled
func (r_ RAWFilter) IsLensCorrectionEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isLensCorrectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isLensCorrectionEnabled */


// A Boolean that indicates whether to enable lens correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islenscorrectionenabled
func (r_ RAWFilter) SetIsLensCorrectionEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsLensCorrectionEnabled:"), value)
}/* debug [instance_properties/setter]: isLensCorrectionEnabled */


// A Boolean that indicates if you can enable lens correction for the current image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islenscorrectionsupported
func (r_ RAWFilter) IsLensCorrectionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isLensCorrectionSupported"))
	return rv
}/* debug [instance_properties/getter]: isLensCorrectionSupported */


// A Boolean that indicates if you can enable lens correction for the current image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islenscorrectionsupported
func (r_ RAWFilter) SetIsLensCorrectionSupported(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsLensCorrectionSupported:"), value)
}/* debug [instance_properties/setter]: isLensCorrectionSupported */


// A Boolean that indicates if the current image supports local tone curve adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islocaltonemapsupported
func (r_ RAWFilter) IsLocalToneMapSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isLocalToneMapSupported"))
	return rv
}/* debug [instance_properties/getter]: isLocalToneMapSupported */


// A Boolean that indicates if the current image supports local tone curve adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/islocaltonemapsupported
func (r_ RAWFilter) SetIsLocalToneMapSupported(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsLocalToneMapSupported:"), value)
}/* debug [instance_properties/setter]: isLocalToneMapSupported */


// A Boolean that indicates if the current image supports luminance noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isluminancenoisereductionsupported
func (r_ RAWFilter) IsLuminanceNoiseReductionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isLuminanceNoiseReductionSupported"))
	return rv
}/* debug [instance_properties/getter]: isLuminanceNoiseReductionSupported */


// A Boolean that indicates if the current image supports luminance noise reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/isluminancenoisereductionsupported
func (r_ RAWFilter) SetIsLuminanceNoiseReductionSupported(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsLuminanceNoiseReductionSupported:"), value)
}/* debug [instance_properties/setter]: isLuminanceNoiseReductionSupported */


// A Boolean that indicates if the current image supports moire artifact reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ismoirereductionsupported
func (r_ RAWFilter) IsMoireReductionSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isMoireReductionSupported"))
	return rv
}/* debug [instance_properties/getter]: isMoireReductionSupported */


// A Boolean that indicates if the current image supports moire artifact reduction adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/ismoirereductionsupported
func (r_ RAWFilter) SetIsMoireReductionSupported(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsMoireReductionSupported:"), value)
}/* debug [instance_properties/setter]: isMoireReductionSupported */


// A Boolean that indicates if the current image supports sharpness adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/issharpnesssupported
func (r_ RAWFilter) IsSharpnessSupported() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isSharpnessSupported"))
	return rv
}/* debug [instance_properties/getter]: isSharpnessSupported */


// A Boolean that indicates if the current image supports sharpness adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirawfilter/issharpnesssupported
func (r_ RAWFilter) SetIsSharpnessSupported(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsSharpnessSupported:"), value)
}/* debug [instance_properties/setter]: isSharpnessSupported */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIRAWFilter */


