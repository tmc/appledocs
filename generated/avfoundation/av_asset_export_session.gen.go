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
	// properties:
	Progress() float32 /* primitive/slice/pointer. */
	AllowsParallelizedExport() bool /* primitive/slice/pointer. */
	SetAllowsParallelizedExport(value bool /* primitive/slice/pointer. */)
	Asset() IAVAsset
	SetAsset(value IAVAsset)
	AudioMix() IAVAudioMix
	SetAudioMix(value IAVAudioMix)
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* not a class type */
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* not a class type */)
	AudioTrackGroupHandling() AssetTrackGroupOutputHandling /* not a class type */
	SetAudioTrackGroupHandling(value AssetTrackGroupOutputHandling /* not a class type */)
	CanPerformMultiplePassesOverSourceMediaData() bool /* primitive/slice/pointer. */
	SetCanPerformMultiplePassesOverSourceMediaData(value bool /* primitive/slice/pointer. */)
	CustomVideoCompositor() VideoCompositing /* not a class type */
	SetCustomVideoCompositor(value VideoCompositing /* not a class type */)
	DirectoryForTemporaryFiles() objc.IObject /* cross-framework: URL */
	SetDirectoryForTemporaryFiles(value objc.IObject /* cross-framework: URL */)
	Error() Error
	SetError(value Error)
	EstimatedOutputFileLength() unsafe.Pointer
	SetEstimatedOutputFileLength(value unsafe.Pointer)
	FileLengthLimit() unsafe.Pointer
	SetFileLengthLimit(value unsafe.Pointer)
	MaxDuration() Time /* not a class type */
	SetMaxDuration(value Time /* not a class type */)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	MetadataItemFilter() MetadataItemFilter /* not a class type */
	SetMetadataItemFilter(value MetadataItemFilter /* not a class type */)
	OutputFileType() objc.IObject /* cross-framework: FileType */
	SetOutputFileType(value objc.IObject /* cross-framework: FileType */)
	OutputURL() objc.IObject /* cross-framework: URL */
	SetOutputURL(value objc.IObject /* cross-framework: URL */)
	PresetName() objc.IObject /* cross-framework: NSString */
	SetPresetName(value objc.IObject /* cross-framework: NSString */)
	ShouldOptimizeForNetworkUse() bool /* primitive/slice/pointer. */
	SetShouldOptimizeForNetworkUse(value bool /* primitive/slice/pointer. */)
	Status() unsafe.Pointer
	SetStatus(value unsafe.Pointer)
	SupportedFileTypes() objc.IObject /* cross-framework: FileType */
	SetSupportedFileTypes(value objc.IObject /* cross-framework: FileType */)
	TimeRange() TimeRange /* not a class type */
	SetTimeRange(value TimeRange /* not a class type */)
	VideoComposition() objc.IObject /* cross-framework: VideoComposition */
	SetVideoComposition(value objc.IObject /* cross-framework: VideoComposition */)
	// methods:
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allExportPresets()
func (ac _AssetExportSessionClass) AllExportPresets() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("allExportPresets"))
	return rv
}


// Returns compatible export presets for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/exportPresets(compatibleWith:)
func (ac _AssetExportSessionClass) ExportPresetsCompatibleWithAsset(asset IAVAsset) []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("exportPresetsCompatibleWithAsset:"), asset)
	return rv
}


// A value that indicates the progress of the export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/progress
func (a_ AssetExportSession) Progress() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("progress"))
	return rv
}


// A Boolean value that indicates whether the session can parallelize its export operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/allowsparallelizedexport
func (a_ AssetExportSession) AllowsParallelizedExport() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsParallelizedExport"))
	return rv
}


// A Boolean value that indicates whether the session can parallelize its export operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/allowsparallelizedexport
func (a_ AssetExportSession) SetAllowsParallelizedExport(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsParallelizedExport:"), value)
}


// An asset that a session exports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/asset
func (a_ AssetExportSession) Asset() IAVAsset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("asset"))
	return rv
}


// An asset that a session exports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/asset
func (a_ AssetExportSession) SetAsset(value IAVAsset) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAsset:"), value)
}


// The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiomix
func (a_ AssetExportSession) AudioMix() IAVAudioMix {
	rv := objc.Send[AudioMix](a_.ID, objc.Sel("audioMix"))
	return rv
}


// The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiomix
func (a_ AssetExportSession) SetAudioMix(value IAVAudioMix) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioMix:"), value)
}


// A processing algorithm for managing audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiotimepitchalgorithm
func (a_ AssetExportSession) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* not a class type */ {
	rv := objc.Send[AudioTimePitchAlgorithm](a_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// A processing algorithm for managing audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiotimepitchalgorithm
func (a_ AssetExportSession) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}


// A policy that defines how the session exports alternate audio tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiotrackgrouphandling
func (a_ AssetExportSession) AudioTrackGroupHandling() AssetTrackGroupOutputHandling /* not a class type */ {
	rv := objc.Send[AssetTrackGroupOutputHandling](a_.ID, objc.Sel("audioTrackGroupHandling"))
	return rv
}


// A policy that defines how the session exports alternate audio tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/audiotrackgrouphandling
func (a_ AssetExportSession) SetAudioTrackGroupHandling(value AssetTrackGroupOutputHandling /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTrackGroupHandling:"), value)
}


// A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/canperformmultiplepassesoversourcemediadata
func (a_ AssetExportSession) CanPerformMultiplePassesOverSourceMediaData() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformMultiplePassesOverSourceMediaData"))
	return rv
}


// A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/canperformmultiplepassesoversourcemediadata
func (a_ AssetExportSession) SetCanPerformMultiplePassesOverSourceMediaData(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanPerformMultiplePassesOverSourceMediaData:"), value)
}


// An optional custom object to use when compositing video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/customvideocompositor
func (a_ AssetExportSession) CustomVideoCompositor() VideoCompositing /* not a class type */ {
	rv := objc.Send[VideoCompositing](a_.ID, objc.Sel("customVideoCompositor"))
	return rv
}


// An optional custom object to use when compositing video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/customvideocompositor
func (a_ AssetExportSession) SetCustomVideoCompositor(value VideoCompositing /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomVideoCompositor:"), value)
}


// A directory suitable to store temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/directoryfortemporaryfiles
func (a_ AssetExportSession) DirectoryForTemporaryFiles() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("directoryForTemporaryFiles"))
	return rv
}


// A directory suitable to store temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/directoryfortemporaryfiles
func (a_ AssetExportSession) SetDirectoryForTemporaryFiles(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDirectoryForTemporaryFiles:"), value)
}


// An optional error object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/error
func (a_ AssetExportSession) Error() Error {
	rv := objc.Send[Error](a_.ID, objc.Sel("error"))
	return rv
}


// An optional error object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/error
func (a_ AssetExportSession) SetError(value Error) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setError:"), value)
}


// The estimated length of the exported file, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/estimatedoutputfilelength
func (a_ AssetExportSession) EstimatedOutputFileLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("estimatedOutputFileLength"))
	return rv
}


// The estimated length of the exported file, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/estimatedoutputfilelength
func (a_ AssetExportSession) SetEstimatedOutputFileLength(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEstimatedOutputFileLength:"), value)
}


// The file length that the output of the session must not exceed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/filelengthlimit
func (a_ AssetExportSession) FileLengthLimit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("fileLengthLimit"))
	return rv
}


// The file length that the output of the session must not exceed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/filelengthlimit
func (a_ AssetExportSession) SetFileLengthLimit(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFileLengthLimit:"), value)
}


// Provides an estimate of the maximum duration of the exported media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/maxduration
func (a_ AssetExportSession) MaxDuration() Time /* not a class type */ {
	rv := objc.Send[Time](a_.ID, objc.Sel("maxDuration"))
	return rv
}


// Provides an estimate of the maximum duration of the exported media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/maxduration
func (a_ AssetExportSession) SetMaxDuration(value Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaxDuration:"), value)
}


// The metadata an export session writes to the output container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/metadata
func (a_ AssetExportSession) Metadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// The metadata an export session writes to the output container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/metadata
func (a_ AssetExportSession) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadata:"), value)
}


// An object the export session uses to filter the metadata items it transfers to the output asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/metadataitemfilter
func (a_ AssetExportSession) MetadataItemFilter() MetadataItemFilter /* not a class type */ {
	rv := objc.Send[MetadataItemFilter](a_.ID, objc.Sel("metadataItemFilter"))
	return rv
}


// An object the export session uses to filter the metadata items it transfers to the output asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/metadataitemfilter
func (a_ AssetExportSession) SetMetadataItemFilter(value MetadataItemFilter /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadataItemFilter:"), value)
}


// The file type of the output an asset export session writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/outputfiletype
func (a_ AssetExportSession) OutputFileType() objc.IObject /* cross-framework: FileType */ {
	rv := objc.Send[FileType](a_.ID, objc.Sel("outputFileType"))
	return rv
}


// The file type of the output an asset export session writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/outputfiletype
func (a_ AssetExportSession) SetOutputFileType(value objc.IObject /* cross-framework: FileType */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFileType:"), value)
}


// A URL where an asset export session writes its output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/outputurl
func (a_ AssetExportSession) OutputURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("outputURL"))
	return rv
}


// A URL where an asset export session writes its output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/outputurl
func (a_ AssetExportSession) SetOutputURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputURL:"), value)
}


// The name of the preset that the asset export session uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/presetname
func (a_ AssetExportSession) PresetName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("presetName"))
	return rv
}


// The name of the preset that the asset export session uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/presetname
func (a_ AssetExportSession) SetPresetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPresetName:"), value)
}


// A Boolean value that indicates whether to optimize the movie for network use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/shouldoptimizefornetworkuse
func (a_ AssetExportSession) ShouldOptimizeForNetworkUse() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}


// A Boolean value that indicates whether to optimize the movie for network use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/shouldoptimizefornetworkuse
func (a_ AssetExportSession) SetShouldOptimizeForNetworkUse(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}


// The status of the export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/status-swift.property
func (a_ AssetExportSession) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("status"))
	return rv
}


// The status of the export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/status-swift.property
func (a_ AssetExportSession) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStatus:"), value)
}


// An array containing the types of files the session can write.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/supportedfiletypes
func (a_ AssetExportSession) SupportedFileTypes() objc.IObject /* cross-framework: FileType */ {
	rv := objc.Send[FileType](a_.ID, objc.Sel("supportedFileTypes"))
	return rv
}


// An array containing the types of files the session can write.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/supportedfiletypes
func (a_ AssetExportSession) SetSupportedFileTypes(value objc.IObject /* cross-framework: FileType */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSupportedFileTypes:"), value)
}


// The time range of the source asset to export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/timerange
func (a_ AssetExportSession) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](a_.ID, objc.Sel("timeRange"))
	return rv
}


// The time range of the source asset to export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/timerange
func (a_ AssetExportSession) SetTimeRange(value TimeRange /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimeRange:"), value)
}


// An optional object that provides instructions for how to composite frames of video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/videocomposition
func (a_ AssetExportSession) VideoComposition() objc.IObject /* cross-framework: VideoComposition */ {
	rv := objc.Send[VideoComposition](a_.ID, objc.Sel("videoComposition"))
	return rv
}


// An optional object that provides instructions for how to composite frames of video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetexportsession/videocomposition
func (a_ AssetExportSession) SetVideoComposition(value objc.IObject /* cross-framework: VideoComposition */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoComposition:"), value)
}



