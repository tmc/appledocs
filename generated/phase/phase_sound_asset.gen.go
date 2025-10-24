// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASESoundAsset */


/* debug [class_header]: Header for PHASESoundAsset */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASESoundAsset */
// An interface definition for the [PHASESoundAsset] class.
type IPHASESoundAsset interface {
	IPHASEAsset
	
/* debug [class_interface_properties]: Properties for PHASESoundAsset */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	Type() PHASEAssetType
	Url() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASESoundAsset */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASESoundAsset */
// Alloc allocates a new instance without initialization.
func (pc _PHASESoundAssetClass) Alloc() PHASESoundAsset {
	rv := objc.Send[PHASESoundAsset](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASESoundAsset */
// A sound resource stored in the asset registry.
//
// This class wraps source audio data that an app intends to play. The framework requires a mixer to play a sound asset, and sound event nodes like combine the asset with a mixer. To provide a sound asset to a sound-event node, refer to the asset by the you pass into the function.


// A sound resource stored in the asset registry.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASESoundAsset *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASESoundAsset */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASESoundAsset */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASESoundAsset */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASESoundAsset */

// A storage buffer for the sound asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundAsset/data
func (p_ PHASESoundAsset) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The type of sound asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundAsset/type
func (p_ PHASESoundAsset) Type() PHASEAssetType {
	rv := objc.Send[PHASEAssetType](p_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The URL of the sound asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundAsset/url
func (p_ PHASESoundAsset) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASESoundAsset */



