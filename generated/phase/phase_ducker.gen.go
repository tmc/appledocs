// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEDucker */


/* debug [class_header]: Header for PHASEDucker */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEDucker */
// An interface definition for the [PHASEDucker] class.
type IPHASEDucker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEDucker */
	// properties:
	AttackCurve() PHASECurveType
	AttackTime() float64
	Gain() float64
	Identifier() objc.IObject /* cross-framework: NSString */
	Active() bool
	ReleaseCurve() PHASECurveType
	ReleaseTime() float64
	SourceGroups() unsafe.Pointer
	TargetGroups() unsafe.Pointer
	IsActive() bool
	SetIsActive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEDucker */
	// methods:
	Activate()
	Deactivate()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEDucker */
// Alloc allocates a new instance without initialization.
func (pc _PHASEDuckerClass) Alloc() PHASEDucker {
	rv := objc.Send[PHASEDucker](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEDucker */
// An object that manages competing sounds.
//
// When a sound plays in any of the source groups, this class lowers the volume of all the target groups so the listener hears the source sound more clearly. You set the source and target using objects; see and .


// An object that manages competing sounds.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEDucker */

// Creates an object that manages competing sounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/init(engine:sourceGroups:targetGroups:gain:attackTime:releaseTime:attackCurve:releaseCurve:)
func NewPHASEDuckerWithEngineSourceGroupsTargetGroupsGainAttackTimeReleaseTimeAttackCurveReleaseCurve(engine IPHASEEngine, sourceGroups unsafe.Pointer, targetGroups unsafe.Pointer, gain float64, attackTime float64, releaseTime float64, attackCurve PHASECurveType, releaseCurve PHASECurveType) PHASEDucker {
	instance := getPHASEDuckerClass().Alloc()
	rv := objc.Send[PHASEDucker](instance.ID, objc.Sel("initWithEngine:sourceGroups:targetGroups:gain:attackTime:releaseTime:attackCurve:releaseCurve:"), engine, sourceGroups, targetGroups, gain, attackTime, releaseTime, attackCurve, releaseCurve)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEDuckerWithEngineSourceGroupsTargetGroupsGainAttackTimeReleaseTimeAttackCurveReleaseCurve */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEDucker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEDucker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEDucker */

// Instructs the ducker to begin altering sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/activate()
func (p_ PHASEDucker) Activate() {
	objc.Send[objc.ID](p_.ID, objc.Sel("activate"))
}/* debug [instance_methods/method]: Activate */


// Stops the ducker from altering sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/deactivate()
func (p_ PHASEDucker) Deactivate() {
	objc.Send[objc.ID](p_.ID, objc.Sel("deactivate"))
}/* debug [instance_methods/method]: Deactivate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEDucker */

// A mathematical curve that shapes transition progress as sound reduction begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/attackCurve
func (p_ PHASEDucker) AttackCurve() PHASECurveType {
	rv := objc.Send[PHASECurveType](p_.ID, objc.Sel("attackCurve"))
	return rv
}/* debug [instance_properties/getter]: attackCurve */


// The amount of time for sound reduction to reach maximum strength.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/attackTime
func (p_ PHASEDucker) AttackTime() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("attackTime"))
	return rv
}/* debug [instance_properties/getter]: attackTime */


// The amount of volume reduction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/gain
func (p_ PHASEDucker) Gain() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("gain"))
	return rv
}/* debug [instance_properties/getter]: gain */


// A unique value for the ducker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/identifier
func (p_ PHASEDucker) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A Boolean value that determines whether the ducker reduces sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/isActive
func (p_ PHASEDucker) Active() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// A mathematical curve that shapes transition progress as sound reduction ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/releaseCurve
func (p_ PHASEDucker) ReleaseCurve() PHASECurveType {
	rv := objc.Send[PHASECurveType](p_.ID, objc.Sel("releaseCurve"))
	return rv
}/* debug [instance_properties/getter]: releaseCurve */


// The amount of time to transition from maximum sound reduction to no reduction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/releaseTime
func (p_ PHASEDucker) ReleaseTime() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("releaseTime"))
	return rv
}/* debug [instance_properties/getter]: releaseTime */


// The sounds that determine volume reduction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/sourceGroups
func (p_ PHASEDucker) SourceGroups() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("sourceGroups"))
	return rv
}/* debug [instance_properties/getter]: sourceGroups */


// The sounds that reduce in volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDucker/targetGroups
func (p_ PHASEDucker) TargetGroups() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("targetGroups"))
	return rv
}/* debug [instance_properties/getter]: targetGroups */


// A Boolean value that determines whether the ducker reduces sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseducker/isactive
func (p_ PHASEDucker) IsActive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean value that determines whether the ducker reduces sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseducker/isactive
func (p_ PHASEDucker) SetIsActive(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEDucker */


