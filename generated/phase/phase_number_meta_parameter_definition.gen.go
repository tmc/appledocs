// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASENumberMetaParameterDefinition] class.
var (
	PHASENumberMetaParameterDefinitionClass     _PHASENumberMetaParameterDefinitionClass
	PHASENumberMetaParameterDefinitionClassOnce sync.Once
)

func getPHASENumberMetaParameterDefinitionClass() _PHASENumberMetaParameterDefinitionClass {
	PHASENumberMetaParameterDefinitionClassOnce.Do(func() {
		PHASENumberMetaParameterDefinitionClass = _PHASENumberMetaParameterDefinitionClass{objc.GetClass("PHASENumberMetaParameterDefinition")}
	})
	return PHASENumberMetaParameterDefinitionClass
}

type _PHASENumberMetaParameterDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASENumberMetaParameterDefinition] class.
type IPHASENumberMetaParameterDefinition interface {
	IPHASEMetaParameterDefinition
	GlobalMetaParameters() PHASEMetaParameter
	SetGlobalMetaParameters(value IPHASEMetaParameter)
	InputMetaParameterDefinition() PHASENumberMetaParameterDefinition
	SetInputMetaParameterDefinition(value IPHASENumberMetaParameterDefinition)
	Maximum() float64
	SetMaximum(value float64)
	Minimum() float64
	SetMinimum(value float64)
	MetaParameters() PHASEMetaParameter
	SetMetaParameters(value IPHASEMetaParameter)
}

// A specification for a metaparameter defined by a number.
//
// Use this class to spawn discrete instances of , for example, a “player speed” metaparameter that the app changes gradually from to . To use a number metaparameter, create an instance of this class and: Register it with the engine by calling , then access the instance of this class in the engine’s dictionary. Pass it to the initializer, , and then access the instance of this class in a sound event’s dictionary. Pass it into the initializer, . Then, access the instance of this class using the mapped parameter’s property.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameterDefinition
type PHASENumberMetaParameterDefinition struct {
	PHASEMetaParameterDefinition
}

// PHASENumberMetaParameterDefinitionFrom constructs a [PHASENumberMetaParameterDefinition] from an unsafe.Pointer.
//
// A specification for a metaparameter defined by a number.
func PHASENumberMetaParameterDefinitionFrom(ptr unsafe.Pointer) PHASENumberMetaParameterDefinition {
	return PHASENumberMetaParameterDefinition{
		PHASEMetaParameterDefinition: PHASEMetaParameterDefinitionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASENumberMetaParameterDefinitionClass) Alloc() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASENumberMetaParameterDefinitionClass) New() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASENumberMetaParameterDefinition) Init() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASENumberMetaParameterDefinition) Autorelease() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASENumberMetaParameterDefinition creates a new PHASENumberMetaParameterDefinition instance.
func NewPHASENumberMetaParameterDefinition() PHASENumberMetaParameterDefinition {
	return getPHASENumberMetaParameterDefinitionClass().New()
}


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASENumberMetaParameterDefinition) GlobalMetaParameters() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}


// SetGlobalMetaParameters sets the value of the globalMetaParameters property.
// A dictionary of metaparameters that all sound event assets share.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASENumberMetaParameterDefinition) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}

// A linear input value to plot on a curve.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasemappedmetaparameterdefinition/inputmetaparameterdefinition
func (p_ PHASENumberMetaParameterDefinition) InputMetaParameterDefinition() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("inputMetaParameterDefinition"))
	return rv
}


// SetInputMetaParameterDefinition sets the value of the inputMetaParameterDefinition property.
// A linear input value to plot on a curve.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasemappedmetaparameterdefinition/inputmetaparameterdefinition
func (p_ PHASENumberMetaParameterDefinition) SetInputMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInputMetaParameterDefinition:"), value)
}

// The highest possible number for the value.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasenumbermetaparameterdefinition/maximum
func (p_ PHASENumberMetaParameterDefinition) Maximum() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("maximum"))
	return rv
}


// SetMaximum sets the value of the maximum property.
// The highest possible number for the value.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasenumbermetaparameterdefinition/maximum
func (p_ PHASENumberMetaParameterDefinition) SetMaximum(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaximum:"), value)
}

// The lowest possible number for the value.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasenumbermetaparameterdefinition/minimum
func (p_ PHASENumberMetaParameterDefinition) Minimum() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minimum"))
	return rv
}


// SetMinimum sets the value of the minimum property.
// The lowest possible number for the value.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasenumbermetaparameterdefinition/minimum
func (p_ PHASENumberMetaParameterDefinition) SetMinimum(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMinimum:"), value)
}

// The object’s meta parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASENumberMetaParameterDefinition) MetaParameters() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}


// SetMetaParameters sets the value of the metaParameters property.
// The object’s meta parameters.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASENumberMetaParameterDefinition) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}



