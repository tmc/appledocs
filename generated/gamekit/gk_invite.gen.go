// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Invite] class.
var (
	InviteClass     _InviteClass
	InviteClassOnce sync.Once
)

func getInviteClass() _InviteClass {
	InviteClassOnce.Do(func() {
		InviteClass = _InviteClass{objc.GetClass("GKInvite")}
	})
	return InviteClass
}

type _InviteClass struct {
	class objc.Class
}

// An interface definition for the [Invite] class.
type IInvite interface {
	objectivec.IObject
	// properties:
	Hosted() bool /* primitive/slice/pointer. */
	PlayerAttributes() uint32 /* not a class type */
	PlayerGroup() uint /* primitive/slice/pointer. */
	Inviter() string /* primitive/slice/pointer. */
	SetInviter(value string /* primitive/slice/pointer. */)
	IsHosted() bool /* primitive/slice/pointer. */
	SetIsHosted(value bool /* primitive/slice/pointer. */)
	Sender() IGKPlayer
	SetSender(value IGKPlayer)
	// methods:
}

// An invitation to join a match sent to the local player from another player.
//
// Your game never directly creates objects. Instead, these objects are created by GameKit and delivered to your game’s matchmaking event handler. The properties of the invitation object describe the match to which another player invites the local player.


// An invitation to join a match sent to the local player from another player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInvite
type Invite struct {
	objectivec.Object
}

// InviteFrom constructs a [Invite] from an unsafe.Pointer.
//
// An invitation to join a match sent to the local player from another player.
func InviteFrom(ptr unsafe.Pointer) Invite {
	return Invite{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _InviteClass) Alloc() Invite {
	rv := objc.Send[Invite](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InviteClass) New() Invite {
	rv := objc.Send[Invite](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ Invite) Init() Invite {
	rv := objc.Send[Invite](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ Invite) Autorelease() Invite {
	rv := objc.Send[Invite](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInvite creates a new Invite instance.
func NewInvite() Invite {
	return getInviteClass().New()
}



// A Boolean value that indicates whether you host the game on your own servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInvite/isHosted
func (i_ Invite) Hosted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("hosted"))
	return rv
}


// The player attributes for the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInvite/playerAttributes
func (i_ Invite) PlayerAttributes() uint32 /* not a class type */ {
	rv := objc.Send[uint32](i_.ID, objc.Sel("playerAttributes"))
	return rv
}


// The player group for the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInvite/playerGroup
func (i_ Invite) PlayerGroup() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("playerGroup"))
	return rv
}


// The identifier for the player who sends the invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkinvite/inviter
func (i_ Invite) Inviter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("inviter"))
	return rv
}


// The identifier for the player who sends the invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkinvite/inviter
func (i_ Invite) SetInviter(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInviter:"), objc.String(value))
}


// A Boolean value that indicates whether you host the game on your own servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkinvite/ishosted
func (i_ Invite) IsHosted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](i_.ID, objc.Sel("isHosted"))
	return rv
}


// A Boolean value that indicates whether you host the game on your own servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkinvite/ishosted
func (i_ Invite) SetIsHosted(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsHosted:"), value)
}


// The player who sends the invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkinvite/sender
func (i_ Invite) Sender() IGKPlayer {
	rv := objc.Send[Player](i_.ID, objc.Sel("sender"))
	return rv
}


// The player who sends the invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkinvite/sender
func (i_ Invite) SetSender(value IGKPlayer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSender:"), value)
}



