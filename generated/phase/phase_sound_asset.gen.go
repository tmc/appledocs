// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHASESoundAsset] class.
var (
	PHASESoundAssetClass     _PHASESoundAssetClass
	PHASESoundAssetClassOnce sync.Once
)

func getPHASESoundAssetClass() _PHASESoundAssetClass {
	PHASESoundAssetClassOnce.Do(func() {
		PHASESoundAssetClass = _PHASESoundAssetClass{objc.GetClass("PHASESoundAsset")}
	})
	return PHASESoundAssetClass
}

type _PHASESoundAssetClass struct {
	class objc.Class
}

// An interface definition for the [PHASESoundAsset] class.
type IPHASESoundAsset interface {
	IPHASEAsset
	Data() foundation.NSData
	Type() PHASEAssetType
	Url() foundation.URL
}

// A sound resource stored in the asset registry.
//
// This class wraps source audio data that an app intends to play. The framework requires a mixer to play a sound asset, and sound event nodes like combine the asset with a mixer. To provide a sound asset to a sound-event node, refer to the asset by the you pass into the function.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundAsset
type PHASESoundAsset struct {
	PHASEAsset
}

// PHASESoundAssetFrom constructs a [PHASESoundAsset] from an unsafe.Pointer.
//
// A sound resource stored in the asset registry.
func PHASESoundAssetFrom(ptr unsafe.Pointer) PHASESoundAsset {
	return PHASESoundAsset{
		PHASEAsset: PHASEAssetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASESoundAssetClass) Alloc() PHASESoundAsset {
	rv := objc.Send[PHASESoundAsset](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASESoundAssetClass) New() PHASESoundAsset {
	rv := objc.Send[PHASESoundAsset](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESoundAsset) Init() PHASESoundAsset {
	rv := objc.Send[PHASESoundAsset](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESoundAsset) Autorelease() PHASESoundAsset {
	rv := objc.Send[PHASESoundAsset](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESoundAsset creates a new PHASESoundAsset instance.
func NewPHASESoundAsset() PHASESoundAsset {
	return getPHASESoundAssetClass().New()
}


// A storage buffer for the sound asset.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundAsset/data
func (p_ PHASESoundAsset) Data() foundation.NSData {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("data"))
	return rv
}

// The type of sound asset.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundAsset/type
func (p_ PHASESoundAsset) Type() PHASEAssetType {
	rv := objc.Send[PHASEAssetType](p_.ID, objc.Sel("type"))
	return rv
}

// The URL of the sound asset.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundAsset/url
func (p_ PHASESoundAsset) Url() foundation.URL {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("url"))
	return rv
}



