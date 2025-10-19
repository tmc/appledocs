// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAssetDownloadConfiguration] class.
var aVAssetDownloadConfigurationClass = _AVAssetDownloadConfigurationClass{objc.GetClass("AVAssetDownloadConfiguration")}

type _AVAssetDownloadConfigurationClass struct {
	class objc.Class
}

// An object that provides the configuration for a download task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration

type AVAssetDownloadConfiguration struct {
	objectivec.Object
}

// AVAssetDownloadConfigurationFrom constructs a [AVAssetDownloadConfiguration] from an unsafe.Pointer.
//
// An object that provides the configuration for a download task.
func AVAssetDownloadConfigurationFrom(ptr unsafe.Pointer) AVAssetDownloadConfiguration {
	return AVAssetDownloadConfiguration{objectivec.Object{objc.ID(ptr)}}
}



