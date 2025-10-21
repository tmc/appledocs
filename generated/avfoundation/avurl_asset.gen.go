// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// An asset that represents media at a local or remote URL.
//
// This class is a concrete subclass of . When you create an asset as shown below, the system creates and returns an instance of . In many cases, this is an appropriate way to create asset instances, but you can also directly instantiate an when you need more fine-grained control over its initialization. The initializer for accepts an options dictionary, which you use to customize the asset’s initialization for your particular purpose. For example, if you’re creating an asset for an HLS stream, you may want to prevent it from retrieving its media when it connects over a cellular network. You can do this by providing the initialization option and value as shown below.
//
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


// The properties of the media extension format reader that decodes the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/mediaExtensionProperties
func (u_ URLAsset) MediaExtensionProperties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("mediaExtensionProperties"))
	return rv
}

// The asset’s associated asset cache, if it exists.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/assetcache
func (u_ URLAsset) AssetCache() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("assetCache"))
	return rv
}


// SetAssetCache sets the value of the assetCache property.
// The asset’s associated asset cache, if it exists.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/assetcache
func (u_ URLAsset) SetAssetCache(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAssetCache:"), value)
}

// A session identifier that the asset sends in HTTP requests that it makes.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/httpsessionidentifier
func (u_ URLAsset) HttpSessionIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("httpSessionIdentifier"))
	return rv
}


// SetHttpSessionIdentifier sets the value of the httpSessionIdentifier property.
// A session identifier that the asset sends in HTTP requests that it makes.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/httpsessionidentifier
func (u_ URLAsset) SetHttpSessionIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpSessionIdentifier:"), value)
}

// A Boolean value that indicates whether you can add this asset as a content key recipient to a content key session.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/mayrequirecontentkeysformediadataprocessing
func (u_ URLAsset) MayRequireContentKeysForMediaDataProcessing() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("mayRequireContentKeysForMediaDataProcessing"))
	return rv
}


// SetMayRequireContentKeysForMediaDataProcessing sets the value of the mayRequireContentKeysForMediaDataProcessing property.
// A Boolean value that indicates whether you can add this asset as a content key recipient to a content key session.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/mayrequirecontentkeysformediadataprocessing
func (u_ URLAsset) SetMayRequireContentKeysForMediaDataProcessing(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMayRequireContentKeysForMediaDataProcessing:"), value)
}

// The resource loader for the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/resourceloader
func (u_ URLAsset) ResourceLoader() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("resourceLoader"))
	return rv
}


// SetResourceLoader sets the value of the resourceLoader property.
// The resource loader for the asset.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/resourceloader
func (u_ URLAsset) SetResourceLoader(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResourceLoader:"), value)
}

// A URL to the asset’s media.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/url
func (u_ URLAsset) Url() foundation.URL {
	rv := objc.Send[foundation.URL](u_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// A URL to the asset’s media.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/url
func (u_ URLAsset) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUrl:"), value)
}

// An array of variants that an asset contains.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/variants
func (u_ URLAsset) Variants() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("variants"))
	return rv
}


// SetVariants sets the value of the variants property.
// An array of variants that an asset contains.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/variants
func (u_ URLAsset) SetVariants(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setVariants:"), value)
}



