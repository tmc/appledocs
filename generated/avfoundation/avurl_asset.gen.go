// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVURLAsset */


/* debug [class_header]: Header for AVURLAsset */
// The class instance for the [URLAsset] class.
var (
	URLAssetClass     _URLAssetClass
	URLAssetClassOnce sync.Once
)

func getURLAssetClass() _URLAssetClass {
	URLAssetClassOnce.Do(func() {
		URLAssetClass = _URLAssetClass{objc.GetClass("AVURLAsset")}
	})
	return URLAssetClass
}

type _URLAssetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLAsset */
// An interface definition for the [URLAsset] class.
type IURLAsset interface {
	IAsset
	
/* debug [class_interface_properties]: Properties for URLAsset */
	// properties:
	AssetCache() IAVAssetCache
	HttpSessionIdentifier() foundation.UUID
	MayRequireContentKeysForMediaDataProcessing() bool
	MediaExtensionProperties() IAVMediaExtensionProperties
	ResourceLoader() IAVAssetResourceLoader
	SidecarURL() objc.IObject /* cross-framework: NSURL */
	URL() objc.IObject /* cross-framework: NSURL */
	Variants() []AssetVariant
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLAsset */
	// methods:
	FindCompatibleTrackForCompositionTrackCompletionHandler(compositionTrack IAVCompositionTrack, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLAsset */
// Alloc allocates a new instance without initialization.
func (uc _URLAssetClass) Alloc() URLAsset {
	rv := objc.Send[URLAsset](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _URLAssetClass) New() URLAsset {
	rv := objc.Send[URLAsset](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLAsset) Init() URLAsset {
	rv := objc.Send[URLAsset](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLAsset) Autorelease() URLAsset {
	rv := objc.Send[URLAsset](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLAsset creates a new URLAsset instance.
func NewURLAsset() URLAsset {
	return getURLAssetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLAsset */
// An asset that represents media at a local or remote URL.
//
// This class is a concrete subclass of . When you create an asset as shown below, the system creates and returns an instance of . In many cases, this is an appropriate way to create asset instances, but you can also directly instantiate an when you need more fine-grained control over its initialization. The initializer for accepts an options dictionary, which you use to customize the asset’s initialization for your particular purpose. For example, if you’re creating an asset for an HLS stream, you may want to prevent it from retrieving its media when it connects over a cellular network. You can do this by providing the initialization option and value as shown below.


// An asset that represents media at a local or remote URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset
type URLAsset struct {
	Asset
}

// URLAssetFrom constructs a [URLAsset] from an unsafe.Pointer.
//
// An asset that represents media at a local or remote URL.
func URLAssetFrom(ptr unsafe.Pointer) URLAsset {
	return URLAsset{
		Asset: AssetFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLAsset */

// Creates an asset that models the media resource at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/init(url:options:)
func NewURLAssetWithURLOptions(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary) URLAsset {
	instance := getURLAssetClass().Alloc()
	rv := objc.Send[URLAsset](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLAssetWithURLOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLAsset */

// Returns an asset that models the media resource found at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/URLAssetWithURL:options:
func (uc _URLAssetClass) URLAssetWithURLOptions(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("URLAssetWithURL:options:"), URL, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=URLAssetWithURLOptions) */


// Returns an array of the MIME types the asset supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/audiovisualMIMETypes()
func (uc _URLAssetClass) AudiovisualMIMETypes() []string {
	rv := objc.Send[[]string](objc.ID(uc.class), objc.Sel("audiovisualMIMETypes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AudiovisualMIMETypes) */


// Returns an array of the file types the asset supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/audiovisualTypes()
func (uc _URLAssetClass) AudiovisualTypes() []string {
	rv := objc.Send[[]string](objc.ID(uc.class), objc.Sel("audiovisualTypes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AudiovisualTypes) */


// Returns a Boolean value that indicates whether the asset is playable with the specified codecs and container type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/isPlayableExtendedMIMEType(_:)
func (uc _URLAssetClass) IsPlayableExtendedMIMEType(extendedMIMEType objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("isPlayableExtendedMIMEType:"), extendedMIMEType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsPlayableExtendedMIMEType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLAsset */

// Provides the content types the AVURLAsset class understands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/audiovisualContentTypes
func (uc _URLAssetClass) AudiovisualContentTypes() []uniformtypeidentifiers.UTType {
	rv := objc.Send[[]uniformtypeidentifiers.UTType](objc.ID(uc.class), objc.Sel("audiovisualContentTypes"))
	return rv
}/* debug [class_properties_class/property]: audiovisualContentTypes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLAsset */

// Loads an asset track from which you can insert any time range into the composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/findCompatibleTrack(for:completionHandler:)
func (u_ URLAsset) FindCompatibleTrackForCompositionTrackCompletionHandler(compositionTrack IAVCompositionTrack, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("findCompatibleTrackForCompositionTrack:completionHandler:"), compositionTrack, completionHandler)
}/* debug [instance_methods/method]: FindCompatibleTrackForCompositionTrackCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLAsset */

// The asset’s associated asset cache, if it exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/assetCache
func (u_ URLAsset) AssetCache() IAVAssetCache {
	rv := objc.Send[AssetCache](u_.ID, objc.Sel("assetCache"))
	return rv
}/* debug [instance_properties/getter]: assetCache */


// Provides the content types the AVURLAsset class understands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/audiovisualContentTypes
func (u_ URLAsset) AudiovisualContentTypes() []uniformtypeidentifiers.UTType {
	rv := objc.Send[[]uniformtypeidentifiers.UTType](u_.ID, objc.Sel("audiovisualContentTypes"))
	return rv
}/* debug [instance_properties/getter]: audiovisualContentTypes */


// A session identifier that the asset sends in HTTP requests that it makes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/httpSessionIdentifier
func (u_ URLAsset) HttpSessionIdentifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](u_.ID, objc.Sel("httpSessionIdentifier"))
	return rv
}/* debug [instance_properties/getter]: httpSessionIdentifier */


// A Boolean value that indicates whether you can add this asset as a content key recipient to a content key session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/mayRequireContentKeysForMediaDataProcessing
func (u_ URLAsset) MayRequireContentKeysForMediaDataProcessing() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("mayRequireContentKeysForMediaDataProcessing"))
	return rv
}/* debug [instance_properties/getter]: mayRequireContentKeysForMediaDataProcessing */


// The properties of the media extension format reader that decodes the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/mediaExtensionProperties
func (u_ URLAsset) MediaExtensionProperties() IAVMediaExtensionProperties {
	rv := objc.Send[MediaExtensionProperties](u_.ID, objc.Sel("mediaExtensionProperties"))
	return rv
}/* debug [instance_properties/getter]: mediaExtensionProperties */


// The resource loader for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/resourceLoader
func (u_ URLAsset) ResourceLoader() IAVAssetResourceLoader {
	rv := objc.Send[AssetResourceLoader](u_.ID, objc.Sel("resourceLoader"))
	return rv
}/* debug [instance_properties/getter]: resourceLoader */


// The sidecar URL used by the MediaExtension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/sidecarURL
func (u_ URLAsset) SidecarURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](u_.ID, objc.Sel("sidecarURL"))
	return rv
}/* debug [instance_properties/getter]: sidecarURL */


// A URL to the asset’s media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/url
func (u_ URLAsset) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](u_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// An array of variants that an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/variants
func (u_ URLAsset) Variants() []AssetVariant {
	rv := objc.Send[[]AssetVariant](u_.ID, objc.Sel("variants"))
	return rv
}/* debug [instance_properties/getter]: variants */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVURLAsset */


