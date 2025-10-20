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




