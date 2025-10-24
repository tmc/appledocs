//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureMovieFileOutput


// A Boolean value that indicates whether the movie file output records video orientation and mirroring information as a metadata track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/recordsVideoOrientationAndMirroringChangesAsMetadataTrack(for:)
func (c_ CaptureMovieFileOutput) RecordsVideoOrientationAndMirroringChangesAsMetadataTrackForConnection(connection IAVCaptureConnection) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("recordsVideoOrientationAndMirroringChangesAsMetadataTrackForConnection:"), connection)
	return rv
}

// Sets whether the movie file output creates a timed metadata track to capture changes to the connection’s video orientation and mirroring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/setRecordsVideoOrientationAndMirroringChangesAsMetadataTrack(_:for:)
func (c_ CaptureMovieFileOutput) SetRecordsVideoOrientationAndMirroringChangesAsMetadataTrackForConnection(doRecordChanges bool, connection IAVCaptureConnection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordsVideoOrientationAndMirroringChanges:asMetadataTrackForConnection:"), doRecordChanges, connection)
}

// Returns a list of supported keys to use in the output settings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/supportedOutputSettingsKeys(for:)
func (c_ CaptureMovieFileOutput) SupportedOutputSettingsKeysForConnection(connection IAVCaptureConnection) []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("supportedOutputSettingsKeysForConnection:"), connection)
	return rv
}

// iOS-only properties

// The video codecs types the output supports for recording movie files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/availableVideoCodecTypes
func (c_ CaptureMovieFileOutput) AvailableVideoCodecTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableVideoCodecTypes"))
	return rv
}




