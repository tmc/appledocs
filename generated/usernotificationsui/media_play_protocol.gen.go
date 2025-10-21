// Code generated from Apple documentation for UserNotificationsUI. DO NOT EDIT.

package usernotificationsui

import "github.com/ebitengine/purego/objc"

// mediaPlayProtocol is the mediaPlay protocol.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to mediaPlay.
var mediaPlayProtocol *objc.Protocol

func init() {
	mediaPlayProtocol = objc.GetProtocol("mediaPlay")
}
