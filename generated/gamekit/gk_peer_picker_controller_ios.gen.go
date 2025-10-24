//go:build darwin && ios

// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for PeerPickerController


// iOS-only properties

// A mask that determines the types of connections a dialog presents to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerController/connectionTypesMask
func (p_ PeerPickerController) ConnectionTypesMask() PeerPickerConnectionType {
	rv := objc.Send[PeerPickerConnectionType](p_.ID, objc.Sel("connectionTypesMask"))
	return rv
}
func (p_ PeerPickerController) SetConnectionTypesMask(value PeerPickerConnectionType) {
	p_.ID.Send(objc.RegisterName("setConnectionTypesMask:"), value)
}

// The delegate of the peer picker controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerController/delegate
func (p_ PeerPickerController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}
func (p_ PeerPickerController) SetDelegate(value unsafe.Pointer) {
	p_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// A Boolean value that indicates whether the picker dialog is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerController/isVisible
func (p_ PeerPickerController) Visible() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("visible"))
	return rv
}





