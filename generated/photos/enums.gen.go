// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

// Enum types and constants
// PHAccessLevel - The app’s level of access to the user’s photo library.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAccessLevel
type PHAccessLevel uint

// PHAssetPlaybackStyle - An enumeration of asset playback styles that dictate how to present an asset to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/PlaybackStyle-swift.enum
type PHAssetPlaybackStyle uint

// PHAssetBurstSelectionType - Bit mask values indicating whether and how an asset is marked as a favorite member of a burst photo sequence. Used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetBurstSelectionType
type PHAssetBurstSelectionType uint

// PHAssetEditOperation - Values identifying possible actions an asset can support, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetEditOperation
type PHAssetEditOperation uint

// PHAssetMediaSubtype - Constants identifying specific variations of asset media, such as panorama or screenshot photos, and time-lapse or high-frame-rate video.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype
type PHAssetMediaSubtype uint

const (
	// PHAssetMediaSubtypePhotoHDR - The asset is a high-dynamic range photo.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/photoHDR
	PHAssetMediaSubtypePhotoHDR PHAssetMediaSubtype = 0
)

// PHAssetMediaType - Identifies the general type of an asset, such as image or video.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaType
type PHAssetMediaType uint

// PHAssetSourceType - The means by which an asset enters the Photos library.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetSourceType
type PHAssetSourceType uint

// PHAuthorizationStatus - Information about your app’s authorization to access the user’s photo library.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAuthorizationStatus
type PHAuthorizationStatus uint

const (
	// PHAuthorizationStatusAuthorized - The user explicitly granted this app access to the photo library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAuthorizationStatus/authorized
	PHAuthorizationStatusAuthorized PHAuthorizationStatus = 0
	// PHAuthorizationStatusLimited - The user authorized this app for limited photo library access.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAuthorizationStatus/limited
	PHAuthorizationStatusLimited PHAuthorizationStatus = 0
	// PHAuthorizationStatusNotDetermined - The user hasn’t set the app’s authorization status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAuthorizationStatus/notDetermined
	PHAuthorizationStatusNotDetermined PHAuthorizationStatus = 0
)

// PHCollectionEditOperation - Values identifying possible actions that a collection can support, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionEditOperation
type PHCollectionEditOperation uint

// PHCollectionListSubtype - Major distinctions between kinds of collection list, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListSubtype
type PHCollectionListSubtype uint

// PHCollectionListType - Major distinctions between kinds of collection list, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListType
type PHCollectionListType uint

// PHImageContentMode - Options for fitting an image’s aspect ratio to a requested size, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageContentMode
type PHImageContentMode uint


