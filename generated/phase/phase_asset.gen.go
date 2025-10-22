// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEAsset] class.
var (
	PHASEAssetClass     _PHASEAssetClass
	PHASEAssetClassOnce sync.Once
)

func getPHASEAssetClass() _PHASEAssetClass {
	PHASEAssetClassOnce.Do(func() {
		PHASEAssetClass = _PHASEAssetClass{objc.GetClass("PHASEAsset")}
	})
	return PHASEAssetClass
}

type _PHASEAssetClass struct {
	class objc.Class
}

// An interface definition for the [PHASEAsset] class.
type IPHASEAsset interface {
	objectivec.IObject
	Identifier() string
}

// A base class that adds a name to framework assets.
//
// Through inheritance, this class adds a string to subclasses, for example, and . PHASE generates objects of this type based on template subclasses. For example, PHASE gives you a when you register a with the asset registry via .
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAsset
type PHASEAsset struct {
	objectivec.Object
}

// PHASEAssetFrom constructs a [PHASEAsset] from an unsafe.Pointer.
//
// A base class that adds a name to framework assets.
func PHASEAssetFrom(ptr unsafe.Pointer) PHASEAsset {
	return PHASEAsset{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEAssetClass) Alloc() PHASEAsset {
	rv := objc.Send[PHASEAsset](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEAssetClass) New() PHASEAsset {
	rv := objc.Send[PHASEAsset](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEAsset) Init() PHASEAsset {
	rv := objc.Send[PHASEAsset](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEAsset) Autorelease() PHASEAsset {
	rv := objc.Send[PHASEAsset](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEAsset creates a new PHASEAsset instance.
func NewPHASEAsset() PHASEAsset {
	return getPHASEAssetClass().New()
}


// A unique name for the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAsset/identifier
func (p_ PHASEAsset) Identifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}



