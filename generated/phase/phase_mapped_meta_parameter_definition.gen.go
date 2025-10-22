// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [PHASEMappedMetaParameterDefinition] class.
type IPHASEMappedMetaParameterDefinition interface {
	IPHASENumberMetaParameterDefinition
	Envelope() PHASEEnvelope
	SetEnvelope(value IPHASEEnvelope)
	InputMetaParameterDefinition() PHASENumberMetaParameterDefinition
	SetInputMetaParameterDefinition(value IPHASENumberMetaParameterDefinition)
}

// A metaparameter that graphs an input value on a set of mathematical curves.
//
// This class takes a metaparameter as input and plots its value on a curve defined by the property. Whereas the envelope’s function in and takes time because the relevant audio starts as its input parameter, in the case of the envelope property for this class, the app has full control over the input metaparameter’s value.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASEMappedMetaParameterDefinitionClass) Alloc() PHASEMappedMetaParameterDefinition {
	rv := objc.Send[PHASEMappedMetaParameterDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a specification for a named metaparameter that the app plots on a graph defined by the given set of curves.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMappedMetaParameterDefinition/init(inputMetaParameterDefinition:envelope:identifier:)
func NewPHASEMappedMetaParameterDefinitionWithInputMetaParameterDefinitionEnvelopeIdentifier(inputMetaParameterDefinition IPHASENumberMetaParameterDefinition, envelope IPHASEEnvelope, identifier string) PHASEMappedMetaParameterDefinition {
	instance := getPHASEMappedMetaParameterDefinitionClass().Alloc()
	rv := objc.Send[PHASEMappedMetaParameterDefinition](instance.ID, objc.Sel("initWithInputMetaParameterDefinition:envelope:identifier:"), inputMetaParameterDefinition, envelope, objc.String(identifier))
	rv.Autorelease()
	return rv
}


// A collection of line segments that curve and connect to form a graph.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasemappedmetaparameterdefinition/envelope
func (p_ PHASEMappedMetaParameterDefinition) Envelope() PHASEEnvelope {
	rv := objc.Send[PHASEEnvelope](p_.ID, objc.Sel("envelope"))
	return rv
}


// SetEnvelope sets the value of the envelope property.
// A collection of line segments that curve and connect to form a graph.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasemappedmetaparameterdefinition/envelope
func (p_ PHASEMappedMetaParameterDefinition) SetEnvelope(value IPHASEEnvelope) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEnvelope:"), value)
}

// A linear input value to plot on a curve.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasemappedmetaparameterdefinition/inputmetaparameterdefinition
func (p_ PHASEMappedMetaParameterDefinition) InputMetaParameterDefinition() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("inputMetaParameterDefinition"))
	return rv
}


// SetInputMetaParameterDefinition sets the value of the inputMetaParameterDefinition property.
// A linear input value to plot on a curve.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasemappedmetaparameterdefinition/inputmetaparameterdefinition
func (p_ PHASEMappedMetaParameterDefinition) SetInputMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInputMetaParameterDefinition:"), value)
}


