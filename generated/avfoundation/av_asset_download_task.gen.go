// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [AssetDownloadTask] class.
var (
	AssetDownloadTaskClass     _AssetDownloadTaskClass
	AssetDownloadTaskClassOnce sync.Once
)

func getAssetDownloadTaskClass() _AssetDownloadTaskClass {
	AssetDownloadTaskClassOnce.Do(func() {
		AssetDownloadTaskClass = _AssetDownloadTaskClass{objc.GetClass("AVAssetDownloadTask")}
	})
	return AssetDownloadTaskClass
}

type _AssetDownloadTaskClass struct {
	class objc.Class
}





// An interface definition for the [AssetDownloadTask] class.
type IAssetDownloadTask interface {
	IURLSessionTask
	

	// properties:
	LoadedTimeRanges() []foundation.Value
	Options() foundation.IDictionary
	URLAsset() IAVURLAsset


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetDownloadTaskClass) Alloc() AssetDownloadTask {
	rv := objc.Send[AssetDownloadTask](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetDownloadTaskClass) New() AssetDownloadTask {
	rv := objc.Send[AssetDownloadTask](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetDownloadTask) Init() AssetDownloadTask {
	rv := objc.Send[AssetDownloadTask](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetDownloadTask) Autorelease() AssetDownloadTask {
	rv := objc.Send[AssetDownloadTask](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetDownloadTask creates a new AssetDownloadTask instance.
func NewAssetDownloadTask() AssetDownloadTask {
	return getAssetDownloadTaskClass().New()
}





// A session used to download HTTP Live Streaming assets.
//
// This class is a subclass of that you use to download HTTP Live Streaming assets. You create instances of this class by calling on the download session.


// A session used to download HTTP Live Streaming assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadTask
type AssetDownloadTask struct {
	URLSessionTask
}

// AssetDownloadTaskFrom constructs a [AssetDownloadTask] from an unsafe.Pointer.
//
// A session used to download HTTP Live Streaming assets.
func AssetDownloadTaskFrom(ptr unsafe.Pointer) AssetDownloadTask {
	return AssetDownloadTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}

























// The time ranges of the downloaded media that are ready for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadTask/loadedTimeRanges
func (a_ AssetDownloadTask) LoadedTimeRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](a_.ID, objc.Sel("loadedTimeRanges"))
	return rv
}


// The configuration options for the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadTask/options
func (a_ AssetDownloadTask) Options() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("options"))
	return rv
}


// The asset that this task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadTask/urlAsset
func (a_ AssetDownloadTask) URLAsset() IAVURLAsset {
	rv := objc.Send[URLAsset](a_.ID, objc.Sel("URLAsset"))
	return rv
}







