// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNAssetSpatialAudioInfo */


/* debug [class_header]: Header for CNAssetSpatialAudioInfo */
// The class instance for the [CNAssetSpatialAudioInfo] class.
var (
	CNAssetSpatialAudioInfoClass     _CNAssetSpatialAudioInfoClass
	CNAssetSpatialAudioInfoClassOnce sync.Once
)

func getCNAssetSpatialAudioInfoClass() _CNAssetSpatialAudioInfoClass {
	CNAssetSpatialAudioInfoClassOnce.Do(func() {
		CNAssetSpatialAudioInfoClass = _CNAssetSpatialAudioInfoClass{objc.GetClass("CNAssetSpatialAudioInfo")}
	})
	return CNAssetSpatialAudioInfoClass
}

type _CNAssetSpatialAudioInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNAssetSpatialAudioInfo */
// An interface definition for the [CNAssetSpatialAudioInfo] class.
type ICNAssetSpatialAudioInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNAssetSpatialAudioInfo */
	// properties:
	DefaultEffectIntensity() float32
	DefaultRenderingStyle() CNSpatialAudioRenderingStyle
	DefaultSpatialAudioTrack() avfoundation.AssetTrack
	SpatialAudioMixMetadata() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNAssetSpatialAudioInfo */
	// methods:
	AssetReaderOutputSettingsForContentType(contentType CNSpatialAudioContentType) foundation.IDictionary
	AssetWriterInputSettingsForContentType(contentType CNSpatialAudioContentType) foundation.IDictionary
	AudioMixWithEffectIntensityRenderingStyle(effectIntensity float32, renderingStyle CNSpatialAudioRenderingStyle) avfoundation.AudioMix
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNAssetSpatialAudioInfo */
// Alloc allocates a new instance without initialization.
func (cc _CNAssetSpatialAudioInfoClass) Alloc() CNAssetSpatialAudioInfo {
	rv := objc.Send[CNAssetSpatialAudioInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNAssetSpatialAudioInfoClass) New() CNAssetSpatialAudioInfo {
	rv := objc.Send[CNAssetSpatialAudioInfo](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNAssetSpatialAudioInfo) Init() CNAssetSpatialAudioInfo {
	rv := objc.Send[CNAssetSpatialAudioInfo](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNAssetSpatialAudioInfo) Autorelease() CNAssetSpatialAudioInfo {
	rv := objc.Send[CNAssetSpatialAudioInfo](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNAssetSpatialAudioInfo creates a new CNAssetSpatialAudioInfo instance.
func NewCNAssetSpatialAudioInfo() CNAssetSpatialAudioInfo {
	return getCNAssetSpatialAudioInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNAssetSpatialAudioInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5
type CNAssetSpatialAudioInfo struct {
	objectivec.Object
}

// CNAssetSpatialAudioInfoFrom constructs a [CNAssetSpatialAudioInfo] from an unsafe.Pointer.
func CNAssetSpatialAudioInfoFrom(ptr unsafe.Pointer) CNAssetSpatialAudioInfo {
	return CNAssetSpatialAudioInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNAssetSpatialAudioInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNAssetSpatialAudioInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/checkIfContainsSpatialAudio:completionHandler:
func (cc _CNAssetSpatialAudioInfoClass) CheckIfContainsSpatialAudioCompletionHandler(asset avfoundation.Asset, completionHandler bool) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("checkIfContainsSpatialAudio:completionHandler:"), asset, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CheckIfContainsSpatialAudioCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/loadFromAsset:completionHandler:
func (cc _CNAssetSpatialAudioInfoClass) LoadFromAssetCompletionHandler(asset avfoundation.Asset, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadFromAsset:completionHandler:"), asset, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadFromAssetCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNAssetSpatialAudioInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/isSupported
func (cc _CNAssetSpatialAudioInfoClass) IsSupported() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("isSupported"))
	return rv
}/* debug [class_properties_class/property]: isSupported */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNAssetSpatialAudioInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/assetReaderOutputSettingsForContentType:
func (c_ CNAssetSpatialAudioInfo) AssetReaderOutputSettingsForContentType(contentType CNSpatialAudioContentType) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("assetReaderOutputSettingsForContentType:"), contentType)
	return rv
}/* debug [instance_methods/method]: AssetReaderOutputSettingsForContentType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/assetWriterInputSettingsForContentType:
func (c_ CNAssetSpatialAudioInfo) AssetWriterInputSettingsForContentType(contentType CNSpatialAudioContentType) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("assetWriterInputSettingsForContentType:"), contentType)
	return rv
}/* debug [instance_methods/method]: AssetWriterInputSettingsForContentType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/audioMixWithEffectIntensity:renderingStyle:
func (c_ CNAssetSpatialAudioInfo) AudioMixWithEffectIntensityRenderingStyle(effectIntensity float32, renderingStyle CNSpatialAudioRenderingStyle) avfoundation.AudioMix {
	rv := objc.Send[avfoundation.AudioMix](c_.ID, objc.Sel("audioMixWithEffectIntensity:renderingStyle:"), effectIntensity, renderingStyle)
	return rv
}/* debug [instance_methods/method]: AudioMixWithEffectIntensityRenderingStyle */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNAssetSpatialAudioInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/defaultEffectIntensity
func (c_ CNAssetSpatialAudioInfo) DefaultEffectIntensity() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("defaultEffectIntensity"))
	return rv
}/* debug [instance_properties/getter]: defaultEffectIntensity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/defaultRenderingStyle
func (c_ CNAssetSpatialAudioInfo) DefaultRenderingStyle() CNSpatialAudioRenderingStyle {
	rv := objc.Send[CNSpatialAudioRenderingStyle](c_.ID, objc.Sel("defaultRenderingStyle"))
	return rv
}/* debug [instance_properties/getter]: defaultRenderingStyle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/defaultSpatialAudioTrack
func (c_ CNAssetSpatialAudioInfo) DefaultSpatialAudioTrack() avfoundation.AssetTrack {
	rv := objc.Send[avfoundation.AssetTrack](c_.ID, objc.Sel("defaultSpatialAudioTrack"))
	return rv
}/* debug [instance_properties/getter]: defaultSpatialAudioTrack */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/isSupported
func (c_ CNAssetSpatialAudioInfo) IsSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSupported"))
	return rv
}/* debug [instance_properties/getter]: isSupported */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/spatialAudioMixMetadata
func (c_ CNAssetSpatialAudioInfo) SpatialAudioMixMetadata() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("spatialAudioMixMetadata"))
	return rv
}/* debug [instance_properties/getter]: spatialAudioMixMetadata */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNAssetSpatialAudioInfo */



