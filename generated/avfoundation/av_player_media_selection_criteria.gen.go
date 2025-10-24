// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerMediaSelectionCriteria */


/* debug [class_header]: Header for AVPlayerMediaSelectionCriteria */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerMediaSelectionCriteria */
// An interface definition for the [PlayerMediaSelectionCriteria] class.
type IPlayerMediaSelectionCriteria interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerMediaSelectionCriteria */
	// properties:
	PreferredLanguages() []string
	PreferredMediaCharacteristics() []string
	PrincipalMediaCharacteristics() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerMediaSelectionCriteria */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerMediaSelectionCriteria */
// Alloc allocates a new instance without initialization.
func (pc _PlayerMediaSelectionCriteriaClass) Alloc() PlayerMediaSelectionCriteria {
	rv := objc.Send[PlayerMediaSelectionCriteria](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerMediaSelectionCriteria */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerMediaSelectionCriteria */

// Creates media selection criteria with the preferred languages and media characteristics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/init(preferredLanguages:preferredMediaCharacteristics:)
func NewPlayerMediaSelectionCriteriaWithPreferredLanguagesPreferredMediaCharacteristics(preferredLanguages []string, preferredMediaCharacteristics []string) PlayerMediaSelectionCriteria {
	instance := getPlayerMediaSelectionCriteriaClass().Alloc()
	rv := objc.Send[PlayerMediaSelectionCriteria](instance.ID, objc.Sel("initWithPreferredLanguages:preferredMediaCharacteristics:"), preferredLanguages, preferredMediaCharacteristics)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerMediaSelectionCriteriaWithPreferredLanguagesPreferredMediaCharacteristics */


// Creates media selection criteria with the principal media characteristics, and preferred languages and media characteristics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/init(principalMediaCharacteristics:preferredLanguages:preferredMediaCharacteristics:)
func NewPlayerMediaSelectionCriteriaWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics(principalMediaCharacteristics []string, preferredLanguages []string, preferredMediaCharacteristics []string) PlayerMediaSelectionCriteria {
	instance := getPlayerMediaSelectionCriteriaClass().Alloc()
	rv := objc.Send[PlayerMediaSelectionCriteria](instance.ID, objc.Sel("initWithPrincipalMediaCharacteristics:preferredLanguages:preferredMediaCharacteristics:"), principalMediaCharacteristics, preferredLanguages, preferredMediaCharacteristics)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerMediaSelectionCriteriaWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerMediaSelectionCriteria */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerMediaSelectionCriteria */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerMediaSelectionCriteria */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerMediaSelectionCriteria */

// An array of language identifiers in preferred order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/preferredLanguages
func (p_ PlayerMediaSelectionCriteria) PreferredLanguages() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("preferredLanguages"))
	return rv
}/* debug [instance_properties/getter]: preferredLanguages */


// An array of media characteristics in preferred order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/preferredMediaCharacteristics
func (p_ PlayerMediaSelectionCriteria) PreferredMediaCharacteristics() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("preferredMediaCharacteristics"))
	return rv
}/* debug [instance_properties/getter]: preferredMediaCharacteristics */


// An array of media characteristics that are essential to select when choosing media with a particular characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/principalMediaCharacteristics
func (p_ PlayerMediaSelectionCriteria) PrincipalMediaCharacteristics() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("principalMediaCharacteristics"))
	return rv
}/* debug [instance_properties/getter]: principalMediaCharacteristics */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerMediaSelectionCriteria */


