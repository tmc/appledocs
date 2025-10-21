// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import "github.com/ebitengine/purego/objc"

// UNNotificationContentProvidingProtocol is the UNNotificationContentProviding protocol.
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+
//
// Use this protocol when registering custom classes that conform to UNNotificationContentProviding.
var UNNotificationContentProvidingProtocol *objc.Protocol

func init() {
	UNNotificationContentProvidingProtocol = objc.GetProtocol("UNNotificationContentProviding")
}
