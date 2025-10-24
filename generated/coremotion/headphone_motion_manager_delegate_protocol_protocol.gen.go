// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PHeadphoneMotionManagerDelegate is the CMHeadphoneMotionManagerDelegate protocol interface.
//
// A set of methods that defines an interface for connecting and disconnecting headphones.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 14.0+
//   - watchOS 7.0+
//
// See: doc://com.apple.coremotion/documentation/CoreMotion/CMHeadphoneMotionManagerDelegate
type PHeadphoneMotionManagerDelegate interface {
	// Optional methods
	HeadphoneMotionManagerDidConnect(manager ICMHeadphoneMotionManager)
	HasHeadphoneMotionManagerDidConnect() bool
	HeadphoneMotionManagerDidDisconnect(manager ICMHeadphoneMotionManager)
	HasHeadphoneMotionManagerDidDisconnect() bool
}

// HeadphoneMotionManagerDelegate is a delegate implementation builder for the PHeadphoneMotionManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type HeadphoneMotionManagerDelegate struct {
	_HeadphoneMotionManagerDidConnect func(manager ICMHeadphoneMotionManager)
	_HeadphoneMotionManagerDidDisconnect func(manager ICMHeadphoneMotionManager)
}

// SetHeadphoneMotionManagerDidConnect sets the handler for the HeadphoneMotionManagerDidConnect delegate method.
//
// Performs a callback to the delegate after you connect headphones.
func (d *HeadphoneMotionManagerDelegate) SetHeadphoneMotionManagerDidConnect(f func(manager ICMHeadphoneMotionManager)) {
	d._HeadphoneMotionManagerDidConnect = f
}

// SetHeadphoneMotionManagerDidDisconnect sets the handler for the HeadphoneMotionManagerDidDisconnect delegate method.
//
// Performs a callback to the delegate after you disconnect headphones.
func (d *HeadphoneMotionManagerDelegate) SetHeadphoneMotionManagerDidDisconnect(f func(manager ICMHeadphoneMotionManager)) {
	d._HeadphoneMotionManagerDidDisconnect = f
}

// HeadphoneMotionManagerDidConnect implements the PHeadphoneMotionManagerDelegate interface.
func (d *HeadphoneMotionManagerDelegate) HeadphoneMotionManagerDidConnect(manager ICMHeadphoneMotionManager) {
	if d._HeadphoneMotionManagerDidConnect != nil {
		d._HeadphoneMotionManagerDidConnect(manager)
	}
}

// HasHeadphoneMotionManagerDidConnect returns true if a handler for HeadphoneMotionManagerDidConnect has been set.
func (d *HeadphoneMotionManagerDelegate) HasHeadphoneMotionManagerDidConnect() bool {
	return d._HeadphoneMotionManagerDidConnect != nil
}

// HeadphoneMotionManagerDidDisconnect implements the PHeadphoneMotionManagerDelegate interface.
func (d *HeadphoneMotionManagerDelegate) HeadphoneMotionManagerDidDisconnect(manager ICMHeadphoneMotionManager) {
	if d._HeadphoneMotionManagerDidDisconnect != nil {
		d._HeadphoneMotionManagerDidDisconnect(manager)
	}
}

// HasHeadphoneMotionManagerDidDisconnect returns true if a handler for HeadphoneMotionManagerDidDisconnect has been set.
func (d *HeadphoneMotionManagerDelegate) HasHeadphoneMotionManagerDidDisconnect() bool {
	return d._HeadphoneMotionManagerDidDisconnect != nil
}
