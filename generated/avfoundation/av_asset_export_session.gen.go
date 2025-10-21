// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssetExportSession] class.
var (
	AssetExportSessionClass     _AssetExportSessionClass
	AssetExportSessionClassOnce sync.Once
)

func getAssetExportSessionClass() _AssetExportSessionClass {
	AssetExportSessionClassOnce.Do(func() {
		AssetExportSessionClass = _AssetExportSessionClass{objc.GetClass("AVAssetExportSession")}
	})
	return AssetExportSessionClass
}

type _AssetExportSessionClass struct {
	class objc.Class
}

// An interface definition for the [AssetExportSession] class.
type IAssetExportSession interface {
	objectivec.IObject
	CancelExport()
	DetermineCompatibleFileTypesWithCompletionHandler(handler unsafe.Pointer)
	ExportAsynchronouslyWithCompletionHandler(handler unsafe.Pointer)
}

// An object that exports assets in a format that you specify using an export preset.
//
// You configure this object to export an instance of by setting an export preset, an output file type, and an output URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession
type AssetExportSession struct {
	objectivec.Object
}

// AssetExportSessionFrom constructs a [AssetExportSession] from an unsafe.Pointer.
//
// An object that exports assets in a format that you specify using an export preset.
func AssetExportSessionFrom(ptr unsafe.Pointer) AssetExportSession {
	return AssetExportSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetExportSessionClass) Alloc() AssetExportSession {
	rv := objc.Send[AssetExportSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetExportSessionClass) New() AssetExportSession {
	rv := objc.Send[AssetExportSession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetExportSession) Init() AssetExportSession {
	rv := objc.Send[AssetExportSession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetExportSession) Autorelease() AssetExportSession {
	rv := objc.Send[AssetExportSession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetExportSession creates a new AssetExportSession instance.
func NewAssetExportSession() AssetExportSession {
	return getAssetExportSessionClass().New()
}


// Returns all available export preset names.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allExportPresets()
func (ac _AssetExportSessionClass) AllExportPresets() []string {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("allExportPresets"))
	return rv
}

// Returns compatible export presets for the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/exportPresets(compatibleWith:)
func (ac _AssetExportSessionClass) ExportPresetsCompatibleWithAsset(asset IAVAsset) []string {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("exportPresetsCompatibleWithAsset:"), asset)
	return rv
}

// Cancels the execution of an export session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/cancelExport()
func (a_ AssetExportSession) CancelExport() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelExport"))
}

// Determines the output file types an asset export session supports writing in its current configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/determineCompatibleFileTypes(completionHandler:)
func (a_ AssetExportSession) DetermineCompatibleFileTypesWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("determineCompatibleFileTypesWithCompletionHandler:"), handler)
}

// Starts the asynchronous execution of an export session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/exportAsynchronously(completionHandler:)
func (a_ AssetExportSession) ExportAsynchronouslyWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("exportAsynchronouslyWithCompletionHandler:"), handler)
}

// A Boolean value that indicates whether the session can parallelize its export operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allowsParallelizedExport
func (a_ AssetExportSession) AllowsParallelizedExport() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsParallelizedExport"))
	return rv
}


// SetAllowsParallelizedExport sets the value of the allowsParallelizedExport property.
// A Boolean value that indicates whether the session can parallelize its export operation.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allowsParallelizedExport
func (a_ AssetExportSession) SetAllowsParallelizedExport(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsParallelizedExport:"), value)
}

// A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/canPerformMultiplePassesOverSourceMediaData
func (a_ AssetExportSession) CanPerformMultiplePassesOverSourceMediaData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformMultiplePassesOverSourceMediaData"))
	return rv
}


// SetCanPerformMultiplePassesOverSourceMediaData sets the value of the canPerformMultiplePassesOverSourceMediaData property.
// A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/canPerformMultiplePassesOverSourceMediaData
func (a_ AssetExportSession) SetCanPerformMultiplePassesOverSourceMediaData(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanPerformMultiplePassesOverSourceMediaData:"), value)
}

// A directory suitable to store temporary files that the export process generates.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/directoryForTemporaryFiles
func (a_ AssetExportSession) DirectoryForTemporaryFiles() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("directoryForTemporaryFiles"))
	return rv
}


// SetDirectoryForTemporaryFiles sets the value of the directoryForTemporaryFiles property.
// A directory suitable to store temporary files that the export process generates.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/directoryForTemporaryFiles
func (a_ AssetExportSession) SetDirectoryForTemporaryFiles(value foundation.IURL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDirectoryForTemporaryFiles:"), value)
}

// The file length that the output of the session must not exceed.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/fileLengthLimit
func (a_ AssetExportSession) FileLengthLimit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("fileLengthLimit"))
	return rv
}


// SetFileLengthLimit sets the value of the fileLengthLimit property.
// The file length that the output of the session must not exceed.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/fileLengthLimit
func (a_ AssetExportSession) SetFileLengthLimit(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFileLengthLimit:"), value)
}

// The file type of the output an asset export session writes.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputFileType
func (a_ AssetExportSession) OutputFileType() FileType {
	rv := objc.Send[FileType](a_.ID, objc.Sel("outputFileType"))
	return rv
}


// SetOutputFileType sets the value of the outputFileType property.
// The file type of the output an asset export session writes.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputFileType
func (a_ AssetExportSession) SetOutputFileType(value FileType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFileType:"), value)
}

// A URL where an asset export session writes its output.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputURL
func (a_ AssetExportSession) OutputURL() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("outputURL"))
	return rv
}


// SetOutputURL sets the value of the outputURL property.
// A URL where an asset export session writes its output.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputURL
func (a_ AssetExportSession) SetOutputURL(value foundation.IURL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputURL:"), value)
}

// A Boolean value that indicates whether to optimize the movie for network use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/shouldOptimizeForNetworkUse
func (a_ AssetExportSession) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}


// SetShouldOptimizeForNetworkUse sets the value of the shouldOptimizeForNetworkUse property.
// A Boolean value that indicates whether to optimize the movie for network use.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/shouldOptimizeForNetworkUse
func (a_ AssetExportSession) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}

// An array containing the types of files the session can write.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/supportedFileTypes
func (a_ AssetExportSession) SupportedFileTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("supportedFileTypes"))
	return rv
}

// The time range of the source asset to export.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/timeRange
func (a_ AssetExportSession) TimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("timeRange"))
	return rv
}


// SetTimeRange sets the value of the timeRange property.
// The time range of the source asset to export.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/timeRange
func (a_ AssetExportSession) SetTimeRange(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimeRange:"), value)
}

// An asset that a session exports.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/asset
func (a_ AssetExportSession) Asset() AVAsset {
	rv := objc.Send[AVAsset](a_.ID, objc.Sel("asset"))
	return rv
}


// SetAsset sets the value of the asset property.
// An asset that a session exports.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/asset
func (a_ AssetExportSession) SetAsset(value IAVAsset) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAsset:"), value)
}

// The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiomix
func (a_ AssetExportSession) AudioMix() AVAudioMix {
	rv := objc.Send[AVAudioMix](a_.ID, objc.Sel("audioMix"))
	return rv
}


// SetAudioMix sets the value of the audioMix property.
// The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiomix
func (a_ AssetExportSession) SetAudioMix(value IAVAudioMix) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioMix:"), value)
}

// A processing algorithm for managing audio pitch for scaled audio edits.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiotimepitchalgorithm
func (a_ AssetExportSession) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Send[AudioTimePitchAlgorithm](a_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// SetAudioTimePitchAlgorithm sets the value of the audioTimePitchAlgorithm property.
// A processing algorithm for managing audio pitch for scaled audio edits.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiotimepitchalgorithm
func (a_ AssetExportSession) SetAudioTimePitchAlgorithm(value IAudioTimePitchAlgorithm) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}

// A policy that defines how the session exports alternate audio tracks.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiotrackgrouphandling
func (a_ AssetExportSession) AudioTrackGroupHandling() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("audioTrackGroupHandling"))
	return rv
}


// SetAudioTrackGroupHandling sets the value of the audioTrackGroupHandling property.
// A policy that defines how the session exports alternate audio tracks.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiotrackgrouphandling
func (a_ AssetExportSession) SetAudioTrackGroupHandling(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTrackGroupHandling:"), value)
}

// An optional custom object to use when compositing video frames.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/customvideocompositor
func (a_ AssetExportSession) CustomVideoCompositor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("customVideoCompositor"))
	return rv
}


// SetCustomVideoCompositor sets the value of the customVideoCompositor property.
// An optional custom object to use when compositing video frames.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/customvideocompositor
func (a_ AssetExportSession) SetCustomVideoCompositor(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomVideoCompositor:"), value)
}

// An optional error object.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/error
func (a_ AssetExportSession) Error() Error {
	rv := objc.Send[Error](a_.ID, objc.Sel("error"))
	return rv
}


// SetError sets the value of the error property.
// An optional error object.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/error
func (a_ AssetExportSession) SetError(value IError) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setError:"), value)
}

// The estimated length of the exported file, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/estimatedoutputfilelength
func (a_ AssetExportSession) EstimatedOutputFileLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("estimatedOutputFileLength"))
	return rv
}


// SetEstimatedOutputFileLength sets the value of the estimatedOutputFileLength property.
// The estimated length of the exported file, in bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/estimatedoutputfilelength
func (a_ AssetExportSession) SetEstimatedOutputFileLength(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEstimatedOutputFileLength:"), value)
}

// Provides an estimate of the maximum duration of the exported media.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/maxduration
func (a_ AssetExportSession) MaxDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("maxDuration"))
	return rv
}


// SetMaxDuration sets the value of the maxDuration property.
// Provides an estimate of the maximum duration of the exported media.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/maxduration
func (a_ AssetExportSession) SetMaxDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaxDuration:"), value)
}

// The metadata an export session writes to the output container file.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/metadata
func (a_ AssetExportSession) Metadata() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// SetMetadata sets the value of the metadata property.
// The metadata an export session writes to the output container file.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/metadata
func (a_ AssetExportSession) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadata:"), value)
}

// An object the export session uses to filter the metadata items it transfers to the output asset.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/metadataitemfilter
func (a_ AssetExportSession) MetadataItemFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("metadataItemFilter"))
	return rv
}


// SetMetadataItemFilter sets the value of the metadataItemFilter property.
// An object the export session uses to filter the metadata items it transfers to the output asset.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/metadataitemfilter
func (a_ AssetExportSession) SetMetadataItemFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadataItemFilter:"), value)
}

// The name of the preset that the asset export session uses.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/presetname
func (a_ AssetExportSession) PresetName() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("presetName"))
	return rv
}


// SetPresetName sets the value of the presetName property.
// The name of the preset that the asset export session uses.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/presetname
func (a_ AssetExportSession) SetPresetName(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPresetName:"), value)
}

// A value that indicates the progress of the export.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/progress
func (a_ AssetExportSession) Progress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("progress"))
	return rv
}


// SetProgress sets the value of the progress property.
// A value that indicates the progress of the export.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/progress
func (a_ AssetExportSession) SetProgress(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProgress:"), value)
}

// The status of the export session.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/status-swift.property
func (a_ AssetExportSession) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
// The status of the export session.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/status-swift.property
func (a_ AssetExportSession) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStatus:"), value)
}

// An optional object that provides instructions for how to composite frames of video.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/videocomposition
func (a_ AssetExportSession) VideoComposition() AVVideoComposition {
	rv := objc.Send[AVVideoComposition](a_.ID, objc.Sel("videoComposition"))
	return rv
}


// SetVideoComposition sets the value of the videoComposition property.
// An optional object that provides instructions for how to composite frames of video.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/videocomposition
func (a_ AssetExportSession) SetVideoComposition(value IAVVideoComposition) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoComposition:"), value)
}



