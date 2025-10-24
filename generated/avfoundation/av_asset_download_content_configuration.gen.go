// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetDownloadContentConfiguration */


/* debug [class_header]: Header for AVAssetDownloadContentConfiguration */
// The class instance for the [AssetDownloadContentConfiguration] class.
var (
	AssetDownloadContentConfigurationClass     _AssetDownloadContentConfigurationClass
	AssetDownloadContentConfigurationClassOnce sync.Once
)

func getAssetDownloadContentConfigurationClass() _AssetDownloadContentConfigurationClass {
	AssetDownloadContentConfigurationClassOnce.Do(func() {
		AssetDownloadContentConfigurationClass = _AssetDownloadContentConfigurationClass{objc.GetClass("AVAssetDownloadContentConfiguration")}
	})
	return AssetDownloadContentConfigurationClass
}

type _AssetDownloadContentConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetDownloadContentConfiguration */
// An interface definition for the [AssetDownloadContentConfiguration] class.
type IAssetDownloadContentConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetDownloadContentConfiguration */
	// properties:
	MediaSelections() []MediaSelection
	SetMediaSelections(value []MediaSelection)
	VariantQualifiers() []AssetVariantQualifier
	SetVariantQualifiers(value []AssetVariantQualifier)
	ArtworkData() foundation.Data
	SetArtworkData(value foundation.Data)
	AuxiliaryContentConfigurations() IAVAssetDownloadContentConfiguration
	SetAuxiliaryContentConfigurations(value IAVAssetDownloadContentConfiguration)
	OptimizesAuxiliaryContentConfigurations() bool
	SetOptimizesAuxiliaryContentConfigurations(value bool)
	PrimaryContentConfiguration() IAVAssetDownloadContentConfiguration
	SetPrimaryContentConfiguration(value IAVAssetDownloadContentConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetDownloadContentConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetDownloadContentConfiguration */
// Alloc allocates a new instance without initialization.
func (ac _AssetDownloadContentConfigurationClass) Alloc() AssetDownloadContentConfiguration {
	rv := objc.Send[AssetDownloadContentConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetDownloadContentConfigurationClass) New() AssetDownloadContentConfiguration {
	rv := objc.Send[AssetDownloadContentConfiguration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetDownloadContentConfiguration) Init() AssetDownloadContentConfiguration {
	rv := objc.Send[AssetDownloadContentConfiguration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetDownloadContentConfiguration) Autorelease() AssetDownloadContentConfiguration {
	rv := objc.Send[AssetDownloadContentConfiguration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetDownloadContentConfiguration creates a new AssetDownloadContentConfiguration instance.
func NewAssetDownloadContentConfiguration() AssetDownloadContentConfiguration {
	return getAssetDownloadContentConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetDownloadContentConfiguration */
// A configuration object that contains variant qualifiers and media options.


// A configuration object that contains variant qualifiers and media options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadContentConfiguration
type AssetDownloadContentConfiguration struct {
	objectivec.Object
}

// AssetDownloadContentConfigurationFrom constructs a [AssetDownloadContentConfiguration] from an unsafe.Pointer.
//
// A configuration object that contains variant qualifiers and media options.
func AssetDownloadContentConfigurationFrom(ptr unsafe.Pointer) AssetDownloadContentConfiguration {
	return AssetDownloadContentConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetDownloadContentConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetDownloadContentConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetDownloadContentConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetDownloadContentConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetDownloadContentConfiguration */

// The media selections of an asset that a task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadContentConfiguration/mediaSelections
func (a_ AssetDownloadContentConfiguration) MediaSelections() []MediaSelection {
	rv := objc.Send[[]MediaSelection](a_.ID, objc.Sel("mediaSelections"))
	return rv
}/* debug [instance_properties/getter]: mediaSelections */


// The media selections of an asset that a task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadContentConfiguration/mediaSelections
func (a_ AssetDownloadContentConfiguration) SetMediaSelections(value []MediaSelection) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaSelections:"), nsArray)
}/* debug [instance_properties/setter]: mediaSelections */


// The variant qualifiers for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadContentConfiguration/variantQualifiers
func (a_ AssetDownloadContentConfiguration) VariantQualifiers() []AssetVariantQualifier {
	rv := objc.Send[[]AssetVariantQualifier](a_.ID, objc.Sel("variantQualifiers"))
	return rv
}/* debug [instance_properties/getter]: variantQualifiers */


// The variant qualifiers for this configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadContentConfiguration/variantQualifiers
func (a_ AssetDownloadContentConfiguration) SetVariantQualifiers(value []AssetVariantQualifier) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setVariantQualifiers:"), nsArray)
}/* debug [instance_properties/setter]: variantQualifiers */


// A data value that represents the asset’s artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/artworkdata
func (a_ AssetDownloadContentConfiguration) ArtworkData() foundation.Data {
	rv := objc.Send[foundation.Data](a_.ID, objc.Sel("artworkData"))
	return rv
}/* debug [instance_properties/getter]: artworkData */


// A data value that represents the asset’s artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/artworkdata
func (a_ AssetDownloadContentConfiguration) SetArtworkData(value foundation.Data) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setArtworkData:"), value)
}/* debug [instance_properties/setter]: artworkData */


// The configuration for the auxiliary content that the task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/auxiliarycontentconfigurations
func (a_ AssetDownloadContentConfiguration) AuxiliaryContentConfigurations() IAVAssetDownloadContentConfiguration {
	rv := objc.Send[AssetDownloadContentConfiguration](a_.ID, objc.Sel("auxiliaryContentConfigurations"))
	return rv
}/* debug [instance_properties/getter]: auxiliaryContentConfigurations */


// The configuration for the auxiliary content that the task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/auxiliarycontentconfigurations
func (a_ AssetDownloadContentConfiguration) SetAuxiliaryContentConfigurations(value IAVAssetDownloadContentConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuxiliaryContentConfigurations:"), value)
}/* debug [instance_properties/setter]: auxiliaryContentConfigurations */


// A Boolean value that indicates whether the task optimizes auxiliary content selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/optimizesauxiliarycontentconfigurations
func (a_ AssetDownloadContentConfiguration) OptimizesAuxiliaryContentConfigurations() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("optimizesAuxiliaryContentConfigurations"))
	return rv
}/* debug [instance_properties/getter]: optimizesAuxiliaryContentConfigurations */


// A Boolean value that indicates whether the task optimizes auxiliary content selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/optimizesauxiliarycontentconfigurations
func (a_ AssetDownloadContentConfiguration) SetOptimizesAuxiliaryContentConfigurations(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOptimizesAuxiliaryContentConfigurations:"), value)
}/* debug [instance_properties/setter]: optimizesAuxiliaryContentConfigurations */


// The configuration for the primary content that the task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/primarycontentconfiguration
func (a_ AssetDownloadContentConfiguration) PrimaryContentConfiguration() IAVAssetDownloadContentConfiguration {
	rv := objc.Send[AssetDownloadContentConfiguration](a_.ID, objc.Sel("primaryContentConfiguration"))
	return rv
}/* debug [instance_properties/getter]: primaryContentConfiguration */


// The configuration for the primary content that the task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadconfiguration/primarycontentconfiguration
func (a_ AssetDownloadContentConfiguration) SetPrimaryContentConfiguration(value IAVAssetDownloadContentConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPrimaryContentConfiguration:"), value)
}/* debug [instance_properties/setter]: primaryContentConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetDownloadContentConfiguration */



