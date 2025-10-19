// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AVAssetDownloadTask] class.
var (
	aVAssetDownloadTaskClass     _AVAssetDownloadTaskClass
	aVAssetDownloadTaskClassOnce sync.Once
)

func getAVAssetDownloadTaskClass() _AVAssetDownloadTaskClass {
	aVAssetDownloadTaskClassOnce.Do(func() {
		aVAssetDownloadTaskClass = _AVAssetDownloadTaskClass{objc.GetClass("AVAssetDownloadTask")}
	})
	return aVAssetDownloadTaskClass
}

type _AVAssetDownloadTaskClass struct {
	class objc.Class
}

// An interface definition for the [AVAssetDownloadTask] class.
type IAVAssetDownloadTask interface {
	foundation.IURLSessionTask
}

// A session used to download HTTP Live Streaming assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadTask
type AVAssetDownloadTask struct {
	foundation.URLSessionTask
}

// AVAssetDownloadTaskFrom constructs a [AVAssetDownloadTask] from an unsafe.Pointer.
//
// A session used to download HTTP Live Streaming assets.
func AVAssetDownloadTaskFrom(ptr unsafe.Pointer) AVAssetDownloadTask {
	return AVAssetDownloadTask{
		URLSessionTask: foundation.URLSessionTaskFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVAssetDownloadTaskClass) Alloc() AVAssetDownloadTask {
	rv := objc.Send[AVAssetDownloadTask](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVAssetDownloadTaskClass) New() AVAssetDownloadTask {
	rv := objc.Send[AVAssetDownloadTask](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAssetDownloadTask) Init() AVAssetDownloadTask {
	rv := objc.Send[AVAssetDownloadTask](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAssetDownloadTask) Autorelease() AVAssetDownloadTask {
	rv := objc.Send[AVAssetDownloadTask](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetDownloadTask creates a new AVAssetDownloadTask instance.
func NewAVAssetDownloadTask() AVAssetDownloadTask {
	return getAVAssetDownloadTaskClass().New()
}




