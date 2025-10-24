// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetVariantQualifier */


/* debug [class_header]: Header for AVAssetVariantQualifier */
// The class instance for the [AssetVariantQualifier] class.
var (
	AssetVariantQualifierClass     _AssetVariantQualifierClass
	AssetVariantQualifierClassOnce sync.Once
)

func getAssetVariantQualifierClass() _AssetVariantQualifierClass {
	AssetVariantQualifierClassOnce.Do(func() {
		AssetVariantQualifierClass = _AssetVariantQualifierClass{objc.GetClass("AVAssetVariantQualifier")}
	})
	return AssetVariantQualifierClass
}

type _AssetVariantQualifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetVariantQualifier */
// An interface definition for the [AssetVariantQualifier] class.
type IAssetVariantQualifier interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetVariantQualifier */
	// properties:
	MediaSelections() IAVMediaSelection
	SetMediaSelections(value IAVMediaSelection)
	VariantQualifiers() IAVAssetVariantQualifier
	SetVariantQualifiers(value IAVAssetVariantQualifier)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetVariantQualifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetVariantQualifier */
// Alloc allocates a new instance without initialization.
func (ac _AssetVariantQualifierClass) Alloc() AssetVariantQualifier {
	rv := objc.Send[AssetVariantQualifier](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetVariantQualifierClass) New() AssetVariantQualifier {
	rv := objc.Send[AssetVariantQualifier](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetVariantQualifier) Init() AssetVariantQualifier {
	rv := objc.Send[AssetVariantQualifier](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetVariantQualifier) Autorelease() AssetVariantQualifier {
	rv := objc.Send[AssetVariantQualifier](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetVariantQualifier creates a new AssetVariantQualifier instance.
func NewAssetVariantQualifier() AssetVariantQualifier {
	return getAssetVariantQualifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetVariantQualifier */
// An object that represents an HTTP Live Streaming asset variant.


// An object that represents an HTTP Live Streaming asset variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier
type AssetVariantQualifier struct {
	objectivec.Object
}

// AssetVariantQualifierFrom constructs a [AssetVariantQualifier] from an unsafe.Pointer.
//
// An object that represents an HTTP Live Streaming asset variant.
func AssetVariantQualifierFrom(ptr unsafe.Pointer) AssetVariantQualifier {
	return AssetVariantQualifier{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetVariantQualifier */

// Creates a variant qualifier with a predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/init(predicate:)
func NewAssetVariantQualifierWithPredicate(predicate foundation.Predicate) AssetVariantQualifier {
	rv := objc.Send[AssetVariantQualifier](objc.ID(getAssetVariantQualifierClass().class), objc.Sel("assetVariantQualifierWithPredicate:"), predicate)
	return rv
}/* debug [class_init_methods/constructor]: NewAssetVariantQualifierWithPredicate */


// Creates a variant qualifier with an asset variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/init(variant:)
func NewAssetVariantQualifierWithVariant(variant IAVAssetVariant) AssetVariantQualifier {
	rv := objc.Send[AssetVariantQualifier](objc.ID(getAssetVariantQualifierClass().class), objc.Sel("assetVariantQualifierWithVariant:"), variant)
	return rv
}/* debug [class_init_methods/constructor]: NewAssetVariantQualifierWithVariant */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetVariantQualifier */

// Returns a qualifer for finding variant with maximum value in the input key path
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/assetVariantQualifierForMaximumValueInKeyPath:
func (ac _AssetVariantQualifierClass) AssetVariantQualifierForMaximumValueInKeyPath(keyPath objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetVariantQualifierForMaximumValueInKeyPath:"), keyPath)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetVariantQualifierForMaximumValueInKeyPath) */


// Returns a qualifer for finding variant with minimum value in the input key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/assetVariantQualifierForMinimumValueInKeyPath:
func (ac _AssetVariantQualifierClass) AssetVariantQualifierForMinimumValueInKeyPath(keyPath objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetVariantQualifierForMinimumValueInKeyPath:"), keyPath)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetVariantQualifierForMinimumValueInKeyPath) */


// Creates a variant qualifier with a predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/init(predicate:)
func (ac _AssetVariantQualifierClass) AssetVariantQualifierWithPredicate(predicate foundation.Predicate) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetVariantQualifierWithPredicate:"), predicate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetVariantQualifierWithPredicate) */


// Creates a variant qualifier with an asset variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/init(variant:)
func (ac _AssetVariantQualifierClass) AssetVariantQualifierWithVariant(variant IAVAssetVariant) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetVariantQualifierWithVariant:"), variant)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetVariantQualifierWithVariant) */


// Creates a predicate for audio sample rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forAudioSampleRate:mediaSelectionOption:operatorType:)
func (ac _AssetVariantQualifierClass) PredicateForAudioSampleRateMediaSelectionOptionOperatorType(sampleRate float64, mediaSelectionOption IAVMediaSelectionOption, operatorType PredicateOperatorType /* not a class type */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForAudioSampleRate:mediaSelectionOption:operatorType:"), sampleRate, mediaSelectionOption, operatorType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForAudioSampleRateMediaSelectionOptionOperatorType) */


// Creates a NSPredicate for audio sample rate which can be used with other NSPredicates to express variant preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forAudioSampleRate:operatorType:)
func (ac _AssetVariantQualifierClass) PredicateForAudioSampleRateOperatorType(sampleRate float64, operatorType PredicateOperatorType /* not a class type */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForAudioSampleRate:operatorType:"), sampleRate, operatorType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForAudioSampleRateOperatorType) */


// Creates a NSPredicate for binaural which can be used with other NSPredicates to express variant preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forBinauralAudio:)
func (ac _AssetVariantQualifierClass) PredicateForBinauralAudio(isBinauralAudio bool) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForBinauralAudio:"), isBinauralAudio)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForBinauralAudio) */


// Creates a predicate for binaural audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forBinauralAudio:mediaSelectionOption:)
func (ac _AssetVariantQualifierClass) PredicateForBinauralAudioMediaSelectionOption(isBinauralAudio bool, mediaSelectionOption IAVMediaSelectionOption) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForBinauralAudio:mediaSelectionOption:"), isBinauralAudio, mediaSelectionOption)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForBinauralAudioMediaSelectionOption) */


// Creates a predicate with a channel count, media selection option, and operator type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forChannelCount:mediaSelectionOption:operatorType:)
func (ac _AssetVariantQualifierClass) PredicateForChannelCountMediaSelectionOptionOperatorType(channelCount int, mediaSelectionOption IAVMediaSelectionOption, operatorType PredicateOperatorType /* not a class type */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForChannelCount:mediaSelectionOption:operatorType:"), channelCount, mediaSelectionOption, operatorType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForChannelCountMediaSelectionOptionOperatorType) */


// Creates a NSPredicate for audio channel count which can be used with other NSPredicates to express variant preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forChannelCount:operatorType:)
func (ac _AssetVariantQualifierClass) PredicateForChannelCountOperatorType(channelCount int, operatorType PredicateOperatorType /* not a class type */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForChannelCount:operatorType:"), channelCount, operatorType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForChannelCountOperatorType) */


// Creates a NSPredicate for immersive audio which can be used with other NSPredicates to express variant preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forDownmixAudio:)
func (ac _AssetVariantQualifierClass) PredicateForDownmixAudio(isDownmixAudio bool) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForDownmixAudio:"), isDownmixAudio)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForDownmixAudio) */


// Creates a predicate for downmix audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forDownmixAudio:mediaSelectionOption:)
func (ac _AssetVariantQualifierClass) PredicateForDownmixAudioMediaSelectionOption(isDownmixAudio bool, mediaSelectionOption IAVMediaSelectionOption) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForDownmixAudio:mediaSelectionOption:"), isDownmixAudio, mediaSelectionOption)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForDownmixAudioMediaSelectionOption) */


// Creates a NSPredicate for immersive audio which can be used with other NSPredicates to express variant preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forImmersiveAudio:)
func (ac _AssetVariantQualifierClass) PredicateForImmersiveAudio(isImmersiveAudio bool) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForImmersiveAudio:"), isImmersiveAudio)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForImmersiveAudio) */


// Creates a predicate for immersive audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forImmersiveAudio:mediaSelectionOption:)
func (ac _AssetVariantQualifierClass) PredicateForImmersiveAudioMediaSelectionOption(isImmersiveAudio bool, mediaSelectionOption IAVMediaSelectionOption) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForImmersiveAudio:mediaSelectionOption:"), isImmersiveAudio, mediaSelectionOption)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForImmersiveAudioMediaSelectionOption) */


// Creates a predicate with a height and operator type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forPresentationHeight:operatorType:)
func (ac _AssetVariantQualifierClass) PredicateForPresentationHeightOperatorType(height float64, operatorType PredicateOperatorType /* not a class type */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForPresentationHeight:operatorType:"), height, operatorType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForPresentationHeightOperatorType) */


// Creates a predicate with a width and operator type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forPresentationWidth:operatorType:)
func (ac _AssetVariantQualifierClass) PredicateForPresentationWidthOperatorType(width float64, operatorType PredicateOperatorType /* not a class type */) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(ac.class), objc.Sel("predicateForPresentationWidth:operatorType:"), width, operatorType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateForPresentationWidthOperatorType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetVariantQualifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetVariantQualifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetVariantQualifier */

// The media selections of an asset that a task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadcontentconfiguration/mediaselections
func (a_ AssetVariantQualifier) MediaSelections() IAVMediaSelection {
	rv := objc.Send[MediaSelection](a_.ID, objc.Sel("mediaSelections"))
	return rv
}/* debug [instance_properties/getter]: mediaSelections */


// The media selections of an asset that a task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadcontentconfiguration/mediaselections
func (a_ AssetVariantQualifier) SetMediaSelections(value IAVMediaSelection) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaSelections:"), value)
}/* debug [instance_properties/setter]: mediaSelections */


// The variant qualifiers for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadcontentconfiguration/variantqualifiers
func (a_ AssetVariantQualifier) VariantQualifiers() IAVAssetVariantQualifier {
	rv := objc.Send[AssetVariantQualifier](a_.ID, objc.Sel("variantQualifiers"))
	return rv
}/* debug [instance_properties/getter]: variantQualifiers */


// The variant qualifiers for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadcontentconfiguration/variantqualifiers
func (a_ AssetVariantQualifier) SetVariantQualifiers(value IAVAssetVariantQualifier) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVariantQualifiers:"), value)
}/* debug [instance_properties/setter]: variantQualifiers */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetVariantQualifier */


