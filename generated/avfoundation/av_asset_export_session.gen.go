// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	

	// properties:
	AllowsParallelizedExport() bool
	SetAllowsParallelizedExport(value bool)
	Asset() IAVAsset
	AudioMix() IAVAudioMix
	SetAudioMix(value IAVAudioMix)
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm)
	AudioTrackGroupHandling() AssetTrackGroupOutputHandling
	SetAudioTrackGroupHandling(value AssetTrackGroupOutputHandling)
	CanPerformMultiplePassesOverSourceMediaData() bool
	SetCanPerformMultiplePassesOverSourceMediaData(value bool)
	CustomVideoCompositor() unsafe.Pointer
	DirectoryForTemporaryFiles() foundation.foundation.INSURL
	SetDirectoryForTemporaryFiles(value foundation.foundation.INSURL)
	Error() foundation.foundation.INSError
	EstimatedOutputFileLength() objectivec.IObject
	FileLengthLimit() objectivec.IObject
	SetFileLengthLimit(value objectivec.IObject)
	MaxDuration() objectivec.IObject
	Metadata() []MetadataItem
	SetMetadata(value []MetadataItem)
	MetadataItemFilter() IAVMetadataItemFilter
	SetMetadataItemFilter(value IAVMetadataItemFilter)
	OutputFileType() FileType
	SetOutputFileType(value FileType)
	OutputURL() foundation.foundation.INSURL
	SetOutputURL(value foundation.foundation.INSURL)
	PresetName() foundation.foundation.INSString
	Progress() float32
	ShouldOptimizeForNetworkUse() bool
	SetShouldOptimizeForNetworkUse(value bool)
	Status() AssetExportSessionStatus
	SupportedFileTypes() []string
	TimeRange() objectivec.IObject
	SetTimeRange(value objectivec.IObject)
	VideoComposition() IAVVideoComposition
	SetVideoComposition(value IAVVideoComposition)


	

	// methods:
	DetermineCompatibleFileTypesWithCompletionHandler(handler unsafe.Pointer)
	EstimateMaximumDurationWithCompletionHandler(handler unsafe.Pointer)
	EstimateOutputFileLengthWithCompletionHandler(handler unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (ac _AssetExportSessionClass) Alloc() AssetExportSession {
	rv := objc.Send[AssetExportSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object that exports assets in a format that you specify using an export preset.
//
// You configure this object to export an instance of by setting an export preset, an output file type, and an output URL.


// An object that exports assets in a format that you specify using an export preset.
//
// [Full Topic]
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






// Creates an export session with a preset configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/init(asset:presetName:)
func NewAssetExportSessionWithAssetPresetName(asset IAVAsset, presetName foundation.foundation.INSString) AssetExportSession {
	instance := getAssetExportSessionClass().Alloc()
	rv := objc.Send[AssetExportSession](instance.ID, objc.Sel("initWithAsset:presetName:"), asset, presetName)
	rv.Autorelease()
	return rv
}







// Returns all available export preset names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allExportPresets()
func (ac _AssetExportSessionClass) AllExportPresets() []string {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("allExportPresets"))
	return rv
}


// Determines an export preset’s compatibility to export the asset in a container of the output file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/determineCompatibility(ofExportPreset:with:outputFileType:completionHandler:)
func (ac _AssetExportSessionClass) DetermineCompatibilityOfExportPresetWithAssetOutputFileTypeCompletionHandler(presetName foundation.foundation.INSString, asset IAVAsset, outputFileType FileType, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("determineCompatibilityOfExportPreset:withAsset:outputFileType:completionHandler:"), presetName, asset, outputFileType, handler)
}


// Returns a new asset export session that uses the specified preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/exportSessionWithAsset:presetName:
func (ac _AssetExportSessionClass) ExportSessionWithAssetPresetName(asset IAVAsset, presetName foundation.foundation.INSString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("exportSessionWithAsset:presetName:"), asset, presetName)
	return rv
}


// Returns compatible export presets for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/exportPresets(compatibleWith:)
func (ac _AssetExportSessionClass) ExportPresetsCompatibleWithAsset(asset IAVAsset) []string {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("exportPresetsCompatibleWithAsset:"), asset)
	return rv
}












// Determines the output file types an asset export session supports writing in its current configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/determineCompatibleFileTypes(completionHandler:)
func (a_ AssetExportSession) DetermineCompatibleFileTypesWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("determineCompatibleFileTypesWithCompletionHandler:"), handler)
}


// Starts estimating the maximum duration of the export while considering the asset, preset, and time range configuration of the export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/estimateMaximumDuration(completionHandler:)
func (a_ AssetExportSession) EstimateMaximumDurationWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("estimateMaximumDurationWithCompletionHandler:"), handler)
}


// Starts estimating the output file length of the export while considering the asset, preset, and time range configuration of the export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/estimateOutputFileLength(completionHandler:)
func (a_ AssetExportSession) EstimateOutputFileLengthWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("estimateOutputFileLengthWithCompletionHandler:"), handler)
}







// A Boolean value that indicates whether the session can parallelize its export operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allowsParallelizedExport
func (a_ AssetExportSession) AllowsParallelizedExport() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsParallelizedExport"))
	return rv
}


// A Boolean value that indicates whether the session can parallelize its export operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allowsParallelizedExport
func (a_ AssetExportSession) SetAllowsParallelizedExport(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsParallelizedExport:"), value)
}


// An asset that a session exports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/asset
func (a_ AssetExportSession) Asset() IAVAsset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("asset"))
	return rv
}


// The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioMix
func (a_ AssetExportSession) AudioMix() IAVAudioMix {
	rv := objc.Send[AudioMix](a_.ID, objc.Sel("audioMix"))
	return rv
}


// The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioMix
func (a_ AssetExportSession) SetAudioMix(value IAVAudioMix) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioMix:"), value)
}


// A processing algorithm for managing audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioTimePitchAlgorithm
func (a_ AssetExportSession) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Send[AudioTimePitchAlgorithm](a_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// A processing algorithm for managing audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioTimePitchAlgorithm
func (a_ AssetExportSession) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}


// A policy that defines how the session exports alternate audio tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioTrackGroupHandling
func (a_ AssetExportSession) AudioTrackGroupHandling() AssetTrackGroupOutputHandling {
	rv := objc.Send[AssetTrackGroupOutputHandling](a_.ID, objc.Sel("audioTrackGroupHandling"))
	return rv
}


// A policy that defines how the session exports alternate audio tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioTrackGroupHandling
func (a_ AssetExportSession) SetAudioTrackGroupHandling(value AssetTrackGroupOutputHandling) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTrackGroupHandling:"), value)
}


// A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/canPerformMultiplePassesOverSourceMediaData
func (a_ AssetExportSession) CanPerformMultiplePassesOverSourceMediaData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformMultiplePassesOverSourceMediaData"))
	return rv
}


// A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/canPerformMultiplePassesOverSourceMediaData
func (a_ AssetExportSession) SetCanPerformMultiplePassesOverSourceMediaData(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanPerformMultiplePassesOverSourceMediaData:"), value)
}


// An optional custom object to use when compositing video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/customVideoCompositor
func (a_ AssetExportSession) CustomVideoCompositor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("customVideoCompositor"))
	return rv
}


// A directory suitable to store temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/directoryForTemporaryFiles
func (a_ AssetExportSession) DirectoryForTemporaryFiles() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("directoryForTemporaryFiles"))
	return rv
}


// A directory suitable to store temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/directoryForTemporaryFiles
func (a_ AssetExportSession) SetDirectoryForTemporaryFiles(value foundation.foundation.INSURL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDirectoryForTemporaryFiles:"), value)
}


// An optional error object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/error
func (a_ AssetExportSession) Error() foundation.foundation.INSError {
	rv := objc.Send[foundation.NSError](a_.ID, objc.Sel("error"))
	return rv
}


// The estimated length of the exported file, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/estimatedOutputFileLength
func (a_ AssetExportSession) EstimatedOutputFileLength() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("estimatedOutputFileLength"))
	return rv
}


// The file length that the output of the session must not exceed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/fileLengthLimit
func (a_ AssetExportSession) FileLengthLimit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("fileLengthLimit"))
	return rv
}


// The file length that the output of the session must not exceed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/fileLengthLimit
func (a_ AssetExportSession) SetFileLengthLimit(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFileLengthLimit:"), value)
}


// Provides an estimate of the maximum duration of the exported media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/maxDuration
func (a_ AssetExportSession) MaxDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("maxDuration"))
	return rv
}


// The metadata an export session writes to the output container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/metadata
func (a_ AssetExportSession) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// The metadata an export session writes to the output container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/metadata
func (a_ AssetExportSession) SetMetadata(value []MetadataItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadata:"), nsArray)
}


// An object the export session uses to filter the metadata items it transfers to the output asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/metadataItemFilter
func (a_ AssetExportSession) MetadataItemFilter() IAVMetadataItemFilter {
	rv := objc.Send[MetadataItemFilter](a_.ID, objc.Sel("metadataItemFilter"))
	return rv
}


// An object the export session uses to filter the metadata items it transfers to the output asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/metadataItemFilter
func (a_ AssetExportSession) SetMetadataItemFilter(value IAVMetadataItemFilter) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadataItemFilter:"), value)
}


// The file type of the output an asset export session writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputFileType
func (a_ AssetExportSession) OutputFileType() FileType {
	rv := objc.Send[FileType](a_.ID, objc.Sel("outputFileType"))
	return rv
}


// The file type of the output an asset export session writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputFileType
func (a_ AssetExportSession) SetOutputFileType(value FileType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFileType:"), value)
}


// A URL where an asset export session writes its output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputURL
func (a_ AssetExportSession) OutputURL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("outputURL"))
	return rv
}


// A URL where an asset export session writes its output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputURL
func (a_ AssetExportSession) SetOutputURL(value foundation.foundation.INSURL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputURL:"), value)
}


// The name of the preset that the asset export session uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/presetName
func (a_ AssetExportSession) PresetName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("presetName"))
	return rv
}


// A value that indicates the progress of the export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/progress
func (a_ AssetExportSession) Progress() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("progress"))
	return rv
}


// A Boolean value that indicates whether to optimize the movie for network use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/shouldOptimizeForNetworkUse
func (a_ AssetExportSession) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}


// A Boolean value that indicates whether to optimize the movie for network use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/shouldOptimizeForNetworkUse
func (a_ AssetExportSession) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}


// The status of the export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/status-swift.property
func (a_ AssetExportSession) Status() AssetExportSessionStatus {
	rv := objc.Send[AssetExportSessionStatus](a_.ID, objc.Sel("status"))
	return rv
}


// An array containing the types of files the session can write.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/supportedFileTypes
func (a_ AssetExportSession) SupportedFileTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("supportedFileTypes"))
	return rv
}


// The time range of the source asset to export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/timeRange
func (a_ AssetExportSession) TimeRange() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("timeRange"))
	return rv
}


// The time range of the source asset to export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/timeRange
func (a_ AssetExportSession) SetTimeRange(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimeRange:"), value)
}


// An optional object that provides instructions for how to composite frames of video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/videoComposition
func (a_ AssetExportSession) VideoComposition() IAVVideoComposition {
	rv := objc.Send[VideoComposition](a_.ID, objc.Sel("videoComposition"))
	return rv
}


// An optional object that provides instructions for how to composite frames of video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/videoComposition
func (a_ AssetExportSession) SetVideoComposition(value IAVVideoComposition) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoComposition:"), value)
}







