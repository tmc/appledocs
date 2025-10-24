// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKTurnBasedParticipant */


/* debug [class_header]: Header for GKTurnBasedParticipant */
// The class instance for the [TurnBasedParticipant] class.
var (
	TurnBasedParticipantClass     _TurnBasedParticipantClass
	TurnBasedParticipantClassOnce sync.Once
)

func getTurnBasedParticipantClass() _TurnBasedParticipantClass {
	TurnBasedParticipantClassOnce.Do(func() {
		TurnBasedParticipantClass = _TurnBasedParticipantClass{objc.GetClass("GKTurnBasedParticipant")}
	})
	return TurnBasedParticipantClass
}

type _TurnBasedParticipantClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TurnBasedParticipant */
// An interface definition for the [TurnBasedParticipant] class.
type ITurnBasedParticipant interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TurnBasedParticipant */
	// properties:
	LastTurnDate() objc.IObject /* cross-framework: NSDate */
	MatchOutcome() TurnBasedMatchOutcome
	SetMatchOutcome(value TurnBasedMatchOutcome)
	Player() IGKPlayer
	PlayerID() objc.IObject /* cross-framework: NSString */
	Status() TurnBasedParticipantStatus
	TimeoutDate() objc.IObject /* cross-framework: NSDate */
	Participants() IGKTurnBasedParticipant
	SetParticipants(value IGKTurnBasedParticipant)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TurnBasedParticipant */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TurnBasedParticipant */
// Alloc allocates a new instance without initialization.
func (tc _TurnBasedParticipantClass) Alloc() TurnBasedParticipant {
	rv := objc.Send[TurnBasedParticipant](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TurnBasedParticipantClass) New() TurnBasedParticipant {
	rv := objc.Send[TurnBasedParticipant](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TurnBasedParticipant) Init() TurnBasedParticipant {
	rv := objc.Send[TurnBasedParticipant](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TurnBasedParticipant) Autorelease() TurnBasedParticipant {
	rv := objc.Send[TurnBasedParticipant](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTurnBasedParticipant creates a new TurnBasedParticipant instance.
func NewTurnBasedParticipant() TurnBasedParticipant {
	return getTurnBasedParticipantClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TurnBasedParticipant */
// A participant in a turn-based match.
//
// A represents a player in a turn-based match that Game Center uses to store and forward match data. In your game, use participant objects to show information about opponents during gameplay. You get objects from the property of a object that GameKit passes to protocol methods. If a participant represents a filled slot in the match, GameKit sets the property and the accordingly. Get more information about a participant, such as the participant’s name and avatar, through the property. Before you end a match, you must set the property for every participant in the match.


// A participant in a turn-based match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant
type TurnBasedParticipant struct {
	objectivec.Object
}

// TurnBasedParticipantFrom constructs a [TurnBasedParticipant] from an unsafe.Pointer.
//
// A participant in a turn-based match.
func TurnBasedParticipantFrom(ptr unsafe.Pointer) TurnBasedParticipant {
	return TurnBasedParticipant{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TurnBasedParticipant *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TurnBasedParticipant */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TurnBasedParticipant */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TurnBasedParticipant */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TurnBasedParticipant */

// The date and time that this participant last took a turn in the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/lastTurnDate
func (t_ TurnBasedParticipant) LastTurnDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("lastTurnDate"))
	return rv
}/* debug [instance_properties/getter]: lastTurnDate */


// The conclusion or results of a participant in a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/matchOutcome
func (t_ TurnBasedParticipant) MatchOutcome() TurnBasedMatchOutcome {
	rv := objc.Send[TurnBasedMatchOutcome](t_.ID, objc.Sel("matchOutcome"))
	return rv
}/* debug [instance_properties/getter]: matchOutcome */


// The conclusion or results of a participant in a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/matchOutcome
func (t_ TurnBasedParticipant) SetMatchOutcome(value TurnBasedMatchOutcome) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMatchOutcome:"), value)
}/* debug [instance_properties/setter]: matchOutcome */


// The player object containing the participant details.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/player
func (t_ TurnBasedParticipant) Player() IGKPlayer {
	rv := objc.Send[Player](t_.ID, objc.Sel("player"))
	return rv
}/* debug [instance_properties/getter]: player */


// The player identifier for this participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/playerID
func (t_ TurnBasedParticipant) PlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("playerID"))
	return rv
}/* debug [instance_properties/getter]: playerID */


// The status of the participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/status-swift.property
func (t_ TurnBasedParticipant) Status() TurnBasedParticipantStatus {
	rv := objc.Send[TurnBasedParticipantStatus](t_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// The date and time that the participant’s turn timed out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/timeoutDate
func (t_ TurnBasedParticipant) TimeoutDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("timeoutDate"))
	return rv
}/* debug [instance_properties/getter]: timeoutDate */


// The players that participate in a turn-based match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/participants
func (t_ TurnBasedParticipant) Participants() IGKTurnBasedParticipant {
	rv := objc.Send[TurnBasedParticipant](t_.ID, objc.Sel("participants"))
	return rv
}/* debug [instance_properties/getter]: participants */


// The players that participate in a turn-based match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/participants
func (t_ TurnBasedParticipant) SetParticipants(value IGKTurnBasedParticipant) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParticipants:"), value)
}/* debug [instance_properties/setter]: participants */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKTurnBasedParticipant */



