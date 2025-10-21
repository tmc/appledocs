// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASESoundEventNodeAsset] class.
var (
	PHASESoundEventNodeAssetClass     _PHASESoundEventNodeAssetClass
	PHASESoundEventNodeAssetClassOnce sync.Once
)

func getPHASESoundEventNodeAssetClass() _PHASESoundEventNodeAssetClass {
	PHASESoundEventNodeAssetClassOnce.Do(func() {
		PHASESoundEventNodeAssetClass = _PHASESoundEventNodeAssetClass{objc.GetClass("PHASESoundEventNodeAsset")}
	})
	return PHASESoundEventNodeAssetClass
}

type _PHASESoundEventNodeAssetClass struct {
	class objc.Class
}

// An interface definition for the [PHASESoundEventNodeAsset] class.
type IPHASESoundEventNodeAsset interface {
	IPHASEAsset
}

// A template object for sounds that can play in reaction to environmental state.
//
// This object refers by name to a collection of sound event nodes that connect to form a tree, or hierarchy. To retrieve an instance of this class, add a sound-event node definition to the asset registry using . Choose the argument from the subclasses in based on the playback features your app requires. To play a single audio asset, register a with only one audio-providing node. Alternatively, to create a sound event that can change its audio based on your app’s current state, register a that contains children. By adding multiple nodes that play varying audio as children to a control node, PHASE plays the right audio for the moment based on control logic that you define. To create a playable sound event from this class, pass to the parameter of the sound event intializer, . Then, invoke the sound event by calling . As an opaque derived object, this class adds no properties to its base class.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventNodeAsset
type PHASESoundEventNodeAsset struct {
	PHASEAsset
}

// PHASESoundEventNodeAssetFrom constructs a [PHASESoundEventNodeAsset] from an unsafe.Pointer.
//
// A template object for sounds that can play in reaction to environmental state.
func PHASESoundEventNodeAssetFrom(ptr unsafe.Pointer) PHASESoundEventNodeAsset {
	return PHASESoundEventNodeAsset{
		PHASEAsset: PHASEAssetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASESoundEventNodeAssetClass) Alloc() PHASESoundEventNodeAsset {
	rv := objc.Send[PHASESoundEventNodeAsset](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASESoundEventNodeAssetClass) New() PHASESoundEventNodeAsset {
	rv := objc.Send[PHASESoundEventNodeAsset](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESoundEventNodeAsset) Init() PHASESoundEventNodeAsset {
	rv := objc.Send[PHASESoundEventNodeAsset](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESoundEventNodeAsset) Autorelease() PHASESoundEventNodeAsset {
	rv := objc.Send[PHASESoundEventNodeAsset](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESoundEventNodeAsset creates a new PHASESoundEventNodeAsset instance.
func NewPHASESoundEventNodeAsset() PHASESoundEventNodeAsset {
	return getPHASESoundEventNodeAssetClass().New()
}


// A unique name for the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseasset/identifier
func (p_ PHASESoundEventNodeAsset) Identifier() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// A unique name for the asset.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseasset/identifier
func (p_ PHASESoundEventNodeAsset) SetIdentifier(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), value)
}



