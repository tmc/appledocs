// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEAssetRegistry */


/* debug [class_header]: Header for PHASEAssetRegistry */
// The class instance for the [PHASEAssetRegistry] class.
var (
	PHASEAssetRegistryClass     _PHASEAssetRegistryClass
	PHASEAssetRegistryClassOnce sync.Once
)

func getPHASEAssetRegistryClass() _PHASEAssetRegistryClass {
	PHASEAssetRegistryClassOnce.Do(func() {
		PHASEAssetRegistryClass = _PHASEAssetRegistryClass{objc.GetClass("PHASEAssetRegistry")}
	})
	return PHASEAssetRegistryClass
}

type _PHASEAssetRegistryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEAssetRegistry */
// An interface definition for the [PHASEAssetRegistry] class.
type IPHASEAssetRegistry interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEAssetRegistry */
	// properties:
	GlobalMetaParameters() foundation.IDictionary
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEAssetRegistry */
	// methods:
	AssetForIdentifier(identifier objc.IObject /* cross-framework: NSString */) IPHASEAsset
	RegisterGlobalMetaParameterError(metaParameterDefinition IPHASEMetaParameterDefinition, error_ unsafe.Pointer) IPHASEGlobalMetaParameterAsset
	RegisterSoundAssetWithDataIdentifierFormatNormalizationModeError(data objc.IObject /* cross-framework: NSData */, identifier objc.IObject /* cross-framework: NSString */, format avfaudio.AudioFormat, normalizationMode PHASENormalizationMode, error_ unsafe.Pointer) IPHASESoundAsset
	RegisterSoundAssetAtURLIdentifierAssetTypeChannelLayoutNormalizationModeError(url objc.IObject /* cross-framework: NSURL */, identifier objc.IObject /* cross-framework: NSString */, assetType PHASEAssetType, channelLayout avfaudio.AudioChannelLayout, normalizationMode PHASENormalizationMode, error_ unsafe.Pointer) IPHASESoundAsset
	RegisterSoundEventAssetWithRootNodeIdentifierError(rootNode IPHASESoundEventNodeDefinition, identifier objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) IPHASESoundEventNodeAsset
	UnregisterAssetWithIdentifierCompletion(identifier objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEAssetRegistry */
// Alloc allocates a new instance without initialization.
func (pc _PHASEAssetRegistryClass) Alloc() PHASEAssetRegistry {
	rv := objc.Send[PHASEAssetRegistry](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEAssetRegistryClass) New() PHASEAssetRegistry {
	rv := objc.Send[PHASEAssetRegistry](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEAssetRegistry) Init() PHASEAssetRegistry {
	rv := objc.Send[PHASEAssetRegistry](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEAssetRegistry) Autorelease() PHASEAssetRegistry {
	rv := objc.Send[PHASEAssetRegistry](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEAssetRegistry creates a new PHASEAssetRegistry instance.
func NewPHASEAssetRegistry() PHASEAssetRegistry {
	return getPHASEAssetRegistryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEAssetRegistry */
// A central repository of audio assets.
//
// This class manages audio by registering two types of assets throughout the app’s life cycle: When you’re done with a sound asset, call to free up its system resources.


// A central repository of audio assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry
type PHASEAssetRegistry struct {
	objectivec.Object
}

// PHASEAssetRegistryFrom constructs a [PHASEAssetRegistry] from an unsafe.Pointer.
//
// A central repository of audio assets.
func PHASEAssetRegistryFrom(ptr unsafe.Pointer) PHASEAssetRegistry {
	return PHASEAssetRegistry{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEAssetRegistry *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEAssetRegistry */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEAssetRegistry */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEAssetRegistry */

// Provides the asset named with the designated identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/asset(forIdentifier:)
func (p_ PHASEAssetRegistry) AssetForIdentifier(identifier objc.IObject /* cross-framework: NSString */) IPHASEAsset {
	rv := objc.Send[PHASEAsset](p_.ID, objc.Sel("assetForIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: AssetForIdentifier */


// Registers a global metaparameter with the asset registry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/registerGlobalMetaParameter(metaParameterDefinition:)
func (p_ PHASEAssetRegistry) RegisterGlobalMetaParameterError(metaParameterDefinition IPHASEMetaParameterDefinition, error_ unsafe.Pointer) IPHASEGlobalMetaParameterAsset {
	rv := objc.Send[PHASEGlobalMetaParameterAsset](p_.ID, objc.Sel("registerGlobalMetaParameter:error:"), metaParameterDefinition, error_)
	return rv
}/* debug [instance_methods/method]: RegisterGlobalMetaParameterError */


// Loads a sound asset from memory and adds it to the engine’s list of registered assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/registerSoundAsset(data:identifier:format:normalizationMode:)
func (p_ PHASEAssetRegistry) RegisterSoundAssetWithDataIdentifierFormatNormalizationModeError(data objc.IObject /* cross-framework: NSData */, identifier objc.IObject /* cross-framework: NSString */, format avfaudio.AudioFormat, normalizationMode PHASENormalizationMode, error_ unsafe.Pointer) IPHASESoundAsset {
	rv := objc.Send[PHASESoundAsset](p_.ID, objc.Sel("registerSoundAssetWithData:identifier:format:normalizationMode:error:"), data, identifier, format, normalizationMode, error_)
	return rv
}/* debug [instance_methods/method]: RegisterSoundAssetWithDataIdentifierFormatNormalizationModeError */


// Loads a sound asset from the argument URL and adds it to the engine’s list of registered assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/registerSoundAsset(url:identifier:assetType:channelLayout:normalizationMode:)
func (p_ PHASEAssetRegistry) RegisterSoundAssetAtURLIdentifierAssetTypeChannelLayoutNormalizationModeError(url objc.IObject /* cross-framework: NSURL */, identifier objc.IObject /* cross-framework: NSString */, assetType PHASEAssetType, channelLayout avfaudio.AudioChannelLayout, normalizationMode PHASENormalizationMode, error_ unsafe.Pointer) IPHASESoundAsset {
	rv := objc.Send[PHASESoundAsset](p_.ID, objc.Sel("registerSoundAssetAtURL:identifier:assetType:channelLayout:normalizationMode:error:"), url, identifier, assetType, channelLayout, normalizationMode, error_)
	return rv
}/* debug [instance_methods/method]: RegisterSoundAssetAtURLIdentifierAssetTypeChannelLayoutNormalizationModeError */


// Registers the root node of the sound event asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/registerSoundEventAsset(rootNode:identifier:)
func (p_ PHASEAssetRegistry) RegisterSoundEventAssetWithRootNodeIdentifierError(rootNode IPHASESoundEventNodeDefinition, identifier objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) IPHASESoundEventNodeAsset {
	rv := objc.Send[PHASESoundEventNodeAsset](p_.ID, objc.Sel("registerSoundEventAssetWithRootNode:identifier:error:"), rootNode, identifier, error_)
	return rv
}/* debug [instance_methods/method]: RegisterSoundEventAssetWithRootNodeIdentifierError */


// Deallocates system memory for a given asset and removes it from the engine’s list of registered assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/unregisterAsset(identifier:completion:)
func (p_ PHASEAssetRegistry) UnregisterAssetWithIdentifierCompletion(identifier objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("unregisterAssetWithIdentifier:completion:"), identifier, handler)
}/* debug [instance_methods/method]: UnregisterAssetWithIdentifierCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEAssetRegistry */

// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/globalMetaParameters
func (p_ PHASEAssetRegistry) GlobalMetaParameters() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}/* debug [instance_properties/getter]: globalMetaParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEAssetRegistry */



