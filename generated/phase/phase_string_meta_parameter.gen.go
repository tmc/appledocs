// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEStringMetaParameter] class.
var (
	PHASEStringMetaParameterClass     _PHASEStringMetaParameterClass
	PHASEStringMetaParameterClassOnce sync.Once
)

func getPHASEStringMetaParameterClass() _PHASEStringMetaParameterClass {
	PHASEStringMetaParameterClassOnce.Do(func() {
		PHASEStringMetaParameterClass = _PHASEStringMetaParameterClass{objc.GetClass("PHASEStringMetaParameter")}
	})
	return PHASEStringMetaParameterClass
}

type _PHASEStringMetaParameterClass struct {
	class objc.Class
}

// An interface definition for the [PHASEStringMetaParameter] class.
type IPHASEStringMetaParameter interface {
	IPHASEMetaParameter
}

// A metaparameter with a text definition that can change over time.
//
// This class contains text that updates, like a “weather” metaparameter that the app changes from “rainy” to “sunny.” To create an instance of this class, first create a , and either: Register it with the engine by calling , then access the instance of this class in the engine’s dictionary. Pass it to the initializer, , and then access the instance of this class in a sound event’s dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStringMetaParameter
type PHASEStringMetaParameter struct {
	PHASEMetaParameter
}

// PHASEStringMetaParameterFrom constructs a [PHASEStringMetaParameter] from an unsafe.Pointer.
//
// A metaparameter with a text definition that can change over time.
func PHASEStringMetaParameterFrom(ptr unsafe.Pointer) PHASEStringMetaParameter {
	return PHASEStringMetaParameter{
		PHASEMetaParameter: PHASEMetaParameterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEStringMetaParameterClass) Alloc() PHASEStringMetaParameter {
	rv := objc.Send[PHASEStringMetaParameter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEStringMetaParameterClass) New() PHASEStringMetaParameter {
	rv := objc.Send[PHASEStringMetaParameter](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEStringMetaParameter) Init() PHASEStringMetaParameter {
	rv := objc.Send[PHASEStringMetaParameter](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEStringMetaParameter) Autorelease() PHASEStringMetaParameter {
	rv := objc.Send[PHASEStringMetaParameter](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEStringMetaParameter creates a new PHASEStringMetaParameter instance.
func NewPHASEStringMetaParameter() PHASEStringMetaParameter {
	return getPHASEStringMetaParameterClass().New()
}


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEStringMetaParameter) GlobalMetaParameters() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}


// SetGlobalMetaParameters sets the value of the globalMetaParameters property.
// A dictionary of metaparameters that all sound event assets share.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEStringMetaParameter) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}

// The object’s meta parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEStringMetaParameter) MetaParameters() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}


// SetMetaParameters sets the value of the metaParameters property.
// The object’s meta parameters.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEStringMetaParameter) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}



