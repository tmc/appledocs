// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PVoiceChatClient is the GKVoiceChatClient protocol interface.
//
// The   protocol is implemented to control the behavior of the   object. The voice chat client has a number of responsibilities:
//
// Availability:
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 3.0+ (Deprecated in 3.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKVoiceChatClient
type PVoiceChatClient interface {
	// Required methods
	ParticipantID() foundation.String/* debug [protocol_interface/required_method]: ParticipantID */
	VoiceChatServiceSendDataToParticipantID(voiceChatService IGKVoiceChatService, data objc.IObject /* cross-framework: NSData */, participantID objc.IObject /* cross-framework: NSString */)/* debug [protocol_interface/required_method]: VoiceChatServiceSendDataToParticipantID */
	// Optional methods
	VoiceChatServiceDidNotStartWithParticipantIDError(voiceChatService IGKVoiceChatService, participantID objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: Error */)
	HasVoiceChatServiceDidNotStartWithParticipantIDError() bool
	VoiceChatServiceDidReceiveInvitationFromParticipantIDCallID(voiceChatService IGKVoiceChatService, participantID objc.IObject /* cross-framework: NSString */, callID int)
	HasVoiceChatServiceDidReceiveInvitationFromParticipantIDCallID() bool
	VoiceChatServiceDidStartWithParticipantID(voiceChatService IGKVoiceChatService, participantID objc.IObject /* cross-framework: NSString */)
	HasVoiceChatServiceDidStartWithParticipantID() bool
	VoiceChatServiceDidStopWithParticipantIDError(voiceChatService IGKVoiceChatService, participantID objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: Error */)
	HasVoiceChatServiceDidStopWithParticipantIDError() bool
	VoiceChatServiceSendRealTimeDataToParticipantID(voiceChatService IGKVoiceChatService, data objc.IObject /* cross-framework: NSData */, participantID objc.IObject /* cross-framework: NSString */)
	HasVoiceChatServiceSendRealTimeDataToParticipantID() bool
}
