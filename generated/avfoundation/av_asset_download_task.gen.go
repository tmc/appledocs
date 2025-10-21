// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	foundation.IURLSessionTask
}

// A session used to download HTTP Live Streaming assets.
//
// This class is a subclass of that you use to download HTTP Live Streaming assets. You create instances of this class by calling on the download session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadTask
type AssetDownloadTask struct {
	foundation.URLSessionTask
}

// AssetDownloadTaskFrom constructs a [AssetDownloadTask] from an unsafe.Pointer.
//
// A session used to download HTTP Live Streaming assets.
func AssetDownloadTaskFrom(ptr unsafe.Pointer) AssetDownloadTask {
	return AssetDownloadTask{
		URLSessionTask: foundation.URLSessionTaskFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetDownloadTaskClass) Alloc() AssetDownloadTask {
	rv := objc.Send[AssetDownloadTask](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The local file URL to where the task downloads the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask/destinationurl
func (a_ AssetDownloadTask) DestinationURL() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("destinationURL"))
	return rv
}


// SetDestinationURL sets the value of the destinationURL property.
// The local file URL to where the task downloads the asset.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask/destinationurl
func (a_ AssetDownloadTask) SetDestinationURL(value foundation.IURL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDestinationURL:"), value)
}

// The time ranges of the downloaded media that are ready for playback.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask/loadedtimeranges
func (a_ AssetDownloadTask) LoadedTimeRanges() foundation.Value {
	rv := objc.Send[foundation.Value](a_.ID, objc.Sel("loadedTimeRanges"))
	return rv
}


// SetLoadedTimeRanges sets the value of the loadedTimeRanges property.
// The time ranges of the downloaded media that are ready for playback.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask/loadedtimeranges
func (a_ AssetDownloadTask) SetLoadedTimeRanges(value foundation.IValue) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLoadedTimeRanges:"), value)
}

// The configuration options for the task.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask/options
func (a_ AssetDownloadTask) Options() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("options"))
	return rv
}


// SetOptions sets the value of the options property.
// The configuration options for the task.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask/options
func (a_ AssetDownloadTask) SetOptions(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOptions:"), value)
}

// The asset that this task downloads.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask/urlasset
func (a_ AssetDownloadTask) UrlAsset() AVURLAsset {
	rv := objc.Send[AVURLAsset](a_.ID, objc.Sel("urlAsset"))
	return rv
}


// SetUrlAsset sets the value of the urlAsset property.
// The asset that this task downloads.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtask/urlasset
func (a_ AssetDownloadTask) SetUrlAsset(value IAVURLAsset) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUrlAsset:"), value)
}



