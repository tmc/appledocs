// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AVAggregateAssetDownloadTask] class.
var (
	aVAggregateAssetDownloadTaskClass     _AVAggregateAssetDownloadTaskClass
	aVAggregateAssetDownloadTaskClassOnce sync.Once
)

func getAVAggregateAssetDownloadTaskClass() _AVAggregateAssetDownloadTaskClass {
	aVAggregateAssetDownloadTaskClassOnce.Do(func() {
		aVAggregateAssetDownloadTaskClass = _AVAggregateAssetDownloadTaskClass{objc.GetClass("AVAggregateAssetDownloadTask")}
	})
	return aVAggregateAssetDownloadTaskClass
}

type _AVAggregateAssetDownloadTaskClass struct {
	class objc.Class
}

// An interface definition for the [AVAggregateAssetDownloadTask] class.
type IAVAggregateAssetDownloadTask interface {
	foundation.IURLSessionTask
}

// A task that downloads multiple media selections for an asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAggregateAssetDownloadTask
type AVAggregateAssetDownloadTask struct {
	foundation.URLSessionTask
}

// AVAggregateAssetDownloadTaskFrom constructs a [AVAggregateAssetDownloadTask] from an unsafe.Pointer.
//
// A task that downloads multiple media selections for an asset.
func AVAggregateAssetDownloadTaskFrom(ptr unsafe.Pointer) AVAggregateAssetDownloadTask {
	return AVAggregateAssetDownloadTask{
		URLSessionTask: foundation.URLSessionTaskFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AVAggregateAssetDownloadTaskClass) Alloc() AVAggregateAssetDownloadTask {
	rv := objc.Send[AVAggregateAssetDownloadTask](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVAggregateAssetDownloadTaskClass) New() AVAggregateAssetDownloadTask {
	rv := objc.Send[AVAggregateAssetDownloadTask](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAggregateAssetDownloadTask) Init() AVAggregateAssetDownloadTask {
	rv := objc.Send[AVAggregateAssetDownloadTask](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAggregateAssetDownloadTask) Autorelease() AVAggregateAssetDownloadTask {
	rv := objc.Send[AVAggregateAssetDownloadTask](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAggregateAssetDownloadTask creates a new AVAggregateAssetDownloadTask instance.
func NewAVAggregateAssetDownloadTask() AVAggregateAssetDownloadTask {
	return getAVAggregateAssetDownloadTaskClass().New()
}




