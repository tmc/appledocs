// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKInvite */


/* debug [class_header]: Header for GKInvite */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Invite */
// An interface definition for the [Invite] class.
type IInvite interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Invite */
	// properties:
	Inviter() objc.IObject /* cross-framework: NSString */
	Hosted() bool
	PlayerAttributes() uint32 /* not a class type */
	PlayerGroup() uint
	Sender() IGKPlayer
	IsHosted() bool
	SetIsHosted(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Invite */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Invite */
// Alloc allocates a new instance without initialization.
func (ic _InviteClass) Alloc() Invite {
	rv := objc.Send[Invite](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Invite */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Invite *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Invite */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Invite */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Invite */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Invite */

// The identifier for the player who sends the invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInvite/inviter
func (i_ Invite) Inviter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("inviter"))
	return rv
}/* debug [instance_properties/getter]: inviter */


// A Boolean value that indicates whether you host the game on your own servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInvite/isHosted
func (i_ Invite) Hosted() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hosted"))
	return rv
}/* debug [instance_properties/getter]: hosted */


// The player attributes for the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInvite/playerAttributes
func (i_ Invite) PlayerAttributes() uint32 /* not a class type */ {
	rv := objc.Send[uint32](i_.ID, objc.Sel("playerAttributes"))
	return rv
}/* debug [instance_properties/getter]: playerAttributes */


// The player group for the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInvite/playerGroup
func (i_ Invite) PlayerGroup() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("playerGroup"))
	return rv
}/* debug [instance_properties/getter]: playerGroup */


// The player who sends the invitation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInvite/sender
func (i_ Invite) Sender() IGKPlayer {
	rv := objc.Send[Player](i_.ID, objc.Sel("sender"))
	return rv
}/* debug [instance_properties/getter]: sender */


// A Boolean value that indicates whether you host the game on your own servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkinvite/ishosted
func (i_ Invite) IsHosted() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isHosted"))
	return rv
}/* debug [instance_properties/getter]: isHosted */


// A Boolean value that indicates whether you host the game on your own servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkinvite/ishosted
func (i_ Invite) SetIsHosted(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsHosted:"), value)
}/* debug [instance_properties/setter]: isHosted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKInvite */



