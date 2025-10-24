//go:build darwin && ios

// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for VoiceChatService


// iOS-only properties

// An object that the voice chat service uses to communicate with remote participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatService/client
func (v_ VoiceChatService) Client() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("client"))
	return rv
}
func (v_ VoiceChatService) SetClient(value unsafe.Pointer) {
	v_.ID.Send(objc.RegisterName("setClient:"), value)
}

// The volume, in decibels (db), being received by the microphone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatService/inputMeterLevel
func (v_ VoiceChatService) InputMeterLevel() float32 {
	rv := objc.Send[float32](v_.ID, objc.Sel("inputMeterLevel"))
	return rv
}

// A Boolean value that indicates whether the microphone’s sound level is being monitored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatService/isInputMeteringEnabled
func (v_ VoiceChatService) InputMeteringEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("inputMeteringEnabled"))
	return rv
}
func (v_ VoiceChatService) SetInputMeteringEnabled(value bool) {
	v_.ID.Send(objc.RegisterName("setInputMeteringEnabled:"), value)
}

// A Boolean value that determines whether the user’s microphone is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatService/isMicrophoneMuted
func (v_ VoiceChatService) MicrophoneMuted() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("microphoneMuted"))
	return rv
}
func (v_ VoiceChatService) SetMicrophoneMuted(value bool) {
	v_.ID.Send(objc.RegisterName("setMicrophoneMuted:"), value)
}

// A Boolean value that indicates whether the voice level of remote participants is monitored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatService/isOutputMeteringEnabled
func (v_ VoiceChatService) OutputMeteringEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("outputMeteringEnabled"))
	return rv
}
func (v_ VoiceChatService) SetOutputMeteringEnabled(value bool) {
	v_.ID.Send(objc.RegisterName("setOutputMeteringEnabled:"), value)
}

// The volume, in decibels (db), being received from all other participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatService/outputMeterLevel
func (v_ VoiceChatService) OutputMeterLevel() float32 {
	rv := objc.Send[float32](v_.ID, objc.Sel("outputMeterLevel"))
	return rv
}

// A float that scales the volume of all remote participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatService/remoteParticipantVolume
func (v_ VoiceChatService) RemoteParticipantVolume() float32 {
	rv := objc.Send[float32](v_.ID, objc.Sel("remoteParticipantVolume"))
	return rv
}
func (v_ VoiceChatService) SetRemoteParticipantVolume(value float32) {
	v_.ID.Send(objc.RegisterName("setRemoteParticipantVolume:"), value)
}







