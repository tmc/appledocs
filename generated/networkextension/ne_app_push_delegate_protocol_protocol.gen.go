// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PNEAppPushDelegate is the NEAppPushDelegate protocol interface.
//
// A protocol that defines how an app push manager instance interacts with the framework.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.networkextension/documentation/NetworkExtension/NEAppPushDelegate
type PNEAppPushDelegate interface {
	// Required methods
	AppPushManagerDidReceiveIncomingCallWithUserInfo(manager INEAppPushManager, userInfo objc.IObject /* cross-framework: NSDictionary */)/* debug [protocol_interface/required_method]: AppPushManagerDidReceiveIncomingCallWithUserInfo */
}

// NEAppPushDelegate is a delegate implementation builder for the PNEAppPushDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type NEAppPushDelegate struct {
	_AppPushManagerDidReceiveIncomingCallWithUserInfo func(manager INEAppPushManager, userInfo objc.IObject /* cross-framework: NSDictionary */)
}

// SetAppPushManagerDidReceiveIncomingCallWithUserInfo sets the handler for the AppPushManagerDidReceiveIncomingCallWithUserInfo delegate method.
//
// A delegate method that the framework invokes when the provider reports an incoming call.
func (d *NEAppPushDelegate) SetAppPushManagerDidReceiveIncomingCallWithUserInfo(f func(manager INEAppPushManager, userInfo objc.IObject /* cross-framework: NSDictionary */)) {
	d._AppPushManagerDidReceiveIncomingCallWithUserInfo = f
}

// AppPushManagerDidReceiveIncomingCallWithUserInfo implements the PNEAppPushDelegate interface.
func (d *NEAppPushDelegate) AppPushManagerDidReceiveIncomingCallWithUserInfo(manager INEAppPushManager, userInfo objc.IObject /* cross-framework: NSDictionary */) {
	if d._AppPushManagerDidReceiveIncomingCallWithUserInfo != nil {
		d._AppPushManagerDidReceiveIncomingCallWithUserInfo(manager, userInfo)
	}
}

// HasAppPushManagerDidReceiveIncomingCallWithUserInfo returns true if a handler for AppPushManagerDidReceiveIncomingCallWithUserInfo has been set.
func (d *NEAppPushDelegate) HasAppPushManagerDidReceiveIncomingCallWithUserInfo() bool {
	return d._AppPushManagerDidReceiveIncomingCallWithUserInfo != nil
}
