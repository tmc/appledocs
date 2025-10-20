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
func (ac _AssetExportSessionClass) ExportPresetsCompatibleWithAsset(asset unsafe.Pointer) []string {
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
func (a_ AssetExportSession) DirectoryForTemporaryFiles() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("directoryForTemporaryFiles"))
	return rv
}


// SetDirectoryForTemporaryFiles sets the value of the directoryForTemporaryFiles property.
// A directory suitable to store temporary files that the export process generates.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/directoryForTemporaryFiles
func (a_ AssetExportSession) SetDirectoryForTemporaryFiles(value unsafe.Pointer) {
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
func (a_ AssetExportSession) OutputFileType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("outputFileType"))
	return rv
}


// SetOutputFileType sets the value of the outputFileType property.
// The file type of the output an asset export session writes.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputFileType
func (a_ AssetExportSession) SetOutputFileType(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFileType:"), value)
}
// A URL where an asset export session writes its output.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputURL
func (a_ AssetExportSession) OutputURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("outputURL"))
	return rv
}


// SetOutputURL sets the value of the outputURL property.
// A URL where an asset export session writes its output.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/outputURL
func (a_ AssetExportSession) SetOutputURL(value unsafe.Pointer) {
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


