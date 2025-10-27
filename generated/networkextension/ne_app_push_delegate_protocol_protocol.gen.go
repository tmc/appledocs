// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (

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
	AppPushManagerDidReceiveIncomingCallWithUserInfo(manager INEAppPushManager, userInfo foundation.foundation.INSDictionary)
}
