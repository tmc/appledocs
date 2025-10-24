// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetDownloadConfiguration */


/* debug [class_header]: Header for AVAssetDownloadConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetDownloadConfiguration */
// An interface definition for the [AssetDownloadConfiguration] class.
type IAssetDownloadConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetDownloadConfiguration */
	// properties:
	ArtworkData() objc.IObject /* cross-framework: NSData */
	SetArtworkData(value objc.IObject /* cross-framework: NSData */)
	AuxiliaryContentConfigurations() []AssetDownloadContentConfiguration
	SetAuxiliaryContentConfigurations(value []AssetDownloadContentConfiguration)
	DownloadsInterstitialAssets() bool
	SetDownloadsInterstitialAssets(value bool)
	OptimizesAuxiliaryContentConfigurations() bool
	SetOptimizesAuxiliaryContentConfigurations(value bool)
	PrimaryContentConfiguration() IAVAssetDownloadContentConfiguration
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetDownloadConfiguration */
	// methods:
	SetInterstitialMediaSelectionCriteriaForMediaCharacteristic(criteria []PlayerMediaSelectionCriteria, mediaCharacteristic MediaCharacteristic /* typedef */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetDownloadConfiguration */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetDownloadConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetDownloadConfiguration */

// Creates a download configuration for a media asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/init(asset:title:)
func NewAssetDownloadConfigurationWithAssetTitle(asset IAVURLAsset, title objc.IObject /* cross-framework: NSString */) AssetDownloadConfiguration {
	rv := objc.Send[AssetDownloadConfiguration](objc.ID(getAssetDownloadConfigurationClass().class), objc.Sel("downloadConfigurationWithAsset:title:"), asset, title)
	return rv
}/* debug [class_init_methods/constructor]: NewAssetDownloadConfigurationWithAssetTitle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetDownloadConfiguration */

// Creates a download configuration for a media asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/init(asset:title:)
func (ac _AssetDownloadConfigurationClass) DownloadConfigurationWithAssetTitle(asset IAVURLAsset, title objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("downloadConfigurationWithAsset:title:"), asset, title)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DownloadConfigurationWithAssetTitle) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetDownloadConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetDownloadConfiguration */

// Sets media selection on interstitials for this asset
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/setInterstitialMediaSelectionCriteria(_:forMediaCharacteristic:)
func (a_ AssetDownloadConfiguration) SetInterstitialMediaSelectionCriteriaForMediaCharacteristic(criteria []PlayerMediaSelectionCriteria, mediaCharacteristic MediaCharacteristic /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInterstitialMediaSelectionCriteria:forMediaCharacteristic:"), criteria, mediaCharacteristic)
}/* debug [instance_methods/method]: SetInterstitialMediaSelectionCriteriaForMediaCharacteristic */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetDownloadConfiguration */

// A data value that represents the asset’s artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/artworkData
func (a_ AssetDownloadConfiguration) ArtworkData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("artworkData"))
	return rv
}/* debug [instance_properties/getter]: artworkData */


// A data value that represents the asset’s artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/artworkData
func (a_ AssetDownloadConfiguration) SetArtworkData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setArtworkData:"), value)
}/* debug [instance_properties/setter]: artworkData */


// The configuration for the auxiliary content that the task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/auxiliaryContentConfigurations
func (a_ AssetDownloadConfiguration) AuxiliaryContentConfigurations() []AssetDownloadContentConfiguration {
	rv := objc.Send[[]AssetDownloadContentConfiguration](a_.ID, objc.Sel("auxiliaryContentConfigurations"))
	return rv
}/* debug [instance_properties/getter]: auxiliaryContentConfigurations */


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
}/* debug [instance_properties/setter]: auxiliaryContentConfigurations */


// Download interstitial assets as listed in the index file. False by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/downloadsInterstitialAssets
func (a_ AssetDownloadConfiguration) DownloadsInterstitialAssets() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("downloadsInterstitialAssets"))
	return rv
}/* debug [instance_properties/getter]: downloadsInterstitialAssets */


// Download interstitial assets as listed in the index file. False by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/downloadsInterstitialAssets
func (a_ AssetDownloadConfiguration) SetDownloadsInterstitialAssets(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDownloadsInterstitialAssets:"), value)
}/* debug [instance_properties/setter]: downloadsInterstitialAssets */


// A Boolean value that indicates whether the task optimizes auxiliary content selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/optimizesAuxiliaryContentConfigurations
func (a_ AssetDownloadConfiguration) OptimizesAuxiliaryContentConfigurations() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("optimizesAuxiliaryContentConfigurations"))
	return rv
}/* debug [instance_properties/getter]: optimizesAuxiliaryContentConfigurations */


// A Boolean value that indicates whether the task optimizes auxiliary content selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/optimizesAuxiliaryContentConfigurations
func (a_ AssetDownloadConfiguration) SetOptimizesAuxiliaryContentConfigurations(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOptimizesAuxiliaryContentConfigurations:"), value)
}/* debug [instance_properties/setter]: optimizesAuxiliaryContentConfigurations */


// The configuration for the primary content that the task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadConfiguration/primaryContentConfiguration
func (a_ AssetDownloadConfiguration) PrimaryContentConfiguration() IAVAssetDownloadContentConfiguration {
	rv := objc.Send[AssetDownloadContentConfiguration](a_.ID, objc.Sel("primaryContentConfiguration"))
	return rv
}/* debug [instance_properties/getter]: primaryContentConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetDownloadConfiguration */


