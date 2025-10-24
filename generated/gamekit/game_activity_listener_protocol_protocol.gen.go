// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"unsafe"
)

// PGameActivityListener is the GKGameActivityListener protocol interface.
//
// An object that responds to activity events.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKGameActivityListener
type PGameActivityListener interface {
	// Optional methods
	PlayerWantsToPlayGameActivityCompletionHandler(player IGKPlayer, activity IGKGameActivity, completionHandler unsafe.Pointer)
	HasPlayerWantsToPlayGameActivityCompletionHandler() bool
}
