// Code generated from Apple documentation for UserNotificationsUI. DO NOT EDIT.

package usernotificationsui

import "github.com/ebitengine/purego/objc"

// didReceiveNotificationResponseProtocol is the didReceiveNotificationResponse: protocol.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to didReceiveNotificationResponse:.
var didReceiveNotificationResponseProtocol *objc.Protocol

func init() {
	didReceiveNotificationResponseProtocol = objc.GetProtocol("didReceiveNotificationResponse:")
}

