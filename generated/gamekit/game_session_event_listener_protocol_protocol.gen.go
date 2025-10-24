// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PGameSessionEventListener is the GKGameSessionEventListener protocol interface.
//
// An event listener that handles game session events.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 10.0+ (Deprecated in 12.0)
//   - iPadOS 10.0+ (Deprecated in 12.0)
//   - macOS 10.12+ (Deprecated in 10.14)
//   - tvOS 10.0+ (Deprecated in 12.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKGameSessionEventListener
type PGameSessionEventListener interface {
	// Optional methods
	SessionDidAddPlayer(session IGKGameSession, player IGKCloudPlayer)
	HasSessionDidAddPlayer() bool
	SessionDidReceiveDataFromPlayer(session IGKGameSession, data objc.IObject /* cross-framework: NSData */, player IGKCloudPlayer)
	HasSessionDidReceiveDataFromPlayer() bool
	SessionDidReceiveMessageWithDataFromPlayer(session IGKGameSession, message objc.IObject /* cross-framework: NSString */, data objc.IObject /* cross-framework: NSData */, player IGKCloudPlayer)
	HasSessionDidReceiveMessageWithDataFromPlayer() bool
	SessionDidRemovePlayer(session IGKGameSession, player IGKCloudPlayer)
	HasSessionDidRemovePlayer() bool
	SessionPlayerDidChangeConnectionState(session IGKGameSession, player IGKCloudPlayer, newState ConnectionState)
	HasSessionPlayerDidChangeConnectionState() bool
	SessionPlayerDidSaveData(session IGKGameSession, player IGKCloudPlayer, data objc.IObject /* cross-framework: NSData */)
	HasSessionPlayerDidSaveData() bool
}
