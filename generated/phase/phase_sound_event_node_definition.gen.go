// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASESoundEventNodeDefinition */


/* debug [class_header]: Header for PHASESoundEventNodeDefinition */
// The class instance for the [PHASESoundEventNodeDefinition] class.
var (
	PHASESoundEventNodeDefinitionClass     _PHASESoundEventNodeDefinitionClass
	PHASESoundEventNodeDefinitionClassOnce sync.Once
)

func getPHASESoundEventNodeDefinitionClass() _PHASESoundEventNodeDefinitionClass {
	PHASESoundEventNodeDefinitionClassOnce.Do(func() {
		PHASESoundEventNodeDefinitionClass = _PHASESoundEventNodeDefinitionClass{objc.GetClass("PHASESoundEventNodeDefinition")}
	})
	return PHASESoundEventNodeDefinitionClass
}

type _PHASESoundEventNodeDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASESoundEventNodeDefinition */
// An interface definition for the [PHASESoundEventNodeDefinition] class.
type IPHASESoundEventNodeDefinition interface {
	IPHASEDefinition
	
/* debug [class_interface_properties]: Properties for PHASESoundEventNodeDefinition */
	// properties:
	Children() []PHASESoundEventNodeDefinition
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	MetaParameters() IPHASEMetaParameter
	SetMetaParameters(value IPHASEMetaParameter)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASESoundEventNodeDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASESoundEventNodeDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASESoundEventNodeDefinitionClass) Alloc() PHASESoundEventNodeDefinition {
	rv := objc.Send[PHASESoundEventNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASESoundEventNodeDefinitionClass) New() PHASESoundEventNodeDefinition {
	rv := objc.Send[PHASESoundEventNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESoundEventNodeDefinition) Init() PHASESoundEventNodeDefinition {
	rv := objc.Send[PHASESoundEventNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESoundEventNodeDefinition) Autorelease() PHASESoundEventNodeDefinition {
	rv := objc.Send[PHASESoundEventNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESoundEventNodeDefinition creates a new PHASESoundEventNodeDefinition instance.
func NewPHASESoundEventNodeDefinition() PHASESoundEventNodeDefinition {
	return getPHASESoundEventNodeDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASESoundEventNodeDefinition */
// A base class for sound event nodes that connect to form a node hierarchy.
//
// This class defines the base functionality for an object that, depending on the derived class’s type, either plays audio or hands off the invocation to one or more other nodes.


// A base class for sound event nodes that connect to form a node hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventNodeDefinition
type PHASESoundEventNodeDefinition struct {
	PHASEDefinition
}

// PHASESoundEventNodeDefinitionFrom constructs a [PHASESoundEventNodeDefinition] from an unsafe.Pointer.
//
// A base class for sound event nodes that connect to form a node hierarchy.
func PHASESoundEventNodeDefinitionFrom(ptr unsafe.Pointer) PHASESoundEventNodeDefinition {
	return PHASESoundEventNodeDefinition{
		PHASEDefinition: PHASEDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASESoundEventNodeDefinition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASESoundEventNodeDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASESoundEventNodeDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASESoundEventNodeDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASESoundEventNodeDefinition */

// An array of child sound event nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventNodeDefinition/children
func (p_ PHASESoundEventNodeDefinition) Children() []PHASESoundEventNodeDefinition {
	rv := objc.Send[[]PHASESoundEventNodeDefinition](p_.ID, objc.Sel("children"))
	return rv
}/* debug [instance_properties/getter]: children */


// A unique name for the definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasedefinition/identifier
func (p_ PHASESoundEventNodeDefinition) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A unique name for the definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasedefinition/identifier
func (p_ PHASESoundEventNodeDefinition) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASESoundEventNodeDefinition) MetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}/* debug [instance_properties/getter]: metaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASESoundEventNodeDefinition) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}/* debug [instance_properties/setter]: metaParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASESoundEventNodeDefinition */



