// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASEGlobalMetaParameterAsset */


/* debug [class_header]: Header for PHASEGlobalMetaParameterAsset */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEGlobalMetaParameterAsset */
// An interface definition for the [PHASEGlobalMetaParameterAsset] class.
type IPHASEGlobalMetaParameterAsset interface {
	IPHASEAsset
	
/* debug [class_interface_properties]: Properties for PHASEGlobalMetaParameterAsset */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	GlobalMetaParameters() IPHASEMetaParameter
	SetGlobalMetaParameters(value IPHASEMetaParameter)
	MetaParameters() IPHASEMetaParameter
	SetMetaParameters(value IPHASEMetaParameter)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEGlobalMetaParameterAsset */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEGlobalMetaParameterAsset */
// Alloc allocates a new instance without initialization.
func (pc _PHASEGlobalMetaParameterAssetClass) Alloc() PHASEGlobalMetaParameterAsset {
	rv := objc.Send[PHASEGlobalMetaParameterAsset](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEGlobalMetaParameterAsset */
// A reference to a registered metaparameter that the app can share with multiple sound events or sources.
//
// The engine’s function returns an instance of this class for a parameter you register. Then, you access the actual metaparameter by using this class’s as the key for metaparameter dictionary, for example, a sound event’s or the asset registry’s . As an opaque derived object, this class adds no properties to the subclass.


// A reference to a registered metaparameter that the app can share with multiple sound events or sources.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEGlobalMetaParameterAsset *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEGlobalMetaParameterAsset */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEGlobalMetaParameterAsset */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEGlobalMetaParameterAsset */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEGlobalMetaParameterAsset */

// A unique name for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseasset/identifier
func (p_ PHASEGlobalMetaParameterAsset) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A unique name for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseasset/identifier
func (p_ PHASEGlobalMetaParameterAsset) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEGlobalMetaParameterAsset) GlobalMetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}/* debug [instance_properties/getter]: globalMetaParameters */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEGlobalMetaParameterAsset) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}/* debug [instance_properties/setter]: globalMetaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEGlobalMetaParameterAsset) MetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}/* debug [instance_properties/getter]: metaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEGlobalMetaParameterAsset) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}/* debug [instance_properties/setter]: metaParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEGlobalMetaParameterAsset */



