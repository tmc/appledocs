// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEMetaParameter] class.
var (
	PHASEMetaParameterClass     _PHASEMetaParameterClass
	PHASEMetaParameterClassOnce sync.Once
)

func getPHASEMetaParameterClass() _PHASEMetaParameterClass {
	PHASEMetaParameterClassOnce.Do(func() {
		PHASEMetaParameterClass = _PHASEMetaParameterClass{objc.GetClass("PHASEMetaParameter")}
	})
	return PHASEMetaParameterClass
}

type _PHASEMetaParameterClass struct {
	class objc.Class
}

// An interface definition for the [PHASEMetaParameter] class.
type IPHASEMetaParameter interface {
	objectivec.IObject
}

// A named parameter with a value that the app can change over time.
//
// Instances of this class provide an app with dynamic control of a sound’s properties. A metaparameter takes a single value as input and may operate on one or more audio characteristics. To change the value of a metaparameter at runtime: Assign a string to a textual metaparameter’s . Adjust the value of a number or mapped metaparameter gradually over a duration by calling .
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMetaParameter
type PHASEMetaParameter struct {
	objectivec.Object
}

// PHASEMetaParameterFrom constructs a [PHASEMetaParameter] from an unsafe.Pointer.
//
// A named parameter with a value that the app can change over time.
func PHASEMetaParameterFrom(ptr unsafe.Pointer) PHASEMetaParameter {
	return PHASEMetaParameter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEMetaParameterClass) Alloc() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEMetaParameterClass) New() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEMetaParameter) Init() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEMetaParameter) Autorelease() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEMetaParameter creates a new PHASEMetaParameter instance.
func NewPHASEMetaParameter() PHASEMetaParameter {
	return getPHASEMetaParameterClass().New()
}




