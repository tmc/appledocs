// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PHASEAssetRegistry] class.
type IPHASEAssetRegistry interface {
	objectivec.IObject
	AssetForIdentifier(identifier string) unsafe.Pointer
	RegisterGlobalMetaParameterError(metaParameterDefinition unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	RegisterSoundAssetWithDataIdentifierFormatNormalizationModeError(data unsafe.Pointer, identifier string, format unsafe.Pointer, normalizationMode unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	RegisterSoundAssetAtURLIdentifierAssetTypeChannelLayoutNormalizationModeError(url unsafe.Pointer, identifier string, assetType unsafe.Pointer, channelLayout unsafe.Pointer, normalizationMode unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	RegisterSoundEventAssetWithRootNodeIdentifierError(rootNode unsafe.Pointer, identifier string, error_ unsafe.Pointer) unsafe.Pointer
	UnregisterAssetWithIdentifierCompletion(identifier string, handler unsafe.Pointer)
}

// A central repository of audio assets.
//
// This class manages audio by registering two types of assets throughout the app’s life cycle: When you’re done with a sound asset, call to free up its system resources.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASEAssetRegistryClass) Alloc() PHASEAssetRegistry {
	rv := objc.Send[PHASEAssetRegistry](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Provides the asset named with the designated identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/asset(forIdentifier:)
func (p_ PHASEAssetRegistry) AssetForIdentifier(identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("assetForIdentifier:"), objc.String(identifier))
	return rv
}

// Registers a global metaparameter with the asset registry.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/registerGlobalMetaParameter(metaParameterDefinition:)
func (p_ PHASEAssetRegistry) RegisterGlobalMetaParameterError(metaParameterDefinition unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("registerGlobalMetaParameter:error:"), metaParameterDefinition, error_)
	return rv
}

// Loads a sound asset from memory and adds it to the engine’s list of registered assets.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/registerSoundAsset(data:identifier:format:normalizationMode:)
func (p_ PHASEAssetRegistry) RegisterSoundAssetWithDataIdentifierFormatNormalizationModeError(data unsafe.Pointer, identifier string, format unsafe.Pointer, normalizationMode unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("registerSoundAssetWithData:identifier:format:normalizationMode:error:"), data, objc.String(identifier), format, normalizationMode, error_)
	return rv
}

// Loads a sound asset from the argument URL and adds it to the engine’s list of registered assets.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/registerSoundAsset(url:identifier:assetType:channelLayout:normalizationMode:)
func (p_ PHASEAssetRegistry) RegisterSoundAssetAtURLIdentifierAssetTypeChannelLayoutNormalizationModeError(url unsafe.Pointer, identifier string, assetType unsafe.Pointer, channelLayout unsafe.Pointer, normalizationMode unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("registerSoundAssetAtURL:identifier:assetType:channelLayout:normalizationMode:error:"), url, objc.String(identifier), assetType, channelLayout, normalizationMode, error_)
	return rv
}

// Registers the root node of the sound event asset.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/registerSoundEventAsset(rootNode:identifier:)
func (p_ PHASEAssetRegistry) RegisterSoundEventAssetWithRootNodeIdentifierError(rootNode unsafe.Pointer, identifier string, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("registerSoundEventAssetWithRootNode:identifier:error:"), rootNode, objc.String(identifier), error_)
	return rv
}

// Deallocates system memory for a given asset and removes it from the engine’s list of registered assets.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/unregisterAsset(identifier:completion:)
func (p_ PHASEAssetRegistry) UnregisterAssetWithIdentifierCompletion(identifier string, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("unregisterAssetWithIdentifier:completion:"), objc.String(identifier), handler)
}

// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAssetRegistry/globalMetaParameters
func (p_ PHASEAssetRegistry) GlobalMetaParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}



