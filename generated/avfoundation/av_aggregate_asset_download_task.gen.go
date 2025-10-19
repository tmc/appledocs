// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVAggregateAssetDownloadTask] class.
var aVAggregateAssetDownloadTaskClass = _AVAggregateAssetDownloadTaskClass{objc.GetClass("AVAggregateAssetDownloadTask")}

type _AVAggregateAssetDownloadTaskClass struct {
	class objc.Class
}

// A task that downloads multiple media selections for an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAggregateAssetDownloadTask

type AVAggregateAssetDownloadTask struct {
	URLSessionTask
}

// AVAggregateAssetDownloadTaskFrom constructs a [AVAggregateAssetDownloadTask] from an unsafe.Pointer.
//
// A task that downloads multiple media selections for an asset.
func AVAggregateAssetDownloadTaskFrom(ptr unsafe.Pointer) AVAggregateAssetDownloadTask {
	return AVAggregateAssetDownloadTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}



