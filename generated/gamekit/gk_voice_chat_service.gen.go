// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// The class allows your application to connect two iOS devices into a voice chat.
//
// Before you can use voice chat, your application must configure an audio session that allows for both play and recording ( ). For more information on audio sessions, see . The voice chat service uses a implemented by your application to find and connect to other participants. Each participant in the chat is identified by a unique string. The client provides a participant identifier for the local user and translates other participant identifiers into connections to other users. The format and mechanism used to translate participant identifiers into network connections is defined by the client. Your application can configure the voice chat service to control the volume level of both local and remote participants and to detect when someone is speaking. To use the voice chat service, your application retrieves the default service and attaches a client to it, then either connects to another participant or waits for them to start a connection.
//
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




