// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AggregateAssetDownloadTask] class.
var (
	AggregateAssetDownloadTaskClass     _AggregateAssetDownloadTaskClass
	AggregateAssetDownloadTaskClassOnce sync.Once
)

func getAggregateAssetDownloadTaskClass() _AggregateAssetDownloadTaskClass {
	AggregateAssetDownloadTaskClassOnce.Do(func() {
		AggregateAssetDownloadTaskClass = _AggregateAssetDownloadTaskClass{objc.GetClass("AVAggregateAssetDownloadTask")}
	})
	return AggregateAssetDownloadTaskClass
}

type _AggregateAssetDownloadTaskClass struct {
	class objc.Class
}

// An interface definition for the [AggregateAssetDownloadTask] class.
type IAggregateAssetDownloadTask interface {
	foundation.IURLSessionTask
	UrlAsset() IAVURLAsset
	SetUrlAsset(value IAVURLAsset)
}

// A task that downloads multiple media selections for an asset.


// A task that downloads multiple media selections for an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAggregateAssetDownloadTask
type AggregateAssetDownloadTask struct {
	foundation.URLSessionTask
}

// AggregateAssetDownloadTaskFrom constructs a [AggregateAssetDownloadTask] from an unsafe.Pointer.
//
// A task that downloads multiple media selections for an asset.
func AggregateAssetDownloadTaskFrom(ptr unsafe.Pointer) AggregateAssetDownloadTask {
	return AggregateAssetDownloadTask{
		URLSessionTask: foundation.URLSessionTaskFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AggregateAssetDownloadTaskClass) Alloc() AggregateAssetDownloadTask {
	rv := objc.Send[AggregateAssetDownloadTask](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AggregateAssetDownloadTaskClass) New() AggregateAssetDownloadTask {
	rv := objc.Send[AggregateAssetDownloadTask](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AggregateAssetDownloadTask) Init() AggregateAssetDownloadTask {
	rv := objc.Send[AggregateAssetDownloadTask](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AggregateAssetDownloadTask) Autorelease() AggregateAssetDownloadTask {
	rv := objc.Send[AggregateAssetDownloadTask](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAggregateAssetDownloadTask creates a new AggregateAssetDownloadTask instance.
func NewAggregateAssetDownloadTask() AggregateAssetDownloadTask {
	return getAggregateAssetDownloadTaskClass().New()
}



// The asset the parent task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaggregateassetdownloadtask/urlasset
func (a_ AggregateAssetDownloadTask) UrlAsset() IAVURLAsset {
	rv := objc.Send[URLAsset](a_.ID, objc.Sel("urlAsset"))
	return rv
}


// The asset the parent task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaggregateassetdownloadtask/urlasset
func (a_ AggregateAssetDownloadTask) SetUrlAsset(value IAVURLAsset) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUrlAsset:"), value)
}



