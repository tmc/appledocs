//go:build darwin && ios

// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for LocalPlayer


// Presents a view controller with a Messages sheet for the player to request friends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/presentFriendRequestCreator(from:)-7j1kn
func (l_ LocalPlayer) PresentFriendRequestCreatorFromViewControllerError(viewController objc.IObject /* cross-framework: ViewController */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("presentFriendRequestCreatorFromViewController:error:"), viewController, error_)
	return rv
}

// iOS-only properties





