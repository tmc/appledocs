// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetExportSession */


/* debug [class_header]: Header for AVAssetExportSession */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetExportSession */
// An interface definition for the [AssetExportSession] class.
type IAssetExportSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetExportSession */
	// properties:
	AllowsParallelizedExport() bool
	SetAllowsParallelizedExport(value bool)
	Asset() IAVAsset
	AudioMix() IAVAudioMix
	SetAudioMix(value IAVAudioMix)
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* typedef */)
	AudioTrackGroupHandling() AssetTrackGroupOutputHandling
	SetAudioTrackGroupHandling(value AssetTrackGroupOutputHandling)
	CanPerformMultiplePassesOverSourceMediaData() bool
	SetCanPerformMultiplePassesOverSourceMediaData(value bool)
	CustomVideoCompositor() unsafe.Pointer
	DirectoryForTemporaryFiles() objc.IObject /* cross-framework: NSURL */
	SetDirectoryForTemporaryFiles(value objc.IObject /* cross-framework: NSURL */)
	Error() Error
	EstimatedOutputFileLength() objectivec.IObject
	FileLengthLimit() objectivec.IObject
	SetFileLengthLimit(value objectivec.IObject)
	MaxDuration() objc.IObject /* cross-framework: Time */
	Metadata() []MetadataItem
	SetMetadata(value []MetadataItem)
	MetadataItemFilter() IAVMetadataItemFilter
	SetMetadataItemFilter(value IAVMetadataItemFilter)
	OutputFileType() FileType /* typedef */
	SetOutputFileType(value FileType /* typedef */)
	OutputURL() objc.IObject /* cross-framework: NSURL */
	SetOutputURL(value objc.IObject /* cross-framework: NSURL */)
	PresetName() objc.IObject /* cross-framework: NSString */
	Progress() float32
	ShouldOptimizeForNetworkUse() bool
	SetShouldOptimizeForNetworkUse(value bool)
	Status() AssetExportSessionStatus
	SupportedFileTypes() []string
	TimeRange() TimeRange /* not a class type */
	SetTimeRange(value TimeRange /* not a class type */)
	VideoComposition() IAVVideoComposition
	SetVideoComposition(value IAVVideoComposition)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetExportSession */
	// methods:
	DetermineCompatibleFileTypesWithCompletionHandler(handler unsafe.Pointer)
	EstimateMaximumDurationWithCompletionHandler(handler unsafe.Pointer)
	EstimateOutputFileLengthWithCompletionHandler(handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetExportSession */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetExportSession */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetExportSession */

// Creates an export session with a preset configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/init(asset:presetName:)
func NewAssetExportSessionWithAssetPresetName(asset IAVAsset, presetName objc.IObject /* cross-framework: NSString */) AssetExportSession {
	instance := getAssetExportSessionClass().Alloc()
	rv := objc.Send[AssetExportSession](instance.ID, objc.Sel("initWithAsset:presetName:"), asset, presetName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAssetExportSessionWithAssetPresetName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetExportSession */

// Returns all available export preset names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allExportPresets()
func (ac _AssetExportSessionClass) AllExportPresets() []string {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("allExportPresets"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AllExportPresets) */


// Determines an export preset’s compatibility to export the asset in a container of the output file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/determineCompatibility(ofExportPreset:with:outputFileType:completionHandler:)
func (ac _AssetExportSessionClass) DetermineCompatibilityOfExportPresetWithAssetOutputFileTypeCompletionHandler(presetName objc.IObject /* cross-framework: NSString */, asset IAVAsset, outputFileType FileType /* typedef */, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("determineCompatibilityOfExportPreset:withAsset:outputFileType:completionHandler:"), presetName, asset, outputFileType, handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DetermineCompatibilityOfExportPresetWithAssetOutputFileTypeCompletionHandler) */


// Returns a new asset export session that uses the specified preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/exportSessionWithAsset:presetName:
func (ac _AssetExportSessionClass) ExportSessionWithAssetPresetName(asset IAVAsset, presetName objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("exportSessionWithAsset:presetName:"), asset, presetName)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExportSessionWithAssetPresetName) */


// Returns compatible export presets for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/exportPresets(compatibleWith:)
func (ac _AssetExportSessionClass) ExportPresetsCompatibleWithAsset(asset IAVAsset) []string {
	rv := objc.Send[[]string](objc.ID(ac.class), objc.Sel("exportPresetsCompatibleWithAsset:"), asset)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExportPresetsCompatibleWithAsset) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetExportSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetExportSession */

// Determines the output file types an asset export session supports writing in its current configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/determineCompatibleFileTypes(completionHandler:)
func (a_ AssetExportSession) DetermineCompatibleFileTypesWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("determineCompatibleFileTypesWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: DetermineCompatibleFileTypesWithCompletionHandler */


// Starts estimating the maximum duration of the export while considering the asset, preset, and time range configuration of the export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/estimateMaximumDuration(completionHandler:)
func (a_ AssetExportSession) EstimateMaximumDurationWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("estimateMaximumDurationWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: EstimateMaximumDurationWithCompletionHandler */


// Starts estimating the output file length of the export while considering the asset, preset, and time range configuration of the export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/estimateOutputFileLength(completionHandler:)
func (a_ AssetExportSession) EstimateOutputFileLengthWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("estimateOutputFileLengthWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: EstimateOutputFileLengthWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetExportSession */

// A Boolean value that indicates whether the session can parallelize its export operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allowsParallelizedExport
func (a_ AssetExportSession) AllowsParallelizedExport() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsParallelizedExport"))
	return rv
}/* debug [instance_properties/getter]: allowsParallelizedExport */


// A Boolean value that indicates whether the session can parallelize its export operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allowsParallelizedExport
func (a_ AssetExportSession) SetAllowsParallelizedExport(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsParallelizedExport:"), value)
}/* debug [instance_properties/setter]: allowsParallelizedExport */


// An asset that a session exports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/asset
func (a_ AssetExportSession) Asset() IAVAsset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("asset"))
	return rv
}/* debug [instance_properties/getter]: asset */


// The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioMix
func (a_ AssetExportSession) AudioMix() IAVAudioMix {
	rv := objc.Send[AudioMix](a_.ID, objc.Sel("audioMix"))
	return rv
}/* debug [instance_properties/getter]: audioMix */


// The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioMix
func (a_ AssetExportSession) SetAudioMix(value IAVAudioMix) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioMix:"), value)
}/* debug [instance_properties/setter]: audioMix */


// A processing algorithm for managing audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioTimePitchAlgorithm
func (a_ AssetExportSession) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}/* debug [instance_properties/getter]: audioTimePitchAlgorithm */


// A processing algorithm for managing audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioTimePitchAlgorithm
func (a_ AssetExportSession) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}/* debug [instance_properties/setter]: audioTimePitchAlgorithm */


// A policy that defines how the session exports alternate audio tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioTrackGroupHandling
func (a_ AssetExportSession) AudioTrackGroupHandling() AssetTrackGroupOutputHandling {
	rv := objc.Send[AssetTrackGroupOutputHandling](a_.ID, objc.Sel("audioTrackGroupHandling"))
	return rv
}/* debug [instance_properties/getter]: audioTrackGroupHandling */


// A policy that defines how the session exports alternate audio tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioTrackGroupHandling
func (a_ AssetExportSession) SetAudioTrackGroupHandling(value AssetTrackGroupOutputHandling) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTrackGroupHandling:"), value)
}/* debug [instance_properties/setter]: audioTrackGroupHandling */


// A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/canPerformMultiplePassesOverSourceMediaData
func (a_ AssetExportSession) CanPerformMultiplePassesOverSourceMediaData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformMultiplePassesOverSourceMediaData"))
	return rv
}/* debug [instance_properties/getter]: canPerformMultiplePassesOverSourceMediaData */


// A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/canPerformMultiplePassesOverSourceMediaData
func (a_ AssetExportSession) SetCanPerformMultiplePassesOverSourceMediaData(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanPerformMultiplePassesOverSourceMediaData:"), value)
}/* debug [instance_properties/setter]: canPerformMultiplePassesOverSourceMediaData */


// An optional custom object to use when compositing video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/customVideoCompositor
func (a_ AssetExportSession) CustomVideoCompositor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("customVideoCompositor"))
	return rv
}/* debug [instance_properties/getter]: customVideoCompositor */


// A directory suitable to store temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/directoryForTemporaryFiles
func (a_ AssetExportSession) DirectoryForTemporaryFiles() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("directoryForTemporaryFiles"))
	return rv
}/* debug [instance_properties/getter]: directoryForTemporaryFiles */


// A directory suitable to store temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/directoryForTemporaryFiles
func (a_ AssetExportSession) SetDirectoryForTemporaryFiles(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDirectoryForTemporaryFiles:"), value)
}/* debug [instance_properties/setter]: directoryForTemporaryFiles */


// An optional error object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/error
func (a_ AssetExportSession) Error() Error {
	rv := objc.Send[Error](a_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The estimated length of the exported file, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/estimatedOutputFileLength
func (a_ AssetExportSession) EstimatedOutputFileLength() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("estimatedOutputFileLength"))
	return rv
}/* debug [instance_properties/getter]: estimatedOutputFileLength */


// The file length that the output of the session must not exceed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/fileLengthLimit
func (a_ AssetExportSession) FileLengthLimit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("fileLengthLimit"))
	return rv
}/* debug [instance_properties/getter]: fileLengthLimit */


// The file length that the output of the session must not exceed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/fileLengthLimit
func (a_ AssetExportSession) SetFileLengthLimit(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFileLengthLimit:"), value)
}/* debug [instance_properties/setter]: fileLengthLimit */


// Provides an estimate of the maximum duration of the exported media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/maxDuration
func (a_ AssetExportSession) MaxDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("maxDuration"))
	return rv
}/* debug [instance_properties/getter]: maxDuration */


// The metadata an export session writes to the output container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/metadata
func (a_ AssetExportSession) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


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
}/* debug [instance_properties/setter]: metadata */


// An object the export session uses to filter the metadata items it transfers to the output asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/metadataItemFilter
func (a_ AssetExportSession) MetadataItemFilter() IAVMetadataItemFilter {
	rv := objc.Send[MetadataItemFilter](a_.ID, objc.Sel("metadataItemFilter"))
	return rv
}/* debug [instance_properties/getter]: metadataItemFilter */


// An object the export session uses to filter the metadata items it transfers to the output asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/metadataItemFilter
func (a_ AssetExportSession) SetMetadataItemFilter(value IAVMetadataItemFilter) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadataItemFilter:"), value)
}/* debug [instance_properties/setter]: metadataItemFilter */


// The file type of the output an asset export session writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputFileType
func (a_ AssetExportSession) OutputFileType() FileType /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("outputFileType"))
	return rv
}/* debug [instance_properties/getter]: outputFileType */


// The file type of the output an asset export session writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputFileType
func (a_ AssetExportSession) SetOutputFileType(value FileType /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFileType:"), value)
}/* debug [instance_properties/setter]: outputFileType */


// A URL where an asset export session writes its output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputURL
func (a_ AssetExportSession) OutputURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("outputURL"))
	return rv
}/* debug [instance_properties/getter]: outputURL */


// A URL where an asset export session writes its output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputURL
func (a_ AssetExportSession) SetOutputURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputURL:"), value)
}/* debug [instance_properties/setter]: outputURL */


// The name of the preset that the asset export session uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/presetName
func (a_ AssetExportSession) PresetName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("presetName"))
	return rv
}/* debug [instance_properties/getter]: presetName */


// A value that indicates the progress of the export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/progress
func (a_ AssetExportSession) Progress() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("progress"))
	return rv
}/* debug [instance_properties/getter]: progress */


// A Boolean value that indicates whether to optimize the movie for network use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/shouldOptimizeForNetworkUse
func (a_ AssetExportSession) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}/* debug [instance_properties/getter]: shouldOptimizeForNetworkUse */


// A Boolean value that indicates whether to optimize the movie for network use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/shouldOptimizeForNetworkUse
func (a_ AssetExportSession) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}/* debug [instance_properties/setter]: shouldOptimizeForNetworkUse */


// The status of the export session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/status-swift.property
func (a_ AssetExportSession) Status() AssetExportSessionStatus {
	rv := objc.Send[AssetExportSessionStatus](a_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// An array containing the types of files the session can write.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/supportedFileTypes
func (a_ AssetExportSession) SupportedFileTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("supportedFileTypes"))
	return rv
}/* debug [instance_properties/getter]: supportedFileTypes */


// The time range of the source asset to export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/timeRange
func (a_ AssetExportSession) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](a_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */


// The time range of the source asset to export.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/timeRange
func (a_ AssetExportSession) SetTimeRange(value TimeRange /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimeRange:"), value)
}/* debug [instance_properties/setter]: timeRange */


// An optional object that provides instructions for how to composite frames of video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/videoComposition
func (a_ AssetExportSession) VideoComposition() IAVVideoComposition {
	rv := objc.Send[VideoComposition](a_.ID, objc.Sel("videoComposition"))
	return rv
}/* debug [instance_properties/getter]: videoComposition */


// An optional object that provides instructions for how to composite frames of video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/videoComposition
func (a_ AssetExportSession) SetVideoComposition(value IAVVideoComposition) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoComposition:"), value)
}/* debug [instance_properties/setter]: videoComposition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetExportSession */


