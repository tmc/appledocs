// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetDownloadConfiguration] class.
var (
	AssetDownloadConfigurationClass     _AssetDownloadConfigurationClass
	AssetDownloadConfigurationClassOnce sync.Once
)

func getAssetDownloadConfigurationClass() _AssetDownloadConfigurationClass {
	AssetDownloadConfigurationClassOnce.Do(func() {
		AssetDownloadConfigurationClass = _AssetDownloadConfigurationClass{objc.GetClass("AVAssetDownloadConfiguration")}
	})
	return AssetDownloadConfigurationClass
}

type _AssetDownloadConfigurationClass struct {
	class objc.Class
}





// An interface definition for the [AssetDownloadConfiguration] class.
type IAssetDownloadConfiguration interface {
	objectivec.IObject
	

	// properties:
	ArtworkData() foundation.foundation.INSData
	SetArtworkData(value foundation.foundation.INSData)
	AuxiliaryContentConfigurations() []AssetDownloadContentConfiguration
	SetAuxiliaryContentConfigurations(value []AssetDownloadContentConfiguration)
	DownloadsInterstitialAssets() bool
	SetDownloadsInterstitialAssets(value bool)
	OptimizesAuxiliaryContentConfigurations() bool
	SetOptimizesAuxiliaryContentConfigurations(value bool)
	PrimaryContentConfiguration() IAVAssetDownloadContentConfiguration


	

	// methods:
	SetInterstitialMediaSelectionCriteriaForMediaCharacteristic(criteria []PlayerMediaSelectionCriteria, mediaCharacteristic MediaCharacteristic)


}





// Alloc allocates a new instance without initialization.
func (ac _AssetDownloadConfigurationClass) Alloc() AssetDownloadConfiguration {
	rv := objc.Send[AssetDownloadConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetDownloadConfigurationClass) New() AssetDownloadConfiguration {
	rv := objc.Send[AssetDownloadConfiguration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetDownloadConfiguration) Init() AssetDownloadConfiguration {
	rv := objc.Send[AssetDownloadConfiguration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetDownloadConfiguration) Autorelease() AssetDownloadConfiguration {
	rv := objc.Send[AssetDownloadConfiguration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetDownloadConfiguration creates a new AssetDownloadConfiguration instance.
func NewAssetDownloadConfiguration() AssetDownloadConfiguration {
	return getAssetDownloadConfigurationClass().New()
}





// An object that provides the configuration for a download task.


// An object that provides the configuration for a download task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration
type AssetDownloadConfiguration struct {
	objectivec.Object
}

// AssetDownloadConfigurationFrom constructs a [AssetDownloadConfiguration] from an unsafe.Pointer.
//
// An object that provides the configuration for a download task.
func AssetDownloadConfigurationFrom(ptr unsafe.Pointer) AssetDownloadConfiguration {
	return AssetDownloadConfiguration{objectivec.Object{objc.ID(ptr)}}
}






// Creates a download configuration for a media asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/init(asset:title:)
func NewAssetDownloadConfigurationWithAssetTitle(asset IAVURLAsset, title foundation.foundation.INSString) AssetDownloadConfiguration {
	rv := objc.Send[AssetDownloadConfiguration](objc.ID(getAssetDownloadConfigurationClass().class), objc.Sel("downloadConfigurationWithAsset:title:"), asset, title)
	return rv
}







// Creates a download configuration for a media asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/init(asset:title:)
func (ac _AssetDownloadConfigurationClass) DownloadConfigurationWithAssetTitle(asset IAVURLAsset, title foundation.foundation.INSString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("downloadConfigurationWithAsset:title:"), asset, title)
	return rv
}












// Sets media selection on interstitials for this asset
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/setInterstitialMediaSelectionCriteria(_:forMediaCharacteristic:)
func (a_ AssetDownloadConfiguration) SetInterstitialMediaSelectionCriteriaForMediaCharacteristic(criteria []PlayerMediaSelectionCriteria, mediaCharacteristic MediaCharacteristic) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInterstitialMediaSelectionCriteria:forMediaCharacteristic:"), criteria, mediaCharacteristic)
}







// A data value that represents the asset’s artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/artworkData
func (a_ AssetDownloadConfiguration) ArtworkData() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("artworkData"))
	return rv
}


// A data value that represents the asset’s artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/artworkData
func (a_ AssetDownloadConfiguration) SetArtworkData(value foundation.foundation.INSData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setArtworkData:"), value)
}


// The configuration for the auxiliary content that the task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/auxiliaryContentConfigurations
func (a_ AssetDownloadConfiguration) AuxiliaryContentConfigurations() []AssetDownloadContentConfiguration {
	rv := objc.Send[[]AssetDownloadContentConfiguration](a_.ID, objc.Sel("auxiliaryContentConfigurations"))
	return rv
}


// The configuration for the auxiliary content that the task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/auxiliaryContentConfigurations
func (a_ AssetDownloadConfiguration) SetAuxiliaryContentConfigurations(value []AssetDownloadContentConfiguration) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuxiliaryContentConfigurations:"), nsArray)
}


// Download interstitial assets as listed in the index file. False by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/downloadsInterstitialAssets
func (a_ AssetDownloadConfiguration) DownloadsInterstitialAssets() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("downloadsInterstitialAssets"))
	return rv
}


// Download interstitial assets as listed in the index file. False by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/downloadsInterstitialAssets
func (a_ AssetDownloadConfiguration) SetDownloadsInterstitialAssets(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDownloadsInterstitialAssets:"), value)
}


// A Boolean value that indicates whether the task optimizes auxiliary content selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/optimizesAuxiliaryContentConfigurations
func (a_ AssetDownloadConfiguration) OptimizesAuxiliaryContentConfigurations() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("optimizesAuxiliaryContentConfigurations"))
	return rv
}


// A Boolean value that indicates whether the task optimizes auxiliary content selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/optimizesAuxiliaryContentConfigurations
func (a_ AssetDownloadConfiguration) SetOptimizesAuxiliaryContentConfigurations(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOptimizesAuxiliaryContentConfigurations:"), value)
}


// The configuration for the primary content that the task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/primaryContentConfiguration
func (a_ AssetDownloadConfiguration) PrimaryContentConfiguration() IAVAssetDownloadContentConfiguration {
	rv := objc.Send[AssetDownloadContentConfiguration](a_.ID, objc.Sel("primaryContentConfiguration"))
	return rv
}







