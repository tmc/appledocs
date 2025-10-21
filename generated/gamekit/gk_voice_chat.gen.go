// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [VoiceChat] class.
type IVoiceChat interface {
	objectivec.IObject
	Stop()
}

// A voice channel that allows players to speak with each other in a multiplayer game.
//
// GameKit provides the underlying mechanism to implement voice chat between players in a multiplayer game. It’s your responsibility to provide player controls and display feedback during the chat. First, configure voice chat by adding the key to the Information Property List and creating an audio session. Then, create a object using the method passing a string that identifies the voice channel. Use the method to connect players to the channel. Use the property to activate the microphone or switch the microphone between channels. Provide a handler using the property to update the interface when a player connects, speaks, or disconnects from a chat. You can also add controls that mute and set the volume using the method and property. Note that if there’s insufficient bandwidth over Wi-Fi to maintain a voice chat, GameKit may disconnect players from the channel or disband a channel.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VoiceChatClass) Alloc() VoiceChat {
	rv := objc.Send[VoiceChat](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Ends communication with other players in a channel.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/stop()
func (v_ VoiceChat) Stop() {
	objc.Send[objc.ID](v_.ID, objc.Sel("stop"))
}

// The players connected to the channel.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/players
func (v_ VoiceChat) Players() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("players"))
	return rv
}


// SetPlayers sets the value of the players property.
// The players connected to the channel.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/players
func (v_ VoiceChat) SetPlayers(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPlayers:"), value)
}

// Handles when a player in the chat changes state.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/playerstateupdatehandler
func (v_ VoiceChat) PlayerStateUpdateHandler() string {
	rv := objc.Send[string](v_.ID, objc.Sel("playerStateUpdateHandler"))
	return rv
}


// SetPlayerStateUpdateHandler sets the value of the playerStateUpdateHandler property.
// Handles when a player in the chat changes state.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/playerstateupdatehandler
func (v_ VoiceChat) SetPlayerStateUpdateHandler(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPlayerStateUpdateHandler:"), objc.String(value))
}

// The name of the voice chat channel.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/name
func (v_ VoiceChat) Name() string {
	rv := objc.Send[string](v_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the voice chat channel.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/name
func (v_ VoiceChat) SetName(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setName:"), objc.String(value))
}

// The volume level for the channel.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/volume
func (v_ VoiceChat) Volume() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("volume"))
	return rv
}


// SetVolume sets the value of the volume property.
// The volume level for the channel.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/volume
func (v_ VoiceChat) SetVolume(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVolume:"), value)
}

// A Boolean value that indicates whether the channel is sampling the microphone.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/isactive
func (v_ VoiceChat) IsActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isActive"))
	return rv
}


// SetIsActive sets the value of the isActive property.
// A Boolean value that indicates whether the channel is sampling the microphone.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/isactive
func (v_ VoiceChat) SetIsActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsActive:"), value)
}

// An array of strings containing the player identifiers for the players connected to the channel.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/playerids
func (v_ VoiceChat) PlayerIDs() string {
	rv := objc.Send[string](v_.ID, objc.Sel("playerIDs"))
	return rv
}


// SetPlayerIDs sets the value of the playerIDs property.
// An array of strings containing the player identifiers for the players connected to the channel.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/playerids
func (v_ VoiceChat) SetPlayerIDs(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPlayerIDs:"), objc.String(value))
}

// A method that handles when a player’s voice chat changes state.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/playervoicechatstatedidchangehandler
func (v_ VoiceChat) PlayerVoiceChatStateDidChangeHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("playerVoiceChatStateDidChangeHandler"))
	return rv
}


// SetPlayerVoiceChatStateDidChangeHandler sets the value of the playerVoiceChatStateDidChangeHandler property.
// A method that handles when a player’s voice chat changes state.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechat/playervoicechatstatedidchangehandler
func (v_ VoiceChat) SetPlayerVoiceChatStateDidChangeHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPlayerVoiceChatStateDidChangeHandler:"), value)
}



