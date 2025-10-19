// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVAssetDownloadTask] class.
var aVAssetDownloadTaskClass = _AVAssetDownloadTaskClass{objc.GetClass("AVAssetDownloadTask")}

type _AVAssetDownloadTaskClass struct {
	class objc.Class
}

// A session used to download HTTP Live Streaming assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadTask

type AVAssetDownloadTask struct {
	URLSessionTask
}

// AVAssetDownloadTaskFrom constructs a [AVAssetDownloadTask] from an unsafe.Pointer.
//
// A session used to download HTTP Live Streaming assets.
func AVAssetDownloadTaskFrom(ptr unsafe.Pointer) AVAssetDownloadTask {
	return AVAssetDownloadTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}



