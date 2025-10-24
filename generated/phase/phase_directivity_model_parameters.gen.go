// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEDirectivityModelParameters] class.
var (
	PHASEDirectivityModelParametersClass     _PHASEDirectivityModelParametersClass
	PHASEDirectivityModelParametersClassOnce sync.Once
)

func getPHASEDirectivityModelParametersClass() _PHASEDirectivityModelParametersClass {
	PHASEDirectivityModelParametersClassOnce.Do(func() {
		PHASEDirectivityModelParametersClass = _PHASEDirectivityModelParametersClass{objc.GetClass("PHASEDirectivityModelParameters")}
	})
	return PHASEDirectivityModelParametersClass
}

type _PHASEDirectivityModelParametersClass struct {
	class objc.Class
}

// An interface definition for the [PHASEDirectivityModelParameters] class.
type IPHASEDirectivityModelParameters interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A base class for objects that direct sound.
//
// Several classes derive from this class that implement a unique strategy to direct sound. Rather than create an instance of this class, instantiate a subclass, such as or .


// A base class for objects that direct sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDirectivityModelParameters
type PHASEDirectivityModelParameters struct {
	objectivec.Object
}

// PHASEDirectivityModelParametersFrom constructs a [PHASEDirectivityModelParameters] from an unsafe.Pointer.
//
// A base class for objects that direct sound.
func PHASEDirectivityModelParametersFrom(ptr unsafe.Pointer) PHASEDirectivityModelParameters {
	return PHASEDirectivityModelParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEDirectivityModelParametersClass) Alloc() PHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEDirectivityModelParametersClass) New() PHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEDirectivityModelParameters) Init() PHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEDirectivityModelParameters) Autorelease() PHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEDirectivityModelParameters creates a new PHASEDirectivityModelParameters instance.
func NewPHASEDirectivityModelParameters() PHASEDirectivityModelParameters {
	return getPHASEDirectivityModelParametersClass().New()
}




