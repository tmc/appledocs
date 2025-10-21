// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PHASEGroup] class.
type IPHASEGroup interface {
	objectivec.IObject
	FadeGainDurationCurveType(gain unsafe.Pointer, duration unsafe.Pointer, curveType unsafe.Pointer)
	FadeRateDurationCurveType(rate unsafe.Pointer, duration unsafe.Pointer, curveType unsafe.Pointer)
	Mute()
	RegisterWithEngine(engine unsafe.Pointer)
	Solo()
	Unmute()
	UnregisterFromEngine()
	Unsolo()
}

// A container that shares audio parameters with a collection of sounds.
//
// With all the sounds it contains, a group shares settings like gain, playback rate, mute, and solo. Groups are nonhierarchical and don’t overlap — that is, each sound event associates with only one group.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASEGroupClass) Alloc() PHASEGroup {
	rv := objc.Send[PHASEGroup](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a group with a unique name.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/init(identifier:)
func NewPHASEGroupWithIdentifier(identifier string) PHASEGroup {
	instance := getPHASEGroupClass().Alloc()
	rv := objc.Send[PHASEGroup](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	rv.Autorelease()
	return rv
}


// Adjusts the volume of the sounds in a group gradually.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/fadeGain(gain:duration:curveType:)
func (p_ PHASEGroup) FadeGainDurationCurveType(gain unsafe.Pointer, duration unsafe.Pointer, curveType unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("fadeGain:duration:curveType:"), gain, duration, curveType)
}

// Adjusts the playback speed of the sounds in a group gradually.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/fadeRate(rate:duration:curveType:)
func (p_ PHASEGroup) FadeRateDurationCurveType(rate unsafe.Pointer, duration unsafe.Pointer, curveType unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("fadeRate:duration:curveType:"), rate, duration, curveType)
}

// Silences the group.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/mute()
func (p_ PHASEGroup) Mute() {
	objc.Send[objc.ID](p_.ID, objc.Sel("mute"))
}

// Adds the group to the engine’s dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/register(engine:)
func (p_ PHASEGroup) RegisterWithEngine(engine unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("registerWithEngine:"), engine)
}

// Silences all other groups.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/solo()
func (p_ PHASEGroup) Solo() {
	objc.Send[objc.ID](p_.ID, objc.Sel("solo"))
}

// Restores the group’s volume.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/unmute()
func (p_ PHASEGroup) Unmute() {
	objc.Send[objc.ID](p_.ID, objc.Sel("unmute"))
}

// Removes the group from the engine’s dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/unregisterFromEngine()
func (p_ PHASEGroup) UnregisterFromEngine() {
	objc.Send[objc.ID](p_.ID, objc.Sel("unregisterFromEngine"))
}

// Restores the other groups’ volume.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/unsolo()
func (p_ PHASEGroup) Unsolo() {
	objc.Send[objc.ID](p_.ID, objc.Sel("unsolo"))
}

// A Boolean value that indicates whether the app silences the group.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegroup/ismuted
func (p_ PHASEGroup) IsMuted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isMuted"))
	return rv
}


// SetIsMuted sets the value of the isMuted property.
// A Boolean value that indicates whether the app silences the group.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegroup/ismuted
func (p_ PHASEGroup) SetIsMuted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsMuted:"), value)
}

// A Boolean value that indicates whether the app silences all groups other than this group.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegroup/issoloed
func (p_ PHASEGroup) IsSoloed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSoloed"))
	return rv
}


// SetIsSoloed sets the value of the isSoloed property.
// A Boolean value that indicates whether the app silences all groups other than this group.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegroup/issoloed
func (p_ PHASEGroup) SetIsSoloed(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSoloed:"), value)
}

// Modifies the volume of the group’s sounds.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/gain
func (p_ PHASEGroup) Gain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("gain"))
	return rv
}


// SetGain sets the value of the gain property.
// Modifies the volume of the group’s sounds.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/gain
func (p_ PHASEGroup) SetGain(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGain:"), value)
}

// A unique name for the group.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/identifier
func (p_ PHASEGroup) Identifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}

// A Boolean value that indicates whether the app silences the group.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/isMuted
func (p_ PHASEGroup) Muted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("muted"))
	return rv
}

// A Boolean value that indicates whether the app silences all groups other than this group.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/isSoloed
func (p_ PHASEGroup) Soloed() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("soloed"))
	return rv
}

// The group’s playback speed.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/rate
func (p_ PHASEGroup) Rate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
// The group’s playback speed.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroup/rate
func (p_ PHASEGroup) SetRate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRate:"), value)
}


