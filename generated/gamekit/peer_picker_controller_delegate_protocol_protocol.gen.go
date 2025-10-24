// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PPeerPickerControllerDelegate is the GKPeerPickerControllerDelegate protocol interface.
//
// The   protocol is implemented on an object to customize the behavior of a   object. The delegate is called by the peer picker to create a session object and to respond as the session is configured by the controller.
//
// Availability:
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKPeerPickerControllerDelegate
type PPeerPickerControllerDelegate interface {
	// Optional methods
	PeerPickerControllerDidConnectPeerToSession(picker IGKPeerPickerController, peerID objc.IObject /* cross-framework: NSString */, session IGKSession)
	HasPeerPickerControllerDidConnectPeerToSession() bool
	PeerPickerControllerDidSelectConnectionType(picker IGKPeerPickerController, type_ PeerPickerConnectionType)
	HasPeerPickerControllerDidSelectConnectionType() bool
	PeerPickerControllerSessionForConnectionType(picker IGKPeerPickerController, type_ PeerPickerConnectionType) Session
	HasPeerPickerControllerSessionForConnectionType() bool
	PeerPickerControllerDidCancel(picker IGKPeerPickerController)
	HasPeerPickerControllerDidCancel() bool
}

// PeerPickerControllerDelegate is a delegate implementation builder for the PPeerPickerControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PeerPickerControllerDelegate struct {
	_PeerPickerControllerDidConnectPeerToSession func(picker IGKPeerPickerController, peerID objc.IObject /* cross-framework: NSString */, session IGKSession)
	_PeerPickerControllerDidSelectConnectionType func(picker IGKPeerPickerController, type_ PeerPickerConnectionType)
	_PeerPickerControllerSessionForConnectionType func(picker IGKPeerPickerController, type_ PeerPickerConnectionType) Session
	_PeerPickerControllerDidCancel func(picker IGKPeerPickerController)
}

// SetPeerPickerControllerDidConnectPeerToSession sets the handler for the PeerPickerControllerDidConnectPeerToSession delegate method.
//
// Tells the delegate that the controller connected a peer to the session.
func (d *PeerPickerControllerDelegate) SetPeerPickerControllerDidConnectPeerToSession(f func(picker IGKPeerPickerController, peerID objc.IObject /* cross-framework: NSString */, session IGKSession)) {
	d._PeerPickerControllerDidConnectPeerToSession = f
}

// SetPeerPickerControllerDidSelectConnectionType sets the handler for the PeerPickerControllerDidSelectConnectionType delegate method.
//
// Tells the delegate that the user selected a connection type.
func (d *PeerPickerControllerDelegate) SetPeerPickerControllerDidSelectConnectionType(f func(picker IGKPeerPickerController, type_ PeerPickerConnectionType)) {
	d._PeerPickerControllerDidSelectConnectionType = f
}

// SetPeerPickerControllerSessionForConnectionType sets the handler for the PeerPickerControllerSessionForConnectionType delegate method.
//
// Asks the delegate to return a session for the specified connection type.
func (d *PeerPickerControllerDelegate) SetPeerPickerControllerSessionForConnectionType(f func(picker IGKPeerPickerController, type_ PeerPickerConnectionType) Session) {
	d._PeerPickerControllerSessionForConnectionType = f
}

// SetPeerPickerControllerDidCancel sets the handler for the PeerPickerControllerDidCancel delegate method.
//
// Tells the delegate that the user canceled the connection attempt.
func (d *PeerPickerControllerDelegate) SetPeerPickerControllerDidCancel(f func(picker IGKPeerPickerController)) {
	d._PeerPickerControllerDidCancel = f
}

// PeerPickerControllerDidConnectPeerToSession implements the PPeerPickerControllerDelegate interface.
func (d *PeerPickerControllerDelegate) PeerPickerControllerDidConnectPeerToSession(picker IGKPeerPickerController, peerID objc.IObject /* cross-framework: NSString */, session IGKSession) {
	if d._PeerPickerControllerDidConnectPeerToSession != nil {
		d._PeerPickerControllerDidConnectPeerToSession(picker, peerID, session)
	}
}

// HasPeerPickerControllerDidConnectPeerToSession returns true if a handler for PeerPickerControllerDidConnectPeerToSession has been set.
func (d *PeerPickerControllerDelegate) HasPeerPickerControllerDidConnectPeerToSession() bool {
	return d._PeerPickerControllerDidConnectPeerToSession != nil
}

// PeerPickerControllerDidSelectConnectionType implements the PPeerPickerControllerDelegate interface.
func (d *PeerPickerControllerDelegate) PeerPickerControllerDidSelectConnectionType(picker IGKPeerPickerController, type_ PeerPickerConnectionType) {
	if d._PeerPickerControllerDidSelectConnectionType != nil {
		d._PeerPickerControllerDidSelectConnectionType(picker, type_)
	}
}

// HasPeerPickerControllerDidSelectConnectionType returns true if a handler for PeerPickerControllerDidSelectConnectionType has been set.
func (d *PeerPickerControllerDelegate) HasPeerPickerControllerDidSelectConnectionType() bool {
	return d._PeerPickerControllerDidSelectConnectionType != nil
}

// PeerPickerControllerSessionForConnectionType implements the PPeerPickerControllerDelegate interface.
func (d *PeerPickerControllerDelegate) PeerPickerControllerSessionForConnectionType(picker IGKPeerPickerController, type_ PeerPickerConnectionType) Session {
	if d._PeerPickerControllerSessionForConnectionType != nil {
		return d._PeerPickerControllerSessionForConnectionType(picker, type_)
	}
	var zero Session
	return zero
}

// HasPeerPickerControllerSessionForConnectionType returns true if a handler for PeerPickerControllerSessionForConnectionType has been set.
func (d *PeerPickerControllerDelegate) HasPeerPickerControllerSessionForConnectionType() bool {
	return d._PeerPickerControllerSessionForConnectionType != nil
}

// PeerPickerControllerDidCancel implements the PPeerPickerControllerDelegate interface.
func (d *PeerPickerControllerDelegate) PeerPickerControllerDidCancel(picker IGKPeerPickerController) {
	if d._PeerPickerControllerDidCancel != nil {
		d._PeerPickerControllerDidCancel(picker)
	}
}

// HasPeerPickerControllerDidCancel returns true if a handler for PeerPickerControllerDidCancel has been set.
func (d *PeerPickerControllerDelegate) HasPeerPickerControllerDidCancel() bool {
	return d._PeerPickerControllerDidCancel != nil
}
