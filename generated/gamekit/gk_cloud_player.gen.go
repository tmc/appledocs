// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class GKCloudPlayer */


/* debug [class_header]: Header for GKCloudPlayer */
// The class instance for the [CloudPlayer] class.
var (
	CloudPlayerClass     _CloudPlayerClass
	CloudPlayerClassOnce sync.Once
)

func getCloudPlayerClass() _CloudPlayerClass {
	CloudPlayerClassOnce.Do(func() {
		CloudPlayerClass = _CloudPlayerClass{objc.GetClass("GKCloudPlayer")}
	})
	return CloudPlayerClass
}

type _CloudPlayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CloudPlayer */
// An interface definition for the [CloudPlayer] class.
type ICloudPlayer interface {
	IBasePlayer
	
/* debug [class_interface_properties]: Properties for CloudPlayer */
	// properties:
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CloudPlayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CloudPlayer */
// Alloc allocates a new instance without initialization.
func (cc _CloudPlayerClass) Alloc() CloudPlayer {
	rv := objc.Send[CloudPlayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CloudPlayerClass) New() CloudPlayer {
	rv := objc.Send[CloudPlayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CloudPlayer) Init() CloudPlayer {
	rv := objc.Send[CloudPlayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CloudPlayer) Autorelease() CloudPlayer {
	rv := objc.Send[CloudPlayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCloudPlayer creates a new CloudPlayer instance.
func NewCloudPlayer() CloudPlayer {
	return getCloudPlayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CloudPlayer */
// The object representing the currently signed-in iCloud user.


// The object representing the currently signed-in iCloud user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKCloudPlayer
type CloudPlayer struct {
	BasePlayer
}

// CloudPlayerFrom constructs a [CloudPlayer] from an unsafe.Pointer.
//
// The object representing the currently signed-in iCloud user.
func CloudPlayerFrom(ptr unsafe.Pointer) CloudPlayer {
	return CloudPlayer{
		BasePlayer: BasePlayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CloudPlayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CloudPlayer */

// Returns player information for the currently signed-in player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKCloudPlayer/getCurrentSignedInPlayer(forContainer:completionHandler:)
func (cc _CloudPlayerClass) GetCurrentSignedInPlayerForContainerCompletionHandler(containerName objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("getCurrentSignedInPlayerForContainer:completionHandler:"), containerName, handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetCurrentSignedInPlayerForContainerCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CloudPlayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CloudPlayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CloudPlayer */

// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (c_ CloudPlayer) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (c_ CloudPlayer) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKCloudPlayer */



