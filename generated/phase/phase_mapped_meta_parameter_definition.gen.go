// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASEMappedMetaParameterDefinition */


/* debug [class_header]: Header for PHASEMappedMetaParameterDefinition */
// The class instance for the [PHASEMappedMetaParameterDefinition] class.
var (
	PHASEMappedMetaParameterDefinitionClass     _PHASEMappedMetaParameterDefinitionClass
	PHASEMappedMetaParameterDefinitionClassOnce sync.Once
)

func getPHASEMappedMetaParameterDefinitionClass() _PHASEMappedMetaParameterDefinitionClass {
	PHASEMappedMetaParameterDefinitionClassOnce.Do(func() {
		PHASEMappedMetaParameterDefinitionClass = _PHASEMappedMetaParameterDefinitionClass{objc.GetClass("PHASEMappedMetaParameterDefinition")}
	})
	return PHASEMappedMetaParameterDefinitionClass
}

type _PHASEMappedMetaParameterDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEMappedMetaParameterDefinition */
// An interface definition for the [PHASEMappedMetaParameterDefinition] class.
type IPHASEMappedMetaParameterDefinition interface {
	IPHASENumberMetaParameterDefinition
	
/* debug [class_interface_properties]: Properties for PHASEMappedMetaParameterDefinition */
	// properties:
	Envelope() IPHASEEnvelope
	InputMetaParameterDefinition() IPHASENumberMetaParameterDefinition
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEMappedMetaParameterDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEMappedMetaParameterDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEMappedMetaParameterDefinitionClass) Alloc() PHASEMappedMetaParameterDefinition {
	rv := objc.Send[PHASEMappedMetaParameterDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEMappedMetaParameterDefinitionClass) New() PHASEMappedMetaParameterDefinition {
	rv := objc.Send[PHASEMappedMetaParameterDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEMappedMetaParameterDefinition) Init() PHASEMappedMetaParameterDefinition {
	rv := objc.Send[PHASEMappedMetaParameterDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEMappedMetaParameterDefinition) Autorelease() PHASEMappedMetaParameterDefinition {
	rv := objc.Send[PHASEMappedMetaParameterDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEMappedMetaParameterDefinition creates a new PHASEMappedMetaParameterDefinition instance.
func NewPHASEMappedMetaParameterDefinition() PHASEMappedMetaParameterDefinition {
	return getPHASEMappedMetaParameterDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEMappedMetaParameterDefinition */
// A metaparameter that graphs an input value on a set of mathematical curves.
//
// This class takes a metaparameter as input and plots its value on a curve defined by the property. Whereas the envelope’s function in and takes time because the relevant audio starts as its input parameter, in the case of the envelope property for this class, the app has full control over the input metaparameter’s value.


// A metaparameter that graphs an input value on a set of mathematical curves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMappedMetaParameterDefinition
type PHASEMappedMetaParameterDefinition struct {
	PHASENumberMetaParameterDefinition
}

// PHASEMappedMetaParameterDefinitionFrom constructs a [PHASEMappedMetaParameterDefinition] from an unsafe.Pointer.
//
// A metaparameter that graphs an input value on a set of mathematical curves.
func PHASEMappedMetaParameterDefinitionFrom(ptr unsafe.Pointer) PHASEMappedMetaParameterDefinition {
	return PHASEMappedMetaParameterDefinition{
		PHASENumberMetaParameterDefinition: PHASENumberMetaParameterDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEMappedMetaParameterDefinition */

// Creates a specification for a metaparameter that the app plots on a graph defined by the given set of curves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMappedMetaParameterDefinition/init(inputMetaParameterDefinition:envelope:)
func NewPHASEMappedMetaParameterDefinitionWithInputMetaParameterDefinitionEnvelope(inputMetaParameterDefinition IPHASENumberMetaParameterDefinition, envelope IPHASEEnvelope) PHASEMappedMetaParameterDefinition {
	instance := getPHASEMappedMetaParameterDefinitionClass().Alloc()
	rv := objc.Send[PHASEMappedMetaParameterDefinition](instance.ID, objc.Sel("initWithInputMetaParameterDefinition:envelope:"), inputMetaParameterDefinition, envelope)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEMappedMetaParameterDefinitionWithInputMetaParameterDefinitionEnvelope */


// Creates a specification for a named metaparameter that the app plots on a graph defined by the given set of curves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMappedMetaParameterDefinition/init(inputMetaParameterDefinition:envelope:identifier:)
func NewPHASEMappedMetaParameterDefinitionWithInputMetaParameterDefinitionEnvelopeIdentifier(inputMetaParameterDefinition IPHASENumberMetaParameterDefinition, envelope IPHASEEnvelope, identifier objc.IObject /* cross-framework: NSString */) PHASEMappedMetaParameterDefinition {
	instance := getPHASEMappedMetaParameterDefinitionClass().Alloc()
	rv := objc.Send[PHASEMappedMetaParameterDefinition](instance.ID, objc.Sel("initWithInputMetaParameterDefinition:envelope:identifier:"), inputMetaParameterDefinition, envelope, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEMappedMetaParameterDefinitionWithInputMetaParameterDefinitionEnvelopeIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEMappedMetaParameterDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEMappedMetaParameterDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEMappedMetaParameterDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEMappedMetaParameterDefinition */

// A collection of line segments that curve and connect to form a graph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMappedMetaParameterDefinition/envelope
func (p_ PHASEMappedMetaParameterDefinition) Envelope() IPHASEEnvelope {
	rv := objc.Send[PHASEEnvelope](p_.ID, objc.Sel("envelope"))
	return rv
}/* debug [instance_properties/getter]: envelope */


// A linear input value to plot on a curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMappedMetaParameterDefinition/inputMetaParameterDefinition
func (p_ PHASEMappedMetaParameterDefinition) InputMetaParameterDefinition() IPHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("inputMetaParameterDefinition"))
	return rv
}/* debug [instance_properties/getter]: inputMetaParameterDefinition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEMappedMetaParameterDefinition */


