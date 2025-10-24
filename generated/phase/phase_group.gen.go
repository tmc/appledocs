// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEGroup */


/* debug [class_header]: Header for PHASEGroup */
// The class instance for the [PHASEGroup] class.
var (
	PHASEGroupClass     _PHASEGroupClass
	PHASEGroupClassOnce sync.Once
)

func getPHASEGroupClass() _PHASEGroupClass {
	PHASEGroupClassOnce.Do(func() {
		PHASEGroupClass = _PHASEGroupClass{objc.GetClass("PHASEGroup")}
	})
	return PHASEGroupClass
}

type _PHASEGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEGroup */
// An interface definition for the [PHASEGroup] class.
type IPHASEGroup interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEGroup */
	// properties:
	Gain() float64
	SetGain(value float64)
	Identifier() objc.IObject /* cross-framework: NSString */
	Muted() bool
	Soloed() bool
	Rate() float64
	SetRate(value float64)
	IsMuted() bool
	SetIsMuted(value bool)
	IsSoloed() bool
	SetIsSoloed(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEGroup */
	// methods:
	FadeGainDurationCurveType(gain float64, duration float64, curveType PHASECurveType)
	FadeRateDurationCurveType(rate float64, duration float64, curveType PHASECurveType)
	Mute()
	RegisterWithEngine(engine IPHASEEngine)
	Solo()
	Unmute()
	UnregisterFromEngine()
	Unsolo()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEGroup */
// Alloc allocates a new instance without initialization.
func (pc _PHASEGroupClass) Alloc() PHASEGroup {
	rv := objc.Send[PHASEGroup](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEGroupClass) New() PHASEGroup {
	rv := objc.Send[PHASEGroup](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEGroup) Init() PHASEGroup {
	rv := objc.Send[PHASEGroup](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEGroup) Autorelease() PHASEGroup {
	rv := objc.Send[PHASEGroup](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEGroup creates a new PHASEGroup instance.
func NewPHASEGroup() PHASEGroup {
	return getPHASEGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEGroup */
// A container that shares audio parameters with a collection of sounds.
//
// With all the sounds it contains, a group shares settings like gain, playback rate, mute, and solo. Groups are nonhierarchical and don’t overlap — that is, each sound event associates with only one group.


// A container that shares audio parameters with a collection of sounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup
type PHASEGroup struct {
	objectivec.Object
}

// PHASEGroupFrom constructs a [PHASEGroup] from an unsafe.Pointer.
//
// A container that shares audio parameters with a collection of sounds.
func PHASEGroupFrom(ptr unsafe.Pointer) PHASEGroup {
	return PHASEGroup{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEGroup */

// Creates a group with a unique name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/init(identifier:)
func NewPHASEGroupWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) PHASEGroup {
	instance := getPHASEGroupClass().Alloc()
	rv := objc.Send[PHASEGroup](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEGroupWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEGroup */

// Adjusts the volume of the sounds in a group gradually.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/fadeGain(gain:duration:curveType:)
func (p_ PHASEGroup) FadeGainDurationCurveType(gain float64, duration float64, curveType PHASECurveType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("fadeGain:duration:curveType:"), gain, duration, curveType)
}/* debug [instance_methods/method]: FadeGainDurationCurveType */


// Adjusts the playback speed of the sounds in a group gradually.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/fadeRate(rate:duration:curveType:)
func (p_ PHASEGroup) FadeRateDurationCurveType(rate float64, duration float64, curveType PHASECurveType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("fadeRate:duration:curveType:"), rate, duration, curveType)
}/* debug [instance_methods/method]: FadeRateDurationCurveType */


// Silences the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/mute()
func (p_ PHASEGroup) Mute() {
	objc.Send[objc.ID](p_.ID, objc.Sel("mute"))
}/* debug [instance_methods/method]: Mute */


// Adds the group to the engine’s dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/register(engine:)
func (p_ PHASEGroup) RegisterWithEngine(engine IPHASEEngine) {
	objc.Send[objc.ID](p_.ID, objc.Sel("registerWithEngine:"), engine)
}/* debug [instance_methods/method]: RegisterWithEngine */


// Silences all other groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/solo()
func (p_ PHASEGroup) Solo() {
	objc.Send[objc.ID](p_.ID, objc.Sel("solo"))
}/* debug [instance_methods/method]: Solo */


// Restores the group’s volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/unmute()
func (p_ PHASEGroup) Unmute() {
	objc.Send[objc.ID](p_.ID, objc.Sel("unmute"))
}/* debug [instance_methods/method]: Unmute */


// Removes the group from the engine’s dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/unregisterFromEngine()
func (p_ PHASEGroup) UnregisterFromEngine() {
	objc.Send[objc.ID](p_.ID, objc.Sel("unregisterFromEngine"))
}/* debug [instance_methods/method]: UnregisterFromEngine */


// Restores the other groups’ volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/unsolo()
func (p_ PHASEGroup) Unsolo() {
	objc.Send[objc.ID](p_.ID, objc.Sel("unsolo"))
}/* debug [instance_methods/method]: Unsolo */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEGroup */

// Modifies the volume of the group’s sounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/gain
func (p_ PHASEGroup) Gain() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("gain"))
	return rv
}/* debug [instance_properties/getter]: gain */


// Modifies the volume of the group’s sounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/gain
func (p_ PHASEGroup) SetGain(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGain:"), value)
}/* debug [instance_properties/setter]: gain */


// A unique name for the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/identifier
func (p_ PHASEGroup) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A Boolean value that indicates whether the app silences the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/isMuted
func (p_ PHASEGroup) Muted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("muted"))
	return rv
}/* debug [instance_properties/getter]: muted */


// A Boolean value that indicates whether the app silences all groups other than this group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/isSoloed
func (p_ PHASEGroup) Soloed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("soloed"))
	return rv
}/* debug [instance_properties/getter]: soloed */


// The group’s playback speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/rate
func (p_ PHASEGroup) Rate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// The group’s playback speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/rate
func (p_ PHASEGroup) SetRate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRate:"), value)
}/* debug [instance_properties/setter]: rate */


// A Boolean value that indicates whether the app silences the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegroup/ismuted
func (p_ PHASEGroup) IsMuted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isMuted"))
	return rv
}/* debug [instance_properties/getter]: isMuted */


// A Boolean value that indicates whether the app silences the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegroup/ismuted
func (p_ PHASEGroup) SetIsMuted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsMuted:"), value)
}/* debug [instance_properties/setter]: isMuted */


// A Boolean value that indicates whether the app silences all groups other than this group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegroup/issoloed
func (p_ PHASEGroup) IsSoloed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSoloed"))
	return rv
}/* debug [instance_properties/getter]: isSoloed */


// A Boolean value that indicates whether the app silences all groups other than this group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegroup/issoloed
func (p_ PHASEGroup) SetIsSoloed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSoloed:"), value)
}/* debug [instance_properties/setter]: isSoloed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEGroup */


