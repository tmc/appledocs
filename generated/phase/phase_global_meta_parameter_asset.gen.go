// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEGlobalMetaParameterAsset] class.
var (
	PHASEGlobalMetaParameterAssetClass     _PHASEGlobalMetaParameterAssetClass
	PHASEGlobalMetaParameterAssetClassOnce sync.Once
)

func getPHASEGlobalMetaParameterAssetClass() _PHASEGlobalMetaParameterAssetClass {
	PHASEGlobalMetaParameterAssetClassOnce.Do(func() {
		PHASEGlobalMetaParameterAssetClass = _PHASEGlobalMetaParameterAssetClass{objc.GetClass("PHASEGlobalMetaParameterAsset")}
	})
	return PHASEGlobalMetaParameterAssetClass
}

type _PHASEGlobalMetaParameterAssetClass struct {
	class objc.Class
}

// An interface definition for the [PHASEGlobalMetaParameterAsset] class.
type IPHASEGlobalMetaParameterAsset interface {
	IPHASEAsset
}

// A reference to a registered metaparameter that the app can share with multiple sound events or sources.
//
// The engine’s function returns an instance of this class for a parameter you register. Then, you access the actual metaparameter by using this class’s as the key for metaparameter dictionary, for example, a sound event’s or the asset registry’s . As an opaque derived object, this class adds no properties to the subclass.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGlobalMetaParameterAsset
type PHASEGlobalMetaParameterAsset struct {
	PHASEAsset
}

// PHASEGlobalMetaParameterAssetFrom constructs a [PHASEGlobalMetaParameterAsset] from an unsafe.Pointer.
//
// A reference to a registered metaparameter that the app can share with multiple sound events or sources.
func PHASEGlobalMetaParameterAssetFrom(ptr unsafe.Pointer) PHASEGlobalMetaParameterAsset {
	return PHASEGlobalMetaParameterAsset{
		PHASEAsset: PHASEAssetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEGlobalMetaParameterAssetClass) Alloc() PHASEGlobalMetaParameterAsset {
	rv := objc.Send[PHASEGlobalMetaParameterAsset](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEGlobalMetaParameterAssetClass) New() PHASEGlobalMetaParameterAsset {
	rv := objc.Send[PHASEGlobalMetaParameterAsset](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEGlobalMetaParameterAsset) Init() PHASEGlobalMetaParameterAsset {
	rv := objc.Send[PHASEGlobalMetaParameterAsset](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEGlobalMetaParameterAsset) Autorelease() PHASEGlobalMetaParameterAsset {
	rv := objc.Send[PHASEGlobalMetaParameterAsset](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEGlobalMetaParameterAsset creates a new PHASEGlobalMetaParameterAsset instance.
func NewPHASEGlobalMetaParameterAsset() PHASEGlobalMetaParameterAsset {
	return getPHASEGlobalMetaParameterAssetClass().New()
}


// A unique name for the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseasset/identifier
func (p_ PHASEGlobalMetaParameterAsset) Identifier() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// A unique name for the asset.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseasset/identifier
func (p_ PHASEGlobalMetaParameterAsset) SetIdentifier(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), value)
}

// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEGlobalMetaParameterAsset) GlobalMetaParameters() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}


// SetGlobalMetaParameters sets the value of the globalMetaParameters property.
// A dictionary of metaparameters that all sound event assets share.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEGlobalMetaParameterAsset) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}

// The object’s meta parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEGlobalMetaParameterAsset) MetaParameters() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}


// SetMetaParameters sets the value of the metaParameters property.
// The object’s meta parameters.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEGlobalMetaParameterAsset) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}



