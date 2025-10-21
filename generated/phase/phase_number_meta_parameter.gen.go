// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASENumberMetaParameter] class.
var (
	PHASENumberMetaParameterClass     _PHASENumberMetaParameterClass
	PHASENumberMetaParameterClassOnce sync.Once
)

func getPHASENumberMetaParameterClass() _PHASENumberMetaParameterClass {
	PHASENumberMetaParameterClassOnce.Do(func() {
		PHASENumberMetaParameterClass = _PHASENumberMetaParameterClass{objc.GetClass("PHASENumberMetaParameter")}
	})
	return PHASENumberMetaParameterClass
}

type _PHASENumberMetaParameterClass struct {
	class objc.Class
}

// An interface definition for the [PHASENumberMetaParameter] class.
type IPHASENumberMetaParameter interface {
	IPHASEMetaParameter
}

// A metaparameter defined by a number that can change over time.
//
// This class contains a number that updates, like a “player speed” metaparameter that the app changes gradually from to . To create an instance of this class, first create a , and either: Register it with the engine by calling , then access the instance of this class in the engine’s dictionary. Pass it to the initializer, , and then access the instance of this class in a sound event’s dictionary. Use it as the input value for a by passing it into the initializer. Then, access the instance of this class using the mapped parameter’s property.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameter
type PHASENumberMetaParameter struct {
	PHASEMetaParameter
}

// PHASENumberMetaParameterFrom constructs a [PHASENumberMetaParameter] from an unsafe.Pointer.
//
// A metaparameter defined by a number that can change over time.
func PHASENumberMetaParameterFrom(ptr unsafe.Pointer) PHASENumberMetaParameter {
	return PHASENumberMetaParameter{
		PHASEMetaParameter: PHASEMetaParameterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASENumberMetaParameterClass) Alloc() PHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASENumberMetaParameterClass) New() PHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASENumberMetaParameter) Init() PHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASENumberMetaParameter) Autorelease() PHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASENumberMetaParameter creates a new PHASENumberMetaParameter instance.
func NewPHASENumberMetaParameter() PHASENumberMetaParameter {
	return getPHASENumberMetaParameterClass().New()
}




