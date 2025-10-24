// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEAsset */


/* debug [class_header]: Header for PHASEAsset */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEAsset */
// An interface definition for the [PHASEAsset] class.
type IPHASEAsset interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEAsset */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEAsset */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEAsset */
// Alloc allocates a new instance without initialization.
func (pc _PHASEAssetClass) Alloc() PHASEAsset {
	rv := objc.Send[PHASEAsset](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEAsset */
// A base class that adds a name to framework assets.
//
// Through inheritance, this class adds a string to subclasses, for example, and . PHASE generates objects of this type based on template subclasses. For example, PHASE gives you a when you register a with the asset registry via .


// A base class that adds a name to framework assets.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEAsset *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEAsset */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEAsset */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEAsset */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEAsset */

// A unique name for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAsset/identifier
func (p_ PHASEAsset) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEAsset */



