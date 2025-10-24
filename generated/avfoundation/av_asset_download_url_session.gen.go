// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AVAssetDownloadURLSession */


/* debug [class_header]: Header for AVAssetDownloadURLSession */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetDownloadURLSession */
// An interface definition for the [AssetDownloadURLSession] class.
type IAssetDownloadURLSession interface {
	foundation.IURLSession
	
/* debug [class_interface_properties]: Properties for AssetDownloadURLSession */
	// properties:
	AVAssetDownloadTaskMediaSelectionKey() objc.IObject /* cross-framework: NSString */
	AVAssetDownloadTaskMediaSelectionPrefersMultichannelKey() objc.IObject /* cross-framework: NSString */
	AVAssetDownloadTaskMinimumRequiredMediaBitrateKey() objc.IObject /* cross-framework: NSString */
	AVAssetDownloadTaskMinimumRequiredPresentationSizeKey() objc.IObject /* cross-framework: NSString */
	AVAssetDownloadTaskPrefersHDRKey() objc.IObject /* cross-framework: NSString */
	AVAssetDownloadTaskPrefersLosslessAudioKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetDownloadURLSession */
	// methods:
	AssetDownloadTaskWithConfiguration(downloadConfiguration IAVAssetDownloadConfiguration) IAssetDownloadTask
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetDownloadURLSession */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetDownloadURLSession */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetDownloadURLSession */

// Creates a URL session to download assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadURLSession/init(configuration:assetDownloadDelegate:delegateQueue:)
func NewAssetDownloadURLSessionWithConfigurationAssetDownloadDelegateDelegateQueue(configuration foundation.URLSessionConfiguration, delegate unsafe.Pointer, delegateQueue foundation.OperationQueue) AssetDownloadURLSession {
	rv := objc.Send[AssetDownloadURLSession](objc.ID(getAssetDownloadURLSessionClass().class), objc.Sel("sessionWithConfiguration:assetDownloadDelegate:delegateQueue:"), configuration, delegate, delegateQueue)
	return rv
}/* debug [class_init_methods/constructor]: NewAssetDownloadURLSessionWithConfigurationAssetDownloadDelegateDelegateQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetDownloadURLSession */

// Creates a URL session to download assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadURLSession/init(configuration:assetDownloadDelegate:delegateQueue:)
func (ac _AssetDownloadURLSessionClass) SessionWithConfigurationAssetDownloadDelegateDelegateQueue(configuration foundation.URLSessionConfiguration, delegate unsafe.Pointer, delegateQueue foundation.OperationQueue) IAssetDownloadURLSession {
	rv := objc.Send[AssetDownloadURLSession](objc.ID(ac.class), objc.Sel("sessionWithConfiguration:assetDownloadDelegate:delegateQueue:"), configuration, delegate, delegateQueue)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SessionWithConfigurationAssetDownloadDelegateDelegateQueue) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetDownloadURLSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetDownloadURLSession */

// Creates a download task that uses the specified configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadURLSession/makeAssetDownloadTask(downloadConfiguration:)
func (a_ AssetDownloadURLSession) AssetDownloadTaskWithConfiguration(downloadConfiguration IAVAssetDownloadConfiguration) IAssetDownloadTask {
	rv := objc.Send[AssetDownloadTask](a_.ID, objc.Sel("assetDownloadTaskWithConfiguration:"), downloadConfiguration)
	return rv
}/* debug [instance_methods/method]: AssetDownloadTaskWithConfiguration */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetDownloadURLSession */

// A key that indicates which media selection to download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskmediaselectionkey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskMediaSelectionKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskMediaSelectionKey"))
	return rv
}/* debug [instance_properties/getter]: AVAssetDownloadTaskMediaSelectionKey */


// A key that indicates whether the task downloads media selections with support for multichannel playback, when available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskmediaselectionprefersmultichannelkey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskMediaSelectionPrefersMultichannelKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskMediaSelectionPrefersMultichannelKey"))
	return rv
}/* debug [instance_properties/getter]: AVAssetDownloadTaskMediaSelectionPrefersMultichannelKey */


// A key that indicates the minimum bit rate of the variant to download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskminimumrequiredmediabitratekey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskMinimumRequiredMediaBitrateKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskMinimumRequiredMediaBitrateKey"))
	return rv
}/* debug [instance_properties/getter]: AVAssetDownloadTaskMinimumRequiredMediaBitrateKey */


// A key that indicates the minimum presentation size of the variant to download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskminimumrequiredpresentationsizekey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskMinimumRequiredPresentationSizeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskMinimumRequiredPresentationSizeKey"))
	return rv
}/* debug [instance_properties/getter]: AVAssetDownloadTaskMinimumRequiredPresentationSizeKey */


// A key that indicates whether the task downloads HDR instead of SDR video, when available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskprefershdrkey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskPrefersHDRKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskPrefersHDRKey"))
	return rv
}/* debug [instance_properties/getter]: AVAssetDownloadTaskPrefersHDRKey */


// A key that indicates whether the task downloads media selections in lossless audio format, when available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadtaskpreferslosslessaudiokey
func (a_ AssetDownloadURLSession) AVAssetDownloadTaskPrefersLosslessAudioKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVAssetDownloadTaskPrefersLosslessAudioKey"))
	return rv
}/* debug [instance_properties/getter]: AVAssetDownloadTaskPrefersLosslessAudioKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetDownloadURLSession */


