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
}

// An object that provides the configuration for a download task.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AssetDownloadConfigurationClass) Alloc() AssetDownloadConfiguration {
	rv := objc.Send[AssetDownloadConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A data value that represents the asset’s artwork.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/artworkdata
func (a_ AssetDownloadConfiguration) ArtworkData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("artworkData"))
	return rv
}


// SetArtworkData sets the value of the artworkData property.
// A data value that represents the asset’s artwork.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/artworkdata
func (a_ AssetDownloadConfiguration) SetArtworkData(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setArtworkData:"), value)
}

// The configuration for the auxiliary content that the task downloads.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/auxiliarycontentconfigurations
func (a_ AssetDownloadConfiguration) AuxiliaryContentConfigurations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("auxiliaryContentConfigurations"))
	return rv
}


// SetAuxiliaryContentConfigurations sets the value of the auxiliaryContentConfigurations property.
// The configuration for the auxiliary content that the task downloads.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/auxiliarycontentconfigurations
func (a_ AssetDownloadConfiguration) SetAuxiliaryContentConfigurations(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuxiliaryContentConfigurations:"), value)
}

// A Boolean value that indicates whether the task optimizes auxiliary content selection.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/optimizesauxiliarycontentconfigurations
func (a_ AssetDownloadConfiguration) OptimizesAuxiliaryContentConfigurations() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("optimizesAuxiliaryContentConfigurations"))
	return rv
}


// SetOptimizesAuxiliaryContentConfigurations sets the value of the optimizesAuxiliaryContentConfigurations property.
// A Boolean value that indicates whether the task optimizes auxiliary content selection.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/optimizesauxiliarycontentconfigurations
func (a_ AssetDownloadConfiguration) SetOptimizesAuxiliaryContentConfigurations(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOptimizesAuxiliaryContentConfigurations:"), value)
}

// The configuration for the primary content that the task downloads.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/primarycontentconfiguration
func (a_ AssetDownloadConfiguration) PrimaryContentConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("primaryContentConfiguration"))
	return rv
}


// SetPrimaryContentConfiguration sets the value of the primaryContentConfiguration property.
// The configuration for the primary content that the task downloads.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/primarycontentconfiguration
func (a_ AssetDownloadConfiguration) SetPrimaryContentConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimaryContentConfiguration:"), value)
}



