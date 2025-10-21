// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [TurnBasedParticipant] class.
type ITurnBasedParticipant interface {
	objectivec.IObject
}

// A participant in a turn-based match.
//
// A represents a player in a turn-based match that Game Center uses to store and forward match data. In your game, use participant objects to show information about opponents during gameplay. You get objects from the property of a object that GameKit passes to protocol methods. If a participant represents a filled slot in the match, GameKit sets the property and the accordingly. Get more information about a participant, such as the participant’s name and avatar, through the property. Before you end a match, you must set the property for every participant in the match.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TurnBasedParticipantClass) Alloc() TurnBasedParticipant {
	rv := objc.Send[TurnBasedParticipant](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The player object containing the participant details.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedparticipant/player
func (t_ TurnBasedParticipant) Player() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("player"))
	return rv
}


// SetPlayer sets the value of the player property.
// The player object containing the participant details.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedparticipant/player
func (t_ TurnBasedParticipant) SetPlayer(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlayer:"), value)
}

// The player identifier for this participant.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedparticipant/playerid
func (t_ TurnBasedParticipant) PlayerID() string {
	rv := objc.Send[string](t_.ID, objc.Sel("playerID"))
	return rv
}


// SetPlayerID sets the value of the playerID property.
// The player identifier for this participant.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedparticipant/playerid
func (t_ TurnBasedParticipant) SetPlayerID(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPlayerID:"), objc.String(value))
}

// The status of the participant.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedparticipant/status-swift.property
func (t_ TurnBasedParticipant) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
// The status of the participant.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedparticipant/status-swift.property
func (t_ TurnBasedParticipant) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStatus:"), value)
}

// The players that participate in a turn-based match.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/participants
func (t_ TurnBasedParticipant) Participants() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("participants"))
	return rv
}


// SetParticipants sets the value of the participants property.
// The players that participate in a turn-based match.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/participants
func (t_ TurnBasedParticipant) SetParticipants(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParticipants:"), value)
}

// The conclusion or results of a participant in a match.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedparticipant/matchoutcome
func (t_ TurnBasedParticipant) MatchOutcome() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("matchOutcome"))
	return rv
}


// SetMatchOutcome sets the value of the matchOutcome property.
// The conclusion or results of a participant in a match.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedparticipant/matchoutcome
func (t_ TurnBasedParticipant) SetMatchOutcome(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMatchOutcome:"), value)
}

// The date and time that this participant last took a turn in the game.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedparticipant/lastturndate
func (t_ TurnBasedParticipant) LastTurnDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("lastTurnDate"))
	return rv
}


// SetLastTurnDate sets the value of the lastTurnDate property.
// The date and time that this participant last took a turn in the game.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedparticipant/lastturndate
func (t_ TurnBasedParticipant) SetLastTurnDate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLastTurnDate:"), value)
}

// The date and time that the participant’s turn timed out.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/timeoutDate
func (t_ TurnBasedParticipant) TimeoutDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("timeoutDate"))
	return rv
}



