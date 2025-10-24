// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKVoiceChatService */


/* debug [class_header]: Header for GKVoiceChatService */
// The class instance for the [VoiceChatService] class.
var (
	VoiceChatServiceClass     _VoiceChatServiceClass
	VoiceChatServiceClassOnce sync.Once
)

func getVoiceChatServiceClass() _VoiceChatServiceClass {
	VoiceChatServiceClassOnce.Do(func() {
		VoiceChatServiceClass = _VoiceChatServiceClass{objc.GetClass("GKVoiceChatService")}
	})
	return VoiceChatServiceClass
}

type _VoiceChatServiceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VoiceChatService */
// An interface definition for the [VoiceChatService] class.
type IVoiceChatService interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VoiceChatService */
	// properties:
	KAudioSessionCategory_PlayAndRecord() int
	SetKAudioSessionCategory_PlayAndRecord(value int)
	IsInputMeteringEnabled() bool
	SetIsInputMeteringEnabled(value bool)
	IsMicrophoneMuted() bool
	SetIsMicrophoneMuted(value bool)
	IsOutputMeteringEnabled() bool
	SetIsOutputMeteringEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VoiceChatService */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VoiceChatService */
// Alloc allocates a new instance without initialization.
func (vc _VoiceChatServiceClass) Alloc() VoiceChatService {
	rv := objc.Send[VoiceChatService](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VoiceChatServiceClass) New() VoiceChatService {
	rv := objc.Send[VoiceChatService](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VoiceChatService) Init() VoiceChatService {
	rv := objc.Send[VoiceChatService](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VoiceChatService) Autorelease() VoiceChatService {
	rv := objc.Send[VoiceChatService](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVoiceChatService creates a new VoiceChatService instance.
func NewVoiceChatService() VoiceChatService {
	return getVoiceChatServiceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VoiceChatService */
// The class allows your application to connect two iOS devices into a voice chat.
//
// Before you can use voice chat, your application must configure an audio session that allows for both play and recording ( ). For more information on audio sessions, see . The voice chat service uses a implemented by your application to find and connect to other participants. Each participant in the chat is identified by a unique string. The client provides a participant identifier for the local user and translates other participant identifiers into connections to other users. The format and mechanism used to translate participant identifiers into network connections is defined by the client. Your application can configure the voice chat service to control the volume level of both local and remote participants and to detect when someone is speaking. To use the voice chat service, your application retrieves the default service and attaches a client to it, then either connects to another participant or waits for them to start a connection.


// The class allows your application to connect two iOS devices into a voice chat.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatService
type VoiceChatService struct {
	objectivec.Object
}

// VoiceChatServiceFrom constructs a [VoiceChatService] from an unsafe.Pointer.
//
// The class allows your application to connect two iOS devices into a voice chat.
func VoiceChatServiceFrom(ptr unsafe.Pointer) VoiceChatService {
	return VoiceChatService{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VoiceChatService *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VoiceChatService */

// Retrieves the singleton chat service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatService/default()
func (vc _VoiceChatServiceClass) DefaultVoiceChatService() IVoiceChatService {
	rv := objc.Send[VoiceChatService](objc.ID(vc.class), objc.Sel("defaultVoiceChatService"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultVoiceChatService) */


// Returns whether voice chat is allowed to be used on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatService/isVoIPAllowed()
func (vc _VoiceChatServiceClass) IsVoIPAllowed() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("isVoIPAllowed"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsVoIPAllowed) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VoiceChatService */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VoiceChatService */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VoiceChatService */

// Allows recording (input) and playback (output) of audio, such as for a VOIP (voice over IP) app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/kAudioSessionCategory_PlayAndRecord
func (v_ VoiceChatService) KAudioSessionCategory_PlayAndRecord() int {
	rv := objc.Send[int](v_.ID, objc.Sel("kAudioSessionCategory_PlayAndRecord"))
	return rv
}/* debug [instance_properties/getter]: kAudioSessionCategory_PlayAndRecord */


// Allows recording (input) and playback (output) of audio, such as for a VOIP (voice over IP) app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/kAudioSessionCategory_PlayAndRecord
func (v_ VoiceChatService) SetKAudioSessionCategory_PlayAndRecord(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setKAudioSessionCategory_PlayAndRecord:"), value)
}/* debug [instance_properties/setter]: kAudioSessionCategory_PlayAndRecord */


// A Boolean value that indicates whether the microphone’s sound level is being monitored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/isinputmeteringenabled
func (v_ VoiceChatService) IsInputMeteringEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isInputMeteringEnabled"))
	return rv
}/* debug [instance_properties/getter]: isInputMeteringEnabled */


// A Boolean value that indicates whether the microphone’s sound level is being monitored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/isinputmeteringenabled
func (v_ VoiceChatService) SetIsInputMeteringEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsInputMeteringEnabled:"), value)
}/* debug [instance_properties/setter]: isInputMeteringEnabled */


// A Boolean value that determines whether the user’s microphone is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/ismicrophonemuted
func (v_ VoiceChatService) IsMicrophoneMuted() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isMicrophoneMuted"))
	return rv
}/* debug [instance_properties/getter]: isMicrophoneMuted */


// A Boolean value that determines whether the user’s microphone is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/ismicrophonemuted
func (v_ VoiceChatService) SetIsMicrophoneMuted(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsMicrophoneMuted:"), value)
}/* debug [instance_properties/setter]: isMicrophoneMuted */


// A Boolean value that indicates whether the voice level of remote participants is monitored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/isoutputmeteringenabled
func (v_ VoiceChatService) IsOutputMeteringEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isOutputMeteringEnabled"))
	return rv
}/* debug [instance_properties/getter]: isOutputMeteringEnabled */


// A Boolean value that indicates whether the voice level of remote participants is monitored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/isoutputmeteringenabled
func (v_ VoiceChatService) SetIsOutputMeteringEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsOutputMeteringEnabled:"), value)
}/* debug [instance_properties/setter]: isOutputMeteringEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKVoiceChatService */


