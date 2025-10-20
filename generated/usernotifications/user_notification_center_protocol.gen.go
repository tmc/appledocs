// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import "github.com/ebitengine/purego/objc"

// userNotificationCenterProtocol is the userNotificationCenter: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.14+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+
//
// Use this protocol when registering custom classes that conform to userNotificationCenter:.
var userNotificationCenterProtocol *objc.Protocol

func init() {
	userNotificationCenterProtocol = objc.GetProtocol("userNotificationCenter:")
}

