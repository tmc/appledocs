// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEDucker] class.
var (
	PHASEDuckerClass     _PHASEDuckerClass
	PHASEDuckerClassOnce sync.Once
)

func getPHASEDuckerClass() _PHASEDuckerClass {
	PHASEDuckerClassOnce.Do(func() {
		PHASEDuckerClass = _PHASEDuckerClass{objc.GetClass("PHASEDucker")}
	})
	return PHASEDuckerClass
}

type _PHASEDuckerClass struct {
	class objc.Class
}

// An interface definition for the [PHASEDucker] class.
type IPHASEDucker interface {
	objectivec.IObject
	Activate()
	Deactivate()
}

// An object that manages competing sounds.
//
// When a sound plays in any of the source groups, this class lowers the volume of all the target groups so the listener hears the source sound more clearly. You set the source and target using objects; see and .
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker
type PHASEDucker struct {
	objectivec.Object
}

// PHASEDuckerFrom constructs a [PHASEDucker] from an unsafe.Pointer.
//
// An object that manages competing sounds.
func PHASEDuckerFrom(ptr unsafe.Pointer) PHASEDucker {
	return PHASEDucker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEDuckerClass) Alloc() PHASEDucker {
	rv := objc.Send[PHASEDucker](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEDuckerClass) New() PHASEDucker {
	rv := objc.Send[PHASEDucker](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEDucker) Init() PHASEDucker {
	rv := objc.Send[PHASEDucker](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEDucker) Autorelease() PHASEDucker {
	rv := objc.Send[PHASEDucker](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEDucker creates a new PHASEDucker instance.
func NewPHASEDucker() PHASEDucker {
	return getPHASEDuckerClass().New()
}




// Creates an object that manages competing sounds.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/init(engine:sourceGroups:targetGroups:gain:attackTime:releaseTime:attackCurve:releaseCurve:)
func NewPHASEDuckerWithEngineSourceGroupsTargetGroupsGainAttackTimeReleaseTimeAttackCurveReleaseCurve(engine unsafe.Pointer, sourceGroups unsafe.Pointer, targetGroups unsafe.Pointer, gain unsafe.Pointer, attackTime unsafe.Pointer, releaseTime unsafe.Pointer, attackCurve unsafe.Pointer, releaseCurve unsafe.Pointer) PHASEDucker {
	instance := getPHASEDuckerClass().Alloc()
	rv := objc.Send[PHASEDucker](instance.ID, objc.Sel("initWithEngine:sourceGroups:targetGroups:gain:attackTime:releaseTime:attackCurve:releaseCurve:"), engine, sourceGroups, targetGroups, gain, attackTime, releaseTime, attackCurve, releaseCurve)
	rv.Autorelease()
	return rv
}


// Instructs the ducker to begin altering sound.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/activate()
func (p_ PHASEDucker) Activate() {
	objc.Send[objc.ID](p_.ID, objc.Sel("activate"))
}

// Stops the ducker from altering sound.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/deactivate()
func (p_ PHASEDucker) Deactivate() {
	objc.Send[objc.ID](p_.ID, objc.Sel("deactivate"))
}

// A mathematical curve that shapes transition progress as sound reduction begins.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/attackCurve
func (p_ PHASEDucker) AttackCurve() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("attackCurve"))
	return rv
}

// The amount of time for sound reduction to reach maximum strength.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/attackTime
func (p_ PHASEDucker) AttackTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("attackTime"))
	return rv
}

// The amount of volume reduction.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/gain
func (p_ PHASEDucker) Gain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("gain"))
	return rv
}

// A unique value for the ducker.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/identifier
func (p_ PHASEDucker) Identifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}

// A Boolean value that determines whether the ducker reduces sound.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/isActive
func (p_ PHASEDucker) Active() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("active"))
	return rv
}

// A mathematical curve that shapes transition progress as sound reduction ends.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/releaseCurve
func (p_ PHASEDucker) ReleaseCurve() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("releaseCurve"))
	return rv
}

// The amount of time to transition from maximum sound reduction to no reduction.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/releaseTime
func (p_ PHASEDucker) ReleaseTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("releaseTime"))
	return rv
}

// The sounds that determine volume reduction.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/sourceGroups
func (p_ PHASEDucker) SourceGroups() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("sourceGroups"))
	return rv
}

// The sounds that reduce in volume.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/targetGroups
func (p_ PHASEDucker) TargetGroups() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("targetGroups"))
	return rv
}


