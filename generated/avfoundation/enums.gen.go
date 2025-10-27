// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation


// Enum types and constants

// Error - An enumeration that defines the errors that framework operations can generate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code
type Error uint

const (
	// ErrorAirPlayControllerRequiresInternet - The AirPlay controller requires an internet connection to function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/airPlayControllerRequiresInternet
	ErrorAirPlayControllerRequiresInternet Error = 0
	// ErrorAirPlayReceiverRequiresInternet - The AirPlay receiver requires an internet connection to function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/airPlayReceiverRequiresInternet
	ErrorAirPlayReceiverRequiresInternet Error = 0
	// ErrorAirPlayReceiverTemporarilyUnavailable - An AirPlay receiver is temporarily unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/airPlayReceiverTemporarilyUnavailable
	ErrorAirPlayReceiverTemporarilyUnavailable Error = 0
	// ErrorApplicationIsNotAuthorized - The app isn’t authorized to play media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/applicationIsNotAuthorized
	ErrorApplicationIsNotAuthorized Error = 0
	// ErrorApplicationIsNotAuthorizedToUseDevice - The user denied this app permission to capture media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/applicationIsNotAuthorizedToUseDevice
	ErrorApplicationIsNotAuthorizedToUseDevice Error = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/autoWhiteBalanceNotLocked
	ErrorAutoWhiteBalanceNotLocked Error = 0
	// ErrorCompositionTrackSegmentsNotContiguous - The composition can’t add the source media because it contains gaps.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/compositionTrackSegmentsNotContiguous
	ErrorCompositionTrackSegmentsNotContiguous Error = 0
	// ErrorContentIsNotAuthorized - The user isn’t authorized to play the media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/contentIsNotAuthorized
	ErrorContentIsNotAuthorized Error = 0
	// ErrorContentIsProtected - The app isn’t authorized to open the media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/contentIsProtected
	ErrorContentIsProtected Error = 0
	// ErrorContentIsUnavailable - The captured content is unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/contentIsUnavailable
	ErrorContentIsUnavailable Error = 0
	// ErrorContentKeyRequestCancelled - The app canceled a request to retrieve a content key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/contentKeyRequestCancelled
	ErrorContentKeyRequestCancelled Error = 0
	// ErrorContentNotUpdated - The system couldn’t update the captured content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/contentNotUpdated
	ErrorContentNotUpdated Error = 0
	// ErrorCreateContentKeyRequestFailed - The app couldn’t create a content key request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/createContentKeyRequestFailed
	ErrorCreateContentKeyRequestFailed Error = 0
	// ErrorDecodeFailed - The system failed to decode the media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/decodeFailed
	ErrorDecodeFailed Error = 0
	// ErrorDecoderNotFound - The system can’t find a suitable decoder for the media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/decoderNotFound
	ErrorDecoderNotFound Error = 0
	// ErrorDecoderTemporarilyUnavailable - A suitable decoder for the media is temporarily available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/decoderTemporarilyUnavailable
	ErrorDecoderTemporarilyUnavailable Error = 0
	// ErrorDeviceAlreadyUsedByAnotherSession - Your app can’t access the device because another session is currently using it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/deviceAlreadyUsedByAnotherSession
	ErrorDeviceAlreadyUsedByAnotherSession Error = 0
	// ErrorDeviceInUseByAnotherApplication - Your app can’t access the device because another app is currently using it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/deviceInUseByAnotherApplication
	ErrorDeviceInUseByAnotherApplication Error = 0
	// ErrorDeviceIsNotAvailableInBackground - You attempted to start a capture session in the background, which isn’t allowed in iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/deviceIsNotAvailableInBackground
	ErrorDeviceIsNotAvailableInBackground Error = 0
	// ErrorDeviceLockedForConfigurationByAnotherProcess - Your app can’t change device settings because another process currently controls the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/deviceLockedForConfigurationByAnotherProcess
	ErrorDeviceLockedForConfigurationByAnotherProcess Error = 0
	// ErrorDeviceNotConnected - You app can’t access the device because it isn’t connected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/deviceNotConnected
	ErrorDeviceNotConnected Error = 0
	// ErrorDeviceWasDisconnected - A previously connected device is no longer accessible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/deviceWasDisconnected
	ErrorDeviceWasDisconnected Error = 0
	// ErrorDiskFull - Recording stopped because the disk is full.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/diskFull
	ErrorDiskFull Error = 0
	// ErrorDisplayWasDisabled - Screen capture failed because the display was inactive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/displayWasDisabled
	ErrorDisplayWasDisabled Error = 0
	// ErrorEncodeFailed - The system couldn’t encode the media data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/encodeFailed
	ErrorEncodeFailed Error = 0
	// ErrorEncoderNotFound - The requested encoder isn’t found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/encoderNotFound
	ErrorEncoderNotFound Error = 0
	// ErrorEncoderTemporarilyUnavailable - An appropriate encoder isn’t currently available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/encoderTemporarilyUnavailable
	ErrorEncoderTemporarilyUnavailable Error = 0
	// ErrorExportFailed - The requested export operation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/exportFailed
	ErrorExportFailed Error = 0
	// ErrorExternalPlaybackNotSupportedForAsset - The current asset doesn’t support playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/externalPlaybackNotSupportedForAsset
	ErrorExternalPlaybackNotSupportedForAsset Error = 0
	// ErrorFailedToLoadMediaData - The system can’t load the requested media data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/failedToLoadMediaData
	ErrorFailedToLoadMediaData Error = 0
	// ErrorFailedToLoadSampleData - The system can’t load the requested sample data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/failedToLoadSampleData
	ErrorFailedToLoadSampleData Error = 0
	// ErrorFailedToParse - The system can’t parse the media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/failedToParse
	ErrorFailedToParse Error = 0
	// ErrorFileAlreadyExists - A file with the same name exists at the location and you can’t overwrite it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/fileAlreadyExists
	ErrorFileAlreadyExists Error = 0
	// ErrorFileFailedToParse - The file is corrupt or in an unrecognized format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/fileFailedToParse
	ErrorFileFailedToParse Error = 0
	// ErrorFileFormatNotRecognized - The system can’t open the file because it’s in an unrecognized format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/fileFormatNotRecognized
	ErrorFileFormatNotRecognized Error = 0
	// ErrorFileTypeDoesNotSupportSampleReferences - The file type doesn’t support sample references.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/fileTypeDoesNotSupportSampleReferences
	ErrorFileTypeDoesNotSupportSampleReferences Error = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/followExternalSyncDeviceTimedOut
	ErrorFollowExternalSyncDeviceTimedOut Error = 0
	// ErrorFormatUnsupported - The current asset format isn’t supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/formatUnsupported
	ErrorFormatUnsupported Error = 0
	// ErrorIncompatibleAsset - You can’t display the media because the device isn’t capable of playing the content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/incompatibleAsset
	ErrorIncompatibleAsset Error = 0
	// ErrorIncorrectlyConfigured - The system is incorrectly configured for the requested operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/incorrectlyConfigured
	ErrorIncorrectlyConfigured Error = 0
	// ErrorInvalidCompositionTrackSegmentDuration - You can’t add the source media because its duration in the destination is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/invalidCompositionTrackSegmentDuration
	ErrorInvalidCompositionTrackSegmentDuration Error = 0
	// ErrorInvalidCompositionTrackSegmentSourceDuration - You can’t add the source media because it has no duration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/invalidCompositionTrackSegmentSourceDuration
	ErrorInvalidCompositionTrackSegmentSourceDuration Error = 0
	// ErrorInvalidCompositionTrackSegmentSourceStartTime - You can’t add the source media because its start time in the destination is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/invalidCompositionTrackSegmentSourceStartTime
	ErrorInvalidCompositionTrackSegmentSourceStartTime Error = 0
	// ErrorInvalidOutputURLPathExtension - The path extension of the output URL is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/invalidOutputURLPathExtension
	ErrorInvalidOutputURLPathExtension Error = 0
	// ErrorInvalidSampleCursor - An invalid sample cursor produced an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/invalidSampleCursor
	ErrorInvalidSampleCursor Error = 0
	// ErrorInvalidSourceMedia - The system couldn’t read the source media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/invalidSourceMedia
	ErrorInvalidSourceMedia Error = 0
	// ErrorInvalidVideoComposition - You attempted to present an unsupported video composition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/invalidVideoComposition
	ErrorInvalidVideoComposition Error = 0
	// ErrorMalformedDepth - The depth data isn’t properly structured.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/malformedDepth
	ErrorMalformedDepth Error = 0
	// ErrorMaximumDurationReached - The recording stopped because it reached the file’s maximum duration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/maximumDurationReached
	ErrorMaximumDurationReached Error = 0
	// ErrorMaximumFileSizeReached - The recording stopped because it reached the file’s maximum size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/maximumFileSizeReached
	ErrorMaximumFileSizeReached Error = 0
	// ErrorMaximumNumberOfSamplesForFileFormatReached - The recording stopped because it reached the file’s maximum number of samples.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/maximumNumberOfSamplesForFileFormatReached
	ErrorMaximumNumberOfSamplesForFileFormatReached Error = 0
	// ErrorMaximumStillImageCaptureRequestsExceeded - Your app can’t take a photo because there are too many unfinished photo capture requests.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/maximumStillImageCaptureRequestsExceeded
	ErrorMaximumStillImageCaptureRequestsExceeded Error = 0
	// ErrorMediaChanged - Recording stopped because the format of the source media changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/mediaChanged
	ErrorMediaChanged Error = 0
	// ErrorMediaDiscontinuity - Recording stopped because there was an interruption in the input media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/mediaDiscontinuity
	ErrorMediaDiscontinuity Error = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/mediaExtensionConflict
	ErrorMediaExtensionConflict Error = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/mediaExtensionDisabled
	ErrorMediaExtensionDisabled Error = 0
	// ErrorMediaServicesWereReset - The system couldn’t perform the operation because media services were unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/mediaServicesWereReset
	ErrorMediaServicesWereReset Error = 0
	// ErrorNoCompatibleAlternatesForExternalDisplay - The system found no compatible external displays.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/noCompatibleAlternatesForExternalDisplay
	ErrorNoCompatibleAlternatesForExternalDisplay Error = 0
	// ErrorNoDataCaptured - The recording failed because the system received no data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/noDataCaptured
	ErrorNoDataCaptured Error = 0
	// ErrorNoImageAtTime - No image is available in the media at the indicated time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/noImageAtTime
	ErrorNoImageAtTime Error = 0
	// ErrorNoLongerPlayable - The asset is no longer playable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/noLongerPlayable
	ErrorNoLongerPlayable Error = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/noSmartFramingsEnabled
	ErrorNoSmartFramingsEnabled Error = 0
	// ErrorNoSourceTrack - The asset doesn’t contain a source track.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/noSourceTrack
	ErrorNoSourceTrack Error = 0
	// ErrorOperationCancelled - The asset handled a request to cancel loading a property value asynchronously.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/operationCancelled
	ErrorOperationCancelled Error = 0
	// ErrorOperationInterrupted - An interruption occurred while performing a reading or writing operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/operationInterrupted
	ErrorOperationInterrupted Error = 0
	// ErrorOperationNotAllowed - The requested operation isn’t allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/operationNotAllowed
	ErrorOperationNotAllowed Error = 0
	// ErrorOperationNotSupportedForAsset - Your app attempted to perform an unsupported operation with the asset.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/operationNotSupportedForAsset
	ErrorOperationNotSupportedForAsset Error = 0
	// ErrorOperationNotSupportedForPreset - Your app attempted to perform an unsupported operation for the current preset.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/operationNotSupportedForPreset
	ErrorOperationNotSupportedForPreset Error = 0
	// ErrorOutOfMemory - The operation couldn’t finish because there isn’t enough memory available to process the media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/outOfMemory
	ErrorOutOfMemory Error = 0
	// ErrorRecordingAlreadyInProgress - Your app attempted to start recording a movie file while an existing recording is underway.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/recordingAlreadyInProgress
	ErrorRecordingAlreadyInProgress Error = 0
	// ErrorReferenceForbiddenByReferencePolicy - The current reference restrictions prevent the system from loading referenced media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/referenceForbiddenByReferencePolicy
	ErrorReferenceForbiddenByReferencePolicy Error = 0
	// ErrorRosettaNotInstalled - The system doesn’t have Rosetta installed and can’t perform the requested operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/rosettaNotInstalled
	ErrorRosettaNotInstalled Error = 0
	// ErrorSandboxExtensionDenied - The system denied issuing the sandbox extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/sandboxExtensionDenied
	ErrorSandboxExtensionDenied Error = 0
	// ErrorScreenCaptureFailed - An unexpected problem occurred that prevented screen capture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/screenCaptureFailed
	ErrorScreenCaptureFailed Error = 0
	// ErrorSegmentStartedWithNonSyncSample - The operation attempted to write a new MPEG-4 segment that didn’t start with a sync sample.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/segmentStartedWithNonSyncSample
	ErrorSegmentStartedWithNonSyncSample Error = 0
	// ErrorServerIncorrectlyConfigured - The configuration of the HTTP server that streams the media resource isn’t correct.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/serverIncorrectlyConfigured
	ErrorServerIncorrectlyConfigured Error = 0
	// ErrorSessionConfigurationChanged - Recording stopped because the configuration of media sources and destinations changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/sessionConfigurationChanged
	ErrorSessionConfigurationChanged Error = 0
	// ErrorSessionHardwareCostOverage - Your app requested too many camera hardware resources.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/sessionHardwareCostOverage
	ErrorSessionHardwareCostOverage Error = 0
	// ErrorSessionNotRunning - The recording couldn’t start because the session isn’t running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/sessionNotRunning
	ErrorSessionNotRunning Error = 0
	// ErrorSessionWasInterrupted - The recording stopped because the system interrupted the audio session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/sessionWasInterrupted
	ErrorSessionWasInterrupted Error = 0
	// ErrorToneMappingFailed - The requested tone mapping failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/toneMappingFailed
	ErrorToneMappingFailed Error = 0
	// ErrorTorchLevelUnavailable - The specified torch level is valid but currently unavailable, possibly due to overheating.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/torchLevelUnavailable
	ErrorTorchLevelUnavailable Error = 0
	// ErrorUndecodableMediaData - The system couldn’t decode the media data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/undecodableMediaData
	ErrorUndecodableMediaData Error = 0
	// ErrorUnknown - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/unknown
	ErrorUnknown Error = 0
	// ErrorUnsupportedDeviceActiveFormat - The capture session doesn’t support the camera device’s active format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/unsupportedDeviceActiveFormat
	ErrorUnsupportedDeviceActiveFormat Error = 0
	// ErrorUnsupportedOutputSettings - Your app requested unsupported output settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/unsupportedOutputSettings
	ErrorUnsupportedOutputSettings Error = 0
	// ErrorVideoCompositorFailed - The compositor couldn’t composite video frames.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code/videoCompositorFailed
	ErrorVideoCompositorFailed Error = 0
)


// CMTagCollectionVideoOutputPreset - Constants that indicate the type of video content to output.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/CMTagCollectionVideoOutputPreset
type CMTagCollectionVideoOutputPreset uint

const (
	// kCMTagCollectionVideoOutputPreset_Monoscopic - An output preset for monoscopic video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/CMTagCollectionVideoOutputPreset/kCMTagCollectionVideoOutputPreset_Monoscopic
	kCMTagCollectionVideoOutputPreset_Monoscopic CMTagCollectionVideoOutputPreset = 0
	// kCMTagCollectionVideoOutputPreset_Stereoscopic - An output preset for stereoscopic video.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/CMTagCollectionVideoOutputPreset/kCMTagCollectionVideoOutputPreset_Stereoscopic
	kCMTagCollectionVideoOutputPreset_Stereoscopic CMTagCollectionVideoOutputPreset = 0
)


// AssetExportSessionStatus - Values that indicate the state of an export session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/Status-swift.enum
type AssetExportSessionStatus uint

const (
	// AssetExportSessionStatusCancelled - You canceled the export.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/Status-swift.enum/cancelled
	AssetExportSessionStatusCancelled AssetExportSessionStatus = 0
	// AssetExportSessionStatusCompleted - The export completes successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/Status-swift.enum/completed
	AssetExportSessionStatusCompleted AssetExportSessionStatus = 0
	// AssetExportSessionStatusExporting - The export is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/Status-swift.enum/exporting
	AssetExportSessionStatusExporting AssetExportSessionStatus = 0
	// AssetExportSessionStatusFailed - The export fails.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/Status-swift.enum/failed
	AssetExportSessionStatusFailed AssetExportSessionStatus = 0
	// AssetExportSessionStatusUnknown - The session status is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/Status-swift.enum/unknown
	AssetExportSessionStatusUnknown AssetExportSessionStatus = 0
	// AssetExportSessionStatusWaiting - The session is waiting to export more data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/Status-swift.enum/waiting
	AssetExportSessionStatusWaiting AssetExportSessionStatus = 0
)


// AssetImageGeneratorResult - Constants that indicate the result of an image generation request.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/Result
type AssetImageGeneratorResult uint

const (
	// AssetImageGeneratorCancelled - A result that indicates you canceled image generation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/Result/cancelled
	AssetImageGeneratorCancelled AssetImageGeneratorResult = 0
	// AssetImageGeneratorFailed - A result that indicates that image generation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/Result/failed
	AssetImageGeneratorFailed AssetImageGeneratorResult = 0
	// AssetImageGeneratorSucceeded - A result that indicates that image generation succeeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/Result/succeeded
	AssetImageGeneratorSucceeded AssetImageGeneratorResult = 0
)


// AssetReaderStatus - Values that represent the possible states of an asset reader.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/Status-swift.enum
type AssetReaderStatus uint

const (
	// AssetReaderStatusCancelled - The asset reader can no longer read samples because you canceled reading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/Status-swift.enum/cancelled
	AssetReaderStatusCancelled AssetReaderStatus = 0
	// AssetReaderStatusCompleted - The asset reader completes reading all samples within its specified time range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/Status-swift.enum/completed
	AssetReaderStatusCompleted AssetReaderStatus = 0
	// AssetReaderStatusFailed - The asset reader can no longer read samples from its asset because of an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/Status-swift.enum/failed
	AssetReaderStatusFailed AssetReaderStatus = 0
	// AssetReaderStatusReading - The asset reader is successfully reading samples from its asset.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/Status-swift.enum/reading
	AssetReaderStatusReading AssetReaderStatus = 0
	// AssetReaderStatusUnknown - The asset reader is in an unknown state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/Status-swift.enum/unknown
	AssetReaderStatusUnknown AssetReaderStatus = 0
)


// AssetReferenceRestrictions - Restrictions to use when resolving references to external media data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReferenceRestrictions
type AssetReferenceRestrictions uint

const (
	// AssetReferenceRestrictionForbidNone - The asset should follow all media references.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReferenceRestrictions/AVAssetReferenceRestrictionForbidNone
	AssetReferenceRestrictionForbidNone AssetReferenceRestrictions = 0
	// AssetReferenceRestrictionDefaultPolicy - The asset should use the default reference restrictions policy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReferenceRestrictions/defaultPolicy
	AssetReferenceRestrictionDefaultPolicy AssetReferenceRestrictions = 0
	// AssetReferenceRestrictionForbidAll - The asset can only reference media stored within its container file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReferenceRestrictions/forbidAll
	AssetReferenceRestrictionForbidAll AssetReferenceRestrictions = 0
	// AssetReferenceRestrictionForbidCrossSiteReference - A remote asset shouldn’t follow references to remote media data stored at a different host.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReferenceRestrictions/forbidCrossSiteReference
	AssetReferenceRestrictionForbidCrossSiteReference AssetReferenceRestrictions = 0
	// AssetReferenceRestrictionForbidLocalReferenceToLocal - A local asset shouldn’t follow references to local media data stored outside its container file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReferenceRestrictions/forbidLocalReferenceToLocal
	AssetReferenceRestrictionForbidLocalReferenceToLocal AssetReferenceRestrictions = 0
	// AssetReferenceRestrictionForbidLocalReferenceToRemote - A local asset shouldn’t follow references to remote media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReferenceRestrictions/forbidLocalReferenceToRemote
	AssetReferenceRestrictionForbidLocalReferenceToRemote AssetReferenceRestrictions = 0
	// AssetReferenceRestrictionForbidRemoteReferenceToLocal - A remote asset shouldn’t follow references to local media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReferenceRestrictions/forbidRemoteReferenceToLocal
	AssetReferenceRestrictionForbidRemoteReferenceToLocal AssetReferenceRestrictions = 0
)


// AssetSegmentType - Constants that define the type of a segment.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentType
type AssetSegmentType uint

const (
	// AssetSegmentTypeInitialization - An initialization segment type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentType/initialization
	AssetSegmentTypeInitialization AssetSegmentType = 0
	// AssetSegmentTypeSeparable - A separable segment type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentType/separable
	AssetSegmentTypeSeparable AssetSegmentType = 0
)


// AssetTrackGroupOutputHandling - A type that specifies policies for how an export session processes alternate tracks in a track group.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackGroupOutputHandling
type AssetTrackGroupOutputHandling uint

const (
	// AssetTrackGroupOutputHandlingDefaultPolicy - The default track group output handling policy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackGroupOutputHandling/AVAssetTrackGroupOutputHandlingDefaultPolicy
	AssetTrackGroupOutputHandlingDefaultPolicy AssetTrackGroupOutputHandling = 0
	// AssetTrackGroupOutputHandlingNone - A policy that doesn’t pass through alternate audio tracks from the source asset during export.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackGroupOutputHandling/AVAssetTrackGroupOutputHandlingNone
	AssetTrackGroupOutputHandlingNone AssetTrackGroupOutputHandling = 0
	// AssetTrackGroupOutputHandlingPreserveAlternateTracks - A policy that passes through alternate audio tracks from the source asset during export.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackGroupOutputHandling/preserveAlternateTracks
	AssetTrackGroupOutputHandlingPreserveAlternateTracks AssetTrackGroupOutputHandling = 0
)


// AssetWriterStatus - Values that indicate the state of an asset writer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum
type AssetWriterStatus uint

const (
	// AssetWriterStatusCancelled - The asset writer canceled the writing operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum/cancelled
	AssetWriterStatusCancelled AssetWriterStatus = 0
	// AssetWriterStatusCompleted - The asset writer finishes writing successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum/completed
	AssetWriterStatusCompleted AssetWriterStatus = 0
	// AssetWriterStatusFailed - The asset writer fails to write the output file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum/failed
	AssetWriterStatusFailed AssetWriterStatus = 0
	// AssetWriterStatusUnknown - The asset writer’s status isn’t known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum/unknown
	AssetWriterStatusUnknown AssetWriterStatus = 0
	// AssetWriterStatusWriting - The asset writer is writing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum/writing
	AssetWriterStatusWriting AssetWriterStatus = 0
)


// AudioSpatializationFormats - A structure that defines the spatialization formats that a player item supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioSpatializationFormats
type AudioSpatializationFormats uint

const (
	// AudioSpatializationFormatNone - A value that indicates the player item doesn’t support audio spatialization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioSpatializationFormats/AVAudioSpatializationFormatNone
	AudioSpatializationFormatNone AudioSpatializationFormats = 0
	// AudioSpatializationFormatMonoAndStereo - A value that indicates the player item only supports mono and stereo layouts for audio spatialization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioSpatializationFormats/monoAndStereo
	AudioSpatializationFormatMonoAndStereo AudioSpatializationFormats = 0
	// AudioSpatializationFormatMonoStereoAndMultichannel - A value that indicates the player item supports mono, stereo, and multichannel layouts for audio spatialization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioSpatializationFormats/monoStereoAndMultichannel
	AudioSpatializationFormatMonoStereoAndMultichannel AudioSpatializationFormats = 0
	// AudioSpatializationFormatMultichannel - A value that indicates the player item only supports multichannel layouts for audio spatialization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioSpatializationFormats/multichannel
	AudioSpatializationFormatMultichannel AudioSpatializationFormats = 0
)


// AuthorizationStatus - Constants that indicate the status of an app’s authorization to capture media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus
type AuthorizationStatus uint

const (
	// AuthorizationStatusAuthorized - A status that indicates the user has explicitly granted an app permission to capture media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/authorized
	AuthorizationStatusAuthorized AuthorizationStatus = 0
	// AuthorizationStatusDenied - A status that indicates the user has explicitly denied an app permission to capture media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/denied
	AuthorizationStatusDenied AuthorizationStatus = 0
	// AuthorizationStatusNotDetermined - A status that indicates the user hasn’t yet granted or denied authorization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/notDetermined
	AuthorizationStatusNotDetermined AuthorizationStatus = 0
	// AuthorizationStatusRestricted - A status that indicates the app isn’t permitted to use media capture devices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus/restricted
	AuthorizationStatusRestricted AuthorizationStatus = 0
)


// CaptionAnimation - Animation options for a caption.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Animation-swift.enum
type CaptionAnimation uint

const (
	// CaptionAnimationCharacterReveal - A character reveal animation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Animation-swift.enum/characterReveal
	CaptionAnimationCharacterReveal CaptionAnimation = 0
	// CaptionAnimationNone - No animation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Animation-swift.enum/none
	CaptionAnimationNone CaptionAnimation = 0
)


// CaptionDecoration - Text decorations for caption text.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Decoration
type CaptionDecoration uint

const (
	// CaptionDecorationNone - No text decoration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionDecoration/AVCaptionDecorationNone
	CaptionDecorationNone CaptionDecoration = 0
	// CaptionDecorationLineThrough - A decoration representing a line through the text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Decoration/lineThrough
	CaptionDecorationLineThrough CaptionDecoration = 0
	// CaptionDecorationOverline - A decoration representing a line over the text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Decoration/overline
	CaptionDecorationOverline CaptionDecoration = 0
	// CaptionDecorationUnderline - A decoration representing a line under the text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Decoration/underline
	CaptionDecorationUnderline CaptionDecoration = 0
)


// CaptionFontStyle - Font styles for caption text.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/FontStyle
type CaptionFontStyle uint

const (
	// CaptionFontStyleItalic - An italic font style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/FontStyle/italic
	CaptionFontStyleItalic CaptionFontStyle = 0
	// CaptionFontStyleNormal - A normal font style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/FontStyle/normal
	CaptionFontStyleNormal CaptionFontStyle = 0
	// CaptionFontStyleUnknown - An unknown font style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/FontStyle/unknown
	CaptionFontStyleUnknown CaptionFontStyle = 0
)


// CaptionFontWeight - Font weights for a caption.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/FontWeight
type CaptionFontWeight uint

const (
	// CaptionFontWeightBold - A bold font weight.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/FontWeight/bold
	CaptionFontWeightBold CaptionFontWeight = 0
	// CaptionFontWeightNormal - A normal font weight.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/FontWeight/normal
	CaptionFontWeightNormal CaptionFontWeight = 0
	// CaptionFontWeightUnknown - An unknown font weight.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/FontWeight/unknown
	CaptionFontWeightUnknown CaptionFontWeight = 0
)


// CaptionTextAlignment - Text alignment options for a caption.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextAlignment-swift.enum
type CaptionTextAlignment uint

const (
	// CaptionTextAlignmentCenter - An alignment of the text to the center of the display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextAlignment-swift.enum/center
	CaptionTextAlignmentCenter CaptionTextAlignment = 0
	// CaptionTextAlignmentEnd - An alignment of the text to the end of the inline progression direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextAlignment-swift.enum/end
	CaptionTextAlignmentEnd CaptionTextAlignment = 0
	// CaptionTextAlignmentLeft - An alignment of the text to the left in horizontal writing mode, and top in vertical writing mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextAlignment-swift.enum/left
	CaptionTextAlignmentLeft CaptionTextAlignment = 0
	// CaptionTextAlignmentRight - An alignment of the text to the right in horizontal writing mode, and bottom in vertical writing mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextAlignment-swift.enum/right
	CaptionTextAlignmentRight CaptionTextAlignment = 0
	// CaptionTextAlignmentStart - An alignment of the text to the start of the inline progression direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextAlignment-swift.enum/start
	CaptionTextAlignmentStart CaptionTextAlignment = 0
)


// CaptionTextCombine - The caption’s supported rendering policy options.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextCombine
type CaptionTextCombine uint

const (
	// CaptionTextCombineAll - An option that combines all of the characters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextCombine/all
	CaptionTextCombineAll CaptionTextCombine = 0
	// CaptionTextCombineFourDigits - An option that combines four consecutive digits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextCombine/fourDigits
	CaptionTextCombineFourDigits CaptionTextCombine = 0
	// CaptionTextCombineNone - An option that doesn’t combine text upright.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextCombine/none
	CaptionTextCombineNone CaptionTextCombine = 0
	// CaptionTextCombineOneDigit - An option that makes one digit upright.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextCombine/oneDigit
	CaptionTextCombineOneDigit CaptionTextCombine = 0
	// CaptionTextCombineThreeDigits - An option that combines three consecutive digits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextCombine/threeDigits
	CaptionTextCombineThreeDigits CaptionTextCombine = 0
	// CaptionTextCombineTwoDigits - An option that combines two consecutive digits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextCombine/twoDigits
	CaptionTextCombineTwoDigits CaptionTextCombine = 0
)


// CaptionConversionValidatorStatus - Constants that indicate the status of a validator.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/Status-swift.enum
type CaptionConversionValidatorStatus uint

const (
	// CaptionConversionValidatorStatusCompleted - A status that indicates the system validation is complete.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/Status-swift.enum/completed
	CaptionConversionValidatorStatusCompleted CaptionConversionValidatorStatus = 0
	// CaptionConversionValidatorStatusStopped - A status that indicates the system validation stopped prior to completion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/Status-swift.enum/stopped
	CaptionConversionValidatorStatusStopped CaptionConversionValidatorStatus = 0
	// CaptionConversionValidatorStatusUnknown - A status that indicates the system didn’t initialize the validation operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/Status-swift.enum/unknown
	CaptionConversionValidatorStatusUnknown CaptionConversionValidatorStatus = 0
	// CaptionConversionValidatorStatusValidating - A status that indicates the system validation is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/Status-swift.enum/validating
	CaptionConversionValidatorStatusValidating CaptionConversionValidatorStatus = 0
)


// CaptionRegionDisplayAlignment - Constants that indicate the alignment of lines in a region.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/DisplayAlignment-swift.enum
type CaptionRegionDisplayAlignment uint

const (
	// CaptionRegionDisplayAlignmentAfter - An alignment that positions lines at the bottom of the block progression direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/DisplayAlignment-swift.enum/after
	CaptionRegionDisplayAlignmentAfter CaptionRegionDisplayAlignment = 0
	// CaptionRegionDisplayAlignmentBefore - An alignment that positions lines at the top of the block progression direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/DisplayAlignment-swift.enum/before
	CaptionRegionDisplayAlignmentBefore CaptionRegionDisplayAlignment = 0
	// CaptionRegionDisplayAlignmentCenter - An alignment that positions lines in the middle of the block progression direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/DisplayAlignment-swift.enum/center
	CaptionRegionDisplayAlignmentCenter CaptionRegionDisplayAlignment = 0
)


// CaptionRegionScroll - Constants that indicate the scrolling effects the system applies to a region.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/Scroll-swift.enum
type CaptionRegionScroll uint

const (
	// CaptionRegionScrollNone - A type that indicates no scrolling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/Scroll-swift.enum/none
	CaptionRegionScrollNone CaptionRegionScroll = 0
	// CaptionRegionScrollRollUp - A type that indicates a roll-up scroll effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/Scroll-swift.enum/rollUp
	CaptionRegionScrollRollUp CaptionRegionScroll = 0
)


// CaptionRegionWritingMode - Constants that indicate the writing mode for a region.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/WritingMode-swift.enum
type CaptionRegionWritingMode uint

const (
	// CaptionRegionWritingModeLeftToRightAndTopToBottom - A left-to-right and top-to-bottom writing mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/WritingMode-swift.enum/leftToRightAndTopToBottom
	CaptionRegionWritingModeLeftToRightAndTopToBottom CaptionRegionWritingMode = 0
	// CaptionRegionWritingModeTopToBottomAndRightToLeft - A top-to-bottom and right-to-left writing mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/WritingMode-swift.enum/topToBottomAndRightToLeft
	CaptionRegionWritingModeTopToBottomAndRightToLeft CaptionRegionWritingMode = 0
)


// CaptionRubyAlignment - Constants that indicate ruby text alignments.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRubyAlignment
type CaptionRubyAlignment uint

const (
	// CaptionRubyAlignmentCenter - An alignment with the ruby text at the center of ruby base.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRubyAlignment/center
	CaptionRubyAlignmentCenter CaptionRubyAlignment = 0
	// CaptionRubyAlignmentDistributeSpaceAround - An alignment with the ruby text so the spaces around each ruby text character are equal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRubyAlignment/distributeSpaceAround
	CaptionRubyAlignmentDistributeSpaceAround CaptionRubyAlignment = 0
	// CaptionRubyAlignmentDistributeSpaceBetween - An alignment with the ruby text so the spaces between the ruby text characters are equal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRubyAlignment/distributeSpaceBetween
	CaptionRubyAlignmentDistributeSpaceBetween CaptionRubyAlignment = 0
	// CaptionRubyAlignmentStart - An alignment with the ruby base and text at the left edge of horizontal text in a left-to-right inline progression, or at top of the vertical text in a top-to-bottom inline progression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRubyAlignment/start
	CaptionRubyAlignmentStart CaptionRubyAlignment = 0
)


// CaptionRubyPosition - Constants that indicate ruby text positions.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRubyPosition
type CaptionRubyPosition uint

const (
	// CaptionRubyPositionAfter - Display ruby text below horizontal text, or to the left of vertical text in a right-to-left block progression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRubyPosition/after
	CaptionRubyPositionAfter CaptionRubyPosition = 0
	// CaptionRubyPositionBefore - Display ruby text above horizontal text, or to the right of vertical text in a right-to-left block progression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRubyPosition/before
	CaptionRubyPositionBefore CaptionRubyPosition = 0
)


// CaptionUnitsType - A structure that defines a units for caption formats.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionUnitsType
type CaptionUnitsType uint

const (
	// CaptionUnitsTypeCells - A cell-based unit type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionUnitsType/cells
	CaptionUnitsTypeCells CaptionUnitsType = 0
	// CaptionUnitsTypePercent - A percentage-based unit type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionUnitsType/percent
	CaptionUnitsTypePercent CaptionUnitsType = 0
	// CaptionUnitsTypeUnspecified - An unspecified unit type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionUnitsType/unspecified
	CaptionUnitsTypeUnspecified CaptionUnitsType = 0
)


// CaptureCameraLensSmudgeDetectionStatus - Constants indicating the current camera lens smudge detection status.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureCameraLensSmudgeDetectionStatus
type CaptureCameraLensSmudgeDetectionStatus uint

const (
	// CaptureCameraLensSmudgeDetectionStatusDisabled - Indicates that the detection is not enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureCameraLensSmudgeDetectionStatus/disabled
	CaptureCameraLensSmudgeDetectionStatusDisabled CaptureCameraLensSmudgeDetectionStatus = 0
	// CaptureCameraLensSmudgeDetectionStatusSmudged - Indicates that the most recent detection found the camera lens to be smudged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureCameraLensSmudgeDetectionStatus/smudged
	CaptureCameraLensSmudgeDetectionStatusSmudged CaptureCameraLensSmudgeDetectionStatus = 0
	// CaptureCameraLensSmudgeDetectionStatusSmudgeNotDetected - Indicates that the most recent detection found no smudge on the camera lens.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureCameraLensSmudgeDetectionStatus/smudgeNotDetected
	CaptureCameraLensSmudgeDetectionStatusSmudgeNotDetected CaptureCameraLensSmudgeDetectionStatus = 0
	// CaptureCameraLensSmudgeDetectionStatusUnknown - Indicates that the detection result has not settled, commonly caused by excessive camera movement or the content of the scene.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureCameraLensSmudgeDetectionStatus/unknown
	CaptureCameraLensSmudgeDetectionStatusUnknown CaptureCameraLensSmudgeDetectionStatus = 0
)


// CaptureColorSpace - An enumeration of color spaces a device can support.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureColorSpace
type CaptureColorSpace uint

const (
	// CaptureColorSpace_AppleLog - The Apple Log Color space, which uses BT2020 as the color primaries, and an Apple-defined Log curve as a transfer function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureColorSpace/appleLog
	CaptureColorSpace_AppleLog CaptureColorSpace = 0
	// CaptureColorSpace_AppleLog2 - The Apple Log 2 Color space, which uses Apple Gamut as the color primaries, and an Apple defined Log curve as a transfer function. When you set this as the active color space on an  , any   or   connected to the same   is made inactive (its   property returns  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureColorSpace/appleLog2
	CaptureColorSpace_AppleLog2 CaptureColorSpace = 0
	// CaptureColorSpace_HLG_BT2020 - The BT.2020 wide color space that uses Illuminant D65 as the white point and Hybrid Log-Gamma (HLG) as the transfer function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureColorSpace/HLG_BT2020
	CaptureColorSpace_HLG_BT2020 CaptureColorSpace = 0
	// CaptureColorSpace_P3_D65 - The P3 D65 wide color space that uses Illuminant D65 as the white point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureColorSpace/P3_D65
	CaptureColorSpace_P3_D65 CaptureColorSpace = 0
	// CaptureColorSpace_sRGB - The standard RGB color space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureColorSpace/sRGB
	CaptureColorSpace_sRGB CaptureColorSpace = 0
)


// CaptureAutoFocusRangeRestriction - Constants to specify the autofocus range of a capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/AutoFocusRangeRestriction-swift.enum
type CaptureAutoFocusRangeRestriction uint

const (
	// CaptureAutoFocusRangeRestrictionFar - The device primarily attempts to focus on subjects far away from the camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/AutoFocusRangeRestriction-swift.enum/far
	CaptureAutoFocusRangeRestrictionFar CaptureAutoFocusRangeRestriction = 0
	// CaptureAutoFocusRangeRestrictionNear - The device primarily attempts to focus on subjects near the camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/AutoFocusRangeRestriction-swift.enum/near
	CaptureAutoFocusRangeRestrictionNear CaptureAutoFocusRangeRestriction = 0
	// CaptureAutoFocusRangeRestrictionNone - The device attempts to focus on objects at any range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/AutoFocusRangeRestriction-swift.enum/none
	CaptureAutoFocusRangeRestrictionNone CaptureAutoFocusRangeRestriction = 0
)


// CaptureCenterStageControlMode - Constants that indicate the current Center Stage control mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CenterStageControlMode-swift.enum
type CaptureCenterStageControlMode uint

const (
	// CaptureCenterStageControlModeApp - The app controls Center Stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CenterStageControlMode-swift.enum/app
	CaptureCenterStageControlModeApp CaptureCenterStageControlMode = 0
	// CaptureCenterStageControlModeCooperative - A user and app cooperatively share control of Center Stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CenterStageControlMode-swift.enum/cooperative
	CaptureCenterStageControlModeCooperative CaptureCenterStageControlMode = 0
	// CaptureCenterStageControlModeUser - The user controls Center Stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CenterStageControlMode-swift.enum/user
	CaptureCenterStageControlModeUser CaptureCenterStageControlMode = 0
)


// CaptureCinematicVideoFocusMode - Constants indicating the focus behavior when recording a Cinematic Video.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CinematicVideoFocusMode
type CaptureCinematicVideoFocusMode uint

const (
	// CaptureCinematicVideoFocusModeNone - Indicates that no focus mode is specified, in which case weak focus is used as default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CinematicVideoFocusMode/none
	CaptureCinematicVideoFocusModeNone CaptureCinematicVideoFocusMode = 0
	// CaptureCinematicVideoFocusModeStrong - Indicates that the subject should remain in focus until it exits the scene.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CinematicVideoFocusMode/strong
	CaptureCinematicVideoFocusModeStrong CaptureCinematicVideoFocusMode = 0
	// CaptureCinematicVideoFocusModeWeak - Indicates that the Cinematic Video algorithm should automatically adjust focus according to the prominence of the subjects in the scene.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CinematicVideoFocusMode/weak
	CaptureCinematicVideoFocusModeWeak CaptureCinematicVideoFocusMode = 0
)


// CaptureExposureMode - Constants that specify the exposure mode of a capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/ExposureMode-swift.enum
type CaptureExposureMode uint

const (
	// CaptureExposureModeAutoExpose - A mode that automatically adjusts the exposure one time, and then locks exposure for the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/ExposureMode-swift.enum/autoExpose
	CaptureExposureModeAutoExpose CaptureExposureMode = 0
	// CaptureExposureModeContinuousAutoExposure - A mode that continuously monitors exposure levels and automatically adjusts exposure when necessary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/ExposureMode-swift.enum/continuousAutoExposure
	CaptureExposureModeContinuousAutoExposure CaptureExposureMode = 0
	// CaptureExposureModeCustom - A mode where an app manually sets the exposure duration and ISO values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/ExposureMode-swift.enum/custom
	CaptureExposureModeCustom CaptureExposureMode = 0
	// CaptureExposureModeLocked - A mode that locks exposure for the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/ExposureMode-swift.enum/locked
	CaptureExposureModeLocked CaptureExposureMode = 0
)


// CaptureFlashMode - Constants that specify the flash modes of a capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/FlashMode-swift.enum
type CaptureFlashMode uint

const (
	// CaptureFlashModeAuto - A mode that indicates the device continuously monitors light levels and uses the flash when necessary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/FlashMode-swift.enum/auto
	CaptureFlashModeAuto CaptureFlashMode = 0
	// CaptureFlashModeOff - A mode that indicates the flash is off.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/FlashMode-swift.enum/off
	CaptureFlashModeOff CaptureFlashMode = 0
	// CaptureFlashModeOn - A mode that indicates the flash is on.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/FlashMode-swift.enum/on
	CaptureFlashModeOn CaptureFlashMode = 0
)


// CaptureFocusMode - Constants to specify the focus mode of a capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/FocusMode-swift.enum
type CaptureFocusMode uint

const (
	// CaptureFocusModeAutoFocus - A mode that automatically adjusts the focus one time, and then locks focus.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/FocusMode-swift.enum/autoFocus
	CaptureFocusModeAutoFocus CaptureFocusMode = 0
	// CaptureFocusModeContinuousAutoFocus - A mode that continuously monitors focus and autofocuses when necessary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/FocusMode-swift.enum/continuousAutoFocus
	CaptureFocusModeContinuousAutoFocus CaptureFocusMode = 0
	// CaptureFocusModeLocked - A mode that locks device focus.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/FocusMode-swift.enum/locked
	CaptureFocusModeLocked CaptureFocusMode = 0
)


// CaptureAutoFocusSystem - An enumeration of auto focus systems.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/AutoFocusSystem-swift.enum
type CaptureAutoFocusSystem uint

const (
	// CaptureAutoFocusSystemContrastDetection - A slower autofocus system based on differences in contrast.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/AutoFocusSystem-swift.enum/contrastDetection
	CaptureAutoFocusSystemContrastDetection CaptureAutoFocusSystem = 0
	// CaptureAutoFocusSystemNone - Autofocus isn’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/AutoFocusSystem-swift.enum/none
	CaptureAutoFocusSystemNone CaptureAutoFocusSystem = 0
	// CaptureAutoFocusSystemPhaseDetection - A faster autofoscus system based on differences in light phase.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/AutoFocusSystem-swift.enum/phaseDetection
	CaptureAutoFocusSystemPhaseDetection CaptureAutoFocusSystem = 0
)


// CaptureLensStabilizationStatus - Constants that indicate the status of optical image stabilization hardware during a bracketed photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/LensStabilizationStatus
type CaptureLensStabilizationStatus uint

const (
	// CaptureLensStabilizationStatusActive - Lens stabilization was active for the full duration of the photo capture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/LensStabilizationStatus/active
	CaptureLensStabilizationStatusActive CaptureLensStabilizationStatus = 0
	// CaptureLensStabilizationStatusOff - Lens stabilization isn’t specified for this photo capture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/LensStabilizationStatus/off
	CaptureLensStabilizationStatusOff CaptureLensStabilizationStatus = 0
	// CaptureLensStabilizationStatusOutOfRange - Lens stabilization was enabled for the photo capture, but device motion or capture duration exceeded the stabilization module’s correction limits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/LensStabilizationStatus/outOfRange
	CaptureLensStabilizationStatusOutOfRange CaptureLensStabilizationStatus = 0
	// CaptureLensStabilizationStatusUnavailable - Lens stabilization was temporarily unavailable during the photo capture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/LensStabilizationStatus/unavailable
	CaptureLensStabilizationStatusUnavailable CaptureLensStabilizationStatus = 0
	// CaptureLensStabilizationStatusUnsupported - Lens stabilization isn’t available on the device or device configuration that captured this photo.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/LensStabilizationStatus/unsupported
	CaptureLensStabilizationStatusUnsupported CaptureLensStabilizationStatus = 0
)


// CaptureMicrophoneMode - Constants that define the available microphone modes.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/MicrophoneMode
type CaptureMicrophoneMode uint

const (
	// CaptureMicrophoneModeStandard - A mode that processes microphone audio with standard voice DSP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/MicrophoneMode/standard
	CaptureMicrophoneModeStandard CaptureMicrophoneMode = 0
	// CaptureMicrophoneModeVoiceIsolation - A mode that processes microphone audio to isolate the voice and attenuate other signals.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/MicrophoneMode/voiceIsolation
	CaptureMicrophoneModeVoiceIsolation CaptureMicrophoneMode = 0
	// CaptureMicrophoneModeWideSpectrum - A mode that minimizes microphone audio processing to capture all sounds in the room.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/MicrophoneMode/wideSpectrum
	CaptureMicrophoneModeWideSpectrum CaptureMicrophoneMode = 0
)


// CaptureDevicePosition - Constants that indicate the physical position of a capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum
type CaptureDevicePosition uint

const (
	// CaptureDevicePositionBack - A position on the subject-facing side of an iOS device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum/back
	CaptureDevicePositionBack CaptureDevicePosition = 0
	// CaptureDevicePositionFront - A position on the user-facing side of an iOS device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum/front
	CaptureDevicePositionFront CaptureDevicePosition = 0
	// CaptureDevicePositionUnspecified - A position that’s unspecified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum/unspecified
	CaptureDevicePositionUnspecified CaptureDevicePosition = 0
)


// CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions - A structure that defines the conditions in which to restrict camera switching.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions-swift.struct
type CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions uint

const (
	// CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone - Disallow switching to a fallback camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions/AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone
	CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 0
	// CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionExposureModeChanged - Restrict switching to a fallback camera only when the device’s exposure mode changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions-swift.struct/exposureModeChanged
	CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionExposureModeChanged CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 0
	// CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionFocusModeChanged - Restrict switching to a fallback camera only when the device’s focus mode changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions-swift.struct/focusModeChanged
	CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionFocusModeChanged CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 0
	// CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionVideoZoomChanged - Restrict switching to a fallback camera only when the device’s video zoom changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions-swift.struct/videoZoomChanged
	CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionVideoZoomChanged CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 0
)


// CapturePrimaryConstituentDeviceSwitchingBehavior - Constants that control when to allow a virtual device to switch its active primary constituent device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum
type CapturePrimaryConstituentDeviceSwitchingBehavior uint

const (
	// CapturePrimaryConstituentDeviceSwitchingBehaviorAuto - The device automatically selects the best camera for the current scene.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/auto
	CapturePrimaryConstituentDeviceSwitchingBehaviorAuto CapturePrimaryConstituentDeviceSwitchingBehavior = 0
	// CapturePrimaryConstituentDeviceSwitchingBehaviorLocked - The device locks camera switching to the active primary constituent device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/locked
	CapturePrimaryConstituentDeviceSwitchingBehaviorLocked CapturePrimaryConstituentDeviceSwitchingBehavior = 0
	// CapturePrimaryConstituentDeviceSwitchingBehaviorRestricted - The device restricts fallback camera selection to certain conditions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/restricted
	CapturePrimaryConstituentDeviceSwitchingBehaviorRestricted CapturePrimaryConstituentDeviceSwitchingBehavior = 0
	// CapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported - The device doesn’t support constituent device switching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum/unsupported
	CapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported CapturePrimaryConstituentDeviceSwitchingBehavior = 0
)


// CaptureSystemPressureFactors - A structure that defines the factors affecting capture system performance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemPressureState-swift.class/Factors-swift.struct
type CaptureSystemPressureFactors uint

const (
	// CaptureSystemPressureFactorNone - System pressure is currently nominal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemPressureFactors/AVCaptureSystemPressureFactorNone
	CaptureSystemPressureFactorNone CaptureSystemPressureFactors = 0
	// CaptureSystemPressureFactorCameraTemperature - The camera module is operating at an elevated temperature.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemPressureState-swift.class/Factors-swift.struct/cameraTemperature
	CaptureSystemPressureFactorCameraTemperature CaptureSystemPressureFactors = 0
	// CaptureSystemPressureFactorDepthModuleTemperature - The module capturing depth information is operating at an elevated temperature.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemPressureState-swift.class/Factors-swift.struct/depthModuleTemperature
	CaptureSystemPressureFactorDepthModuleTemperature CaptureSystemPressureFactors = 0
	// CaptureSystemPressureFactorPeakPower - The system’s peak power requirements exceed the battery’s current capacity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemPressureState-swift.class/Factors-swift.struct/peakPower
	CaptureSystemPressureFactorPeakPower CaptureSystemPressureFactors = 0
	// CaptureSystemPressureFactorSystemTemperature - The entire system is under elevated thermal load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemPressureState-swift.class/Factors-swift.struct/systemTemperature
	CaptureSystemPressureFactorSystemTemperature CaptureSystemPressureFactors = 0
)


// CaptureSystemUserInterface - Constants that describe the capture device configuration user interfaces.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemUserInterface
type CaptureSystemUserInterface uint

const (
	// CaptureSystemUserInterfaceMicrophoneModes - The system user interface for selecting microphone modes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemUserInterface/microphoneModes
	CaptureSystemUserInterfaceMicrophoneModes CaptureSystemUserInterface = 0
	// CaptureSystemUserInterfaceVideoEffects - The system user interface for changing the state of video effects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemUserInterface/videoEffects
	CaptureSystemUserInterfaceVideoEffects CaptureSystemUserInterface = 0
)


// CaptureTorchMode - Constants to specify the capture device’s torch mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/TorchMode-swift.enum
type CaptureTorchMode uint

const (
	// CaptureTorchModeAuto - The capture device continuously monitors light levels and uses the torch when necessary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/TorchMode-swift.enum/auto
	CaptureTorchModeAuto CaptureTorchMode = 0
	// CaptureTorchModeOff - The capture device torch is always off.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/TorchMode-swift.enum/off
	CaptureTorchModeOff CaptureTorchMode = 0
	// CaptureTorchModeOn - The capture device torch is always on.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/TorchMode-swift.enum/on
	CaptureTorchModeOn CaptureTorchMode = 0
)


// CaptureDeviceTransportControlsPlaybackMode - Constants that indicate the transport control’s current mode of playback, if it has one.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/TransportControlsPlaybackMode-swift.enum
type CaptureDeviceTransportControlsPlaybackMode uint

const (
	// CaptureDeviceTransportControlsNotPlayingMode - A value that indicates that the tape transport isn’t threaded through the play head.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/TransportControlsPlaybackMode-swift.enum/notPlaying
	CaptureDeviceTransportControlsNotPlayingMode CaptureDeviceTransportControlsPlaybackMode = 0
	// CaptureDeviceTransportControlsPlayingMode - A value that indicates that the tape transport is threaded through the play head.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/TransportControlsPlaybackMode-swift.enum/playing
	CaptureDeviceTransportControlsPlayingMode CaptureDeviceTransportControlsPlaybackMode = 0
)


// CaptureWhiteBalanceMode - Constants to specify the white balance mode of a capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceMode-swift.enum
type CaptureWhiteBalanceMode uint

const (
	// CaptureWhiteBalanceModeAutoWhiteBalance - A mode that automatically manages white balance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceMode-swift.enum/autoWhiteBalance
	CaptureWhiteBalanceModeAutoWhiteBalance CaptureWhiteBalanceMode = 0
	// CaptureWhiteBalanceModeContinuousAutoWhiteBalance - A mode that continuously monitors white balance and adjusts when necessary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceMode-swift.enum/continuousAutoWhiteBalance
	CaptureWhiteBalanceModeContinuousAutoWhiteBalance CaptureWhiteBalanceMode = 0
	// CaptureWhiteBalanceModeLocked - A mode that locks the white balance state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceMode-swift.enum/locked
	CaptureWhiteBalanceModeLocked CaptureWhiteBalanceMode = 0
)


// CaptureMultichannelAudioMode - Constants that indicate the modes of multichannel audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultichannelAudioMode
type CaptureMultichannelAudioMode uint

const (
	// CaptureMultichannelAudioModeFirstOrderAmbisonics - An audio mode that indicates the recording uses first-order ambisonics.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultichannelAudioMode/firstOrderAmbisonics
	CaptureMultichannelAudioModeFirstOrderAmbisonics CaptureMultichannelAudioMode = 0
	// CaptureMultichannelAudioModeNone - A mode that indicates there’s no multichannel audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultichannelAudioMode/none
	CaptureMultichannelAudioModeNone CaptureMultichannelAudioMode = 0
	// CaptureMultichannelAudioModeStereo - A mode that indicates the recording uses stereo audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultichannelAudioMode/stereo
	CaptureMultichannelAudioModeStereo CaptureMultichannelAudioMode = 0
)


// CaptureOutputDataDroppedReason - Constants that define reasons for why the system dropped a frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/DataDroppedReason
type CaptureOutputDataDroppedReason uint

const (
	// CaptureOutputDataDroppedReasonDiscontinuity - The system dropped data because the device providing data experienced a discontinuity, and the output lost an unknown number of data objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/DataDroppedReason/discontinuity
	CaptureOutputDataDroppedReasonDiscontinuity CaptureOutputDataDroppedReason = 0
	// CaptureOutputDataDroppedReasonLateData - The system dropped data because you’ve configured capture output to drop data when delegate queue is in a blocked state, and there’s data to deliver.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/DataDroppedReason/lateData
	CaptureOutputDataDroppedReasonLateData CaptureOutputDataDroppedReason = 0
	// CaptureOutputDataDroppedReasonNone - The system didn’t drop data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/DataDroppedReason/none
	CaptureOutputDataDroppedReasonNone CaptureOutputDataDroppedReason = 0
	// CaptureOutputDataDroppedReasonOutOfBuffers - The system dropped data because the capture output exhausted its internal pool of memory buffers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/DataDroppedReason/outOfBuffers
	CaptureOutputDataDroppedReasonOutOfBuffers CaptureOutputDataDroppedReason = 0
)


// CapturePhotoOutputCaptureReadiness - Constants that indicate whether the output is ready to receive capture requests.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/CaptureReadiness-swift.enum
type CapturePhotoOutputCaptureReadiness uint

const (
	// CapturePhotoOutputCaptureReadinessNotReadyMomentarily - Indicates that the output isn’t ready to receive requests, but may be ready shortly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/CaptureReadiness-swift.enum/notReadyMomentarily
	CapturePhotoOutputCaptureReadinessNotReadyMomentarily CapturePhotoOutputCaptureReadiness = 0
	// CapturePhotoOutputCaptureReadinessNotReadyWaitingForCapture - Indicates that the output isn’t ready to receive requests for a longer duration because it’s busy capturing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/CaptureReadiness-swift.enum/notReadyWaitingForCapture
	CapturePhotoOutputCaptureReadinessNotReadyWaitingForCapture CapturePhotoOutputCaptureReadiness = 0
	// CapturePhotoOutputCaptureReadinessNotReadyWaitingForProcessing - Indicates that the output isn’t ready to receive requests for a longer duration because it’s busy processing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/CaptureReadiness-swift.enum/notReadyWaitingForProcessing
	CapturePhotoOutputCaptureReadinessNotReadyWaitingForProcessing CapturePhotoOutputCaptureReadiness = 0
	// CapturePhotoOutputCaptureReadinessReady - Indicates that the output is ready to receive new requests.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/CaptureReadiness-swift.enum/ready
	CapturePhotoOutputCaptureReadinessReady CapturePhotoOutputCaptureReadiness = 0
	// CapturePhotoOutputCaptureReadinessSessionNotRunning - Indicates that the session isn’t running and the output isn’t ready to receive requests.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/CaptureReadiness-swift.enum/sessionNotRunning
	CapturePhotoOutputCaptureReadinessSessionNotRunning CapturePhotoOutputCaptureReadiness = 0
)


// CapturePhotoQualityPrioritization - Constants that indicate how to prioritize photo quality relative to capture speed.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/QualityPrioritization
type CapturePhotoQualityPrioritization uint

const (
	// CapturePhotoQualityPrioritizationBalanced - Priority is balanced between photo quality and speed of delivery.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/QualityPrioritization/balanced
	CapturePhotoQualityPrioritizationBalanced CapturePhotoQualityPrioritization = 0
	// CapturePhotoQualityPrioritizationQuality - Photo quality is most important, even at the expense of shot-to-shot time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/QualityPrioritization/quality
	CapturePhotoQualityPrioritizationQuality CapturePhotoQualityPrioritization = 0
	// CapturePhotoQualityPrioritizationSpeed - Speed of photo delivery is most important, even at the expense of quality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/QualityPrioritization/speed
	CapturePhotoQualityPrioritizationSpeed CapturePhotoQualityPrioritization = 0
)


// CaptureSessionInterruptionReason - Constants identifying the reason a capture session was interrupted, found in an 
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason
type CaptureSessionInterruptionReason uint

const (
	// CaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient - An interruption caused by the audio hardware temporarily being made unavailable (for example, for a phone call or alarm).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/audioDeviceInUseByAnotherClient
	CaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonSensitiveContentMitigationActivated - An interruption caused by a   when it detects sensitive content on an associated  .  To resume your capture session, call your analyzer’s   method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/sensitiveContentMitigationActivated
	CaptureSessionInterruptionReasonSensitiveContentMitigationActivated CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient - An interruption caused by the video device temporarily being made unavailable (for example, when used by another capture session).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceInUseByAnotherClient
	CaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure - An interruption due to system pressure, such as thermal duress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableDueToSystemPressure
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground - An interruption caused by the app being sent to the background while using a camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableInBackground
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground CaptureSessionInterruptionReason = 0
	// CaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps - An interruption caused when your app is running in Slide Over, Split View, or Picture in Picture mode on iPad.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason/videoDeviceNotAvailableWithMultipleForegroundApps
	CaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps CaptureSessionInterruptionReason = 0
)


// CaptureTimecodeSourceType - Defines possible sources for generating timecode in using a timecode generator.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/SourceType-swift.enum
type CaptureTimecodeSourceType uint

const (
	// CaptureTimecodeSourceTypeExternal - Synchronizes timecode to an external timecode data stream. Ideal for professional audio and video synchronization with external quarter-frame MIDI or HID timecode hardware.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/SourceType-swift.enum/external
	CaptureTimecodeSourceTypeExternal CaptureTimecodeSourceType = 0
	// CaptureTimecodeSourceTypeFrameCount - No internal or external source is adopted. Timecodes are zero-based, sequentially generated frame counts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/SourceType-swift.enum/frameCount
	CaptureTimecodeSourceTypeFrameCount CaptureTimecodeSourceType = 0
	// CaptureTimecodeSourceTypeRealTimeClock - Synchronizes timecode to the system clock for real-time applications. Useful for live events or scenarios requiring alignment with the actual time of day.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/SourceType-swift.enum/realTimeClock
	CaptureTimecodeSourceTypeRealTimeClock CaptureTimecodeSourceType = 0
)


// CaptureTimecodeGeneratorSynchronizationStatus - Constants defining the synchronization status of a timecode generator .
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/SynchronizationStatus
type CaptureTimecodeGeneratorSynchronizationStatus uint

const (
	// CaptureTimecodeGeneratorSynchronizationStatusNotRequired - The timecode generator does not require active synchronization for a given source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/SynchronizationStatus/notRequired
	CaptureTimecodeGeneratorSynchronizationStatusNotRequired CaptureTimecodeGeneratorSynchronizationStatus = 0
	// CaptureTimecodeGeneratorSynchronizationStatusSourceSelected - A timecode source has been selected, but synchronization has not yet started.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/SynchronizationStatus/sourceSelected
	CaptureTimecodeGeneratorSynchronizationStatusSourceSelected CaptureTimecodeGeneratorSynchronizationStatus = 0
	// CaptureTimecodeGeneratorSynchronizationStatusSourceUnavailable - The timecode generator has failed to establish a connection with a given source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/SynchronizationStatus/sourceUnavailable
	CaptureTimecodeGeneratorSynchronizationStatusSourceUnavailable CaptureTimecodeGeneratorSynchronizationStatus = 0
	// CaptureTimecodeGeneratorSynchronizationStatusSourceUnsupported - The timecode generator is receiving data from the source in an unrecognized format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/SynchronizationStatus/sourceUnsupported
	CaptureTimecodeGeneratorSynchronizationStatusSourceUnsupported CaptureTimecodeGeneratorSynchronizationStatus = 0
	// CaptureTimecodeGeneratorSynchronizationStatusSynchronized - The timecode generator is successfully synchronized to the selected source, maintaining active timing alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/SynchronizationStatus/synchronized
	CaptureTimecodeGeneratorSynchronizationStatusSynchronized CaptureTimecodeGeneratorSynchronizationStatus = 0
	// CaptureTimecodeGeneratorSynchronizationStatusSynchronizing - The timecode generator is actively synchronizing to the selected source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/SynchronizationStatus/synchronizing
	CaptureTimecodeGeneratorSynchronizationStatusSynchronizing CaptureTimecodeGeneratorSynchronizationStatus = 0
	// CaptureTimecodeGeneratorSynchronizationStatusTimedOut - The synchronization has timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/SynchronizationStatus/timedOut
	CaptureTimecodeGeneratorSynchronizationStatusTimedOut CaptureTimecodeGeneratorSynchronizationStatus = 0
	// CaptureTimecodeGeneratorSynchronizationStatusUnknown - The initial state before a source is selected or during error conditions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/SynchronizationStatus/unknown
	CaptureTimecodeGeneratorSynchronizationStatusUnknown CaptureTimecodeGeneratorSynchronizationStatus = 0
)


// CaptureVideoOrientation - Constants indicating video orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoOrientation
type CaptureVideoOrientation uint

const (
	// CaptureVideoOrientationLandscapeLeft - Indicates that video should be oriented horizontally, top on the right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoOrientation/landscapeLeft
	CaptureVideoOrientationLandscapeLeft CaptureVideoOrientation = 0
	// CaptureVideoOrientationLandscapeRight - Indicates that video should be oriented horizontally, top on the left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoOrientation/landscapeRight
	CaptureVideoOrientationLandscapeRight CaptureVideoOrientation = 0
	// CaptureVideoOrientationPortrait - Indicates that video should be oriented vertically, top at the top.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoOrientation/portrait
	CaptureVideoOrientationPortrait CaptureVideoOrientation = 0
	// CaptureVideoOrientationPortraitUpsideDown - Indicates that video should be oriented vertically, top at the bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoOrientation/portraitUpsideDown
	CaptureVideoOrientationPortraitUpsideDown CaptureVideoOrientation = 0
)


// CaptureVideoStabilizationMode - An enumeration of video stabilization modes that capture devices and formats support.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoStabilizationMode
type CaptureVideoStabilizationMode uint

const (
	// CaptureVideoStabilizationModeAuto - A mode that indicates the system chooses the most appropriate video stabilization mode for the device and format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoStabilizationMode/auto
	CaptureVideoStabilizationModeAuto CaptureVideoStabilizationMode = 0
	// CaptureVideoStabilizationModeCinematic - A mode that uses the cinematic stabilization algorithm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoStabilizationMode/cinematic
	CaptureVideoStabilizationModeCinematic CaptureVideoStabilizationMode = 0
	// CaptureVideoStabilizationModeCinematicExtended - A mode that uses the extended cinematic stabilization algorithm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoStabilizationMode/cinematicExtended
	CaptureVideoStabilizationModeCinematicExtended CaptureVideoStabilizationMode = 0
	// CaptureVideoStabilizationModeCinematicExtendedEnhanced - A mode that stabilizes video using the enhanced extended cinematic stabilization algorithm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoStabilizationMode/cinematicExtendedEnhanced
	CaptureVideoStabilizationModeCinematicExtendedEnhanced CaptureVideoStabilizationMode = 0
	// CaptureVideoStabilizationModeLowLatency - Indicates that video should be stabilized using the low latency stabilization algorithm. Low Latency stabilization has a reduced field of view. Enabling low latency stabilization introduces no additional latency into the video capture pipeline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoStabilizationMode/lowLatency
	CaptureVideoStabilizationModeLowLatency CaptureVideoStabilizationMode = 0
	// CaptureVideoStabilizationModeOff - A mode that doesn’t stabilize video capture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoStabilizationMode/off
	CaptureVideoStabilizationModeOff CaptureVideoStabilizationMode = 0
	// CaptureVideoStabilizationModePreviewOptimized - A mode that uses the preview optimized stabilization algorithm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoStabilizationMode/previewOptimized
	CaptureVideoStabilizationModePreviewOptimized CaptureVideoStabilizationMode = 0
	// CaptureVideoStabilizationModeStandard - A mode that uses the standard algorithm.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoStabilizationMode/standard
	CaptureVideoStabilizationModeStandard CaptureVideoStabilizationMode = 0
)


// ContentAuthorizationStatus - A value representing the status of a content authorization request.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentAuthorizationStatus
type ContentAuthorizationStatus uint

const (
	// ContentAuthorizationBusy - The last call to request content authorization couldn’t be completed because another asset is currently attempting authorization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentAuthorizationStatus/busy
	ContentAuthorizationBusy ContentAuthorizationStatus = 0
	// ContentAuthorizationCancelled - The last call to request content authorization was cancelled by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentAuthorizationStatus/cancelled
	ContentAuthorizationCancelled ContentAuthorizationStatus = 0
	// ContentAuthorizationCompleted - The last completed call to request content authorization completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentAuthorizationStatus/completed
	ContentAuthorizationCompleted ContentAuthorizationStatus = 0
	// ContentAuthorizationNotAvailable - The last call to request content authorization couldn’t be completed because there was no known mechanism by which to attempt authorization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentAuthorizationStatus/notAvailable
	ContentAuthorizationNotAvailable ContentAuthorizationStatus = 0
	// ContentAuthorizationNotPossible - The last call to request content authorization couldn’t be completed in a non-recoverable way.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentAuthorizationStatus/notPossible
	ContentAuthorizationNotPossible ContentAuthorizationStatus = 0
	// ContentAuthorizationTimedOut - The last call to request content authorization was cancelled because the timeout interval was reached.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentAuthorizationStatus/timedOut
	ContentAuthorizationTimedOut ContentAuthorizationStatus = 0
	// ContentAuthorizationUnknown - The content authorization content request hasn’t completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentAuthorizationStatus/unknown
	ContentAuthorizationUnknown ContentAuthorizationStatus = 0
)


// ContentKeyRequestStatus - The status for a content key request.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/Status-swift.enum
type ContentKeyRequestStatus uint

const (
	// ContentKeyRequestStatusCancelled - The key request was canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/Status-swift.enum/cancelled
	ContentKeyRequestStatusCancelled ContentKeyRequestStatus = 0
	// ContentKeyRequestStatusFailed - The key request failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/Status-swift.enum/failed
	ContentKeyRequestStatusFailed ContentKeyRequestStatus = 0
	// ContentKeyRequestStatusReceivedResponse - The key request was received, and the key is in use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/Status-swift.enum/receivedResponse
	ContentKeyRequestStatusReceivedResponse ContentKeyRequestStatus = 0
	// ContentKeyRequestStatusRenewed - The key request was renewed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/Status-swift.enum/renewed
	ContentKeyRequestStatusRenewed ContentKeyRequestStatus = 0
	// ContentKeyRequestStatusRequestingResponse - The key request was just created.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/Status-swift.enum/requestingResponse
	ContentKeyRequestStatusRequestingResponse ContentKeyRequestStatus = 0
	// ContentKeyRequestStatusRetried - The key request was retried.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/Status-swift.enum/retried
	ContentKeyRequestStatusRetried ContentKeyRequestStatus = 0
)


// DelegatingPlaybackCoordinatorRateChangeOptions - Constants that define rate change options.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorRateChangeOptions
type DelegatingPlaybackCoordinatorRateChangeOptions uint

const (
	// DelegatingPlaybackCoordinatorRateChangeOptionPlayImmediately - Indicates that the coordinator should begin playback as soon as possible, regardless of other participant’s readiness or suspensions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorRateChangeOptions/playImmediately
	DelegatingPlaybackCoordinatorRateChangeOptionPlayImmediately DelegatingPlaybackCoordinatorRateChangeOptions = 0
)


// DelegatingPlaybackCoordinatorSeekOptions - Constants that define seek options.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekOptions
type DelegatingPlaybackCoordinatorSeekOptions uint

const (
	// DelegatingPlaybackCoordinatorSeekOptionResumeImmediately - An option that Indicates that the coordinator needs to resume playback as soon as possible, regardless of other participant’s readiness or suspensions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekOptions/resumeImmediately
	DelegatingPlaybackCoordinatorSeekOptionResumeImmediately DelegatingPlaybackCoordinatorSeekOptions = 0
)


// DepthDataAccuracy - Values indicating the general accuracy of a depth data map.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/Accuracy
type DepthDataAccuracy uint

const (
	// DepthDataAccuracyAbsolute - Values within the depth map are absolutely accurate within the physical world.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/Accuracy/absolute
	DepthDataAccuracyAbsolute DepthDataAccuracy = 0
	// DepthDataAccuracyRelative - Values within the depth data map are usable for foreground/background separation, but are not absolutely accurate in the physical world.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/Accuracy/relative
	DepthDataAccuracyRelative DepthDataAccuracy = 0
)


// DepthDataQuality - Values indicating the overall quality of a depth data map.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/Quality
type DepthDataQuality uint

const (
	// DepthDataQualityHigh - The depth map is a good candidate for rendering high-quality depth effects or reconstructing a 3D scene.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/Quality/high
	DepthDataQualityHigh DepthDataQuality = 0
	// DepthDataQualityLow - The depth map is a poor candidate for rendering high-quality depth effects or reconstructing a 3D scene.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/Quality/low
	DepthDataQualityLow DepthDataQuality = 0
)


// ExternalContentProtectionStatus - Constants that specify whether sufficient protection exists to display the content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalContentProtectionStatus
type ExternalContentProtectionStatus uint

const (
	// ExternalContentProtectionStatusInsufficient - A status that indicates insufficient protections exists for display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalContentProtectionStatus/insufficient
	ExternalContentProtectionStatusInsufficient ExternalContentProtectionStatus = 0
	// ExternalContentProtectionStatusPending - A status that indicates content protections are pending.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalContentProtectionStatus/pending
	ExternalContentProtectionStatusPending ExternalContentProtectionStatus = 0
	// ExternalContentProtectionStatusSufficient - A status that indicates sufficient protections exists for display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalContentProtectionStatus/sufficient
	ExternalContentProtectionStatusSufficient ExternalContentProtectionStatus = 0
)


// ExternalSyncDeviceStatus - Connection state of an external sync device
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDeviceStatus
type ExternalSyncDeviceStatus uint

const (
	// ExternalSyncDeviceStatusActiveSync - Indicates that the   object is running and that the clock property on   is calibrated to the external sync signal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDeviceStatus/activeSync
	ExternalSyncDeviceStatusActiveSync ExternalSyncDeviceStatus = 0
	// ExternalSyncDeviceStatusCalibrating - Indicates that the external sync signal is connected and that the AVExternalSyncDevice object is calibrating to follow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDeviceStatus/calibrating
	ExternalSyncDeviceStatusCalibrating ExternalSyncDeviceStatus = 0
	// ExternalSyncDeviceStatusFreeRunSync - Indicates that the AVExternalSyncDevice was calibrated to follow the external sync, but the sync signal has been lost. The camera will continue to match the last signal it received, but sync is not guaranteed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDeviceStatus/freeRunSync
	ExternalSyncDeviceStatusFreeRunSync ExternalSyncDeviceStatus = 0
	// ExternalSyncDeviceStatusReady - Indicates that a device supporting external sync is connected, but calibration has not started.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDeviceStatus/ready
	ExternalSyncDeviceStatusReady ExternalSyncDeviceStatus = 0
	// ExternalSyncDeviceStatusUnavailable - Indicates that external sync signal is not connected, or has transitioned to a state that is not recoverable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDeviceStatus/unavailable
	ExternalSyncDeviceStatusUnavailable ExternalSyncDeviceStatus = 0
)


// KeyValueStatus - Values that indicate the loaded status of a property.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVKeyValueStatus
type KeyValueStatus uint

const (
	// KeyValueStatusCancelled - You canceled loading a property value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVKeyValueStatus/cancelled
	KeyValueStatusCancelled KeyValueStatus = 0
	// KeyValueStatusFailed - The system is unable to load the property value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVKeyValueStatus/failed
	KeyValueStatusFailed KeyValueStatus = 0
	// KeyValueStatusLoaded - The property value is ready to use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVKeyValueStatus/loaded
	KeyValueStatusLoaded KeyValueStatus = 0
	// KeyValueStatusLoading - The system is loading the property value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVKeyValueStatus/loading
	KeyValueStatusLoading KeyValueStatus = 0
	// KeyValueStatusUnknown - The property value’s status is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVKeyValueStatus/unknown
	KeyValueStatusUnknown KeyValueStatus = 0
)


// MovieWritingOptions - A structure that defines options to control the writing of a movie header to a destination URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovieWritingOptions
type MovieWritingOptions uint

const (
	// MovieWritingAddMovieHeaderToDestination - The new movie header overwrites any existing movie header.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovieWritingOptions/addMovieHeaderToDestination
	MovieWritingAddMovieHeaderToDestination MovieWritingOptions = 0
	// MovieWritingTruncateDestinationToMovieHeaderOnly - The movie header overwrites all existing data and creates a pure reference movie file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovieWritingOptions/truncateDestinationToMovieHeaderOnly
	MovieWritingTruncateDestinationToMovieHeaderOnly MovieWritingOptions = 0
)


// PlayerActionAtItemEnd - The actions a player can take when it finishes playing.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/ActionAtItemEnd-swift.enum
type PlayerActionAtItemEnd uint

const (
	// PlayerActionAtItemEndAdvance - The player should advance to the next item, if there is one.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/ActionAtItemEnd-swift.enum/advance
	PlayerActionAtItemEndAdvance PlayerActionAtItemEnd = 0
	// PlayerActionAtItemEndNone - The player should do nothing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/ActionAtItemEnd-swift.enum/none
	PlayerActionAtItemEndNone PlayerActionAtItemEnd = 0
	// PlayerActionAtItemEndPause - The player should pause playing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/ActionAtItemEnd-swift.enum/pause
	PlayerActionAtItemEndPause PlayerActionAtItemEnd = 0
)


// PlayerHDRMode - A bitfield type that specifies an HDR mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/HDRMode
type PlayerHDRMode uint

const (
	// PlayerHDRModeDolbyVision - The Dolby Vision HDR mode is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/HDRMode/dolbyVision
	PlayerHDRModeDolbyVision PlayerHDRMode = 0
	// PlayerHDRModeHDR10 - The HDR10 HDR mode is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/HDRMode/hdr10
	PlayerHDRModeHDR10 PlayerHDRMode = 0
	// PlayerHDRModeHLG - The Hybrid Log-Gamma HDR mode is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/HDRMode/hlg
	PlayerHDRModeHLG PlayerHDRMode = 0
)


// PlayerNetworkResourcePriority - This defines the network resource priority for a player.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/NetworkResourcePriority-swift.enum
type PlayerNetworkResourcePriority uint

const (
	// PlayerNetworkResourcePriorityDefault - The default priority level given to a player for loading network resources. Use this when the player requires an optimal level of network resources and streaming in high-quality resolution is ideal. Players with AVPlayerNetworkResourcePriorityHigh will take precedence over this player. This player will take precedence over players with AVPlayerNetworkResourcePriorityLow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/NetworkResourcePriority-swift.enum/default
	PlayerNetworkResourcePriorityDefault PlayerNetworkResourcePriority = 0
	// PlayerNetworkResourcePriorityHigh - Indicates a high priority level for loading network resources. Use this when the player requires a high level of network resources and streaming in high-quality resolution is crucial. This player will take precedence over other lower priority players.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/NetworkResourcePriority-swift.enum/high
	PlayerNetworkResourcePriorityHigh PlayerNetworkResourcePriority = 0
	// PlayerNetworkResourcePriorityLow - Indicates a low priority level for loading network resources. Use this when the player requires minimal network bandwidth and streaming in high-quality resolution is not crucial. Other players with higher priority will take precedence over this player.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/NetworkResourcePriority-swift.enum/low
	PlayerNetworkResourcePriorityLow PlayerNetworkResourcePriority = 0
)


// PlayerStatus - Status values that indicate whether a player can successfully play media.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/Status-swift.enum
type PlayerStatus uint

const (
	// PlayerStatusFailed - A value that indicates the player can no longer play media due to an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/Status-swift.enum/failed
	PlayerStatusFailed PlayerStatus = 0
	// PlayerStatusReadyToPlay - A value that indicates the player is ready to media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/Status-swift.enum/readyToPlay
	PlayerStatusReadyToPlay PlayerStatus = 0
	// PlayerStatusUnknown - A value that indicates a player hasn’t attempted to load media for playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/Status-swift.enum/unknown
	PlayerStatusUnknown PlayerStatus = 0
)


// PlayerTimeControlStatus - Constants that indicate the state of playback control.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum
type PlayerTimeControlStatus uint

const (
	// PlayerTimeControlStatusPaused - A state that indicates the player paused playback indefinitely.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum/paused
	PlayerTimeControlStatusPaused PlayerTimeControlStatus = 0
	// PlayerTimeControlStatusPlaying - A state that indicates that the player is currently playing media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum/playing
	PlayerTimeControlStatusPlaying PlayerTimeControlStatus = 0
	// PlayerTimeControlStatusWaitingToPlayAtSpecifiedRate - A state that indicates that the player is waiting for network conditions to improve before it can start or resume playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum/waitingToPlayAtSpecifiedRate
	PlayerTimeControlStatusWaitingToPlayAtSpecifiedRate PlayerTimeControlStatus = 0
)


// PlayerAudiovisualBackgroundPlaybackPolicy - Policies that describe playback behavior when an app transitions to the background while playing video.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerAudiovisualBackgroundPlaybackPolicy
type PlayerAudiovisualBackgroundPlaybackPolicy uint

const (
	// PlayerAudiovisualBackgroundPlaybackPolicyAutomatic - The system decides whether playback continues.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerAudiovisualBackgroundPlaybackPolicy/automatic
	PlayerAudiovisualBackgroundPlaybackPolicyAutomatic PlayerAudiovisualBackgroundPlaybackPolicy = 0
	// PlayerAudiovisualBackgroundPlaybackPolicyContinuesIfPossible - The app continues playback, if possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerAudiovisualBackgroundPlaybackPolicy/continuesIfPossible
	PlayerAudiovisualBackgroundPlaybackPolicyContinuesIfPossible PlayerAudiovisualBackgroundPlaybackPolicy = 0
	// PlayerAudiovisualBackgroundPlaybackPolicyPauses - The app pauses playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerAudiovisualBackgroundPlaybackPolicy/pauses
	PlayerAudiovisualBackgroundPlaybackPolicyPauses PlayerAudiovisualBackgroundPlaybackPolicy = 0
)


// PlayerInterstitialEventRestrictions - Constants that define restrictions on the playback of interstitial content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/Restrictions-swift.struct
type PlayerInterstitialEventRestrictions uint

const (
	// PlayerInterstitialEventRestrictionDefaultPolicy - The default restriction policy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventRestrictions/AVPlayerInterstitialEventRestrictionDefaultPolicy
	PlayerInterstitialEventRestrictionDefaultPolicy PlayerInterstitialEventRestrictions = 0
	// PlayerInterstitialEventRestrictionNone - A value that indicates no restrictions on playback of primary or interstitial content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventRestrictions/AVPlayerInterstitialEventRestrictionNone
	PlayerInterstitialEventRestrictionNone PlayerInterstitialEventRestrictions = 0
	// PlayerInterstitialEventRestrictionConstrainsSeekingForwardInPrimaryContent - A restriction that indicates the event doesn’t allow seeking forward within an interstitial item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/Restrictions-swift.struct/constrainsSeekingForwardInPrimaryContent
	PlayerInterstitialEventRestrictionConstrainsSeekingForwardInPrimaryContent PlayerInterstitialEventRestrictions = 0
	// PlayerInterstitialEventRestrictionRequiresPlaybackAtPreferredRateForAdvancement - A restriction that indicates the event doesn’t allow advancing the current time within an interstitial item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/Restrictions-swift.struct/requiresPlaybackAtPreferredRateForAdvancement
	PlayerInterstitialEventRestrictionRequiresPlaybackAtPreferredRateForAdvancement PlayerInterstitialEventRestrictions = 0
)


// PlayerInterstitialEventSkippableEventState - These constants describe the state for a skippable AVPlayerInterstitialEvent.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/SkippableEventState
type PlayerInterstitialEventSkippableEventState uint

const (
	// PlayerInterstitialEventSkippableEventStateEligible - Indicates that the interstitial event is currently skippable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/SkippableEventState/eligible
	PlayerInterstitialEventSkippableEventStateEligible PlayerInterstitialEventSkippableEventState = 0
	// PlayerInterstitialEventSkippableEventStateNoLongerEligible - Indicates that the interstitial event is no longer eligible to be skipped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/SkippableEventState/noLongerEligible
	PlayerInterstitialEventSkippableEventStateNoLongerEligible PlayerInterstitialEventSkippableEventState = 0
	// PlayerInterstitialEventSkippableEventStateNotSkippable - Indicates that the interstitial event is not skippable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/SkippableEventState/notSkippable
	PlayerInterstitialEventSkippableEventStateNotSkippable PlayerInterstitialEventSkippableEventState = 0
	// PlayerInterstitialEventSkippableEventStateNotYetEligible - Indicates that the interstitial event will eventually become eligible to be skipped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/SkippableEventState/notYetEligible
	PlayerInterstitialEventSkippableEventStateNotYetEligible PlayerInterstitialEventSkippableEventState = 0
)


// PlayerInterstitialEventTimelineOccupancy - Constants that specify how an event occupies time on an integrated timeline.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/TimelineOccupancy-swift.enum
type PlayerInterstitialEventTimelineOccupancy uint

const (
	// PlayerInterstitialEventTimelineOccupancyFill - The event fills the integrated timeline with the duration of this event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/TimelineOccupancy-swift.enum/fill
	PlayerInterstitialEventTimelineOccupancyFill PlayerInterstitialEventTimelineOccupancy = 0
	// PlayerInterstitialEventTimelineOccupancySinglePoint - The event occupies a single point on the integrated timeline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/TimelineOccupancy-swift.enum/singlePoint
	PlayerInterstitialEventTimelineOccupancySinglePoint PlayerInterstitialEventTimelineOccupancy = 0
)


// PlayerInterstitialEventAssetListResponseStatus - Constants that describe the status of the asset list response for an interstitial event.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventAssetListResponseStatus
type PlayerInterstitialEventAssetListResponseStatus uint

const (
	// PlayerInterstitialEventAssetListResponseStatusAvailable - Indicates that a valid asset list response is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventAssetListResponseStatus/available
	PlayerInterstitialEventAssetListResponseStatusAvailable PlayerInterstitialEventAssetListResponseStatus = 0
	// PlayerInterstitialEventAssetListResponseStatusCleared - Indicates that the system cleared the asset list response.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventAssetListResponseStatus/cleared
	PlayerInterstitialEventAssetListResponseStatusCleared PlayerInterstitialEventAssetListResponseStatus = 0
	// PlayerInterstitialEventAssetListResponseStatusUnavailable - Indicates that the asset list response is unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventAssetListResponseStatus/unavailable
	PlayerInterstitialEventAssetListResponseStatusUnavailable PlayerInterstitialEventAssetListResponseStatus = 0
)


// PlayerItemStatus - The statuses for a player item.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum
type PlayerItemStatus uint

const (
	// PlayerItemStatusFailed - The item no longer plays due to an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/failed
	PlayerItemStatusFailed PlayerItemStatus = 0
	// PlayerItemStatusReadyToPlay - The item is ready to play.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/readyToPlay
	PlayerItemStatusReadyToPlay PlayerItemStatus = 0
	// PlayerItemStatusUnknown - The item’s status is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum/unknown
	PlayerItemStatusUnknown PlayerItemStatus = 0
)


// PlayerItemSegmentType - Constants that specify the type of segment.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/SegmentType-swift.enum
type PlayerItemSegmentType uint

const (
	// PlayerItemSegmentTypeInterstitial - A segment that represents playback of an interstitial event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/SegmentType-swift.enum/interstitial
	PlayerItemSegmentTypeInterstitial PlayerItemSegmentType = 0
	// PlayerItemSegmentTypePrimary - A segment that represents playback of a primary item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/SegmentType-swift.enum/primary
	PlayerItemSegmentTypePrimary PlayerItemSegmentType = 0
)


// PlayerLooperItemOrdering - Constants that define the ordering of items in a player looper.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/ItemOrdering
type PlayerLooperItemOrdering uint

const (
	// PlayerLooperItemOrderingLoopingItemsFollowExistingItems - Indicates to insert replica items after any existing items in the specified player’s queue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/ItemOrdering/loopingItemsFollowExistingItems
	PlayerLooperItemOrderingLoopingItemsFollowExistingItems PlayerLooperItemOrdering = 0
	// PlayerLooperItemOrderingLoopingItemsPrecedeExistingItems - Indicates to insert replica items before any existing items in the specified player’s queue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/ItemOrdering/loopingItemsPrecedeExistingItems
	PlayerLooperItemOrderingLoopingItemsPrecedeExistingItems PlayerLooperItemOrdering = 0
)


// PlayerLooperStatus - Status constants that indicate whether a looper can successfully perform looping playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/Status-swift.enum
type PlayerLooperStatus uint

const (
	// PlayerLooperStatusCancelled - The app canceled looping on the player.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/Status-swift.enum/cancelled
	PlayerLooperStatusCancelled PlayerLooperStatus = 0
	// PlayerLooperStatusFailed - The looper isn’t able to perform looping playback due to an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/Status-swift.enum/failed
	PlayerLooperStatusFailed PlayerLooperStatus = 0
	// PlayerLooperStatusReady - The looper is ready to perform looping playback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/Status-swift.enum/ready
	PlayerLooperStatusReady PlayerLooperStatus = 0
	// PlayerLooperStatusUnknown - The status isn’t known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/Status-swift.enum/unknown
	PlayerLooperStatusUnknown PlayerLooperStatus = 0
)


// QueuedSampleBufferRenderingStatus - The statuses for sample buffer rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRenderingStatus
type QueuedSampleBufferRenderingStatus uint

const (
	// QueuedSampleBufferRenderingStatusFailed - The object can no longer render sample buffers because of an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRenderingStatus/failed
	QueuedSampleBufferRenderingStatusFailed QueuedSampleBufferRenderingStatus = 0
	// QueuedSampleBufferRenderingStatusRendering - The object is rendering the sample buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRenderingStatus/rendering
	QueuedSampleBufferRenderingStatusRendering QueuedSampleBufferRenderingStatus = 0
	// QueuedSampleBufferRenderingStatusUnknown - The object doesn’t have any sample buffers enqueued.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRenderingStatus/unknown
	QueuedSampleBufferRenderingStatusUnknown QueuedSampleBufferRenderingStatus = 0
)


// SampleBufferRequestDirection - The modes that describe the buffer request direction.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/Direction-swift.enum
type SampleBufferRequestDirection uint

const (
	// SampleBufferRequestDirectionForward - The number of following samples may be zero or greater.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/Direction-swift.enum/forward
	SampleBufferRequestDirectionForward SampleBufferRequestDirection = 0
	// SampleBufferRequestDirectionNone - A single sample will be loaded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/Direction-swift.enum/none
	SampleBufferRequestDirectionNone SampleBufferRequestDirection = 0
	// SampleBufferRequestDirectionReverse - The number of previous samples may be zero or greater.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/Direction-swift.enum/reverse
	SampleBufferRequestDirectionReverse SampleBufferRequestDirection = 0
)


// SampleBufferRequestMode - The modes in which a sample buffer generator processes a request.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/Mode-swift.enum
type SampleBufferRequestMode uint

const (
	// SampleBufferRequestModeImmediate - A mode that indicates that sample buffer creation requests load data as soon as possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/Mode-swift.enum/immediate
	SampleBufferRequestModeImmediate SampleBufferRequestMode = 0
	// SampleBufferRequestModeOpportunistic - A mode that indicates that opportunistic sample buffer creation requests load data as soon as possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/Mode-swift.enum/opportunistic
	SampleBufferRequestModeOpportunistic SampleBufferRequestMode = 0
	// SampleBufferRequestModeScheduled - A mode that indicates that sample buffer creation requests load data according to a scheduled deadline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/Mode-swift.enum/scheduled
	SampleBufferRequestModeScheduled SampleBufferRequestMode = 0
)


// VariantPreferences - Defines the preferences the player item uses when selecting variant playlists.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVariantPreferences
type VariantPreferences uint

const (
	// VariantPreferenceNone - Indicates that the player item uses the default behavior for determining variant playlist selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVariantPreferences/AVVariantPreferenceNone
	VariantPreferenceNone VariantPreferences = 0
	// VariantPreferenceScalabilityToLosslessAudio - A preference that indicates the player item supports variant playlists that contain losslessly encoded audio when sufficient bandwidth is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVariantPreferences/scalabilityToLosslessAudio
	VariantPreferenceScalabilityToLosslessAudio VariantPreferences = 0
)


// VideoFieldMode - Constants that indicate which interlacing modes the connection applies to video flowing through it.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoFieldMode
type VideoFieldMode uint

const (
	// VideoFieldModeBoth - A value that indicates that a video connection passes both the top and bottom video fields.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoFieldMode/both
	VideoFieldModeBoth VideoFieldMode = 0
	// VideoFieldModeBottomOnly - A value that indicates that a video connection only passes the bottom video field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoFieldMode/bottomOnly
	VideoFieldModeBottomOnly VideoFieldMode = 0
	// VideoFieldModeDeinterlace - A value that indicates that a video connection deinterlaces the top and bottom video fields.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoFieldMode/deinterlace
	VideoFieldModeDeinterlace VideoFieldMode = 0
	// VideoFieldModeTopOnly - A value that indicates that a video connection only passes the top video field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoFieldMode/topOnly
	VideoFieldModeTopOnly VideoFieldMode = 0
)


