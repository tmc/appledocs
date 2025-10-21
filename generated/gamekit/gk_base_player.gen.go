// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BasePlayer] class.
var (
	BasePlayerClass     _BasePlayerClass
	BasePlayerClassOnce sync.Once
)

func getBasePlayerClass() _BasePlayerClass {
	BasePlayerClassOnce.Do(func() {
		BasePlayerClass = _BasePlayerClass{objc.GetClass("GKBasePlayer")}
	})
	return BasePlayerClass
}

type _BasePlayerClass struct {
	class objc.Class
}

// An interface definition for the [BasePlayer] class.
type IBasePlayer interface {
	objectivec.IObject
}

// A class that provides common data and methods for the different player objects.
//
// is the abstract superclass for the classes that represent the local player running your app and remote players who may join their games. Use the subclass to initialize the local player who runs your app on their device. Then you can access the local player’s nickname, avatar, leaderboards, and achievements. You can also invite other players ( objects), and send information between players.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKBasePlayer
type BasePlayer struct {
	objectivec.Object
}

// BasePlayerFrom constructs a [BasePlayer] from an unsafe.Pointer.
//
// A class that provides common data and methods for the different player objects.
func BasePlayerFrom(ptr unsafe.Pointer) BasePlayer {
	return BasePlayer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BasePlayerClass) Alloc() BasePlayer {
	rv := objc.Send[BasePlayer](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BasePlayerClass) New() BasePlayer {
	rv := objc.Send[BasePlayer](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BasePlayer) Init() BasePlayer {
	rv := objc.Send[BasePlayer](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BasePlayer) Autorelease() BasePlayer {
	rv := objc.Send[BasePlayer](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBasePlayer creates a new BasePlayer instance.
func NewBasePlayer() BasePlayer {
	return getBasePlayerClass().New()
}


// A unique identifier for a player.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKBasePlayer/playerID
func (b_ BasePlayer) PlayerID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("playerID"))
	return rv
}



