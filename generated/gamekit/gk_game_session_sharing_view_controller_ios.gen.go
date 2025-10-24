//go:build darwin && ios

// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for GameSessionSharingViewController


// iOS-only properties

// The delegate for the sharing view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionSharingViewController/delegate
func (g_ GameSessionSharingViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("delegate"))
	return rv
}
func (g_ GameSessionSharingViewController) SetDelegate(value unsafe.Pointer) {
	g_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The game session associated with the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionSharingViewController/session
func (g_ GameSessionSharingViewController) Session() IGKGameSession {
	rv := objc.Send[GameSession](g_.ID, objc.Sel("session"))
	return rv
}




