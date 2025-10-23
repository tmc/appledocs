// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [URLAsset] class.
type IURLAsset interface {
	IAsset
	// properties:
	AssetCache() AssetCache /* not a class type */
	HttpSessionIdentifier() UUID /* not a class type */
	MayRequireContentKeysForMediaDataProcessing() bool /* primitive/slice/pointer */
	MediaExtensionProperties() IAVMediaExtensionProperties
	ResourceLoader() IAVAssetResourceLoader
	SidecarURL() foundation.URL /* not a class type */
	URL() foundation.URL /* not a class type */
	Variants() []AssetVariant /* primitive/slice/pointer */
	// methods:
	FindCompatibleTrackForCompositionTrackCompletionHandler(compositionTrack CompositionTrack /* not a class type */, completionHandler unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (uc _URLAssetClass) Alloc() URLAsset {
	rv := objc.Send[URLAsset](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates an asset that models the media resource at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/init(url:options:)
func NewURLAssetWithURLOptions(URL foundation.URL /* not a class type */, options foundation.IDictionary /* already interface */) URLAsset {
	instance := getURLAssetClass().Alloc()
	rv := objc.Send[URLAsset](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	rv.Autorelease()
	return rv
}



// Returns an asset that models the media resource found at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/URLAssetWithURL:options:
func (uc _URLAssetClass) URLAssetWithURLOptions(URL foundation.URL /* not a class type */, options foundation.IDictionary /* already interface */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("URLAssetWithURL:options:"), URL, options)
	return rv
}


// Returns an array of the MIME types the asset supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/audiovisualMIMETypes()
func (uc _URLAssetClass) AudiovisualMIMETypes() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](objc.ID(uc.class), objc.Sel("audiovisualMIMETypes"))
	return rv
}


// Returns an array of the file types the asset supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/audiovisualTypes()
func (uc _URLAssetClass) AudiovisualTypes() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](objc.ID(uc.class), objc.Sel("audiovisualTypes"))
	return rv
}


// Returns a Boolean value that indicates whether the asset is playable with the specified codecs and container type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/isPlayableExtendedMIMEType(_:)
func (uc _URLAssetClass) IsPlayableExtendedMIMEType(extendedMIMEType string /* primitive/slice/pointer */) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("isPlayableExtendedMIMEType:"), objc.String(extendedMIMEType))
	return rv
}


// Provides the content types the AVURLAsset class understands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/audiovisualContentTypes
func (uc _URLAssetClass) AudiovisualContentTypes() []objectivec.IObject /* already interface */ {
	rv := objc.Send[[]objectivec.IObject](objc.ID(uc.class), objc.Sel("audiovisualContentTypes"))
	return rv
}

// Loads an asset track from which you can insert any time range into the composition track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/findCompatibleTrack(for:completionHandler:)
func (u_ URLAsset) FindCompatibleTrackForCompositionTrackCompletionHandler(compositionTrack CompositionTrack /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("findCompatibleTrackForCompositionTrack:completionHandler:"), compositionTrack, completionHandler)
}


// The asset’s associated asset cache, if it exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/assetCache
func (u_ URLAsset) AssetCache() AssetCache /* not a class type */ {
	rv := objc.Send[AssetCache](u_.ID, objc.Sel("assetCache"))
	return rv
}


// Provides the content types the AVURLAsset class understands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/audiovisualContentTypes
func (u_ URLAsset) AudiovisualContentTypes() []objectivec.IObject /* already interface */ {
	rv := objc.Send[[]objectivec.IObject](u_.ID, objc.Sel("audiovisualContentTypes"))
	return rv
}


// A session identifier that the asset sends in HTTP requests that it makes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/httpSessionIdentifier
func (u_ URLAsset) HttpSessionIdentifier() UUID /* not a class type */ {
	rv := objc.Send[UUID](u_.ID, objc.Sel("httpSessionIdentifier"))
	return rv
}


// A Boolean value that indicates whether you can add this asset as a content key recipient to a content key session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/mayRequireContentKeysForMediaDataProcessing
func (u_ URLAsset) MayRequireContentKeysForMediaDataProcessing() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("mayRequireContentKeysForMediaDataProcessing"))
	return rv
}


// The properties of the media extension format reader that decodes the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/mediaExtensionProperties
func (u_ URLAsset) MediaExtensionProperties() IAVMediaExtensionProperties {
	rv := objc.Send[MediaExtensionProperties](u_.ID, objc.Sel("mediaExtensionProperties"))
	return rv
}


// The resource loader for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/resourceLoader
func (u_ URLAsset) ResourceLoader() IAVAssetResourceLoader {
	rv := objc.Send[AssetResourceLoader](u_.ID, objc.Sel("resourceLoader"))
	return rv
}


// The sidecar URL used by the MediaExtension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/sidecarURL
func (u_ URLAsset) SidecarURL() foundation.URL /* not a class type */ {
	rv := objc.Send[foundation.URL](u_.ID, objc.Sel("sidecarURL"))
	return rv
}


// A URL to the asset’s media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/url
func (u_ URLAsset) URL() foundation.URL /* not a class type */ {
	rv := objc.Send[foundation.URL](u_.ID, objc.Sel("URL"))
	return rv
}


// An array of variants that an asset contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/variants
func (u_ URLAsset) Variants() []AssetVariant /* primitive/slice/pointer */ {
	rv := objc.Send[[]AssetVariant](u_.ID, objc.Sel("variants"))
	return rv
}


