// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

// Enum types and constants
// RPCameraPosition - The position of the camera being accessed.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPCameraPosition
type RPCameraPosition uint

// RPPreviewViewControllerMode - The modes used to determine whether the preview view controller or the share screen appears when editing a replay.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPPreviewViewControllerMode
type RPPreviewViewControllerMode uint

// RPRecordingErrorCode - The ReplayKit error domain codes.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode
type RPRecordingErrorCode uint

const (
	// RPRecordingErrorActivePhoneCall - Unable to record due to an active phone call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/activePhoneCall
	RPRecordingErrorActivePhoneCall RPRecordingErrorCode = 0
	// RPRecordingErrorAttemptToStartInRecordingState - Attempted to start a recording that’s already in a recording state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/attemptToStartInRecordingState
	RPRecordingErrorAttemptToStartInRecordingState RPRecordingErrorCode = 0
	// RPRecordingErrorAttemptToStopNonRecording - Attempted to stop a recording that’s not in a recording state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/attemptToStopNonRecording
	RPRecordingErrorAttemptToStopNonRecording RPRecordingErrorCode = 0
	// RPRecordingErrorBroadcastInvalidSession - Attempted to start a broadcast without a prior session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/broadcastInvalidSession
	RPRecordingErrorBroadcastInvalidSession RPRecordingErrorCode = 0
	// RPRecordingErrorBroadcastSetupFailed - The broadcast set up failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/broadcastSetupFailed
	RPRecordingErrorBroadcastSetupFailed RPRecordingErrorCode = 0
	// RPRecordingErrorCarPlay - Failed to start recording because CarPlay is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/carPlay
	RPRecordingErrorCarPlay RPRecordingErrorCode = 0
	// RPRecordingErrorCodeSuccessful - Successfully saved the recording to the Camera Roll.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/codeSuccessful
	RPRecordingErrorCodeSuccessful RPRecordingErrorCode = 0
	// RPRecordingErrorContentResize - Recording interrupted by multitasking and content resizing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/contentResize
	RPRecordingErrorContentResize RPRecordingErrorCode = 0
	// RPRecordingErrorDisabled - Recording disabled via parental controls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/disabled
	RPRecordingErrorDisabled RPRecordingErrorCode = 0
	// RPRecordingErrorEntitlements - Recording failed due to missing entitlements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/entitlements
	RPRecordingErrorEntitlements RPRecordingErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/exportClipToURLInProgress
	RPRecordingErrorExportClipToURLInProgress RPRecordingErrorCode = 0
	// RPRecordingErrorFailed - Recording error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failed
	RPRecordingErrorFailed RPRecordingErrorCode = 0
	// RPRecordingErrorFailedApplicationConnectionInterrupted - The recording failed because the app’s connection was interrupted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedApplicationConnectionInterrupted
	RPRecordingErrorFailedApplicationConnectionInterrupted RPRecordingErrorCode = 0
	// RPRecordingErrorFailedApplicationConnectionInvalid - The recording failed because the app’s connection is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedApplicationConnectionInvalid
	RPRecordingErrorFailedApplicationConnectionInvalid RPRecordingErrorCode = 0
	// RPRecordingErrorFailedAssetWriterExportCanceled - The recording failed because the user canceled the export.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedAssetWriterExportCanceled
	RPRecordingErrorFailedAssetWriterExportCanceled RPRecordingErrorCode = 0
	// RPRecordingErrorFailedAssetWriterExportFailed - The recording failed due to an error exporting the movie.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedAssetWriterExportFailed
	RPRecordingErrorFailedAssetWriterExportFailed RPRecordingErrorCode = 0
	// RPRecordingErrorFailedAssetWriterFailedToSave - The recording failed due to an asset writer failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedAssetWriterFailedToSave
	RPRecordingErrorFailedAssetWriterFailedToSave RPRecordingErrorCode = 0
	// RPRecordingErrorFailedAssetWriterInWrongState - The recording failed because the asset writer is in an invalid state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedAssetWriterInWrongState
	RPRecordingErrorFailedAssetWriterInWrongState RPRecordingErrorCode = 0
	// RPRecordingErrorFailedIncorrectTimeStamps - The recording failed due to malformed start and end time intervals.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedIncorrectTimeStamps
	RPRecordingErrorFailedIncorrectTimeStamps RPRecordingErrorCode = 0
	// RPRecordingErrorFailedMediaServicesFailure - The recording failed due to a mediaservices daemon failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedMediaServicesFailure
	RPRecordingErrorFailedMediaServicesFailure RPRecordingErrorCode = 0
	// RPRecordingErrorFailedNoAssetWriter - The recording failed because there is no asset writer available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedNoAssetWriter
	RPRecordingErrorFailedNoAssetWriter RPRecordingErrorCode = 0
	// RPRecordingErrorFailedNoMatchingApplicationContext - The context identifier doesn’t match the app identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedNoMatchingApplicationContext
	RPRecordingErrorFailedNoMatchingApplicationContext RPRecordingErrorCode = 0
	// RPRecordingErrorFailedToObtainURL - The recording failed due to a failure to obtain the URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedToObtainURL
	RPRecordingErrorFailedToObtainURL RPRecordingErrorCode = 0
	// RPRecordingErrorFailedToProcessFirstSample - The recording failed because the asset writer failed to process the first media sample.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedToProcessFirstSample
	RPRecordingErrorFailedToProcessFirstSample RPRecordingErrorCode = 0
	// RPRecordingErrorFailedToRemoveFile - The recording failed because the temporary file wasn’t removed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedToRemoveFile
	RPRecordingErrorFailedToRemoveFile RPRecordingErrorCode = 0
	// RPRecordingErrorFailedToSave - The recording failed to save.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedToSave
	RPRecordingErrorFailedToSave RPRecordingErrorCode = 0
	// RPRecordingErrorFailedToStart - Recording failed to start.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedToStart
	RPRecordingErrorFailedToStart RPRecordingErrorCode = 0
	// RPRecordingErrorFailedToStartCaptureStack - The system failed to configure the app for A/V recording.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/failedToStartCaptureStack
	RPRecordingErrorFailedToStartCaptureStack RPRecordingErrorCode = 0
	// RPRecordingErrorFilePermissions - The recording failed due to a file permission error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/filePermissions
	RPRecordingErrorFilePermissions RPRecordingErrorCode = 0
	// RPRecordingErrorInsufficientStorage - Not enough storage available on the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/insufficientStorage
	RPRecordingErrorInsufficientStorage RPRecordingErrorCode = 0
	// RPRecordingErrorInterrupted - Recording interrupted by another app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/interrupted
	RPRecordingErrorInterrupted RPRecordingErrorCode = 0
	// RPRecordingErrorInvalidParameter - The recording failed because of an invalid parameter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/invalidParameter
	RPRecordingErrorInvalidParameter RPRecordingErrorCode = 0
	// RPRecordingErrorPhotoFailure - Failed saving the video the Camera Roll.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/photoFailure
	RPRecordingErrorPhotoFailure RPRecordingErrorCode = 0
	// RPRecordingErrorRecordingInvalidSession - Attempted to start an invalid recording session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/recordingInvalidSession
	RPRecordingErrorRecordingInvalidSession RPRecordingErrorCode = 0
	// RPRecordingErrorSystemDormancy - Recording forced to end by the user pressing the power button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/systemDormancy
	RPRecordingErrorSystemDormancy RPRecordingErrorCode = 0
	// RPRecordingErrorUnknown - Error cause unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/unknown
	RPRecordingErrorUnknown RPRecordingErrorCode = 0
	// RPRecordingErrorUserDeclined - User declined recording request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/userDeclined
	RPRecordingErrorUserDeclined RPRecordingErrorCode = 0
	// RPRecordingErrorVideoMixingFailure - The recording failed due to an A/V mixing failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPRecordingErrorCode/videoMixingFailure
	RPRecordingErrorVideoMixingFailure RPRecordingErrorCode = 0
)

// RPSampleBufferType - The type of media clip sample being buffered.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPSampleBufferType
type RPSampleBufferType uint


