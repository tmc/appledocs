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
}

// An invitation to join a match sent to the local player from another player.
//
// Your game never directly creates objects. Instead, these objects are created by GameKit and delivered to your game’s matchmaking event handler. The properties of the invitation object describe the match to which another player invites the local player.
//
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


// The player group for the match.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInvite/playerGroup
func (i_ Invite) PlayerGroup() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("playerGroup"))
	return rv
}



