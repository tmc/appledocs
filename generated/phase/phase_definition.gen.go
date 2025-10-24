// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEDefinition */


/* debug [class_header]: Header for PHASEDefinition */
// The class instance for the [PHASEDefinition] class.
var (
	PHASEDefinitionClass     _PHASEDefinitionClass
	PHASEDefinitionClassOnce sync.Once
)

func getPHASEDefinitionClass() _PHASEDefinitionClass {
	PHASEDefinitionClassOnce.Do(func() {
		PHASEDefinitionClass = _PHASEDefinitionClass{objc.GetClass("PHASEDefinition")}
	})
	return PHASEDefinitionClass
}

type _PHASEDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEDefinition */
// An interface definition for the [PHASEDefinition] class.
type IPHASEDefinition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEDefinition */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	GlobalMetaParameters() IPHASEMetaParameter
	SetGlobalMetaParameters(value IPHASEMetaParameter)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEDefinitionClass) Alloc() PHASEDefinition {
	rv := objc.Send[PHASEDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEDefinitionClass) New() PHASEDefinition {
	rv := objc.Send[PHASEDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEDefinition) Init() PHASEDefinition {
	rv := objc.Send[PHASEDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEDefinition) Autorelease() PHASEDefinition {
	rv := objc.Send[PHASEDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEDefinition creates a new PHASEDefinition instance.
func NewPHASEDefinition() PHASEDefinition {
	return getPHASEDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEDefinition */
// A base class that adds a name to framework definitions.
//
// Various PHASE classes derive from this class, for example, , , and . This class represents a template from which PHASE creates concrete subclasses at runtime. For example, when you register a global metaparameter definition using , PHASE returns a subclass, , that identifies a usable metaparameter by name. To access the usable metaparameter, pass the into the dictionary.


// A base class that adds a name to framework definitions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDefinition
type PHASEDefinition struct {
	objectivec.Object
}

// PHASEDefinitionFrom constructs a [PHASEDefinition] from an unsafe.Pointer.
//
// A base class that adds a name to framework definitions.
func PHASEDefinitionFrom(ptr unsafe.Pointer) PHASEDefinition {
	return PHASEDefinition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEDefinition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEDefinition */

// A unique name for the definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDefinition/identifier
func (p_ PHASEDefinition) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEDefinition) GlobalMetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}/* debug [instance_properties/getter]: globalMetaParameters */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEDefinition) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}/* debug [instance_properties/setter]: globalMetaParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEDefinition */



