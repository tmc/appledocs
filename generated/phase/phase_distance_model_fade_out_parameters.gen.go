// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEDistanceModelFadeOutParameters] class.
var (
	PHASEDistanceModelFadeOutParametersClass     _PHASEDistanceModelFadeOutParametersClass
	PHASEDistanceModelFadeOutParametersClassOnce sync.Once
)

func getPHASEDistanceModelFadeOutParametersClass() _PHASEDistanceModelFadeOutParametersClass {
	PHASEDistanceModelFadeOutParametersClassOnce.Do(func() {
		PHASEDistanceModelFadeOutParametersClass = _PHASEDistanceModelFadeOutParametersClass{objc.GetClass("PHASEDistanceModelFadeOutParameters")}
	})
	return PHASEDistanceModelFadeOutParametersClass
}

type _PHASEDistanceModelFadeOutParametersClass struct {
	class objc.Class
}

// An interface definition for the [PHASEDistanceModelFadeOutParameters] class.
type IPHASEDistanceModelFadeOutParameters interface {
	objectivec.IObject
}

// A distance over which the framework fades out sound.
//
// For spatial sound output, the framework stops playing a sound when its distance from the listener surpases . The framework gradually fades out the sound’s volume as the distance between the source and listener approaches . Likewise, the framework gradually fades in the sound as the distance between the source and listener approaches . A object provides an instance of this class to a spatial mixer; for more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelFadeOutParameters
type PHASEDistanceModelFadeOutParameters struct {
	objectivec.Object
}

// PHASEDistanceModelFadeOutParametersFrom constructs a [PHASEDistanceModelFadeOutParameters] from an unsafe.Pointer.
//
// A distance over which the framework fades out sound.
func PHASEDistanceModelFadeOutParametersFrom(ptr unsafe.Pointer) PHASEDistanceModelFadeOutParameters {
	return PHASEDistanceModelFadeOutParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEDistanceModelFadeOutParametersClass) Alloc() PHASEDistanceModelFadeOutParameters {
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEDistanceModelFadeOutParametersClass) New() PHASEDistanceModelFadeOutParameters {
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEDistanceModelFadeOutParameters) Init() PHASEDistanceModelFadeOutParameters {
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEDistanceModelFadeOutParameters) Autorelease() PHASEDistanceModelFadeOutParameters {
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEDistanceModelFadeOutParameters creates a new PHASEDistanceModelFadeOutParameters instance.
func NewPHASEDistanceModelFadeOutParameters() PHASEDistanceModelFadeOutParameters {
	return getPHASEDistanceModelFadeOutParametersClass().New()
}




// Creates a distance beyond which sound sources stop playing.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelFadeOutParameters/init(cullDistance:)
func NewPHASEDistanceModelFadeOutParametersWithCullDistance(cullDistance unsafe.Pointer) PHASEDistanceModelFadeOutParameters {
	instance := getPHASEDistanceModelFadeOutParametersClass().Alloc()
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](instance.ID, objc.Sel("initWithCullDistance:"), cullDistance)
	rv.Autorelease()
	return rv
}


// The distance beyond which the framework doesn’t process the sound.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDistanceModelFadeOutParameters/cullDistance
func (p_ PHASEDistanceModelFadeOutParameters) CullDistance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cullDistance"))
	return rv
}

// A distance over which the framework fades out the mixer’s sound.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasedistancemodelparameters/fadeoutparameters
func (p_ PHASEDistanceModelFadeOutParameters) FadeOutParameters() PHASEDistanceModelFadeOutParameters {
	rv := objc.Send[PHASEDistanceModelFadeOutParameters](p_.ID, objc.Sel("fadeOutParameters"))
	return rv
}


// SetFadeOutParameters sets the value of the fadeOutParameters property.
// A distance over which the framework fades out the mixer’s sound.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasedistancemodelparameters/fadeoutparameters
func (p_ PHASEDistanceModelFadeOutParameters) SetFadeOutParameters(value IPHASEDistanceModelFadeOutParameters) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFadeOutParameters:"), value)
}


