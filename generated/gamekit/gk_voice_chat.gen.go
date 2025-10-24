// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKVoiceChat */


/* debug [class_header]: Header for GKVoiceChat */
// The class instance for the [VoiceChat] class.
var (
	VoiceChatClass     _VoiceChatClass
	VoiceChatClassOnce sync.Once
)

func getVoiceChatClass() _VoiceChatClass {
	VoiceChatClassOnce.Do(func() {
		VoiceChatClass = _VoiceChatClass{objc.GetClass("GKVoiceChat")}
	})
	return VoiceChatClass
}

type _VoiceChatClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VoiceChat */
// An interface definition for the [VoiceChat] class.
type IVoiceChat interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VoiceChat */
	// properties:
	Active() bool
	SetActive(value bool)
	Name() objc.IObject /* cross-framework: NSString */
	PlayerIDs() []string
	Players() []Player
	PlayerStateUpdateHandler() unsafe.Pointer
	SetPlayerStateUpdateHandler(value unsafe.Pointer)
	PlayerVoiceChatStateDidChangeHandler() func(unsafe.Pointer, unsafe.Pointer)
	SetPlayerVoiceChatStateDidChangeHandler(value func(unsafe.Pointer, unsafe.Pointer))
	Volume() float32
	SetVolume(value float32)
	IsActive() bool
	SetIsActive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VoiceChat */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VoiceChat */
// Alloc allocates a new instance without initialization.
func (vc _VoiceChatClass) Alloc() VoiceChat {
	rv := objc.Send[VoiceChat](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VoiceChatClass) New() VoiceChat {
	rv := objc.Send[VoiceChat](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VoiceChat) Init() VoiceChat {
	rv := objc.Send[VoiceChat](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VoiceChat) Autorelease() VoiceChat {
	rv := objc.Send[VoiceChat](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVoiceChat creates a new VoiceChat instance.
func NewVoiceChat() VoiceChat {
	return getVoiceChatClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VoiceChat */
// A voice channel that allows players to speak with each other in a multiplayer game.
//
// GameKit provides the underlying mechanism to implement voice chat between players in a multiplayer game. It’s your responsibility to provide player controls and display feedback during the chat. First, configure voice chat by adding the key to the Information Property List and creating an audio session. Then, create a object using the method passing a string that identifies the voice channel. Use the method to connect players to the channel. Use the property to activate the microphone or switch the microphone between channels. Provide a handler using the property to update the interface when a player connects, speaks, or disconnects from a chat. You can also add controls that mute and set the volume using the method and property. Note that if there’s insufficient bandwidth over Wi-Fi to maintain a voice chat, GameKit may disconnect players from the channel or disband a channel.


// A voice channel that allows players to speak with each other in a multiplayer game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat
type VoiceChat struct {
	objectivec.Object
}

// VoiceChatFrom constructs a [VoiceChat] from an unsafe.Pointer.
//
// A voice channel that allows players to speak with each other in a multiplayer game.
func VoiceChatFrom(ptr unsafe.Pointer) VoiceChat {
	return VoiceChat{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VoiceChat *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VoiceChat */

// Returns whether voice chat is available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/isVoIPAllowed()
func (vc _VoiceChatClass) IsVoIPAllowed() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("isVoIPAllowed"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsVoIPAllowed) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VoiceChat */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VoiceChat */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VoiceChat */

// A Boolean value that indicates whether the channel is sampling the microphone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/isActive
func (v_ VoiceChat) Active() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// A Boolean value that indicates whether the channel is sampling the microphone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/isActive
func (v_ VoiceChat) SetActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setActive:"), value)
}/* debug [instance_properties/setter]: active */


// The name of the voice chat channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/name
func (v_ VoiceChat) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// An array of strings containing the player identifiers for the players connected to the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/playerIDs
func (v_ VoiceChat) PlayerIDs() []string {
	rv := objc.Send[[]string](v_.ID, objc.Sel("playerIDs"))
	return rv
}/* debug [instance_properties/getter]: playerIDs */


// The players connected to the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/players
func (v_ VoiceChat) Players() []Player {
	rv := objc.Send[[]Player](v_.ID, objc.Sel("players"))
	return rv
}/* debug [instance_properties/getter]: players */


// Handles when a player in the chat changes state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/playerStateUpdateHandler
func (v_ VoiceChat) PlayerStateUpdateHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("playerStateUpdateHandler"))
	return rv
}/* debug [instance_properties/getter]: playerStateUpdateHandler */


// Handles when a player in the chat changes state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/playerStateUpdateHandler
func (v_ VoiceChat) SetPlayerStateUpdateHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPlayerStateUpdateHandler:"), value)
}/* debug [instance_properties/setter]: playerStateUpdateHandler */


// A method that handles when a player’s voice chat changes state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/playerVoiceChatStateDidChangeHandler
func (v_ VoiceChat) PlayerVoiceChatStateDidChangeHandler() func(unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer)](v_.ID, objc.Sel("playerVoiceChatStateDidChangeHandler"))
	return rv
}/* debug [instance_properties/getter]: playerVoiceChatStateDidChangeHandler */


// A method that handles when a player’s voice chat changes state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/playerVoiceChatStateDidChangeHandler
func (v_ VoiceChat) SetPlayerVoiceChatStateDidChangeHandler(value func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPlayerVoiceChatStateDidChangeHandler:"), value)
}/* debug [instance_properties/setter]: playerVoiceChatStateDidChangeHandler */


// The volume level for the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/volume
func (v_ VoiceChat) Volume() float32 {
	rv := objc.Send[float32](v_.ID, objc.Sel("volume"))
	return rv
}/* debug [instance_properties/getter]: volume */


// The volume level for the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/volume
func (v_ VoiceChat) SetVolume(value float32) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVolume:"), value)
}/* debug [instance_properties/setter]: volume */


// A Boolean value that indicates whether the channel is sampling the microphone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/isactive
func (v_ VoiceChat) IsActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean value that indicates whether the channel is sampling the microphone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/isactive
func (v_ VoiceChat) SetIsActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKVoiceChat */



