// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHAsset] class.
var (
	PHAssetClass     _PHAssetClass
	PHAssetClassOnce sync.Once
)

func getPHAssetClass() _PHAssetClass {
	PHAssetClassOnce.Do(func() {
		PHAssetClass = _PHAssetClass{objc.GetClass("PHAsset")}
	})
	return PHAssetClass
}

type _PHAssetClass struct {
	class objc.Class
}

// An interface definition for the [PHAsset] class.
type IPHAsset interface {
	IPHObject
	CanPerformEditOperation(editOperation unsafe.Pointer) bool
	CancelContentEditingInputRequest(requestID unsafe.Pointer)
	RequestContentEditingInputWithOptionsCompletionHandler(options unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer
}

// A representation of an image, video, or Live Photo in the Photos library.
//
// You fetch assets to begin working with them. Use the class methods listed in Fetching Assets to retrieve one or more instances representing the assets you want to display or edit. Assets contain only metadata. The underlying image or video data for any given asset might not be stored on the local device. However, depending on how you plan to use this data, you may not need to download all of it. If you need to populate a collection view with thumbnail images, the Photos framework can manage downloading, generating, and caching thumbnails for each asset. For details, see . Asset objects are immutable. To edit an asset’s metadata (such as marking it as a favorite photo), create a object within a photo library change block. For more details on using change requests and change blocks to update the photo library, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset
type PHAsset struct {
	PHObject
}

// PHAssetFrom constructs a [PHAsset] from an unsafe.Pointer.
//
// A representation of an image, video, or Live Photo in the Photos library.
func PHAssetFrom(ptr unsafe.Pointer) PHAsset {
	return PHAsset{
		PHObject: PHObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAssetClass) Alloc() PHAsset {
	rv := objc.Send[PHAsset](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAssetClass) New() PHAsset {
	rv := objc.Send[PHAsset](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAsset) Init() PHAsset {
	rv := objc.Send[PHAsset](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAsset) Autorelease() PHAsset {
	rv := objc.Send[PHAsset](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAsset creates a new PHAsset instance.
func NewPHAsset() PHAsset {
	return getPHAssetClass().New()
}


// Retrieves assets from the specified asset collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/fetchAssets(in:options:)
func (pc _PHAssetClass) FetchAssetsInAssetCollectionOptions(assetCollection unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchAssetsInAssetCollection:options:"), assetCollection, options)
	return rv
}

// Retrieves all assets matching the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/fetchAssets(with:)
func (pc _PHAssetClass) FetchAssetsWithOptions(options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchAssetsWithOptions:"), options)
	return rv
}

// Retrieves assets with the specified media type.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/fetchAssets(with:options:)
func (pc _PHAssetClass) FetchAssetsWithMediaTypeOptions(mediaType unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchAssetsWithMediaType:options:"), mediaType, options)
	return rv
}

// Retrieves assets using URLs provided by the Assets Library framework.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/fetchAssets(withALAssetURLs:options:)
func (pc _PHAssetClass) FetchAssetsWithALAssetURLsOptions(assetURLs unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchAssetsWithALAssetURLs:options:"), assetURLs, options)
	return rv
}

// Retrieves assets with the specified burst photo sequence identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/fetchAssets(withBurstIdentifier:options:)
func (pc _PHAssetClass) FetchAssetsWithBurstIdentifierOptions(burstIdentifier string, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchAssetsWithBurstIdentifier:options:"), objc.String(burstIdentifier), options)
	return rv
}

// Retrieves assets with the specified local-device-specific unique identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/fetchAssets(withLocalIdentifiers:options:)
func (pc _PHAssetClass) FetchAssetsWithLocalIdentifiersOptions(identifiers unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchAssetsWithLocalIdentifiers:options:"), identifiers, options)
	return rv
}

// Retrieves assets marked as key assets in the specified asset collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/fetchKeyAssets(in:options:)
func (pc _PHAssetClass) FetchKeyAssetsInAssetCollectionOptions(assetCollection unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchKeyAssetsInAssetCollection:options:"), assetCollection, options)
	return rv
}

// Returns whether the asset supports the specified editing operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/canPerform(_:)
func (p_ PHAsset) CanPerformEditOperation(editOperation unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPerformEditOperation:"), editOperation)
	return rv
}

// Cancels a request for editing the asset’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/cancelContentEditingInputRequest(_:)
func (p_ PHAsset) CancelContentEditingInputRequest(requestID unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancelContentEditingInputRequest:"), requestID)
}

// Requests asset information for beginning a content editing session.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/requestContentEditingInput(with:completionHandler:)
func (p_ PHAsset) RequestContentEditingInputWithOptionsCompletionHandler(options unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("requestContentEditingInputWithOptions:completionHandler:"), options, completionHandler)
	return rv
}

// The date and time this asset was added to the photo library (from the device that was used to add this asset)
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/addedDate
func (p_ PHAsset) AddedDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("addedDate"))
	return rv
}

// The identifier that describes the adjustment format.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/adjustmentFormatIdentifier
func (p_ PHAsset) AdjustmentFormatIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("adjustmentFormatIdentifier"))
	return rv
}

// The unique identifier shared by photo assets from the same burst sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/burstIdentifier
func (p_ PHAsset) BurstIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("burstIdentifier"))
	return rv
}

// The selection type of the asset in a burst photo sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/burstSelectionTypes
func (p_ PHAsset) BurstSelectionTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("burstSelectionTypes"))
	return rv
}

// The type of image or video data that is presented for the asset
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/contentType
func (p_ PHAsset) ContentType() UTType {
	rv := objc.Send[UTType](p_.ID, objc.Sel("contentType"))
	return rv
}

// The date and time of the asset’s creation.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/creationDate
func (p_ PHAsset) CreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("creationDate"))
	return rv
}

// The duration, in seconds, of the video asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/duration
func (p_ PHAsset) Duration() TimeInterval {
	rv := objc.Send[TimeInterval](p_.ID, objc.Sel("duration"))
	return rv
}

// A Boolean value that indicates whether the asset contains adjustment data.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/hasAdjustments
func (p_ PHAsset) HasAdjustments() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasAdjustments"))
	return rv
}

// A Boolean value that indicates whether the user marks the asset as a favorite.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/isFavorite
func (p_ PHAsset) Favorite() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("favorite"))
	return rv
}

// A Boolean value that indicates whether the user hides the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/isHidden
func (p_ PHAsset) Hidden() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hidden"))
	return rv
}

// A Boolean value that indicates whether the user hides the sync failure message.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/isSyncFailureHidden
func (p_ PHAsset) SyncFailureHidden() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("syncFailureHidden"))
	return rv
}

// The location information for the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/location
func (p_ PHAsset) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("location"))
	return rv
}

// The subtypes of the asset, identifying special kinds of assets, such as panoramic photo or high-frame-rate video.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/mediaSubtypes
func (p_ PHAsset) MediaSubtypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mediaSubtypes"))
	return rv
}

// The type of the asset, such as video or audio.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/mediaType
func (p_ PHAsset) MediaType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mediaType"))
	return rv
}

// The date and time of the asset’s last modification.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/modificationDate
func (p_ PHAsset) ModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("modificationDate"))
	return rv
}

// The height, in pixels, of the asset’s image or video data.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/pixelHeight
func (p_ PHAsset) PixelHeight() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("pixelHeight"))
	return rv
}

// The width, in pixels, of the asset’s image or video data.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/pixelWidth
func (p_ PHAsset) PixelWidth() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("pixelWidth"))
	return rv
}

// An enumerated value that describes how to present an asset to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/playbackStyle-swift.property
func (p_ PHAsset) PlaybackStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("playbackStyle"))
	return rv
}

// A Boolean value that indicates whether the asset is the representative photo from a burst photo sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/representsBurst
func (p_ PHAsset) RepresentsBurst() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("representsBurst"))
	return rv
}

// The means by which the asset enters the user’s Photos library.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/sourceType
func (p_ PHAsset) SourceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("sourceType"))
	return rv
}



