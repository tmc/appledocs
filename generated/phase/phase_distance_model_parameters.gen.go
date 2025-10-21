// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEDistanceModelParameters] class.
var (
	PHASEDistanceModelParametersClass     _PHASEDistanceModelParametersClass
	PHASEDistanceModelParametersClassOnce sync.Once
)

func getPHASEDistanceModelParametersClass() _PHASEDistanceModelParametersClass {
	PHASEDistanceModelParametersClassOnce.Do(func() {
		PHASEDistanceModelParametersClass = _PHASEDistanceModelParametersClass{objc.GetClass("PHASEDistanceModelParameters")}
	})
	return PHASEDistanceModelParametersClass
}

type _PHASEDistanceModelParametersClass struct {
	class objc.Class
}

// An interface definition for the [PHASEDistanceModelParameters] class.
type IPHASEDistanceModelParameters interface {
	objectivec.IObject
}

// A base class for a sound’s rate of change over distance.
//
// When your app outputs sound with a 3D position and orientation, designate a subclass of this class to indicate the manner in which PHASE changes sound with distance. Assign an instance of either or , depending on your app’s needs, to the class’s property.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelParameters
type PHASEDistanceModelParameters struct {
	objectivec.Object
}

// PHASEDistanceModelParametersFrom constructs a [PHASEDistanceModelParameters] from an unsafe.Pointer.
//
// A base class for a sound’s rate of change over distance.
func PHASEDistanceModelParametersFrom(ptr unsafe.Pointer) PHASEDistanceModelParameters {
	return PHASEDistanceModelParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEDistanceModelParametersClass) Alloc() PHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEDistanceModelParametersClass) New() PHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEDistanceModelParameters) Init() PHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEDistanceModelParameters) Autorelease() PHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEDistanceModelParameters creates a new PHASEDistanceModelParameters instance.
func NewPHASEDistanceModelParameters() PHASEDistanceModelParameters {
	return getPHASEDistanceModelParametersClass().New()
}


// An effect that changes sound as it carries over a distance.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/distancemodelparameters
func (p_ PHASEDistanceModelParameters) DistanceModelParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("distanceModelParameters"))
	return rv
}


// SetDistanceModelParameters sets the value of the distanceModelParameters property.
// An effect that changes sound as it carries over a distance.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/distancemodelparameters
func (p_ PHASEDistanceModelParameters) SetDistanceModelParameters(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDistanceModelParameters:"), value)
}

// A distance over which the framework fades out the mixer’s sound.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelParameters/fadeOutParameters
func (p_ PHASEDistanceModelParameters) FadeOutParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("fadeOutParameters"))
	return rv
}


// SetFadeOutParameters sets the value of the fadeOutParameters property.
// A distance over which the framework fades out the mixer’s sound.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelParameters/fadeOutParameters
func (p_ PHASEDistanceModelParameters) SetFadeOutParameters(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFadeOutParameters:"), value)
}



