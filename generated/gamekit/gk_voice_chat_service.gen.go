// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [VoiceChatService] class.
type IVoiceChatService interface {
	objectivec.IObject
	KAudioSessionCategory_PlayAndRecord() int
	SetKAudioSessionCategory_PlayAndRecord(value int)
	Client() unsafe.Pointer
	SetClient(value unsafe.Pointer)
	InputMeterLevel() float32
	SetInputMeterLevel(value float32)
	IsInputMeteringEnabled() bool
	SetIsInputMeteringEnabled(value bool)
	IsMicrophoneMuted() bool
	SetIsMicrophoneMuted(value bool)
	IsOutputMeteringEnabled() bool
	SetIsOutputMeteringEnabled(value bool)
	OutputMeterLevel() float32
	SetOutputMeterLevel(value float32)
	RemoteParticipantVolume() float32
	SetRemoteParticipantVolume(value float32)
}

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

// Alloc allocates a new instance without initialization.
func (vc _VoiceChatServiceClass) Alloc() VoiceChatService {
	rv := objc.Send[VoiceChatService](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Allows recording (input) and playback (output) of audio, such as for a VOIP (voice over IP) app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/kAudioSessionCategory_PlayAndRecord
func (v_ VoiceChatService) KAudioSessionCategory_PlayAndRecord() int {
	rv := objc.Send[int](v_.ID, objc.Sel("kAudioSessionCategory_PlayAndRecord"))
	return rv
}


// Allows recording (input) and playback (output) of audio, such as for a VOIP (voice over IP) app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/kAudioSessionCategory_PlayAndRecord
func (v_ VoiceChatService) SetKAudioSessionCategory_PlayAndRecord(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setKAudioSessionCategory_PlayAndRecord:"), value)
}


// An object that the voice chat service uses to communicate with remote participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/client
func (v_ VoiceChatService) Client() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("client"))
	return rv
}


// An object that the voice chat service uses to communicate with remote participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/client
func (v_ VoiceChatService) SetClient(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setClient:"), value)
}


// The volume, in decibels (db), being received by the microphone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/inputmeterlevel
func (v_ VoiceChatService) InputMeterLevel() float32 {
	rv := objc.Send[float32](v_.ID, objc.Sel("inputMeterLevel"))
	return rv
}


// The volume, in decibels (db), being received by the microphone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/inputmeterlevel
func (v_ VoiceChatService) SetInputMeterLevel(value float32) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setInputMeterLevel:"), value)
}


// A Boolean value that indicates whether the microphone’s sound level is being monitored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/isinputmeteringenabled
func (v_ VoiceChatService) IsInputMeteringEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isInputMeteringEnabled"))
	return rv
}


// A Boolean value that indicates whether the microphone’s sound level is being monitored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/isinputmeteringenabled
func (v_ VoiceChatService) SetIsInputMeteringEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsInputMeteringEnabled:"), value)
}


// A Boolean value that determines whether the user’s microphone is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/ismicrophonemuted
func (v_ VoiceChatService) IsMicrophoneMuted() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isMicrophoneMuted"))
	return rv
}


// A Boolean value that determines whether the user’s microphone is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/ismicrophonemuted
func (v_ VoiceChatService) SetIsMicrophoneMuted(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsMicrophoneMuted:"), value)
}


// A Boolean value that indicates whether the voice level of remote participants is monitored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/isoutputmeteringenabled
func (v_ VoiceChatService) IsOutputMeteringEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isOutputMeteringEnabled"))
	return rv
}


// A Boolean value that indicates whether the voice level of remote participants is monitored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/isoutputmeteringenabled
func (v_ VoiceChatService) SetIsOutputMeteringEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsOutputMeteringEnabled:"), value)
}


// The volume, in decibels (db), being received from all other participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/outputmeterlevel
func (v_ VoiceChatService) OutputMeterLevel() float32 {
	rv := objc.Send[float32](v_.ID, objc.Sel("outputMeterLevel"))
	return rv
}


// The volume, in decibels (db), being received from all other participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/outputmeterlevel
func (v_ VoiceChatService) SetOutputMeterLevel(value float32) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setOutputMeterLevel:"), value)
}


// A float that scales the volume of all remote participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/remoteparticipantvolume
func (v_ VoiceChatService) RemoteParticipantVolume() float32 {
	rv := objc.Send[float32](v_.ID, objc.Sel("remoteParticipantVolume"))
	return rv
}


// A float that scales the volume of all remote participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkvoicechatservice/remoteparticipantvolume
func (v_ VoiceChatService) SetRemoteParticipantVolume(value float32) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setRemoteParticipantVolume:"), value)
}




