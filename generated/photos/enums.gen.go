// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

// Enum types and constants
// PHAccessLevel - The app’s level of access to the user’s photo library.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAccessLevel
type PHAccessLevel uint

const (
// PHAccessLevelAddOnly - A value that indicates the app may only add to the user’s photo library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAccessLevel/addOnly
PHAccessLevelAddOnly PHAccessLevel = 0
// PHAccessLevelReadWrite - A value that indicates the app can read from and write to the user’s photo library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAccessLevel/readWrite
PHAccessLevelReadWrite PHAccessLevel = 0
)

// PHAssetPlaybackStyle - An enumeration of asset playback styles that dictate how to present an asset to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/PlaybackStyle-swift.enum
type PHAssetPlaybackStyle uint

const (
// PHAssetPlaybackStyleImage - An enumeration indicating that the asset should be displayed as a still image.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/PlaybackStyle-swift.enum/image
PHAssetPlaybackStyleImage PHAssetPlaybackStyle = 0
// PHAssetPlaybackStyleImageAnimated - An enumeration indicating that the asset should be displayed as an animated image.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/PlaybackStyle-swift.enum/imageAnimated
PHAssetPlaybackStyleImageAnimated PHAssetPlaybackStyle = 0
// PHAssetPlaybackStyleLivePhoto - An enumeration indicating that the asset should be displayed as a Live Photo.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/PlaybackStyle-swift.enum/livePhoto
PHAssetPlaybackStyleLivePhoto PHAssetPlaybackStyle = 0
// PHAssetPlaybackStyleUnsupported - An enumeration indicating that the asset has an unsupported or undefined media playback type.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/PlaybackStyle-swift.enum/unsupported
PHAssetPlaybackStyleUnsupported PHAssetPlaybackStyle = 0
// PHAssetPlaybackStyleVideo - An enumeration indicating that the asset should be displayed as a video.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/PlaybackStyle-swift.enum/video
PHAssetPlaybackStyleVideo PHAssetPlaybackStyle = 0
// PHAssetPlaybackStyleVideoLooping - An enumeration indicating that the asset should be displayed as a looping video.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAsset/PlaybackStyle-swift.enum/videoLooping
PHAssetPlaybackStyleVideoLooping PHAssetPlaybackStyle = 0
)

// PHAssetBurstSelectionType - Bit mask values indicating whether and how an asset is marked as a favorite member of a burst photo sequence. Used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetBurstSelectionType
type PHAssetBurstSelectionType uint

// PHAssetCollectionSubtype - Minor distinctions between kinds of asset collections, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionSubtype
type PHAssetCollectionSubtype uint

const (
// PHAssetCollectionSubtypeAlbumMyPhotoStream - The user’s personal iCloud Photo Stream.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionSubtype/albumMyPhotoStream
PHAssetCollectionSubtypeAlbumMyPhotoStream PHAssetCollectionSubtype = 0
// PHAssetCollectionSubtypeAlbumRegular - An album created in the Photos app.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionSubtype/albumRegular
PHAssetCollectionSubtypeAlbumRegular PHAssetCollectionSubtype = 0
// PHAssetCollectionSubtypeSmartAlbumBursts - A Smart Album that groups all burst photo sequences in the photo library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionSubtype/smartAlbumBursts
PHAssetCollectionSubtypeSmartAlbumBursts PHAssetCollectionSubtype = 0
// PHAssetCollectionSubtypeSmartAlbumLivePhotos - A Smart Album that groups all Live Photos assets.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionSubtype/smartAlbumLivePhotos
PHAssetCollectionSubtypeSmartAlbumLivePhotos PHAssetCollectionSubtype = 0
// PHAssetCollectionSubtypeSmartAlbumRAW - A Smart Album that groups all RAW assets in the photo library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionSubtype/smartAlbumRAW
PHAssetCollectionSubtypeSmartAlbumRAW PHAssetCollectionSubtype = 0
// PHAssetCollectionSubtypeSmartAlbumRecentlyAdded - A Smart Album that groups all recently added assets in the photo library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionSubtype/smartAlbumRecentlyAdded
PHAssetCollectionSubtypeSmartAlbumRecentlyAdded PHAssetCollectionSubtype = 0
// PHAssetCollectionSubtypeSmartAlbumTimelapses - A Smart Album that groups all time-lapse videos in the photo library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionSubtype/smartAlbumTimelapses
PHAssetCollectionSubtypeSmartAlbumTimelapses PHAssetCollectionSubtype = 0
// PHAssetCollectionSubtypeSmartAlbumUnableToUpload - A Smart Album that groups all assets that the system can’t upload to iCloud.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionSubtype/smartAlbumUnableToUpload
PHAssetCollectionSubtypeSmartAlbumUnableToUpload PHAssetCollectionSubtype = 0
// PHAssetCollectionSubtypeSmartAlbumUserLibrary - A Smart Album that groups all assets that originate in the user’s own library (as opposed to assets from iCloud Shared Albums).
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionSubtype/smartAlbumUserLibrary
PHAssetCollectionSubtypeSmartAlbumUserLibrary PHAssetCollectionSubtype = 0
)

// PHAssetCollectionType - Major distinctions between kinds of asset collections, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionType
type PHAssetCollectionType uint

const (
// PHAssetCollectionTypeAlbum - An album in the Photos app.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionType/album
PHAssetCollectionTypeAlbum PHAssetCollectionType = 0
// PHAssetCollectionTypeMoment - A moment in the Photos app.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionType/moment
PHAssetCollectionTypeMoment PHAssetCollectionType = 0
// PHAssetCollectionTypeSmartAlbum - A smart album whose contents update dynamically.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionType/smartAlbum
PHAssetCollectionTypeSmartAlbum PHAssetCollectionType = 0
)

// PHAssetEditOperation - Values identifying possible actions an asset can support, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetEditOperation
type PHAssetEditOperation uint

const (
// PHAssetEditOperationContent - The asset’s photo or video content can be edited.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetEditOperation/content
PHAssetEditOperationContent PHAssetEditOperation = 0
// PHAssetEditOperationDelete - The asset can be deleted from the photo library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetEditOperation/delete
PHAssetEditOperationDelete PHAssetEditOperation = 0
// PHAssetEditOperationProperties - The asset’s metadata properties can be edited.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetEditOperation/properties
PHAssetEditOperationProperties PHAssetEditOperation = 0
)

// PHAssetMediaSubtype - Constants identifying specific variations of asset media, such as panorama or screenshot photos, and time-lapse or high-frame-rate video.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype
type PHAssetMediaSubtype uint

const (
// PHAssetMediaSubtypeNone - The asset has no subtype.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/PHAssetMediaSubtypeNone
PHAssetMediaSubtypeNone PHAssetMediaSubtype = 0
// PHAssetMediaSubtypePhotoDepthEffect - The asset is a photo captured with the Camera app’s Portrait mode depth effect.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/photoDepthEffect
PHAssetMediaSubtypePhotoDepthEffect PHAssetMediaSubtype = 0
// PHAssetMediaSubtypePhotoHDR - The asset is a high-dynamic range photo.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/photoHDR
PHAssetMediaSubtypePhotoHDR PHAssetMediaSubtype = 0
// PHAssetMediaSubtypePhotoLive - The asset is a Live Photo that includes movement and sounds from the moments just before and after its capture.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/photoLive
PHAssetMediaSubtypePhotoLive PHAssetMediaSubtype = 0
// PHAssetMediaSubtypePhotoPanorama - The asset is a large-format panorama photo.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/photoPanorama
PHAssetMediaSubtypePhotoPanorama PHAssetMediaSubtype = 0
// PHAssetMediaSubtypePhotoScreenshot - The asset is an image captured with the device’s screenshot feature.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/photoScreenshot
PHAssetMediaSubtypePhotoScreenshot PHAssetMediaSubtype = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/spatialMedia
PHAssetMediaSubtypeSpatialMedia PHAssetMediaSubtype = 0
// PHAssetMediaSubtypeVideoCinematic - The asset is a cinematic video.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/videoCinematic
PHAssetMediaSubtypeVideoCinematic PHAssetMediaSubtype = 0
// PHAssetMediaSubtypeVideoHighFrameRate - The asset is a high-frame-rate video.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/videoHighFrameRate
PHAssetMediaSubtypeVideoHighFrameRate PHAssetMediaSubtype = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/videoScreenRecording
PHAssetMediaSubtypeVideoScreenRecording PHAssetMediaSubtype = 0
// PHAssetMediaSubtypeVideoStreamed - The asset is a video with contents that always stream over a network connection.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/videoStreamed
PHAssetMediaSubtypeVideoStreamed PHAssetMediaSubtype = 0
// PHAssetMediaSubtypeVideoTimelapse - The asset is a time-lapse video.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaSubtype/videoTimelapse
PHAssetMediaSubtypeVideoTimelapse PHAssetMediaSubtype = 0
)

// PHAssetMediaType - Identifies the general type of an asset, such as image or video.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaType
type PHAssetMediaType uint

const (
// PHAssetMediaTypeAudio - The asset is an audio file.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaType/audio
PHAssetMediaTypeAudio PHAssetMediaType = 0
// PHAssetMediaTypeImage - The asset is a photo or other static image.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaType/image
PHAssetMediaTypeImage PHAssetMediaType = 0
// PHAssetMediaTypeUnknown - The asset’s type is unknown.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaType/unknown
PHAssetMediaTypeUnknown PHAssetMediaType = 0
// PHAssetMediaTypeVideo - The asset is a video file.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetMediaType/video
PHAssetMediaTypeVideo PHAssetMediaType = 0
)

// PHAssetResourceType - Describes the relationship of an asset resource to its owning asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceType
type PHAssetResourceType uint

const (
// PHAssetResourceTypeAdjustmentBasePairedVideo - Provides an unaltered version of the video data for a Live Photo asset for use in reconstructing recent edits.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceType/adjustmentBasePairedVideo
PHAssetResourceTypeAdjustmentBasePairedVideo PHAssetResourceType = 0
// PHAssetResourceTypeAdjustmentBasePhoto - Provides an unaltered version of its photo asset for use in for use in reconstructing recent edits.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceType/adjustmentBasePhoto
PHAssetResourceTypeAdjustmentBasePhoto PHAssetResourceType = 0
// PHAssetResourceTypeAdjustmentBaseVideo - Provides an unaltered version of its video asset.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceType/adjustmentBaseVideo
PHAssetResourceTypeAdjustmentBaseVideo PHAssetResourceType = 0
// PHAssetResourceTypeAlternatePhoto - Provides photo data that isn’t the primary form of its asset.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceType/alternatePhoto
PHAssetResourceTypeAlternatePhoto PHAssetResourceType = 0
// PHAssetResourceTypeAudio - Provides the original audio data for its asset.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceType/audio
PHAssetResourceTypeAudio PHAssetResourceType = 0
// PHAssetResourceTypeFullSizePairedVideo - Provides the current video data component of a Live Photo asset.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceType/fullSizePairedVideo
PHAssetResourceTypeFullSizePairedVideo PHAssetResourceType = 0
// PHAssetResourceTypeFullSizePhoto - Provides a modified version of the original photo asset.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceType/fullSizePhoto
PHAssetResourceTypeFullSizePhoto PHAssetResourceType = 0
// PHAssetResourceTypeFullSizeVideo - Provides a modified version of the original video asset.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceType/fullSizeVideo
PHAssetResourceTypeFullSizeVideo PHAssetResourceType = 0
// PHAssetResourceTypePairedVideo - Provides the original video data component of a Live Photo asset.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceType/pairedVideo
PHAssetResourceTypePairedVideo PHAssetResourceType = 0
// PHAssetResourceTypeVideo - Provides the original video data for its asset.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceType/video
PHAssetResourceTypeVideo PHAssetResourceType = 0
)

// PHAssetResourceUploadJobAction - These actions correspond with the types of fetches we can make on a PHAssetResourceUploadJob and the actions we can also take on those jobs.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/Action
type PHAssetResourceUploadJobAction uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/Action/acknowledge
PHAssetResourceUploadJobActionAcknowledge PHAssetResourceUploadJobAction = 0
// PHAssetResourceUploadJobActionRetry - Where PHAssetResourceUploadJobState = (success OR fail).
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/Action/retry
PHAssetResourceUploadJobActionRetry PHAssetResourceUploadJobAction = 0
)

// PHAssetResourceUploadJobState - The states of an upload job.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/State-swift.enum
type PHAssetResourceUploadJobState uint

const (
// PHAssetResourceUploadJobStateFailed - A request has been made to send the asset resource to the destination, but has not yet been fulfilled.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/State-swift.enum/failed
PHAssetResourceUploadJobStateFailed PHAssetResourceUploadJobState = 0
// PHAssetResourceUploadJobStatePending - The job has been registered.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/State-swift.enum/pending
PHAssetResourceUploadJobStatePending PHAssetResourceUploadJobState = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/State-swift.enum/registered
PHAssetResourceUploadJobStateRegistered PHAssetResourceUploadJobState = 0
// PHAssetResourceUploadJobStateSucceeded - The job has failed to send over.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceUploadJob/State-swift.enum/succeeded
PHAssetResourceUploadJobStateSucceeded PHAssetResourceUploadJobState = 0
)

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
// PHAuthorizationStatusDenied - The user explicitly denied this app access to the photo library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAuthorizationStatus/denied
PHAuthorizationStatusDenied PHAuthorizationStatus = 0
// PHAuthorizationStatusLimited - The user authorized this app for limited photo library access.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAuthorizationStatus/limited
PHAuthorizationStatusLimited PHAuthorizationStatus = 0
// PHAuthorizationStatusNotDetermined - The user hasn’t set the app’s authorization status.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAuthorizationStatus/notDetermined
PHAuthorizationStatusNotDetermined PHAuthorizationStatus = 0
// PHAuthorizationStatusRestricted - The app isn’t authorized to access the photo library, and the user can’t grant such permission.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAuthorizationStatus/restricted
PHAuthorizationStatusRestricted PHAuthorizationStatus = 0
)

// PHCollectionEditOperation - Values identifying possible actions that a collection can support, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionEditOperation
type PHCollectionEditOperation uint

const (
// PHCollectionEditOperationAddContent - The collection supports adding items that already exist elsewhere in the photo library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionEditOperation/addContent
PHCollectionEditOperationAddContent PHCollectionEditOperation = 0
// PHCollectionEditOperationCreateContent - The collection supports creating new items.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionEditOperation/createContent
PHCollectionEditOperationCreateContent PHCollectionEditOperation = 0
// PHCollectionEditOperationDelete - The collection itself can be deleted.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionEditOperation/delete
PHCollectionEditOperationDelete PHCollectionEditOperation = 0
// PHCollectionEditOperationDeleteContent - The collection supports deleting the items it contains.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionEditOperation/deleteContent
PHCollectionEditOperationDeleteContent PHCollectionEditOperation = 0
// PHCollectionEditOperationRearrangeContent - The collection supports reordering the arrangement of items it contains.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionEditOperation/rearrangeContent
PHCollectionEditOperationRearrangeContent PHCollectionEditOperation = 0
// PHCollectionEditOperationRemoveContent - The collection supports removing the items it contains.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionEditOperation/removeContent
PHCollectionEditOperationRemoveContent PHCollectionEditOperation = 0
// PHCollectionEditOperationRename - The collection itself can be renamed.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionEditOperation/rename
PHCollectionEditOperationRename PHCollectionEditOperation = 0
)

// PHCollectionListSubtype - Major distinctions between kinds of collection list, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListSubtype
type PHCollectionListSubtype uint

const (
// PHCollectionListSubtypeAny - Use this value to fetch collection lists of all possible subtypes.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListSubtype/any
PHCollectionListSubtypeAny PHCollectionListSubtype = 0
// PHCollectionListSubtypeMomentListCluster - The collection list is a moment cluster, grouping several related moments.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListSubtype/momentListCluster
PHCollectionListSubtypeMomentListCluster PHCollectionListSubtype = 0
// PHCollectionListSubtypeMomentListYear - The collection list is a moment year, grouping all moments from one or more calendar years.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListSubtype/momentListYear
PHCollectionListSubtypeMomentListYear PHCollectionListSubtype = 0
// PHCollectionListSubtypeRegularFolder - The collection list is a folder containing albums or other folders.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListSubtype/regularFolder
PHCollectionListSubtypeRegularFolder PHCollectionListSubtype = 0
// PHCollectionListSubtypeSmartFolderEvents - The collection list is a smart folder containing one or more Events synced from iPhoto.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListSubtype/smartFolderEvents
PHCollectionListSubtypeSmartFolderEvents PHCollectionListSubtype = 0
// PHCollectionListSubtypeSmartFolderFaces - The collection list is a smart folder containing one or more Faces synced from iPhoto.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListSubtype/smartFolderFaces
PHCollectionListSubtypeSmartFolderFaces PHCollectionListSubtype = 0
)

// PHCollectionListType - Major distinctions between kinds of collection list, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListType
type PHCollectionListType uint

const (
// PHCollectionListTypeFolder - A folder containing asset collections of type   or  .
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListType/folder
PHCollectionListTypeFolder PHCollectionListType = 0
// PHCollectionListTypeMomentList - A group of asset collections of type  .
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListType/momentList
PHCollectionListTypeMomentList PHCollectionListType = 0
// PHCollectionListTypeSmartFolder - A smart folder synced to the device from .
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListType/smartFolder
PHCollectionListTypeSmartFolder PHCollectionListType = 0
)

// PHImageContentMode - Options for fitting an image’s aspect ratio to a requested size, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageContentMode
type PHImageContentMode uint

// PHImageRequestOptionsDeliveryMode - Options for delivering requested image data, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsDeliveryMode
type PHImageRequestOptionsDeliveryMode uint

const (
// PHImageRequestOptionsDeliveryModeFastFormat - Photos provides only a fast-loading image, possibly sacrificing image quality.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsDeliveryMode/fastFormat
PHImageRequestOptionsDeliveryModeFastFormat PHImageRequestOptionsDeliveryMode = 0
// PHImageRequestOptionsDeliveryModeHighQualityFormat - Photos provides only the highest-quality image available, regardless of how much time it takes to load.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsDeliveryMode/highQualityFormat
PHImageRequestOptionsDeliveryModeHighQualityFormat PHImageRequestOptionsDeliveryMode = 0
// PHImageRequestOptionsDeliveryModeOpportunistic - Photos automatically provides one or more results in order to balance image quality and responsiveness.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsDeliveryMode/opportunistic
PHImageRequestOptionsDeliveryModeOpportunistic PHImageRequestOptionsDeliveryMode = 0
)

// PHImageRequestOptionsResizeMode - Options for how to resize the requested image to fit a target size, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsResizeMode
type PHImageRequestOptionsResizeMode uint

const (
// PHImageRequestOptionsResizeModeExact - Photos resizes the image to match the target size exactly.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsResizeMode/exact
PHImageRequestOptionsResizeModeExact PHImageRequestOptionsResizeMode = 0
// PHImageRequestOptionsResizeModeFast - Photos efficiently resizes the image to a size similar to, or slightly larger than, the target size.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsResizeMode/fast
PHImageRequestOptionsResizeModeFast PHImageRequestOptionsResizeMode = 0
// PHImageRequestOptionsResizeModeNone - Photos does not resize the image asset.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsResizeMode/none
PHImageRequestOptionsResizeModeNone PHImageRequestOptionsResizeMode = 0
)

// PHImageRequestOptionsVersion - Options for requesting an image asset with or without adjustments, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsVersion
type PHImageRequestOptionsVersion uint

const (
// PHImageRequestOptionsVersionCurrent - Request the most recent version of the image asset (the one that reflects all edits).
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsVersion/current
PHImageRequestOptionsVersionCurrent PHImageRequestOptionsVersion = 0
// PHImageRequestOptionsVersionOriginal - Request the original, highest-fidelity version of the image asset.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsVersion/original
PHImageRequestOptionsVersionOriginal PHImageRequestOptionsVersion = 0
// PHImageRequestOptionsVersionUnadjusted - Request a version of the image asset without adjustments.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageRequestOptionsVersion/unadjusted
PHImageRequestOptionsVersionUnadjusted PHImageRequestOptionsVersion = 0
)

// PHLivePhotoEditingErrorCode - Error codes for Live Photo editing errors.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingErrorCode
type PHLivePhotoEditingErrorCode uint

// PHLivePhotoFrameType - Identifiers for the type of frame image to be processed. Used with the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoFrameType
type PHLivePhotoFrameType uint

// PHPhotosError - Error codes for framework operations.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotosError-swift.struct/Code
type PHPhotosError uint

const (
// PHPhotosErrorInvalid - An error that indicates the operation isn’t valid.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotosError-c.enum/PHPhotosErrorInvalid
PHPhotosErrorInvalid PHPhotosError = 0
// PHPhotosErrorChangeNotSupported - An error that indicates the system doesn’t support the change request configuration.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotosError-swift.struct/Code/changeNotSupported
PHPhotosErrorChangeNotSupported PHPhotosError = 0
// PHPhotosErrorLibraryVolumeOffline - An error that indicates the photo library isn’t available because the file system volume that stores it isn’t mounted.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotosError-swift.struct/Code/libraryVolumeOffline
PHPhotosErrorLibraryVolumeOffline PHPhotosError = 0
// PHPhotosErrorPersistentChangeTokenExpired - An error that indicates the library state is older than the available history of persistent changes.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotosError-swift.struct/Code/persistentChangeTokenExpired
PHPhotosErrorPersistentChangeTokenExpired PHPhotosError = 0
// PHPhotosErrorRelinquishingLibraryBundleToWriter - An error that indicates the photo library isn’t available because the user moves, renames, or deletes the system’s photo library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotosError-swift.struct/Code/relinquishingLibraryBundleToWriter
PHPhotosErrorRelinquishingLibraryBundleToWriter PHPhotosError = 0
// PHPhotosErrorSwitchingSystemPhotoLibrary - An error that indicates the photo library isn’t available because the user switches the system’s photo library.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotosError-swift.struct/Code/switchingSystemPhotoLibrary
PHPhotosErrorSwitchingSystemPhotoLibrary PHPhotosError = 0
// PHPhotosErrorUserCancelled - An error that indicates the user cancels the asset retrieval or editing request.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotosError-swift.struct/Code/userCancelled
PHPhotosErrorUserCancelled PHPhotosError = 0
)

// PHVideoRequestOptionsDeliveryMode - Options for delivering requested video data, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHVideoRequestOptionsDeliveryMode
type PHVideoRequestOptionsDeliveryMode uint

const (
// PHVideoRequestOptionsDeliveryModeAutomatic - Photos automatically determines which quality of video data to provide based on the request and current conditions.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHVideoRequestOptionsDeliveryMode/automatic
PHVideoRequestOptionsDeliveryModeAutomatic PHVideoRequestOptionsDeliveryMode = 0
// PHVideoRequestOptionsDeliveryModeFastFormat - Photos provides whatever quality of video can be most quickly loaded.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHVideoRequestOptionsDeliveryMode/fastFormat
PHVideoRequestOptionsDeliveryModeFastFormat PHVideoRequestOptionsDeliveryMode = 0
// PHVideoRequestOptionsDeliveryModeHighQualityFormat - Photos provides only the highest quality video available.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHVideoRequestOptionsDeliveryMode/highQualityFormat
PHVideoRequestOptionsDeliveryModeHighQualityFormat PHVideoRequestOptionsDeliveryMode = 0
// PHVideoRequestOptionsDeliveryModeMediumQualityFormat - Photos provides a video of moderate quality unless a higher quality version is locally cached.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHVideoRequestOptionsDeliveryMode/mediumQualityFormat
PHVideoRequestOptionsDeliveryModeMediumQualityFormat PHVideoRequestOptionsDeliveryMode = 0
)

// PHVideoRequestOptionsVersion - Options for requesting a video asset with or without adjustments, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHVideoRequestOptionsVersion
type PHVideoRequestOptionsVersion uint

const (
// PHVideoRequestOptionsVersionCurrent - Request the most recent version of the video asset, reflecting all edits.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHVideoRequestOptionsVersion/current
PHVideoRequestOptionsVersionCurrent PHVideoRequestOptionsVersion = 0
// PHVideoRequestOptionsVersionOriginal - Request a version of the video asset without adjustments.
//
	// [Full Topic]: https://developer.apple.com/documentation/Photos/PHVideoRequestOptionsVersion/original
PHVideoRequestOptionsVersionOriginal PHVideoRequestOptionsVersion = 0
)


