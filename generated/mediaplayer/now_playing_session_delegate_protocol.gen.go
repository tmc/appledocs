// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import "github.com/ebitengine/purego/objc"

// NowPlayingSessionDelegateProtocol is the MPNowPlayingSessionDelegate protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to MPNowPlayingSessionDelegate.
var NowPlayingSessionDelegateProtocol *objc.Protocol

func init() {
	NowPlayingSessionDelegateProtocol = objc.GetProtocol("MPNowPlayingSessionDelegate")
}

