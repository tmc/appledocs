// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKBasePlayer */


/* debug [class_header]: Header for GKBasePlayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BasePlayer */
// An interface definition for the [BasePlayer] class.
type IBasePlayer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BasePlayer */
	// properties:
	DisplayName() objc.IObject /* cross-framework: NSString */
	PlayerID() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BasePlayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BasePlayer */
// Alloc allocates a new instance without initialization.
func (bc _BasePlayerClass) Alloc() BasePlayer {
	rv := objc.Send[BasePlayer](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BasePlayer */
// A class that provides common data and methods for the different player objects.
//
// is the abstract superclass for the classes that represent the local player running your app and remote players who may join their games. Use the subclass to initialize the local player who runs your app on their device. Then you can access the local player’s nickname, avatar, leaderboards, and achievements. You can also invite other players ( objects), and send information between players.


// A class that provides common data and methods for the different player objects.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BasePlayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BasePlayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BasePlayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BasePlayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BasePlayer */

// The Game Center profile name for a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKBasePlayer/displayName
func (b_ BasePlayer) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// A unique identifier for a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKBasePlayer/playerID
func (b_ BasePlayer) PlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("playerID"))
	return rv
}/* debug [instance_properties/getter]: playerID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKBasePlayer */



