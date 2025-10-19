// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetExportSession] class.
var aVAssetExportSessionClass = _AVAssetExportSessionClass{objc.GetClass("AVAssetExportSession")}

type _AVAssetExportSessionClass struct {
	class objc.Class
}

// An object that exports assets in a format that you specify using an export preset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession

type AVAssetExportSession struct {
	objectivec.Object
}

// AVAssetExportSessionFrom constructs a [AVAssetExportSession] from an unsafe.Pointer.
//
// An object that exports assets in a format that you specify using an export preset.
func AVAssetExportSessionFrom(ptr unsafe.Pointer) AVAssetExportSession {
	return AVAssetExportSession{objectivec.Object{objc.ID(ptr)}}
}

// Returns all available export preset names. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allExportPresets()
func (ac _AVAssetExportSessionClass) AllExportPresets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("allExportPresets"))
	return rv
}
// Returns compatible export presets for the asset. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/exportPresets(compatibleWith:)
func (ac _AVAssetExportSessionClass) ExportPresetsCompatibleWithAsset(asset unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("exportPresetsCompatibleWithAsset:"), asset)
	return rv
}
// Starts the asynchronous execution of an export session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/exportAsynchronously(completionHandler:)
func (a_ AVAssetExportSession) ExportAsynchronouslyWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("exportAsynchronouslyWithCompletionHandler:"), handler)
}


