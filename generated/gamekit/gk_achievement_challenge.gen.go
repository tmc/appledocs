// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GKAchievementChallenge */


/* debug [class_header]: Header for GKAchievementChallenge */
// The class instance for the [AchievementChallenge] class.
var (
	AchievementChallengeClass     _AchievementChallengeClass
	AchievementChallengeClassOnce sync.Once
)

func getAchievementChallengeClass() _AchievementChallengeClass {
	AchievementChallengeClassOnce.Do(func() {
		AchievementChallengeClass = _AchievementChallengeClass{objc.GetClass("GKAchievementChallenge")}
	})
	return AchievementChallengeClass
}

type _AchievementChallengeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AchievementChallenge */
// An interface definition for the [AchievementChallenge] class.
type IAchievementChallenge interface {
	IChallenge
	
/* debug [class_interface_properties]: Properties for AchievementChallenge */
	// properties:
	Achievement() IGKAchievement
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AchievementChallenge */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AchievementChallenge */
// Alloc allocates a new instance without initialization.
func (ac _AchievementChallengeClass) Alloc() AchievementChallenge {
	rv := objc.Send[AchievementChallenge](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AchievementChallengeClass) New() AchievementChallenge {
	rv := objc.Send[AchievementChallenge](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AchievementChallenge) Init() AchievementChallenge {
	rv := objc.Send[AchievementChallenge](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AchievementChallenge) Autorelease() AchievementChallenge {
	rv := objc.Send[AchievementChallenge](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAchievementChallenge creates a new AchievementChallenge instance.
func NewAchievementChallenge() AchievementChallenge {
	return getAchievementChallengeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AchievementChallenge */
// A type of challenge where a player must earn another player’s achievement.


// A type of challenge where a player must earn another player’s achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementChallenge
type AchievementChallenge struct {
	Challenge
}

// AchievementChallengeFrom constructs a [AchievementChallenge] from an unsafe.Pointer.
//
// A type of challenge where a player must earn another player’s achievement.
func AchievementChallengeFrom(ptr unsafe.Pointer) AchievementChallenge {
	return AchievementChallenge{
		Challenge: ChallengeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AchievementChallenge *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AchievementChallenge */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AchievementChallenge */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AchievementChallenge */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AchievementChallenge */

// The achievement that the player must earn to complete the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementChallenge/achievement
func (a_ AchievementChallenge) Achievement() IGKAchievement {
	rv := objc.Send[Achievement](a_.ID, objc.Sel("achievement"))
	return rv
}/* debug [instance_properties/getter]: achievement */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (a_ AchievementChallenge) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (a_ AchievementChallenge) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKAchievementChallenge */



