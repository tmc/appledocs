// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)





// The class instance for the [AssetDownloadURLSession] class.
var (
	AssetDownloadURLSessionClass     _AssetDownloadURLSessionClass
	AssetDownloadURLSessionClassOnce sync.Once
)

func getAssetDownloadURLSessionClass() _AssetDownloadURLSessionClass {
	AssetDownloadURLSessionClassOnce.Do(func() {
		AssetDownloadURLSessionClass = _AssetDownloadURLSessionClass{objc.GetClass("AVAssetDownloadURLSession")}
	})
	return AssetDownloadURLSessionClass
}

type _AssetDownloadURLSessionClass struct {
	class objc.Class
}





// An interface definition for the [AssetDownloadURLSession] class.
type IAssetDownloadURLSession interface {
	foundation.IURLSession
	

	// properties:
	AVAssetDownloadTaskMediaSelectionKey() objc.IObject /* cross-framework: NSString */
	AVAssetDownloadTaskMediaSelectionPrefersMultichannelKey() objc.IObject /* cross-framework: NSString */
	AVAssetDownloadTaskMinimumRequiredMediaBitrateKey() objc.IObject /* cross-framework: NSString */
	AVAssetDownloadTaskMinimumRequiredPresentationSizeKey() objc.IObject /* cross-framework: NSString */
	AVAssetDownloadTaskPrefersHDRKey() objc.IObject /* cross-framework: NSString */
	AVAssetDownloadTaskPrefersLosslessAudioKey() objc.IObject /* cross-framework: NSString */


	

	// methods:
	AssetDownloadTaskWithConfiguration(downloadConfiguration IAVAssetDownloadConfiguration) IAssetDownloadTask


}





// Alloc allocates a new instance without initialization.
func (ac _AssetDownloadURLSessionClass) Alloc() AssetDownloadURLSession {
	rv := objc.Send[AssetDownloadURLSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetDownloadURLSessionClass) New() AssetDownloadURLSession {
	rv := objc.Send[AssetDownloadURLSession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetDownloadURLSession) Init() AssetDownloadURLSession {
	rv := objc.Send[AssetDownloadURLSession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetDownloadURLSession) Autorelease() AssetDownloadURLSession {
	rv := objc.Send[AssetDownloadURLSession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetDownloadURLSession creates a new AssetDownloadURLSession instance.
func NewAssetDownloadURLSession() AssetDownloadURLSession {
	return getAssetDownloadURLSessionClass().New()
}





// A URL session that creates and executes asset download tasks.


// A URL session that creates and executes asset download tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadURLSession
type AssetDownloadURLSession struct {
	foundation.URLSession
}

// AssetDownloadURLSessionFrom constructs a [AssetDownloadURLSession] from an unsafe.Pointer.
//
// A URL session that creates and executes asset download tasks.
func AssetDownloadURLSessionFrom(ptr unsafe.Pointer) AssetDownloadURLSession {
	return AssetDownloadURLSession{
		URLSession: foundation.URLSessionFrom(ptr),
	}
}






// Creates a URL session to download assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadURLSession/init(configuration:assetDownloadDelegate:delegateQueue:)
func NewAssetDownloadURLSessionWithConfigurationAssetDownloadDelegateDelegateQueue(configuration foundation.URLSessionConfiguration, delegate unsafe.Pointer, delegateQueue foundation.OperationQueue) AssetDownloadURLSession {
	rv := objc.Send[AssetDownloadURLSession](objc.ID(getAssetDownloadURLSessionClass().class), objc.Sel("sessionWithConfiguration:assetDownloadDelegate:delegateQueue:"), configuration, delegate, delegateQueue)
	return rv
}







// Creates a URL session to download assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadURLSession/init(configuration:assetDownloadDelegate:delegateQueue:)
func (ac _AssetDownloadURLSessionClass) SessionWithConfigurationAssetDownloadDelegateDelegateQueue(configuration foundation.URLSessionConfiguration, delegate unsafe.Pointer, delegateQueue foundation.OperationQueue) IAssetDownloadURLSession {
	rv := objc.Send[AssetDownloadURLSession](objc.ID(ac.class), objc.Sel("sessionWithConfiguration:assetDownloadDelegate:delegateQueue:"), configuration, delegate, delegateQueue)
	return rv
}












// Creates a download task that uses the specified configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadURLSession/makeAssetDownloadTask(downloadConfiguration:)
func (a_ AssetDownloadURLSession) AssetDownloadTaskWithConfiguration(downloadConfiguration IAVAssetDownloadConfiguration) IAssetDownloadTask {
	rv := objc.Send[AssetDownloadTask](a_.ID, objc.Sel("assetDownloadTaskWithConfiguration:"), downloadConfiguration)
	return rv
}







// A key that indicates which media selection to download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskmediaselectionkey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskMediaSelectionKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskMediaSelectionKey"))
	return rv
}


// A key that indicates whether the task downloads media selections with support for multichannel playback, when available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskmediaselectionprefersmultichannelkey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskMediaSelectionPrefersMultichannelKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskMediaSelectionPrefersMultichannelKey"))
	return rv
}


// A key that indicates the minimum bit rate of the variant to download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskminimumrequiredmediabitratekey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskMinimumRequiredMediaBitrateKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskMinimumRequiredMediaBitrateKey"))
	return rv
}


// A key that indicates the minimum presentation size of the variant to download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskminimumrequiredpresentationsizekey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskMinimumRequiredPresentationSizeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskMinimumRequiredPresentationSizeKey"))
	return rv
}


// A key that indicates whether the task downloads HDR instead of SDR video, when available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskprefershdrkey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskPrefersHDRKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskPrefersHDRKey"))
	return rv
}


// A key that indicates whether the task downloads media selections in lossless audio format, when available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskpreferslosslessaudiokey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskPrefersLosslessAudioKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskPrefersLosslessAudioKey"))
	return rv
}







