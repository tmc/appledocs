// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayerMediaSelectionCriteria] class.
var (
	PlayerMediaSelectionCriteriaClass     _PlayerMediaSelectionCriteriaClass
	PlayerMediaSelectionCriteriaClassOnce sync.Once
)

func getPlayerMediaSelectionCriteriaClass() _PlayerMediaSelectionCriteriaClass {
	PlayerMediaSelectionCriteriaClassOnce.Do(func() {
		PlayerMediaSelectionCriteriaClass = _PlayerMediaSelectionCriteriaClass{objc.GetClass("AVPlayerMediaSelectionCriteria")}
	})
	return PlayerMediaSelectionCriteriaClass
}

type _PlayerMediaSelectionCriteriaClass struct {
	class objc.Class
}

// An interface definition for the [PlayerMediaSelectionCriteria] class.
type IPlayerMediaSelectionCriteria interface {
	objectivec.IObject
	// properties:
	PreferredLanguages() objc.IObject /* cross-framework: NSString */
	SetPreferredLanguages(value objc.IObject /* cross-framework: NSString */)
	PreferredMediaCharacteristics() MediaCharacteristic /* not a class type */
	SetPreferredMediaCharacteristics(value MediaCharacteristic /* not a class type */)
	PrincipalMediaCharacteristics() MediaCharacteristic /* not a class type */
	SetPrincipalMediaCharacteristics(value MediaCharacteristic /* not a class type */)
	// methods:
}

// An object that specifies the preferred languages and media characteristics for a player.
//
// An instance of this object represents the languages and media characteristics of assets that contain media selection options that a player attempts to select automatically when preparing and playing items. It lists the languages and media characteristics in their preferred order.


// An object that specifies the preferred languages and media characteristics for a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria
type PlayerMediaSelectionCriteria struct {
	objectivec.Object
}

// PlayerMediaSelectionCriteriaFrom constructs a [PlayerMediaSelectionCriteria] from an unsafe.Pointer.
//
// An object that specifies the preferred languages and media characteristics for a player.
func PlayerMediaSelectionCriteriaFrom(ptr unsafe.Pointer) PlayerMediaSelectionCriteria {
	return PlayerMediaSelectionCriteria{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerMediaSelectionCriteriaClass) Alloc() PlayerMediaSelectionCriteria {
	rv := objc.Send[PlayerMediaSelectionCriteria](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerMediaSelectionCriteriaClass) New() PlayerMediaSelectionCriteria {
	rv := objc.Send[PlayerMediaSelectionCriteria](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerMediaSelectionCriteria) Init() PlayerMediaSelectionCriteria {
	rv := objc.Send[PlayerMediaSelectionCriteria](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerMediaSelectionCriteria) Autorelease() PlayerMediaSelectionCriteria {
	rv := objc.Send[PlayerMediaSelectionCriteria](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerMediaSelectionCriteria creates a new PlayerMediaSelectionCriteria instance.
func NewPlayerMediaSelectionCriteria() PlayerMediaSelectionCriteria {
	return getPlayerMediaSelectionCriteriaClass().New()
}



// An array of language identifiers in preferred order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria/preferredlanguages
func (p_ PlayerMediaSelectionCriteria) PreferredLanguages() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("preferredLanguages"))
	return rv
}


// An array of language identifiers in preferred order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria/preferredlanguages
func (p_ PlayerMediaSelectionCriteria) SetPreferredLanguages(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredLanguages:"), value)
}


// An array of media characteristics in preferred order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria/preferredmediacharacteristics
func (p_ PlayerMediaSelectionCriteria) PreferredMediaCharacteristics() MediaCharacteristic /* not a class type */ {
	rv := objc.Send[MediaCharacteristic](p_.ID, objc.Sel("preferredMediaCharacteristics"))
	return rv
}


// An array of media characteristics in preferred order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria/preferredmediacharacteristics
func (p_ PlayerMediaSelectionCriteria) SetPreferredMediaCharacteristics(value MediaCharacteristic /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredMediaCharacteristics:"), value)
}


// An array of media characteristics that are essential to select when choosing media with a particular characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria/principalmediacharacteristics
func (p_ PlayerMediaSelectionCriteria) PrincipalMediaCharacteristics() MediaCharacteristic /* not a class type */ {
	rv := objc.Send[MediaCharacteristic](p_.ID, objc.Sel("principalMediaCharacteristics"))
	return rv
}


// An array of media characteristics that are essential to select when choosing media with a particular characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayermediaselectioncriteria/principalmediacharacteristics
func (p_ PlayerMediaSelectionCriteria) SetPrincipalMediaCharacteristics(value MediaCharacteristic /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrincipalMediaCharacteristics:"), value)
}



